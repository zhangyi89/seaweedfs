// Code generated by protoc-gen-go.
// source: iam.proto
// DO NOT EDIT!

/*
Package iam_pb is a generated protocol buffer package.

It is generated from these files:
	iam.proto

It has these top-level messages:
	S3ApiConfiguration
	Identity
	Credential
*/
package iam_pb

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

type S3ApiConfiguration struct {
	Identities []*Identity `protobuf:"bytes,1,rep,name=identities" json:"identities,omitempty"`
}

func (m *S3ApiConfiguration) Reset()                    { *m = S3ApiConfiguration{} }
func (m *S3ApiConfiguration) String() string            { return proto.CompactTextString(m) }
func (*S3ApiConfiguration) ProtoMessage()               {}
func (*S3ApiConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *S3ApiConfiguration) GetIdentities() []*Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

type Identity struct {
	Name        string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Credentials []*Credential `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
	Actions     []string      `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *Identity) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Credential struct {
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey" json:"secret_key,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Credential) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Credential) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func init() {
	proto.RegisterType((*S3ApiConfiguration)(nil), "iam_pb.S3ApiConfiguration")
	proto.RegisterType((*Identity)(nil), "iam_pb.Identity")
	proto.RegisterType((*Credential)(nil), "iam_pb.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedIdentityAccessManagement service

type SeaweedIdentityAccessManagementClient interface {
}

type seaweedIdentityAccessManagementClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedIdentityAccessManagementClient(cc *grpc.ClientConn) SeaweedIdentityAccessManagementClient {
	return &seaweedIdentityAccessManagementClient{cc}
}

// Server API for SeaweedIdentityAccessManagement service

type SeaweedIdentityAccessManagementServer interface {
}

func RegisterSeaweedIdentityAccessManagementServer(s *grpc.Server, srv SeaweedIdentityAccessManagementServer) {
	s.RegisterService(&_SeaweedIdentityAccessManagement_serviceDesc, srv)
}

var _SeaweedIdentityAccessManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iam_pb.SeaweedIdentityAccessManagement",
	HandlerType: (*SeaweedIdentityAccessManagementServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "iam.proto",
}

func init() { proto.RegisterFile("iam.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x69, 0x23, 0xb5, 0x99, 0x5e, 0xca, 0x9c, 0xf6, 0xa0, 0x18, 0x73, 0xca, 0x29, 0x48,
	0xeb, 0x1f, 0xa8, 0x05, 0xa1, 0x16, 0x41, 0xd2, 0x1f, 0x50, 0xa6, 0xdb, 0x69, 0x19, 0xec, 0x6e,
	0x42, 0x76, 0x45, 0xf2, 0xef, 0x25, 0xbb, 0x46, 0x7b, 0xdb, 0x7d, 0xdf, 0x7b, 0xb3, 0x3b, 0x0f,
	0x52, 0x21, 0x53, 0x36, 0x6d, 0xed, 0x6b, 0x9c, 0x08, 0x99, 0x7d, 0x73, 0xc8, 0x5f, 0x01, 0x77,
	0xcb, 0x55, 0x23, 0xeb, 0xda, 0x9e, 0xe4, 0xfc, 0xd5, 0x92, 0x97, 0xda, 0xe2, 0x13, 0x80, 0x1c,
	0xd9, 0x7a, 0xf1, 0xc2, 0x4e, 0x8d, 0xb2, 0xa4, 0x98, 0x2d, 0xe6, 0x65, 0x8c, 0x94, 0x9b, 0x48,
	0xba, 0xea, 0xca, 0x93, 0x5b, 0x98, 0x0e, 0x3a, 0x22, 0xdc, 0x58, 0x32, 0xac, 0x46, 0xd9, 0xa8,
	0x48, 0xab, 0x70, 0xc6, 0x67, 0x98, 0xe9, 0x96, 0x83, 0x83, 0x2e, 0x4e, 0x8d, 0xc3, 0x48, 0x1c,
	0x46, 0xae, 0xff, 0x50, 0x75, 0x6d, 0x43, 0x05, 0xb7, 0xa4, 0xfb, 0x1f, 0x39, 0x95, 0x64, 0x49,
	0x91, 0x56, 0xc3, 0x35, 0x7f, 0x03, 0xf8, 0x0f, 0xe1, 0x3d, 0x00, 0x69, 0xcd, 0xce, 0xed, 0x3f,
	0xb9, 0xfb, 0x7d, 0x37, 0x8d, 0xca, 0x96, 0xbb, 0x1e, 0x3b, 0xd6, 0x2d, 0xfb, 0x80, 0xc7, 0x11,
	0x47, 0x65, 0xcb, 0xdd, 0xe2, 0x11, 0x1e, 0x76, 0x4c, 0xdf, 0xcc, 0xc7, 0x61, 0x85, 0x55, 0x88,
	0xbe, 0x93, 0xa5, 0x33, 0x1b, 0xb6, 0xfe, 0xe5, 0x0e, 0xe6, 0x2e, 0x5a, 0x4e, 0xae, 0xd4, 0x17,
	0xe9, 0xb5, 0xe9, 0x86, 0xcc, 0x47, 0x5f, 0xe6, 0x61, 0x12, 0x3a, 0x5d, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x83, 0x4f, 0x61, 0x03, 0x60, 0x01, 0x00, 0x00,
}
