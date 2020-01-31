package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/identity/x/identity/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data types.GenesisState) []abci.ValidatorUpdate {

	// TODO: Define logic for when you would like to initalize a new genesis

	return []abci.ValidatorUpdate{}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data types.GenesisState) {

	// TODO: Define logic for exporting state
	return types.NewGenesisState()
}
