// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: qotd.proto

package qotd

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

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qotd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_qotd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_qotd_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Quote  string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qotd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_qotd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_qotd_proto_rawDescGZIP(), []int{1}
}

func (x *GetResp) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetResp) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

var File_qotd_proto protoreflect.FileDescriptor

var file_qotd_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x71, 0x6f, 0x74, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x71, 0x6f,
	0x74, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x32, 0x30, 0x0a,
	0x04, 0x51, 0x4f, 0x54, 0x44, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x51, 0x4f, 0x54, 0x44,
	0x12, 0x0c, 0x2e, 0x71, 0x6f, 0x74, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d,
	0x2e, 0x71, 0x6f, 0x74, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f,
	0x46, 0x6f, 0x72, 0x44, 0x65, 0x76, 0x4f, 0x70, 0x73, 0x2f, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x52,
	0x50, 0x43, 0x2f, 0x71, 0x6f, 0x74, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qotd_proto_rawDescOnce sync.Once
	file_qotd_proto_rawDescData = file_qotd_proto_rawDesc
)

func file_qotd_proto_rawDescGZIP() []byte {
	file_qotd_proto_rawDescOnce.Do(func() {
		file_qotd_proto_rawDescData = protoimpl.X.CompressGZIP(file_qotd_proto_rawDescData)
	})
	return file_qotd_proto_rawDescData
}

var file_qotd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qotd_proto_goTypes = []interface{}{
	(*GetReq)(nil),  // 0: qotd.GetReq
	(*GetResp)(nil), // 1: qotd.GetResp
}
var file_qotd_proto_depIdxs = []int32{
	0, // 0: qotd.QOTD.GetQOTD:input_type -> qotd.GetReq
	1, // 1: qotd.QOTD.GetQOTD:output_type -> qotd.GetResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qotd_proto_init() }
func file_qotd_proto_init() {
	if File_qotd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qotd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_qotd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
			RawDescriptor: file_qotd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qotd_proto_goTypes,
		DependencyIndexes: file_qotd_proto_depIdxs,
		MessageInfos:      file_qotd_proto_msgTypes,
	}.Build()
	File_qotd_proto = out.File
	file_qotd_proto_rawDesc = nil
	file_qotd_proto_goTypes = nil
	file_qotd_proto_depIdxs = nil
}
