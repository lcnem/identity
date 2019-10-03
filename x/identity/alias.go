package identity

import (
	"github.com/lcnem/identity/x/identity/internal/types"
	"github.com/lcnem/identity/x/identity/keeper"
)

// nolint
const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
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
	ErrImportConflict = types.ErrImportConflict
)

type (
	// Keeper keeper.Keeper
	Keeper = keeper.Keeper
	// MsgSet types.MsgSet
	MsgSet = types.MsgSet
	// MsgImport types.MsgImport
	MsgImport = types.MsgImport
	// MsgDelete types.MsgDelete
	MsgDelete = types.MsgDelete
)
