package keeper

import (
	"context"

	"github.com/BitCannaGlobal/bcna/x/burn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BurnCoinsAction(goCtx context.Context, msg *types.MsgBurnCoinsAction) (*types.MsgBurnCoinsActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, msg.Coins); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Coins); err != nil {
		return nil, err
	}

	return &types.MsgBurnCoinsActionResponse{}, nil
}
