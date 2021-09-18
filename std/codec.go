package std

import (
	"github.com/cosmos/osmosis-sdk/codec"
	"github.com/cosmos/osmosis-sdk/codec/types"
	cryptocodec "github.com/cosmos/osmosis-sdk/crypto/codec"
	sdk "github.com/cosmos/osmosis-sdk/types"
	txtypes "github.com/cosmos/osmosis-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
