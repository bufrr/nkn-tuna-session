// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pb/session.proto

package pb

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

type SessionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SessionMetadata) Reset() {
	*x = SessionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionMetadata) ProtoMessage() {}

func (x *SessionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionMetadata.ProtoReflect.Descriptor instead.
func (*SessionMetadata) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{0}
}

func (x *SessionMetadata) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_pb_session_proto protoreflect.FileDescriptor

var file_pb_session_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_session_proto_rawDescOnce sync.Once
	file_pb_session_proto_rawDescData = file_pb_session_proto_rawDesc
)

func file_pb_session_proto_rawDescGZIP() []byte {
	file_pb_session_proto_rawDescOnce.Do(func() {
		file_pb_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_session_proto_rawDescData)
	})
	return file_pb_session_proto_rawDescData
}

var file_pb_session_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_session_proto_goTypes = []interface{}{
	(*SessionMetadata)(nil), // 0: pb.SessionMetadata
}
var file_pb_session_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_session_proto_init() }
func file_pb_session_proto_init() {
	if File_pb_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionMetadata); i {
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
			RawDescriptor: file_pb_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_session_proto_goTypes,
		DependencyIndexes: file_pb_session_proto_depIdxs,
		MessageInfos:      file_pb_session_proto_msgTypes,
	}.Build()
	File_pb_session_proto = out.File
	file_pb_session_proto_rawDesc = nil
	file_pb_session_proto_goTypes = nil
	file_pb_session_proto_depIdxs = nil
}
