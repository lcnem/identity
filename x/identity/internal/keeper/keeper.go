package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/identity/x/identity/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc *codec.Codec // The wire codec for binary encoding/decoding.

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
}

// NewKeeper creates new instances of the identity Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Get get
func (k Keeper) Get(ctx sdk.Context, address sdk.AccAddress) map[string]string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(address.String()))

	m := map[string]string{}
	json.Unmarshal(bz, &m)

	return m
}

// Set set
func (k Keeper) Set(ctx sdk.Context, address sdk.AccAddress, keyValuePairs map[string]string) {
	store := ctx.KVStore(k.storeKey)

	m := k.Get(ctx, address)
	for key, value := range keyValuePairs {
		if len(value) == 0 {
			delete(m, key)
		} else {
			m[key] = value
		}
	}

	bz, _ := json.Marshal(m)
	store.Set([]byte(address.String()), bz)
}

// Import import
func (k Keeper) Import(ctx sdk.Context, fromAddress sdk.AccAddress, toAddress sdk.AccAddress) sdk.Error {
	from := k.Get(ctx, fromAddress)
	to := k.Get(ctx, toAddress)

	for key := range from {
		_, exist := to[key]
		if exist {
			return types.ErrImportConflict()
		}
	}
	k.Set(ctx, toAddress, from)

	return nil
}

// Delete delete
func (k Keeper) Delete(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(address.String()))
}
