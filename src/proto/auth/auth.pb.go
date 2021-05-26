// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

package pbAuth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserRegisterReq struct {
	UID                  string   `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Gender               int32    `protobuf:"varint,4,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Birth                string   `protobuf:"bytes,6,opt,name=Birth,proto3" json:"Birth,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=Email,proto3" json:"Email,omitempty"`
	Ex                   string   `protobuf:"bytes,8,opt,name=Ex,proto3" json:"Ex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterReq) Reset()         { *m = UserRegisterReq{} }
func (m *UserRegisterReq) String() string { return proto.CompactTextString(m) }
func (*UserRegisterReq) ProtoMessage()    {}
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{0}
}

func (m *UserRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterReq.Unmarshal(m, b)
}
func (m *UserRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterReq.Marshal(b, m, deterministic)
}
func (m *UserRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterReq.Merge(m, src)
}
func (m *UserRegisterReq) XXX_Size() int {
	return xxx_messageInfo_UserRegisterReq.Size(m)
}
func (m *UserRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterReq proto.InternalMessageInfo

func (m *UserRegisterReq) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *UserRegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRegisterReq) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserRegisterReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserRegisterReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserRegisterReq) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

func (m *UserRegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRegisterReq) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

type UserRegisterResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterResp) Reset()         { *m = UserRegisterResp{} }
func (m *UserRegisterResp) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResp) ProtoMessage()    {}
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{1}
}

func (m *UserRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResp.Unmarshal(m, b)
}
func (m *UserRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResp.Marshal(b, m, deterministic)
}
func (m *UserRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResp.Merge(m, src)
}
func (m *UserRegisterResp) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResp.Size(m)
}
func (m *UserRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResp proto.InternalMessageInfo

func (m *UserRegisterResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UserTokenReq struct {
	Platform             int32    `protobuf:"varint,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	UID                  string   `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenReq) Reset()         { *m = UserTokenReq{} }
func (m *UserTokenReq) String() string { return proto.CompactTextString(m) }
func (*UserTokenReq) ProtoMessage()    {}
func (*UserTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{2}
}

func (m *UserTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenReq.Unmarshal(m, b)
}
func (m *UserTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenReq.Marshal(b, m, deterministic)
}
func (m *UserTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenReq.Merge(m, src)
}
func (m *UserTokenReq) XXX_Size() int {
	return xxx_messageInfo_UserTokenReq.Size(m)
}
func (m *UserTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenReq proto.InternalMessageInfo

func (m *UserTokenReq) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *UserTokenReq) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

type UserTokenResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	ExpiredTime          int64    `protobuf:"varint,4,opt,name=ExpiredTime,proto3" json:"ExpiredTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenResp) Reset()         { *m = UserTokenResp{} }
func (m *UserTokenResp) String() string { return proto.CompactTextString(m) }
func (*UserTokenResp) ProtoMessage()    {}
func (*UserTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{3}
}

func (m *UserTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenResp.Unmarshal(m, b)
}
func (m *UserTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenResp.Marshal(b, m, deterministic)
}
func (m *UserTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenResp.Merge(m, src)
}
func (m *UserTokenResp) XXX_Size() int {
	return xxx_messageInfo_UserTokenResp.Size(m)
}
func (m *UserTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenResp proto.InternalMessageInfo

func (m *UserTokenResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *UserTokenResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserTokenResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserTokenResp) GetExpiredTime() int64 {
	if m != nil {
		return m.ExpiredTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRegisterReq)(nil), "pbAuth.UserRegisterReq")
	proto.RegisterType((*UserRegisterResp)(nil), "pbAuth.UserRegisterResp")
	proto.RegisterType((*UserTokenReq)(nil), "pbAuth.UserTokenReq")
	proto.RegisterType((*UserTokenResp)(nil), "pbAuth.UserTokenResp")
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor_712ec48c1eaf43a2) }

var fileDescriptor_712ec48c1eaf43a2 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd3, 0x42, 0x0b, 0x0c, 0x22, 0x64, 0x82, 0xba, 0xe1, 0x44, 0x7a, 0xe2, 0x60, 0x30,
	0xd1, 0x8b, 0x89, 0x5e, 0x40, 0x1b, 0xc3, 0x01, 0x63, 0x2a, 0x5c, 0xbc, 0x15, 0x58, 0xa1, 0x91,
	0xb2, 0x75, 0xb7, 0x24, 0x78, 0xf6, 0xa1, 0x7c, 0x3d, 0x33, 0xbb, 0x5b, 0x82, 0x86, 0x4b, 0x3b,
	0xff, 0xb7, 0x33, 0xd9, 0xfd, 0x67, 0x06, 0x9a, 0xf1, 0x36, 0x5f, 0x5d, 0xd1, 0xa7, 0x9f, 0x49,
	0x91, 0x0b, 0xf4, 0xb3, 0xd9, 0x60, 0x9b, 0xaf, 0x82, 0x1f, 0x07, 0x9a, 0x53, 0xc5, 0x65, 0xc4,
	0x97, 0x89, 0xca, 0xe9, 0xff, 0x89, 0x2d, 0x28, 0x4d, 0x47, 0x8f, 0xcc, 0xe9, 0x3a, 0xbd, 0x5a,
	0x44, 0x21, 0x22, 0x94, 0x9f, 0xe3, 0x94, 0x33, 0x57, 0x23, 0x1d, 0x13, 0x1b, 0xcd, 0xc5, 0x86,
	0x95, 0x0c, 0xa3, 0x18, 0xcf, 0xc1, 0x7f, 0xe2, 0x9b, 0x05, 0x97, 0xac, 0xdc, 0x75, 0x7a, 0x5e,
	0x64, 0x15, 0xf1, 0xb1, 0x98, 0x25, 0x6b, 0xce, 0x3c, 0x9d, 0x6d, 0x15, 0xb6, 0xc1, 0x1b, 0x26,
	0x32, 0x5f, 0x31, 0x5f, 0x63, 0x23, 0x88, 0x86, 0x69, 0x9c, 0xac, 0x59, 0xc5, 0x50, 0x2d, 0xf0,
	0x14, 0xdc, 0x70, 0xc7, 0xaa, 0x1a, 0xb9, 0xe1, 0x2e, 0xb8, 0x84, 0xd6, 0xdf, 0x87, 0xab, 0x0c,
	0x19, 0x54, 0x5e, 0xb7, 0xf3, 0x39, 0x57, 0x4a, 0xbf, 0xbe, 0x1a, 0x15, 0x32, 0xb8, 0x87, 0x13,
	0xca, 0x9e, 0x88, 0x0f, 0xbe, 0x21, 0x8f, 0x1d, 0xa8, 0xbe, 0xac, 0xe3, 0xfc, 0x5d, 0xc8, 0x54,
	0xa7, 0x7a, 0xd1, 0x5e, 0x17, 0xfe, 0xdd, 0xbd, 0xff, 0xe0, 0x0b, 0x1a, 0x07, 0xd5, 0xe6, 0xa2,
	0x50, 0xca, 0x07, 0xb1, 0xe0, 0xb6, 0xba, 0x90, 0x64, 0x35, 0x94, 0x72, 0xac, 0x96, 0xb6, 0xde,
	0x2a, 0x32, 0xa5, 0xcb, 0x6d, 0xbf, 0x8c, 0xc0, 0x2e, 0xd4, 0xc3, 0x5d, 0x96, 0x48, 0xbe, 0x98,
	0x24, 0x29, 0xd7, 0x5d, 0x2b, 0x45, 0x87, 0xe8, 0xfa, 0xdb, 0x81, 0x32, 0x4d, 0x0a, 0x07, 0xc6,
	0x41, 0xe1, 0x17, 0x2f, 0xfa, 0x66, 0x84, 0xfd, 0x7f, 0xe3, 0xeb, 0xb0, 0xe3, 0x07, 0x2a, 0xc3,
	0x5b, 0xa8, 0xed, 0x6d, 0x60, 0xfb, 0x30, 0xad, 0xe8, 0x4b, 0xe7, 0xec, 0x08, 0x55, 0xd9, 0xb0,
	0xf1, 0x56, 0xa7, 0xe5, 0xb9, 0x33, 0x87, 0x33, 0x5f, 0x2f, 0xd1, 0xcd, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x03, 0x0f, 0xb1, 0xb3, 0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, "/pbAuth.Auth/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := c.cc.Invoke(ctx, "/pbAuth.Auth/UserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) UserRegister(ctx context.Context, req *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (*UnimplementedAuthServer) UserToken(ctx context.Context, req *UserTokenReq) (*UserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbAuth.Auth/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbAuth.Auth/UserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbAuth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _Auth_UserRegister_Handler,
		},
		{
			MethodName: "UserToken",
			Handler:    _Auth_UserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
