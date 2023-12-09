// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: model/omemo.proto

package protobuf

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

type OMEMOMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N          *uint32 `protobuf:"varint,1,req,name=n" json:"n,omitempty"`
	Pn         *uint32 `protobuf:"varint,2,req,name=pn" json:"pn,omitempty"`
	DhPub      []byte  `protobuf:"bytes,3,req,name=dh_pub,json=dhPub" json:"dh_pub,omitempty"`
	Ciphertext []byte  `protobuf:"bytes,4,opt,name=ciphertext" json:"ciphertext,omitempty"`
}

func (x *OMEMOMessage) Reset() {
	*x = OMEMOMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_omemo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMEMOMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMEMOMessage) ProtoMessage() {}

func (x *OMEMOMessage) ProtoReflect() protoreflect.Message {
	mi := &file_model_omemo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMEMOMessage.ProtoReflect.Descriptor instead.
func (*OMEMOMessage) Descriptor() ([]byte, []int) {
	return file_model_omemo_proto_rawDescGZIP(), []int{0}
}

func (x *OMEMOMessage) GetN() uint32 {
	if x != nil && x.N != nil {
		return *x.N
	}
	return 0
}

func (x *OMEMOMessage) GetPn() uint32 {
	if x != nil && x.Pn != nil {
		return *x.Pn
	}
	return 0
}

func (x *OMEMOMessage) GetDhPub() []byte {
	if x != nil {
		return x.DhPub
	}
	return nil
}

func (x *OMEMOMessage) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type OMEMOAuthenticatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac     []byte `protobuf:"bytes,1,req,name=mac" json:"mac,omitempty"`
	Message []byte `protobuf:"bytes,2,req,name=message" json:"message,omitempty"` // Byte-encoding of an OMEMOMessage
}

func (x *OMEMOAuthenticatedMessage) Reset() {
	*x = OMEMOAuthenticatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_omemo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMEMOAuthenticatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMEMOAuthenticatedMessage) ProtoMessage() {}

func (x *OMEMOAuthenticatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_model_omemo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMEMOAuthenticatedMessage.ProtoReflect.Descriptor instead.
func (*OMEMOAuthenticatedMessage) Descriptor() ([]byte, []int) {
	return file_model_omemo_proto_rawDescGZIP(), []int{1}
}

func (x *OMEMOAuthenticatedMessage) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

func (x *OMEMOAuthenticatedMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type OMEMOKeyExchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PkId    *uint32                    `protobuf:"varint,1,req,name=pk_id,json=pkId" json:"pk_id,omitempty"`
	SpkId   *uint32                    `protobuf:"varint,2,req,name=spk_id,json=spkId" json:"spk_id,omitempty"`
	Ik      []byte                     `protobuf:"bytes,3,req,name=ik" json:"ik,omitempty"`
	Ek      []byte                     `protobuf:"bytes,4,req,name=ek" json:"ek,omitempty"`
	Message *OMEMOAuthenticatedMessage `protobuf:"bytes,5,req,name=message" json:"message,omitempty"`
}

func (x *OMEMOKeyExchange) Reset() {
	*x = OMEMOKeyExchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_omemo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMEMOKeyExchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMEMOKeyExchange) ProtoMessage() {}

func (x *OMEMOKeyExchange) ProtoReflect() protoreflect.Message {
	mi := &file_model_omemo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMEMOKeyExchange.ProtoReflect.Descriptor instead.
func (*OMEMOKeyExchange) Descriptor() ([]byte, []int) {
	return file_model_omemo_proto_rawDescGZIP(), []int{2}
}

func (x *OMEMOKeyExchange) GetPkId() uint32 {
	if x != nil && x.PkId != nil {
		return *x.PkId
	}
	return 0
}

func (x *OMEMOKeyExchange) GetSpkId() uint32 {
	if x != nil && x.SpkId != nil {
		return *x.SpkId
	}
	return 0
}

func (x *OMEMOKeyExchange) GetIk() []byte {
	if x != nil {
		return x.Ik
	}
	return nil
}

func (x *OMEMOKeyExchange) GetEk() []byte {
	if x != nil {
		return x.Ek
	}
	return nil
}

func (x *OMEMOKeyExchange) GetMessage() *OMEMOAuthenticatedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_model_omemo_proto protoreflect.FileDescriptor

var file_model_omemo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0c, 0x4f, 0x4d, 0x45, 0x4d, 0x4f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x01,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x70,
	0x6e, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x05, 0x64, 0x68, 0x50, 0x75, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x47, 0x0a, 0x19, 0x4f, 0x4d, 0x45, 0x4d,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0c, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x94, 0x01, 0x0a, 0x10, 0x4f, 0x4d, 0x45, 0x4d, 0x4f, 0x4b, 0x65, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6b, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x70, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x6b,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6b, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x02,
	0x65, 0x6b, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4f, 0x4d, 0x45, 0x4d, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x6d, 0x65, 0x6c, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x6f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_model_omemo_proto_rawDescOnce sync.Once
	file_model_omemo_proto_rawDescData = file_model_omemo_proto_rawDesc
)

func file_model_omemo_proto_rawDescGZIP() []byte {
	file_model_omemo_proto_rawDescOnce.Do(func() {
		file_model_omemo_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_omemo_proto_rawDescData)
	})
	return file_model_omemo_proto_rawDescData
}

var file_model_omemo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_omemo_proto_goTypes = []interface{}{
	(*OMEMOMessage)(nil),              // 0: OMEMOMessage
	(*OMEMOAuthenticatedMessage)(nil), // 1: OMEMOAuthenticatedMessage
	(*OMEMOKeyExchange)(nil),          // 2: OMEMOKeyExchange
}
var file_model_omemo_proto_depIdxs = []int32{
	1, // 0: OMEMOKeyExchange.message:type_name -> OMEMOAuthenticatedMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_omemo_proto_init() }
func file_model_omemo_proto_init() {
	if File_model_omemo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_omemo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMEMOMessage); i {
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
		file_model_omemo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMEMOAuthenticatedMessage); i {
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
		file_model_omemo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMEMOKeyExchange); i {
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
			RawDescriptor: file_model_omemo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_omemo_proto_goTypes,
		DependencyIndexes: file_model_omemo_proto_depIdxs,
		MessageInfos:      file_model_omemo_proto_msgTypes,
	}.Build()
	File_model_omemo_proto = out.File
	file_model_omemo_proto_rawDesc = nil
	file_model_omemo_proto_goTypes = nil
	file_model_omemo_proto_depIdxs = nil
}
