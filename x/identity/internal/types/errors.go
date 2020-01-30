package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// nolint
var (
	ErrInvalidKey     = sdkerrors.Register(ModuleName, 1, "Key must start with [a-z] and available characters are [a-z],[0-9],_")
	ErrImportConflict = sdkerrors.Register(ModuleName, 2, "Conflict was occured in import")
)
