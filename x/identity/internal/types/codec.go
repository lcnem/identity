package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSet{}, "identity/MsgSet", nil)
	cdc.RegisterConcrete(MsgImport{}, "identity/MsgImport", nil)
	cdc.RegisterConcrete(MsgDelete{}, "identity/MsgDelete", nil)
}
