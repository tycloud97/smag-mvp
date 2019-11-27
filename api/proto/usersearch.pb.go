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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Bio                  string   `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Followings           []*User  `protobuf:"bytes,6,rep,name=followings,proto3" json:"followings,omitempty"`
	Followers            []*User  `protobuf:"bytes,7,rep,name=followers,proto3" json:"followers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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

func (m *User) GetFollowings() []*User {
	if m != nil {
		return m.Followings
	}
	return nil
}

func (m *User) GetFollowers() []*User {
	if m != nil {
		return m.Followers
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
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{1}
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

type UserNameRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNameRequest) Reset()         { *m = UserNameRequest{} }
func (m *UserNameRequest) String() string { return proto.CompactTextString(m) }
func (*UserNameRequest) ProtoMessage()    {}
func (*UserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{2}
}
func (m *UserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNameRequest.Unmarshal(m, b)
}
func (m *UserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNameRequest.Marshal(b, m, deterministic)
}
func (dst *UserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNameRequest.Merge(dst, src)
}
func (m *UserNameRequest) XXX_Size() int {
	return xxx_messageInfo_UserNameRequest.Size(m)
}
func (m *UserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserNameRequest proto.InternalMessageInfo

func (m *UserNameRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{3}
}
func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (dst *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(dst, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type InstaPost struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId               string   `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ImgUrl               string   `protobuf:"bytes,3,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	Caption              string   `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	TaggedUsers          []*User  `protobuf:"bytes,5,rep,name=tagged_users,json=taggedUsers,proto3" json:"tagged_users,omitempty"`
	ShortCode            string   `protobuf:"bytes,6,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstaPost) Reset()         { *m = InstaPost{} }
func (m *InstaPost) String() string { return proto.CompactTextString(m) }
func (*InstaPost) ProtoMessage()    {}
func (*InstaPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{4}
}
func (m *InstaPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstaPost.Unmarshal(m, b)
}
func (m *InstaPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstaPost.Marshal(b, m, deterministic)
}
func (dst *InstaPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstaPost.Merge(dst, src)
}
func (m *InstaPost) XXX_Size() int {
	return xxx_messageInfo_InstaPost.Size(m)
}
func (m *InstaPost) XXX_DiscardUnknown() {
	xxx_messageInfo_InstaPost.DiscardUnknown(m)
}

var xxx_messageInfo_InstaPost proto.InternalMessageInfo

func (m *InstaPost) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstaPost) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *InstaPost) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *InstaPost) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *InstaPost) GetTaggedUsers() []*User {
	if m != nil {
		return m.TaggedUsers
	}
	return nil
}

func (m *InstaPost) GetShortCode() string {
	if m != nil {
		return m.ShortCode
	}
	return ""
}

type InstaPostsResponse struct {
	UserId               string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InstaPosts           []*InstaPost `protobuf:"bytes,2,rep,name=insta_posts,json=instaPosts,proto3" json:"insta_posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InstaPostsResponse) Reset()         { *m = InstaPostsResponse{} }
func (m *InstaPostsResponse) String() string { return proto.CompactTextString(m) }
func (*InstaPostsResponse) ProtoMessage()    {}
func (*InstaPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{5}
}
func (m *InstaPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstaPostsResponse.Unmarshal(m, b)
}
func (m *InstaPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstaPostsResponse.Marshal(b, m, deterministic)
}
func (dst *InstaPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstaPostsResponse.Merge(dst, src)
}
func (m *InstaPostsResponse) XXX_Size() int {
	return xxx_messageInfo_InstaPostsResponse.Size(m)
}
func (m *InstaPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstaPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstaPostsResponse proto.InternalMessageInfo

func (m *InstaPostsResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *InstaPostsResponse) GetInstaPosts() []*InstaPost {
	if m != nil {
		return m.InstaPosts
	}
	return nil
}

type FaceSearchRequest struct {
	Base64EncodedPicture string   `protobuf:"bytes,1,opt,name=base64EncodedPicture,proto3" json:"base64EncodedPicture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceSearchRequest) Reset()         { *m = FaceSearchRequest{} }
func (m *FaceSearchRequest) String() string { return proto.CompactTextString(m) }
func (*FaceSearchRequest) ProtoMessage()    {}
func (*FaceSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{6}
}
func (m *FaceSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceSearchRequest.Unmarshal(m, b)
}
func (m *FaceSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceSearchRequest.Marshal(b, m, deterministic)
}
func (dst *FaceSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceSearchRequest.Merge(dst, src)
}
func (m *FaceSearchRequest) XXX_Size() int {
	return xxx_messageInfo_FaceSearchRequest.Size(m)
}
func (m *FaceSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FaceSearchRequest proto.InternalMessageInfo

func (m *FaceSearchRequest) GetBase64EncodedPicture() string {
	if m != nil {
		return m.Base64EncodedPicture
	}
	return ""
}

type Face struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	PostId               int32    `protobuf:"varint,5,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	FullImageSrc         string   `protobuf:"bytes,6,opt,name=full_image_src,json=fullImageSrc,proto3" json:"full_image_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Face) Reset()         { *m = Face{} }
func (m *Face) String() string { return proto.CompactTextString(m) }
func (*Face) ProtoMessage()    {}
func (*Face) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{7}
}
func (m *Face) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Face.Unmarshal(m, b)
}
func (m *Face) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Face.Marshal(b, m, deterministic)
}
func (dst *Face) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Face.Merge(dst, src)
}
func (m *Face) XXX_Size() int {
	return xxx_messageInfo_Face.Size(m)
}
func (m *Face) XXX_DiscardUnknown() {
	xxx_messageInfo_Face.DiscardUnknown(m)
}

var xxx_messageInfo_Face proto.InternalMessageInfo

func (m *Face) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Face) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Face) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Face) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Face) GetPostId() int32 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *Face) GetFullImageSrc() string {
	if m != nil {
		return m.FullImageSrc
	}
	return ""
}

