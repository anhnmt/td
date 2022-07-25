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

// UploadSaveFilePartRequest represents TL type `upload.saveFilePart#b304a621`.
// Saves a part of file for further sending to one of the methods.
//
// See https://core.telegram.org/method/upload.saveFilePart for reference.
type UploadSaveFilePartRequest struct {
	// Random file identifier created by the client
	FileID int64
	// Numerical order of a part
	FilePart int
	// Binary data, contend of a part
	Bytes []byte
}

// UploadSaveFilePartRequestTypeID is TL type id of UploadSaveFilePartRequest.
const UploadSaveFilePartRequestTypeID = 0xb304a621

// Ensuring interfaces in compile-time for UploadSaveFilePartRequest.
var (
	_ bin.Encoder     = &UploadSaveFilePartRequest{}
	_ bin.Decoder     = &UploadSaveFilePartRequest{}
	_ bin.BareEncoder = &UploadSaveFilePartRequest{}
	_ bin.BareDecoder = &UploadSaveFilePartRequest{}
)

func (s *UploadSaveFilePartRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FileID == 0) {
		return false
	}
	if !(s.FilePart == 0) {
		return false
	}
	if !(s.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *UploadSaveFilePartRequest) String() string {
	if s == nil {
		return "UploadSaveFilePartRequest(nil)"
	}
	type Alias UploadSaveFilePartRequest
	return fmt.Sprintf("UploadSaveFilePartRequest%+v", Alias(*s))
}

// FillFrom fills UploadSaveFilePartRequest from given interface.
func (s *UploadSaveFilePartRequest) FillFrom(from interface {
	GetFileID() (value int64)
	GetFilePart() (value int)
	GetBytes() (value []byte)
}) {
	s.FileID = from.GetFileID()
	s.FilePart = from.GetFilePart()
	s.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadSaveFilePartRequest) TypeID() uint32 {
	return UploadSaveFilePartRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadSaveFilePartRequest) TypeName() string {
	return "upload.saveFilePart"
}

// TypeInfo returns info about TL type.
func (s *UploadSaveFilePartRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.saveFilePart",
		ID:   UploadSaveFilePartRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "FilePart",
			SchemaName: "file_part",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *UploadSaveFilePartRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveFilePart#b304a621 as nil")
	}
	b.PutID(UploadSaveFilePartRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *UploadSaveFilePartRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveFilePart#b304a621 as nil")
	}
	b.PutLong(s.FileID)
	b.PutInt(s.FilePart)
	b.PutBytes(s.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (s *UploadSaveFilePartRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveFilePart#b304a621 to nil")
	}
	if err := b.ConsumeID(UploadSaveFilePartRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *UploadSaveFilePartRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveFilePart#b304a621 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field file_id: %w", err)
		}
		s.FileID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field file_part: %w", err)
		}
		s.FilePart = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field bytes: %w", err)
		}
		s.Bytes = value
	}
	return nil
}

// GetFileID returns value of FileID field.
func (s *UploadSaveFilePartRequest) GetFileID() (value int64) {
	if s == nil {
		return
	}
	return s.FileID
}

// GetFilePart returns value of FilePart field.
func (s *UploadSaveFilePartRequest) GetFilePart() (value int) {
	if s == nil {
		return
	}
	return s.FilePart
}

// GetBytes returns value of Bytes field.
func (s *UploadSaveFilePartRequest) GetBytes() (value []byte) {
	if s == nil {
		return
	}
	return s.Bytes
}

// UploadSaveFilePart invokes method upload.saveFilePart#b304a621 returning error if any.
// Saves a part of file for further sending to one of the methods.
//
// Possible errors:
//  400 FILE_PART_EMPTY: The provided file part is empty.
//  400 FILE_PART_INVALID: The file part number is invalid.
//
// See https://core.telegram.org/method/upload.saveFilePart for reference.
// Can be used by bots.
func (c *Client) UploadSaveFilePart(ctx context.Context, request *UploadSaveFilePartRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
