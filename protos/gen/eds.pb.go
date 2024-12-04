// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: eds.proto

package edsv1

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

type AddNewDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorEmail string `protobuf:"bytes,1,opt,name=creator_email,json=creatorEmail,proto3" json:"creator_email,omitempty"`
}

func (x *AddNewDocumentRequest) Reset() {
	*x = AddNewDocumentRequest{}
	mi := &file_eds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddNewDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewDocumentRequest) ProtoMessage() {}

func (x *AddNewDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewDocumentRequest.ProtoReflect.Descriptor instead.
func (*AddNewDocumentRequest) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{0}
}

func (x *AddNewDocumentRequest) GetCreatorEmail() string {
	if x != nil {
		return x.CreatorEmail
	}
	return ""
}

type AddNewDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrCode string `protobuf:"bytes,1,opt,name=qr_code,json=qrCode,proto3" json:"qr_code,omitempty"`
}

func (x *AddNewDocumentResponse) Reset() {
	*x = AddNewDocumentResponse{}
	mi := &file_eds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddNewDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewDocumentResponse) ProtoMessage() {}

func (x *AddNewDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewDocumentResponse.ProtoReflect.Descriptor instead.
func (*AddNewDocumentResponse) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{1}
}

func (x *AddNewDocumentResponse) GetQrCode() string {
	if x != nil {
		return x.QrCode
	}
	return ""
}

type SendDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId  string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	SignersMail []string `protobuf:"bytes,2,rep,name=signers_mail,json=signersMail,proto3" json:"signers_mail,omitempty"`
	FilesUrl    []string `protobuf:"bytes,3,rep,name=files_url,json=filesUrl,proto3" json:"files_url,omitempty"`
}

func (x *SendDocumentRequest) Reset() {
	*x = SendDocumentRequest{}
	mi := &file_eds_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDocumentRequest) ProtoMessage() {}

func (x *SendDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDocumentRequest.ProtoReflect.Descriptor instead.
func (*SendDocumentRequest) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{2}
}

func (x *SendDocumentRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *SendDocumentRequest) GetSignersMail() []string {
	if x != nil {
		return x.SignersMail
	}
	return nil
}

func (x *SendDocumentRequest) GetFilesUrl() []string {
	if x != nil {
		return x.FilesUrl
	}
	return nil
}

type SendDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendDocumentResponse) Reset() {
	*x = SendDocumentResponse{}
	mi := &file_eds_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDocumentResponse) ProtoMessage() {}

func (x *SendDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDocumentResponse.ProtoReflect.Descriptor instead.
func (*SendDocumentResponse) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{3}
}

func (x *SendDocumentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetDocumentByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId string `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *GetDocumentByIDRequest) Reset() {
	*x = GetDocumentByIDRequest{}
	mi := &file_eds_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDocumentByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentByIDRequest) ProtoMessage() {}

func (x *GetDocumentByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentByIDRequest.ProtoReflect.Descriptor instead.
func (*GetDocumentByIDRequest) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{4}
}

func (x *GetDocumentByIDRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

type GetDocumentByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Creator    *User    `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Signers    []*User  `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
	FilesUrl   []string `protobuf:"bytes,4,rep,name=files_url,json=filesUrl,proto3" json:"files_url,omitempty"`
}

func (x *GetDocumentByIDResponse) Reset() {
	*x = GetDocumentByIDResponse{}
	mi := &file_eds_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDocumentByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentByIDResponse) ProtoMessage() {}

func (x *GetDocumentByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentByIDResponse.ProtoReflect.Descriptor instead.
func (*GetDocumentByIDResponse) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{5}
}

func (x *GetDocumentByIDResponse) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *GetDocumentByIDResponse) GetCreator() *User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *GetDocumentByIDResponse) GetSigners() []*User {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *GetDocumentByIDResponse) GetFilesUrl() []string {
	if x != nil {
		return x.FilesUrl
	}
	return nil
}

type SignDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId       string `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	SignerMail       string `protobuf:"bytes,2,opt,name=signer_mail,json=signerMail,proto3" json:"signer_mail,omitempty"`
	SignatureFileUrl string `protobuf:"bytes,3,opt,name=signature_file_url,json=signatureFileUrl,proto3" json:"signature_file_url,omitempty"`
}

func (x *SignDocumentRequest) Reset() {
	*x = SignDocumentRequest{}
	mi := &file_eds_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignDocumentRequest) ProtoMessage() {}

