// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/user_apiv1.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username             string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirmation string `protobuf:"bytes,4,opt,name=passwordConfirmation,proto3" json:"passwordConfirmation,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_apiv1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_apiv1_proto_msgTypes[0]
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
	return file_proto_user_apiv1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetPasswordConfirmation() string {
	if x != nil {
		return x.PasswordConfirmation
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username         string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Gender           int32  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Email            string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ThumbnailUrl     string `protobuf:"bytes,5,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	SelfIntroduction string `protobuf:"bytes,6,opt,name=self_introduction,json=selfIntroduction,proto3" json:"self_introduction,omitempty"`
	Lastname         string `protobuf:"bytes,7,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Firstname        string `protobuf:"bytes,8,opt,name=firstname,proto3" json:"firstname,omitempty"`
	LastnameKana     string `protobuf:"bytes,9,opt,name=lastname_kana,json=lastnameKana,proto3" json:"lastname_kana,omitempty"`
	FirstnameKana    string `protobuf:"bytes,10,opt,name=firstname_kana,json=firstnameKana,proto3" json:"firstname_kana,omitempty"`
	PostalCode       string `protobuf:"bytes,11,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Prefecture       string `protobuf:"bytes,12,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
	City             string `protobuf:"bytes,13,opt,name=city,proto3" json:"city,omitempty"`
	AddressLine1     string `protobuf:"bytes,14,opt,name=address_line1,json=addressLine1,proto3" json:"address_line1,omitempty"`
	AddressLine2     string `protobuf:"bytes,15,opt,name=address_line2,json=addressLine2,proto3" json:"address_line2,omitempty"`
	PhoneNumber      string `protobuf:"bytes,16,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Role             int32  `protobuf:"varint,17,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt        string `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_apiv1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_apiv1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_apiv1_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *UserResponse) GetSelfIntroduction() string {
	if x != nil {
		return x.SelfIntroduction
	}
	return ""
}

func (x *UserResponse) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserResponse) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserResponse) GetLastnameKana() string {
	if x != nil {
		return x.LastnameKana
	}
	return ""
}

func (x *UserResponse) GetFirstnameKana() string {
	if x != nil {
		return x.FirstnameKana
	}
	return ""
}

func (x *UserResponse) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UserResponse) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *UserResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserResponse) GetAddressLine1() string {
	if x != nil {
		return x.AddressLine1
	}
	return ""
}

func (x *UserResponse) GetAddressLine2() string {
	if x != nil {
		return x.AddressLine2
	}
	return ""
}

func (x *UserResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *UserResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_proto_user_apiv1_proto protoreflect.FileDescriptor

var file_proto_user_apiv1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd4, 0x04, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6b, 0x61, 0x6e, 0x61, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x4b, 0x61, 0x6e,
	0x61, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6b,
	0x61, 0x6e, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x4b, 0x61, 0x6e, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x31, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x32, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x60, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x24,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6c,
	0x6d, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_apiv1_proto_rawDescOnce sync.Once
	file_proto_user_apiv1_proto_rawDescData = file_proto_user_apiv1_proto_rawDesc
)

func file_proto_user_apiv1_proto_rawDescGZIP() []byte {
	file_proto_user_apiv1_proto_rawDescOnce.Do(func() {
		file_proto_user_apiv1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_apiv1_proto_rawDescData)
	})
	return file_proto_user_apiv1_proto_rawDescData
}

var file_proto_user_apiv1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_apiv1_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil), // 0: proto.CreateUserRequest
	(*UserResponse)(nil),      // 1: proto.UserResponse
}
var file_proto_user_apiv1_proto_depIdxs = []int32{
	0, // 0: proto.UserService.CreateUser:input_type -> proto.CreateUserRequest
	1, // 1: proto.UserService.CreateUser:output_type -> proto.UserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_apiv1_proto_init() }
func file_proto_user_apiv1_proto_init() {
	if File_proto_user_apiv1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_apiv1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_user_apiv1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_proto_user_apiv1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_apiv1_proto_goTypes,
		DependencyIndexes: file_proto_user_apiv1_proto_depIdxs,
		MessageInfos:      file_proto_user_apiv1_proto_msgTypes,
	}.Build()
	File_proto_user_apiv1_proto = out.File
	file_proto_user_apiv1_proto_rawDesc = nil
	file_proto_user_apiv1_proto_goTypes = nil
	file_proto_user_apiv1_proto_depIdxs = nil
}
