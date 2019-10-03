package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeImportConflict sdk.CodeType = 1
)

// ErrImportConflict ErrImportConflict
func ErrImportConflict() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeImportConflict, "Conflict was occured in import")
}
