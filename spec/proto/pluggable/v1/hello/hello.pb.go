// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc2
// source: spec/proto/pluggable/v1/hello/hello.proto

package hello

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	common "mosn.io/layotto/spec/proto/pluggable/v1/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config      *common.Config    `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Type        string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	HelloString string            `protobuf:"bytes,3,opt,name=hello_string,json=helloString,proto3" json:"hello_string,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelloConfig) Reset() {
	*x = HelloConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloConfig) ProtoMessage() {}

func (x *HelloConfig) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloConfig.ProtoReflect.Descriptor instead.
func (*HelloConfig) Descriptor() ([]byte, []int) {
	return file_spec_proto_pluggable_v1_hello_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloConfig) GetConfig() *common.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *HelloConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HelloConfig) GetHelloString() string {
	if x != nil {
		return x.HelloString
	}
	return ""
}

func (x *HelloConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_spec_proto_pluggable_v1_hello_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloString string `protobuf:"bytes,1,opt,name=hello_string,json=helloString,proto3" json:"hello_string,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_spec_proto_pluggable_v1_hello_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetHelloString() string {
	if x != nil {
		return x.HelloString
	}
	return ""
}

var File_spec_proto_pluggable_v1_hello_hello_proto protoreflect.FileDescriptor

var file_spec_proto_pluggable_v1_hello_hello_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0xba, 0x01, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x4a, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x65, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x72, 0x0a, 0x1d, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x1c, 0x50, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x33, 0x6d, 0x6f, 0x73, 0x6e, 0x2e, 0x69, 0x6f, 0x2f, 0x6c, 0x61, 0x79, 0x6f,
	0x74, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_proto_pluggable_v1_hello_hello_proto_rawDescOnce sync.Once
	file_spec_proto_pluggable_v1_hello_hello_proto_rawDescData = file_spec_proto_pluggable_v1_hello_hello_proto_rawDesc
)

func file_spec_proto_pluggable_v1_hello_hello_proto_rawDescGZIP() []byte {
	file_spec_proto_pluggable_v1_hello_hello_proto_rawDescOnce.Do(func() {
		file_spec_proto_pluggable_v1_hello_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_proto_pluggable_v1_hello_hello_proto_rawDescData)
	})
	return file_spec_proto_pluggable_v1_hello_hello_proto_rawDescData
}

var file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spec_proto_pluggable_v1_hello_hello_proto_goTypes = []interface{}{
	(*HelloConfig)(nil),   // 0: spec.proto.pluggable.v1.hello.HelloConfig
	(*HelloRequest)(nil),  // 1: spec.proto.pluggable.v1.hello.HelloRequest
	(*HelloResponse)(nil), // 2: spec.proto.pluggable.v1.hello.HelloResponse
	nil,                   // 3: spec.proto.pluggable.v1.hello.HelloConfig.MetadataEntry
	(*common.Config)(nil), // 4: spec.proto.pluggable.v1.common.Config
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_spec_proto_pluggable_v1_hello_hello_proto_depIdxs = []int32{
	4, // 0: spec.proto.pluggable.v1.hello.HelloConfig.config:type_name -> spec.proto.pluggable.v1.common.Config
	3, // 1: spec.proto.pluggable.v1.hello.HelloConfig.metadata:type_name -> spec.proto.pluggable.v1.hello.HelloConfig.MetadataEntry
	0, // 2: spec.proto.pluggable.v1.hello.Hello.Init:input_type -> spec.proto.pluggable.v1.hello.HelloConfig
	1, // 3: spec.proto.pluggable.v1.hello.Hello.SayHello:input_type -> spec.proto.pluggable.v1.hello.HelloRequest
	5, // 4: spec.proto.pluggable.v1.hello.Hello.Init:output_type -> google.protobuf.Empty
	2, // 5: spec.proto.pluggable.v1.hello.Hello.SayHello:output_type -> spec.proto.pluggable.v1.hello.HelloResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spec_proto_pluggable_v1_hello_hello_proto_init() }
func file_spec_proto_pluggable_v1_hello_hello_proto_init() {
	if File_spec_proto_pluggable_v1_hello_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloConfig); i {
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
		file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_spec_proto_pluggable_v1_hello_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_proto_pluggable_v1_hello_hello_proto_goTypes,
		DependencyIndexes: file_spec_proto_pluggable_v1_hello_hello_proto_depIdxs,
		MessageInfos:      file_spec_proto_pluggable_v1_hello_hello_proto_msgTypes,
	}.Build()
	File_spec_proto_pluggable_v1_hello_hello_proto = out.File
	file_spec_proto_pluggable_v1_hello_hello_proto_rawDesc = nil
	file_spec_proto_pluggable_v1_hello_hello_proto_goTypes = nil
	file_spec_proto_pluggable_v1_hello_hello_proto_depIdxs = nil
}