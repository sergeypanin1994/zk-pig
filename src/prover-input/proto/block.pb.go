// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: src/prover-input/proto/block.proto

package proto

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

type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *Header                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions  []*Transaction         `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Uncles        []*Header              `protobuf:"bytes,3,rep,name=uncles,proto3" json:"uncles,omitempty"`
	Withdrawals   []*Withdrawal          `protobuf:"bytes,4,rep,name=withdrawals,proto3" json:"withdrawals,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_src_prover_input_proto_block_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_src_prover_input_proto_block_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_src_prover_input_proto_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetUncles() []*Header {
	if x != nil {
		return x.Uncles
	}
	return nil
}

func (x *Block) GetWithdrawals() []*Withdrawal {
	if x != nil {
		return x.Withdrawals
	}
	return nil
}

type Header struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ParentHash       []byte                 `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Sha3Uncles       []byte                 `protobuf:"bytes,2,opt,name=sha3_uncles,json=sha3Uncles,proto3" json:"sha3_uncles,omitempty"`
	Miner            []byte                 `protobuf:"bytes,3,opt,name=miner,proto3" json:"miner,omitempty"`
	Root             []byte                 `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	TransactionsRoot []byte                 `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	ReceiptsRoot     []byte                 `protobuf:"bytes,6,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`
	LogsBloom        []byte                 `protobuf:"bytes,7,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom,omitempty"`
	Difficulty       []byte                 `protobuf:"bytes,8,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Number           []byte                 `protobuf:"bytes,9,opt,name=number,proto3" json:"number,omitempty"`
	GasLimit         uint64                 `protobuf:"varint,10,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed          uint64                 `protobuf:"varint,11,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Timestamp        uint64                 `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExtraData        []byte                 `protobuf:"bytes,13,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	MixHash          []byte                 `protobuf:"bytes,14,opt,name=mix_hash,json=mixHash,proto3" json:"mix_hash,omitempty"`
	Nonce            uint64                 `protobuf:"varint,15,opt,name=nonce,proto3" json:"nonce,omitempty"`
	BaseFeePerGas    []byte                 `protobuf:"bytes,16,opt,name=base_fee_per_gas,json=baseFeePerGas,proto3,oneof" json:"base_fee_per_gas,omitempty"`
	WithdrawalsRoot  []byte                 `protobuf:"bytes,17,opt,name=withdrawals_root,json=withdrawalsRoot,proto3,oneof" json:"withdrawals_root,omitempty"`
	BlobGasUsed      *uint64                `protobuf:"varint,18,opt,name=blob_gas_used,json=blobGasUsed,proto3,oneof" json:"blob_gas_used,omitempty"`
	ExcessBlobGas    *uint64                `protobuf:"varint,19,opt,name=excess_blob_gas,json=excessBlobGas,proto3,oneof" json:"excess_blob_gas,omitempty"`
	ParentBeaconRoot []byte                 `protobuf:"bytes,20,opt,name=parent_beacon_root,json=parentBeaconRoot,proto3,oneof" json:"parent_beacon_root,omitempty"`
	RequestsRoot     []byte                 `protobuf:"bytes,21,opt,name=requests_root,json=requestsRoot,proto3,oneof" json:"requests_root,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_src_prover_input_proto_block_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_src_prover_input_proto_block_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_src_prover_input_proto_block_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *Header) GetSha3Uncles() []byte {
	if x != nil {
		return x.Sha3Uncles
	}
	return nil
}

func (x *Header) GetMiner() []byte {
	if x != nil {
		return x.Miner
	}
	return nil
}

func (x *Header) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *Header) GetTransactionsRoot() []byte {
	if x != nil {
		return x.TransactionsRoot
	}
	return nil
}

func (x *Header) GetReceiptsRoot() []byte {
	if x != nil {
		return x.ReceiptsRoot
	}
	return nil
}

func (x *Header) GetLogsBloom() []byte {
	if x != nil {
		return x.LogsBloom
	}
	return nil
}

func (x *Header) GetDifficulty() []byte {
	if x != nil {
		return x.Difficulty
	}
	return nil
}

func (x *Header) GetNumber() []byte {
	if x != nil {
		return x.Number
	}
	return nil
}

func (x *Header) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Header) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *Header) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetExtraData() []byte {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

func (x *Header) GetMixHash() []byte {
	if x != nil {
		return x.MixHash
	}
	return nil
}

func (x *Header) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Header) GetBaseFeePerGas() []byte {
	if x != nil {
		return x.BaseFeePerGas
	}
	return nil
}

func (x *Header) GetWithdrawalsRoot() []byte {
	if x != nil {
		return x.WithdrawalsRoot
	}
	return nil
}

func (x *Header) GetBlobGasUsed() uint64 {
	if x != nil && x.BlobGasUsed != nil {
		return *x.BlobGasUsed
	}
	return 0
}

func (x *Header) GetExcessBlobGas() uint64 {
	if x != nil && x.ExcessBlobGas != nil {
		return *x.ExcessBlobGas
	}
	return 0
}

func (x *Header) GetParentBeaconRoot() []byte {
	if x != nil {
		return x.ParentBeaconRoot
	}
	return nil
}

func (x *Header) GetRequestsRoot() []byte {
	if x != nil {
		return x.RequestsRoot
	}
	return nil
}

type Withdrawal struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Index          uint64                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	ValidatorIndex uint64                 `protobuf:"varint,2,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Address        []byte                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Amount         uint64                 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Withdrawal) Reset() {
	*x = Withdrawal{}
	mi := &file_src_prover_input_proto_block_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Withdrawal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdrawal) ProtoMessage() {}

