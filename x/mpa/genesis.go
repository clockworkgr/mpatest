package mpa

import (
	"fmt"

	"github.com/chain/mpatest/x/mpa/keeper"
	"github.com/chain/mpatest/x/mpa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the position
	for _, elem := range genState.PositionList {
		k.SetPosition(ctx, *elem)
	}

	// Set position count
	k.SetPositionCount(ctx, uint64(len(genState.PositionList)))

	// check if the module account exists
	moduleAcc := k.GetTransferAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all position
	positionList := k.GetAllPosition(ctx)
	for _, elem := range positionList {
		elem := elem
		genesis.PositionList = append(genesis.PositionList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
