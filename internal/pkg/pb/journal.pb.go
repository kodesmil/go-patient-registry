// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/journal.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	query "github.com/infobloxopen/atlas-app-toolkit/query"
	resource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/types"
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

type JournalEntry_Severity int32

const (
	JournalEntry_NONE JournalEntry_Severity = 0
	JournalEntry_LOW  JournalEntry_Severity = 1
	JournalEntry_MID  JournalEntry_Severity = 2
	JournalEntry_HIGH JournalEntry_Severity = 3
)

var JournalEntry_Severity_name = map[int32]string{
	0: "NONE",
	1: "LOW",
	2: "MID",
	3: "HIGH",
}

var JournalEntry_Severity_value = map[string]int32{
	"NONE": 0,
	"LOW":  1,
	"MID":  2,
	"HIGH": 3,
}

func (x JournalEntry_Severity) String() string {
	return proto.EnumName(JournalEntry_Severity_name, int32(x))
}

func (JournalEntry_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{2, 0}
}

type JournalSubjectType struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JournalSubjectType) Reset()         { *m = JournalSubjectType{} }
func (m *JournalSubjectType) String() string { return proto.CompactTextString(m) }
func (*JournalSubjectType) ProtoMessage()    {}
func (*JournalSubjectType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{0}
}

func (m *JournalSubjectType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JournalSubjectType.Unmarshal(m, b)
}
func (m *JournalSubjectType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JournalSubjectType.Marshal(b, m, deterministic)
}
func (m *JournalSubjectType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JournalSubjectType.Merge(m, src)
}
func (m *JournalSubjectType) XXX_Size() int {
	return xxx_messageInfo_JournalSubjectType.Size(m)
}
func (m *JournalSubjectType) XXX_DiscardUnknown() {
	xxx_messageInfo_JournalSubjectType.DiscardUnknown(m)
}

var xxx_messageInfo_JournalSubjectType proto.InternalMessageInfo

func (m *JournalSubjectType) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *JournalSubjectType) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JournalSubjectType) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *JournalSubjectType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JournalSubject struct {
	Id                   *resource.Identifier  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name                 string                `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Types                []*JournalSubjectType `protobuf:"bytes,5,rep,name=types,proto3" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *JournalSubject) Reset()         { *m = JournalSubject{} }
func (m *JournalSubject) String() string { return proto.CompactTextString(m) }
func (*JournalSubject) ProtoMessage()    {}
func (*JournalSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{1}
}

func (m *JournalSubject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JournalSubject.Unmarshal(m, b)
}
func (m *JournalSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JournalSubject.Marshal(b, m, deterministic)
}
func (m *JournalSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JournalSubject.Merge(m, src)
}
func (m *JournalSubject) XXX_Size() int {
	return xxx_messageInfo_JournalSubject.Size(m)
}
func (m *JournalSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_JournalSubject.DiscardUnknown(m)
}

var xxx_messageInfo_JournalSubject proto.InternalMessageInfo

func (m *JournalSubject) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *JournalSubject) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JournalSubject) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *JournalSubject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JournalSubject) GetTypes() []*JournalSubjectType {
	if m != nil {
		return m.Types
	}
	return nil
}

type JournalEntry struct {
	Id                   *resource.Identifier  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Day                  *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`
	Severity             JournalEntry_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=model.JournalEntry_Severity" json:"severity,omitempty"`
	Note                 string                `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *JournalEntry) Reset()         { *m = JournalEntry{} }
func (m *JournalEntry) String() string { return proto.CompactTextString(m) }
func (*JournalEntry) ProtoMessage()    {}
func (*JournalEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{2}
}

func (m *JournalEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JournalEntry.Unmarshal(m, b)
}
func (m *JournalEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JournalEntry.Marshal(b, m, deterministic)
}
func (m *JournalEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JournalEntry.Merge(m, src)
}
func (m *JournalEntry) XXX_Size() int {
	return xxx_messageInfo_JournalEntry.Size(m)
}
func (m *JournalEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_JournalEntry.DiscardUnknown(m)
}

var xxx_messageInfo_JournalEntry proto.InternalMessageInfo

func (m *JournalEntry) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *JournalEntry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JournalEntry) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *JournalEntry) GetDay() *timestamp.Timestamp {
	if m != nil {
		return m.Day
	}
	return nil
}

func (m *JournalEntry) GetSeverity() JournalEntry_Severity {
	if m != nil {
		return m.Severity
	}
	return JournalEntry_NONE
}

func (m *JournalEntry) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

type CreateJournalEntryRequest struct {
	Payload              *JournalEntry `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateJournalEntryRequest) Reset()         { *m = CreateJournalEntryRequest{} }
