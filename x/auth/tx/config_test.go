package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/osmosis-sdk/codec"
	codectypes "github.com/cosmos/osmosis-sdk/codec/types"
	"github.com/cosmos/osmosis-sdk/std"
	"github.com/cosmos/osmosis-sdk/testutil/testdata"
	sdk "github.com/cosmos/osmosis-sdk/types"
	"github.com/cosmos/osmosis-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
