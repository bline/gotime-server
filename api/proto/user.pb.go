// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmailOrId struct {
	// Types that are valid to be assigned to UserIdent:
	//	*EmailOrId_Email
	//	*EmailOrId_UserID
	UserIdent isEmailOrId_UserIdent `protobuf_oneof:"UserIdent"`
}

func (m *EmailOrId) Reset()                    { *m = EmailOrId{} }
func (m *EmailOrId) String() string            { return proto.CompactTextString(m) }
func (*EmailOrId) ProtoMessage()               {}
func (*EmailOrId) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isEmailOrId_UserIdent interface {
	isEmailOrId_UserIdent()
}

type EmailOrId_Email struct {
	Email string `protobuf:"bytes,1,opt,name=Email,oneof"`
}
type EmailOrId_UserID struct {
	UserID int64 `protobuf:"varint,2,opt,name=UserID,oneof"`
}

func (*EmailOrId_Email) isEmailOrId_UserIdent()  {}
func (*EmailOrId_UserID) isEmailOrId_UserIdent() {}

func (m *EmailOrId) GetUserIdent() isEmailOrId_UserIdent {
	if m != nil {
		return m.UserIdent
	}
	return nil
}

func (m *EmailOrId) GetEmail() string {
	if x, ok := m.GetUserIdent().(*EmailOrId_Email); ok {
		return x.Email
	}
	return ""
}

func (m *EmailOrId) GetUserID() int64 {
	if x, ok := m.GetUserIdent().(*EmailOrId_UserID); ok {
		return x.UserID
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EmailOrId) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EmailOrId_OneofMarshaler, _EmailOrId_OneofUnmarshaler, _EmailOrId_OneofSizer, []interface{}{
		(*EmailOrId_Email)(nil),
		(*EmailOrId_UserID)(nil),
	}
}

func _EmailOrId_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EmailOrId)
	// UserIdent
	switch x := m.UserIdent.(type) {
	case *EmailOrId_Email:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Email)
	case *EmailOrId_UserID:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.UserID))
	case nil:
	default:
		return fmt.Errorf("EmailOrId.UserIdent has unexpected type %T", x)
	}
	return nil
}

func _EmailOrId_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EmailOrId)
	switch tag {
	case 1: // UserIdent.Email
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.UserIdent = &EmailOrId_Email{x}
		return true, err
	case 2: // UserIdent.UserID
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.UserIdent = &EmailOrId_UserID{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _EmailOrId_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EmailOrId)
	// UserIdent
	switch x := m.UserIdent.(type) {
	case *EmailOrId_Email:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Email)))
		n += len(x.Email)
	case *EmailOrId_UserID:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.UserID))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LockUserRequest struct {
	Ident *EmailOrId `protobuf:"bytes,1,opt,name=Ident" json:"Ident,omitempty"`
}

func (m *LockUserRequest) Reset()                    { *m = LockUserRequest{} }
func (m *LockUserRequest) String() string            { return proto.CompactTextString(m) }
func (*LockUserRequest) ProtoMessage()               {}
func (*LockUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *LockUserRequest) GetIdent() *EmailOrId {
	if m != nil {
		return m.Ident
	}
	return nil
}

type DisableUserRequest struct {
	Ident *EmailOrId `protobuf:"bytes,1,opt,name=Ident" json:"Ident,omitempty"`
}

func (m *DisableUserRequest) Reset()                    { *m = DisableUserRequest{} }
func (m *DisableUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DisableUserRequest) ProtoMessage()               {}
func (*DisableUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DisableUserRequest) GetIdent() *EmailOrId {
	if m != nil {
		return m.Ident
	}
	return nil
}

type DeleteUserRequest struct {
	Ident *EmailOrId `protobuf:"bytes,1,opt,name=Ident" json:"Ident,omitempty"`
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DeleteUserRequest) GetIdent() *EmailOrId {
	if m != nil {
		return m.Ident
	}
	return nil
}

type DeleteUserResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *DeleteUserResponse) Reset()                    { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()               {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DeleteUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetUserRequest struct {
	Ident *EmailOrId `protobuf:"bytes,1,opt,name=Ident" json:"Ident,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GetUserRequest) GetIdent() *EmailOrId {
	if m != nil {
		return m.Ident
	}
	return nil
}

type GetUsersRequest struct {
	SearchOptions *SearchOptions `protobuf:"bytes,1,opt,name=SearchOptions" json:"SearchOptions,omitempty"`
	IsDeleted     bool           `protobuf:"varint,2,opt,name=IsDeleted" json:"IsDeleted,omitempty"`
	IsDisabled    bool           `protobuf:"varint,3,opt,name=IsDisabled" json:"IsDisabled,omitempty"`
	IsLockedOut   bool           `protobuf:"varint,4,opt,name=IsLockedOut" json:"IsLockedOut,omitempty"`
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *GetUsersRequest) GetSearchOptions() *SearchOptions {
	if m != nil {
		return m.SearchOptions
	}
	return nil
}

func (m *GetUsersRequest) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *GetUsersRequest) GetIsDisabled() bool {
	if m != nil {
		return m.IsDisabled
	}
	return false
}

func (m *GetUsersRequest) GetIsLockedOut() bool {
	if m != nil {
		return m.IsLockedOut
	}
	return false
}

type User struct {
	ID          int64                      `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Email       string                     `protobuf:"bytes,2,opt,name=Email" json:"Email,omitempty"`
	Token       string                     `protobuf:"bytes,3,opt,name=Token" json:"Token,omitempty"`
	LastLogin   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=LastLogin" json:"LastLogin,omitempty"`
	LockedOut   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=LockedOut" json:"LockedOut,omitempty"`
	IsDisabled  bool                       `protobuf:"varint,6,opt,name=IsDisabled" json:"IsDisabled,omitempty"`
	IsDeleted   bool                       `protobuf:"varint,7,opt,name=IsDeleted" json:"IsDeleted,omitempty"`
	IsLockedOut bool                       `protobuf:"varint,8,opt,name=IsLockedOut" json:"IsLockedOut,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetLastLogin() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func (m *User) GetLockedOut() *google_protobuf.Timestamp {
	if m != nil {
		return m.LockedOut
	}
	return nil
}

func (m *User) GetIsDisabled() bool {
	if m != nil {
		return m.IsDisabled
	}
	return false
}

func (m *User) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *User) GetIsLockedOut() bool {
	if m != nil {
		return m.IsLockedOut
	}
	return false
}

func init() {
	proto.RegisterType((*EmailOrId)(nil), "api.EmailOrId")
	proto.RegisterType((*LockUserRequest)(nil), "api.LockUserRequest")
	proto.RegisterType((*DisableUserRequest)(nil), "api.DisableUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "api.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "api.DeleteUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "api.GetUserRequest")
	proto.RegisterType((*GetUsersRequest)(nil), "api.GetUsersRequest")
	proto.RegisterType((*User)(nil), "api.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Accounts service

type AccountsClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (Accounts_GetUsersClient, error)
	DisableUser(ctx context.Context, in *DisableUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/api.Accounts/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (Accounts_GetUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Accounts_serviceDesc.Streams[0], c.cc, "/api.Accounts/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Accounts_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type accountsGetUsersClient struct {
	grpc.ClientStream
}

func (x *accountsGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountsClient) DisableUser(ctx context.Context, in *DisableUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/api.Accounts/DisableUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/api.Accounts/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/api.Accounts/LockUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service

type AccountsServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GetUsers(*GetUsersRequest, Accounts_GetUsersServer) error
	DisableUser(context.Context, *DisableUserRequest) (*SimpleResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*SimpleResponse, error)
	LockUser(context.Context, *LockUserRequest) (*SimpleResponse, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Accounts/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsServer).GetUsers(m, &accountsGetUsersServer{stream})
}

type Accounts_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type accountsGetUsersServer struct {
	grpc.ServerStream
}

func (x *accountsGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _Accounts_DisableUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).DisableUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Accounts/DisableUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).DisableUser(ctx, req.(*DisableUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Accounts/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_LockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).LockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Accounts/LockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).LockUser(ctx, req.(*LockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Accounts_GetUser_Handler,
		},
		{
			MethodName: "DisableUser",
			Handler:    _Accounts_DisableUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Accounts_DeleteUser_Handler,
		},
		{
			MethodName: "LockUser",
			Handler:    _Accounts_LockUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _Accounts_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x92, 0xa6, 0xb1, 0x27, 0x90, 0x8a, 0x21, 0x2a, 0x96, 0x85, 0x20, 0xb2, 0x38, 0x94,
	0x03, 0x2e, 0x0a, 0x02, 0x5a, 0x38, 0x81, 0x8c, 0x88, 0x51, 0x50, 0xa4, 0x4d, 0xf9, 0x00, 0xc7,
	0x5e, 0x82, 0x55, 0xdb, 0x6b, 0xbc, 0xeb, 0x2f, 0x41, 0xe2, 0x27, 0xf8, 0x49, 0x94, 0x59, 0x3b,
	0x71, 0x9c, 0x1e, 0x28, 0x37, 0xcf, 0xcc, 0x7b, 0x9e, 0xd9, 0x37, 0x6f, 0x00, 0x4a, 0xc9, 0x0b,
	0x37, 0x2f, 0x84, 0x12, 0xd8, 0x0b, 0xf2, 0xd8, 0xbe, 0x17, 0x8a, 0x34, 0x15, 0x99, 0x4e, 0xd9,
	0x4f, 0xd7, 0x42, 0xac, 0x13, 0x7e, 0x41, 0xd1, 0xaa, 0xfc, 0x7e, 0xa1, 0xe2, 0x94, 0x4b, 0x15,
	0xa4, 0xb9, 0x06, 0x38, 0x5f, 0xc0, 0xfc, 0x94, 0x06, 0x71, 0xb2, 0x28, 0xfc, 0x08, 0xcf, 0xa0,
	0x4f, 0x81, 0xd5, 0x99, 0x74, 0xce, 0xcd, 0xd9, 0x11, 0xd3, 0x21, 0x5a, 0x70, 0xf2, 0x4d, 0xf2,
	0xc2, 0xf7, 0xac, 0xee, 0xa4, 0x73, 0xde, 0x9b, 0x1d, 0xb1, 0x2a, 0xfe, 0x38, 0x04, 0x93, 0xbe,
	0x22, 0x9e, 0x29, 0xe7, 0x2d, 0x9c, 0xce, 0x45, 0x78, 0xb3, 0x49, 0x30, 0xfe, 0xb3, 0xe4, 0x52,
	0xe1, 0x33, 0xe8, 0x53, 0x8d, 0xfe, 0x38, 0x9c, 0x8e, 0xdc, 0x20, 0x8f, 0xdd, 0x6d, 0x43, 0xa6,
	0x8b, 0xce, 0x3b, 0x40, 0x2f, 0x96, 0xc1, 0x2a, 0xe1, 0x77, 0xe7, 0x5e, 0xc1, 0x03, 0x8f, 0x27,
	0x5c, 0xfd, 0x07, 0x75, 0x06, 0xd8, 0xa4, 0xca, 0x5c, 0x64, 0x92, 0xa3, 0x05, 0x83, 0x65, 0x19,
	0x86, 0x5c, 0x4a, 0x62, 0x1b, 0xac, 0x0e, 0x37, 0x95, 0xaf, 0x5c, 0xca, 0x60, 0xcd, 0x49, 0x07,
	0x93, 0xd5, 0xa1, 0xf3, 0x06, 0x46, 0x9f, 0xb9, 0xba, 0xfb, 0x04, 0x7f, 0x3a, 0x70, 0x5a, 0x11,
	0x65, 0xcd, 0xbc, 0x84, 0xfb, 0x4b, 0x1e, 0x14, 0xe1, 0x8f, 0x45, 0xae, 0x62, 0x91, 0xc9, 0xea,
	0x0f, 0x48, 0x7f, 0xd8, 0xab, 0xb0, 0x7d, 0x20, 0x3e, 0x06, 0xd3, 0x97, 0xfa, 0x45, 0x11, 0x4d,
	0x68, 0xb0, 0x5d, 0x02, 0x9f, 0x00, 0xf8, 0xb2, 0x92, 0x39, 0xb2, 0x7a, 0x54, 0x6e, 0x64, 0x70,
	0x02, 0x43, 0x5f, 0x6e, 0xf6, 0xc7, 0xa3, 0x45, 0xa9, 0xac, 0x63, 0x02, 0x34, 0x53, 0xce, 0xef,
	0x2e, 0x1c, 0x6f, 0x46, 0xc5, 0x11, 0x74, 0x7d, 0x8f, 0xe6, 0xea, 0xb1, 0xae, 0xef, 0xe1, 0xb8,
	0xf6, 0x8d, 0x96, 0xa5, 0x72, 0xcd, 0x18, 0xfa, 0xd7, 0xe2, 0x86, 0x67, 0xd4, 0xcb, 0x64, 0x3a,
	0xc0, 0x4b, 0x30, 0xe7, 0x81, 0x54, 0x73, 0xb1, 0x8e, 0x33, 0x6a, 0x32, 0x9c, 0xda, 0xae, 0x76,
	0xa9, 0x5b, 0xbb, 0xd4, 0xbd, 0xae, 0x5d, 0xca, 0x76, 0x60, 0x62, 0x6e, 0xc7, 0xeb, 0xff, 0x03,
	0xb3, 0x06, 0xb7, 0x9e, 0x7e, 0x72, 0xf0, 0xf4, 0x3d, 0xe1, 0x06, 0x6d, 0xe1, 0x5a, 0xc2, 0x18,
	0x07, 0xc2, 0x4c, 0x7f, 0x75, 0xc1, 0xf8, 0x10, 0x86, 0xa2, 0xcc, 0x94, 0xc4, 0xe7, 0x30, 0xa8,
	0x56, 0x8a, 0x0f, 0x69, 0x67, 0xfb, 0xce, 0xb0, 0x4d, 0x4a, 0x52, 0xfd, 0x05, 0x18, 0xf5, 0xf6,
	0x71, 0xdc, 0xc4, 0xca, 0x43, 0xf0, 0xcb, 0x0e, 0xbe, 0x87, 0x61, 0xe3, 0x4c, 0xf0, 0x11, 0xd5,
	0x0e, 0x0f, 0xc7, 0xd6, 0x6d, 0x97, 0x71, 0x9a, 0x27, 0x7c, 0x6b, 0xeb, 0x2b, 0x80, 0x9d, 0xd9,
	0xf1, 0x4c, 0x73, 0xdb, 0x87, 0x73, 0x3b, 0xf5, 0x35, 0x18, 0xf5, 0x5d, 0x57, 0x63, 0xb6, 0xce,
	0xfc, 0x56, 0xda, 0xea, 0x84, 0x96, 0xf2, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x0e,
	0x79, 0x69, 0xa3, 0x04, 0x00, 0x00,
}
