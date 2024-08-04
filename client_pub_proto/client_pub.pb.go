// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: client_pub.proto

package client_api_proto

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

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cpf         string `protobuf:"bytes,2,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClientRequest) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *CreateClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateClientRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateClientRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpf         string `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreateClientResponse) Reset() {
	*x = CreateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientResponse) ProtoMessage() {}

func (x *CreateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientResponse.ProtoReflect.Descriptor instead.
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateClientResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClientResponse) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *CreateClientResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateClientResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateClientResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateClientResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetClientByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetClientByIDRequest) Reset() {
	*x = GetClientByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByIDRequest) ProtoMessage() {}

func (x *GetClientByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByIDRequest.ProtoReflect.Descriptor instead.
func (*GetClientByIDRequest) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{2}
}

func (x *GetClientByIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetClientByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpf         string `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *GetClientByIDResponse) Reset() {
	*x = GetClientByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByIDResponse) ProtoMessage() {}

func (x *GetClientByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByIDResponse.ProtoReflect.Descriptor instead.
func (*GetClientByIDResponse) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientByIDResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetClientByIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetClientByIDResponse) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *GetClientByIDResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetClientByIDResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetClientByIDResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetClientByIDResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetClientByCPFRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpf string `protobuf:"bytes,1,opt,name=cpf,proto3" json:"cpf,omitempty"`
}

func (x *GetClientByCPFRequest) Reset() {
	*x = GetClientByCPFRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByCPFRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByCPFRequest) ProtoMessage() {}

func (x *GetClientByCPFRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByCPFRequest.ProtoReflect.Descriptor instead.
func (*GetClientByCPFRequest) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{4}
}

func (x *GetClientByCPFRequest) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

type GetClientByCPFResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpf         string `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *GetClientByCPFResponse) Reset() {
	*x = GetClientByCPFResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_pub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByCPFResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByCPFResponse) ProtoMessage() {}

func (x *GetClientByCPFResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_pub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByCPFResponse.ProtoReflect.Descriptor instead.
func (*GetClientByCPFResponse) Descriptor() ([]byte, []int) {
	return file_client_pub_proto_rawDescGZIP(), []int{5}
}

func (x *GetClientByCPFResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetClientByCPFResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetClientByCPFResponse) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *GetClientByCPFResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetClientByCPFResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetClientByCPFResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetClientByCPFResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_client_pub_proto protoreflect.FileDescriptor

var file_client_pub_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x43, 0x50, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66,
	0x22, 0xc2, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x43, 0x50, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x43,
	0x50, 0x46, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x43, 0x50, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x50, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x28, 0x5a, 0x26, 0x68, 0x66, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_client_pub_proto_rawDescOnce sync.Once
	file_client_pub_proto_rawDescData = file_client_pub_proto_rawDesc
)

func file_client_pub_proto_rawDescGZIP() []byte {
	file_client_pub_proto_rawDescOnce.Do(func() {
		file_client_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_pub_proto_rawDescData)
	})
	return file_client_pub_proto_rawDescData
}

var file_client_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_client_pub_proto_goTypes = []any{
	(*CreateClientRequest)(nil),    // 0: CreateClientRequest
	(*CreateClientResponse)(nil),   // 1: CreateClientResponse
	(*GetClientByIDRequest)(nil),   // 2: GetClientByIDRequest
	(*GetClientByIDResponse)(nil),  // 3: GetClientByIDResponse
	(*GetClientByCPFRequest)(nil),  // 4: GetClientByCPFRequest
	(*GetClientByCPFResponse)(nil), // 5: GetClientByCPFResponse
}
var file_client_pub_proto_depIdxs = []int32{
	2, // 0: Client.GetClientByID:input_type -> GetClientByIDRequest
	4, // 1: Client.GetClientByCPF:input_type -> GetClientByCPFRequest
	0, // 2: Client.CreateClient:input_type -> CreateClientRequest
	3, // 3: Client.GetClientByID:output_type -> GetClientByIDResponse
	5, // 4: Client.GetClientByCPF:output_type -> GetClientByCPFResponse
	1, // 5: Client.CreateClient:output_type -> CreateClientResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_pub_proto_init() }
func file_client_pub_proto_init() {
	if File_client_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateClientRequest); i {
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
		file_client_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateClientResponse); i {
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
		file_client_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetClientByIDRequest); i {
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
		file_client_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetClientByIDResponse); i {
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
		file_client_pub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetClientByCPFRequest); i {
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
		file_client_pub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetClientByCPFResponse); i {
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
			RawDescriptor: file_client_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_pub_proto_goTypes,
		DependencyIndexes: file_client_pub_proto_depIdxs,
		MessageInfos:      file_client_pub_proto_msgTypes,
	}.Build()
	File_client_pub_proto = out.File
	file_client_pub_proto_rawDesc = nil
	file_client_pub_proto_goTypes = nil
	file_client_pub_proto_depIdxs = nil
}
