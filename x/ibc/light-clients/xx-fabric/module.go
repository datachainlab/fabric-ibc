package fabric

import (
	"github.com/datachainlab/fabric-ibc/x/ibc/light-clients/xx-fabric/types"
)

// Name returns the IBC client name
func Name() string {
	return types.ModuleName
}