func (m *CreateJournalEntryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJournalEntryRequest) ProtoMessage()    {}
func (*CreateJournalEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{3}
}

func (m *CreateJournalEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJournalEntryRequest.Unmarshal(m, b)
}
func (m *CreateJournalEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJournalEntryRequest.Marshal(b, m, deterministic)
}
func (m *CreateJournalEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJournalEntryRequest.Merge(m, src)
}
func (m *CreateJournalEntryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJournalEntryRequest.Size(m)
}
func (m *CreateJournalEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJournalEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJournalEntryRequest proto.InternalMessageInfo

func (m *CreateJournalEntryRequest) GetPayload() *JournalEntry {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateJournalEntryResponse struct {
	Result               *JournalEntry `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateJournalEntryResponse) Reset()         { *m = CreateJournalEntryResponse{} }
func (m *CreateJournalEntryResponse) String() string { return proto.CompactTextString(m) }
func (*CreateJournalEntryResponse) ProtoMessage()    {}
func (*CreateJournalEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{4}
}

func (m *CreateJournalEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJournalEntryResponse.Unmarshal(m, b)
}
func (m *CreateJournalEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJournalEntryResponse.Marshal(b, m, deterministic)
}
func (m *CreateJournalEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJournalEntryResponse.Merge(m, src)
}
func (m *CreateJournalEntryResponse) XXX_Size() int {
	return xxx_messageInfo_CreateJournalEntryResponse.Size(m)
}
func (m *CreateJournalEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJournalEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJournalEntryResponse proto.InternalMessageInfo

func (m *CreateJournalEntryResponse) GetResult() *JournalEntry {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReadJournalEntryRequest struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReadJournalEntryRequest) Reset()         { *m = ReadJournalEntryRequest{} }
func (m *ReadJournalEntryRequest) String() string { return proto.CompactTextString(m) }
func (*ReadJournalEntryRequest) ProtoMessage()    {}
func (*ReadJournalEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{5}
}

func (m *ReadJournalEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadJournalEntryRequest.Unmarshal(m, b)
}
func (m *ReadJournalEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadJournalEntryRequest.Marshal(b, m, deterministic)
}
func (m *ReadJournalEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadJournalEntryRequest.Merge(m, src)
}
func (m *ReadJournalEntryRequest) XXX_Size() int {
	return xxx_messageInfo_ReadJournalEntryRequest.Size(m)
}
func (m *ReadJournalEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadJournalEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadJournalEntryRequest proto.InternalMessageInfo

func (m *ReadJournalEntryRequest) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

type ReadJournalEntryResponse struct {
	Result               *JournalEntry `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadJournalEntryResponse) Reset()         { *m = ReadJournalEntryResponse{} }
func (m *ReadJournalEntryResponse) String() string { return proto.CompactTextString(m) }
func (*ReadJournalEntryResponse) ProtoMessage()    {}
func (*ReadJournalEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{6}
}

func (m *ReadJournalEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadJournalEntryResponse.Unmarshal(m, b)
}
func (m *ReadJournalEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadJournalEntryResponse.Marshal(b, m, deterministic)
}
func (m *ReadJournalEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadJournalEntryResponse.Merge(m, src)
}
func (m *ReadJournalEntryResponse) XXX_Size() int {
	return xxx_messageInfo_ReadJournalEntryResponse.Size(m)
}
func (m *ReadJournalEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadJournalEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadJournalEntryResponse proto.InternalMessageInfo

func (m *ReadJournalEntryResponse) GetResult() *JournalEntry {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateJournalEntryRequest struct {
	Payload              *JournalEntry `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateJournalEntryRequest) Reset()         { *m = UpdateJournalEntryRequest{} }
func (m *UpdateJournalEntryRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateJournalEntryRequest) ProtoMessage()    {}
func (*UpdateJournalEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{7}
}

func (m *UpdateJournalEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJournalEntryRequest.Unmarshal(m, b)
}
func (m *UpdateJournalEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJournalEntryRequest.Marshal(b, m, deterministic)
}
func (m *UpdateJournalEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJournalEntryRequest.Merge(m, src)
}
func (m *UpdateJournalEntryRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateJournalEntryRequest.Size(m)
}
func (m *UpdateJournalEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJournalEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJournalEntryRequest proto.InternalMessageInfo

func (m *UpdateJournalEntryRequest) GetPayload() *JournalEntry {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UpdateJournalEntryResponse struct {
	Result               *JournalEntry `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateJournalEntryResponse) Reset()         { *m = UpdateJournalEntryResponse{} }
func (m *UpdateJournalEntryResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateJournalEntryResponse) ProtoMessage()    {}
func (*UpdateJournalEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{8}
}

func (m *UpdateJournalEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJournalEntryResponse.Unmarshal(m, b)
}
func (m *UpdateJournalEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJournalEntryResponse.Marshal(b, m, deterministic)
}
func (m *UpdateJournalEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJournalEntryResponse.Merge(m, src)
}
func (m *UpdateJournalEntryResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateJournalEntryResponse.Size(m)
}
func (m *UpdateJournalEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJournalEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJournalEntryResponse proto.InternalMessageInfo

func (m *UpdateJournalEntryResponse) GetResult() *JournalEntry {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteJournalEntryRequest struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteJournalEntryRequest) Reset()         { *m = DeleteJournalEntryRequest{} }
func (m *DeleteJournalEntryRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteJournalEntryRequest) ProtoMessage()    {}
func (*DeleteJournalEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{9}
}

func (m *DeleteJournalEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteJournalEntryRequest.Unmarshal(m, b)
}
func (m *DeleteJournalEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteJournalEntryRequest.Marshal(b, m, deterministic)
}
func (m *DeleteJournalEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteJournalEntryRequest.Merge(m, src)
}
func (m *DeleteJournalEntryRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteJournalEntryRequest.Size(m)
}
func (m *DeleteJournalEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteJournalEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteJournalEntryRequest proto.InternalMessageInfo

func (m *DeleteJournalEntryRequest) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

type DeleteJournalEntryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteJournalEntryResponse) Reset()         { *m = DeleteJournalEntryResponse{} }
func (m *DeleteJournalEntryResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteJournalEntryResponse) ProtoMessage()    {}
func (*DeleteJournalEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{10}
}

func (m *DeleteJournalEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteJournalEntryResponse.Unmarshal(m, b)
}
func (m *DeleteJournalEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteJournalEntryResponse.Marshal(b, m, deterministic)
}
func (m *DeleteJournalEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteJournalEntryResponse.Merge(m, src)
}
func (m *DeleteJournalEntryResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteJournalEntryResponse.Size(m)
}
func (m *DeleteJournalEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteJournalEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteJournalEntryResponse proto.InternalMessageInfo

type ListJournalEntryRequest struct {
	Filter               *query.Filtering      `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              *query.Sorting        `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Fields               *query.FieldSelection `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	Paging               *query.Pagination     `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListJournalEntryRequest) Reset()         { *m = ListJournalEntryRequest{} }
func (m *ListJournalEntryRequest) String() string { return proto.CompactTextString(m) }
func (*ListJournalEntryRequest) ProtoMessage()    {}
func (*ListJournalEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{11}
}

func (m *ListJournalEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJournalEntryRequest.Unmarshal(m, b)
}
func (m *ListJournalEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJournalEntryRequest.Marshal(b, m, deterministic)
}
func (m *ListJournalEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJournalEntryRequest.Merge(m, src)
}
func (m *ListJournalEntryRequest) XXX_Size() int {
	return xxx_messageInfo_ListJournalEntryRequest.Size(m)
}
func (m *ListJournalEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJournalEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJournalEntryRequest proto.InternalMessageInfo

func (m *ListJournalEntryRequest) GetFilter() *query.Filtering {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListJournalEntryRequest) GetOrderBy() *query.Sorting {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListJournalEntryRequest) GetFields() *query.FieldSelection {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListJournalEntryRequest) GetPaging() *query.Pagination {
	if m != nil {
		return m.Paging
	}
	return nil
}

type ListJournalEntryResponse struct {
	Results              []*JournalEntry `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListJournalEntryResponse) Reset()         { *m = ListJournalEntryResponse{} }
func (m *ListJournalEntryResponse) String() string { return proto.CompactTextString(m) }
func (*ListJournalEntryResponse) ProtoMessage()    {}
func (*ListJournalEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f294d657bf37c, []int{12}
}

func (m *ListJournalEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJournalEntryResponse.Unmarshal(m, b)
}
func (m *ListJournalEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJournalEntryResponse.Marshal(b, m, deterministic)
}
func (m *ListJournalEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJournalEntryResponse.Merge(m, src)
}
func (m *ListJournalEntryResponse) XXX_Size() int {
	return xxx_messageInfo_ListJournalEntryResponse.Size(m)
}
func (m *ListJournalEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJournalEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJournalEntryResponse proto.InternalMessageInfo

func (m *ListJournalEntryResponse) GetResults() []*JournalEntry {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterEnum("model.JournalEntry_Severity", JournalEntry_Severity_name, JournalEntry_Severity_value)
	proto.RegisterType((*JournalSubjectType)(nil), "model.JournalSubjectType")
	proto.RegisterType((*JournalSubject)(nil), "model.JournalSubject")
	proto.RegisterType((*JournalEntry)(nil), "model.JournalEntry")
	proto.RegisterType((*CreateJournalEntryRequest)(nil), "model.CreateJournalEntryRequest")
	proto.RegisterType((*CreateJournalEntryResponse)(nil), "model.CreateJournalEntryResponse")
	proto.RegisterType((*ReadJournalEntryRequest)(nil), "model.ReadJournalEntryRequest")
	proto.RegisterType((*ReadJournalEntryResponse)(nil), "model.ReadJournalEntryResponse")
	proto.RegisterType((*UpdateJournalEntryRequest)(nil), "model.UpdateJournalEntryRequest")
	proto.RegisterType((*UpdateJournalEntryResponse)(nil), "model.UpdateJournalEntryResponse")
	proto.RegisterType((*DeleteJournalEntryRequest)(nil), "model.DeleteJournalEntryRequest")
	proto.RegisterType((*DeleteJournalEntryResponse)(nil), "model.DeleteJournalEntryResponse")
	proto.RegisterType((*ListJournalEntryRequest)(nil), "model.ListJournalEntryRequest")
	proto.RegisterType((*ListJournalEntryResponse)(nil), "model.ListJournalEntryResponse")
}

func init() {
	proto.RegisterFile("github.com/kodesmil/ks-model/journal.proto", fileDescriptor_f15f294d657bf37c)
}

var fileDescriptor_f15f294d657bf37c = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
	0x17, 0xfe, 0x65, 0x3b, 0xb6, 0x7b, 0xfa, 0x23, 0x30, 0x38, 0x14, 0x91, 0x8d, 0x60, 0xf5, 0x0c,
	0x0c, 0x30, 0xd6, 0x59, 0x0c, 0xb2, 0x0d, 0x58, 0x73, 0xb5, 0x66, 0xcd, 0x5a, 0x07, 0x5d, 0x3b,
	0xc8, 0x1d, 0x86, 0xee, 0x26, 0xa0, 0xa5, 0x63, 0x8f, 0xb5, 0x2c, 0xb2, 0x14, 0x55, 0x54, 0xaf,
	0xb0, 0x37, 0x4a, 0xde, 0x65, 0xc0, 0x2e, 0xf6, 0x06, 0x7b, 0x81, 0x41, 0x14, 0x2d, 0x24, 0xb6,
	0xe5, 0x66, 0xc5, 0x80, 0x61, 0x37, 0x36, 0x2d, 0x7e, 0xdf, 0x77, 0x0e, 0xbf, 0x73, 0x8e, 0x68,
	0xf8, 0x6c, 0xce, 0xf5, 0x2f, 0xe9, 0xd4, 0x0b, 0xc4, 0x92, 0x2e, 0x44, 0x88, 0xc9, 0x92, 0x47,
	0x74, 0x91, 0x8c, 0x96, 0x22, 0xc4, 0x88, 0xbe, 0x16, 0xa9, 0x8a, 0x59, 0xe4, 0x49, 0x25, 0xb4,
	0x20, 0x7b, 0xe6, 0x61, 0xaf, 0x3f, 0x17, 0x62, 0x1e, 0x21, 0x35, 0x0f, 0xa7, 0xe9, 0x8c, 0xce,
	0x38, 0x46, 0xe1, 0xc5, 0x92, 0x25, 0x8b, 0x02, 0xd8, 0xbb, 0xbf, 0x8e, 0xd0, 0x7c, 0x89, 0x89,
	0x66, 0x4b, 0x69, 0x01, 0x87, 0x16, 0xc0, 0x24, 0xa7, 0x2c, 0x8e, 0x85, 0x66, 0x9a, 0x8b, 0x38,
	0xb1, 0xbb, 0x0f, 0xcc, 0x57, 0x30, 0x9a, 0x63, 0x3c, 0x7a, 0xcb, 0x22, 0x1e, 0x32, 0x8d, 0x74,
	0x63, 0x61, 0xc1, 0x27, 0xd7, 0x0e, 0xc0, 0xe3, 0x99, 0x98, 0x46, 0xe2, 0x9d, 0x90, 0x18, 0xd3,
	0x6b, 0x22, 0x73, 0xa1, 0x96, 0x54, 0x48, 0x13, 0x83, 0xe6, 0x3f, 0x2c, 0xf7, 0xe1, 0x6d, 0xb9,
	0x3a, 0x93, 0x98, 0x14, 0x9f, 0x96, 0x7a, 0x5e, 0x45, 0x65, 0x3a, 0x62, 0xc9, 0x88, 0x49, 0x39,
	0xd2, 0x42, 0x44, 0x0b, 0xae, 0xe9, 0x9b, 0x14, 0x55, 0x46, 0x03, 0x11, 0x45, 0x18, 0xe4, 0x29,
	0x5c, 0x08, 0x89, 0x8a, 0x69, 0xa1, 0x56, 0x5a, 0x67, 0xb7, 0xd7, 0x52, 0x32, 0xa0, 0x0a, 0x13,
	0x91, 0xaa, 0x00, 0xcb, 0x45, 0x21, 0x33, 0xf8, 0xcd, 0x01, 0x72, 0x5e, 0x14, 0x6c, 0x92, 0x4e,
	0x5f, 0x63, 0xa0, 0x5f, 0x66, 0x12, 0xc9, 0x57, 0x50, 0xe3, 0xa1, 0xeb, 0xf4, 0x9d, 0xe1, 0xdd,
	0xe3, 0x7b, 0x9e, 0x91, 0xf4, 0x94, 0x0c, 0xbc, 0x71, 0x88, 0xb1, 0xe6, 0x33, 0x8e, 0xea, 0x74,
	0xff, 0xea, 0xb2, 0x0b, 0xd0, 0x26, 0x0d, 0x8d, 0xef, 0xf4, 0xd0, 0xf1, 0x6b, 0x3c, 0x24, 0x0f,
	0x01, 0x02, 0x85, 0x4c, 0x63, 0x78, 0xc1, 0xb4, 0x5b, 0x33, 0xf4, 0x9e, 0x57, 0xd4, 0xcd, 0x5b,
	0x15, 0xd6, 0x7b, 0xb9, 0x2a, 0xac, 0x7f, 0xc7, 0xa2, 0x1f, 0xe9, 0x9c, 0x9a, 0xca, 0x70, 0x45,
	0xad, 0xbf, 0x9f, 0x6a, 0xd1, 0x8f, 0x34, 0x21, 0xd0, 0x88, 0xd9, 0x12, 0xdd, 0x46, 0xdf, 0x19,
	0xde, 0xf1, 0xcd, 0xfa, 0xa4, 0x79, 0x75, 0xd9, 0xad, 0xb5, 0x9d, 0xc1, 0xaf, 0x35, 0xd8, 0xbf,
	0x79, 0xbe, 0xff, 0xfc, 0xd9, 0x08, 0x85, 0x3d, 0xd3, 0x55, 0xee, 0x5e, 0xbf, 0x3e, 0xbc, 0x7b,
	0xdc, 0xf5, 0xcc, 0x88, 0x79, 0x9b, 0x65, 0xf4, 0x0b, 0x5c, 0x69, 0xc6, 0x9f, 0x35, 0xf8, 0xbf,
	0x45, 0x9d, 0xc5, 0x5a, 0x65, 0x7f, 0xcb, 0x8a, 0x34, 0xe5, 0xe1, 0xbf, 0x6a, 0xc5, 0xe7, 0x50,
	0x0f, 0x59, 0x66, 0x9c, 0xd8, 0xcd, 0xc9, 0x61, 0xe4, 0x6b, 0x68, 0x27, 0xf8, 0x16, 0x15, 0xd7,
	0x99, 0xbb, 0xd7, 0x77, 0x86, 0xfb, 0xc7, 0x87, 0x37, 0x7d, 0x32, 0x0e, 0x78, 0x13, 0x8b, 0xf1,
	0x4b, 0xb4, 0xb1, 0x5c, 0x68, 0x74, 0x5b, 0xd6, 0x72, 0xa1, 0x71, 0x70, 0x04, 0xed, 0x15, 0x92,
	0xb4, 0xa1, 0xf1, 0xfc, 0xc5, 0xf3, 0xb3, 0xce, 0xff, 0x48, 0x0b, 0xea, 0xcf, 0x5e, 0xfc, 0xd4,
	0x71, 0xf2, 0xc5, 0xf7, 0xe3, 0xc7, 0x9d, 0x5a, 0xbe, 0xf7, 0x74, 0xfc, 0xe4, 0x69, 0xa7, 0x7e,
	0xd2, 0xbe, 0xba, 0xec, 0x36, 0xda, 0x4e, 0xdf, 0x19, 0x9c, 0x43, 0xf7, 0x5b, 0x73, 0xfe, 0xeb,
	0x81, 0x7d, 0x7c, 0x93, 0x62, 0xa2, 0xc9, 0x08, 0x5a, 0x92, 0x65, 0x91, 0x60, 0xab, 0x32, 0x7c,
	0xb4, 0x25, 0x4b, 0x7f, 0x85, 0x19, 0x8c, 0xa1, 0xb7, 0x4d, 0x2b, 0x91, 0x22, 0x4e, 0x90, 0x3c,
	0x80, 0xa6, 0xc2, 0x24, 0x8d, 0xf4, 0x2e, 0x2d, 0x0b, 0x19, 0x7c, 0x03, 0x07, 0x3e, 0xb2, 0x70,
	0x5b, 0x52, 0x9f, 0xbe, 0xb7, 0x2d, 0xf2, 0x36, 0x18, 0x3c, 0x01, 0x77, 0x53, 0xe1, 0x43, 0x52,
	0x39, 0x87, 0xee, 0x8f, 0xa6, 0xcc, 0xff, 0x8c, 0x43, 0xdb, 0xb4, 0x3e, 0x24, 0xad, 0x53, 0xe8,
	0x3e, 0xc6, 0x08, 0xb7, 0xa7, 0x75, 0x4b, 0x8f, 0x0e, 0xa1, 0xb7, 0x4d, 0xa3, 0x48, 0x67, 0xf0,
	0x87, 0x03, 0x07, 0xcf, 0x78, 0xa2, 0xb7, 0x05, 0xa0, 0xd0, 0x9c, 0xf1, 0x48, 0xa3, 0xb2, 0x41,
	0x0e, 0xbc, 0xd5, 0x6b, 0xde, 0x63, 0x92, 0x7b, 0xdf, 0x99, 0x3d, 0x1e, 0xcf, 0x7d, 0x0b, 0x23,
	0x47, 0xd0, 0x16, 0x2a, 0x44, 0x75, 0x31, 0xcd, 0xec, 0x4c, 0xde, 0xbb, 0x49, 0x99, 0x08, 0xa5,
	0x73, 0x42, 0xcb, 0xc0, 0x4e, 0x33, 0xf2, 0x65, 0x1e, 0x02, 0xa3, 0x30, 0xb1, 0x83, 0x78, 0xb8,
	0x1e, 0x02, 0xa3, 0x70, 0x82, 0xf6, 0x06, 0xf2, 0x2d, 0x96, 0x1c, 0x41, 0x53, 0xb2, 0x39, 0x8f,
	0xe7, 0x76, 0x14, 0xdd, 0x9b, 0xac, 0x1f, 0xf2, 0x3d, 0x56, 0x30, 0x0a, 0xdc, 0x60, 0x0c, 0xee,
	0xe6, 0x29, 0x6d, 0x45, 0x46, 0xd0, 0x2a, 0xec, 0x4e, 0x5c, 0xc7, 0xbc, 0xce, 0xb6, 0x97, 0xd7,
	0x62, 0x8e, 0x7f, 0xaf, 0x43, 0xcb, 0xee, 0x90, 0x57, 0x40, 0x36, 0x87, 0x81, 0xf4, 0x2d, 0xbf,
	0x72, 0xe6, 0x7a, 0x9f, 0xec, 0x40, 0xd8, 0xac, 0x26, 0xd0, 0x59, 0x6f, 0x6d, 0xf2, 0xb1, 0xa5,
	0x55, 0x4c, 0x4d, 0xef, 0x7e, 0xe5, 0xbe, 0x15, 0x7d, 0x05, 0x64, 0xb3, 0x35, 0xcb, 0x7c, 0x2b,
	0x27, 0xa0, 0xcc, 0x77, 0x47, 0x5f, 0x2f, 0x81, 0x6c, 0xb6, 0x59, 0x29, 0x5d, 0xd9, 0xc5, 0xa5,
	0xf4, 0x8e, 0x1e, 0x25, 0x57, 0x97, 0xdd, 0xfd, 0xb5, 0x7b, 0x63, 0x02, 0x9d, 0xf5, 0x82, 0x96,
	0xf6, 0x54, 0xf4, 0x73, 0x69, 0x4f, 0x55, 0x27, 0xf4, 0xec, 0x2d, 0x75, 0x7a, 0xf4, 0xb3, 0x57,
	0xf1, 0xff, 0x72, 0xca, 0x82, 0x05, 0xc6, 0x21, 0xe5, 0xb1, 0xc6, 0x9c, 0x4f, 0xe5, 0x62, 0x4e,
	0xe5, 0x74, 0xda, 0x34, 0x97, 0xc0, 0x17, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xf7, 0x2c,
	0xc2, 0x98, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JournalClient is the client API for Journal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JournalClient interface {
	CreateJournalEntry(ctx context.Context, in *CreateJournalEntryRequest, opts ...grpc.CallOption) (*CreateJournalEntryResponse, error)
	ReadJournalEntry(ctx context.Context, in *ReadJournalEntryRequest, opts ...grpc.CallOption) (*ReadJournalEntryResponse, error)
	UpdateJournalEntry(ctx context.Context, in *UpdateJournalEntryRequest, opts ...grpc.CallOption) (*UpdateJournalEntryResponse, error)
	DeleteJournalEntry(ctx context.Context, in *DeleteJournalEntryRequest, opts ...grpc.CallOption) (*DeleteJournalEntryResponse, error)
	ListJournalEntry(ctx context.Context, in *ListJournalEntryRequest, opts ...grpc.CallOption) (*ListJournalEntryResponse, error)
}

type journalClient struct {
	cc *grpc.ClientConn
}

func NewJournalClient(cc *grpc.ClientConn) JournalClient {
	return &journalClient{cc}
}

func (c *journalClient) CreateJournalEntry(ctx context.Context, in *CreateJournalEntryRequest, opts ...grpc.CallOption) (*CreateJournalEntryResponse, error) {
	out := new(CreateJournalEntryResponse)
	err := c.cc.Invoke(ctx, "/model.Journal/CreateJournalEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalClient) ReadJournalEntry(ctx context.Context, in *ReadJournalEntryRequest, opts ...grpc.CallOption) (*ReadJournalEntryResponse, error) {
	out := new(ReadJournalEntryResponse)
	err := c.cc.Invoke(ctx, "/model.Journal/ReadJournalEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalClient) UpdateJournalEntry(ctx context.Context, in *UpdateJournalEntryRequest, opts ...grpc.CallOption) (*UpdateJournalEntryResponse, error) {
	out := new(UpdateJournalEntryResponse)
	err := c.cc.Invoke(ctx, "/model.Journal/UpdateJournalEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalClient) DeleteJournalEntry(ctx context.Context, in *DeleteJournalEntryRequest, opts ...grpc.CallOption) (*DeleteJournalEntryResponse, error) {
	out := new(DeleteJournalEntryResponse)
	err := c.cc.Invoke(ctx, "/model.Journal/DeleteJournalEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalClient) ListJournalEntry(ctx context.Context, in *ListJournalEntryRequest, opts ...grpc.CallOption) (*ListJournalEntryResponse, error) {
	out := new(ListJournalEntryResponse)
	err := c.cc.Invoke(ctx, "/model.Journal/ListJournalEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JournalServer is the server API for Journal service.
type JournalServer interface {
	CreateJournalEntry(context.Context, *CreateJournalEntryRequest) (*CreateJournalEntryResponse, error)
	ReadJournalEntry(context.Context, *ReadJournalEntryRequest) (*ReadJournalEntryResponse, error)
	UpdateJournalEntry(context.Context, *UpdateJournalEntryRequest) (*UpdateJournalEntryResponse, error)
	DeleteJournalEntry(context.Context, *DeleteJournalEntryRequest) (*DeleteJournalEntryResponse, error)
	ListJournalEntry(context.Context, *ListJournalEntryRequest) (*ListJournalEntryResponse, error)
}

func RegisterJournalServer(s *grpc.Server, srv JournalServer) {
	s.RegisterService(&_Journal_serviceDesc, srv)
}

func _Journal_CreateJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServer).CreateJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Journal/CreateJournalEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServer).CreateJournalEntry(ctx, req.(*CreateJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Journal_ReadJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServer).ReadJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Journal/ReadJournalEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServer).ReadJournalEntry(ctx, req.(*ReadJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Journal_UpdateJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServer).UpdateJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Journal/UpdateJournalEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServer).UpdateJournalEntry(ctx, req.(*UpdateJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Journal_DeleteJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServer).DeleteJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Journal/DeleteJournalEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServer).DeleteJournalEntry(ctx, req.(*DeleteJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Journal_ListJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServer).ListJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Journal/ListJournalEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServer).ListJournalEntry(ctx, req.(*ListJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Journal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Journal",
	HandlerType: (*JournalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJournalEntry",
			Handler:    _Journal_CreateJournalEntry_Handler,
		},
		{
			MethodName: "ReadJournalEntry",
			Handler:    _Journal_ReadJournalEntry_Handler,
		},
		{
			MethodName: "UpdateJournalEntry",
			Handler:    _Journal_UpdateJournalEntry_Handler,
		},
		{
			MethodName: "DeleteJournalEntry",
			Handler:    _Journal_DeleteJournalEntry_Handler,
		},
		{
			MethodName: "ListJournalEntry",
			Handler:    _Journal_ListJournalEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kodesmil/ks-model/journal.proto",
}
