package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidKey     sdk.CodeType = 1
	CodeImportConflict sdk.CodeType = 2
)

// ErrInvalidKey ErrInvalidKey
func ErrInvalidKey() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidKey, "Key must start with [a-z] and available characters are [a-z],[0-9],-")
}

// ErrImportConflict ErrImportConflict
func ErrImportConflict() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeImportConflict, "Conflict was occured in import")
}
