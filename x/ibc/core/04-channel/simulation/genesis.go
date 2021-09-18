package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/osmosis-sdk/types/simulation"
	"github.com/cosmos/osmosis-sdk/x/ibc/core/04-channel/types"
)

// GenChannelGenesis returns the default channel genesis state.
func GenChannelGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
