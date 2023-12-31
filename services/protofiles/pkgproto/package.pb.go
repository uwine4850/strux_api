// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: package.proto

package pkgproto

import (
	baseproto "github.com/uwine4850/strux_api/services/protofiles/baseproto"
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

type MutateShowVersionBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *baseproto.BaseResponse `protobuf:"bytes,1,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
	Versions     []string                `protobuf:"bytes,2,rep,name=Versions,proto3" json:"Versions,omitempty"`
}

func (x *MutateShowVersionBaseResponse) Reset() {
	*x = MutateShowVersionBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateShowVersionBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateShowVersionBaseResponse) ProtoMessage() {}

func (x *MutateShowVersionBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateShowVersionBaseResponse.ProtoReflect.Descriptor instead.
func (*MutateShowVersionBaseResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{0}
}

func (x *MutateShowVersionBaseResponse) GetBaseResponse() *baseproto.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *MutateShowVersionBaseResponse) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

type RequestShowVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	PkgName  string `protobuf:"bytes,2,opt,name=PkgName,proto3" json:"PkgName,omitempty"`
}

func (x *RequestShowVersions) Reset() {
	*x = RequestShowVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestShowVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestShowVersions) ProtoMessage() {}

func (x *RequestShowVersions) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestShowVersions.ProtoReflect.Descriptor instead.
func (*RequestShowVersions) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{1}
}

func (x *RequestShowVersions) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestShowVersions) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

type MutateDownloadBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse       *baseproto.BaseResponse `protobuf:"bytes,1,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
	UplFiles           *UploadDirInfo          `protobuf:"bytes,2,opt,name=UplFiles,proto3" json:"UplFiles,omitempty"`
	UplDirInfo         *UploadDirInfo          `protobuf:"bytes,3,opt,name=UplDirInfo,proto3" json:"UplDirInfo,omitempty"`
	UploadDirsInfoJson []byte                  `protobuf:"bytes,4,opt,name=UploadDirsInfoJson,proto3" json:"UploadDirsInfoJson,omitempty"`
}

func (x *MutateDownloadBaseResponse) Reset() {
	*x = MutateDownloadBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateDownloadBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateDownloadBaseResponse) ProtoMessage() {}

func (x *MutateDownloadBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateDownloadBaseResponse.ProtoReflect.Descriptor instead.
func (*MutateDownloadBaseResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{2}
}

func (x *MutateDownloadBaseResponse) GetBaseResponse() *baseproto.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *MutateDownloadBaseResponse) GetUplFiles() *UploadDirInfo {
	if x != nil {
		return x.UplFiles
	}
	return nil
}

func (x *MutateDownloadBaseResponse) GetUplDirInfo() *UploadDirInfo {
	if x != nil {
		return x.UplDirInfo
	}
	return nil
}

func (x *MutateDownloadBaseResponse) GetUploadDirsInfoJson() []byte {
	if x != nil {
		return x.UploadDirsInfoJson
	}
	return nil
}

type RequestDownloadPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	PkgName  string `protobuf:"bytes,2,opt,name=PkgName,proto3" json:"PkgName,omitempty"`
	Version  string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *RequestDownloadPackage) Reset() {
	*x = RequestDownloadPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDownloadPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDownloadPackage) ProtoMessage() {}

func (x *RequestDownloadPackage) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDownloadPackage.ProtoReflect.Descriptor instead.
func (*RequestDownloadPackage) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{3}
}

func (x *RequestDownloadPackage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestDownloadPackage) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

func (x *RequestDownloadPackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RequestPackageExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	PkgName  string `protobuf:"bytes,2,opt,name=PkgName,proto3" json:"PkgName,omitempty"`
	Version  string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *RequestPackageExists) Reset() {
	*x = RequestPackageExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPackageExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPackageExists) ProtoMessage() {}

func (x *RequestPackageExists) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPackageExists.ProtoReflect.Descriptor instead.
func (*RequestPackageExists) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{4}
}

func (x *RequestPackageExists) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestPackageExists) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

