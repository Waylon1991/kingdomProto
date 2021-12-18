// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: Protocol/client/client.proto

package client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "protocol/common"
	errcode "protocol/errcode"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//登录走的http ✔
type CP int32

const (
	CP_Default           CP = 0
	CP_Error_Code_Resp   CP = 1   //统一的错误码返回
	CP_Valid_Key_Req     CP = 2   //验证 key ✔
	CP_Heart_Ping_Req    CP = 10  //心跳请求 空协议体 ✔
	CP_Heart_Ping_Resp   CP = 11  //心跳返回 空协议体
	CP_Enter_Lobby_Resp  CP = 101 //进入大厅返回 ✔
	CP_Create_Room_Req   CP = 102 //创建房间
	CP_Create_Room_Resp  CP = 103
	CP_Join_Room_Req     CP = 104 //加入房间
	CP_Enter_Match_Req   CP = 105 //进入匹配 ✔
	CP_Enter_Match_Wait  CP = 106 //进入匹配等待
	CP_Cancel_Match      CP = 107 //取消匹配
	CP_Cancel_Match_Resp CP = 108 //取消匹配结果
	CP_Match_Suc_Resp    CP = 109 //匹配成功返回
	CP_Match_Fail_Resp   CP = 110 //匹配失败返回
)

// Enum value maps for CP.
var (
	CP_name = map[int32]string{
		0:   "Default",
		1:   "Error_Code_Resp",
		2:   "Valid_Key_Req",
		10:  "Heart_Ping_Req",
		11:  "Heart_Ping_Resp",
		101: "Enter_Lobby_Resp",
		102: "Create_Room_Req",
		103: "Create_Room_Resp",
		104: "Join_Room_Req",
		105: "Enter_Match_Req",
		106: "Enter_Match_Wait",
		107: "Cancel_Match",
		108: "Cancel_Match_Resp",
		109: "Match_Suc_Resp",
		110: "Match_Fail_Resp",
	}
	CP_value = map[string]int32{
		"Default":           0,
		"Error_Code_Resp":   1,
		"Valid_Key_Req":     2,
		"Heart_Ping_Req":    10,
		"Heart_Ping_Resp":   11,
		"Enter_Lobby_Resp":  101,
		"Create_Room_Req":   102,
		"Create_Room_Resp":  103,
		"Join_Room_Req":     104,
		"Enter_Match_Req":   105,
		"Enter_Match_Wait":  106,
		"Cancel_Match":      107,
		"Cancel_Match_Resp": 108,
		"Match_Suc_Resp":    109,
		"Match_Fail_Resp":   110,
	}
)

func (x CP) Enum() *CP {
	p := new(CP)
	*p = x
	return p
}

func (x CP) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CP) Descriptor() protoreflect.EnumDescriptor {
	return file_Protocol_client_client_proto_enumTypes[0].Descriptor()
}

func (CP) Type() protoreflect.EnumType {
	return &file_Protocol_client_client_proto_enumTypes[0]
}

func (x CP) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CP.Descriptor instead.
func (CP) EnumDescriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{0}
}

type ErrorCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code errcode.Error `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.errcode.Error" json:"code,omitempty"` //错误码
}

func (x *ErrorCode) Reset() {
	*x = ErrorCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorCode) ProtoMessage() {}

func (x *ErrorCode) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorCode.ProtoReflect.Descriptor instead.
func (*ErrorCode) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorCode) GetCode() errcode.Error {
	if x != nil {
		return x.Code
	}
	return errcode.Error(0)
}

type ValidKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //用户 id
	Key    []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                      //key
}

func (x *ValidKeyReq) Reset() {
	*x = ValidKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidKeyReq) ProtoMessage() {}

func (x *ValidKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidKeyReq.ProtoReflect.Descriptor instead.
func (*ValidKeyReq) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{1}
}

func (x *ValidKeyReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ValidKeyReq) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type EnterLobbyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //用户 id
}

func (x *EnterLobbyResp) Reset() {
	*x = EnterLobbyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterLobbyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterLobbyResp) ProtoMessage() {}

func (x *EnterLobbyResp) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterLobbyResp.ProtoReflect.Descriptor instead.
func (*EnterLobbyResp) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{2}
}

