package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/osmosis-sdk/codec"
	cryptoAmino "github.com/cosmos/osmosis-sdk/crypto/codec"
	"github.com/cosmos/osmosis-sdk/testutil/testdata"
	sdk "github.com/cosmos/osmosis-sdk/types"
	"github.com/cosmos/osmosis-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/osmosis-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "osmosis-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
