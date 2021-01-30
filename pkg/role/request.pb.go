//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/role/pb/request.proto

package role

import (
	proto "github.com/golang/protobuf/proto"
	page "github.com/infraboard/mcube/pb/page"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
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

// CreateRoleRequest 创建应用请求
type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色类型
	Type RoleType `protobuf:"varint,1,opt,name=type,proto3,enum=keyauth.role.RoleType" json:"type" bson:"type"`
	// 应用名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name" validate:"required,lte=30"`
	// 应用简单的描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" bson:"description" validate:"lte=400"`
	// 读权限
	Permissions []*Permission `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty" bson:"permissions"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRoleRequest) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Permission 权限
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 效力
	Effect EffectType `protobuf:"varint,1,opt,name=effect,proto3,enum=keyauth.role.EffectType" json:"effect,omitempty" bson:"effect"`
	// 资源列表
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty" bson:"resource_name"`
	// 维度
	LabelKey string `protobuf:"bytes,3,opt,name=label_key,json=labelKey,proto3" json:"label_key,omitempty" bson:"label_key"`
	// 适配所有值
	MatchAll bool `protobuf:"varint,4,opt,name=match_all,json=matchAll,proto3" json:"match_all" bson:"match_all"`
	// 标识值
	LabelValues []string `protobuf:"bytes,5,rep,name=label_values,json=labelValues,proto3" json:"label_values,omitempty" bson:"label_values"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetEffect() EffectType {
	if x != nil {
		return x.Effect
	}
	return EffectType_ALLOW
}

func (x *Permission) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Permission) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *Permission) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *Permission) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

// QueryRoleRequest 列表查询
type QueryRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page" bson:"page"`
	Type            RoleType          `protobuf:"varint,2,opt,name=type,proto3,enum=keyauth.role.RoleType" json:"type" bson:"type"`
	WithPermissions bool              `protobuf:"varint,3,opt,name=with_permissions,json=withPermissions,proto3" json:"with_permissions" bson:"with_permissions"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRoleRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

func (x *QueryRoleRequest) GetWithPermissions() bool {
	if x != nil {
		return x.WithPermissions
	}
	return false
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required,lte=64"`
	WithPermissions bool     `protobuf:"varint,3,opt,name=with_permissions,json=withPermissions,proto3" json:"with_permissions" bson:"with_permissions"`
	Type            RoleType `protobuf:"varint,4,opt,name=type,proto3,enum=keyauth.role.RoleType" json:"type" bson:"type"`
}

func (x *DescribeRoleRequest) Reset() {
	*x = DescribeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRoleRequest) ProtoMessage() {}

func (x *DescribeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRoleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeRoleRequest) GetWithPermissions() bool {
	if x != nil {
		return x.WithPermissions
	}
	return false
}

func (x *DescribeRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

// DeleteRoleRequest role删除
type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required,lte=64"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pkg_role_pb_request_proto protoreflect.FileDescriptor

var file_pkg_role_pb_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1d,
	0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x56, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x42, 0xc2, 0xde, 0x1f, 0x3e, 0x0a, 0x3c, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x2c,
	0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74,
	0x65, 0x3d, 0x33, 0x30, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x48, 0xc2, 0xde, 0x1f, 0x44, 0x0a, 0x42, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x22, 0x6c, 0x74, 0x65, 0x3d, 0x34, 0x30, 0x30, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31, 0x0a, 0x2f, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbd, 0x03, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x06, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x5e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x39, 0xc2, 0xde, 0x1f, 0x35, 0x0a, 0x33, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2c,
	0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xc2, 0xde, 0x1f,
	0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b,
	0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x6b, 0x65, 0x79, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x27, 0xc2, 0xde, 0x1f,
	0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61,
	0x6c, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x61, 0x6c, 0x6c, 0x22, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x5a,
	0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x37, 0xc2, 0xde, 0x1f, 0x33, 0x0a, 0x31, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x85, 0x02, 0x0a, 0x10, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1d, 0xc2, 0xde, 0x1f,
	0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x60, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31,
	0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xaf, 0x02, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xc2, 0xde, 0x1f, 0x32, 0x0a,
	0x30, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x36, 0x34,
	0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31, 0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xc2, 0xde, 0x1f, 0x30, 0x0a, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x36, 0x34, 0x22, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_role_pb_request_proto_rawDescOnce sync.Once
	file_pkg_role_pb_request_proto_rawDescData = file_pkg_role_pb_request_proto_rawDesc
)

func file_pkg_role_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_role_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_role_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_role_pb_request_proto_rawDescData)
	})
	return file_pkg_role_pb_request_proto_rawDescData
}

var file_pkg_role_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_role_pb_request_proto_goTypes = []interface{}{
	(*CreateRoleRequest)(nil),   // 0: keyauth.role.CreateRoleRequest
	(*Permission)(nil),          // 1: keyauth.role.Permission
	(*QueryRoleRequest)(nil),    // 2: keyauth.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil), // 3: keyauth.role.DescribeRoleRequest
	(*DeleteRoleRequest)(nil),   // 4: keyauth.role.DeleteRoleRequest
	(RoleType)(0),               // 5: keyauth.role.RoleType
	(EffectType)(0),             // 6: keyauth.role.EffectType
	(*page.PageRequest)(nil),    // 7: page.PageRequest
}
var file_pkg_role_pb_request_proto_depIdxs = []int32{
	5, // 0: keyauth.role.CreateRoleRequest.type:type_name -> keyauth.role.RoleType
	1, // 1: keyauth.role.CreateRoleRequest.permissions:type_name -> keyauth.role.Permission
	6, // 2: keyauth.role.Permission.effect:type_name -> keyauth.role.EffectType
	7, // 3: keyauth.role.QueryRoleRequest.page:type_name -> page.PageRequest
	5, // 4: keyauth.role.QueryRoleRequest.type:type_name -> keyauth.role.RoleType
	5, // 5: keyauth.role.DescribeRoleRequest.type:type_name -> keyauth.role.RoleType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_role_pb_request_proto_init() }
func file_pkg_role_pb_request_proto_init() {
	if File_pkg_role_pb_request_proto != nil {
		return
	}
	file_pkg_role_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_role_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_pkg_role_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
			RawDescriptor: file_pkg_role_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_role_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_role_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_role_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_role_pb_request_proto = out.File
	file_pkg_role_pb_request_proto_rawDesc = nil
	file_pkg_role_pb_request_proto_goTypes = nil
	file_pkg_role_pb_request_proto_depIdxs = nil
}
