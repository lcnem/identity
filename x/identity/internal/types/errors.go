package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrInvalidKey      = sdkerrors.Register(ModuleName, 1, "Available characters are a-z,0-9,_")
	ErrIsNotRegistered = sdkerrors.Register(ModuleName, 2, "Address is not registered")
)
