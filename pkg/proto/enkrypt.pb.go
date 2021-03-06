// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/proto/enkrypt.proto

package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EncryptedFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *EncryptedFiles) Reset() {
	*x = EncryptedFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_enkrypt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedFiles) ProtoMessage() {}

func (x *EncryptedFiles) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_enkrypt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedFiles.ProtoReflect.Descriptor instead.
func (*EncryptedFiles) Descriptor() ([]byte, []int) {
	return file_pkg_proto_enkrypt_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptedFiles) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_enkrypt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_enkrypt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_pkg_proto_enkrypt_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_enkrypt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_enkrypt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_proto_enkrypt_proto_rawDescGZIP(), []int{2}
}

var File_pkg_proto_enkrypt_proto protoreflect.FileDescriptor

var file_pkg_proto_enkrypt_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x6b, 0x72,
	0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x26, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x69, 0x0a,
	0x07, 0x45, 0x6e, 0x6b, 0x72, 0x79, 0x70, 0x74, 0x12, 0x39, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_enkrypt_proto_rawDescOnce sync.Once
	file_pkg_proto_enkrypt_proto_rawDescData = file_pkg_proto_enkrypt_proto_rawDesc
)

func file_pkg_proto_enkrypt_proto_rawDescGZIP() []byte {
	file_pkg_proto_enkrypt_proto_rawDescOnce.Do(func() {
		file_pkg_proto_enkrypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_enkrypt_proto_rawDescData)
	})
	return file_pkg_proto_enkrypt_proto_rawDescData
}

var file_pkg_proto_enkrypt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_enkrypt_proto_goTypes = []interface{}{
	(*EncryptedFiles)(nil), // 0: proto.EncryptedFiles
	(*File)(nil),           // 1: proto.File
	(*Empty)(nil),          // 2: proto.Empty
}
var file_pkg_proto_enkrypt_proto_depIdxs = []int32{
	2, // 0: proto.Enkrypt.ListEncryptedFiles:input_type -> proto.Empty
	1, // 1: proto.Enkrypt.GetFile:input_type -> proto.File
	0, // 2: proto.Enkrypt.ListEncryptedFiles:output_type -> proto.EncryptedFiles
	1, // 3: proto.Enkrypt.GetFile:output_type -> proto.File
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_enkrypt_proto_init() }
func file_pkg_proto_enkrypt_proto_init() {
	if File_pkg_proto_enkrypt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_enkrypt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedFiles); i {
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
		file_pkg_proto_enkrypt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_pkg_proto_enkrypt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_pkg_proto_enkrypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_enkrypt_proto_goTypes,
		DependencyIndexes: file_pkg_proto_enkrypt_proto_depIdxs,
		MessageInfos:      file_pkg_proto_enkrypt_proto_msgTypes,
	}.Build()
	File_pkg_proto_enkrypt_proto = out.File
	file_pkg_proto_enkrypt_proto_rawDesc = nil
	file_pkg_proto_enkrypt_proto_goTypes = nil
	file_pkg_proto_enkrypt_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnkryptClient is the client API for Enkrypt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnkryptClient interface {
	ListEncryptedFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EncryptedFiles, error)
	GetFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
}

type enkryptClient struct {
	cc grpc.ClientConnInterface
}

func NewEnkryptClient(cc grpc.ClientConnInterface) EnkryptClient {
	return &enkryptClient{cc}
}

func (c *enkryptClient) ListEncryptedFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EncryptedFiles, error) {
	out := new(EncryptedFiles)
	err := c.cc.Invoke(ctx, "/proto.Enkrypt/ListEncryptedFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enkryptClient) GetFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/proto.Enkrypt/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnkryptServer is the server API for Enkrypt service.
type EnkryptServer interface {
	ListEncryptedFiles(context.Context, *Empty) (*EncryptedFiles, error)
	GetFile(context.Context, *File) (*File, error)
}

// UnimplementedEnkryptServer can be embedded to have forward compatible implementations.
type UnimplementedEnkryptServer struct {
}

func (*UnimplementedEnkryptServer) ListEncryptedFiles(context.Context, *Empty) (*EncryptedFiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEncryptedFiles not implemented")
}
func (*UnimplementedEnkryptServer) GetFile(context.Context, *File) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

func RegisterEnkryptServer(s *grpc.Server, srv EnkryptServer) {
	s.RegisterService(&_Enkrypt_serviceDesc, srv)
}

func _Enkrypt_ListEncryptedFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnkryptServer).ListEncryptedFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Enkrypt/ListEncryptedFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnkryptServer).ListEncryptedFiles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enkrypt_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnkryptServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Enkrypt/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnkryptServer).GetFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

var _Enkrypt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Enkrypt",
	HandlerType: (*EnkryptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEncryptedFiles",
			Handler:    _Enkrypt_ListEncryptedFiles_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _Enkrypt_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/enkrypt.proto",
}
