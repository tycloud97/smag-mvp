// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/usersearch.proto

package proto

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

type User struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RealName             string   `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Bio                  string   `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	FollowingsCount      uint64   `protobuf:"varint,5,opt,name=followings_count,json=followingsCount,proto3" json:"followings_count,omitempty"`
	FollowersCount       uint64   `protobuf:"varint,6,opt,name=followers_count,json=followersCount,proto3" json:"followers_count,omitempty"`
	FollowingsUsers      []*User  `protobuf:"bytes,7,rep,name=followings_users,json=followingsUsers,proto3" json:"followings_users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_68e38b08d2cbe7af, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetFollowingsCount() uint64 {
	if m != nil {
		return m.FollowingsCount
	}
	return 0
}

func (m *User) GetFollowersCount() uint64 {
	if m != nil {
		return m.FollowersCount
	}
	return 0
}

func (m *User) GetFollowingsUsers() []*User {
	if m != nil {
		return m.FollowingsUsers
	}
	return nil
}

type UserSearchResponse struct {
	UserList             []*User  `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSearchResponse) Reset()         { *m = UserSearchResponse{} }
func (m *UserSearchResponse) String() string { return proto.CompactTextString(m) }
func (*UserSearchResponse) ProtoMessage()    {}
func (*UserSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_68e38b08d2cbe7af, []int{1}
}
func (m *UserSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchResponse.Unmarshal(m, b)
}
func (m *UserSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchResponse.Marshal(b, m, deterministic)
}
func (dst *UserSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchResponse.Merge(dst, src)
}
func (m *UserSearchResponse) XXX_Size() int {
	return xxx_messageInfo_UserSearchResponse.Size(m)
}
func (m *UserSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSearchResponse proto.InternalMessageInfo

func (m *UserSearchResponse) GetUserList() []*User {
	if m != nil {
		return m.UserList
	}
	return nil
}

type UserName struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserName) Reset()         { *m = UserName{} }
func (m *UserName) String() string { return proto.CompactTextString(m) }
func (*UserName) ProtoMessage()    {}
func (*UserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_68e38b08d2cbe7af, []int{2}
}
func (m *UserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserName.Unmarshal(m, b)
}
func (m *UserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserName.Marshal(b, m, deterministic)
}
func (dst *UserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserName.Merge(dst, src)
}
func (m *UserName) XXX_Size() int {
	return xxx_messageInfo_UserName.Size(m)
}
func (m *UserName) XXX_DiscardUnknown() {
	xxx_messageInfo_UserName.DiscardUnknown(m)
}

var xxx_messageInfo_UserName proto.InternalMessageInfo

func (m *UserName) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserSearchResponse)(nil), "proto.UserSearchResponse")
	proto.RegisterType((*UserName)(nil), "proto.UserName")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSearchServiceClient is the client API for UserSearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSearchServiceClient interface {
	GetUserWithUsername(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*UserSearchResponse, error)
}

type userSearchServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserSearchServiceClient(cc *grpc.ClientConn) UserSearchServiceClient {
	return &userSearchServiceClient{cc}
}

