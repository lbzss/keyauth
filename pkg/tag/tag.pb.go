//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/tag/pb/tag.proto

package tag

import (
	proto "github.com/golang/protobuf/proto"
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

// Tag is 标签
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间`
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// Tag所属的域
	Domain string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建人
	Creater string `protobuf:"bytes,5,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// tag生效范围
	ScopeType ScopeType `protobuf:"varint,6,opt,name=scope_type,json=scopeType,proto3,enum=keyauth.tag.ScopeType" json:"scope_type" bson:"scope_type"`
	// tag属于哪个namespace
	Namespace string `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 建名称
	KeyName string `protobuf:"bytes,8,opt,name=key_name,json=keyName,proto3" json:"key_name" bson:"key_name"`
	// 建标识
	KeyLabel string `protobuf:"bytes,9,opt,name=key_label,json=keyLabel,proto3" json:"key_label" bson:"key_label"`
	// 建描述
	KeyDesc string `protobuf:"bytes,10,opt,name=key_desc,json=keyDesc,proto3" json:"key_desc" bson:"key_desc"`
	// 值来源
	ValueFrom ValueFrom `protobuf:"varint,11,opt,name=value_from,json=valueFrom,proto3,enum=keyauth.tag.ValueFrom" json:"value_from" bson:"value_from"`
	// http 获取Tag 值的参数
	HttpFromOption *HTTPFromOption `protobuf:"bytes,12,opt,name=http_from_option,json=httpFromOption,proto3" json:"http_from_option" bson:"http_from_option"`
	Values         []*TagValue     `protobuf:"bytes,13,rep,name=values,proto3" json:"values" bson:"-"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Tag) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Tag) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Tag) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Tag) GetScopeType() ScopeType {
	if x != nil {
		return x.ScopeType
	}
	return ScopeType_NAMESPACE
}

func (x *Tag) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Tag) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *Tag) GetKeyLabel() string {
	if x != nil {
		return x.KeyLabel
	}
	return ""
}

func (x *Tag) GetKeyDesc() string {
	if x != nil {
		return x.KeyDesc
	}
	return ""
}

func (x *Tag) GetValueFrom() ValueFrom {
	if x != nil {
		return x.ValueFrom
	}
	return ValueFrom_MANUAL
}

func (x *Tag) GetHttpFromOption() *HTTPFromOption {
	if x != nil {
		return x.HttpFromOption
	}
	return nil
}

func (x *Tag) GetValues() []*TagValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type TagValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag Value ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间`
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 创建人
	Creater string `protobuf:"bytes,4,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 关联的Tag key
	KeyId string `protobuf:"bytes,5,opt,name=key_id,json=keyId,proto3" json:"key_id" bson:"key_id"`
	// String 类型的值
	Value *ValueOption `protobuf:"bytes,6,opt,name=value,proto3" json:"value" bson:"value"`
}

func (x *TagValue) Reset() {
	*x = TagValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagValue) ProtoMessage() {}

func (x *TagValue) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagValue.ProtoReflect.Descriptor instead.
func (*TagValue) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagValue) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *TagValue) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *TagValue) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *TagValue) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *TagValue) GetValue() *ValueOption {
	if x != nil {
		return x.Value
	}
	return nil
}

type TagSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64  `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Tag `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *TagSet) Reset() {
	*x = TagSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagSet) ProtoMessage() {}

func (x *TagSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagSet.ProtoReflect.Descriptor instead.
func (*TagSet) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{2}
}

func (x *TagSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TagSet) GetItems() []*Tag {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_tag_pb_tag_proto protoreflect.FileDescriptor

var file_pkg_tag_pb_tag_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x74, 0x61, 0x67, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x07, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x2a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a,
	0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde,
	0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x44, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12,
	0x60, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x29, 0xc2, 0xde, 0x1f,
	0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x45, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21,
	0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x6b, 0x65,
	0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x40, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x22, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x60, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x29,
	0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x7c, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x35, 0xc2, 0xde, 0x1f,
	0x31, 0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x52, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1c, 0xc2, 0xde, 0x1f, 0x18, 0x0a,
	0x16, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0x8c, 0x03, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f,
	0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x4f, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x88,
	0x01, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x47, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61,
	0x67, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74,
	0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_tag_pb_tag_proto_rawDescOnce sync.Once
	file_pkg_tag_pb_tag_proto_rawDescData = file_pkg_tag_pb_tag_proto_rawDesc
)

func file_pkg_tag_pb_tag_proto_rawDescGZIP() []byte {
	file_pkg_tag_pb_tag_proto_rawDescOnce.Do(func() {
		file_pkg_tag_pb_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_tag_pb_tag_proto_rawDescData)
	})
	return file_pkg_tag_pb_tag_proto_rawDescData
}

var file_pkg_tag_pb_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_tag_pb_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),            // 0: keyauth.tag.Tag
	(*TagValue)(nil),       // 1: keyauth.tag.TagValue
	(*TagSet)(nil),         // 2: keyauth.tag.TagSet
	(ScopeType)(0),         // 3: keyauth.tag.ScopeType
	(ValueFrom)(0),         // 4: keyauth.tag.ValueFrom
	(*HTTPFromOption)(nil), // 5: keyauth.tag.HTTPFromOption
	(*ValueOption)(nil),    // 6: keyauth.tag.ValueOption
}
var file_pkg_tag_pb_tag_proto_depIdxs = []int32{
	3, // 0: keyauth.tag.Tag.scope_type:type_name -> keyauth.tag.ScopeType
	4, // 1: keyauth.tag.Tag.value_from:type_name -> keyauth.tag.ValueFrom
	5, // 2: keyauth.tag.Tag.http_from_option:type_name -> keyauth.tag.HTTPFromOption
	1, // 3: keyauth.tag.Tag.values:type_name -> keyauth.tag.TagValue
	6, // 4: keyauth.tag.TagValue.value:type_name -> keyauth.tag.ValueOption
	0, // 5: keyauth.tag.TagSet.items:type_name -> keyauth.tag.Tag
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_tag_pb_tag_proto_init() }
func file_pkg_tag_pb_tag_proto_init() {
	if File_pkg_tag_pb_tag_proto != nil {
		return
	}
	file_pkg_tag_pb_enum_proto_init()
	file_pkg_tag_pb_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_tag_pb_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_pkg_tag_pb_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagValue); i {
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
		file_pkg_tag_pb_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagSet); i {
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
			RawDescriptor: file_pkg_tag_pb_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_tag_pb_tag_proto_goTypes,
		DependencyIndexes: file_pkg_tag_pb_tag_proto_depIdxs,
		MessageInfos:      file_pkg_tag_pb_tag_proto_msgTypes,
	}.Build()
	File_pkg_tag_pb_tag_proto = out.File
	file_pkg_tag_pb_tag_proto_rawDesc = nil
	file_pkg_tag_pb_tag_proto_goTypes = nil
	file_pkg_tag_pb_tag_proto_depIdxs = nil
}
