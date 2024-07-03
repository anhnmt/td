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

// InputPaidMediaTypePhoto represents TL type `inputPaidMediaTypePhoto#d299fd1a`.
type InputPaidMediaTypePhoto struct {
}

// InputPaidMediaTypePhotoTypeID is TL type id of InputPaidMediaTypePhoto.
const InputPaidMediaTypePhotoTypeID = 0xd299fd1a

// construct implements constructor of InputPaidMediaTypeClass.
func (i InputPaidMediaTypePhoto) construct() InputPaidMediaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputPaidMediaTypePhoto.
var (
	_ bin.Encoder     = &InputPaidMediaTypePhoto{}
	_ bin.Decoder     = &InputPaidMediaTypePhoto{}
	_ bin.BareEncoder = &InputPaidMediaTypePhoto{}
	_ bin.BareDecoder = &InputPaidMediaTypePhoto{}

	_ InputPaidMediaTypeClass = &InputPaidMediaTypePhoto{}
)

func (i *InputPaidMediaTypePhoto) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaidMediaTypePhoto) String() string {
	if i == nil {
		return "InputPaidMediaTypePhoto(nil)"
	}
	type Alias InputPaidMediaTypePhoto
	return fmt.Sprintf("InputPaidMediaTypePhoto%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaidMediaTypePhoto) TypeID() uint32 {
	return InputPaidMediaTypePhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaidMediaTypePhoto) TypeName() string {
	return "inputPaidMediaTypePhoto"
}

// TypeInfo returns info about TL type.
func (i *InputPaidMediaTypePhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaidMediaTypePhoto",
		ID:   InputPaidMediaTypePhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaidMediaTypePhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypePhoto#d299fd1a as nil")
	}
	b.PutID(InputPaidMediaTypePhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaidMediaTypePhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypePhoto#d299fd1a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaidMediaTypePhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypePhoto#d299fd1a to nil")
	}
	if err := b.ConsumeID(InputPaidMediaTypePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaidMediaTypePhoto#d299fd1a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaidMediaTypePhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypePhoto#d299fd1a to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputPaidMediaTypePhoto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypePhoto#d299fd1a as nil")
	}
	b.ObjStart()
	b.PutID("inputPaidMediaTypePhoto")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputPaidMediaTypePhoto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypePhoto#d299fd1a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputPaidMediaTypePhoto"); err != nil {
				return fmt.Errorf("unable to decode inputPaidMediaTypePhoto#d299fd1a: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// InputPaidMediaTypeVideo represents TL type `inputPaidMediaTypeVideo#b053f9fc`.
type InputPaidMediaTypeVideo struct {
	// Duration of the video, in seconds
	Duration int32
	// True, if the video is supposed to be streamed
	SupportsStreaming bool
}

// InputPaidMediaTypeVideoTypeID is TL type id of InputPaidMediaTypeVideo.
const InputPaidMediaTypeVideoTypeID = 0xb053f9fc

// construct implements constructor of InputPaidMediaTypeClass.
func (i InputPaidMediaTypeVideo) construct() InputPaidMediaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputPaidMediaTypeVideo.
var (
	_ bin.Encoder     = &InputPaidMediaTypeVideo{}
	_ bin.Decoder     = &InputPaidMediaTypeVideo{}
	_ bin.BareEncoder = &InputPaidMediaTypeVideo{}
	_ bin.BareDecoder = &InputPaidMediaTypeVideo{}

	_ InputPaidMediaTypeClass = &InputPaidMediaTypeVideo{}
)

func (i *InputPaidMediaTypeVideo) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Duration == 0) {
		return false
	}
	if !(i.SupportsStreaming == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaidMediaTypeVideo) String() string {
	if i == nil {
		return "InputPaidMediaTypeVideo(nil)"
	}
	type Alias InputPaidMediaTypeVideo
	return fmt.Sprintf("InputPaidMediaTypeVideo%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaidMediaTypeVideo) TypeID() uint32 {
	return InputPaidMediaTypeVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaidMediaTypeVideo) TypeName() string {
	return "inputPaidMediaTypeVideo"
}

// TypeInfo returns info about TL type.
func (i *InputPaidMediaTypeVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaidMediaTypeVideo",
		ID:   InputPaidMediaTypeVideoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "SupportsStreaming",
			SchemaName: "supports_streaming",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaidMediaTypeVideo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypeVideo#b053f9fc as nil")
	}
	b.PutID(InputPaidMediaTypeVideoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaidMediaTypeVideo) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypeVideo#b053f9fc as nil")
	}
	b.PutInt32(i.Duration)
	b.PutBool(i.SupportsStreaming)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaidMediaTypeVideo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypeVideo#b053f9fc to nil")
	}
	if err := b.ConsumeID(InputPaidMediaTypeVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaidMediaTypeVideo) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypeVideo#b053f9fc to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: field duration: %w", err)
		}
		i.Duration = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: field supports_streaming: %w", err)
		}
		i.SupportsStreaming = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputPaidMediaTypeVideo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaidMediaTypeVideo#b053f9fc as nil")
	}
	b.ObjStart()
	b.PutID("inputPaidMediaTypeVideo")
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(i.Duration)
	b.Comma()
	b.FieldStart("supports_streaming")
	b.PutBool(i.SupportsStreaming)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputPaidMediaTypeVideo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaidMediaTypeVideo#b053f9fc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputPaidMediaTypeVideo"); err != nil {
				return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: field duration: %w", err)
			}
			i.Duration = value
		case "supports_streaming":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode inputPaidMediaTypeVideo#b053f9fc: field supports_streaming: %w", err)
			}
			i.SupportsStreaming = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (i *InputPaidMediaTypeVideo) GetDuration() (value int32) {
	if i == nil {
		return
	}
	return i.Duration
}

