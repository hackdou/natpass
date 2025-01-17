// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: connect.proto

package network

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

type ConnectRequestType int32

const (
	ConnectRequest_tcp   ConnectRequestType = 0
	ConnectRequest_udp   ConnectRequestType = 1
	ConnectRequest_shell ConnectRequestType = 2
	ConnectRequest_vnc   ConnectRequestType = 3
)

// Enum value maps for ConnectRequestType.
var (
	ConnectRequestType_name = map[int32]string{
		0: "tcp",
		1: "udp",
		2: "shell",
		3: "vnc",
	}
	ConnectRequestType_value = map[string]int32{
		"tcp":   0,
		"udp":   1,
		"shell": 2,
		"vnc":   3,
	}
)

func (x ConnectRequestType) Enum() *ConnectRequestType {
	p := new(ConnectRequestType)
	*p = x
	return p
}

func (x ConnectRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_connect_proto_enumTypes[0].Descriptor()
}

func (ConnectRequestType) Type() protoreflect.EnumType {
	return &file_connect_proto_enumTypes[0]
}

func (x ConnectRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectRequestType.Descriptor instead.
func (ConnectRequestType) EnumDescriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{3, 0}
}

type ConnectAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectAddr) Reset() {
	*x = ConnectAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAddr) ProtoMessage() {}

func (x *ConnectAddr) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAddr.ProtoReflect.Descriptor instead.
func (*ConnectAddr) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectAddr) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ConnectAddr) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ConnectShell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exec string   `protobuf:"bytes,1,opt,name=exec,proto3" json:"exec,omitempty"`
	Env  []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *ConnectShell) Reset() {
	*x = ConnectShell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectShell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectShell) ProtoMessage() {}

func (x *ConnectShell) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectShell.ProtoReflect.Descriptor instead.
func (*ConnectShell) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectShell) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

func (x *ConnectShell) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

type ConnectVnc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fps     uint32 `protobuf:"varint,1,opt,name=fps,proto3" json:"fps,omitempty"`
	Quality uint32 `protobuf:"varint,2,opt,name=quality,proto3" json:"quality,omitempty"` // image quality, percent
	Cursor  bool   `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`   // show cursor
}

func (x *ConnectVnc) Reset() {
	*x = ConnectVnc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectVnc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectVnc) ProtoMessage() {}

func (x *ConnectVnc) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectVnc.ProtoReflect.Descriptor instead.
func (*ConnectVnc) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectVnc) GetFps() uint32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *ConnectVnc) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *ConnectVnc) GetCursor() bool {
	if x != nil {
		return x.Cursor
	}
	return false
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XType ConnectRequestType `protobuf:"varint,2,opt,name=_type,json=Type,proto3,enum=network.ConnectRequestType" json:"_type,omitempty"`
	// Types that are assignable to Payload:
	//	*ConnectRequest_Caddr
	//	*ConnectRequest_Cshell
	//	*ConnectRequest_Cvnc
	Payload isConnectRequest_Payload `protobuf_oneof:"payload"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectRequest) GetXType() ConnectRequestType {
	if x != nil {
		return x.XType
	}
	return ConnectRequest_tcp
}

func (m *ConnectRequest) GetPayload() isConnectRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ConnectRequest) GetCaddr() *ConnectAddr {
	if x, ok := x.GetPayload().(*ConnectRequest_Caddr); ok {
		return x.Caddr
	}
	return nil
}

func (x *ConnectRequest) GetCshell() *ConnectShell {
	if x, ok := x.GetPayload().(*ConnectRequest_Cshell); ok {
		return x.Cshell
	}
	return nil
}

func (x *ConnectRequest) GetCvnc() *ConnectVnc {
	if x, ok := x.GetPayload().(*ConnectRequest_Cvnc); ok {
		return x.Cvnc
	}
	return nil
}

type isConnectRequest_Payload interface {
	isConnectRequest_Payload()
}

type ConnectRequest_Caddr struct {
	Caddr *ConnectAddr `protobuf:"bytes,10,opt,name=caddr,proto3,oneof"`
}

type ConnectRequest_Cshell struct {
	Cshell *ConnectShell `protobuf:"bytes,11,opt,name=cshell,proto3,oneof"`
}

type ConnectRequest_Cvnc struct {
	Cvnc *ConnectVnc `protobuf:"bytes,12,opt,name=cvnc,proto3,oneof"`
}

func (*ConnectRequest_Caddr) isConnectRequest_Payload() {}

func (*ConnectRequest_Cshell) isConnectRequest_Payload() {}

func (*ConnectRequest_Cvnc) isConnectRequest_Payload() {}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ConnectResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_connect_proto protoreflect.FileDescriptor

var file_connect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x36, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x35, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x68, 0x65, 0x6c,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x51, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x5f, 0x76, 0x6e, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x9f, 0x02, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x61, 0x64, 0x64, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x48, 0x00, 0x52,
	0x06, 0x63, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x76, 0x6e, 0x63, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x6e, 0x63, 0x48, 0x00, 0x52, 0x04, 0x63,
	0x76, 0x6e, 0x63, 0x22, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x74,
	0x63, 0x70, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x64, 0x70, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x76, 0x6e, 0x63, 0x10,
	0x03, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x34, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connect_proto_rawDescOnce sync.Once
	file_connect_proto_rawDescData = file_connect_proto_rawDesc
)

func file_connect_proto_rawDescGZIP() []byte {
	file_connect_proto_rawDescOnce.Do(func() {
		file_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_proto_rawDescData)
	})
	return file_connect_proto_rawDescData
}

var file_connect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_connect_proto_goTypes = []interface{}{
	(ConnectRequestType)(0), // 0: network.connect_request.type
	(*ConnectAddr)(nil),     // 1: network.connect_addr
	(*ConnectShell)(nil),    // 2: network.connect_shell
	(*ConnectVnc)(nil),      // 3: network.connect_vnc
	(*ConnectRequest)(nil),  // 4: network.connect_request
	(*ConnectResponse)(nil), // 5: network.connect_response
}
var file_connect_proto_depIdxs = []int32{
	0, // 0: network.connect_request._type:type_name -> network.connect_request.type
	1, // 1: network.connect_request.caddr:type_name -> network.connect_addr
	2, // 2: network.connect_request.cshell:type_name -> network.connect_shell
	3, // 3: network.connect_request.cvnc:type_name -> network.connect_vnc
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_connect_proto_init() }
func file_connect_proto_init() {
	if File_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectAddr); i {
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
		file_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectShell); i {
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
		file_connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectVnc); i {
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
		file_connect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_connect_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
	file_connect_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ConnectRequest_Caddr)(nil),
		(*ConnectRequest_Cshell)(nil),
		(*ConnectRequest_Cvnc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connect_proto_goTypes,
		DependencyIndexes: file_connect_proto_depIdxs,
		EnumInfos:         file_connect_proto_enumTypes,
		MessageInfos:      file_connect_proto_msgTypes,
	}.Build()
	File_connect_proto = out.File
	file_connect_proto_rawDesc = nil
	file_connect_proto_goTypes = nil
	file_connect_proto_depIdxs = nil
}
