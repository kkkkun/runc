// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: tty.proto

package images

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

type TtyType int32

const (
	TtyType_UNKNOWN TtyType = 0
	TtyType_PTY     TtyType = 1
	TtyType_CONSOLE TtyType = 2
	TtyType_VT      TtyType = 3
	TtyType_CTTY    TtyType = 4
	TtyType_EXT_TTY TtyType = 5
	TtyType_SERIAL  TtyType = 6
)

// Enum value maps for TtyType.
var (
	TtyType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PTY",
		2: "CONSOLE",
		3: "VT",
		4: "CTTY",
		5: "EXT_TTY",
		6: "SERIAL",
	}
	TtyType_value = map[string]int32{
		"UNKNOWN": 0,
		"PTY":     1,
		"CONSOLE": 2,
		"VT":      3,
		"CTTY":    4,
		"EXT_TTY": 5,
		"SERIAL":  6,
	}
)

func (x TtyType) Enum() *TtyType {
	p := new(TtyType)
	*p = x
	return p
}

func (x TtyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TtyType) Descriptor() protoreflect.EnumDescriptor {
	return file_tty_proto_enumTypes[0].Descriptor()
}

func (TtyType) Type() protoreflect.EnumType {
	return &file_tty_proto_enumTypes[0]
}

func (x TtyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TtyType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TtyType(num)
	return nil
}

// Deprecated: Use TtyType.Descriptor instead.
func (TtyType) EnumDescriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{0}
}

type WinsizeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WsRow    *uint32 `protobuf:"varint,1,req,name=ws_row,json=wsRow" json:"ws_row,omitempty"`
	WsCol    *uint32 `protobuf:"varint,2,req,name=ws_col,json=wsCol" json:"ws_col,omitempty"`
	WsXpixel *uint32 `protobuf:"varint,3,req,name=ws_xpixel,json=wsXpixel" json:"ws_xpixel,omitempty"`
	WsYpixel *uint32 `protobuf:"varint,4,req,name=ws_ypixel,json=wsYpixel" json:"ws_ypixel,omitempty"`
}

func (x *WinsizeEntry) Reset() {
	*x = WinsizeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinsizeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinsizeEntry) ProtoMessage() {}

func (x *WinsizeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinsizeEntry.ProtoReflect.Descriptor instead.
func (*WinsizeEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{0}
}

func (x *WinsizeEntry) GetWsRow() uint32 {
	if x != nil && x.WsRow != nil {
		return *x.WsRow
	}
	return 0
}

func (x *WinsizeEntry) GetWsCol() uint32 {
	if x != nil && x.WsCol != nil {
		return *x.WsCol
	}
	return 0
}

func (x *WinsizeEntry) GetWsXpixel() uint32 {
	if x != nil && x.WsXpixel != nil {
		return *x.WsXpixel
	}
	return 0
}

func (x *WinsizeEntry) GetWsYpixel() uint32 {
	if x != nil && x.WsYpixel != nil {
		return *x.WsYpixel
	}
	return 0
}

type TermiosEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CIflag  *uint32  `protobuf:"varint,1,req,name=c_iflag,json=cIflag" json:"c_iflag,omitempty"`
	COflag  *uint32  `protobuf:"varint,2,req,name=c_oflag,json=cOflag" json:"c_oflag,omitempty"`
	CCflag  *uint32  `protobuf:"varint,3,req,name=c_cflag,json=cCflag" json:"c_cflag,omitempty"`
	CLflag  *uint32  `protobuf:"varint,4,req,name=c_lflag,json=cLflag" json:"c_lflag,omitempty"`
	CLine   *uint32  `protobuf:"varint,5,req,name=c_line,json=cLine" json:"c_line,omitempty"`
	CIspeed *uint32  `protobuf:"varint,6,req,name=c_ispeed,json=cIspeed" json:"c_ispeed,omitempty"`
	COspeed *uint32  `protobuf:"varint,7,req,name=c_ospeed,json=cOspeed" json:"c_ospeed,omitempty"`
	CCc     []uint32 `protobuf:"varint,8,rep,name=c_cc,json=cCc" json:"c_cc,omitempty"`
}

func (x *TermiosEntry) Reset() {
	*x = TermiosEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermiosEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermiosEntry) ProtoMessage() {}

