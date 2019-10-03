package identity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSet:
			return handleMsgSet(ctx, keeper, msg)
		case MsgImport:
			return handleMsgImport(ctx, keeper, msg)
		case MsgDelete:
			return handleMsgDelete(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized identity Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSet(ctx sdk.Context, keeper Keeper, msg MsgSet) sdk.Result {
	keeper.Set(ctx, msg.Address, msg.Key, msg.Value)
	return sdk.Result{}
}

func handleMsgImport(ctx sdk.Context, keeper Keeper, msg MsgImport) sdk.Result {
	err := keeper.Import(ctx, msg.FromAddress, msg.ToAddress)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgDelete(ctx sdk.Context, keeper Keeper, msg MsgDelete) sdk.Result {
	keeper.Delete(ctx, msg.Address)

	return sdk.Result{}
}
