package types_test

import (
	"github.com/cosmos/osmosis-sdk/simapp"
)

var (
	app         = simapp.Setup(false)
	appCodec, _ = simapp.MakeCodecs()
)
