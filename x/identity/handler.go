package identity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// NewHandler returns a handler
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSet:
			return handleMsgSet(ctx, keeper, msg)
		case MsgImport:
			return handleMsgImport(ctx, keeper, msg)
		case MsgDelete:
			return handleMsgDelete(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized identity Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgSet(ctx sdk.Context, keeper Keeper, msg MsgSet) (*sdk.Result, error) {
	keeper.Set(ctx, msg.Address, msg.KeyValuePairs)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgImport(ctx sdk.Context, keeper Keeper, msg MsgImport) (*sdk.Result, error) {
	err := keeper.Import(ctx, msg.FromAddress, msg.ToAddress)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgDelete(ctx sdk.Context, keeper Keeper, msg MsgDelete) (*sdk.Result, error) {
	keeper.Delete(ctx, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
