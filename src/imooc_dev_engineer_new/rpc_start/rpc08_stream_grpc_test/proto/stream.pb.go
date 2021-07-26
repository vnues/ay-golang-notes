// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: stream.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamRequestData) Reset() {
	*x = StreamRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequestData) ProtoMessage() {}

func (x *StreamRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequestData.ProtoReflect.Descriptor instead.
func (*StreamRequestData) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequestData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamResponseData) Reset() {
	*x = StreamResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponseData) ProtoMessage() {}

func (x *StreamResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponseData.ProtoReflect.Descriptor instead.
func (*StreamResponseData) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponseData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27,
	0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xb3, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12, 0x38, 0x0a,
	0x09, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_proto_goTypes = []interface{}{
	(*StreamRequestData)(nil),  // 0: StreamRequestData
	(*StreamResponseData)(nil), // 1: StreamResponseData
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: Greeter.GetStream:input_type -> StreamRequestData
	0, // 1: Greeter.PutStream:input_type -> StreamRequestData
	0, // 2: Greeter.AllStream:input_type -> StreamRequestData
	1, // 3: Greeter.GetStream:output_type -> StreamResponseData
	1, // 4: Greeter.PutStream:output_type -> StreamResponseData
	1, // 5: Greeter.AllStream:output_type -> StreamResponseData
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequestData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	//服务器端流模式
	GetStream(ctx context.Context, in *StreamRequestData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error)
	//客户端流模式
	PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error)
	//双向流模式
	AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetStream(ctx context.Context, in *StreamRequestData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/Greeter/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GetStreamClient interface {
	Recv() (*StreamResponseData, error)
	grpc.ClientStream
}

type greeterGetStreamClient struct {
	grpc.ClientStream
}

func (x *greeterGetStreamClient) Recv() (*StreamResponseData, error) {
	m := new(StreamResponseData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/Greeter/PutStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterPutStreamClient{stream}
	return x, nil
}

type Greeter_PutStreamClient interface {
	Send(*StreamRequestData) error
	CloseAndRecv() (*StreamResponseData, error)
	grpc.ClientStream
}

type greeterPutStreamClient struct {
	grpc.ClientStream
}

func (x *greeterPutStreamClient) Send(m *StreamRequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterPutStreamClient) CloseAndRecv() (*StreamResponseData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponseData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[2], "/Greeter/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterAllStreamClient{stream}
	return x, nil
}

type Greeter_AllStreamClient interface {
	Send(*StreamRequestData) error
	Recv() (*StreamResponseData, error)
	grpc.ClientStream
}

type greeterAllStreamClient struct {
	grpc.ClientStream
}

func (x *greeterAllStreamClient) Send(m *StreamRequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterAllStreamClient) Recv() (*StreamResponseData, error) {
	m := new(StreamResponseData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	//服务器端流模式
	GetStream(*StreamRequestData, Greeter_GetStreamServer) error
	//客户端流模式
	PutStream(Greeter_PutStreamServer) error
	//双向流模式
	AllStream(Greeter_AllStreamServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) GetStream(*StreamRequestData, Greeter_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (*UnimplementedGreeterServer) PutStream(Greeter_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}
func (*UnimplementedGreeterServer) AllStream(Greeter_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequestData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetStream(m, &greeterGetStreamServer{stream})
}

type Greeter_GetStreamServer interface {
	Send(*StreamResponseData) error
	grpc.ServerStream
}

type greeterGetStreamServer struct {
	grpc.ServerStream
}

func (x *greeterGetStreamServer) Send(m *StreamResponseData) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).PutStream(&greeterPutStreamServer{stream})
}

type Greeter_PutStreamServer interface {
	SendAndClose(*StreamResponseData) error
	Recv() (*StreamRequestData, error)
	grpc.ServerStream
}

type greeterPutStreamServer struct {
	grpc.ServerStream
}

func (x *greeterPutStreamServer) SendAndClose(m *StreamResponseData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterPutStreamServer) Recv() (*StreamRequestData, error) {
	m := new(StreamRequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).AllStream(&greeterAllStreamServer{stream})
}

type Greeter_AllStreamServer interface {
	Send(*StreamResponseData) error
	Recv() (*StreamRequestData, error)
	grpc.ServerStream
}

type greeterAllStreamServer struct {
	grpc.ServerStream
}

func (x *greeterAllStreamServer) Send(m *StreamResponseData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterAllStreamServer) Recv() (*StreamRequestData, error) {
	m := new(StreamRequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Greeter_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutStream",
			Handler:       _Greeter_PutStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _Greeter_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
