// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: mock/v1/ssh.proto

package mockv1

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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{1}
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HostConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host         *Host  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port         int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Workdir      string `protobuf:"bytes,3,opt,name=workdir,proto3" json:"workdir,omitempty"`
	Passphrase   string `protobuf:"bytes,4,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	PrivateKey   []byte `protobuf:"bytes,5,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	AuthUsername string `protobuf:"bytes,6,opt,name=auth_username,json=authUsername,proto3" json:"auth_username,omitempty"`
	AuthPassword string `protobuf:"bytes,7,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	AuthPrivkey  []byte `protobuf:"bytes,8,opt,name=auth_privkey,json=authPrivkey,proto3" json:"auth_privkey,omitempty"`
}

func (x *HostConfig) Reset() {
	*x = HostConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostConfig) ProtoMessage() {}

func (x *HostConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostConfig.ProtoReflect.Descriptor instead.
func (*HostConfig) Descriptor() ([]byte, []int) {
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{2}
}

func (x *HostConfig) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *HostConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *HostConfig) GetWorkdir() string {
	if x != nil {
		return x.Workdir
	}
	return ""
}

func (x *HostConfig) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

func (x *HostConfig) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *HostConfig) GetAuthUsername() string {
	if x != nil {
		return x.AuthUsername
	}
	return ""
}

func (x *HostConfig) GetAuthPassword() string {
	if x != nil {
		return x.AuthPassword
	}
	return ""
}

func (x *HostConfig) GetAuthPrivkey() []byte {
	if x != nil {
		return x.AuthPrivkey
	}
	return nil
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Eq:
	//
	//	*Query_Name
	//	*Query_Port
	Eq isQuery_Eq `protobuf_oneof:"eq"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{3}
}

func (m *Query) GetEq() isQuery_Eq {
	if m != nil {
		return m.Eq
	}
	return nil
}

func (x *Query) GetName() string {
	if x, ok := x.GetEq().(*Query_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Query) GetPort() int32 {
	if x, ok := x.GetEq().(*Query_Port); ok {
		return x.Port
	}
	return 0
}

type isQuery_Eq interface {
	isQuery_Eq()
}

type Query_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type Query_Port struct {
	Port int32 `protobuf:"varint,2,opt,name=port,proto3,oneof"`
}

func (*Query_Name) isQuery_Eq() {}

func (*Query_Port) isQuery_Eq() {}

type BundleChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BundleChunk) Reset() {
	*x = BundleChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleChunk) ProtoMessage() {}

func (x *BundleChunk) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleChunk.ProtoReflect.Descriptor instead.
func (*BundleChunk) Descriptor() ([]byte, []int) {
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{4}
}

func (x *BundleChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_v1_ssh_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_mock_v1_ssh_proto_msgTypes[5]
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
	return file_mock_v1_ssh_proto_rawDescGZIP(), []int{5}
}

var File_mock_v1_ssh_proto protoreflect.FileDescriptor

var file_mock_v1_ssh_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x8b, 0x02, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x21, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x64, 0x69, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x6b, 0x65, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x69, 0x76, 0x6b, 0x65,
	0x79, 0x22, 0x39, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x04, 0x0a, 0x02, 0x65, 0x71, 0x22, 0x21, 0x0a, 0x0b,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xfb, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x63,
	0x6b, 0x53, 0x53, 0x48, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x4e,
	0x65, 0x77, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x06,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x08, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01, 0x12,
	0x25, 0x0a, 0x04, 0x46, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x95, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x61, 0x75, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x74, 0x61, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x70, 0x6f, 0x72, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x63,
	0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4d, 0x6f, 0x63, 0x6b,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4d, 0x6f, 0x63, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x4d, 0x6f, 0x63, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4d, 0x6f, 0x63, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mock_v1_ssh_proto_rawDescOnce sync.Once
	file_mock_v1_ssh_proto_rawDescData = file_mock_v1_ssh_proto_rawDesc
)

func file_mock_v1_ssh_proto_rawDescGZIP() []byte {
	file_mock_v1_ssh_proto_rawDescOnce.Do(func() {
		file_mock_v1_ssh_proto_rawDescData = protoimpl.X.CompressGZIP(file_mock_v1_ssh_proto_rawDescData)
	})
	return file_mock_v1_ssh_proto_rawDescData
}

var file_mock_v1_ssh_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mock_v1_ssh_proto_goTypes = []any{
	(*Command)(nil),     // 0: mock.v1.Command
	(*Host)(nil),        // 1: mock.v1.Host
	(*HostConfig)(nil),  // 2: mock.v1.HostConfig
	(*Query)(nil),       // 3: mock.v1.Query
	(*BundleChunk)(nil), // 4: mock.v1.BundleChunk
	(*Empty)(nil),       // 5: mock.v1.Empty
}
var file_mock_v1_ssh_proto_depIdxs = []int32{
	1, // 0: mock.v1.HostConfig.host:type_name -> mock.v1.Host
	2, // 1: mock.v1.MockSSHService.New:input_type -> mock.v1.HostConfig
	3, // 2: mock.v1.MockSSHService.Lookup:input_type -> mock.v1.Query
	1, // 3: mock.v1.MockSSHService.Commands:input_type -> mock.v1.Host
	1, // 4: mock.v1.MockSSHService.Filesystem:input_type -> mock.v1.Host
	1, // 5: mock.v1.MockSSHService.Free:input_type -> mock.v1.Host
	2, // 6: mock.v1.MockSSHService.New:output_type -> mock.v1.HostConfig
	2, // 7: mock.v1.MockSSHService.Lookup:output_type -> mock.v1.HostConfig
	0, // 8: mock.v1.MockSSHService.Commands:output_type -> mock.v1.Command
	4, // 9: mock.v1.MockSSHService.Filesystem:output_type -> mock.v1.BundleChunk
	5, // 10: mock.v1.MockSSHService.Free:output_type -> mock.v1.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mock_v1_ssh_proto_init() }
func file_mock_v1_ssh_proto_init() {
	if File_mock_v1_ssh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mock_v1_ssh_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Command); i {
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
		file_mock_v1_ssh_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Host); i {
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
		file_mock_v1_ssh_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HostConfig); i {
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
		file_mock_v1_ssh_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Query); i {
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
		file_mock_v1_ssh_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BundleChunk); i {
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
		file_mock_v1_ssh_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
	file_mock_v1_ssh_proto_msgTypes[3].OneofWrappers = []any{
		(*Query_Name)(nil),
		(*Query_Port)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mock_v1_ssh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mock_v1_ssh_proto_goTypes,
		DependencyIndexes: file_mock_v1_ssh_proto_depIdxs,
		MessageInfos:      file_mock_v1_ssh_proto_msgTypes,
	}.Build()
	File_mock_v1_ssh_proto = out.File
	file_mock_v1_ssh_proto_rawDesc = nil
	file_mock_v1_ssh_proto_goTypes = nil
	file_mock_v1_ssh_proto_depIdxs = nil
}
