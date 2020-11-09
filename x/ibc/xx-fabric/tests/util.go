package tests

import (
	"path/filepath"

	"github.com/golang/protobuf/proto"
	msppb "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/msp"
)

func GetLocalMspConfig(mspsDir string, mspID string) (*msppb.MSPConfig, *factory.FactoryOpts, error) {
	bccspConf := getDefaultBccspConfig()
	mconf, err := getLocalMspConfig(mspsDir, mspID, bccspConf)
	return mconf, bccspConf, err
}

func GetLocalVerifyingMspConfig(mspsDir string, mspID string) (*msppb.MSPConfig, error) {
	mconf, _, err := GetLocalMspConfig(mspsDir, mspID)
	if err != nil {
		return nil, err
	}
	err = ToVerifyingConfig(mconf)
	if err != nil {
		return nil, err
	}
	return mconf, nil
}

func GetLocalMsp(mspsDir string, mspID string) (msp.MSP, error) {
	mspConf, bccspConf, err := GetLocalMspConfig(mspsDir, mspID)
	if err != nil {
		return nil, err
	}
	return SetupLocalMsp(mspConf, bccspConf)
}

func SetupLocalMsp(mspConf *msppb.MSPConfig, bccspConf *factory.FactoryOpts) (msp.MSP, error) {
	cryptoProvider, err := (&factory.SWFactory{}).Get(bccspConf)
	if err != nil {
		return nil, err
	}
	opts := msp.Options[msp.ProviderTypeToString(msp.FABRIC)]
	m, err := msp.New(opts, cryptoProvider)
	if err != nil {
		return nil, err
	}
	if err := m.Setup(mspConf); err != nil {
		return nil, err
	}
	return m, nil
}

func ToVerifyingConfig(mconf *msppb.MSPConfig) error {
	var conf msppb.FabricMSPConfig
	err := proto.Unmarshal(mconf.Config, &conf)
	if err != nil {
		return err
	}
	conf.SigningIdentity = nil
	confb, err := proto.Marshal(&conf)
	if err != nil {
		return err
	}
	mconf.Config = confb
	return nil
}

func getLocalMspConfig(mspsDir string, mspID string, bccspConfig *factory.FactoryOpts) (*msppb.MSPConfig, error) {
	dir := filepath.Join(mspsDir, mspID)
	return msp.GetLocalMspConfig(dir, bccspConfig, mspID)
}

func getDefaultBccspConfig() *factory.FactoryOpts {
	return factory.GetDefaultOpts()
}
