package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "gimmie/testutil/keeper"
    "gimmie/x/gimmie/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.GimmieKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
