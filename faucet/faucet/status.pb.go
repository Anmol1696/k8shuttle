// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: faucet/status.proto

package faucet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucet_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_faucet_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_faucet_status_proto_rawDescGZIP(), []int{0}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type AddressBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance []*Coin `protobuf:"bytes,2,rep,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AddressBalance) Reset() {
	*x = AddressBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucet_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBalance) ProtoMessage() {}

func (x *AddressBalance) ProtoReflect() protoreflect.Message {
	mi := &file_faucet_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBalance.ProtoReflect.Descriptor instead.
func (*AddressBalance) Descriptor() ([]byte, []int) {
	return file_faucet_status_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBalance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressBalance) GetBalance() []*Coin {
	if x != nil {
		return x.Balance
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	NodeUrl         string            `protobuf:"bytes,2,opt,name=node_url,json=nodeUrl,proto3" json:"node_url,omitempty"`
	ChainId         string            `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainTokens     []string          `protobuf:"bytes,4,rep,name=chain_tokens,json=chainTokens,proto3" json:"chain_tokens,omitempty"`
	AvailableTokens []string          `protobuf:"bytes,5,rep,name=available_tokens,json=availableTokens,proto3" json:"available_tokens,omitempty"`
	Holder          *AddressBalance   `protobuf:"bytes,6,opt,name=holder,proto3" json:"holder,omitempty"`
	Distributors    []*AddressBalance `protobuf:"bytes,7,rep,name=distributors,proto3" json:"distributors,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucet_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_faucet_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_faucet_status_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *State) GetNodeUrl() string {
	if x != nil {
		return x.NodeUrl
	}
	return ""
}

func (x *State) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *State) GetChainTokens() []string {
	if x != nil {
		return x.ChainTokens
	}
	return nil
}

func (x *State) GetAvailableTokens() []string {
	if x != nil {
		return x.AvailableTokens
	}
	return nil
}

func (x *State) GetHolder() *AddressBalance {
	if x != nil {
		return x.Holder
	}
	return nil
}

func (x *State) GetDistributors() []*AddressBalance {
	if x != nil {
		return x.Distributors
	}
	return nil
}

var File_faucet_status_proto protoreflect.FileDescriptor

var file_faucet_status_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x34, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x26, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x3a, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x7c, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2d, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x46, 0x61, 0x75, 0x63,
	0x65, 0x74, 0xca, 0x02, 0x06, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0xe2, 0x02, 0x12, 0x46, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x06, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_faucet_status_proto_rawDescOnce sync.Once
	file_faucet_status_proto_rawDescData = file_faucet_status_proto_rawDesc
)

func file_faucet_status_proto_rawDescGZIP() []byte {
	file_faucet_status_proto_rawDescOnce.Do(func() {
		file_faucet_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_faucet_status_proto_rawDescData)
	})
	return file_faucet_status_proto_rawDescData
}

var file_faucet_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_faucet_status_proto_goTypes = []interface{}{
	(*Coin)(nil),           // 0: faucet.Coin
	(*AddressBalance)(nil), // 1: faucet.AddressBalance
	(*State)(nil),          // 2: faucet.State
}
var file_faucet_status_proto_depIdxs = []int32{
	0, // 0: faucet.AddressBalance.balance:type_name -> faucet.Coin
	1, // 1: faucet.State.holder:type_name -> faucet.AddressBalance
	1, // 2: faucet.State.distributors:type_name -> faucet.AddressBalance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_faucet_status_proto_init() }
func file_faucet_status_proto_init() {
	if File_faucet_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faucet_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file_faucet_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBalance); i {
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
		file_faucet_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_faucet_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_faucet_status_proto_goTypes,
		DependencyIndexes: file_faucet_status_proto_depIdxs,
		MessageInfos:      file_faucet_status_proto_msgTypes,
	}.Build()
	File_faucet_status_proto = out.File
	file_faucet_status_proto_rawDesc = nil
	file_faucet_status_proto_goTypes = nil
	file_faucet_status_proto_depIdxs = nil
}
