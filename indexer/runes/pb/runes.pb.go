// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.27.1
// source: indexer/runes/pb/runes.proto

package pb

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

type Uint128 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lo uint64 `protobuf:"varint,1,opt,name=lo,proto3" json:"lo,omitempty"`
	Hi uint64 `protobuf:"varint,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *Uint128) Reset() {
	*x = Uint128{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint128) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint128) ProtoMessage() {}

func (x *Uint128) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint128.ProtoReflect.Descriptor instead.
func (*Uint128) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{0}
}

func (x *Uint128) GetLo() uint64 {
	if x != nil {
		return x.Lo
	}
	return 0
}

func (x *Uint128) GetHi() uint64 {
	if x != nil {
		return x.Hi
	}
	return 0
}

type Uint8 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Uint8) Reset() {
	*x = Uint8{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint8) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint8) ProtoMessage() {}

func (x *Uint8) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint8.ProtoReflect.Descriptor instead.
func (*Uint8) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{1}
}

func (x *Uint8) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type RuneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Tx    uint32 `protobuf:"varint,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *RuneId) Reset() {
	*x = RuneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuneId) ProtoMessage() {}

func (x *RuneId) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuneId.ProtoReflect.Descriptor instead.
func (*RuneId) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{2}
}

func (x *RuneId) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *RuneId) GetTx() uint32 {
	if x != nil {
		return x.Tx
	}
	return 0
}

type Rune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Uint128 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Rune) Reset() {
	*x = Rune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rune) ProtoMessage() {}

func (x *Rune) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rune.ProtoReflect.Descriptor instead.
func (*Rune) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{3}
}

func (x *Rune) GetValue() *Uint128 {
	if x != nil {
		return x.Value
	}
	return nil
}

type InscriptionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InscriptionId) Reset() {
	*x = InscriptionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InscriptionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InscriptionId) ProtoMessage() {}

func (x *InscriptionId) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InscriptionId.ProtoReflect.Descriptor instead.
func (*InscriptionId) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{4}
}

func (x *InscriptionId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SpacedRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rune    *Rune  `protobuf:"bytes,1,opt,name=rune,proto3" json:"rune,omitempty"`
	Spacers uint32 `protobuf:"varint,2,opt,name=Spacers,proto3" json:"Spacers,omitempty"`
}

func (x *SpacedRune) Reset() {
	*x = SpacedRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacedRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacedRune) ProtoMessage() {}

func (x *SpacedRune) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacedRune.ProtoReflect.Descriptor instead.
func (*SpacedRune) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{5}
}

func (x *SpacedRune) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

func (x *SpacedRune) GetSpacers() uint32 {
	if x != nil {
		return x.Spacers
	}
	return 0
}

type Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Symbol) Reset() {
	*x = Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Symbol) ProtoMessage() {}

func (x *Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Symbol.ProtoReflect.Descriptor instead.
func (*Symbol) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{6}
}

func (x *Symbol) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Terms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      *Uint128 `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Cap         *Uint128 `protobuf:"bytes,2,opt,name=cap,proto3" json:"cap,omitempty"`
	StartHeight *Uint128 `protobuf:"bytes,3,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight   *Uint128 `protobuf:"bytes,4,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	StartOffset uint64   `protobuf:"varint,5,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	EndOffset   uint64   `protobuf:"varint,6,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
}

func (x *Terms) Reset() {
	*x = Terms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terms) ProtoMessage() {}

func (x *Terms) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terms.ProtoReflect.Descriptor instead.
func (*Terms) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{7}
}

func (x *Terms) GetAmount() *Uint128 {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Terms) GetCap() *Uint128 {
	if x != nil {
		return x.Cap
	}
	return nil
}

func (x *Terms) GetStartHeight() *Uint128 {
	if x != nil {
		return x.StartHeight
	}
	return nil
}

func (x *Terms) GetEndHeight() *Uint128 {
	if x != nil {
		return x.EndHeight
	}
	return nil
}

func (x *Terms) GetStartOffset() uint64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *Terms) GetEndOffset() uint64 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

type RuneEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuneId       *RuneId        `protobuf:"bytes,1,opt,name=runeId,proto3" json:"runeId,omitempty"`
	Burned       *Uint128       `protobuf:"bytes,2,opt,name=burned,proto3" json:"burned,omitempty"`
	Divisibility *Uint8         `protobuf:"bytes,3,opt,name=divisibility,proto3" json:"divisibility,omitempty"`
	Etching      []byte         `protobuf:"bytes,4,opt,name=etching,proto3" json:"etching,omitempty"`
	Parent       *InscriptionId `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	Mints        *Uint128       `protobuf:"bytes,6,opt,name=mints,proto3" json:"mints,omitempty"`
	Number       uint64         `protobuf:"varint,7,opt,name=number,proto3" json:"number,omitempty"`
	Premine      *Uint128       `protobuf:"bytes,8,opt,name=premine,proto3" json:"premine,omitempty"`
	SpacedRune   *SpacedRune    `protobuf:"bytes,9,opt,name=spaced_rune,json=spacedRune,proto3" json:"spaced_rune,omitempty"`
	Symbol       *Symbol        `protobuf:"bytes,10,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Terms        *Terms         `protobuf:"bytes,11,opt,name=terms,proto3" json:"terms,omitempty"`
	Timestamp    uint64         `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Turbo        bool           `protobuf:"varint,13,opt,name=turbo,proto3" json:"turbo,omitempty"`
}

func (x *RuneEntry) Reset() {
	*x = RuneEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_runes_pb_runes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuneEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuneEntry) ProtoMessage() {}

func (x *RuneEntry) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_runes_pb_runes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuneEntry.ProtoReflect.Descriptor instead.
func (*RuneEntry) Descriptor() ([]byte, []int) {
	return file_indexer_runes_pb_runes_proto_rawDescGZIP(), []int{8}
}

func (x *RuneEntry) GetRuneId() *RuneId {
	if x != nil {
		return x.RuneId
	}
	return nil
}

func (x *RuneEntry) GetBurned() *Uint128 {
	if x != nil {
		return x.Burned
	}
	return nil
}

func (x *RuneEntry) GetDivisibility() *Uint8 {
	if x != nil {
		return x.Divisibility
	}
	return nil
}

func (x *RuneEntry) GetEtching() []byte {
	if x != nil {
		return x.Etching
	}
	return nil
}

func (x *RuneEntry) GetParent() *InscriptionId {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *RuneEntry) GetMints() *Uint128 {
	if x != nil {
		return x.Mints
	}
	return nil
}

func (x *RuneEntry) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *RuneEntry) GetPremine() *Uint128 {
	if x != nil {
		return x.Premine
	}
	return nil
}

func (x *RuneEntry) GetSpacedRune() *SpacedRune {
	if x != nil {
		return x.SpacedRune
	}
	return nil
}

func (x *RuneEntry) GetSymbol() *Symbol {
	if x != nil {
		return x.Symbol
	}
	return nil
}

func (x *RuneEntry) GetTerms() *Terms {
	if x != nil {
		return x.Terms
	}
	return nil
}

func (x *RuneEntry) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RuneEntry) GetTurbo() bool {
	if x != nil {
		return x.Turbo
	}
	return false
}

var File_indexer_runes_pb_runes_proto protoreflect.FileDescriptor

var file_indexer_runes_pb_runes_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x07, 0x55, 0x69, 0x6e, 0x74,
	0x31, 0x32, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x6c, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x68, 0x69, 0x22, 0x1d, 0x0a, 0x05, 0x55, 0x69, 0x6e, 0x74, 0x38, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x74, 0x78, 0x22, 0x2f, 0x0a, 0x04, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72,
	0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x32, 0x38, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x0a, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65,
	0x73, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x72, 0x73, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x81, 0x02, 0x0a, 0x05, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x12, 0x29, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x31, 0x32, 0x38, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x63,
	0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75,
	0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x32, 0x38, 0x52, 0x03, 0x63, 0x61, 0x70,
	0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65,
	0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x32, 0x38, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x32, 0x38, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8a, 0x04, 0x0a, 0x09, 0x52,
	0x75, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75,
	0x6e, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x52, 0x06, 0x72, 0x75, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x62, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69,
	0x6e, 0x74, 0x31, 0x32, 0x38, 0x52, 0x06, 0x62, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x33, 0x0a,
	0x0c, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x38, 0x52, 0x0c, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x6d, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31, 0x32, 0x38, 0x52,
	0x05, 0x6d, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x31,
	0x32, 0x38, 0x52, 0x07, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x64, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x52, 0x75,
	0x6e, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x05,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x05, 0x74, 0x65,
	0x72, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x72, 0x75, 0x6e, 0x65,
	0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indexer_runes_pb_runes_proto_rawDescOnce sync.Once
	file_indexer_runes_pb_runes_proto_rawDescData = file_indexer_runes_pb_runes_proto_rawDesc
)

func file_indexer_runes_pb_runes_proto_rawDescGZIP() []byte {
	file_indexer_runes_pb_runes_proto_rawDescOnce.Do(func() {
		file_indexer_runes_pb_runes_proto_rawDescData = protoimpl.X.CompressGZIP(file_indexer_runes_pb_runes_proto_rawDescData)
	})
	return file_indexer_runes_pb_runes_proto_rawDescData
}

var file_indexer_runes_pb_runes_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_indexer_runes_pb_runes_proto_goTypes = []interface{}{
	(*Uint128)(nil),       // 0: pb.runes.Uint128
	(*Uint8)(nil),         // 1: pb.runes.Uint8
	(*RuneId)(nil),        // 2: pb.runes.RuneId
	(*Rune)(nil),          // 3: pb.runes.Rune
	(*InscriptionId)(nil), // 4: pb.runes.InscriptionId
	(*SpacedRune)(nil),    // 5: pb.runes.SpacedRune
	(*Symbol)(nil),        // 6: pb.runes.Symbol
	(*Terms)(nil),         // 7: pb.runes.Terms
	(*RuneEntry)(nil),     // 8: pb.runes.RuneEntry
}
var file_indexer_runes_pb_runes_proto_depIdxs = []int32{
	0,  // 0: pb.runes.Rune.value:type_name -> pb.runes.Uint128
	3,  // 1: pb.runes.SpacedRune.rune:type_name -> pb.runes.Rune
	0,  // 2: pb.runes.Terms.amount:type_name -> pb.runes.Uint128
	0,  // 3: pb.runes.Terms.cap:type_name -> pb.runes.Uint128
	0,  // 4: pb.runes.Terms.start_height:type_name -> pb.runes.Uint128
	0,  // 5: pb.runes.Terms.end_height:type_name -> pb.runes.Uint128
	2,  // 6: pb.runes.RuneEntry.runeId:type_name -> pb.runes.RuneId
	0,  // 7: pb.runes.RuneEntry.burned:type_name -> pb.runes.Uint128
	1,  // 8: pb.runes.RuneEntry.divisibility:type_name -> pb.runes.Uint8
	4,  // 9: pb.runes.RuneEntry.parent:type_name -> pb.runes.InscriptionId
	0,  // 10: pb.runes.RuneEntry.mints:type_name -> pb.runes.Uint128
	0,  // 11: pb.runes.RuneEntry.premine:type_name -> pb.runes.Uint128
	5,  // 12: pb.runes.RuneEntry.spaced_rune:type_name -> pb.runes.SpacedRune
	6,  // 13: pb.runes.RuneEntry.symbol:type_name -> pb.runes.Symbol
	7,  // 14: pb.runes.RuneEntry.terms:type_name -> pb.runes.Terms
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_indexer_runes_pb_runes_proto_init() }
func file_indexer_runes_pb_runes_proto_init() {
	if File_indexer_runes_pb_runes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indexer_runes_pb_runes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint128); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint8); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuneId); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rune); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InscriptionId); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacedRune); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Symbol); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terms); i {
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
		file_indexer_runes_pb_runes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuneEntry); i {
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
			RawDescriptor: file_indexer_runes_pb_runes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indexer_runes_pb_runes_proto_goTypes,
		DependencyIndexes: file_indexer_runes_pb_runes_proto_depIdxs,
		MessageInfos:      file_indexer_runes_pb_runes_proto_msgTypes,
	}.Build()
	File_indexer_runes_pb_runes_proto = out.File
	file_indexer_runes_pb_runes_proto_rawDesc = nil
	file_indexer_runes_pb_runes_proto_goTypes = nil
	file_indexer_runes_pb_runes_proto_depIdxs = nil
}
