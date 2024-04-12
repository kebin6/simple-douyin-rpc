// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: douyin.proto

package douyin

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

type BaseUUIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseUUIDResp) Reset() {
	*x = BaseUUIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUUIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUUIDResp) ProtoMessage() {}

func (x *BaseUUIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUUIDResp.ProtoReflect.Descriptor instead.
func (*BaseUUIDResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{0}
}

func (x *BaseUUIDResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaseUUIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SignatureReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url"`
}

func (x *SignatureReq) Reset() {
	*x = SignatureReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureReq) ProtoMessage() {}

func (x *SignatureReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureReq.ProtoReflect.Descriptor instead.
func (*SignatureReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{1}
}

func (x *SignatureReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type BaseIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDResp) Reset() {
	*x = BaseIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDResp) ProtoMessage() {}

func (x *BaseIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDResp.ProtoReflect.Descriptor instead.
func (*BaseIDResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{2}
}

func (x *BaseIDResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IDsUint32Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsUint32Req) Reset() {
	*x = IDsUint32Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsUint32Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsUint32Req) ProtoMessage() {}

func (x *IDsUint32Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsUint32Req.ProtoReflect.Descriptor instead.
func (*IDsUint32Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{3}
}

func (x *IDsUint32Req) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BaseIDInt64Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDInt64Resp) Reset() {
	*x = BaseIDInt64Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDInt64Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDInt64Resp) ProtoMessage() {}

func (x *BaseIDInt64Resp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDInt64Resp.ProtoReflect.Descriptor instead.
func (*BaseIDInt64Resp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{4}
}

func (x *BaseIDInt64Resp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDInt64Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PageInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
}

func (x *PageInfoReq) Reset() {
	*x = PageInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReq) ProtoMessage() {}

func (x *PageInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReq.ProtoReflect.Descriptor instead.
func (*PageInfoReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{5}
}

func (x *PageInfoReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfoReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type UUIDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids"`
}

func (x *UUIDsReq) Reset() {
	*x = UUIDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDsReq) ProtoMessage() {}

func (x *UUIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDsReq.ProtoReflect.Descriptor instead.
func (*UUIDsReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{6}
}

func (x *UUIDsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *UUIDReq) Reset() {
	*x = UUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDReq) ProtoMessage() {}

func (x *UUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDReq.ProtoReflect.Descriptor instead.
func (*UUIDReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{7}
}

func (x *UUIDReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SignatureResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature"`
	ClientKey string `protobuf:"bytes,2,opt,name=client_key,json=clientKey,proto3" json:"client_key"`
}

func (x *SignatureResp) Reset() {
	*x = SignatureResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureResp) ProtoMessage() {}

func (x *SignatureResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureResp.ProtoReflect.Descriptor instead.
func (*SignatureResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{8}
}

func (x *SignatureResp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SignatureResp) GetClientKey() string {
	if x != nil {
		return x.ClientKey
	}
	return ""
}

// base message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[9]
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
	return file_douyin_proto_rawDescGZIP(), []int{9}
}

type IDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsReq) Reset() {
	*x = IDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsReq) ProtoMessage() {}

func (x *IDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsReq.ProtoReflect.Descriptor instead.
func (*IDsReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{10}
}

func (x *IDsReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BaseIDUint32Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDUint32Resp) Reset() {
	*x = BaseIDUint32Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDUint32Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDUint32Resp) ProtoMessage() {}

func (x *BaseIDUint32Resp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDUint32Resp.ProtoReflect.Descriptor instead.
func (*BaseIDUint32Resp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{11}
}

func (x *BaseIDUint32Resp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDUint32Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IDInt64Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDInt64Req) Reset() {
	*x = IDInt64Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDInt64Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDInt64Req) ProtoMessage() {}

func (x *IDInt64Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDInt64Req.ProtoReflect.Descriptor instead.
func (*IDInt64Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{12}
}

func (x *IDInt64Req) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{13}
}

func (x *IDReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDInt32Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDInt32Req) Reset() {
	*x = IDInt32Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDInt32Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDInt32Req) ProtoMessage() {}

func (x *IDInt32Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDInt32Req.ProtoReflect.Descriptor instead.
func (*IDInt32Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{14}
}

func (x *IDInt32Req) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDsInt32Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsInt32Req) Reset() {
	*x = IDsInt32Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsInt32Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsInt32Req) ProtoMessage() {}

func (x *IDsInt32Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsInt32Req.ProtoReflect.Descriptor instead.
func (*IDsInt32Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{15}
}

func (x *IDsInt32Req) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type IDUint32Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDUint32Req) Reset() {
	*x = IDUint32Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDUint32Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDUint32Req) ProtoMessage() {}

func (x *IDUint32Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDUint32Req.ProtoReflect.Descriptor instead.
func (*IDUint32Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{16}
}

func (x *IDUint32Req) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BaseIDInt32Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDInt32Resp) Reset() {
	*x = BaseIDInt32Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDInt32Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDInt32Resp) ProtoMessage() {}

func (x *BaseIDInt32Resp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDInt32Resp.ProtoReflect.Descriptor instead.
func (*BaseIDInt32Resp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{17}
}

func (x *BaseIDInt32Resp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDInt32Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IDsInt64Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsInt64Req) Reset() {
	*x = IDsInt64Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsInt64Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsInt64Req) ProtoMessage() {}

func (x *IDsInt64Req) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsInt64Req.ProtoReflect.Descriptor instead.
func (*IDsInt64Req) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{18}
}

func (x *IDsInt64Req) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{19}
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_douyin_proto protoreflect.FileDescriptor

var file_douyin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x22, 0x30, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2e, 0x0a, 0x0a, 0x42, 0x61,
	0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0c, 0x49, 0x44,
	0x73, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x0f,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x3e, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x1c, 0x0a, 0x08, 0x55, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x19, 0x0a, 0x07, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x34, 0x0a,
	0x10, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x44, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x44,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0b, 0x49, 0x44, 0x73, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x44, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1f, 0x0a,
	0x0b, 0x49, 0x44, 0x73, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x1c,
	0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x76, 0x0a, 0x06,
	0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x0d, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_douyin_proto_rawDescOnce sync.Once
	file_douyin_proto_rawDescData = file_douyin_proto_rawDesc
)

func file_douyin_proto_rawDescGZIP() []byte {
	file_douyin_proto_rawDescOnce.Do(func() {
		file_douyin_proto_rawDescData = protoimpl.X.CompressGZIP(file_douyin_proto_rawDescData)
	})
	return file_douyin_proto_rawDescData
}

var file_douyin_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_douyin_proto_goTypes = []interface{}{
	(*BaseUUIDResp)(nil),     // 0: douyin.BaseUUIDResp
	(*SignatureReq)(nil),     // 1: douyin.SignatureReq
	(*BaseIDResp)(nil),       // 2: douyin.BaseIDResp
	(*IDsUint32Req)(nil),     // 3: douyin.IDsUint32Req
	(*BaseIDInt64Resp)(nil),  // 4: douyin.BaseIDInt64Resp
	(*PageInfoReq)(nil),      // 5: douyin.PageInfoReq
	(*UUIDsReq)(nil),         // 6: douyin.UUIDsReq
	(*UUIDReq)(nil),          // 7: douyin.UUIDReq
	(*SignatureResp)(nil),    // 8: douyin.SignatureResp
	(*Empty)(nil),            // 9: douyin.Empty
	(*IDsReq)(nil),           // 10: douyin.IDsReq
	(*BaseIDUint32Resp)(nil), // 11: douyin.BaseIDUint32Resp
	(*IDInt64Req)(nil),       // 12: douyin.IDInt64Req
	(*IDReq)(nil),            // 13: douyin.IDReq
	(*IDInt32Req)(nil),       // 14: douyin.IDInt32Req
	(*IDsInt32Req)(nil),      // 15: douyin.IDsInt32Req
	(*IDUint32Req)(nil),      // 16: douyin.IDUint32Req
	(*BaseIDInt32Resp)(nil),  // 17: douyin.BaseIDInt32Resp
	(*IDsInt64Req)(nil),      // 18: douyin.IDsInt64Req
	(*BaseResp)(nil),         // 19: douyin.BaseResp
}
var file_douyin_proto_depIdxs = []int32{
	9,  // 0: douyin.Douyin.initDatabase:input_type -> douyin.Empty
	1,  // 1: douyin.Douyin.GetSignature:input_type -> douyin.SignatureReq
	19, // 2: douyin.Douyin.initDatabase:output_type -> douyin.BaseResp
	8,  // 3: douyin.Douyin.GetSignature:output_type -> douyin.SignatureResp
	2,  // [2:4] is the sub-list for method output_type
	0,  // [0:2] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_douyin_proto_init() }
func file_douyin_proto_init() {
	if File_douyin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_douyin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUUIDResp); i {
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
		file_douyin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureReq); i {
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
		file_douyin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDResp); i {
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
		file_douyin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsUint32Req); i {
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
		file_douyin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDInt64Resp); i {
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
		file_douyin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReq); i {
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
		file_douyin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDsReq); i {
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
		file_douyin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDReq); i {
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
		file_douyin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureResp); i {
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
		file_douyin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_douyin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsReq); i {
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
		file_douyin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDUint32Resp); i {
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
		file_douyin_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDInt64Req); i {
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
		file_douyin_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_douyin_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDInt32Req); i {
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
		file_douyin_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsInt32Req); i {
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
		file_douyin_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDUint32Req); i {
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
		file_douyin_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDInt32Resp); i {
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
		file_douyin_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsInt64Req); i {
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
		file_douyin_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
			RawDescriptor: file_douyin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_douyin_proto_goTypes,
		DependencyIndexes: file_douyin_proto_depIdxs,
		MessageInfos:      file_douyin_proto_msgTypes,
	}.Build()
	File_douyin_proto = out.File
	file_douyin_proto_rawDesc = nil
	file_douyin_proto_goTypes = nil
	file_douyin_proto_depIdxs = nil
}