func (x *Withdrawal) ProtoReflect() protoreflect.Message {
	mi := &file_src_prover_input_proto_block_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdrawal.ProtoReflect.Descriptor instead.
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return file_src_prover_input_proto_block_proto_rawDescGZIP(), []int{2}
}

func (x *Withdrawal) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Withdrawal) GetValidatorIndex() uint64 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *Withdrawal) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Withdrawal) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_src_prover_input_proto_block_proto protoreflect.FileDescriptor

var file_src_prover_input_proto_block_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x25, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x0a, 0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x75,
	0x6e, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x0b, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x22, 0xcd, 0x06, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x33, 0x5f, 0x75,
	0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x68, 0x61,
	0x33, 0x55, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x6f,
	0x74, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6f,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f,
	0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61,
	0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75,
	0x73, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x69, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x69, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x67, 0x61, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x61,
	0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e,
	0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x47, 0x61, 0x73,
	0x55, 0x73, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x65, 0x73, 0x73, 0x42, 0x6c, 0x6f, 0x62, 0x47, 0x61,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x04, 0x52, 0x10, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x52, 0x6f, 0x6f, 0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x05,
	0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x67, 0x61,
	0x73, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x22, 0x7d, 0x0a, 0x0a, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27,
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6b, 0x72, 0x74, 0x2d, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x7a, 0x6b, 0x2d, 0x70, 0x69, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_prover_input_proto_block_proto_rawDescOnce sync.Once
	file_src_prover_input_proto_block_proto_rawDescData = file_src_prover_input_proto_block_proto_rawDesc
)

func file_src_prover_input_proto_block_proto_rawDescGZIP() []byte {
	file_src_prover_input_proto_block_proto_rawDescOnce.Do(func() {
		file_src_prover_input_proto_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_prover_input_proto_block_proto_rawDescData)
	})
	return file_src_prover_input_proto_block_proto_rawDescData
}

var file_src_prover_input_proto_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_prover_input_proto_block_proto_goTypes = []any{
	(*Block)(nil),       // 0: input.Block
	(*Header)(nil),      // 1: input.Header
	(*Withdrawal)(nil),  // 2: input.Withdrawal
	(*Transaction)(nil), // 3: input.Transaction
}
var file_src_prover_input_proto_block_proto_depIdxs = []int32{
	1, // 0: input.Block.header:type_name -> input.Header
	3, // 1: input.Block.transactions:type_name -> input.Transaction
	1, // 2: input.Block.uncles:type_name -> input.Header
	2, // 3: input.Block.withdrawals:type_name -> input.Withdrawal
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_src_prover_input_proto_block_proto_init() }
func file_src_prover_input_proto_block_proto_init() {
	if File_src_prover_input_proto_block_proto != nil {
		return
	}
	file_src_prover_input_proto_transaction_proto_init()
	file_src_prover_input_proto_block_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_prover_input_proto_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_prover_input_proto_block_proto_goTypes,
		DependencyIndexes: file_src_prover_input_proto_block_proto_depIdxs,
		MessageInfos:      file_src_prover_input_proto_block_proto_msgTypes,
	}.Build()
	File_src_prover_input_proto_block_proto = out.File
	file_src_prover_input_proto_block_proto_rawDesc = nil
	file_src_prover_input_proto_block_proto_goTypes = nil
	file_src_prover_input_proto_block_proto_depIdxs = nil
}
