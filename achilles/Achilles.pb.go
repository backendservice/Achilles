// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Achilles.proto

package achilles

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

// The request message containing the user's name.
type AchillesRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AchillesRequest) Reset()         { *m = AchillesRequest{} }
func (m *AchillesRequest) String() string { return proto.CompactTextString(m) }
func (*AchillesRequest) ProtoMessage()    {}
func (*AchillesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Achilles_4a8e8ff1b0665b83, []int{0}
}
func (m *AchillesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AchillesRequest.Unmarshal(m, b)
}
func (m *AchillesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AchillesRequest.Marshal(b, m, deterministic)
}
func (dst *AchillesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AchillesRequest.Merge(dst, src)
}
func (m *AchillesRequest) XXX_Size() int {
	return xxx_messageInfo_AchillesRequest.Size(m)
}
func (m *AchillesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AchillesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AchillesRequest proto.InternalMessageInfo

func (m *AchillesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type AchillesReply struct {
	Score                string   `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AchillesReply) Reset()         { *m = AchillesReply{} }
func (m *AchillesReply) String() string { return proto.CompactTextString(m) }
func (*AchillesReply) ProtoMessage()    {}
func (*AchillesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_Achilles_4a8e8ff1b0665b83, []int{1}
}
func (m *AchillesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AchillesReply.Unmarshal(m, b)
}
func (m *AchillesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AchillesReply.Marshal(b, m, deterministic)
}
func (dst *AchillesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AchillesReply.Merge(dst, src)
}
func (m *AchillesReply) XXX_Size() int {
	return xxx_messageInfo_AchillesReply.Size(m)
}
func (m *AchillesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AchillesReply.DiscardUnknown(m)
}

var xxx_messageInfo_AchillesReply proto.InternalMessageInfo

func (m *AchillesReply) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func init() {
	proto.RegisterType((*AchillesRequest)(nil), "achilles.AchillesRequest")
	proto.RegisterType((*AchillesReply)(nil), "achilles.AchillesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ComputingMinScoreClient is the client API for ComputingMinScore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComputingMinScoreClient interface {
	// Sends a greeting
	Compute(ctx context.Context, in *AchillesRequest, opts ...grpc.CallOption) (*AchillesReply, error)
}

type computingMinScoreClient struct {
	cc *grpc.ClientConn
}

func NewComputingMinScoreClient(cc *grpc.ClientConn) ComputingMinScoreClient {
	return &computingMinScoreClient{cc}
}

func (c *computingMinScoreClient) Compute(ctx context.Context, in *AchillesRequest, opts ...grpc.CallOption) (*AchillesReply, error) {
	out := new(AchillesReply)
	err := c.cc.Invoke(ctx, "/achilles.ComputingMinScore/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputingMinScoreServer is the server API for ComputingMinScore service.
type ComputingMinScoreServer interface {
	// Sends a greeting
	Compute(context.Context, *AchillesRequest) (*AchillesReply, error)
}

func RegisterComputingMinScoreServer(s *grpc.Server, srv ComputingMinScoreServer) {
	s.RegisterService(&_ComputingMinScore_serviceDesc, srv)
}

func _ComputingMinScore_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AchillesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputingMinScoreServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achilles.ComputingMinScore/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputingMinScoreServer).Compute(ctx, req.(*AchillesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComputingMinScore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "achilles.ComputingMinScore",
	HandlerType: (*ComputingMinScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _ComputingMinScore_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Achilles.proto",
}

func init() { proto.RegisterFile("Achilles.proto", fileDescriptor_Achilles_4a8e8ff1b0665b83) }

var fileDescriptor_Achilles_4a8e8ff1b0665b83 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0x4c, 0xce, 0xc8,
	0xcc, 0xc9, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x84, 0xf2, 0x95,
	0x54, 0xb9, 0xf8, 0x61, 0x72, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x2a, 0x17,
	0x2f, 0x42, 0x59, 0x41, 0x4e, 0xa5, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x72, 0x7e, 0x11, 0x4c, 0x15,
	0x84, 0x63, 0x14, 0xc2, 0x25, 0xe8, 0x9c, 0x9f, 0x5b, 0x50, 0x5a, 0x92, 0x99, 0x97, 0xee, 0x9b,
	0x99, 0x17, 0x0c, 0x12, 0x14, 0xb2, 0xe7, 0x62, 0x87, 0x08, 0xa6, 0x0a, 0x49, 0xea, 0xc1, 0x2c,
	0xd6, 0x43, 0xb3, 0x55, 0x4a, 0x1c, 0x9b, 0x54, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x93, 0x01, 0x97,
	0x52, 0x72, 0x7e, 0xae, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x52, 0x62, 0x72, 0x76,
	0x6a, 0x5e, 0x4a, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0x2a, 0x5c, 0xb1, 0x13, 0x07, 0x8c, 0x15,
	0xc0, 0x98, 0xc4, 0x06, 0xf6, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x42, 0x29, 0x21,
	0xf8, 0x00, 0x00, 0x00,
}
