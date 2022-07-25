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

// UploadGetFileRequest represents TL type `upload.getFile#be5335be`.
// Returns content of a whole file or its part.
//
// See https://core.telegram.org/method/upload.getFile for reference.
type UploadGetFileRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable some checks on limit and offset values, useful for example to stream videos by
	// keyframes
	Precise bool
	// Whether the current client supports CDN downloads¹
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	CDNSupported bool
	// File location
	Location InputFileLocationClass
	// Number of bytes to be skipped
	Offset int64
	// Number of bytes to be returned
	Limit int
}

// UploadGetFileRequestTypeID is TL type id of UploadGetFileRequest.
const UploadGetFileRequestTypeID = 0xbe5335be

// Ensuring interfaces in compile-time for UploadGetFileRequest.
var (
	_ bin.Encoder     = &UploadGetFileRequest{}
	_ bin.Decoder     = &UploadGetFileRequest{}
	_ bin.BareEncoder = &UploadGetFileRequest{}
	_ bin.BareDecoder = &UploadGetFileRequest{}
)

func (g *UploadGetFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Precise == false) {
		return false
	}
	if !(g.CDNSupported == false) {
		return false
	}
	if !(g.Location == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetFileRequest) String() string {
	if g == nil {
		return "UploadGetFileRequest(nil)"
	}
	type Alias UploadGetFileRequest
	return fmt.Sprintf("UploadGetFileRequest%+v", Alias(*g))
}

// FillFrom fills UploadGetFileRequest from given interface.
func (g *UploadGetFileRequest) FillFrom(from interface {
	GetPrecise() (value bool)
	GetCDNSupported() (value bool)
	GetLocation() (value InputFileLocationClass)
	GetOffset() (value int64)
	GetLimit() (value int)
}) {
	g.Precise = from.GetPrecise()
	g.CDNSupported = from.GetCDNSupported()
	g.Location = from.GetLocation()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadGetFileRequest) TypeID() uint32 {
	return UploadGetFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadGetFileRequest) TypeName() string {
	return "upload.getFile"
}

// TypeInfo returns info about TL type.
func (g *UploadGetFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.getFile",
		ID:   UploadGetFileRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Precise",
			SchemaName: "precise",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "CDNSupported",
			SchemaName: "cdn_supported",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *UploadGetFileRequest) SetFlags() {
	if !(g.Precise == false) {
		g.Flags.Set(0)
	}
	if !(g.CDNSupported == false) {
		g.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (g *UploadGetFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFile#be5335be as nil")
	}
	b.PutID(UploadGetFileRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UploadGetFileRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFile#be5335be as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFile#be5335be: field flags: %w", err)
	}
	if g.Location == nil {
		return fmt.Errorf("unable to encode upload.getFile#be5335be: field location is nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFile#be5335be: field location: %w", err)
	}
	b.PutLong(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFile#be5335be to nil")
	}
	if err := b.ConsumeID(UploadGetFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getFile#be5335be: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UploadGetFileRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFile#be5335be to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode upload.getFile#be5335be: field flags: %w", err)
		}
	}
	g.Precise = g.Flags.Has(0)
	g.CDNSupported = g.Flags.Has(1)
	{
		value, err := DecodeInputFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#be5335be: field location: %w", err)
		}
		g.Location = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#be5335be: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#be5335be: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetPrecise sets value of Precise conditional field.
func (g *UploadGetFileRequest) SetPrecise(value bool) {
	if value {
		g.Flags.Set(0)
		g.Precise = true
	} else {
		g.Flags.Unset(0)
		g.Precise = false
	}
}

// GetPrecise returns value of Precise conditional field.
func (g *UploadGetFileRequest) GetPrecise() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// SetCDNSupported sets value of CDNSupported conditional field.
func (g *UploadGetFileRequest) SetCDNSupported(value bool) {
	if value {
		g.Flags.Set(1)
		g.CDNSupported = true
	} else {
		g.Flags.Unset(1)
		g.CDNSupported = false
	}
}

// GetCDNSupported returns value of CDNSupported conditional field.
func (g *UploadGetFileRequest) GetCDNSupported() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// GetLocation returns value of Location field.
func (g *UploadGetFileRequest) GetLocation() (value InputFileLocationClass) {
	if g == nil {
		return
	}
	return g.Location
}

// GetOffset returns value of Offset field.
func (g *UploadGetFileRequest) GetOffset() (value int64) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *UploadGetFileRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// UploadGetFile invokes method upload.getFile#be5335be returning error if any.
// Returns content of a whole file or its part.
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 FILE_REFERENCE_*: The file reference expired, it must be refreshed.
//  406 FILEREF_UPGRADE_NEEDED: The client has to be updated in order to support file references.
//  400 FILE_ID_INVALID: The provided file id is invalid.
//  400 FILE_REFERENCE_EXPIRED: File reference expired, it must be refetched as described in the documentation.
//  400 LIMIT_INVALID: The provided limit is invalid.
//  400 LOCATION_INVALID: The provided location is invalid.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//  400 OFFSET_INVALID: The provided offset is invalid.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/upload.getFile for reference.
// Can be used by bots.
func (c *Client) UploadGetFile(ctx context.Context, request *UploadGetFileRequest) (UploadFileClass, error) {
	var result UploadFileBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.File, nil
}
