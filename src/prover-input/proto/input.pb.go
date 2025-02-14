// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: src/prover-input/proto/input.proto

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

// ProverInput contains the minimal data needed for block execution and proof validation
type ProverInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Blocks        []*Block               `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Witness       *Witness               `protobuf:"bytes,3,opt,name=witness,proto3" json:"witness,omitempty"`
	ChainConfig   *ChainConfig           `protobuf:"bytes,4,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProverInput) Reset() {
	*x = ProverInput{}
	mi := &file_src_prover_input_proto_input_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProverInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProverInput) ProtoMessage() {}

func (x *ProverInput) ProtoReflect() protoreflect.Message {
	mi := &file_src_prover_input_proto_input_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProverInput.ProtoReflect.Descriptor instead.
func (*ProverInput) Descriptor() ([]byte, []int) {
	return file_src_prover_input_proto_input_proto_rawDescGZIP(), []int{0}
}

func (x *ProverInput) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProverInput) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *ProverInput) GetWitness() *Witness {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *ProverInput) GetChainConfig() *ChainConfig {
	if x != nil {
		return x.ChainConfig
	}
	return nil
}

type Witness struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         [][]byte               `protobuf:"bytes,1,rep,name=state,proto3" json:"state,omitempty"`
	Ancestors     []*Header              `protobuf:"bytes,2,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
	Codes         [][]byte               `protobuf:"bytes,3,rep,name=codes,proto3" json:"codes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Witness) Reset() {
	*x = Witness{}
	mi := &file_src_prover_input_proto_input_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Witness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Witness) ProtoMessage() {}

func (x *Witness) ProtoReflect() protoreflect.Message {
	mi := &file_src_prover_input_proto_input_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Witness.ProtoReflect.Descriptor instead.
func (*Witness) Descriptor() ([]byte, []int) {
	return file_src_prover_input_proto_input_proto_rawDescGZIP(), []int{1}
}

func (x *Witness) GetState() [][]byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Witness) GetAncestors() []*Header {
	if x != nil {
		return x.Ancestors
	}
	return nil
}

func (x *Witness) GetCodes() [][]byte {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_src_prover_input_proto_input_proto protoreflect.FileDescriptor

var file_src_prover_input_proto_input_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x77, 0x69,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x07, 0x77, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x62, 0x0a, 0x07, 0x57,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x09,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6b,
	0x72, 0x74, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x7a, 0x6b, 0x2d, 0x70, 0x69, 0x67, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_prover_input_proto_input_proto_rawDescOnce sync.Once
	file_src_prover_input_proto_input_proto_rawDescData = file_src_prover_input_proto_input_proto_rawDesc
)

func file_src_prover_input_proto_input_proto_rawDescGZIP() []byte {
	file_src_prover_input_proto_input_proto_rawDescOnce.Do(func() {
		file_src_prover_input_proto_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_prover_input_proto_input_proto_rawDescData)
	})
	return file_src_prover_input_proto_input_proto_rawDescData
}

var file_src_prover_input_proto_input_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_prover_input_proto_input_proto_goTypes = []any{
	(*ProverInput)(nil), // 0: input.ProverInput
	(*Witness)(nil),     // 1: input.Witness
	(*Block)(nil),       // 2: input.Block
	(*ChainConfig)(nil), // 3: input.ChainConfig
	(*Header)(nil),      // 4: input.Header
}
var file_src_prover_input_proto_input_proto_depIdxs = []int32{
	2, // 0: input.ProverInput.blocks:type_name -> input.Block
	1, // 1: input.ProverInput.witness:type_name -> input.Witness
	3, // 2: input.ProverInput.chain_config:type_name -> input.ChainConfig
	4, // 3: input.Witness.ancestors:type_name -> input.Header
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_src_prover_input_proto_input_proto_init() }
func file_src_prover_input_proto_input_proto_init() {
	if File_src_prover_input_proto_input_proto != nil {
		return
	}
	file_src_prover_input_proto_block_proto_init()
	file_src_prover_input_proto_chain_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_prover_input_proto_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_prover_input_proto_input_proto_goTypes,
		DependencyIndexes: file_src_prover_input_proto_input_proto_depIdxs,
		MessageInfos:      file_src_prover_input_proto_input_proto_msgTypes,
	}.Build()
	File_src_prover_input_proto_input_proto = out.File
	file_src_prover_input_proto_input_proto_rawDesc = nil
	file_src_prover_input_proto_input_proto_goTypes = nil
	file_src_prover_input_proto_input_proto_depIdxs = nil
}
