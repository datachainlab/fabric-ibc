package testing

import (
	"strings"

	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/datachainlab/fabric-ibc/chaincode"
	"github.com/datachainlab/fabric-ibc/commitment"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func queryEndorseCommitment(ctx contractapi.TransactionContextInterface, cc *chaincode.IBCChaincode, key []byte) (*commitment.CommitmentEntry, error) {
	k := string(key)

	if strings.Contains(k, string(host.KeyClientState())) {
		parts := strings.Split(k, "/")
		clientID := parts[1]

		return cc.EndorseClientState(ctx, clientID)
	}

	if strings.HasPrefix(k, string(host.KeyConnectionPrefix)) {
		parts := strings.Split(k, "/")
		connectionID := parts[1]

		return cc.EndorseConnectionState(ctx, connectionID)
	}

	if strings.Contains(k, "consensusStates/") {
		parts := strings.Split(k, "/")
		clientID := parts[1]
		height, err := clienttypes.ParseHeight(parts[3])
		if err != nil {
			return nil, err
		}
		return cc.EndorseConsensusStateCommitment(ctx, clientID, height.VersionHeight)
	}

	panic("not implemented error")
}