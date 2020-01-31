package identity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/identity/x/identity/internal/types"
)

// NewHandler creates an sdk.Handler for all the identity type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		case MsgSet:
			return handleMsgSet(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgSet does x
func handleMsgSet(ctx sdk.Context, msg MsgSet, k Keeper) (*sdk.Result, error) {

	err := k.Set(ctx, msg.Address, msg.KeyValuePairs)
	if err != nil {
		return nil, types.ErrIsNotRegistered
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Address.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handdeImport does x
func handleMsgImport(ctx sdk.Context, msg MsgImport, k Keeper) (*sdk.Result, error) {

	err := k.Import(ctx, msg.FromAddress, msg.ToAddress)
	if err != nil {
		return nil, types.ErrIsNotRegistered
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ToAddress.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handdeDelete does x
func handleMsgDelete(ctx sdk.Context, msg MsgDelete, k Keeper) sdk.Result {

	k.Delete(ctx, msg.Address)

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Address.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
