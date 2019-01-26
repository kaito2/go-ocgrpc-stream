// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc-stream.proto

package grpcstream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SampleRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleRequest) Reset()         { *m = SampleRequest{} }
func (m *SampleRequest) String() string { return proto.CompactTextString(m) }
func (*SampleRequest) ProtoMessage()    {}
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a243c71aba31818, []int{0}
}

func (m *SampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleRequest.Unmarshal(m, b)
}
func (m *SampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleRequest.Marshal(b, m, deterministic)
}
func (m *SampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleRequest.Merge(m, src)
}
func (m *SampleRequest) XXX_Size() int {
	return xxx_messageInfo_SampleRequest.Size(m)
}
func (m *SampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleRequest proto.InternalMessageInfo

func (m *SampleRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SampleResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleResponse) Reset()         { *m = SampleResponse{} }
func (m *SampleResponse) String() string { return proto.CompactTextString(m) }
func (*SampleResponse) ProtoMessage()    {}
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a243c71aba31818, []int{1}
}

func (m *SampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleResponse.Unmarshal(m, b)
}
func (m *SampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleResponse.Marshal(b, m, deterministic)
}
func (m *SampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleResponse.Merge(m, src)
}
func (m *SampleResponse) XXX_Size() int {
	return xxx_messageInfo_SampleResponse.Size(m)
}
func (m *SampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SampleResponse proto.InternalMessageInfo

func (m *SampleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SampleRequest)(nil), "grpcstream.SampleRequest")
	proto.RegisterType((*SampleResponse)(nil), "grpcstream.SampleResponse")
}

func init() { proto.RegisterFile("grpc-stream.proto", fileDescriptor_2a243c71aba31818) }

var fileDescriptor_2a243c71aba31818 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0xd6, 0x2d, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x09, 0x41, 0x44, 0x94, 0x34, 0xb9, 0x78, 0x83, 0x13, 0x73, 0x0b, 0x72, 0x52, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x2d, 0x2e, 0x3e, 0x98, 0xd2, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0xdc, 0x6a, 0x8d, 0xe2, 0xb8, 0x78, 0x20, 0x6a, 0x83, 0xc1, 0xd6, 0x08,
	0xf9, 0x71, 0x09, 0xba, 0xa7, 0x96, 0xa0, 0x69, 0x97, 0xd4, 0x43, 0x38, 0x44, 0x0f, 0xc5, 0x15,
	0x52, 0x52, 0xd8, 0xa4, 0x20, 0xda, 0x94, 0x18, 0x0c, 0x18, 0x93, 0xd8, 0xc0, 0x3e, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x24, 0x47, 0x6a, 0xfb, 0xde, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleStreamClient is the client API for SampleStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleStreamClient interface {
	GetSampleResponse(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (SampleStream_GetSampleResponseClient, error)
}

type sampleStreamClient struct {
	cc *grpc.ClientConn
}

func NewSampleStreamClient(cc *grpc.ClientConn) SampleStreamClient {
	return &sampleStreamClient{cc}
}

func (c *sampleStreamClient) GetSampleResponse(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (SampleStream_GetSampleResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SampleStream_serviceDesc.Streams[0], "/grpcstream.SampleStream/GetSampleResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleStreamGetSampleResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SampleStream_GetSampleResponseClient interface {
	Recv() (*SampleResponse, error)
	grpc.ClientStream
}

type sampleStreamGetSampleResponseClient struct {
	grpc.ClientStream
}

func (x *sampleStreamGetSampleResponseClient) Recv() (*SampleResponse, error) {
	m := new(SampleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleStreamServer is the server API for SampleStream service.
type SampleStreamServer interface {
	GetSampleResponse(*SampleRequest, SampleStream_GetSampleResponseServer) error
}

func RegisterSampleStreamServer(s *grpc.Server, srv SampleStreamServer) {
	s.RegisterService(&_SampleStream_serviceDesc, srv)
}

func _SampleStream_GetSampleResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SampleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SampleStreamServer).GetSampleResponse(m, &sampleStreamGetSampleResponseServer{stream})
}

type SampleStream_GetSampleResponseServer interface {
	Send(*SampleResponse) error
	grpc.ServerStream
}

type sampleStreamGetSampleResponseServer struct {
	grpc.ServerStream
}

func (x *sampleStreamGetSampleResponseServer) Send(m *SampleResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SampleStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcstream.SampleStream",
	HandlerType: (*SampleStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSampleResponse",
			Handler:       _SampleStream_GetSampleResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc-stream.proto",
}