type FaceSearchResponse struct {
	Faces                []*Face  `protobuf:"bytes,1,rep,name=faces,proto3" json:"faces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceSearchResponse) Reset()         { *m = FaceSearchResponse{} }
func (m *FaceSearchResponse) String() string { return proto.CompactTextString(m) }
func (*FaceSearchResponse) ProtoMessage()    {}
func (*FaceSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_f9fab015ff81956f, []int{8}
}
func (m *FaceSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceSearchResponse.Unmarshal(m, b)
}
func (m *FaceSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceSearchResponse.Marshal(b, m, deterministic)
}
func (dst *FaceSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceSearchResponse.Merge(dst, src)
}
func (m *FaceSearchResponse) XXX_Size() int {
	return xxx_messageInfo_FaceSearchResponse.Size(m)
}
func (m *FaceSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FaceSearchResponse proto.InternalMessageInfo

func (m *FaceSearchResponse) GetFaces() []*Face {
	if m != nil {
		return m.Faces
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserSearchResponse)(nil), "proto.UserSearchResponse")
	proto.RegisterType((*UserNameRequest)(nil), "proto.UserNameRequest")
	proto.RegisterType((*UserIdRequest)(nil), "proto.UserIdRequest")
	proto.RegisterType((*InstaPost)(nil), "proto.InstaPost")
	proto.RegisterType((*InstaPostsResponse)(nil), "proto.InstaPostsResponse")
	proto.RegisterType((*FaceSearchRequest)(nil), "proto.FaceSearchRequest")
	proto.RegisterType((*Face)(nil), "proto.Face")
	proto.RegisterType((*FaceSearchResponse)(nil), "proto.FaceSearchResponse")
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
	GetAllUsersLikeUsername(ctx context.Context, in *UserNameRequest, opts ...grpc.CallOption) (*UserSearchResponse, error)
	GetUserWithUsername(ctx context.Context, in *UserNameRequest, opts ...grpc.CallOption) (*User, error)
	GetInstaPostsWithUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*InstaPostsResponse, error)
	SearchSimilarFaces(ctx context.Context, in *FaceSearchRequest, opts ...grpc.CallOption) (*FaceSearchResponse, error)
}

type userSearchServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserSearchServiceClient(cc *grpc.ClientConn) UserSearchServiceClient {
	return &userSearchServiceClient{cc}
}

func (c *userSearchServiceClient) GetAllUsersLikeUsername(ctx context.Context, in *UserNameRequest, opts ...grpc.CallOption) (*UserSearchResponse, error) {
	out := new(UserSearchResponse)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetAllUsersLikeUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSearchServiceClient) GetUserWithUsername(ctx context.Context, in *UserNameRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetUserWithUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSearchServiceClient) GetInstaPostsWithUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*InstaPostsResponse, error) {
	out := new(InstaPostsResponse)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetInstaPostsWithUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSearchServiceClient) SearchSimilarFaces(ctx context.Context, in *FaceSearchRequest, opts ...grpc.CallOption) (*FaceSearchResponse, error) {
	out := new(FaceSearchResponse)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/SearchSimilarFaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSearchServiceServer is the server API for UserSearchService service.
type UserSearchServiceServer interface {
	GetAllUsersLikeUsername(context.Context, *UserNameRequest) (*UserSearchResponse, error)
	GetUserWithUsername(context.Context, *UserNameRequest) (*User, error)
	GetInstaPostsWithUserId(context.Context, *UserIdRequest) (*InstaPostsResponse, error)
	SearchSimilarFaces(context.Context, *FaceSearchRequest) (*FaceSearchResponse, error)
}

func RegisterUserSearchServiceServer(s *grpc.Server, srv UserSearchServiceServer) {
	s.RegisterService(&_UserSearchService_serviceDesc, srv)
}

func _UserSearchService_GetAllUsersLikeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).GetAllUsersLikeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/GetAllUsersLikeUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).GetAllUsersLikeUsername(ctx, req.(*UserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSearchService_GetUserWithUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNameRequest)
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
		return srv.(UserSearchServiceServer).GetUserWithUsername(ctx, req.(*UserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSearchService_GetInstaPostsWithUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).GetInstaPostsWithUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/GetInstaPostsWithUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).GetInstaPostsWithUserId(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSearchService_SearchSimilarFaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaceSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).SearchSimilarFaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/SearchSimilarFaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).SearchSimilarFaces(ctx, req.(*FaceSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserSearchService",
	HandlerType: (*UserSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsersLikeUsername",
			Handler:    _UserSearchService_GetAllUsersLikeUsername_Handler,
		},
		{
			MethodName: "GetUserWithUsername",
			Handler:    _UserSearchService_GetUserWithUsername_Handler,
		},
		{
			MethodName: "GetInstaPostsWithUserId",
			Handler:    _UserSearchService_GetInstaPostsWithUserId_Handler,
		},
		{
			MethodName: "SearchSimilarFaces",
			Handler:    _UserSearchService_SearchSimilarFaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/usersearch.proto",
}

func init() {
	proto.RegisterFile("api/proto/usersearch.proto", fileDescriptor_usersearch_f9fab015ff81956f)
}

var fileDescriptor_usersearch_f9fab015ff81956f = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xd3, 0x4a,
	0x10, 0xc6, 0xe5, 0xb4, 0x76, 0x4e, 0x26, 0x3d, 0xa5, 0x5d, 0xaa, 0x76, 0x29, 0x42, 0x2a, 0x16,
	0x17, 0x41, 0x48, 0xa9, 0x28, 0x08, 0x24, 0x2e, 0x90, 0x10, 0x82, 0x2a, 0xa8, 0x42, 0x95, 0xab,
	0x8a, 0x4b, 0xb3, 0xb5, 0xa7, 0xce, 0x0a, 0x27, 0x1b, 0x76, 0xd7, 0xfd, 0xf3, 0x14, 0xbc, 0x0c,
	0xe2, 0x5d, 0x78, 0x1b, 0x34, 0xbb, 0x8e, 0xeb, 0x34, 0x45, 0x5c, 0xc5, 0xf3, 0x7d, 0xe3, 0xf1,
	0xcc, 0x6f, 0x26, 0xb0, 0x2b, 0x66, 0x72, 0x7f, 0xa6, 0x95, 0x55, 0xfb, 0x95, 0x41, 0x6d, 0x50,
	0xe8, 0x6c, 0x3c, 0x74, 0x02, 0x0b, 0xdd, 0x4f, 0xfc, 0x3b, 0x80, 0xd5, 0x53, 0x83, 0x9a, 0xad,
	0x43, 0x47, 0xe6, 0x3c, 0xd8, 0x0b, 0x06, 0xbd, 0xa4, 0x23, 0x73, 0xf6, 0x10, 0x7a, 0xf4, 0x4e,
	0x3a, 0x15, 0x13, 0xe4, 0x1d, 0x27, 0xff, 0x47, 0xc2, 0x67, 0x31, 0x41, 0x32, 0x35, 0x8a, 0xd2,
	0x9b, 0x2b, 0xde, 0x24, 0xc1, 0x99, 0x1b, 0xb0, 0x72, 0x26, 0x15, 0x5f, 0x75, 0x32, 0x3d, 0xb2,
	0x47, 0x00, 0xe2, 0x42, 0x58, 0xa1, 0xd3, 0x4a, 0x97, 0x3c, 0x74, 0x46, 0xcf, 0x2b, 0xa7, 0xba,
	0x64, 0xcf, 0x00, 0xce, 0x55, 0x59, 0xaa, 0x4b, 0x39, 0x2d, 0x0c, 0x8f, 0xf6, 0x56, 0x06, 0xfd,
	0x83, 0xbe, 0x6f, 0x73, 0x48, 0xbd, 0x25, 0x2d, 0x9b, 0x3d, 0x85, 0x9e, 0x8f, 0x50, 0x1b, 0xde,
	0x5d, 0xce, 0xbd, 0x71, 0xe3, 0xb7, 0xc0, 0x48, 0x3a, 0x71, 0x63, 0x27, 0x68, 0x66, 0x6a, 0x6a,
	0x90, 0x0d, 0xea, 0xc1, 0x4a, 0x69, 0x2c, 0x0f, 0x96, 0x0b, 0xb8, 0x29, 0x8f, 0xa4, 0xb1, 0xf1,
	0x10, 0xee, 0x9d, 0xd6, 0x13, 0x27, 0xf8, 0xbd, 0x42, 0x63, 0x17, 0xa9, 0x04, 0x8b, 0x54, 0xe2,
	0x01, 0xfc, 0x4f, 0xf9, 0xa3, 0x7c, 0x9e, 0xbd, 0x03, 0x5d, 0x97, 0xdd, 0x80, 0x8d, 0x2a, 0xe7,
	0xc7, 0xbf, 0x02, 0xe8, 0x8d, 0xa6, 0xc6, 0x8a, 0x63, 0x65, 0xec, 0x12, 0xfa, 0x1d, 0xe8, 0xce,
	0x94, 0xb1, 0xf4, 0x9a, 0x07, 0x1f, 0x51, 0x38, 0x72, 0x86, 0x9c, 0x14, 0x0e, 0xa2, 0x87, 0x1e,
	0xc9, 0x49, 0x41, 0x04, 0x39, 0x74, 0x33, 0x31, 0xb3, 0x52, 0x4d, 0x6b, 0xec, 0xf3, 0x90, 0x0d,
	0x61, 0xcd, 0x8a, 0xa2, 0xc0, 0x3c, 0x75, 0x17, 0xc0, 0xc3, 0xe5, 0x81, 0xfb, 0x3e, 0x81, 0x9e,
	0x0d, 0xad, 0xca, 0x8c, 0x95, 0xb6, 0x69, 0xa6, 0x72, 0xe4, 0x91, 0x5f, 0x95, 0x53, 0xde, 0xab,
	0x1c, 0xe3, 0xaf, 0xc0, 0x9a, 0xbe, 0x4d, 0x83, 0xf4, 0x6f, 0x73, 0xb2, 0xe7, 0xd0, 0x97, 0x94,
	0x9e, 0xd2, 0x00, 0x86, 0x77, 0xdc, 0xc7, 0x37, 0xea, 0x8f, 0x37, 0x85, 0x12, 0x90, 0x4d, 0xcd,
	0xf8, 0x10, 0x36, 0x3f, 0x8a, 0x0c, 0xe7, 0x4b, 0xf3, 0x20, 0x0f, 0x60, 0xeb, 0x4c, 0x18, 0x7c,
	0xf5, 0xf2, 0xc3, 0x94, 0xfa, 0xca, 0x8f, 0x65, 0x66, 0x2b, 0x3d, 0xdf, 0xc0, 0x9d, 0x5e, 0xfc,
	0x23, 0x80, 0x55, 0xaa, 0xc4, 0xd6, 0x20, 0xb8, 0x72, 0x99, 0x61, 0x12, 0x5c, 0x51, 0x74, 0xed,
	0xb0, 0x86, 0x49, 0x70, 0xcd, 0xb6, 0x20, 0xbc, 0x94, 0xb9, 0x1d, 0x3b, 0x9e, 0x61, 0xe2, 0x03,
	0xb6, 0x0d, 0xd1, 0x18, 0x65, 0x31, 0xb6, 0x8e, 0x66, 0x98, 0xd4, 0x51, 0x7b, 0x31, 0xa1, 0x37,
	0xea, 0xc5, 0x3c, 0x81, 0xf5, 0xf3, 0xaa, 0x2c, 0x53, 0x39, 0x11, 0x05, 0xa6, 0x46, 0x67, 0x35,
	0xb9, 0x35, 0x52, 0x47, 0x24, 0x9e, 0xe8, 0x2c, 0x7e, 0x0d, 0xac, 0x3d, 0x5a, 0x0d, 0xef, 0x31,
	0x84, 0xe7, 0x22, 0x43, 0x73, 0xeb, 0x16, 0x29, 0x33, 0xf1, 0xce, 0xc1, 0xcf, 0x0e, 0x6c, 0xde,
	0x5c, 0xf2, 0x09, 0xea, 0x0b, 0x99, 0x21, 0x3b, 0x82, 0x9d, 0x43, 0xb4, 0xef, 0xca, 0xd2, 0x6d,
	0xee, 0x48, 0x7e, 0x43, 0x7a, 0xa0, 0xcb, 0x64, 0xdb, 0xad, 0xfd, 0xb6, 0xce, 0x77, 0xf7, 0x41,
	0x4b, 0xbf, 0xd5, 0xc6, 0x1b, 0xb8, 0x7f, 0x88, 0x96, 0x8c, 0x2f, 0xd2, 0x8e, 0xff, 0x59, 0xa9,
	0x7d, 0x41, 0xec, 0x93, 0xeb, 0xe4, 0xe6, 0x30, 0xe6, 0x15, 0x46, 0x39, 0xdb, 0x6a, 0xe5, 0x35,
	0x7f, 0x8c, 0xa6, 0x8f, 0x3b, 0x6e, 0x69, 0x04, 0xac, 0x1e, 0x53, 0x4e, 0x64, 0x29, 0x34, 0x71,
	0x30, 0x8c, 0xb7, 0xa8, 0x2c, 0x9c, 0x46, 0x53, 0x6a, 0x99, 0xec, 0x59, 0xe4, 0x9c, 0x17, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x29, 0x3b, 0xfa, 0xec, 0x07, 0x05, 0x00, 0x00,
}
