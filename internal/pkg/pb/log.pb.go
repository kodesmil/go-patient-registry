// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/log.proto

package pb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/atlas-app-toolkit/query"
	_ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/protobuf/field_mask"
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

type LogActivity struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ip                   string               `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogActivity) Reset()         { *m = LogActivity{} }
func (m *LogActivity) String() string { return proto.CompactTextString(m) }
func (*LogActivity) ProtoMessage()    {}
func (*LogActivity) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b168eac4135d350, []int{0}
}

func (m *LogActivity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogActivity.Unmarshal(m, b)
}
func (m *LogActivity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogActivity.Marshal(b, m, deterministic)
}
func (m *LogActivity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogActivity.Merge(m, src)
}
func (m *LogActivity) XXX_Size() int {
	return xxx_messageInfo_LogActivity.Size(m)
}
func (m *LogActivity) XXX_DiscardUnknown() {
	xxx_messageInfo_LogActivity.DiscardUnknown(m)
}

var xxx_messageInfo_LogActivity proto.InternalMessageInfo

func (m *LogActivity) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogActivity) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LogActivity) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *LogActivity) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func init() {
	proto.RegisterType((*LogActivity)(nil), "model.LogActivity")
}

func init() {
	proto.RegisterFile("github.com/kodesmil/ks-model/log.proto", fileDescriptor_1b168eac4135d350)
}

var fileDescriptor_1b168eac4135d350 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x49, 0x5c, 0x4b, 0x9b, 0x8a, 0x48, 0x4e, 0xdb, 0x45, 0x30, 0x78, 0x90, 0x05, 0x49,
	0x46, 0xf4, 0xd4, 0xde, 0xb6, 0xe0, 0x45, 0x3c, 0x2d, 0x9e, 0xbc, 0x2c, 0x2f, 0x33, 0x6f, 0xe3,
	0x23, 0x93, 0x79, 0xe3, 0xcc, 0x4b, 0x71, 0x3f, 0x5a, 0x7b, 0xf1, 0xab, 0xc9, 0x66, 0x67, 0xa5,
	0x28, 0xa5, 0xbd, 0x24, 0x93, 0x97, 0xdf, 0xef, 0xff, 0x1e, 0xbc, 0x29, 0xde, 0x75, 0x24, 0x3f,
	0xc6, 0xb6, 0xd1, 0x3c, 0xa8, 0x9e, 0x0d, 0xc6, 0x81, 0xac, 0xea, 0x63, 0x3d, 0xb0, 0x41, 0xab,
	0x2c, 0x77, 0x8d, 0x0f, 0x2c, 0x5c, 0x3e, 0x9f, 0x0a, 0x8b, 0xaa, 0x63, 0xee, 0x2c, 0xaa, 0xa9,
	0xd8, 0x8e, 0x5b, 0xb5, 0x25, 0xb4, 0x66, 0x33, 0x40, 0xec, 0x0f, 0xe0, 0xe2, 0xcd, 0xbf, 0x84,
	0xd0, 0x80, 0x51, 0x60, 0xf0, 0x09, 0x78, 0x9d, 0x00, 0xf0, 0xa4, 0xc0, 0x39, 0x16, 0x10, 0x62,
	0x17, 0xd3, 0xdf, 0xf7, 0xd3, 0x4b, 0xd7, 0x1d, 0xba, 0xfa, 0x06, 0x2c, 0x19, 0x10, 0x54, 0xff,
	0x1d, 0x12, 0x7c, 0x75, 0x6f, 0x78, 0x72, 0x5b, 0x6e, 0x2d, 0xff, 0x62, 0x8f, 0x4e, 0xdd, 0x0b,
	0xe9, 0x38, 0x0c, 0x8a, 0xfd, 0xd4, 0x43, 0xed, 0x3f, 0x92, 0x7b, 0xf9, 0x54, 0x57, 0x76, 0x1e,
	0xe3, 0xe1, 0x99, 0xd4, 0x2f, 0x0f, 0xa9, 0x20, 0x16, 0x62, 0x0d, 0xde, 0xd7, 0xc2, 0x6c, 0x7b,
	0x12, 0xf5, 0x73, 0xc4, 0xb0, 0x53, 0x9a, 0xad, 0x45, 0xbd, 0x1f, 0x61, 0xc3, 0x1e, 0x03, 0x08,
	0x87, 0x63, 0xd6, 0xe7, 0xa7, 0x67, 0x05, 0xaf, 0x55, 0xc0, 0xc8, 0x63, 0xd0, 0xf8, 0xf7, 0x70,
	0x88, 0x79, 0xfb, 0x3b, 0x2b, 0xce, 0xbf, 0x72, 0xb7, 0xd2, 0x42, 0x37, 0x24, 0xbb, 0xb2, 0x2a,
	0x72, 0x32, 0xf3, 0xac, 0xca, 0x96, 0xb3, 0xeb, 0x57, 0x77, 0xb7, 0x17, 0x2f, 0x8a, 0xa2, 0x3c,
	0x89, 0x18, 0x08, 0xec, 0x32, 0x5b, 0xe7, 0x64, 0xca, 0xcb, 0xa2, 0xd0, 0x01, 0x41, 0xd0, 0x6c,
	0x40, 0xe6, 0x79, 0x95, 0x2d, 0xcf, 0x3f, 0x2e, 0x9a, 0xc3, 0x6e, 0x9a, 0xe3, 0xf2, 0x9a, 0x6f,
	0xc7, 0xe5, 0xad, 0xcf, 0x12, 0xbd, 0x92, 0xbd, 0x3a, 0x7a, 0x73, 0x54, 0x9f, 0x3d, 0xae, 0x26,
	0x7a, 0x25, 0xe5, 0xcb, 0x22, 0x27, 0x3f, 0x9f, 0x55, 0xd9, 0xf2, 0x6c, 0x9d, 0x93, 0xbf, 0x3a,
	0xbd, 0xbb, 0xbd, 0x98, 0x9d, 0x66, 0x55, 0x76, 0xfd, 0xe1, 0x7b, 0xf3, 0xc0, 0x55, 0x6c, 0x41,
	0xf7, 0xe8, 0x8c, 0x22, 0x27, 0x18, 0x1c, 0x58, 0xe5, 0xfb, 0x4e, 0xf9, 0xb6, 0x3d, 0x99, 0x5a,
	0x7d, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x77, 0x45, 0xfc, 0x61, 0xc3, 0x02, 0x00, 0x00,
}