// GetSupportsStreaming returns value of SupportsStreaming field.
func (i *InputPaidMediaTypeVideo) GetSupportsStreaming() (value bool) {
	if i == nil {
		return
	}
	return i.SupportsStreaming
}

// InputPaidMediaTypeClassName is schema name of InputPaidMediaTypeClass.
const InputPaidMediaTypeClassName = "InputPaidMediaType"

// InputPaidMediaTypeClass represents InputPaidMediaType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputPaidMediaType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputPaidMediaTypePhoto: // inputPaidMediaTypePhoto#d299fd1a
//	case *tdapi.InputPaidMediaTypeVideo: // inputPaidMediaTypeVideo#b053f9fc
//	default: panic(v)
//	}
type InputPaidMediaTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPaidMediaTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeInputPaidMediaType implements binary de-serialization for InputPaidMediaTypeClass.
func DecodeInputPaidMediaType(buf *bin.Buffer) (InputPaidMediaTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPaidMediaTypePhotoTypeID:
		// Decoding inputPaidMediaTypePhoto#d299fd1a.
		v := InputPaidMediaTypePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", err)
		}
		return &v, nil
	case InputPaidMediaTypeVideoTypeID:
		// Decoding inputPaidMediaTypeVideo#b053f9fc.
		v := InputPaidMediaTypeVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputPaidMediaType implements binary de-serialization for InputPaidMediaTypeClass.
func DecodeTDLibJSONInputPaidMediaType(buf tdjson.Decoder) (InputPaidMediaTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputPaidMediaTypePhoto":
		// Decoding inputPaidMediaTypePhoto#d299fd1a.
		v := InputPaidMediaTypePhoto{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", err)
		}
		return &v, nil
	case "inputPaidMediaTypeVideo":
		// Decoding inputPaidMediaTypeVideo#b053f9fc.
		v := InputPaidMediaTypeVideo{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaidMediaTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputPaidMediaType boxes the InputPaidMediaTypeClass providing a helper.
type InputPaidMediaTypeBox struct {
	InputPaidMediaType InputPaidMediaTypeClass
}

// Decode implements bin.Decoder for InputPaidMediaTypeBox.
func (b *InputPaidMediaTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaidMediaTypeBox to nil")
	}
	v, err := DecodeInputPaidMediaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaidMediaType = v
	return nil
}

// Encode implements bin.Encode for InputPaidMediaTypeBox.
func (b *InputPaidMediaTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPaidMediaType == nil {
		return fmt.Errorf("unable to encode InputPaidMediaTypeClass as nil")
	}
	return b.InputPaidMediaType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputPaidMediaTypeBox.
func (b *InputPaidMediaTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaidMediaTypeBox to nil")
	}
	v, err := DecodeTDLibJSONInputPaidMediaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaidMediaType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputPaidMediaTypeBox.
func (b *InputPaidMediaTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputPaidMediaType == nil {
		return fmt.Errorf("unable to encode InputPaidMediaTypeClass as nil")
	}
	return b.InputPaidMediaType.EncodeTDLibJSON(buf)
}
