// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/backbone/main.proto

package backbone

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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string  `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Reps  []int64 `protobuf:"varint,2,rep,packed,name=reps,proto3" json:"reps,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backbone_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backbone_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_backbone_main_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Test) GetReps() []int64 {
	if x != nil {
		return x.Reps
	}
	return nil
}

var File_proto_backbone_main_proto protoreflect.FileDescriptor

var file_proto_backbone_main_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e, 0x65,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x63,
	0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x22, 0x30, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x04, 0x72, 0x65, 0x70, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_backbone_main_proto_rawDescOnce sync.Once
	file_proto_backbone_main_proto_rawDescData = file_proto_backbone_main_proto_rawDesc
)

func file_proto_backbone_main_proto_rawDescGZIP() []byte {
	file_proto_backbone_main_proto_rawDescOnce.Do(func() {
		file_proto_backbone_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_backbone_main_proto_rawDescData)
	})
	return file_proto_backbone_main_proto_rawDescData
}

var file_proto_backbone_main_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_backbone_main_proto_goTypes = []any{
	(*Test)(nil), // 0: backbone.Test
}
var file_proto_backbone_main_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_backbone_main_proto_init() }
func file_proto_backbone_main_proto_init() {
	if File_proto_backbone_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_backbone_main_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_proto_backbone_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_backbone_main_proto_goTypes,
		DependencyIndexes: file_proto_backbone_main_proto_depIdxs,
		MessageInfos:      file_proto_backbone_main_proto_msgTypes,
	}.Build()
	File_proto_backbone_main_proto = out.File
	file_proto_backbone_main_proto_rawDesc = nil
	file_proto_backbone_main_proto_goTypes = nil
	file_proto_backbone_main_proto_depIdxs = nil
}
