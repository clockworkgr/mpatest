package keeper

import (
	"github.com/chain/mpatest/x/mpa/types"
)

var _ types.QueryServer = Keeper{}