func (x *SignDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignDocumentRequest.ProtoReflect.Descriptor instead.
func (*SignDocumentRequest) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{6}
}

func (x *SignDocumentRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *SignDocumentRequest) GetSignerMail() string {
	if x != nil {
		return x.SignerMail
	}
	return ""
}

func (x *SignDocumentRequest) GetSignatureFileUrl() string {
	if x != nil {
		return x.SignatureFileUrl
	}
	return ""
}

type SignDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SignDocumentResponse) Reset() {
	*x = SignDocumentResponse{}
	mi := &file_eds_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignDocumentResponse) ProtoMessage() {}

func (x *SignDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignDocumentResponse.ProtoReflect.Descriptor instead.
func (*SignDocumentResponse) Descriptor() ([]byte, []int) {
	return file_eds_proto_rawDescGZIP(), []int{7}
}

func (x *SignDocumentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserName  string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_eds_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_eds_proto_msgTypes[8]
	if x != nil {
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
	return file_eds_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_eds_proto protoreflect.FileDescriptor

var file_eds_proto_rawDesc = []byte{
	0x0a, 0x09, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x64, 0x73,
	0x22, 0x3c, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x31,
	0x0a, 0x16, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x76, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65, 0x6e,
	0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x39, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x65, 0x64, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x53,
	0x69, 0x67, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x4d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xa8, 0x02, 0x0a, 0x03, 0x45, 0x44, 0x53, 0x12, 0x49, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x4e, 0x65, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x64,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x4e, 0x65, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65,
	0x64, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x65, 0x64, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x66, 0x69, 0x66, 0x31,
	0x39, 0x38, 0x32, 0x2f, 0x65, 0x64, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x65, 0x64, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eds_proto_rawDescOnce sync.Once
	file_eds_proto_rawDescData = file_eds_proto_rawDesc
)

func file_eds_proto_rawDescGZIP() []byte {
	file_eds_proto_rawDescOnce.Do(func() {
		file_eds_proto_rawDescData = protoimpl.X.CompressGZIP(file_eds_proto_rawDescData)
	})
	return file_eds_proto_rawDescData
}

var file_eds_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_eds_proto_goTypes = []any{
	(*AddNewDocumentRequest)(nil),   // 0: eds.AddNewDocumentRequest
	(*AddNewDocumentResponse)(nil),  // 1: eds.AddNewDocumentResponse
	(*SendDocumentRequest)(nil),     // 2: eds.SendDocumentRequest
	(*SendDocumentResponse)(nil),    // 3: eds.SendDocumentResponse
	(*GetDocumentByIDRequest)(nil),  // 4: eds.GetDocumentByIDRequest
	(*GetDocumentByIDResponse)(nil), // 5: eds.GetDocumentByIDResponse
	(*SignDocumentRequest)(nil),     // 6: eds.SignDocumentRequest
	(*SignDocumentResponse)(nil),    // 7: eds.SignDocumentResponse
	(*User)(nil),                    // 8: eds.User
}
var file_eds_proto_depIdxs = []int32{
	8, // 0: eds.GetDocumentByIDResponse.creator:type_name -> eds.User
	8, // 1: eds.GetDocumentByIDResponse.signers:type_name -> eds.User
	0, // 2: eds.EDS.AddNewDocument:input_type -> eds.AddNewDocumentRequest
	4, // 3: eds.EDS.GetDocumentByID:input_type -> eds.GetDocumentByIDRequest
	2, // 4: eds.EDS.SendDocument:input_type -> eds.SendDocumentRequest
	6, // 5: eds.EDS.SignDocument:input_type -> eds.SignDocumentRequest
	1, // 6: eds.EDS.AddNewDocument:output_type -> eds.AddNewDocumentResponse
	5, // 7: eds.EDS.GetDocumentByID:output_type -> eds.GetDocumentByIDResponse
	3, // 8: eds.EDS.SendDocument:output_type -> eds.SendDocumentResponse
	7, // 9: eds.EDS.SignDocument:output_type -> eds.SignDocumentResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eds_proto_init() }
func file_eds_proto_init() {
	if File_eds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eds_proto_goTypes,
		DependencyIndexes: file_eds_proto_depIdxs,
		MessageInfos:      file_eds_proto_msgTypes,
	}.Build()
	File_eds_proto = out.File
	file_eds_proto_rawDesc = nil
	file_eds_proto_goTypes = nil
	file_eds_proto_depIdxs = nil
}
