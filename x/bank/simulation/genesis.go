package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/osmosis-sdk/types"
	"github.com/cosmos/osmosis-sdk/types/module"
	"github.com/cosmos/osmosis-sdk/x/bank/types"
)

// RandomGenesisDefaultSendParam computes randomized allow all send transfers param for the bank module
func RandomGenesisDefaultSendParam(r *rand.Rand) bool {
	// 90% chance of transfers being enable or P(a) = 0.9 for success
	return r.Int63n(101) <= 90
}

// RandomGenesisSendParams randomized Parameters for the bank module
func RandomGenesisSendParams(r *rand.Rand) types.SendEnabledParams {
	params := types.DefaultParams()
	// 90% chance of transfers being DefaultSendEnabled=true or P(a) = 0.9 for success
	// 50% of the time add an additional denom specific record (P(b) = 0.475 = 0.5 * 0.95)
	if r.Int63n(101) <= 50 {
		// set send enabled 95% of the time
		bondEnabled := r.Int63n(101) <= 95
		params = params.SetSendEnabledParam(
			sdk.DefaultBondDenom,
			bondEnabled)
	}

	// overall probability of enabled for bond denom is 94.75% (P(a)+P(b) - P(a)*P(b))
	return params.SendEnabled
}

func GenesisCoins(bondDenom string, initialStake int64) sdk.Coins {
	coins := make(sdk.Coins, 4)
	for i := range coins {
		coins[i] = sdk.NewCoin(fmt.Sprintf("%s%02d", bondDenom, i), sdk.NewInt(initialStake))
	}
	coins = coins.Add(sdk.NewInt64Coin(bondDenom, initialStake))
	return coins
}

// RandomGenesisBalances returns a slice of account balances. Each account has
// a balance of simState.InitialStake for sdk.DefaultBondDenom.
func RandomGenesisBalances(simState *module.SimulationState) []types.Balance {
	genesisBalances := []types.Balance{}

	for _, acc := range simState.Accounts {
		genesisBalances = append(genesisBalances, types.Balance{
			Address: acc.Address.String(),
			Coins:   GenesisCoins(sdk.DefaultBondDenom, simState.InitialStake),
		})
	}

	return genesisBalances
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) {
	var sendEnabledParams types.SendEnabledParams
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.KeySendEnabled), &sendEnabledParams, simState.Rand,
		func(r *rand.Rand) { sendEnabledParams = RandomGenesisSendParams(r) },
	)

	var defaultSendEnabledParam bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.KeyDefaultSendEnabled), &defaultSendEnabledParam, simState.Rand,
		func(r *rand.Rand) { defaultSendEnabledParam = RandomGenesisDefaultSendParam(r) },
	)

	numAccs := int64(len(simState.Accounts))

	supply := GenesisCoins(sdk.DefaultBondDenom, simState.InitialStake)
	for i, coin := range supply {
		supply[i] = sdk.NewCoin(coin.Denom, coin.Amount.MulRaw(numAccs))
	}
	supply = supply.Add(sdk.NewInt64Coin(sdk.DefaultBondDenom, simState.InitialStake*simState.NumBonded))

	bankGenesis := types.GenesisState{
		Params: types.Params{
			SendEnabled:        sendEnabledParams,
			DefaultSendEnabled: defaultSendEnabledParam,
		},
		Balances: RandomGenesisBalances(simState),
		Supply:   supply,
	}

	paramsBytes, err := json.MarshalIndent(&bankGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated bank parameters:\n%s\n", paramsBytes)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
}
