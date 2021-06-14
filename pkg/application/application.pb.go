//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/application/pb/application.proto

package application

import (
	_ "github.com/infraboard/mcube/cmd/protoc-gen-go-ext/extension/tag"
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

// Application is oauth2's client: https://tools.ietf.org/html/rfc6749#section-2
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 是否是内建应用
	BuildIn bool `protobuf:"varint,2,opt,name=build_in,json=buildIn,proto3" json:"build_in" bson:"build_in"`
	// 所处于域
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 应用属于那个用户
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user" bson:"user"`
	// 应用创建的时间
	CreateAt int64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 应用更新的时间
	UpdateAt int64 `protobuf:"varint,6,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 应用客户端ID
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id" bson:"client_id"`
	// 应用客户端秘钥
	ClientSecret string `protobuf:"bytes,8,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" bson:"client_secret"`
	// 是否冻结应用, 冻结应用后, 该应用无法通过凭证获取访问凭证(token)
	Locked bool `protobuf:"varint,9,opt,name=locked,proto3" json:"locked" bson:"locked"`
	// 应用名称
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name" bson:"name"`
	// 应用的网站地址
	Website string `protobuf:"bytes,11,opt,name=website,proto3" json:"website" bson:"website"`
	// 应用的LOGO
	LogoImage string `protobuf:"bytes,12,opt,name=logo_image,json=logoImage,proto3" json:"logo_image" bson:"logo_image"`
	// 应用简单的描述
	Description string `protobuf:"bytes,13,opt,name=description,proto3" json:"description" bson:"description"`
	// 应用重定向URI, Oauht2时需要该参数
	RedirectUri string `protobuf:"bytes,14,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri" bson:"redirect_uri"`
	// 应用申请的token的过期时间
	AccessTokenExpireSecond int64 `protobuf:"varint,15,opt,name=access_token_expire_second,json=accessTokenExpireSecond,proto3" json:"access_token_expire_second" bson:"access_token_expire_second"`
	// 刷新token过期时间
	RefreshTokenExpireSecond int64 `protobuf:"varint,16,opt,name=refresh_token_expire_second,json=refreshTokenExpireSecond,proto3" json:"refresh_token_expire_second" bson:"refresh_token_expire_second"`
	// 客户端类型
	ClientType ClientType `protobuf:"varint,17,opt,name=client_type,json=clientType,proto3,enum=keyauth.application.ClientType" json:"client_type" bson:"client_type"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application) GetBuildIn() bool {
	if x != nil {
		return x.BuildIn
	}
	return false
}

func (x *Application) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Application) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Application) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Application) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Application) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Application) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Application) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Application) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *Application) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Application) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *Application) GetAccessTokenExpireSecond() int64 {
	if x != nil {
		return x.AccessTokenExpireSecond
	}
	return 0
}

func (x *Application) GetRefreshTokenExpireSecond() int64 {
	if x != nil {
		return x.RefreshTokenExpireSecond
	}
	return 0
}

func (x *Application) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_CONFIDENTIAL
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Application `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_application_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Application {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_application_pb_application_proto protoreflect.FileDescriptor

var file_pkg_application_pb_application_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x0a, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x22, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x44, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc2,
	0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x52, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x21, 0xc2, 0xde,
	0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x52,
	0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f,
	0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22,
	0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xc2,
	0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x6f, 0x67, 0x6f, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x22, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x69, 0x12, 0x86, 0x01, 0x0a, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x42, 0x49, 0xc2, 0xde, 0x1f, 0x45, 0x0a,
	0x43, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x22, 0x52, 0x17, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x8a, 0x01,
	0x0a, 0x1b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x4b, 0xc2, 0xde, 0x1f, 0x47, 0x0a, 0x45, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22,
	0x52, 0x18, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x6d, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x57, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a,
	0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_application_pb_application_proto_rawDescOnce sync.Once
	file_pkg_application_pb_application_proto_rawDescData = file_pkg_application_pb_application_proto_rawDesc
)

func file_pkg_application_pb_application_proto_rawDescGZIP() []byte {
	file_pkg_application_pb_application_proto_rawDescOnce.Do(func() {
		file_pkg_application_pb_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_application_pb_application_proto_rawDescData)
	})
	return file_pkg_application_pb_application_proto_rawDescData
}

var file_pkg_application_pb_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_application_pb_application_proto_goTypes = []interface{}{
	(*Application)(nil), // 0: keyauth.application.Application
	(*Set)(nil),         // 1: keyauth.application.Set
	(ClientType)(0),     // 2: keyauth.application.ClientType
}
var file_pkg_application_pb_application_proto_depIdxs = []int32{
	2, // 0: keyauth.application.Application.client_type:type_name -> keyauth.application.ClientType
	0, // 1: keyauth.application.Set.items:type_name -> keyauth.application.Application
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_application_pb_application_proto_init() }
func file_pkg_application_pb_application_proto_init() {
	if File_pkg_application_pb_application_proto != nil {
		return
	}
	file_pkg_application_pb_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_application_pb_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_pkg_application_pb_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
			RawDescriptor: file_pkg_application_pb_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_application_pb_application_proto_goTypes,
		DependencyIndexes: file_pkg_application_pb_application_proto_depIdxs,
		MessageInfos:      file_pkg_application_pb_application_proto_msgTypes,
	}.Build()
	File_pkg_application_pb_application_proto = out.File
	file_pkg_application_pb_application_proto_rawDesc = nil
	file_pkg_application_pb_application_proto_goTypes = nil
	file_pkg_application_pb_application_proto_depIdxs = nil
}
