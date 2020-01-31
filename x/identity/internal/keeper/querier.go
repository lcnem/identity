package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/identity/x/identity/internal/types"
)

// NewQuerier creates a new querier for identity clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAddress:
			return queryAddress(ctx, k, path[1])
			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown identity query endpoint")
		}
	}
}

func queryAddress(ctx sdk.Context, k Keeper, address string) ([]byte, error) {
	params, err := k.Get(ctx, address)
	if err != nil {
		return nil, types.ErrIsNotRegistered
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
