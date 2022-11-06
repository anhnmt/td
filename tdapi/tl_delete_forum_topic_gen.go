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

// DeleteForumTopicRequest represents TL type `deleteForumTopic#6f285cb8`.
type DeleteForumTopicRequest struct {
	// Identifier of the chat
	ChatID int64
	// Message thread identifier of the forum topic
	MessageThreadID int64
}

// DeleteForumTopicRequestTypeID is TL type id of DeleteForumTopicRequest.
const DeleteForumTopicRequestTypeID = 0x6f285cb8

// Ensuring interfaces in compile-time for DeleteForumTopicRequest.
var (
	_ bin.Encoder     = &DeleteForumTopicRequest{}
	_ bin.Decoder     = &DeleteForumTopicRequest{}
	_ bin.BareEncoder = &DeleteForumTopicRequest{}
	_ bin.BareDecoder = &DeleteForumTopicRequest{}
)

func (d *DeleteForumTopicRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.MessageThreadID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteForumTopicRequest) String() string {
	if d == nil {
		return "DeleteForumTopicRequest(nil)"
	}
	type Alias DeleteForumTopicRequest
	return fmt.Sprintf("DeleteForumTopicRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteForumTopicRequest) TypeID() uint32 {
	return DeleteForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteForumTopicRequest) TypeName() string {
	return "deleteForumTopic"
}

// TypeInfo returns info about TL type.
func (d *DeleteForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteForumTopic",
		ID:   DeleteForumTopicRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteForumTopicRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteForumTopic#6f285cb8 as nil")
	}
	b.PutID(DeleteForumTopicRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteForumTopic#6f285cb8 as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutInt53(d.MessageThreadID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteForumTopicRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteForumTopic#6f285cb8 to nil")
	}
	if err := b.ConsumeID(DeleteForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteForumTopic#6f285cb8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: field message_thread_id: %w", err)
		}
		d.MessageThreadID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteForumTopicRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteForumTopic#6f285cb8 as nil")
	}
	b.ObjStart()
	b.PutID("deleteForumTopic")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(d.MessageThreadID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteForumTopicRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteForumTopic#6f285cb8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteForumTopic"); err != nil {
				return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: field chat_id: %w", err)
			}
			d.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteForumTopic#6f285cb8: field message_thread_id: %w", err)
			}
			d.MessageThreadID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteForumTopicRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (d *DeleteForumTopicRequest) GetMessageThreadID() (value int64) {
	if d == nil {
		return
	}
	return d.MessageThreadID
}

// DeleteForumTopic invokes method deleteForumTopic#6f285cb8 returning error if any.
func (c *Client) DeleteForumTopic(ctx context.Context, request *DeleteForumTopicRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
