// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mgmt.proto

package mgmt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Server state in the system map.
type JoinResp_State int32

const (
	// Server in the system.
	JoinResp_IN JoinResp_State = 0
	// Server excluded from the system.
	JoinResp_OUT JoinResp_State = 1
)

var JoinResp_State_name = map[int32]string{
	0: "IN",
	1: "OUT",
}
var JoinResp_State_value = map[string]int32{
	"IN":  0,
	"OUT": 1,
}

func (x JoinResp_State) String() string {
	return proto.EnumName(JoinResp_State_name, int32(x))
}
func (JoinResp_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{1, 0}
}

type JoinReq struct {
	// Server UUID.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Server rank desired, if not -1.
	Rank uint32 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	// Server CaRT base URI (i.e., for context 0).
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// Server CaRT context count.
	Nctxs uint32 `protobuf:"varint,4,opt,name=nctxs,proto3" json:"nctxs,omitempty"`
	// Server management address.
	Addr                 string   `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinReq) Reset()         { *m = JoinReq{} }
func (m *JoinReq) String() string { return proto.CompactTextString(m) }
func (*JoinReq) ProtoMessage()    {}
func (*JoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{0}
}
func (m *JoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinReq.Unmarshal(m, b)
}
func (m *JoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinReq.Marshal(b, m, deterministic)
}
func (dst *JoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinReq.Merge(dst, src)
}
func (m *JoinReq) XXX_Size() int {
	return xxx_messageInfo_JoinReq.Size(m)
}
func (m *JoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinReq proto.InternalMessageInfo

func (m *JoinReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *JoinReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *JoinReq) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *JoinReq) GetNctxs() uint32 {
	if m != nil {
		return m.Nctxs
	}
	return 0
}

func (m *JoinReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type JoinResp struct {
	Status DaosRequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=mgmt.DaosRequestStatus" json:"status,omitempty"`
	// Server rank assigned.
	Rank                 uint32         `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	State                JoinResp_State `protobuf:"varint,3,opt,name=state,proto3,enum=mgmt.JoinResp_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinResp) Reset()         { *m = JoinResp{} }
func (m *JoinResp) String() string { return proto.CompactTextString(m) }
func (*JoinResp) ProtoMessage()    {}
func (*JoinResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{1}
}
func (m *JoinResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResp.Unmarshal(m, b)
}
func (m *JoinResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResp.Marshal(b, m, deterministic)
}
func (dst *JoinResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResp.Merge(dst, src)
}
func (m *JoinResp) XXX_Size() int {
	return xxx_messageInfo_JoinResp.Size(m)
}
func (m *JoinResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResp proto.InternalMessageInfo

func (m *JoinResp) GetStatus() DaosRequestStatus {
	if m != nil {
		return m.Status
	}
	return DaosRequestStatus_SUCCESS
}

func (m *JoinResp) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *JoinResp) GetState() JoinResp_State {
	if m != nil {
		return m.State
	}
	return JoinResp_IN
}

type GetAttachInfoReq struct {
	// System name. For daos_agent only.
	Sys                  string   `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttachInfoReq) Reset()         { *m = GetAttachInfoReq{} }
func (m *GetAttachInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoReq) ProtoMessage()    {}
func (*GetAttachInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{2}
}
func (m *GetAttachInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoReq.Unmarshal(m, b)
}
func (m *GetAttachInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoReq.Marshal(b, m, deterministic)
}
func (dst *GetAttachInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoReq.Merge(dst, src)
}
func (m *GetAttachInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoReq.Size(m)
}
func (m *GetAttachInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoReq proto.InternalMessageInfo

func (m *GetAttachInfoReq) GetSys() string {
	if m != nil {
		return m.Sys
	}
	return ""
}

type GetAttachInfoResp struct {
	Status DaosRequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=mgmt.DaosRequestStatus" json:"status,omitempty"`
	// CaRT PSRs of the system group.
	Psrs                 []*GetAttachInfoResp_Psr `protobuf:"bytes,2,rep,name=psrs,proto3" json:"psrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetAttachInfoResp) Reset()         { *m = GetAttachInfoResp{} }
func (m *GetAttachInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoResp) ProtoMessage()    {}
func (*GetAttachInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{3}
}
func (m *GetAttachInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoResp.Unmarshal(m, b)
}
func (m *GetAttachInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoResp.Marshal(b, m, deterministic)
}
func (dst *GetAttachInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoResp.Merge(dst, src)
}
func (m *GetAttachInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoResp.Size(m)
}
func (m *GetAttachInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoResp proto.InternalMessageInfo

func (m *GetAttachInfoResp) GetStatus() DaosRequestStatus {
	if m != nil {
		return m.Status
	}
	return DaosRequestStatus_SUCCESS
}

func (m *GetAttachInfoResp) GetPsrs() []*GetAttachInfoResp_Psr {
	if m != nil {
		return m.Psrs
	}
	return nil
}

// CaRT PSR.
type GetAttachInfoResp_Psr struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttachInfoResp_Psr) Reset()         { *m = GetAttachInfoResp_Psr{} }
func (m *GetAttachInfoResp_Psr) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoResp_Psr) ProtoMessage()    {}
func (*GetAttachInfoResp_Psr) Descriptor() ([]byte, []int) {
	return fileDescriptor_mgmt_3492d618868502ca, []int{3, 0}
}
func (m *GetAttachInfoResp_Psr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Unmarshal(m, b)
}
func (m *GetAttachInfoResp_Psr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Marshal(b, m, deterministic)
}
func (dst *GetAttachInfoResp_Psr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoResp_Psr.Merge(dst, src)
}
func (m *GetAttachInfoResp_Psr) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Size(m)
}
func (m *GetAttachInfoResp_Psr) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoResp_Psr.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoResp_Psr proto.InternalMessageInfo

func (m *GetAttachInfoResp_Psr) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *GetAttachInfoResp_Psr) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinReq)(nil), "mgmt.JoinReq")
	proto.RegisterType((*JoinResp)(nil), "mgmt.JoinResp")
	proto.RegisterType((*GetAttachInfoReq)(nil), "mgmt.GetAttachInfoReq")
	proto.RegisterType((*GetAttachInfoResp)(nil), "mgmt.GetAttachInfoResp")
	proto.RegisterType((*GetAttachInfoResp_Psr)(nil), "mgmt.GetAttachInfoResp.Psr")
	proto.RegisterEnum("mgmt.JoinResp_State", JoinResp_State_name, JoinResp_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MgmtSvcClient is the client API for MgmtSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MgmtSvcClient interface {
	// Join the server described by JoinReq to the system.
	Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (*JoinResp, error)
	// Get the information required by libdaos to attach to the system.
	GetAttachInfo(ctx context.Context, in *GetAttachInfoReq, opts ...grpc.CallOption) (*GetAttachInfoResp, error)
}

type mgmtSvcClient struct {
	cc *grpc.ClientConn
}

func NewMgmtSvcClient(cc *grpc.ClientConn) MgmtSvcClient {
	return &mgmtSvcClient{cc}
}

func (c *mgmtSvcClient) Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (*JoinResp, error) {
	out := new(JoinResp)
	err := c.cc.Invoke(ctx, "/mgmt.MgmtSvc/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtSvcClient) GetAttachInfo(ctx context.Context, in *GetAttachInfoReq, opts ...grpc.CallOption) (*GetAttachInfoResp, error) {
	out := new(GetAttachInfoResp)
	err := c.cc.Invoke(ctx, "/mgmt.MgmtSvc/GetAttachInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtSvcServer is the server API for MgmtSvc service.
type MgmtSvcServer interface {
	// Join the server described by JoinReq to the system.
	Join(context.Context, *JoinReq) (*JoinResp, error)
	// Get the information required by libdaos to attach to the system.
	GetAttachInfo(context.Context, *GetAttachInfoReq) (*GetAttachInfoResp, error)
}

func RegisterMgmtSvcServer(s *grpc.Server, srv MgmtSvcServer) {
	s.RegisterService(&_MgmtSvc_serviceDesc, srv)
}

func _MgmtSvc_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtSvcServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mgmt.MgmtSvc/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtSvcServer).Join(ctx, req.(*JoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtSvc_GetAttachInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtSvcServer).GetAttachInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mgmt.MgmtSvc/GetAttachInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtSvcServer).GetAttachInfo(ctx, req.(*GetAttachInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MgmtSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mgmt.MgmtSvc",
	HandlerType: (*MgmtSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _MgmtSvc_Join_Handler,
		},
		{
			MethodName: "GetAttachInfo",
			Handler:    _MgmtSvc_GetAttachInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mgmt.proto",
}

func init() { proto.RegisterFile("mgmt.proto", fileDescriptor_mgmt_3492d618868502ca) }

var fileDescriptor_mgmt_3492d618868502ca = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xed, 0x4a, 0xc3, 0x30,
	0x14, 0x5d, 0xfa, 0xb1, 0xb9, 0x2b, 0x1b, 0x35, 0x0c, 0x17, 0xea, 0x9f, 0x11, 0x04, 0x87, 0x42,
	0x07, 0xf5, 0x09, 0x14, 0x41, 0x26, 0xa8, 0x23, 0xd3, 0x07, 0xa8, 0x5b, 0x9d, 0x43, 0xda, 0xae,
	0xb9, 0xe9, 0xd0, 0x07, 0xf1, 0x05, 0x7c, 0x52, 0x49, 0x32, 0x87, 0x73, 0xf3, 0x8f, 0xff, 0x4e,
	0x92, 0xd3, 0x73, 0xce, 0x3d, 0xb7, 0x00, 0xd9, 0x2c, 0x53, 0xd1, 0x42, 0x16, 0xaa, 0xa0, 0x9e,
	0xc6, 0x61, 0x13, 0xe5, 0xd2, 0x5e, 0xf0, 0x0c, 0x1a, 0x37, 0xc5, 0x3c, 0x17, 0x69, 0x49, 0x29,
	0x78, 0x55, 0x35, 0x9f, 0x32, 0xd2, 0x23, 0xfd, 0xa6, 0x30, 0x58, 0xdf, 0xc9, 0x24, 0x7f, 0x65,
	0x4e, 0x8f, 0xf4, 0x5b, 0xc2, 0x60, 0x1a, 0x80, 0x5b, 0xc9, 0x39, 0x73, 0x0d, 0x4d, 0x43, 0xda,
	0x01, 0x3f, 0x9f, 0xa8, 0x37, 0x64, 0x9e, 0xa1, 0xd9, 0x83, 0xfe, 0x36, 0x99, 0x4e, 0x25, 0xf3,
	0xad, 0x9e, 0xc6, 0xfc, 0x83, 0xc0, 0x9e, 0xf5, 0xc3, 0x05, 0x1d, 0x40, 0x1d, 0x55, 0xa2, 0x2a,
	0x34, 0x96, 0xed, 0xb8, 0x1b, 0x99, 0xa4, 0x57, 0x49, 0x81, 0x22, 0x2d, 0xab, 0x14, 0xd5, 0xd8,
	0x3c, 0x8b, 0x15, 0x6d, 0x67, 0x9a, 0x53, 0xf0, 0xf5, 0x6b, 0x6a, 0xf2, 0xb4, 0xe3, 0x8e, 0xd5,
	0xf8, 0xf6, 0x88, 0xb4, 0x42, 0x2a, 0x2c, 0x85, 0x33, 0xf0, 0xcd, 0x99, 0xd6, 0xc1, 0x19, 0xde,
	0x05, 0x35, 0xda, 0x00, 0xf7, 0xfe, 0xf1, 0x21, 0x20, 0xfc, 0x18, 0x82, 0xeb, 0x54, 0x5d, 0x28,
	0x95, 0x4c, 0x5e, 0x86, 0xf9, 0x73, 0xa1, 0xfb, 0x08, 0xc0, 0xc5, 0x77, 0x5c, 0xd5, 0xa1, 0x21,
	0xff, 0x24, 0x70, 0xf0, 0x8b, 0xf6, 0x9f, 0x31, 0x06, 0xe0, 0x2d, 0x50, 0x22, 0x73, 0x7a, 0x6e,
	0x7f, 0x3f, 0x3e, 0xb2, 0xf4, 0x2d, 0xdd, 0x68, 0x84, 0x52, 0x18, 0x62, 0x78, 0x06, 0xee, 0x08,
	0xe5, 0x7a, 0x7c, 0xb2, 0xbd, 0x0c, 0x67, 0xbd, 0x8c, 0x78, 0x09, 0x8d, 0xdb, 0x59, 0xa6, 0xc6,
	0xcb, 0x09, 0x3d, 0x01, 0x4f, 0x17, 0x41, 0x5b, 0x3f, 0x4b, 0x29, 0xc3, 0xf6, 0x66, 0x47, 0xbc,
	0x46, 0x2f, 0xa1, 0xb5, 0xe1, 0x4f, 0x0f, 0x77, 0x86, 0x2a, 0xc3, 0xee, 0x1f, 0x61, 0x79, 0xed,
	0xa9, 0x6e, 0x7e, 0xa8, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x54, 0xa6, 0x10, 0x6f,
	0x02, 0x00, 0x00,
}