func (c *userSearchServiceClient) GetUserWithUsername(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*UserSearchResponse, error) {
	out := new(UserSearchResponse)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetUserWithUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSearchServiceServer is the server API for UserSearchService service.
type UserSearchServiceServer interface {
	GetUserWithUsername(context.Context, *UserName) (*UserSearchResponse, error)
}

func RegisterUserSearchServiceServer(s *grpc.Server, srv UserSearchServiceServer) {
	s.RegisterService(&_UserSearchService_serviceDesc, srv)
}

func _UserSearchService_GetUserWithUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).GetUserWithUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/GetUserWithUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).GetUserWithUsername(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserSearchService",
	HandlerType: (*UserSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserWithUsername",
			Handler:    _UserSearchService_GetUserWithUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/usersearch.proto",
}

func init() {
	proto.RegisterFile("api/proto/usersearch.proto", fileDescriptor_usersearch_68e38b08d2cbe7af)
}

var fileDescriptor_usersearch_68e38b08d2cbe7af = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x4d, 0x6b, 0x33, 0x05, 0x5b, 0xd7, 0xcb, 0x5a, 0x11, 0x42, 0x2e, 0x8d, 0x97,
	0x16, 0x2a, 0x78, 0x14, 0xc4, 0x83, 0x17, 0xf1, 0x90, 0x12, 0x3d, 0x86, 0x6d, 0x18, 0xed, 0xc2,
	0x36, 0x1b, 0x76, 0x37, 0xf5, 0xd7, 0x0b, 0x32, 0x9b, 0x6a, 0x53, 0xc4, 0xd3, 0x4e, 0xbe, 0xf7,
	0x66, 0x26, 0xf3, 0x60, 0x2a, 0x6a, 0xb9, 0xa8, 0x8d, 0x76, 0x7a, 0xd1, 0x58, 0x34, 0x16, 0x85,
	0x29, 0x37, 0x73, 0x0f, 0x58, 0xdf, 0x3f, 0xc9, 0x57, 0x00, 0x61, 0x6e, 0xd1, 0xb0, 0x2b, 0x88,
	0xc8, 0x53, 0x54, 0x62, 0x8b, 0x3c, 0x88, 0x83, 0x34, 0xca, 0x86, 0x04, 0x5e, 0xc4, 0x16, 0x49,
	0x34, 0x28, 0x54, 0x2b, 0x9e, 0xb4, 0x22, 0x01, 0x2f, 0x4e, 0xa0, 0xb7, 0x96, 0x9a, 0xf7, 0x3c,
	0xa6, 0x92, 0x5d, 0x03, 0x88, 0x9d, 0x70, 0xc2, 0x14, 0x8d, 0x51, 0x3c, 0xf4, 0x42, 0xd4, 0x92,
	0xdc, 0x28, 0x76, 0x03, 0x93, 0x77, 0xad, 0x94, 0xfe, 0x94, 0xd5, 0x87, 0x2d, 0x4a, 0xdd, 0x54,
	0x8e, 0xf7, 0xe3, 0x20, 0x0d, 0xb3, 0xf1, 0x81, 0x3f, 0x12, 0x66, 0x33, 0xd8, 0x23, 0x34, 0x3f,
	0xce, 0x81, 0x77, 0x9e, 0xfd, 0xe2, 0xd6, 0x78, 0x77, 0x34, 0xd3, 0x5f, 0xcb, 0x4f, 0xe3, 0x5e,
	0x3a, 0x5a, 0x8e, 0xda, 0x83, 0xe7, 0x74, 0x65, 0x77, 0x01, 0x7d, 0xdb, 0xe4, 0x1e, 0x18, 0x15,
	0x2b, 0x1f, 0x4d, 0x86, 0xb6, 0xd6, 0x95, 0x45, 0x96, 0xee, 0xc3, 0x50, 0xd2, 0x3a, 0x1e, 0xfc,
	0x1d, 0xe3, 0x93, 0x79, 0x96, 0xd6, 0x25, 0x33, 0x18, 0xe6, 0x9d, 0x94, 0xfe, 0x8d, 0x70, 0xf9,
	0x0a, 0xe7, 0x87, 0x45, 0x2b, 0x34, 0x3b, 0x59, 0x22, 0x7b, 0x80, 0x8b, 0x27, 0x74, 0xc4, 0xdf,
	0xa4, 0xdb, 0xd0, 0x4b, 0xbd, 0x6c, 0xdc, 0xd9, 0x45, 0xcd, 0xd3, 0xcb, 0x0e, 0x38, 0xfe, 0xd5,
	0xf5, 0xc0, 0x2b, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb6, 0xab, 0x18, 0xec, 0x01,
	0x00, 0x00,
}
