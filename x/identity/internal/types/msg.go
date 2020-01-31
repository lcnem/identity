package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgSet{}

// MsgSet - struct for unjailing jailed validator
type MsgSet struct {
	Address       sdk.AccAddress    `json:"address"`
	KeyValuePairs map[string]string `json:"key_value_pairs"`
}

// NewMsgSet creates a new MsgSet instance
func NewMsgSet(address sdk.AccAddress, keyValuePairs map[string]string) MsgSet {
	return MsgSet{
		Address:       address,
		KeyValuePairs: keyValuePairs,
	}
}

const setConst = "set"

// nolint
func (msg MsgSet) Route() string { return RouterKey }
func (msg MsgSet) Type() string  { return setConst }
func (msg MsgSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Address)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSet) ValidateBasic() error {
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing address")
	}

	exp := regexp.MustCompile("^[a-z_][a-z_0-9]*$")
	for key := range msg.KeyValuePairs {
		if !exp.Match([]byte(key)) {
			return ErrInvalidKey
		}
	}

	return nil
}

//verify interface at compile time
var _ sdk.Msg = &MsgImport{}

// MsgImport - struct for unjailing jailed validator
type MsgImport struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
}

// NewMsgImport creates a new MsgImport instance
func NewMsgImport(fromAddress sdk.AccAddress, toAddress sdk.AccAddress) MsgImport {
	return MsgImport{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

const importConst = "import"

// nolint
func (msg MsgImport) Route() string { return RouterKey }
func (msg MsgImport) Type() string  { return importConst }
func (msg MsgImport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ToAddress)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgImport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgImport) ValidateBasic() error {
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	if msg.ToAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing to address")
	}
	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgDelete{}

// MsgDelete - struct for unjailing jailed validator
type MsgDelete struct {
	Address sdk.AccAddress `json:"address"`
}

// NewMsgDelete creates a new MsgDelete instance
func NewMsgDelete(address sdk.AccAddress) MsgDelete {
	return MsgDelete{
		Address: address,
	}
}

const deleteConst = "delete"

// nolint
func (msg MsgDelete) Route() string { return RouterKey }
func (msg MsgDelete) Type() string  { return deleteConst }
func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Address)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDelete) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDelete) ValidateBasic() error {
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing address")
	}
	return nil
}
