// Code generated with goa v3.12.4, DO NOT EDIT.
//
// front protocol buffer definition
//
// Command:
// $ goa gen game-service/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: goadesign_goagen_front.proto

package frontpb

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

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_front_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_front_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_front_proto_rawDescGZIP(), []int{0}
}

type StoredItemCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []*StoredItem `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *StoredItemCollection) Reset() {
	*x = StoredItemCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_front_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredItemCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredItemCollection) ProtoMessage() {}

func (x *StoredItemCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_front_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredItemCollection.ProtoReflect.Descriptor instead.
func (*StoredItemCollection) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_front_proto_rawDescGZIP(), []int{1}
}

func (x *StoredItemCollection) GetField() []*StoredItem {
	if x != nil {
		return x.Field
	}
	return nil
}

// A StoredItem describes an item stored with an ID
type StoredItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique id of the item.
	Id string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// Name of item
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of item
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Amount of damage the item can do
	Damage uint32 `protobuf:"varint,4,opt,name=damage,proto3" json:"damage,omitempty"`
	// Amount of healing the item can provide
	Healing uint32 `protobuf:"varint,5,opt,name=healing,proto3" json:"healing,omitempty"`
	// Amount of protection the item can provide
	Protection uint32 `protobuf:"varint,6,opt,name=protection,proto3" json:"protection,omitempty"`
}

func (x *StoredItem) Reset() {
	*x = StoredItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_front_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredItem) ProtoMessage() {}

func (x *StoredItem) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_front_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredItem.ProtoReflect.Descriptor instead.
func (*StoredItem) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_front_proto_rawDescGZIP(), []int{2}
}

func (x *StoredItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoredItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredItem) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *StoredItem) GetDamage() uint32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *StoredItem) GetHealing() uint32 {
	if x != nil {
		return x.Healing
	}
	return 0
}

func (x *StoredItem) GetProtection() uint32 {
	if x != nil {
		return x.Protection
	}
	return 0
}

var File_goadesign_goagen_front_proto protoreflect.FileDescriptor

var file_goadesign_goagen_front_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x4a, 0x0a, 0x05, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_front_proto_rawDescOnce sync.Once
	file_goadesign_goagen_front_proto_rawDescData = file_goadesign_goagen_front_proto_rawDesc
)

func file_goadesign_goagen_front_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_front_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_front_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_front_proto_rawDescData)
	})
	return file_goadesign_goagen_front_proto_rawDescData
}

var file_goadesign_goagen_front_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goadesign_goagen_front_proto_goTypes = []interface{}{
	(*ListItemsRequest)(nil),     // 0: front.ListItemsRequest
	(*StoredItemCollection)(nil), // 1: front.StoredItemCollection
	(*StoredItem)(nil),           // 2: front.StoredItem
}
var file_goadesign_goagen_front_proto_depIdxs = []int32{
	2, // 0: front.StoredItemCollection.field:type_name -> front.StoredItem
	0, // 1: front.Front.ListItems:input_type -> front.ListItemsRequest
	1, // 2: front.Front.ListItems:output_type -> front.StoredItemCollection
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_front_proto_init() }
func file_goadesign_goagen_front_proto_init() {
	if File_goadesign_goagen_front_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_front_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
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
		file_goadesign_goagen_front_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredItemCollection); i {
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
		file_goadesign_goagen_front_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredItem); i {
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
	file_goadesign_goagen_front_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goadesign_goagen_front_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_front_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_front_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_front_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_front_proto = out.File
	file_goadesign_goagen_front_proto_rawDesc = nil
	file_goadesign_goagen_front_proto_goTypes = nil
	file_goadesign_goagen_front_proto_depIdxs = nil
}
