//@proto_id=1000

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.3
// source: Proto/Common.proto

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

type ErrorAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdId int32  `protobuf:"varint,1,opt,name=CmdId,proto3" json:"CmdId,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *ErrorAck) Reset() {
	*x = ErrorAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorAck) ProtoMessage() {}

func (x *ErrorAck) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorAck.ProtoReflect.Descriptor instead.
func (*ErrorAck) Descriptor() ([]byte, []int) {
	return file_Proto_Common_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorAck) GetCmdId() int32 {
	if x != nil {
		return x.CmdId
	}
	return 0
}

func (x *ErrorAck) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdCode        int32 `protobuf:"varint,1,opt,name=CmdCode,proto3" json:"CmdCode,omitempty"`
	ProtocolSwitch int32 `protobuf:"varint,2,opt,name=ProtocolSwitch,proto3" json:"ProtocolSwitch,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_Proto_Common_proto_rawDescGZIP(), []int{1}
}

func (x *PingReq) GetCmdCode() int32 {
	if x != nil {
		return x.CmdCode
	}
	return 0
}

func (x *PingReq) GetProtocolSwitch() int32 {
	if x != nil {
		return x.ProtocolSwitch
	}
	return 0
}

type PingAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *PingAck) Reset() {
	*x = PingAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingAck) ProtoMessage() {}

func (x *PingAck) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingAck.ProtoReflect.Descriptor instead.
func (*PingAck) Descriptor() ([]byte, []int) {
	return file_Proto_Common_proto_rawDescGZIP(), []int{2}
}

func (x *PingAck) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type TestNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int32 `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *TestNtf) Reset() {
	*x = TestNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNtf) ProtoMessage() {}

func (x *TestNtf) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNtf.ProtoReflect.Descriptor instead.
func (*TestNtf) Descriptor() ([]byte, []int) {
	return file_Proto_Common_proto_rawDescGZIP(), []int{3}
}

func (x *TestNtf) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_Proto_Common_proto protoreflect.FileDescriptor

var file_Proto_Common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x32, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x4b, 0x0a, 0x07,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x1d, 0x0a, 0x07, 0x50, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74,
	0x4e, 0x74, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0xaa, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_Common_proto_rawDescOnce sync.Once
	file_Proto_Common_proto_rawDescData = file_Proto_Common_proto_rawDesc
)

func file_Proto_Common_proto_rawDescGZIP() []byte {
	file_Proto_Common_proto_rawDescOnce.Do(func() {
		file_Proto_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_Common_proto_rawDescData)
	})
	return file_Proto_Common_proto_rawDescData
}

var file_Proto_Common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Proto_Common_proto_goTypes = []interface{}{
	(*ErrorAck)(nil), // 0: pb.ErrorAck
	(*PingReq)(nil),  // 1: pb.PingReq
	(*PingAck)(nil),  // 2: pb.PingAck
	(*TestNtf)(nil),  // 3: pb.TestNtf
}
var file_Proto_Common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Proto_Common_proto_init() }
func file_Proto_Common_proto_init() {
	if File_Proto_Common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_Common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorAck); i {
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
		file_Proto_Common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_Proto_Common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingAck); i {
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
		file_Proto_Common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNtf); i {
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
			RawDescriptor: file_Proto_Common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Proto_Common_proto_goTypes,
		DependencyIndexes: file_Proto_Common_proto_depIdxs,
		MessageInfos:      file_Proto_Common_proto_msgTypes,
	}.Build()
	File_Proto_Common_proto = out.File
	file_Proto_Common_proto_rawDesc = nil
	file_Proto_Common_proto_goTypes = nil
	file_Proto_Common_proto_depIdxs = nil
}
