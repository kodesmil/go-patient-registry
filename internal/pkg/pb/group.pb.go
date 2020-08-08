// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/group.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	query "github.com/infobloxopen/atlas-app-toolkit/query"
	resource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	types "github.com/infobloxopen/protoc-gen-gorm/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

type Group struct {
	Id                   *types.UUIDValue     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Notes                string               `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
	ProfileId            *resource.Identifier `protobuf:"bytes,4,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{0}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() *types.UUIDValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Group) GetProfileId() *resource.Identifier {
	if m != nil {
		return m.ProfileId
	}
	return nil
}

type CreateGroupRequest struct {
	Payload              *Group   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{1}
}

func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(m, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetPayload() *Group {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateGroupResponse struct {
	Result               *Group   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupResponse) Reset()         { *m = CreateGroupResponse{} }
func (m *CreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGroupResponse) ProtoMessage()    {}
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{2}
}

func (m *CreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupResponse.Unmarshal(m, b)
}
func (m *CreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupResponse.Merge(m, src)
}
func (m *CreateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGroupResponse.Size(m)
}
func (m *CreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupResponse proto.InternalMessageInfo

func (m *CreateGroupResponse) GetResult() *Group {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReadGroupRequest struct {
	Id                   *types.UUIDValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReadGroupRequest) Reset()         { *m = ReadGroupRequest{} }
func (m *ReadGroupRequest) String() string { return proto.CompactTextString(m) }
func (*ReadGroupRequest) ProtoMessage()    {}
func (*ReadGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{3}
}

func (m *ReadGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroupRequest.Unmarshal(m, b)
}
func (m *ReadGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroupRequest.Marshal(b, m, deterministic)
}
func (m *ReadGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroupRequest.Merge(m, src)
}
func (m *ReadGroupRequest) XXX_Size() int {
	return xxx_messageInfo_ReadGroupRequest.Size(m)
}
func (m *ReadGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroupRequest proto.InternalMessageInfo

func (m *ReadGroupRequest) GetId() *types.UUIDValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type ReadGroupResponse struct {
	Result               *Group   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadGroupResponse) Reset()         { *m = ReadGroupResponse{} }
func (m *ReadGroupResponse) String() string { return proto.CompactTextString(m) }
func (*ReadGroupResponse) ProtoMessage()    {}
func (*ReadGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{4}
}

func (m *ReadGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroupResponse.Unmarshal(m, b)
}
func (m *ReadGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroupResponse.Marshal(b, m, deterministic)
}
func (m *ReadGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroupResponse.Merge(m, src)
}
func (m *ReadGroupResponse) XXX_Size() int {
	return xxx_messageInfo_ReadGroupResponse.Size(m)
}
func (m *ReadGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroupResponse proto.InternalMessageInfo

func (m *ReadGroupResponse) GetResult() *Group {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateGroupRequest struct {
	Payload              *Group   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{5}
}

func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(m, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetPayload() *Group {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UpdateGroupResponse struct {
	Result               *Group   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupResponse) Reset()         { *m = UpdateGroupResponse{} }
func (m *UpdateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupResponse) ProtoMessage()    {}
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{6}
}

func (m *UpdateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupResponse.Unmarshal(m, b)
}
func (m *UpdateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupResponse.Marshal(b, m, deterministic)
}
func (m *UpdateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupResponse.Merge(m, src)
}
func (m *UpdateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupResponse.Size(m)
}
func (m *UpdateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupResponse proto.InternalMessageInfo

func (m *UpdateGroupResponse) GetResult() *Group {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteGroupRequest struct {
	Id                   *types.UUIDValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{7}
}

func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(m, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetId() *types.UUIDValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type DeleteGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGroupResponse) Reset()         { *m = DeleteGroupResponse{} }
func (m *DeleteGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupResponse) ProtoMessage()    {}
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{8}
}

func (m *DeleteGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupResponse.Unmarshal(m, b)
}
func (m *DeleteGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupResponse.Marshal(b, m, deterministic)
}
func (m *DeleteGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupResponse.Merge(m, src)
}
func (m *DeleteGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupResponse.Size(m)
}
func (m *DeleteGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupResponse proto.InternalMessageInfo

type ListGroupRequest struct {
	Filter               *query.Filtering      `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              *query.Sorting        `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Fields               *query.FieldSelection `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	Paging               *query.Pagination     `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListGroupRequest) Reset()         { *m = ListGroupRequest{} }
func (m *ListGroupRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupRequest) ProtoMessage()    {}
func (*ListGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{9}
}

func (m *ListGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupRequest.Unmarshal(m, b)
}
func (m *ListGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupRequest.Marshal(b, m, deterministic)
}
func (m *ListGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupRequest.Merge(m, src)
}
func (m *ListGroupRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupRequest.Size(m)
}
func (m *ListGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupRequest proto.InternalMessageInfo

func (m *ListGroupRequest) GetFilter() *query.Filtering {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListGroupRequest) GetOrderBy() *query.Sorting {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListGroupRequest) GetFields() *query.FieldSelection {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListGroupRequest) GetPaging() *query.Pagination {
	if m != nil {
		return m.Paging
	}
	return nil
}

type ListGroupResponse struct {
	Results              []*Group `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupResponse) Reset()         { *m = ListGroupResponse{} }
func (m *ListGroupResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupResponse) ProtoMessage()    {}
func (*ListGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f361d419742be504, []int{10}
}

func (m *ListGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupResponse.Unmarshal(m, b)
}
func (m *ListGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupResponse.Marshal(b, m, deterministic)
}
func (m *ListGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupResponse.Merge(m, src)
}
func (m *ListGroupResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupResponse.Size(m)
}
func (m *ListGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupResponse proto.InternalMessageInfo

func (m *ListGroupResponse) GetResults() []*Group {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "model.Group")
	proto.RegisterType((*CreateGroupRequest)(nil), "model.CreateGroupRequest")
	proto.RegisterType((*CreateGroupResponse)(nil), "model.CreateGroupResponse")
	proto.RegisterType((*ReadGroupRequest)(nil), "model.ReadGroupRequest")
	proto.RegisterType((*ReadGroupResponse)(nil), "model.ReadGroupResponse")
	proto.RegisterType((*UpdateGroupRequest)(nil), "model.UpdateGroupRequest")
	proto.RegisterType((*UpdateGroupResponse)(nil), "model.UpdateGroupResponse")
	proto.RegisterType((*DeleteGroupRequest)(nil), "model.DeleteGroupRequest")
	proto.RegisterType((*DeleteGroupResponse)(nil), "model.DeleteGroupResponse")
	proto.RegisterType((*ListGroupRequest)(nil), "model.ListGroupRequest")
	proto.RegisterType((*ListGroupResponse)(nil), "model.ListGroupResponse")
}

func init() {
	proto.RegisterFile("github.com/kodesmil/ks-model/group.proto", fileDescriptor_f361d419742be504)
}

var fileDescriptor_f361d419742be504 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0x85, 0xbc, 0x8e, 0x9c, 0x65, 0x7a, 0xd8, 0xe5, 0x76, 0xb1, 0xb2, 0x11, 0xb4, 0x86, 0xd0,
	0x06, 0x2e, 0x0a, 0x8b, 0x41, 0x9a, 0x8b, 0x93, 0x00, 0x01, 0xd2, 0xb4, 0x85, 0x8b, 0x1c, 0x0a,
	0x07, 0xe9, 0xa1, 0x17, 0x83, 0x16, 0xc7, 0x2a, 0x61, 0x8a, 0x64, 0x48, 0x2a, 0xa8, 0x7f, 0x5a,
	0x82, 0xfe, 0x90, 0x5e, 0xfa, 0x5f, 0x0a, 0x49, 0xb4, 0xe1, 0xaf, 0xa0, 0x09, 0xf6, 0x62, 0x51,
	0x9c, 0xf7, 0xe6, 0xcd, 0x0c, 0x9f, 0x29, 0xd4, 0xcb, 0xb8, 0xfb, 0xb3, 0x98, 0x24, 0xa9, 0xca,
	0xc9, 0x4c, 0x31, 0xb0, 0x39, 0x17, 0x64, 0x66, 0xfb, 0xb9, 0x62, 0x20, 0x48, 0x66, 0x54, 0xa1,
	0x13, 0x6d, 0x94, 0x53, 0x78, 0xaf, 0xda, 0xea, 0x74, 0x33, 0xa5, 0x32, 0x01, 0xa4, 0xda, 0x9c,
	0x14, 0x53, 0x32, 0xe5, 0x20, 0xd8, 0x38, 0xa7, 0x76, 0x56, 0x03, 0x3b, 0x5f, 0x6f, 0x22, 0x1c,
	0xcf, 0xc1, 0x3a, 0x9a, 0xfb, 0x4c, 0x9d, 0x43, 0x0f, 0xa0, 0x9a, 0x13, 0x2a, 0xa5, 0x72, 0xd4,
	0x71, 0x25, 0xad, 0x8f, 0x7e, 0x5f, 0x3d, 0xd2, 0x7e, 0x06, 0xb2, 0xff, 0x40, 0x05, 0x67, 0xd4,
	0x01, 0xd9, 0x5a, 0x78, 0xf0, 0xd9, 0x4a, 0xf9, 0x5c, 0x4e, 0xd5, 0x44, 0xa8, 0xbf, 0x94, 0x06,
	0x49, 0x56, 0x92, 0x64, 0xca, 0xe4, 0x44, 0xe9, 0x4a, 0x83, 0x94, 0x2f, 0x9e, 0x3b, 0x78, 0x29,
	0xd7, 0xcd, 0x35, 0xd8, 0xfa, 0xd7, 0x53, 0x7f, 0x7d, 0x8e, 0x4a, 0x9d, 0xa0, 0xb6, 0x4f, 0xb5,
	0xee, 0x3b, 0xa5, 0xc4, 0x8c, 0x3b, 0x72, 0x5f, 0x80, 0x99, 0x93, 0x54, 0x09, 0x01, 0x69, 0x59,
	0xc2, 0x58, 0x69, 0x30, 0xd4, 0x29, 0xb3, 0xc8, 0xf5, 0xd3, 0xcb, 0x73, 0x19, 0x9d, 0x12, 0x03,
	0x56, 0x15, 0x26, 0x85, 0xe5, 0xa2, 0x4e, 0x13, 0xff, 0x1d, 0xa0, 0xbd, 0x5f, 0xca, 0xe3, 0xc2,
	0x43, 0xd4, 0xe0, 0x2c, 0x0a, 0xba, 0x41, 0xef, 0xe0, 0xe4, 0x63, 0x52, 0x35, 0x5c, 0xd7, 0x7e,
	0x77, 0x37, 0xbc, 0xfe, 0x9d, 0x8a, 0x02, 0xae, 0xe2, 0xa7, 0xc7, 0xf6, 0x57, 0xe8, 0x10, 0x37,
	0x8b, 0x82, 0xb3, 0x5e, 0x70, 0x86, 0xcb, 0xe7, 0x38, 0x03, 0x59, 0x96, 0x06, 0xe3, 0x87, 0xd3,
	0xde, 0x77, 0xa3, 0x06, 0x67, 0x18, 0xa3, 0xa6, 0xa4, 0x39, 0x44, 0x8d, 0x6e, 0xd0, 0x7b, 0x3b,
	0xaa, 0xd6, 0xf8, 0x4b, 0xb4, 0x27, 0x95, 0x03, 0x1b, 0xbd, 0xa9, 0x36, 0xeb, 0x17, 0x7c, 0x8a,
	0x90, 0x36, 0x6a, 0xca, 0x05, 0x8c, 0x39, 0x8b, 0x9a, 0x5e, 0xbc, 0x6a, 0x21, 0x31, 0x3a, 0x4d,
	0x86, 0x0c, 0xa4, 0xe3, 0x53, 0x0e, 0x66, 0xf4, 0xd6, 0x03, 0x87, 0xec, 0x6c, 0xff, 0xe9, 0xb1,
	0xdd, 0xdc, 0x0f, 0xba, 0x41, 0x7c, 0x81, 0xf0, 0x8f, 0x06, 0xa8, 0x83, 0xaa, 0x87, 0x11, 0xdc,
	0x17, 0x60, 0x1d, 0x3e, 0x42, 0x2d, 0x4d, 0xe7, 0x42, 0xd1, 0x45, 0x3f, 0x5f, 0x24, 0x95, 0x0b,
	0x93, 0x1a, 0xb5, 0x08, 0xc6, 0xe7, 0xe8, 0xc3, 0x1a, 0xdb, 0x6a, 0x25, 0x2d, 0xe0, 0x6f, 0x50,
	0x68, 0xc0, 0x16, 0xc2, 0xed, 0x64, 0xfb, 0x58, 0x3c, 0x40, 0xef, 0x46, 0x40, 0xd9, 0x9a, 0xf0,
	0xb7, 0xff, 0x3b, 0xc3, 0x72, 0x3e, 0xf1, 0x00, 0xbd, 0x5f, 0xa1, 0xbe, 0x4a, 0xf5, 0x02, 0xe1,
	0x3b, 0xcd, 0x3e, 0xa3, 0xe1, 0x35, 0xf6, 0xab, 0xa4, 0xcf, 0x11, 0xbe, 0x06, 0x01, 0x1b, 0xd2,
	0x2f, 0x6c, 0xf9, 0x23, 0xfa, 0xb0, 0x46, 0xae, 0x95, 0xe3, 0x7f, 0x03, 0xf4, 0xee, 0x86, 0x5b,
	0xb7, 0x96, 0x92, 0xa0, 0x70, 0xca, 0x85, 0x03, 0xe3, 0xd3, 0x7e, 0x4a, 0x16, 0x06, 0x4f, 0xa8,
	0xe6, 0xc9, 0xcf, 0x55, 0x8c, 0xcb, 0x6c, 0xe4, 0x61, 0xf8, 0x18, 0xed, 0x2b, 0xc3, 0xc0, 0x8c,
	0x27, 0xf3, 0xca, 0x73, 0x65, 0x25, 0x6b, 0x94, 0x5b, 0x65, 0x5c, 0x49, 0x68, 0x55, 0xb0, 0xab,
	0x39, 0x3e, 0x2d, 0x25, 0x40, 0xb0, 0xda, 0x8e, 0x07, 0x27, 0x87, 0x9b, 0x12, 0x20, 0xd8, 0x2d,
	0xf8, 0xff, 0xde, 0xc8, 0x63, 0xf1, 0x31, 0x0a, 0x35, 0xcd, 0xb8, 0xcc, 0xbc, 0x53, 0xa3, 0x75,
	0xd6, 0x6f, 0x65, 0x8c, 0xd6, 0x8c, 0x1a, 0x17, 0x9f, 0xa3, 0xf7, 0x2b, 0xed, 0xf9, 0x71, 0x1f,
	0xa1, 0x56, 0x3d, 0x52, 0x1b, 0x05, 0xdd, 0x37, 0xdb, 0xa7, 0xe5, 0x83, 0x27, 0xff, 0x34, 0x50,
	0x58, 0x6d, 0x59, 0x7c, 0x89, 0xc2, 0xda, 0xa9, 0xb8, 0xed, 0xb1, 0xdb, 0xb6, 0xef, 0x74, 0x76,
	0x85, 0xbc, 0xe6, 0x00, 0x35, 0x4b, 0xcb, 0xe1, 0x4f, 0x1e, 0xb3, 0x69, 0xdd, 0x4e, 0xb4, 0x1d,
	0xf0, 0xd4, 0x4b, 0x14, 0xd6, 0xa6, 0x59, 0x6a, 0x6f, 0x3b, 0x70, 0xa9, 0xbd, 0xcb, 0x5e, 0x37,
	0x28, 0xac, 0xcf, 0x7e, 0x99, 0x60, 0xdb, 0x47, 0xcb, 0x04, 0xbb, 0x5c, 0x72, 0xf0, 0xf4, 0xd8,
	0x6e, 0x2d, 0xee, 0xa9, 0x01, 0x6a, 0x96, 0x23, 0x5d, 0x76, 0xb2, 0x69, 0x9f, 0x65, 0x27, 0x5b,
	0x83, 0xef, 0xf8, 0x7b, 0x23, 0x0a, 0xae, 0x8e, 0xff, 0x48, 0x9e, 0xf9, 0x82, 0x4d, 0x68, 0x3a,
	0x03, 0xc9, 0x08, 0x97, 0x0e, 0x8c, 0xa4, 0x82, 0xe8, 0x59, 0x46, 0xf4, 0x64, 0x12, 0x56, 0xf7,
	0xe5, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x84, 0xa4, 0x0c, 0xfa, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupsClient is the client API for Groups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupsClient interface {
	Create(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	Read(ctx context.Context, in *ReadGroupRequest, opts ...grpc.CallOption) (*ReadGroupResponse, error)
	Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	Delete(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	List(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error)
}

type groupsClient struct {
	cc *grpc.ClientConn
}

func NewGroupsClient(cc *grpc.ClientConn) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) Create(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/model.Groups/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Read(ctx context.Context, in *ReadGroupRequest, opts ...grpc.CallOption) (*ReadGroupResponse, error) {
	out := new(ReadGroupResponse)
	err := c.cc.Invoke(ctx, "/model.Groups/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, "/model.Groups/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Delete(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, "/model.Groups/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) List(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error) {
	out := new(ListGroupResponse)
	err := c.cc.Invoke(ctx, "/model.Groups/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsServer is the server API for Groups service.
type GroupsServer interface {
	Create(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	Read(context.Context, *ReadGroupRequest) (*ReadGroupResponse, error)
	Update(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	Delete(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	List(context.Context, *ListGroupRequest) (*ListGroupResponse, error)
}

func RegisterGroupsServer(s *grpc.Server, srv GroupsServer) {
	s.RegisterService(&_Groups_serviceDesc, srv)
}

func _Groups_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Groups/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Create(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Groups/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Read(ctx, req.(*ReadGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Groups/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Update(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Groups/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Delete(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Groups/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).List(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Groups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Groups_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Groups_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Groups_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Groups_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Groups_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kodesmil/ks-model/group.proto",
}
