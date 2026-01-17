package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "gimmie/testutil/keeper"
    "gimmie/x/gimmie/types"
    "gimmie/x/gimmie/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.GimmieKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}