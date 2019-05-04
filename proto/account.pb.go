// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Code                 uint32         `protobuf:"fixed32,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *AccountResult `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetData() *AccountResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type AccountResult struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ActualName           string   `protobuf:"bytes,4,opt,name=actualName,proto3" json:"actualName,omitempty"`
	CreatedByUserId      uint32   `protobuf:"varint,5,opt,name=createdByUserId,proto3" json:"createdByUserId,omitempty"`
	Department           string   `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	RoleId               string   `protobuf:"bytes,9,opt,name=roleId,proto3" json:"roleId,omitempty"`
	IsAdmin              uint32   `protobuf:"varint,10,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	Status               uint32   `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            int64    `protobuf:"varint,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,13,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Valid                int32    `protobuf:"varint,14,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResult) Reset()         { *m = AccountResult{} }
func (m *AccountResult) String() string { return proto.CompactTextString(m) }
func (*AccountResult) ProtoMessage()    {}
func (*AccountResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *AccountResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResult.Unmarshal(m, b)
}
func (m *AccountResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResult.Marshal(b, m, deterministic)
}
func (m *AccountResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResult.Merge(m, src)
}
func (m *AccountResult) XXX_Size() int {
	return xxx_messageInfo_AccountResult.Size(m)
}
func (m *AccountResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResult.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResult proto.InternalMessageInfo

func (m *AccountResult) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AccountResult) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *AccountResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountResult) GetActualName() string {
	if m != nil {
		return m.ActualName
	}
	return ""
}

func (m *AccountResult) GetCreatedByUserId() uint32 {
	if m != nil {
		return m.CreatedByUserId
	}
	return 0
}

func (m *AccountResult) GetDepartment() string {
	if m != nil {
		return m.Department
	}
	return ""
}

func (m *AccountResult) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountResult) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AccountResult) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *AccountResult) GetIsAdmin() uint32 {
	if m != nil {
		return m.IsAdmin
	}
	return 0
}

func (m *AccountResult) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AccountResult) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AccountResult) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AccountResult) GetValid() int32 {
	if m != nil {
		return m.Valid
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "account.LoginRequest")
	proto.RegisterType((*Response)(nil), "account.Response")
	proto.RegisterType((*AccountResult)(nil), "account.AccountResult")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x5f, 0x8b, 0xd4, 0x30,
	0x14, 0xc5, 0x6d, 0xa7, 0x9d, 0xce, 0xdc, 0x9d, 0xae, 0x18, 0x75, 0x09, 0x8b, 0x48, 0x29, 0x08,
	0xc5, 0x87, 0x01, 0x77, 0x51, 0x7c, 0x12, 0x66, 0x1e, 0x84, 0x05, 0x91, 0x25, 0xe0, 0x07, 0x88,
	0xcd, 0x65, 0x27, 0xd0, 0x36, 0xb5, 0x49, 0x94, 0xfd, 0xe8, 0xbe, 0x49, 0xfe, 0xb4, 0xbb, 0xee,
	0x53, 0x73, 0x7e, 0x27, 0xf7, 0xe6, 0x26, 0x3d, 0x50, 0xf2, 0xb6, 0x55, 0x76, 0x30, 0xfb, 0x71,
	0x52, 0x46, 0x91, 0x22, 0xca, 0xfa, 0x2b, 0xec, 0xbe, 0xa9, 0x3b, 0x39, 0x30, 0xfc, 0x65, 0x51,
	0x1b, 0x72, 0x09, 0x1b, 0xab, 0x71, 0x1a, 0x78, 0x8f, 0x34, 0xa9, 0x92, 0x66, 0xcb, 0x16, 0xed,
	0xbc, 0x91, 0x6b, 0xfd, 0x47, 0x4d, 0x82, 0xa6, 0xc1, 0x9b, 0x75, 0x2d, 0x60, 0xc3, 0x50, 0x8f,
	0x6a, 0xd0, 0x48, 0x08, 0x64, 0xad, 0x12, 0xa1, 0xbe, 0x60, 0x7e, 0x4d, 0x28, 0x14, 0x3d, 0x6a,
	0xcd, 0xef, 0x30, 0x96, 0xce, 0x92, 0xbc, 0x87, 0x4c, 0x70, 0xc3, 0xe9, 0xaa, 0x4a, 0x9a, 0xb3,
	0xab, 0x8b, 0xfd, 0x3c, 0xe8, 0x21, 0x7c, 0x19, 0x6a, 0xdb, 0x19, 0xe6, 0xf7, 0xd4, 0x7f, 0x53,
	0x28, 0xff, 0xe3, 0xe4, 0x1c, 0x52, 0x29, 0xfc, 0x49, 0x19, 0x4b, 0xa5, 0x70, 0x67, 0x5b, 0x2b,
	0xe7, 0xf9, 0xfc, 0xda, 0x31, 0x7f, 0x9f, 0x55, 0x60, 0xfe, 0x2e, 0x6f, 0x01, 0x78, 0x6b, 0x2c,
	0xef, 0xbe, 0x3b, 0x27, 0xf3, 0xce, 0x23, 0x42, 0x1a, 0x78, 0xde, 0x4e, 0xc8, 0x0d, 0x8a, 0xe3,
	0xfd, 0x0f, 0x8d, 0xd3, 0x8d, 0xa0, 0x79, 0x95, 0x34, 0x25, 0x7b, 0x8a, 0x5d, 0x27, 0x81, 0x23,
	0x9f, 0x4c, 0x8f, 0x83, 0xa1, 0xeb, 0xd0, 0xe9, 0x81, 0x90, 0x57, 0x90, 0x63, 0xcf, 0x65, 0x47,
	0x0b, 0x6f, 0x05, 0xe1, 0xe8, 0x78, 0x52, 0x03, 0xd2, 0x4d, 0xa0, 0x5e, 0x90, 0x0b, 0x58, 0x4f,
	0xaa, 0xc3, 0x1b, 0x41, 0xb7, 0x1e, 0x47, 0xe5, 0x5e, 0x4f, 0xea, 0x83, 0xe8, 0xe5, 0x40, 0xc1,
	0x4f, 0x31, 0x4b, 0x57, 0xa1, 0x0d, 0x37, 0x56, 0xd3, 0x33, 0x6f, 0x44, 0x45, 0xde, 0xc0, 0x36,
	0x0e, 0x7a, 0x30, 0x74, 0x57, 0x25, 0xcd, 0x8a, 0x3d, 0x00, 0xe7, 0xda, 0x51, 0x44, 0xb7, 0x0c,
	0xee, 0x02, 0xdc, 0x6c, 0xbf, 0x79, 0x27, 0x05, 0x3d, 0xaf, 0x92, 0x26, 0x67, 0x41, 0x5c, 0x7d,
	0x81, 0x22, 0x3e, 0x3d, 0xb9, 0x86, 0xdc, 0x87, 0x86, 0xbc, 0x5e, 0xfe, 0xd6, 0xe3, 0x10, 0x5d,
	0xbe, 0x58, 0xf0, 0x9c, 0x89, 0xfa, 0xd9, 0xf1, 0x1d, 0xbc, 0x6c, 0x55, 0xbf, 0x3f, 0x9d, 0xee,
	0x3f, 0x7e, 0xfe, 0xf4, 0x61, 0xde, 0x71, 0xdc, 0xc5, 0xa6, 0xb7, 0x2e, 0x97, 0xb7, 0xc9, 0xcf,
	0xb5, 0x0f, 0xe8, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xc0, 0xa8, 0x89, 0xb1, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/account.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Login(context.Context, *LoginRequest) (*Response, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Login(ctx context.Context, req *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
