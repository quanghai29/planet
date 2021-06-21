package keeper

import (
	"github.com/quanghai29/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
