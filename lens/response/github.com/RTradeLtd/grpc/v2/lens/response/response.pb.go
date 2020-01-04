// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package response

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

type Index struct {
	// id is the identifier of the indexed object according to the lens system
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// data is miscellaneous data associated with the response
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// keywords are the keywords that you can use to search for this object within lens
	Keywords             []string `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Index) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Index) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type Results struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
}

func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type Object struct {
	// name is the "name" of the object, such as an IPFS content hash
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MimeType string `protobuf:"bytes,2,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// type is the type of the object, such as IPLD
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{2}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Object) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Object) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Object) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Index)(nil), "response.Index")
	proto.RegisterType((*Results)(nil), "response.Results")
	proto.RegisterType((*Object)(nil), "response.Object")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0xbb, 0xee, 0x6e, 0x47, 0x59, 0x24, 0xa7, 0xe0, 0xa9, 0xf4, 0x54, 0x16, 0x69,
	0x60, 0xc5, 0x3f, 0xe0, 0x45, 0x04, 0x41, 0x08, 0x7b, 0xf2, 0x96, 0x36, 0x43, 0xb7, 0xba, 0x6d,
	0x4a, 0x32, 0xab, 0xf6, 0xdf, 0x4b, 0xd2, 0x6d, 0x6f, 0xef, 0x9b, 0x49, 0xde, 0x7b, 0x0c, 0xec,
	0x2c, 0xba, 0xc1, 0xf4, 0x0e, 0xcb, 0xc1, 0x1a, 0x32, 0x6c, 0x3b, 0x73, 0xfe, 0x0a, 0x37, 0x6f,
	0xbd, 0xc6, 0x3f, 0xb6, 0x83, 0xb8, 0xd5, 0x3c, 0xca, 0xa2, 0x22, 0x95, 0x71, 0xab, 0x19, 0x83,
	0x95, 0x56, 0xa4, 0x78, 0x92, 0x45, 0xc5, 0x9d, 0x0c, 0x9a, 0x3d, 0xc0, 0xf6, 0x1b, 0xc7, 0x5f,
	0x63, 0xb5, 0xe3, 0x71, 0x96, 0x14, 0xa9, 0x5c, 0x38, 0x7f, 0x86, 0x8d, 0x44, 0x77, 0x39, 0x93,
	0x63, 0x7b, 0xd8, 0x98, 0xea, 0x0b, 0x6b, 0x72, 0x3c, 0xca, 0x92, 0xe2, 0xf6, 0x70, 0x5f, 0x2e,
	0xf9, 0x1f, 0x61, 0x21, 0xe7, 0x07, 0xf9, 0x09, 0xd6, 0xd3, 0xc8, 0x07, 0xf6, 0xaa, 0xc3, 0x6b,
	0x85, 0xa0, 0x7d, 0x60, 0xd7, 0x76, 0x78, 0x1c, 0x07, 0xe4, 0x71, 0x98, 0x2f, 0xec, 0x77, 0xb5,
	0x22, 0x6c, 0x8c, 0x1d, 0x43, 0xc9, 0x54, 0x2e, 0xec, 0xbd, 0xc8, 0xff, 0x59, 0x4d, 0x5e, 0x5e,
	0xbf, 0x3c, 0x7e, 0xee, 0x9b, 0x96, 0x4e, 0x97, 0xaa, 0xac, 0x4d, 0x27, 0xe4, 0xd1, 0x2a, 0x8d,
	0xef, 0xa4, 0x45, 0x63, 0x87, 0x5a, 0xfc, 0x1c, 0xc4, 0x19, 0x7b, 0x27, 0xe6, 0x9e, 0xd5, 0x3a,
	0x1c, 0xea, 0xe9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x50, 0xbb, 0xbf, 0x3a, 0x01, 0x00, 0x00,
}
