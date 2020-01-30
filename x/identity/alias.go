package identity

import (
	"github.com/lcnem/identity/x/identity/internal/keeper"
	"github.com/lcnem/identity/x/identity/internal/types"
)

// nolint
const (
	ModuleName             = types.ModuleName
	StoreKey                           = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	RouterKey              = types.RouterKey
	AttributeValueCategory = types.AttributeValueCategory
)

// nolint
var (
	NewKeeper         = keeper.NewKeeper
	NewQuerier        = keeper.NewQuerier
	NewMsgSet         = types.NewMsgSet
	NewMsgImport      = types.NewMsgImport
	NewMsgDelete      = types.NewMsgDelete
	ModuleCdc         = types.ModuleCdc
	RegisterCodec     = types.RegisterCodec
	ErrInvalidKey     = types.ErrInvalidKey
	ErrImportConflict = types.ErrImportConflict
)

// nolint
type (
	Keeper    = keeper.Keeper
	MsgSet    = types.MsgSet
	MsgImport = types.MsgImport
	MsgDelete = types.MsgDelete
)