func (x *EnterLobbyResp) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //用户 id
	Mode   int32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`                   //模式
}

func (x *CreateRoomReq) Reset() {
	*x = CreateRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomReq) ProtoMessage() {}

func (x *CreateRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomReq.ProtoReflect.Descriptor instead.
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoomReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRoomReq) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type JoinRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int32 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` //用户 id
	Mode   int32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`                   //模式
}

func (x *JoinRoomReq) Reset() {
	*x = JoinRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomReq) ProtoMessage() {}

func (x *JoinRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomReq.ProtoReflect.Descriptor instead.
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRoomReq) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *JoinRoomReq) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

//请求匹配
type EnterMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode common.RoomMode `protobuf:"varint,1,opt,name=mode,proto3,enum=protocol.common.RoomMode" json:"mode,omitempty"` //模式
}

func (x *EnterMatchReq) Reset() {
	*x = EnterMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMatchReq) ProtoMessage() {}

func (x *EnterMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMatchReq.ProtoReflect.Descriptor instead.
func (*EnterMatchReq) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{5}
}

func (x *EnterMatchReq) GetMode() common.RoomMode {
	if x != nil {
		return x.Mode
	}
	return common.RoomMode(0)
}

//等待匹配中
type MatchWait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode    common.RoomMode `protobuf:"varint,1,opt,name=mode,proto3,enum=protocol.common.RoomMode" json:"mode,omitempty"` //模式
	Total   int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`                             // 匹配人数
	PreTime int32           `protobuf:"varint,3,opt,name=pre_time,json=preTime,proto3" json:"pre_time,omitempty"`          //预计时间
}

func (x *MatchWait) Reset() {
	*x = MatchWait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchWait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchWait) ProtoMessage() {}

func (x *MatchWait) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchWait.ProtoReflect.Descriptor instead.
func (*MatchWait) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{6}
}

func (x *MatchWait) GetMode() common.RoomMode {
	if x != nil {
		return x.Mode
	}
	return common.RoomMode(0)
}

func (x *MatchWait) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MatchWait) GetPreTime() int32 {
	if x != nil {
		return x.PreTime
	}
	return 0
}

//取消匹配
type CancelMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelMatchReq) Reset() {
	*x = CancelMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchReq) ProtoMessage() {}

func (x *CancelMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMatchReq.ProtoReflect.Descriptor instead.
func (*CancelMatchReq) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{7}
}

//取消匹配返回 有可能取消失败
type CancelMatchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suc bool `protobuf:"varint,1,opt,name=suc,proto3" json:"suc,omitempty"` //取消是否成功
}

func (x *CancelMatchResp) Reset() {
	*x = CancelMatchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchResp) ProtoMessage() {}

func (x *CancelMatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMatchResp.ProtoReflect.Descriptor instead.
func (*CancelMatchResp) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{8}
}

func (x *CancelMatchResp) GetSuc() bool {
	if x != nil {
		return x.Suc
	}
	return false
}

type SelectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //用户信息
	HeroId int32 `protobuf:"varint,2,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"` //英雄信息
}

func (x *SelectInfo) Reset() {
	*x = SelectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectInfo) ProtoMessage() {}

func (x *SelectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectInfo.ProtoReflect.Descriptor instead.
func (*SelectInfo) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{9}
}

func (x *SelectInfo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SelectInfo) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

//匹配成功
type MatchSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode  common.RoomMode `protobuf:"varint,1,opt,name=mode,proto3,enum=protocol.common.RoomMode" json:"mode,omitempty"` //模式
	Heros []*common.Hero  `protobuf:"bytes,2,rep,name=heros,proto3" json:"heros,omitempty"`                              //英雄列表
	Infos []*SelectInfo   `protobuf:"bytes,3,rep,name=infos,proto3" json:"infos,omitempty"`                              //选择信息
}

func (x *MatchSuccess) Reset() {
	*x = MatchSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchSuccess) ProtoMessage() {}

func (x *MatchSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchSuccess.ProtoReflect.Descriptor instead.
func (*MatchSuccess) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{10}
}

func (x *MatchSuccess) GetMode() common.RoomMode {
	if x != nil {
		return x.Mode
	}
	return common.RoomMode(0)
}

func (x *MatchSuccess) GetHeros() []*common.Hero {
	if x != nil {
		return x.Heros
	}
	return nil
}

func (x *MatchSuccess) GetInfos() []*SelectInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

//匹配失败
type MatchFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *MatchFailed) Reset() {
	*x = MatchFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protocol_client_client_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchFailed) ProtoMessage() {}

func (x *MatchFailed) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_client_client_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchFailed.ProtoReflect.Descriptor instead.
func (*MatchFailed) Descriptor() ([]byte, []int) {
	return file_Protocol_client_client_proto_rawDescGZIP(), []int{11}
}

func (x *MatchFailed) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_Protocol_client_client_proto protoreflect.FileDescriptor

