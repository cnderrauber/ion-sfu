// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/media/media.proto

package media

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Info struct {
	Dc                   string   `protobuf:"bytes,1,opt,name=dc,proto3" json:"dc,omitempty"`
	Nid                  string   `protobuf:"bytes,2,opt,name=nid,proto3" json:"nid,omitempty"`
	Rid                  string   `protobuf:"bytes,3,opt,name=rid,proto3" json:"rid,omitempty"`
	Uid                  string   `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Mid                  string   `protobuf:"bytes,5,opt,name=mid,proto3" json:"mid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3655907f88b3485, []int{0}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *Info) GetNid() string {
	if m != nil {
		return m.Nid
	}
	return ""
}

func (m *Info) GetRid() string {
	if m != nil {
		return m.Rid
	}
	return ""
}

func (m *Info) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Info) GetMid() string {
	if m != nil {
		return m.Mid
	}
	return ""
}

type Stream struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tracks               []*Track `protobuf:"bytes,2,rep,name=tracks,proto3" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stream) Reset()         { *m = Stream{} }
func (m *Stream) String() string { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()    {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3655907f88b3485, []int{1}
}

func (m *Stream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stream.Unmarshal(m, b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
}
func (m *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(m, src)
}
func (m *Stream) XXX_Size() int {
	return xxx_messageInfo_Stream.Size(m)
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Stream) GetTracks() []*Track {
	if m != nil {
		return m.Tracks
	}
	return nil
}

type Track struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ssrc                 uint32   `protobuf:"varint,3,opt,name=ssrc,proto3" json:"ssrc,omitempty"`
	Payload              uint32   `protobuf:"varint,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Codec                string   `protobuf:"bytes,6,opt,name=codec,proto3" json:"codec,omitempty"`
	Fmtp                 string   `protobuf:"bytes,7,opt,name=fmtp,proto3" json:"fmtp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Track) Reset()         { *m = Track{} }
func (m *Track) String() string { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()    {}
func (*Track) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3655907f88b3485, []int{2}
}

func (m *Track) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Track.Unmarshal(m, b)
}
func (m *Track) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Track.Marshal(b, m, deterministic)
}
func (m *Track) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Track.Merge(m, src)
}
func (m *Track) XXX_Size() int {
	return xxx_messageInfo_Track.Size(m)
}
func (m *Track) XXX_DiscardUnknown() {
	xxx_messageInfo_Track.DiscardUnknown(m)
}

var xxx_messageInfo_Track proto.InternalMessageInfo

func (m *Track) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Track) GetSsrc() uint32 {
	if m != nil {
		return m.Ssrc
	}
	return 0
}

func (m *Track) GetPayload() uint32 {
	if m != nil {
		return m.Payload
	}
	return 0
}

func (m *Track) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Track) GetCodec() string {
	if m != nil {
		return m.Codec
	}
	return ""
}

func (m *Track) GetFmtp() string {
	if m != nil {
		return m.Fmtp
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "media.Info")
	proto.RegisterType((*Stream)(nil), "media.Stream")
	proto.RegisterType((*Track)(nil), "media.Track")
}

func init() { proto.RegisterFile("pkg/proto/media/media.proto", fileDescriptor_c3655907f88b3485) }

var fileDescriptor_c3655907f88b3485 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4b, 0x04, 0x31,
	0x14, 0xc4, 0xb9, 0xfd, 0x77, 0xf8, 0xf4, 0x44, 0x82, 0x45, 0xc0, 0xe6, 0x58, 0x04, 0xcf, 0xc2,
	0x5d, 0xd0, 0xde, 0xc2, 0xce, 0x76, 0xb5, 0xb2, 0x10, 0x72, 0x79, 0xbb, 0xe7, 0xe3, 0xcc, 0x26,
	0x64, 0xb3, 0xc5, 0xd5, 0x7e, 0x71, 0x49, 0xde, 0x5d, 0xa3, 0x4d, 0x98, 0xf9, 0x65, 0xc2, 0x0c,
	0x81, 0x1b, 0xb7, 0xdf, 0xb5, 0xce, 0xdb, 0x60, 0x5b, 0xd3, 0x23, 0x29, 0x3e, 0x9b, 0x44, 0x44,
	0x99, 0x4c, 0xfd, 0x09, 0xc5, 0xeb, 0x38, 0x58, 0x71, 0x09, 0x19, 0x6a, 0xb9, 0x58, 0x2f, 0x36,
	0x67, 0x5d, 0x86, 0x5a, 0x5c, 0x41, 0x3e, 0x12, 0xca, 0x2c, 0x81, 0x28, 0x23, 0xf1, 0x84, 0x32,
	0x67, 0xe2, 0x99, 0xcc, 0x84, 0xb2, 0x60, 0x32, 0x33, 0x31, 0x84, 0xb2, 0x64, 0x62, 0x08, 0xeb,
	0x67, 0xa8, 0xde, 0x82, 0xef, 0x95, 0x89, 0x0d, 0x84, 0xa7, 0x06, 0x42, 0x71, 0x0b, 0x55, 0xf0,
	0x4a, 0xef, 0x27, 0x99, 0xad, 0xf3, 0xcd, 0xf9, 0xe3, 0x45, 0xc3, 0xf3, 0xde, 0x23, 0xec, 0x8e,
	0x77, 0xf5, 0xcf, 0x02, 0xca, 0x44, 0xfe, 0xbd, 0x17, 0x50, 0x4c, 0x93, 0xd7, 0x69, 0xd0, 0xaa,
	0x4b, 0x5a, 0x48, 0x58, 0x3a, 0x75, 0xf8, 0xb6, 0x8a, 0x57, 0xad, 0xba, 0x93, 0x8d, 0xe9, 0x70,
	0x70, 0xfd, 0x71, 0x5a, 0xd2, 0xe2, 0x1a, 0x4a, 0x6d, 0xb1, 0xd7, 0xb2, 0x4a, 0x90, 0x4d, 0x4c,
	0x0e, 0x26, 0x38, 0xb9, 0xe4, 0x64, 0xd4, 0x2f, 0xf7, 0x1f, 0x77, 0x3b, 0x0a, 0x5f, 0xf3, 0xb6,
	0xd1, 0xd6, 0xb4, 0x8e, 0xec, 0xd8, 0x92, 0x1d, 0x1f, 0xa6, 0x61, 0x6e, 0xff, 0xfc, 0xf1, 0xb6,
	0x4a, 0xe6, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x96, 0xad, 0xfe, 0x7d, 0x01, 0x00, 0x00,
}
