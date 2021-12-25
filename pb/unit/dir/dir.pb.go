// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pb/unit/dir/dir.proto

package dir

import (
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

type Dir struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *Spec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Dir) Reset() {
	*x = Dir{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_dir_dir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dir) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dir) ProtoMessage() {}

func (x *Dir) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_dir_dir_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dir.ProtoReflect.Descriptor instead.
func (*Dir) Descriptor() ([]byte, []int) {
	return file_pb_unit_dir_dir_proto_rawDescGZIP(), []int{0}
}

func (x *Dir) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size        string `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	WrittenDate string `protobuf:"bytes,2,opt,name=written_date,json=writtenDate,proto3" json:"written_date,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_dir_dir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_dir_dir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_pb_unit_dir_dir_proto_rawDescGZIP(), []int{1}
}

func (x *Spec) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Spec) GetWrittenDate() string {
	if x != nil {
		return x.WrittenDate
	}
	return ""
}

var File_pb_unit_dir_dir_proto protoreflect.FileDescriptor

var file_pb_unit_dir_dir_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x64, 0x69, 0x72, 0x2f, 0x64, 0x69,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x62, 0x2e, 0x75, 0x6e, 0x69, 0x74,
	0x2e, 0x64, 0x69, 0x72, 0x22, 0x2c, 0x0a, 0x03, 0x44, 0x69, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x75,
	0x6e, 0x69, 0x74, 0x2e, 0x64, 0x69, 0x72, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0x3d, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x67, 0x6c, 0x69, 0x64, 0x65, 0x31, 0x30, 0x30, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x64, 0x69, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_unit_dir_dir_proto_rawDescOnce sync.Once
	file_pb_unit_dir_dir_proto_rawDescData = file_pb_unit_dir_dir_proto_rawDesc
)

func file_pb_unit_dir_dir_proto_rawDescGZIP() []byte {
	file_pb_unit_dir_dir_proto_rawDescOnce.Do(func() {
		file_pb_unit_dir_dir_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_unit_dir_dir_proto_rawDescData)
	})
	return file_pb_unit_dir_dir_proto_rawDescData
}

var file_pb_unit_dir_dir_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_unit_dir_dir_proto_goTypes = []interface{}{
	(*Dir)(nil),  // 0: pb.unit.dir.Dir
	(*Spec)(nil), // 1: pb.unit.dir.Spec
}
var file_pb_unit_dir_dir_proto_depIdxs = []int32{
	1, // 0: pb.unit.dir.Dir.spec:type_name -> pb.unit.dir.Spec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_unit_dir_dir_proto_init() }
func file_pb_unit_dir_dir_proto_init() {
	if File_pb_unit_dir_dir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_unit_dir_dir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dir); i {
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
		file_pb_unit_dir_dir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_pb_unit_dir_dir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_unit_dir_dir_proto_goTypes,
		DependencyIndexes: file_pb_unit_dir_dir_proto_depIdxs,
		MessageInfos:      file_pb_unit_dir_dir_proto_msgTypes,
	}.Build()
	File_pb_unit_dir_dir_proto = out.File
	file_pb_unit_dir_dir_proto_rawDesc = nil
	file_pb_unit_dir_dir_proto_goTypes = nil
	file_pb_unit_dir_dir_proto_depIdxs = nil
}
