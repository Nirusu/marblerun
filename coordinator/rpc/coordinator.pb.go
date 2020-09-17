// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: coordinator.proto

package rpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ActivationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: sending the quote via metadata/context would be cleaner.
	Quote      []byte `protobuf:"bytes,1,opt,name=Quote,proto3" json:"Quote,omitempty"`
	CSR        []byte `protobuf:"bytes,2,opt,name=CSR,proto3" json:"CSR,omitempty"`
	MarbleType string `protobuf:"bytes,3,opt,name=MarbleType,proto3" json:"MarbleType,omitempty"`
	UUID       string `protobuf:"bytes,4,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *ActivationReq) Reset() {
	*x = ActivationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationReq) ProtoMessage() {}

func (x *ActivationReq) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationReq.ProtoReflect.Descriptor instead.
func (*ActivationReq) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *ActivationReq) GetQuote() []byte {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *ActivationReq) GetCSR() []byte {
	if x != nil {
		return x.CSR
	}
	return nil
}

func (x *ActivationReq) GetMarbleType() string {
	if x != nil {
		return x.MarbleType
	}
	return ""
}

func (x *ActivationReq) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type ActivationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters *Parameters `protobuf:"bytes,1,opt,name=Parameters,proto3" json:"Parameters,omitempty"`
}

func (x *ActivationResp) Reset() {
	*x = ActivationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationResp) ProtoMessage() {}

func (x *ActivationResp) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationResp.ProtoReflect.Descriptor instead.
func (*ActivationResp) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *ActivationResp) GetParameters() *Parameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files map[string][]byte `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env   map[string]string `protobuf:"bytes,2,rep,name=Env,proto3" json:"Env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Argv  []string          `protobuf:"bytes,3,rep,name=Argv,proto3" json:"Argv,omitempty"`
}

func (x *Parameters) Reset() {
	*x = Parameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters) ProtoMessage() {}

func (x *Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters.ProtoReflect.Descriptor instead.
func (*Parameters) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *Parameters) GetFiles() map[string][]byte {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Parameters) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Parameters) GetArgv() []string {
	if x != nil {
		return x.Argv
	}
	return nil
}

var File_coordinator_proto protoreflect.FileDescriptor

var file_coordinator_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x6b, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x53, 0x52, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x43, 0x53,
	0x52, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x61, 0x72, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x41, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x03, 0x45, 0x6e, 0x76,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x45, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x76, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x67, 0x76, 0x1a, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x3d, 0x0a, 0x06, 0x4d,
	0x61, 0x72, 0x62, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x65, 0x73,
	0x73, 0x73, 0x79, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coordinator_proto_rawDescOnce sync.Once
	file_coordinator_proto_rawDescData = file_coordinator_proto_rawDesc
)

func file_coordinator_proto_rawDescGZIP() []byte {
	file_coordinator_proto_rawDescOnce.Do(func() {
		file_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_coordinator_proto_rawDescData)
	})
	return file_coordinator_proto_rawDescData
}

var file_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_coordinator_proto_goTypes = []interface{}{
	(*ActivationReq)(nil),  // 0: rpc.ActivationReq
	(*ActivationResp)(nil), // 1: rpc.ActivationResp
	(*Parameters)(nil),     // 2: rpc.Parameters
	nil,                    // 3: rpc.Parameters.FilesEntry
	nil,                    // 4: rpc.Parameters.EnvEntry
}
var file_coordinator_proto_depIdxs = []int32{
	2, // 0: rpc.ActivationResp.Parameters:type_name -> rpc.Parameters
	3, // 1: rpc.Parameters.Files:type_name -> rpc.Parameters.FilesEntry
	4, // 2: rpc.Parameters.Env:type_name -> rpc.Parameters.EnvEntry
	0, // 3: rpc.Marble.Activate:input_type -> rpc.ActivationReq
	1, // 4: rpc.Marble.Activate:output_type -> rpc.ActivationResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coordinator_proto_init() }
func file_coordinator_proto_init() {
	if File_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationReq); i {
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
		file_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationResp); i {
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
		file_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameters); i {
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
			RawDescriptor: file_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coordinator_proto_goTypes,
		DependencyIndexes: file_coordinator_proto_depIdxs,
		MessageInfos:      file_coordinator_proto_msgTypes,
	}.Build()
	File_coordinator_proto = out.File
	file_coordinator_proto_rawDesc = nil
	file_coordinator_proto_goTypes = nil
	file_coordinator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MarbleClient is the client API for Marble service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarbleClient interface {
	// Activate activates a marble in the mesh.
	Activate(ctx context.Context, in *ActivationReq, opts ...grpc.CallOption) (*ActivationResp, error)
}

type marbleClient struct {
	cc grpc.ClientConnInterface
}

func NewMarbleClient(cc grpc.ClientConnInterface) MarbleClient {
	return &marbleClient{cc}
}

func (c *marbleClient) Activate(ctx context.Context, in *ActivationReq, opts ...grpc.CallOption) (*ActivationResp, error) {
	out := new(ActivationResp)
	err := c.cc.Invoke(ctx, "/rpc.Marble/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarbleServer is the server API for Marble service.
type MarbleServer interface {
	// Activate activates a marble in the mesh.
	Activate(context.Context, *ActivationReq) (*ActivationResp, error)
}

// UnimplementedMarbleServer can be embedded to have forward compatible implementations.
type UnimplementedMarbleServer struct {
}

func (*UnimplementedMarbleServer) Activate(context.Context, *ActivationReq) (*ActivationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}

func RegisterMarbleServer(s *grpc.Server, srv MarbleServer) {
	s.RegisterService(&_Marble_serviceDesc, srv)
}

func _Marble_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarbleServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Marble/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarbleServer).Activate(ctx, req.(*ActivationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Marble_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Marble",
	HandlerType: (*MarbleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _Marble_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}
