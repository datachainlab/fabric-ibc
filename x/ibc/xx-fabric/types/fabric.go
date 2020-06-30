package types

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
)

func GetLocalDeserializer() msp.IdentityDeserializer {
	// TODO fix msp config
	// setup the MSP manager so that we can sign/verify
	mgr, err := LoadMSPs(DefaultConfig())
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	return mgr
}

func GetPolicyEvaluator(policyBytes []byte) (policies.Policy, error) {
	var ap peer.ApplicationPolicy
	if err := proto.Unmarshal(policyBytes, &ap); err != nil {
		return nil, err
	}
	sigp := ap.Type.(*peer.ApplicationPolicy_SignaturePolicy)
	pp := cauthdsl.EnvelopeBasedPolicyProvider{Deserializer: GetLocalDeserializer()}
	return pp.NewPolicy(sigp.SignaturePolicy)
}

func makeSignedDataList(pr *Proof) []*protoutil.SignedData {
	var sigSet []*protoutil.SignedData
	for i := 0; i < len(pr.Signatures); i++ {
		msg := make([]byte, len(pr.Proposal)+len(pr.Identities[i]))
		copy(msg[:len(pr.Proposal)], pr.Proposal)
		copy(msg[len(pr.Proposal):], pr.Identities[i])

		sigSet = append(
			sigSet,
			&protoutil.SignedData{
				Data:      msg,
				Identity:  pr.Identities[i],
				Signature: pr.Signatures[i],
			},
		)
	}
	return sigSet
}

func GetTxReadWriteSetFromProposalResponsePayload(proposal []byte) (*peer.ChaincodeID, *rwset.TxReadWriteSet, error) {
	var payload peer.ProposalResponsePayload
	if err := proto.Unmarshal(proposal, &payload); err != nil {
		return nil, nil, err
	}
	var cact peer.ChaincodeAction
	if err := proto.Unmarshal(payload.Extension, &cact); err != nil {
		return nil, nil, err
	}
	var result rwset.TxReadWriteSet
	if err := proto.Unmarshal(cact.Results, &result); err != nil {
		return nil, nil, err
	}
	return cact.GetChaincodeId(), &result, nil
}

func EnsureWriteSetIncludesCommitment(set []*rwset.NsReadWriteSet, nsIdx, rwsIdx uint32, targetKey string, expectValue []byte) (bool, error) {
	ws := set[nsIdx].Rwset
	var rws kvrwset.KVRWSet
	if err := proto.Unmarshal(ws, &rws); err != nil {
		return false, err
	}
	if len(rws.Writes) <= int(rwsIdx) {
		return false, fmt.Errorf("not found index '%v'(length=%v)", int(rwsIdx), len(rws.Writes))
	}

	w := rws.Writes[rwsIdx]
	if targetKey != w.Key {
		return false, fmt.Errorf("unexpected key '%v'(expected=%v)", w.Key, targetKey)
	}
	if bytes.Equal(w.Value, expectValue) {
		return true, nil
	} else {
		return false, nil
	}
}

func VerifyEndorsement(ccID peer.ChaincodeID, policyBytes []byte, proof Proof, path string, value []byte) (bool, error) {
	sigSet := makeSignedDataList(&proof)
	policy, err := GetPolicyEvaluator(policyBytes)
	if err != nil {
		return false, err
	}
	if err := policy.EvaluateSignedData(sigSet); err != nil {
		return false, err
	}

	id, rwset, err := GetTxReadWriteSetFromProposalResponsePayload(proof.Proposal)
	if err != nil {
		return false, err
	}
	if !equalChaincodeID(ccID, *id) {
		return false, fmt.Errorf("unexpected chaincodID: %v", *id)
	}
	return EnsureWriteSetIncludesCommitment(rwset.GetNsRwset(), proof.NsIndex, proof.WriteSetIndex, path, value)
}

func equalChaincodeID(a, b peer.ChaincodeID) bool {
	if a.Name == b.Name && a.Path == b.Path && a.Version == b.Version {
		return true
	} else {
		return false
	}
}

func VerifyChaincodeHeader(clientState ClientState, h ChaincodeHeader) error {
	lastci := clientState.LastChaincodeInfo
	ok, err := VerifyEndorsement(clientState.LastChaincodeInfo.GetFabricChaincodeID(), lastci.EndorsementPolicy, h.Proof, MakeSequenceCommitmentEntryKey(h.Sequence.Value), h.Sequence.Bytes())
	if err != nil {
		return err
	} else if !ok {
		return errors.New("failed to verify the endorsement")
	}
	return nil
}

func VerifyChaincodeInfo(clientState ClientState, info ChaincodeInfo) error {
	// TODO implement
	return nil
}

func LoadMSPs(conf Config) (msp.MSPManager, error) {
	mgr := msp.NewMSPManager()
	msps := []msp.MSP{}
	// if in local, this would depend `peer.localMspType` config
	opts := msp.Options[msp.ProviderTypeToString(msp.FABRIC)]
	for _, id := range conf.MSPIDs {
		dir := filepath.Join(conf.MSPsDir, id)
		mspConf, err := msp.GetLocalMspConfig(dir, nil, id)
		if err != nil {
			return nil, err
		}
		m, err := msp.New(opts, factory.GetDefault())
		if err != nil {
			return nil, err
		}
		m.Setup(mspConf)
		msps = append(msps, m)
	}
	err := mgr.Setup(msps)
	if err != nil {
		return nil, err
	}
	return mgr, nil
}