func (x *TermiosEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermiosEntry.ProtoReflect.Descriptor instead.
func (*TermiosEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{1}
}

func (x *TermiosEntry) GetCIflag() uint32 {
	if x != nil && x.CIflag != nil {
		return *x.CIflag
	}
	return 0
}

func (x *TermiosEntry) GetCOflag() uint32 {
	if x != nil && x.COflag != nil {
		return *x.COflag
	}
	return 0
}

func (x *TermiosEntry) GetCCflag() uint32 {
	if x != nil && x.CCflag != nil {
		return *x.CCflag
	}
	return 0
}

func (x *TermiosEntry) GetCLflag() uint32 {
	if x != nil && x.CLflag != nil {
		return *x.CLflag
	}
	return 0
}

func (x *TermiosEntry) GetCLine() uint32 {
	if x != nil && x.CLine != nil {
		return *x.CLine
	}
	return 0
}

func (x *TermiosEntry) GetCIspeed() uint32 {
	if x != nil && x.CIspeed != nil {
		return *x.CIspeed
	}
	return 0
}

func (x *TermiosEntry) GetCOspeed() uint32 {
	if x != nil && x.COspeed != nil {
		return *x.COspeed
	}
	return 0
}

func (x *TermiosEntry) GetCCc() []uint32 {
	if x != nil {
		return x.CCc
	}
	return nil
}

type TtyPtyEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *uint32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
}

func (x *TtyPtyEntry) Reset() {
	*x = TtyPtyEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtyPtyEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtyPtyEntry) ProtoMessage() {}

func (x *TtyPtyEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtyPtyEntry.ProtoReflect.Descriptor instead.
func (*TtyPtyEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{2}
}

func (x *TtyPtyEntry) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type TtyDataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TtyId *uint32 `protobuf:"varint,1,req,name=tty_id,json=ttyId" json:"tty_id,omitempty"`
	Data  []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
}

func (x *TtyDataEntry) Reset() {
	*x = TtyDataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtyDataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtyDataEntry) ProtoMessage() {}

func (x *TtyDataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtyDataEntry.ProtoReflect.Descriptor instead.
func (*TtyDataEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{3}
}

func (x *TtyDataEntry) GetTtyId() uint32 {
	if x != nil && x.TtyId != nil {
		return *x.TtyId
	}
	return 0
}

func (x *TtyDataEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TtyInfoEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type       *TtyType `protobuf:"varint,2,req,name=type,enum=criu.TtyType" json:"type,omitempty"`
	Locked     *bool    `protobuf:"varint,3,req,name=locked" json:"locked,omitempty"` // Unix98 PTY only
	Exclusive  *bool    `protobuf:"varint,4,req,name=exclusive" json:"exclusive,omitempty"`
	PacketMode *bool    `protobuf:"varint,5,req,name=packet_mode,json=packetMode" json:"packet_mode,omitempty"` // Unix98 PTY only
	Sid        *uint32  `protobuf:"varint,6,req,name=sid" json:"sid,omitempty"`
	Pgrp       *uint32  `protobuf:"varint,7,req,name=pgrp" json:"pgrp,omitempty"`
	// Convenient for printing errors and such, with this
	// device encoded we can figure out major and minor
	// numbers.
	Rdev          *uint32       `protobuf:"varint,8,req,name=rdev" json:"rdev,omitempty"`
	Termios       *TermiosEntry `protobuf:"bytes,9,opt,name=termios" json:"termios,omitempty"`
	TermiosLocked *TermiosEntry `protobuf:"bytes,10,opt,name=termios_locked,json=termiosLocked" json:"termios_locked,omitempty"`
	Winsize       *WinsizeEntry `protobuf:"bytes,11,opt,name=winsize" json:"winsize,omitempty"`
	// These are optional fields which presence depends on
	// TTY type.
	Pty *TtyPtyEntry `protobuf:"bytes,12,opt,name=pty" json:"pty,omitempty"`
	Dev *uint32      `protobuf:"varint,13,opt,name=dev" json:"dev,omitempty"`
	Uid *uint32      `protobuf:"varint,14,opt,name=uid" json:"uid,omitempty"`
	Gid *uint32      `protobuf:"varint,15,opt,name=gid" json:"gid,omitempty"`
}

func (x *TtyInfoEntry) Reset() {
	*x = TtyInfoEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtyInfoEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtyInfoEntry) ProtoMessage() {}

func (x *TtyInfoEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtyInfoEntry.ProtoReflect.Descriptor instead.
func (*TtyInfoEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{4}
}

func (x *TtyInfoEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *TtyInfoEntry) GetType() TtyType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TtyType_UNKNOWN
}

func (x *TtyInfoEntry) GetLocked() bool {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return false
}

func (x *TtyInfoEntry) GetExclusive() bool {
	if x != nil && x.Exclusive != nil {
		return *x.Exclusive
	}
	return false
}

func (x *TtyInfoEntry) GetPacketMode() bool {
	if x != nil && x.PacketMode != nil {
		return *x.PacketMode
	}
	return false
}

func (x *TtyInfoEntry) GetSid() uint32 {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return 0
}

func (x *TtyInfoEntry) GetPgrp() uint32 {
	if x != nil && x.Pgrp != nil {
		return *x.Pgrp
	}
	return 0
}

func (x *TtyInfoEntry) GetRdev() uint32 {
	if x != nil && x.Rdev != nil {
		return *x.Rdev
	}
	return 0
}

func (x *TtyInfoEntry) GetTermios() *TermiosEntry {
	if x != nil {
		return x.Termios
	}
	return nil
}

func (x *TtyInfoEntry) GetTermiosLocked() *TermiosEntry {
	if x != nil {
		return x.TermiosLocked
	}
	return nil
}

func (x *TtyInfoEntry) GetWinsize() *WinsizeEntry {
	if x != nil {
		return x.Winsize
	}
	return nil
}

func (x *TtyInfoEntry) GetPty() *TtyPtyEntry {
	if x != nil {
		return x.Pty
	}
	return nil
}

func (x *TtyInfoEntry) GetDev() uint32 {
	if x != nil && x.Dev != nil {
		return *x.Dev
	}
	return 0
}

func (x *TtyInfoEntry) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *TtyInfoEntry) GetGid() uint32 {
	if x != nil && x.Gid != nil {
		return *x.Gid
	}
	return 0
}

type TtyFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint32    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	TtyInfoId *uint32    `protobuf:"varint,2,req,name=tty_info_id,json=ttyInfoId" json:"tty_info_id,omitempty"`
	Flags     *uint32    `protobuf:"varint,3,req,name=flags" json:"flags,omitempty"`
	Fown      *FownEntry `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
	// optional sint32		mnt_id		= 5 [default = 0];
	RegfId *uint32 `protobuf:"varint,6,opt,name=regf_id,json=regfId" json:"regf_id,omitempty"`
}

func (x *TtyFileEntry) Reset() {
	*x = TtyFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tty_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtyFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtyFileEntry) ProtoMessage() {}

func (x *TtyFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tty_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtyFileEntry.ProtoReflect.Descriptor instead.
func (*TtyFileEntry) Descriptor() ([]byte, []int) {
	return file_tty_proto_rawDescGZIP(), []int{5}
}

func (x *TtyFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *TtyFileEntry) GetTtyInfoId() uint32 {
	if x != nil && x.TtyInfoId != nil {
		return *x.TtyInfoId
	}
	return 0
}

func (x *TtyFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *TtyFileEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *TtyFileEntry) GetRegfId() uint32 {
	if x != nil && x.RegfId != nil {
		return *x.RegfId
	}
	return 0
}

var File_tty_proto protoreflect.FileDescriptor

var file_tty_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x72, 0x69,
	0x75, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66,
	0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0d, 0x77, 0x69, 0x6e,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73,
	0x5f, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x73, 0x52, 0x6f,
	0x77, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x77, 0x73, 0x43, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x73, 0x5f, 0x78,
	0x70, 0x69, 0x78, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x73, 0x58,
	0x70, 0x69, 0x78, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x73, 0x5f, 0x79, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x73, 0x59, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x22, 0xd3, 0x01, 0x0a, 0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6f, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x5f, 0x69, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x49, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x5f, 0x6f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06,
	0x63, 0x4f, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x5f, 0x63, 0x66, 0x6c, 0x61,
	0x67, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x43, 0x66, 0x6c, 0x61, 0x67, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x5f, 0x6c, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x06, 0x63, 0x4c, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x5f, 0x69, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x07, 0x63, 0x49, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x5f,
	0x6f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x4f,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x11, 0x0a, 0x04, 0x63, 0x5f, 0x63, 0x63, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x43, 0x63, 0x22, 0x25, 0x0a, 0x0d, 0x74, 0x74, 0x79, 0x5f,
	0x70, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x3b, 0x0a, 0x0e, 0x74, 0x74, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x74, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x03, 0x0a,
	0x0e, 0x74, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x63, 0x72, 0x69, 0x75, 0x2e, 0x54, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x09, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x67, 0x72, 0x70, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x67, 0x72, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x64, 0x65, 0x76, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x64, 0x65, 0x76, 0x12, 0x2d, 0x0a, 0x07, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6f, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x69, 0x75, 0x2e, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6f, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6f, 0x73, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x69,
	0x75, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6f, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6f, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x2d,
	0x0a, 0x07, 0x77, 0x69, 0x6e, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x72, 0x69, 0x75, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a,
	0x03, 0x70, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x69,
	0x75, 0x2e, 0x74, 0x74, 0x79, 0x5f, 0x70, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x70, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x64, 0x65, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x67, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x74,
	0x74, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x74, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x09, 0x74, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x05, 0xd2, 0x3f,
	0x02, 0x08, 0x01, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x66, 0x6f,
	0x77, 0x6e, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x69, 0x75, 0x2e,
	0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x65, 0x67, 0x66, 0x49, 0x64, 0x2a, 0x57, 0x0a, 0x07, 0x54, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x54, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x54, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x54, 0x54, 0x59, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x54,
	0x5f, 0x54, 0x54, 0x59, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c,
	0x10, 0x06,
}

var (
	file_tty_proto_rawDescOnce sync.Once
	file_tty_proto_rawDescData = file_tty_proto_rawDesc
)

func file_tty_proto_rawDescGZIP() []byte {
	file_tty_proto_rawDescOnce.Do(func() {
		file_tty_proto_rawDescData = protoimpl.X.CompressGZIP(file_tty_proto_rawDescData)
	})
	return file_tty_proto_rawDescData
}

var file_tty_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tty_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tty_proto_goTypes = []interface{}{
	(TtyType)(0),         // 0: criu.TtyType
	(*WinsizeEntry)(nil), // 1: criu.winsize_entry
	(*TermiosEntry)(nil), // 2: criu.termios_entry
	(*TtyPtyEntry)(nil),  // 3: criu.tty_pty_entry
	(*TtyDataEntry)(nil), // 4: criu.tty_data_entry
	(*TtyInfoEntry)(nil), // 5: criu.tty_info_entry
	(*TtyFileEntry)(nil), // 6: criu.tty_file_entry
	(*FownEntry)(nil),    // 7: criu.fown_entry
}
var file_tty_proto_depIdxs = []int32{
	0, // 0: criu.tty_info_entry.type:type_name -> criu.TtyType
	2, // 1: criu.tty_info_entry.termios:type_name -> criu.termios_entry
	2, // 2: criu.tty_info_entry.termios_locked:type_name -> criu.termios_entry
	1, // 3: criu.tty_info_entry.winsize:type_name -> criu.winsize_entry
	3, // 4: criu.tty_info_entry.pty:type_name -> criu.tty_pty_entry
	7, // 5: criu.tty_file_entry.fown:type_name -> criu.fown_entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tty_proto_init() }
func file_tty_proto_init() {
	if File_tty_proto != nil {
		return
	}
	file_opts_proto_init()
	file_fown_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinsizeEntry); i {
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
		file_tty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermiosEntry); i {
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
		file_tty_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtyPtyEntry); i {
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
		file_tty_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtyDataEntry); i {
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
		file_tty_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtyInfoEntry); i {
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
		file_tty_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtyFileEntry); i {
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
			RawDescriptor: file_tty_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tty_proto_goTypes,
		DependencyIndexes: file_tty_proto_depIdxs,
		EnumInfos:         file_tty_proto_enumTypes,
		MessageInfos:      file_tty_proto_msgTypes,
	}.Build()
	File_tty_proto = out.File
	file_tty_proto_rawDesc = nil
	file_tty_proto_goTypes = nil
	file_tty_proto_depIdxs = nil
}