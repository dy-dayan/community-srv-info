// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomicid.proto

package dayan_common_srv_atomicid

import (
	fmt "fmt"
	idl "github.com/dy-dayan/community-srv-info/idl"
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

type GetIDReq struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDReq) Reset()         { *m = GetIDReq{} }
func (m *GetIDReq) String() string { return proto.CompactTextString(m) }
func (*GetIDReq) ProtoMessage()    {}
func (*GetIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_99422b2753e559c3, []int{0}
}

func (m *GetIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDReq.Unmarshal(m, b)
}
func (m *GetIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDReq.Marshal(b, m, deterministic)
}
func (m *GetIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDReq.Merge(m, src)
}
func (m *GetIDReq) XXX_Size() int {
	return xxx_messageInfo_GetIDReq.Size(m)
}
func (m *GetIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDReq proto.InternalMessageInfo

func (m *GetIDReq) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type GetIDResp struct {
	BaseResp             *idl.Resp `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	Id                   int64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetIDResp) Reset()         { *m = GetIDResp{} }
func (m *GetIDResp) String() string { return proto.CompactTextString(m) }
func (*GetIDResp) ProtoMessage()    {}
func (*GetIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_99422b2753e559c3, []int{1}
}

func (m *GetIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDResp.Unmarshal(m, b)
}
func (m *GetIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDResp.Marshal(b, m, deterministic)
}
func (m *GetIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDResp.Merge(m, src)
}
func (m *GetIDResp) XXX_Size() int {
	return xxx_messageInfo_GetIDResp.Size(m)
}
func (m *GetIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDResp proto.InternalMessageInfo

func (m *GetIDResp) GetBaseResp() *idl.Resp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func (m *GetIDResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetIDReq)(nil), "dayan.common.srv.atomicid.GetIDReq")
	proto.RegisterType((*GetIDResp)(nil), "dayan.common.srv.atomicid.GetIDResp")
}

func init() { proto.RegisterFile("atomicid.proto", fileDescriptor_99422b2753e559c3) }

var fileDescriptor_99422b2753e559c3 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2c, 0xc9, 0xcf,
	0xcd, 0x4c, 0xce, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0x49, 0xac, 0x4c,
	0xcc, 0xd3, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0x29, 0x90,
	0xe2, 0x4a, 0x4a, 0x2c, 0x4e, 0x85, 0x28, 0x53, 0x52, 0xe0, 0xe2, 0x70, 0x4f, 0x2d, 0xf1, 0x74,
	0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe1, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x9c, 0xb9, 0x38, 0xa1, 0x2a, 0x8a, 0x0b, 0x84, 0xd4, 0xb8,
	0x38, 0x9c, 0x12, 0x8b, 0x53, 0x41, 0x6c, 0xb0, 0x2a, 0x6e, 0x23, 0x2e, 0x3d, 0xb0, 0x69, 0x20,
	0x91, 0x20, 0xb8, 0x9c, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x53, 0x66, 0x8a, 0x51, 0x1c, 0x17, 0x87, 0x23, 0xd8, 0x7a, 0x4f, 0x17, 0xa1, 0x20, 0x2e,
	0x56, 0xb0, 0x81, 0x42, 0xca, 0x7a, 0x38, 0xdd, 0xa8, 0x07, 0x73, 0x94, 0x94, 0x0a, 0x61, 0x45,
	0xc5, 0x05, 0x49, 0x6c, 0x60, 0xdf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x4f, 0x8a,
	0x18, 0x06, 0x01, 0x00, 0x00,
}