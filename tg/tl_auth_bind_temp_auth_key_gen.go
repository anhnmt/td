// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// AuthBindTempAuthKeyRequest represents TL type `auth.bindTempAuthKey#cdd42a05`.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization
// key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a
// time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
type AuthBindTempAuthKeyRequest struct {
	// Permanent auth_key_id to bind to
	PermAuthKeyID int64
	// Random long from Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	Nonce int64
	// Unix timestamp to invalidate temporary key, see Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	ExpiresAt int
	// See Generating encrypted_message¹
	//
	// Links:
	//  1) https://core.telegram.org#generating-encrypted-message
	EncryptedMessage []byte
}

// AuthBindTempAuthKeyRequestTypeID is TL type id of AuthBindTempAuthKeyRequest.
const AuthBindTempAuthKeyRequestTypeID = 0xcdd42a05

// Ensuring interfaces in compile-time for AuthBindTempAuthKeyRequest.
var (
	_ bin.Encoder     = &AuthBindTempAuthKeyRequest{}
	_ bin.Decoder     = &AuthBindTempAuthKeyRequest{}
	_ bin.BareEncoder = &AuthBindTempAuthKeyRequest{}
	_ bin.BareDecoder = &AuthBindTempAuthKeyRequest{}
)

func (b *AuthBindTempAuthKeyRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.PermAuthKeyID == 0) {
		return false
	}
	if !(b.Nonce == 0) {
		return false
	}
	if !(b.ExpiresAt == 0) {
		return false
	}
	if !(b.EncryptedMessage == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *AuthBindTempAuthKeyRequest) String() string {
	if b == nil {
		return "AuthBindTempAuthKeyRequest(nil)"
	}
	type Alias AuthBindTempAuthKeyRequest
	return fmt.Sprintf("AuthBindTempAuthKeyRequest%+v", Alias(*b))
}

// FillFrom fills AuthBindTempAuthKeyRequest from given interface.
func (b *AuthBindTempAuthKeyRequest) FillFrom(from interface {
	GetPermAuthKeyID() (value int64)
	GetNonce() (value int64)
	GetExpiresAt() (value int)
	GetEncryptedMessage() (value []byte)
}) {
	b.PermAuthKeyID = from.GetPermAuthKeyID()
	b.Nonce = from.GetNonce()
	b.ExpiresAt = from.GetExpiresAt()
	b.EncryptedMessage = from.GetEncryptedMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthBindTempAuthKeyRequest) TypeID() uint32 {
	return AuthBindTempAuthKeyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthBindTempAuthKeyRequest) TypeName() string {
	return "auth.bindTempAuthKey"
}

// TypeInfo returns info about TL type.
func (b *AuthBindTempAuthKeyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.bindTempAuthKey",
		ID:   AuthBindTempAuthKeyRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PermAuthKeyID",
			SchemaName: "perm_auth_key_id",
		},
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ExpiresAt",
			SchemaName: "expires_at",
		},
		{
			Name:       "EncryptedMessage",
			SchemaName: "encrypted_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *AuthBindTempAuthKeyRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode auth.bindTempAuthKey#cdd42a05 as nil")
	}
	buf.PutID(AuthBindTempAuthKeyRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *AuthBindTempAuthKeyRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode auth.bindTempAuthKey#cdd42a05 as nil")
	}
	buf.PutLong(b.PermAuthKeyID)
	buf.PutLong(b.Nonce)
	buf.PutInt(b.ExpiresAt)
	buf.PutBytes(b.EncryptedMessage)
	return nil
}

// Decode implements bin.Decoder.
func (b *AuthBindTempAuthKeyRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode auth.bindTempAuthKey#cdd42a05 to nil")
	}
	if err := buf.ConsumeID(AuthBindTempAuthKeyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *AuthBindTempAuthKeyRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode auth.bindTempAuthKey#cdd42a05 to nil")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field perm_auth_key_id: %w", err)
		}
		b.PermAuthKeyID = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field nonce: %w", err)
		}
		b.Nonce = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field expires_at: %w", err)
		}
		b.ExpiresAt = value
	}
	{
		value, err := buf.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field encrypted_message: %w", err)
		}
		b.EncryptedMessage = value
	}
	return nil
}

// GetPermAuthKeyID returns value of PermAuthKeyID field.
func (b *AuthBindTempAuthKeyRequest) GetPermAuthKeyID() (value int64) {
	if b == nil {
		return
	}
	return b.PermAuthKeyID
}

// GetNonce returns value of Nonce field.
func (b *AuthBindTempAuthKeyRequest) GetNonce() (value int64) {
	if b == nil {
		return
	}
	return b.Nonce
}

// GetExpiresAt returns value of ExpiresAt field.
func (b *AuthBindTempAuthKeyRequest) GetExpiresAt() (value int) {
	if b == nil {
		return
	}
	return b.ExpiresAt
}

// GetEncryptedMessage returns value of EncryptedMessage field.
func (b *AuthBindTempAuthKeyRequest) GetEncryptedMessage() (value []byte) {
	if b == nil {
		return
	}
	return b.EncryptedMessage
}

// AuthBindTempAuthKey invokes method auth.bindTempAuthKey#cdd42a05 returning error if any.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization
// key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a
// time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// Possible errors:
//  400 ENCRYPTED_MESSAGE_INVALID: Encrypted message invalid.
//  400 TEMP_AUTH_KEY_ALREADY_BOUND: The passed temporary key is already bound to another perm_auth_key_id.
//  400 TEMP_AUTH_KEY_EMPTY: No temporary auth key provided.
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
// Can be used by bots.
func (c *Client) AuthBindTempAuthKey(ctx context.Context, request *AuthBindTempAuthKeyRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
