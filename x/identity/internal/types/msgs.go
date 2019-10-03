package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSet defines the Set message
type MsgSet struct {
	Address sdk.AccAddress `json:"address"`
	Key     string         `json:"key"`
	Value   string         `json:"value"`
}

// NewMsgSet is a constructor function
func NewMsgSet(address sdk.AccAddress, key string, value string) MsgSet {
	return MsgSet{
		Address: address,
		Key:     key,
		Value:   value,
	}
}

// Route should return the name of the module
func (msg MsgSet) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSet) Type() string { return "set" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSet) ValidateBasic() sdk.Error {
	if !regexp.MustCompile("^[a-z]([a-z0-9]|-)*$").Match([]byte(msg.Key)) {
		return ErrInvalidKey()
	}
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// MsgImport defines a Import message
type MsgImport struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
}

// NewMsgImport is a constructor function
func NewMsgImport(fromAddress sdk.AccAddress, toAddress sdk.AccAddress) MsgImport {
	return MsgImport{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

// Route should return the name of the module
func (msg MsgImport) Route() string { return RouterKey }

// Type should return the action
func (msg MsgImport) Type() string { return "import" }

// ValidateBasic runs stateless checks on the message
func (msg MsgImport) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.FromAddress.String())
	}
	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.ToAddress.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgImport) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgImport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ToAddress}
}

// MsgDelete defines a Delete message
type MsgDelete struct {
	Address sdk.AccAddress `json:"address"`
}

// NewMsgDelete is a constructor function
func NewMsgDelete(address sdk.AccAddress) MsgDelete {
	return MsgDelete{
		Address: address,
	}
}

// Route should return the name of the module
func (msg MsgDelete) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDelete) Type() string { return "delete" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDelete) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDelete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