var file_Protocol_client_client_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x29, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3e, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x22, 0x6b, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x57, 0x61, 0x69,
	0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x22, 0x23, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x75, 0x63, 0x22, 0x3e, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x68, 0x65, 0x72, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x05,
	0x68, 0x65, 0x72, 0x6f, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a,
	0xb3, 0x02, 0x0a, 0x02, 0x43, 0x50, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x43, 0x6f, 0x64,
	0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x4b, 0x65, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x5f, 0x52, 0x65, 0x71, 0x10, 0x0a, 0x12,
	0x13, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x6f, 0x6f, 0x6d, 0x5f, 0x52, 0x65, 0x71, 0x10, 0x66, 0x12,
	0x14, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x6f, 0x6f, 0x6d, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x10, 0x67, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x5f, 0x52, 0x6f,
	0x6f, 0x6d, 0x5f, 0x52, 0x65, 0x71, 0x10, 0x68, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x52, 0x65, 0x71, 0x10, 0x69, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x57, 0x61, 0x69,
	0x74, 0x10, 0x6a, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x10, 0x6b, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x10, 0x6c, 0x12, 0x12, 0x0a, 0x0e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x53, 0x75, 0x63, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x10, 0x6d,
	0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x46, 0x61, 0x69, 0x6c, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x10, 0x6e, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Protocol_client_client_proto_rawDescOnce sync.Once
	file_Protocol_client_client_proto_rawDescData = file_Protocol_client_client_proto_rawDesc
)

func file_Protocol_client_client_proto_rawDescGZIP() []byte {
	file_Protocol_client_client_proto_rawDescOnce.Do(func() {
		file_Protocol_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_Protocol_client_client_proto_rawDescData)
	})
	return file_Protocol_client_client_proto_rawDescData
}

var file_Protocol_client_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Protocol_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_Protocol_client_client_proto_goTypes = []interface{}{
	(CP)(0),                 // 0: protocol.client.CP
	(*ErrorCode)(nil),       // 1: protocol.client.ErrorCode
	(*ValidKeyReq)(nil),     // 2: protocol.client.ValidKeyReq
	(*EnterLobbyResp)(nil),  // 3: protocol.client.EnterLobbyResp
	(*CreateRoomReq)(nil),   // 4: protocol.client.CreateRoomReq
	(*JoinRoomReq)(nil),     // 5: protocol.client.JoinRoomReq
	(*EnterMatchReq)(nil),   // 6: protocol.client.EnterMatchReq
	(*MatchWait)(nil),       // 7: protocol.client.MatchWait
	(*CancelMatchReq)(nil),  // 8: protocol.client.CancelMatchReq
	(*CancelMatchResp)(nil), // 9: protocol.client.CancelMatchResp
	(*SelectInfo)(nil),      // 10: protocol.client.SelectInfo
	(*MatchSuccess)(nil),    // 11: protocol.client.MatchSuccess
	(*MatchFailed)(nil),     // 12: protocol.client.MatchFailed
	(errcode.Error)(0),      // 13: protocol.errcode.Error
	(common.RoomMode)(0),    // 14: protocol.common.RoomMode
	(*common.Hero)(nil),     // 15: protocol.common.Hero
}
var file_Protocol_client_client_proto_depIdxs = []int32{
	13, // 0: protocol.client.ErrorCode.code:type_name -> protocol.errcode.Error
	14, // 1: protocol.client.EnterMatchReq.mode:type_name -> protocol.common.RoomMode
	14, // 2: protocol.client.MatchWait.mode:type_name -> protocol.common.RoomMode
	14, // 3: protocol.client.MatchSuccess.mode:type_name -> protocol.common.RoomMode
	15, // 4: protocol.client.MatchSuccess.heros:type_name -> protocol.common.Hero
	10, // 5: protocol.client.MatchSuccess.infos:type_name -> protocol.client.SelectInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_Protocol_client_client_proto_init() }
func file_Protocol_client_client_proto_init() {
	if File_Protocol_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Protocol_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorCode); i {
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
		file_Protocol_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidKeyReq); i {
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
		file_Protocol_client_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterLobbyResp); i {
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
		file_Protocol_client_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomReq); i {
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
		file_Protocol_client_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomReq); i {
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
		file_Protocol_client_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterMatchReq); i {
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
		file_Protocol_client_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchWait); i {
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
		file_Protocol_client_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMatchReq); i {
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
		file_Protocol_client_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMatchResp); i {
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
		file_Protocol_client_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectInfo); i {
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
		file_Protocol_client_client_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchSuccess); i {
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
		file_Protocol_client_client_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchFailed); i {
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
			RawDescriptor: file_Protocol_client_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Protocol_client_client_proto_goTypes,
		DependencyIndexes: file_Protocol_client_client_proto_depIdxs,
		EnumInfos:         file_Protocol_client_client_proto_enumTypes,
		MessageInfos:      file_Protocol_client_client_proto_msgTypes,
	}.Build()
	File_Protocol_client_client_proto = out.File
	file_Protocol_client_client_proto_rawDesc = nil
	file_Protocol_client_client_proto_goTypes = nil
	file_Protocol_client_client_proto_depIdxs = nil
}
