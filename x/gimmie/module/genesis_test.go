package gimmie_test

import (
	"testing"

	keepertest "gimmie/testutil/keeper"
	"gimmie/testutil/nullify"
	"gimmie/x/gimmie/module"
	"gimmie/x/gimmie/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GimmieKeeper(t)
	gimmie.InitGenesis(ctx, k, genesisState)
	got := gimmie.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
