package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chain/mpatest/x/mpa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) CreatePosition(goCtx context.Context, msg *types.MsgCreatePosition) (*types.MsgCreatePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var position = types.Position{
		Creator:    msg.Creator,
		Collateral: msg.Collateral,
		Supply:     msg.Supply,
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	feeCoins, err := sdk.ParseCoinsNormalized(strconv.Itoa(int(msg.Collateral)) + "token")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err := k.bankKeeper.SendCoins(ctx, creatorAddress, moduleAcct, feeCoins); err != nil {
		return nil, err
	}

	mintCoin := sdk.NewCoin("mpausd", sdk.NewIntFromUint64(2000))
	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintCoin)); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, creatorAddress, sdk.NewCoins(mintCoin)); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	id := k.AppendPosition(
		ctx,
		position,
	)

	return &types.MsgCreatePositionResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePosition(goCtx context.Context, msg *types.MsgUpdatePosition) (*types.MsgUpdatePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var position = types.Position{
		Creator:    msg.Creator,
		Id:         msg.Id,
		Collateral: msg.Collateral,
		Supply:     msg.Supply,
	}

	// Checks that the element exists
	if !k.HasPosition(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPositionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPosition(ctx, position)

	return &types.MsgUpdatePositionResponse{}, nil
}

func (k msgServer) DeletePosition(goCtx context.Context, msg *types.MsgDeletePosition) (*types.MsgDeletePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasPosition(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPositionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePosition(ctx, msg.Id)

	return &types.MsgDeletePositionResponse{}, nil
}
