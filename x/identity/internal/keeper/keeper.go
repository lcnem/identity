package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/identity/x/identity/internal/types"
)

// Keeper of the identity store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a identity keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey:   key,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) Get(ctx sdk.Context, key string) (map[string]string, error) {
	store := ctx.KVStore(k.storeKey)
	var item map[string]string
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (k Keeper) set(ctx sdk.Context, key string, value map[string]string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set([]byte(key), bz)
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// Set Set
func (k Keeper) Set(ctx sdk.Context, address sdk.AccAddress, keyValuePairs map[string]string) map[string]string {
	m, err := k.Get(ctx, address.String())
	if err != nil {
		m = map[string]string{}
	}
	for key, value := range keyValuePairs {
		if len(value) == 0 {
			delete(m, key)
		} else {
			m[key] = value
		}
	}

	k.set(ctx, address.String(), m)

	return m
}

// Import import
func (k Keeper) Import(ctx sdk.Context, fromAddress sdk.AccAddress, toAddress sdk.AccAddress) error {
	from, err := k.Get(ctx, fromAddress.String())
	if err != nil {
		return types.ErrIsNotRegistered
	}

	k.Set(ctx, toAddress, from)

	return nil
}

// Delete delete
func (k Keeper) Delete(ctx sdk.Context, address sdk.AccAddress) {
	k.delete(ctx, address.String())
}
