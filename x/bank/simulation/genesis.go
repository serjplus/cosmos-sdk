package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/serjplus/cosmos-sdk/types"
	"github.com/serjplus/cosmos-sdk/types/module"
	"github.com/serjplus/cosmos-sdk/x/bank/internal/types"
)

// Simulation parameter constants
const (
	SendEnabled = "send_enabled"
)

// GenSendEnabled randomized SendEnabled
func GenSendEnabled(r *rand.Rand) bool {
	return r.Int63n(101) <= 95 // 95% chance of transfers being enabled
}

// RandomGenesisAccounts returns a slice of account balances. Each account has
// a balance of simState.InitialStake for sdk.DefaultBondDenom.
func RandomGenesisBalances(simState *module.SimulationState) []types.Balance {
	genesisBalances := []types.Balance{}

	for _, acc := range simState.Accounts {
		genesisBalances = append(genesisBalances, types.Balance{
			Address: acc.Address,
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(simState.InitialStake))),
		})
	}

	return genesisBalances
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) {
	var sendEnabled bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, SendEnabled, &sendEnabled, simState.Rand,
		func(r *rand.Rand) { sendEnabled = GenSendEnabled(r) },
	)

	bankGenesis := types.NewGenesisState(sendEnabled, RandomGenesisBalances(simState))

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(bankGenesis)
}
