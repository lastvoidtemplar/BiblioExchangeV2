// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: upload-client.proto

package upload_client

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *CallbackRequest) Reset() {
	*x = CallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRequest) ProtoMessage() {}

func (x *CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRequest.ProtoReflect.Descriptor instead.
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return file_upload_client_proto_rawDescGZIP(), []int{0}
}

func (x *CallbackRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type CallbackSuccRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CallbackSuccRequest) Reset() {
	*x = CallbackSuccRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackSuccRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackSuccRequest) ProtoMessage() {}

func (x *CallbackSuccRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackSuccRequest.ProtoReflect.Descriptor instead.
func (*CallbackSuccRequest) Descriptor() ([]byte, []int) {
	return file_upload_client_proto_rawDescGZIP(), []int{1}
}

func (x *CallbackSuccRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *CallbackSuccRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_upload_client_proto protoreflect.FileDescriptor

var file_upload_client_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x13, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x75, 0x63, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32,
	0xfb, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x50, 0x0a, 0x12, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53,
	0x75, 0x63, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x4e, 0x0a, 0x14, 0x4f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x57, 0x68, 0x65,
	0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x49, 0x0a, 0x0f, 0x4f, 0x6e, 0x55, 0x72, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x19, 0x5a,
	0x17, 0x2e, 0x2e, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_client_proto_rawDescOnce sync.Once
	file_upload_client_proto_rawDescData = file_upload_client_proto_rawDesc
)

func file_upload_client_proto_rawDescGZIP() []byte {
	file_upload_client_proto_rawDescOnce.Do(func() {
		file_upload_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_client_proto_rawDescData)
	})
	return file_upload_client_proto_rawDescData
}

var file_upload_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_upload_client_proto_goTypes = []interface{}{
	(*CallbackRequest)(nil),     // 0: upload_client.CallbackRequest
	(*CallbackSuccRequest)(nil), // 1: upload_client.CallbackSuccRequest
	(*empty.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_upload_client_proto_depIdxs = []int32{
	1, // 0: upload_client.UploadClient.OnSuccessfulUpload:input_type -> upload_client.CallbackSuccRequest
	0, // 1: upload_client.UploadClient.OnErrorWhenUploading:input_type -> upload_client.CallbackRequest
	0, // 2: upload_client.UploadClient.OnUrlExpiration:input_type -> upload_client.CallbackRequest
	2, // 3: upload_client.UploadClient.OnSuccessfulUpload:output_type -> google.protobuf.Empty
	2, // 4: upload_client.UploadClient.OnErrorWhenUploading:output_type -> google.protobuf.Empty
	2, // 5: upload_client.UploadClient.OnUrlExpiration:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_client_proto_init() }
func file_upload_client_proto_init() {
	if File_upload_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackRequest); i {
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
		file_upload_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackSuccRequest); i {
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
			RawDescriptor: file_upload_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_client_proto_goTypes,
		DependencyIndexes: file_upload_client_proto_depIdxs,
		MessageInfos:      file_upload_client_proto_msgTypes,
	}.Build()
	File_upload_client_proto = out.File
	file_upload_client_proto_rawDesc = nil
	file_upload_client_proto_goTypes = nil
	file_upload_client_proto_depIdxs = nil
}
