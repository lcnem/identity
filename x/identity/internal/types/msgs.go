package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// nolint
const RouterKey = ModuleName

//nolint
type MsgSet struct {
	Address       sdk.AccAddress    `json:"address"`
	KeyValuePairs map[string]string `json:"key_value_pairs"`
}

//nolint
func NewMsgSet(address sdk.AccAddress, keyValuePairs map[string]string) MsgSet {
	return MsgSet{
		Address:       address,
		KeyValuePairs: keyValuePairs,
	}
}

// nolint
func (msg MsgSet) Route() string { return RouterKey }

// nolint
func (msg MsgSet) Type() string { return "set" }

// nolint
func (msg MsgSet) ValidateBasic() error {
	var exp = regexp.MustCompile("^[a-z_][a-z_0-9]*$")
	for key := range msg.KeyValuePairs {
		if !exp.Match([]byte(key)) {
			return ErrInvalidKey
		}
	}

	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
	}

	return nil
}

// nolint
func (msg MsgSet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// nolint
func (msg MsgSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// nolint
type MsgImport struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
}

// nolint
func NewMsgImport(fromAddress sdk.AccAddress, toAddress sdk.AccAddress) MsgImport {
	return MsgImport{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

// nolint
func (msg MsgImport) Route() string { return RouterKey }

// nolint
func (msg MsgImport) Type() string { return "import" }

// nolint
func (msg MsgImport) ValidateBasic() error {
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.FromAddress.String())
	}
	if msg.ToAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ToAddress.String())
	}

	return nil
}

// nolint
func (msg MsgImport) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// nolint
func (msg MsgImport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ToAddress}
}

// nolint
type MsgDelete struct {
	Address sdk.AccAddress `json:"address"`
}

// nolint
func NewMsgDelete(address sdk.AccAddress) MsgDelete {
	return MsgDelete{
		Address: address,
	}
}

// nolint
func (msg MsgDelete) Route() string { return RouterKey }

// nolint
func (msg MsgDelete) Type() string { return "delete" }

// nolint
func (msg MsgDelete) ValidateBasic() error {
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
	}

	return nil
}

// nolint
func (msg MsgDelete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// nolint
func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
