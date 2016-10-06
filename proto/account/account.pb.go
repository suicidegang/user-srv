// Code generated by protoc-gen-go.
// source: account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	User
	CreateRequest
	CreateResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created  string `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	Updated  string `protobuf:"bytes,5,opt,name=updated" json:"updated,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateRequest struct {
	User     *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Account{hdlr}, opts...))
}

type Account struct {
	AccountHandler
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x6d, 0xa7, 0x9d, 0xd1, 0x2b, 0x53, 0x25, 0xb8, 0x88, 0x5d, 0x95, 0xac, 0x0a, 0x42,
	0x16, 0x15, 0xdc, 0x8b, 0xe0, 0x03, 0x04, 0x7c, 0x80, 0xd8, 0xdc, 0x45, 0xc1, 0x36, 0x31, 0x3f,
	0xb8, 0xf1, 0xe1, 0x25, 0x89, 0x19, 0xe8, 0xf2, 0xbb, 0x5f, 0x92, 0x73, 0x72, 0xe1, 0x2c, 0xe7,
	0x59, 0x87, 0xcd, 0x73, 0x63, 0xb5, 0xd7, 0xec, 0x17, 0x9a, 0x0f, 0x87, 0x96, 0x74, 0x50, 0x2f,
	0x8a, 0x56, 0x43, 0x35, 0x1e, 0x44, 0xbd, 0x28, 0xd2, 0xc3, 0x75, 0x70, 0x68, 0x37, 0xb9, 0x22,
	0xad, 0x87, 0x6a, 0xbc, 0x11, 0x17, 0x26, 0x0f, 0xd0, 0xe2, 0x2a, 0x97, 0x2f, 0x7a, 0x48, 0x22,
	0x03, 0xa1, 0x70, 0x9a, 0x2d, 0x4a, 0x8f, 0x8a, 0x36, 0x69, 0x5e, 0x30, 0x9a, 0x60, 0x54, 0x32,
	0x6d, 0x36, 0xff, 0xc8, 0xde, 0xe1, 0xfc, 0x96, 0x0e, 0x09, 0xfc, 0x0e, 0xe8, 0x3c, 0x79, 0x84,
	0x26, 0xc6, 0xa4, 0x22, 0xb7, 0x53, 0xcb, 0x63, 0x37, 0x91, 0x46, 0xb1, 0x91, 0x91, 0xce, 0xfd,
	0x68, 0xab, 0x4a, 0xa3, 0xc2, 0xec, 0x1e, 0xba, 0xf2, 0x8e, 0x33, 0x7a, 0x73, 0x38, 0xbd, 0xc0,
	0xe9, 0x35, 0x7f, 0x94, 0x3c, 0xc1, 0x31, 0x4b, 0xd2, 0xf1, 0x5d, 0x5a, 0x7f, 0xc7, 0xf7, 0xb7,
	0xd8, 0xd5, 0xe7, 0x31, 0xad, 0xe5, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x01, 0xd4, 0x3d, 0xfe,
	0x27, 0x01, 0x00, 0x00,
}