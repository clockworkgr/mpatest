package keeper

import (
	"github.com/chain/mpatest/x/mpatest/types"
)

var _ types.QueryServer = Keeper{}
