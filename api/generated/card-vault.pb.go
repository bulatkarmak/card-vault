// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.2
// source: api/card-vault/card-vault.proto

package generated

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone          string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email          string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Age            int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	PassportNumber string `protobuf:"bytes,5,opt,name=passport_number,json=passportNumber,proto3" json:"passport_number,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateUserRequest) GetPassportNumber() string {
	if x != nil {
		return x.PassportNumber
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Age            int32  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	PassportNumber string `protobuf:"bytes,6,opt,name=passport_number,json=passportNumber,proto3" json:"passport_number,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserRequest) GetPassportNumber() string {
	if x != nil {
		return x.PassportNumber
	}
	return ""
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNumber string `protobuf:"bytes,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpiryDate string `protobuf:"bytes,2,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	UserId     int32  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BankId     int32  `protobuf:"varint,4,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
}

func (x *CreateCardRequest) Reset() {
	*x = CreateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardRequest) ProtoMessage() {}

func (x *CreateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardRequest.ProtoReflect.Descriptor instead.
func (*CreateCardRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CreateCardRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *CreateCardRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCardRequest) GetBankId() int32 {
	if x != nil {
		return x.BankId
	}
	return 0
}

type CreateCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateCardResponse) Reset() {
	*x = CreateCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardResponse) ProtoMessage() {}

func (x *CreateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardResponse.ProtoReflect.Descriptor instead.
func (*CreateCardResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCardResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CardNumber string `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpiryDate string `protobuf:"bytes,3,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	UserId     int32  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BankId     int32  `protobuf:"varint,5,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
}

func (x *UpdateCardRequest) Reset() {
	*x = UpdateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardRequest) ProtoMessage() {}

func (x *UpdateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardRequest.ProtoReflect.Descriptor instead.
func (*UpdateCardRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCardRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *UpdateCardRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *UpdateCardRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateCardRequest) GetBankId() int32 {
	if x != nil {
		return x.BankId
	}
	return 0
}

type UpdateCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateCardResponse) Reset() {
	*x = UpdateCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardResponse) ProtoMessage() {}

func (x *UpdateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardResponse.ProtoReflect.Descriptor instead.
func (*UpdateCardResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCardRequest) Reset() {
	*x = DeleteCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardRequest) ProtoMessage() {}

func (x *DeleteCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteCardRequest) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCardRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteCardResponse) Reset() {
	*x = DeleteCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_card_vault_card_vault_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardResponse) ProtoMessage() {}

func (x *DeleteCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_card_vault_card_vault_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardResponse.ProtoReflect.Descriptor instead.
func (*DeleteCardResponse) Descriptor() ([]byte, []int) {
	return file_api_card_vault_card_vault_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteCardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_card_vault_card_vault_proto protoreflect.FileDescriptor

var file_api_card_vault_card_vault_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2d, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2f, 0x63, 0x61, 0x72, 0x64, 0x2d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x8e, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3c, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x22,
	0x2c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0xcf, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x1c, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_card_vault_card_vault_proto_rawDescOnce sync.Once
	file_api_card_vault_card_vault_proto_rawDescData = file_api_card_vault_card_vault_proto_rawDesc
)

func file_api_card_vault_card_vault_proto_rawDescGZIP() []byte {
	file_api_card_vault_card_vault_proto_rawDescOnce.Do(func() {
		file_api_card_vault_card_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_card_vault_card_vault_proto_rawDescData)
	})
	return file_api_card_vault_card_vault_proto_rawDescData
}

var file_api_card_vault_card_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_card_vault_card_vault_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: cardvault.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: cardvault.CreateUserResponse
	(*UpdateUserRequest)(nil),  // 2: cardvault.UpdateUserRequest
	(*UpdateUserResponse)(nil), // 3: cardvault.UpdateUserResponse
	(*DeleteUserRequest)(nil),  // 4: cardvault.DeleteUserRequest
	(*DeleteUserResponse)(nil), // 5: cardvault.DeleteUserResponse
	(*CreateCardRequest)(nil),  // 6: cardvault.CreateCardRequest
	(*CreateCardResponse)(nil), // 7: cardvault.CreateCardResponse
	(*UpdateCardRequest)(nil),  // 8: cardvault.UpdateCardRequest
	(*UpdateCardResponse)(nil), // 9: cardvault.UpdateCardResponse
	(*DeleteCardRequest)(nil),  // 10: cardvault.DeleteCardRequest
	(*DeleteCardResponse)(nil), // 11: cardvault.DeleteCardResponse
}
var file_api_card_vault_card_vault_proto_depIdxs = []int32{
	0,  // 0: cardvault.UserService.CreateUser:input_type -> cardvault.CreateUserRequest
	2,  // 1: cardvault.UserService.UpdateUser:input_type -> cardvault.UpdateUserRequest
	4,  // 2: cardvault.UserService.DeleteUser:input_type -> cardvault.DeleteUserRequest
	6,  // 3: cardvault.UserService.CreateCard:input_type -> cardvault.CreateCardRequest
	8,  // 4: cardvault.UserService.UpdateCard:input_type -> cardvault.UpdateCardRequest
	10, // 5: cardvault.UserService.DeleteCard:input_type -> cardvault.DeleteCardRequest
	1,  // 6: cardvault.UserService.CreateUser:output_type -> cardvault.CreateUserResponse
	3,  // 7: cardvault.UserService.UpdateUser:output_type -> cardvault.UpdateUserResponse
	5,  // 8: cardvault.UserService.DeleteUser:output_type -> cardvault.DeleteUserResponse
	7,  // 9: cardvault.UserService.CreateCard:output_type -> cardvault.CreateCardResponse
	9,  // 10: cardvault.UserService.UpdateCard:output_type -> cardvault.UpdateCardResponse
	11, // 11: cardvault.UserService.DeleteCard:output_type -> cardvault.DeleteCardResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_card_vault_card_vault_proto_init() }
func file_api_card_vault_card_vault_proto_init() {
	if File_api_card_vault_card_vault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_card_vault_card_vault_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserResponse); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserResponse); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserResponse); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCardRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCardResponse); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCardRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCardResponse); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCardRequest); i {
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
		file_api_card_vault_card_vault_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCardResponse); i {
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
			RawDescriptor: file_api_card_vault_card_vault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_card_vault_card_vault_proto_goTypes,
		DependencyIndexes: file_api_card_vault_card_vault_proto_depIdxs,
		MessageInfos:      file_api_card_vault_card_vault_proto_msgTypes,
	}.Build()
	File_api_card_vault_card_vault_proto = out.File
	file_api_card_vault_card_vault_proto_rawDesc = nil
	file_api_card_vault_card_vault_proto_goTypes = nil
	file_api_card_vault_card_vault_proto_depIdxs = nil
}
