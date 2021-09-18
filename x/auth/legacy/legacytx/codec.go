package legacytx

import (
	"github.com/cosmos/osmosis-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "osmosis-sdk/StdTx", nil)
}
