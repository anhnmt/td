// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// Usernames represents TL type `usernames#bf343063`.
type Usernames struct {
	// List of active usernames; the first one must be shown as the primary username. The
	// order of active usernames can be changed with reorderActiveUsernames or
	// reorderSupergroupActiveUsernames
	ActiveUsernames []string
	// List of currently disabled usernames; the username can be activated with
	// toggleUsernameIsActive/toggleSupergroupUsernameIsActive
	DisabledUsernames []string
	// The active username, which can be changed with setUsername/setSupergroupUsername
	EditableUsername string
}

// UsernamesTypeID is TL type id of Usernames.
const UsernamesTypeID = 0xbf343063

// Ensuring interfaces in compile-time for Usernames.
var (
	_ bin.Encoder     = &Usernames{}
	_ bin.Decoder     = &Usernames{}
	_ bin.BareEncoder = &Usernames{}
	_ bin.BareDecoder = &Usernames{}
)

func (u *Usernames) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ActiveUsernames == nil) {
		return false
	}
	if !(u.DisabledUsernames == nil) {
		return false
	}
	if !(u.EditableUsername == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *Usernames) String() string {
	if u == nil {
		return "Usernames(nil)"
	}
	type Alias Usernames
	return fmt.Sprintf("Usernames%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Usernames) TypeID() uint32 {
	return UsernamesTypeID
}

// TypeName returns name of type in TL schema.
func (*Usernames) TypeName() string {
	return "usernames"
}

// TypeInfo returns info about TL type.
func (u *Usernames) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "usernames",
		ID:   UsernamesTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ActiveUsernames",
			SchemaName: "active_usernames",
		},
		{
			Name:       "DisabledUsernames",
			SchemaName: "disabled_usernames",
		},
		{
			Name:       "EditableUsername",
			SchemaName: "editable_username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *Usernames) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode usernames#bf343063 as nil")
	}
	b.PutID(UsernamesTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *Usernames) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode usernames#bf343063 as nil")
	}
	b.PutInt(len(u.ActiveUsernames))
	for _, v := range u.ActiveUsernames {
		b.PutString(v)
	}
	b.PutInt(len(u.DisabledUsernames))
	for _, v := range u.DisabledUsernames {
		b.PutString(v)
	}
	b.PutString(u.EditableUsername)
	return nil
}

// Decode implements bin.Decoder.
func (u *Usernames) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode usernames#bf343063 to nil")
	}
	if err := b.ConsumeID(UsernamesTypeID); err != nil {
		return fmt.Errorf("unable to decode usernames#bf343063: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *Usernames) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode usernames#bf343063 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode usernames#bf343063: field active_usernames: %w", err)
		}

		if headerLen > 0 {
			u.ActiveUsernames = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: field active_usernames: %w", err)
			}
			u.ActiveUsernames = append(u.ActiveUsernames, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode usernames#bf343063: field disabled_usernames: %w", err)
		}

		if headerLen > 0 {
			u.DisabledUsernames = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: field disabled_usernames: %w", err)
			}
			u.DisabledUsernames = append(u.DisabledUsernames, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode usernames#bf343063: field editable_username: %w", err)
		}
		u.EditableUsername = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *Usernames) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode usernames#bf343063 as nil")
	}
	b.ObjStart()
	b.PutID("usernames")
	b.Comma()
	b.FieldStart("active_usernames")
	b.ArrStart()
	for _, v := range u.ActiveUsernames {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("disabled_usernames")
	b.ArrStart()
	for _, v := range u.DisabledUsernames {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("editable_username")
	b.PutString(u.EditableUsername)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *Usernames) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode usernames#bf343063 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("usernames"); err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: %w", err)
			}
		case "active_usernames":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode usernames#bf343063: field active_usernames: %w", err)
				}
				u.ActiveUsernames = append(u.ActiveUsernames, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: field active_usernames: %w", err)
			}
		case "disabled_usernames":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode usernames#bf343063: field disabled_usernames: %w", err)
				}
				u.DisabledUsernames = append(u.DisabledUsernames, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: field disabled_usernames: %w", err)
			}
		case "editable_username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode usernames#bf343063: field editable_username: %w", err)
			}
			u.EditableUsername = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetActiveUsernames returns value of ActiveUsernames field.
func (u *Usernames) GetActiveUsernames() (value []string) {
	if u == nil {
		return
	}
	return u.ActiveUsernames
}

// GetDisabledUsernames returns value of DisabledUsernames field.
func (u *Usernames) GetDisabledUsernames() (value []string) {
	if u == nil {
		return
	}
	return u.DisabledUsernames
}

// GetEditableUsername returns value of EditableUsername field.
func (u *Usernames) GetEditableUsername() (value string) {
	if u == nil {
		return
	}
	return u.EditableUsername
}
