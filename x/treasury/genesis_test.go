package treasury_test

import (
	"testing"

	keepertest "github.com/mineplex/mineplex-chain/testutil/keeper"
	"github.com/mineplex/mineplex-chain/testutil/nullify"
	"github.com/mineplex/mineplex-chain/x/treasury"
	"github.com/mineplex/mineplex-chain/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TreasuryKeeper(t)
	treasury.InitGenesis(ctx, *k, genesisState)
	got := treasury.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
