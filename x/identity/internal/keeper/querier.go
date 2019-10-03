package keeper

import (
	"encoding/json"

	"github.com/lcnem/identity/x/identity/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QueryAddress:
			return queryAddress(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown securities query endpoint")
		}
	}
}

func queryAddress(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(string(req.Data))

	if err != nil {
		return nil, sdk.ErrInvalidAddress(address.String())
	}

	obj := keeper.Get(ctx, address)
	res, _ := json.Marshall(obj)

	return res, nil
}
