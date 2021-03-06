package exported_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	sdk "github.com/serjplus/cosmos-sdk/types"
	"github.com/serjplus/cosmos-sdk/x/auth/exported"
	authtypes "github.com/serjplus/cosmos-sdk/x/auth/types"
)

func TestGenesisAccountsContains(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	acc := authtypes.NewBaseAccount(addr, secp256k1.GenPrivKey().PubKey(), 0, 0)

	genAccounts := exported.GenesisAccounts{}
	require.False(t, genAccounts.Contains(acc.GetAddress()))

	genAccounts = append(genAccounts, acc)
	require.True(t, genAccounts.Contains(acc.GetAddress()))
}