func (x *RequestPackageExists) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type UploadDirInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string           `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	FileNames []string         `protobuf:"bytes,2,rep,name=FileNames,proto3" json:"FileNames,omitempty"`
	InnerDir  []*UploadDirInfo `protobuf:"bytes,3,rep,name=InnerDir,proto3" json:"InnerDir,omitempty"`
}

func (x *UploadDirInfo) Reset() {
	*x = UploadDirInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDirInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDirInfo) ProtoMessage() {}

func (x *UploadDirInfo) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDirInfo.ProtoReflect.Descriptor instead.
func (*UploadDirInfo) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{5}
}

func (x *UploadDirInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadDirInfo) GetFileNames() []string {
	if x != nil {
		return x.FileNames
	}
	return nil
}

func (x *UploadDirInfo) GetInnerDir() []*UploadDirInfo {
	if x != nil {
		return x.InnerDir
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UploadFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName      string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileBytesData []byte `protobuf:"bytes,2,opt,name=FileBytesData,proto3" json:"FileBytesData,omitempty"`
}

func (x *UploadFile) Reset() {
	*x = UploadFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFile) ProtoMessage() {}

func (x *UploadFile) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFile.ProtoReflect.Descriptor instead.
func (*UploadFile) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{7}
}

func (x *UploadFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFile) GetFileBytesData() []byte {
	if x != nil {
		return x.FileBytesData
	}
	return nil
}

type RequestUploadPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User          `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	UplFiles   []*UploadFile  `protobuf:"bytes,2,rep,name=UplFiles,proto3" json:"UplFiles,omitempty"`
	UplDirInfo *UploadDirInfo `protobuf:"bytes,3,opt,name=UplDirInfo,proto3" json:"UplDirInfo,omitempty"`
	Version    string         `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *RequestUploadPackage) Reset() {
	*x = RequestUploadPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUploadPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUploadPackage) ProtoMessage() {}

func (x *RequestUploadPackage) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUploadPackage.ProtoReflect.Descriptor instead.
func (*RequestUploadPackage) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{8}
}

func (x *RequestUploadPackage) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RequestUploadPackage) GetUplFiles() []*UploadFile {
	if x != nil {
		return x.UplFiles
	}
	return nil
}

func (x *RequestUploadPackage) GetUplDirInfo() *UploadDirInfo {
	if x != nil {
		return x.UplDirInfo
	}
	return nil
}

func (x *RequestUploadPackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_package_proto protoreflect.FileDescriptor

var file_package_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1d, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x68, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6b, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6b, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x55, 0x70, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x55, 0x70, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x55, 0x70, 0x6c, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x4a, 0x73,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x69, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x16,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6b,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6b, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x76,
	0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x08, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x69, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x44, 0x69, 0x72, 0x22, 0x3e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbf, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x55, 0x70, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x44, 0x69, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6b, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x55, 0x70, 0x6c, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xd8, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0d, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x1a, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0f,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x1a, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0c, 0x53, 0x68, 0x6f,
	0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6b, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x77,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x27, 0x2e, 0x70, 0x6b, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x77, 0x69, 0x6e, 0x65, 0x34, 0x38, 0x35, 0x30, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x78, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x6b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_package_proto_rawDescOnce sync.Once
	file_package_proto_rawDescData = file_package_proto_rawDesc
)

func file_package_proto_rawDescGZIP() []byte {
	file_package_proto_rawDescOnce.Do(func() {
		file_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_proto_rawDescData)
	})
	return file_package_proto_rawDescData
}

var file_package_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_package_proto_goTypes = []interface{}{
	(*MutateShowVersionBaseResponse)(nil), // 0: pkgproto.MutateShowVersionBaseResponse
	(*RequestShowVersions)(nil),           // 1: pkgproto.RequestShowVersions
	(*MutateDownloadBaseResponse)(nil),    // 2: pkgproto.MutateDownloadBaseResponse
	(*RequestDownloadPackage)(nil),        // 3: pkgproto.RequestDownloadPackage
	(*RequestPackageExists)(nil),          // 4: pkgproto.RequestPackageExists
	(*UploadDirInfo)(nil),                 // 5: pkgproto.UploadDirInfo
	(*User)(nil),                          // 6: pkgproto.User
	(*UploadFile)(nil),                    // 7: pkgproto.UploadFile
	(*RequestUploadPackage)(nil),          // 8: pkgproto.RequestUploadPackage
	(*baseproto.BaseResponse)(nil),        // 9: baseproto.BaseResponse
}
var file_package_proto_depIdxs = []int32{
	9,  // 0: pkgproto.MutateShowVersionBaseResponse.BaseResponse:type_name -> baseproto.BaseResponse
	9,  // 1: pkgproto.MutateDownloadBaseResponse.BaseResponse:type_name -> baseproto.BaseResponse
	5,  // 2: pkgproto.MutateDownloadBaseResponse.UplFiles:type_name -> pkgproto.UploadDirInfo
	5,  // 3: pkgproto.MutateDownloadBaseResponse.UplDirInfo:type_name -> pkgproto.UploadDirInfo
	5,  // 4: pkgproto.UploadDirInfo.InnerDir:type_name -> pkgproto.UploadDirInfo
	6,  // 5: pkgproto.RequestUploadPackage.User:type_name -> pkgproto.User
	7,  // 6: pkgproto.RequestUploadPackage.UplFiles:type_name -> pkgproto.UploadFile
	5,  // 7: pkgproto.RequestUploadPackage.UplDirInfo:type_name -> pkgproto.UploadDirInfo
	8,  // 8: pkgproto.Package.UploadPackage:input_type -> pkgproto.RequestUploadPackage
	4,  // 9: pkgproto.Package.ExistsPackage:input_type -> pkgproto.RequestPackageExists
	3,  // 10: pkgproto.Package.DownloadPackage:input_type -> pkgproto.RequestDownloadPackage
	1,  // 11: pkgproto.Package.ShowVersions:input_type -> pkgproto.RequestShowVersions
	9,  // 12: pkgproto.Package.UploadPackage:output_type -> baseproto.BaseResponse
	9,  // 13: pkgproto.Package.ExistsPackage:output_type -> baseproto.BaseResponse
	2,  // 14: pkgproto.Package.DownloadPackage:output_type -> pkgproto.MutateDownloadBaseResponse
	0,  // 15: pkgproto.Package.ShowVersions:output_type -> pkgproto.MutateShowVersionBaseResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_package_proto_init() }
func file_package_proto_init() {
	if File_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateShowVersionBaseResponse); i {
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
		file_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestShowVersions); i {
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
		file_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateDownloadBaseResponse); i {
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
		file_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDownloadPackage); i {
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
		file_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPackageExists); i {
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
		file_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDirInfo); i {
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
		file_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFile); i {
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
		file_package_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUploadPackage); i {
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
			RawDescriptor: file_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_proto_goTypes,
		DependencyIndexes: file_package_proto_depIdxs,
		MessageInfos:      file_package_proto_msgTypes,
	}.Build()
	File_package_proto = out.File
	file_package_proto_rawDesc = nil
	file_package_proto_goTypes = nil
	file_package_proto_depIdxs = nil
}
