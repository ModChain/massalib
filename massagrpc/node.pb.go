// Copyright (c) 2023 MASSA LABS <info@massa.net>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: massa/model/v1/node.proto

package massagrpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConnectionType enum
type ConnectionType int32

const (
	ConnectionType_CONNECTION_TYPE_UNSPECIFIED ConnectionType = 0 // Default enum value
	ConnectionType_CONNECTION_TYPE_INCOMING    ConnectionType = 1 // Incoming connection
	ConnectionType_CONNECTION_TYPE_OUTGOING    ConnectionType = 2 // Outgoing connection
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_UNSPECIFIED",
		1: "CONNECTION_TYPE_INCOMING",
		2: "CONNECTION_TYPE_OUTGOING",
	}
	ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_UNSPECIFIED": 0,
		"CONNECTION_TYPE_INCOMING":    1,
		"CONNECTION_TYPE_OUTGOING":    2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_massa_model_v1_node_proto_enumTypes[0].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_massa_model_v1_node_proto_enumTypes[0]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_massa_model_v1_node_proto_rawDescGZIP(), []int{0}
}

// Node status
type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Our node id
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Optional node ip
	NodeIp string `protobuf:"bytes,2,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	// Node version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Now
	CurrentTime *NativeTime `protobuf:"bytes,4,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	// Current cycle
	CurrentCycle uint64 `protobuf:"varint,5,opt,name=current_cycle,json=currentCycle,proto3" json:"current_cycle,omitempty"`
	// Current cycle starting timestamp
	CurrentCycleTime *NativeTime `protobuf:"bytes,6,opt,name=current_cycle_time,json=currentCycleTime,proto3" json:"current_cycle_time,omitempty"`
	// Next cycle starting timestamp
	NextCycleTime *NativeTime `protobuf:"bytes,7,opt,name=next_cycle_time,json=nextCycleTime,proto3" json:"next_cycle_time,omitempty"`
	// Connected nodes
	ConnectedNodes []*ConnectedNode `protobuf:"bytes,8,rep,name=connected_nodes,json=connectedNodes,proto3" json:"connected_nodes,omitempty"`
	// Last executed final slot
	LastExecutedFinalSlot *Slot `protobuf:"bytes,9,opt,name=last_executed_final_slot,json=lastExecutedFinalSlot,proto3" json:"last_executed_final_slot,omitempty"`
	// Last executed speculative slot
	LastExecutedSpeculativeSlot *Slot `protobuf:"bytes,10,opt,name=last_executed_speculative_slot,json=lastExecutedSpeculativeSlot,proto3" json:"last_executed_speculative_slot,omitempty"`
	// The hash of the XOF final state hash
	FinalStateFingerprint string `protobuf:"bytes,11,opt,name=final_state_fingerprint,json=finalStateFingerprint,proto3" json:"final_state_fingerprint,omitempty"`
	// Consensus stats
	ConsensusStats *ConsensusStats `protobuf:"bytes,12,opt,name=consensus_stats,json=consensusStats,proto3" json:"consensus_stats,omitempty"`
	// Pool stats (operation count and endorsement count)
	PoolStats *PoolStats `protobuf:"bytes,13,opt,name=pool_stats,json=poolStats,proto3" json:"pool_stats,omitempty"`
	// Network stats
	NetworkStats *NetworkStats `protobuf:"bytes,14,opt,name=network_stats,json=networkStats,proto3" json:"network_stats,omitempty"`
	// Execution stats
	ExecutionStats *ExecutionStats `protobuf:"bytes,15,opt,name=execution_stats,json=executionStats,proto3" json:"execution_stats,omitempty"`
	// Compact configuration
	Config *CompactConfig `protobuf:"bytes,16,opt,name=config,proto3" json:"config,omitempty"`
	// Chain id
	ChainId uint64 `protobuf:"varint,17,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Current mip version
	CurrentMipVersion uint32 `protobuf:"varint,18,opt,name=current_mip_version,json=currentMipVersion,proto3" json:"current_mip_version,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	mi := &file_massa_model_v1_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_massa_model_v1_node_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_massa_model_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatus) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeStatus) GetNodeIp() string {
	if x != nil {
		return x.NodeIp
	}
	return ""
}

func (x *NodeStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeStatus) GetCurrentTime() *NativeTime {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

func (x *NodeStatus) GetCurrentCycle() uint64 {
	if x != nil {
		return x.CurrentCycle
	}
	return 0
}

func (x *NodeStatus) GetCurrentCycleTime() *NativeTime {
	if x != nil {
		return x.CurrentCycleTime
	}
	return nil
}

func (x *NodeStatus) GetNextCycleTime() *NativeTime {
	if x != nil {
		return x.NextCycleTime
	}
	return nil
}

func (x *NodeStatus) GetConnectedNodes() []*ConnectedNode {
	if x != nil {
		return x.ConnectedNodes
	}
	return nil
}

func (x *NodeStatus) GetLastExecutedFinalSlot() *Slot {
	if x != nil {
		return x.LastExecutedFinalSlot
	}
	return nil
}

func (x *NodeStatus) GetLastExecutedSpeculativeSlot() *Slot {
	if x != nil {
		return x.LastExecutedSpeculativeSlot
	}
	return nil
}

func (x *NodeStatus) GetFinalStateFingerprint() string {
	if x != nil {
		return x.FinalStateFingerprint
	}
	return ""
}

func (x *NodeStatus) GetConsensusStats() *ConsensusStats {
	if x != nil {
		return x.ConsensusStats
	}
	return nil
}

func (x *NodeStatus) GetPoolStats() *PoolStats {
	if x != nil {
		return x.PoolStats
	}
	return nil
}

func (x *NodeStatus) GetNetworkStats() *NetworkStats {
	if x != nil {
		return x.NetworkStats
	}
	return nil
}

func (x *NodeStatus) GetExecutionStats() *ExecutionStats {
	if x != nil {
		return x.ExecutionStats
	}
	return nil
}

func (x *NodeStatus) GetConfig() *CompactConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *NodeStatus) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *NodeStatus) GetCurrentMipVersion() uint32 {
	if x != nil {
		return x.CurrentMipVersion
	}
	return 0
}

// Connected node
type ConnectedNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node id
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node ip
	NodeIp string `protobuf:"bytes,2,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	// Connection type
	ConnectionType ConnectionType `protobuf:"varint,3,opt,name=connection_type,json=connectionType,proto3,enum=massa.model.v1.ConnectionType" json:"connection_type,omitempty"`
}

func (x *ConnectedNode) Reset() {
	*x = ConnectedNode{}
	mi := &file_massa_model_v1_node_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectedNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedNode) ProtoMessage() {}

func (x *ConnectedNode) ProtoReflect() protoreflect.Message {
	mi := &file_massa_model_v1_node_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedNode.ProtoReflect.Descriptor instead.
func (*ConnectedNode) Descriptor() ([]byte, []int) {
	return file_massa_model_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectedNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ConnectedNode) GetNodeIp() string {
	if x != nil {
		return x.NodeIp
	}
	return ""
}

func (x *ConnectedNode) GetConnectionType() ConnectionType {
	if x != nil {
		return x.ConnectionType
	}
	return ConnectionType_CONNECTION_TYPE_UNSPECIFIED
}

// Compact configuration
type CompactConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time in milliseconds when the blockclique started.
	GenesisTimestamp *NativeTime `protobuf:"bytes,1,opt,name=genesis_timestamp,json=genesisTimestamp,proto3" json:"genesis_timestamp,omitempty"`
	// TESTNET: time when the blockclique is ended.
	EndTimestamp *NativeTime `protobuf:"bytes,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	// Number of threads
	ThreadCount uint32 `protobuf:"varint,3,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
	// Time between the periods in the same thread.
	T0 *NativeTime `protobuf:"bytes,4,opt,name=t0,proto3" json:"t0,omitempty"`
	// Threshold for fitness.
	DeltaF0 uint64 `protobuf:"varint,5,opt,name=delta_f0,json=deltaF0,proto3" json:"delta_f0,omitempty"`
	// Maximum operation validity period count
	OperationValidityPeriods uint64 `protobuf:"varint,6,opt,name=operation_validity_periods,json=operationValidityPeriods,proto3" json:"operation_validity_periods,omitempty"`
	// cycle duration in periods
	PeriodsPerCycle uint64 `protobuf:"varint,7,opt,name=periods_per_cycle,json=periodsPerCycle,proto3" json:"periods_per_cycle,omitempty"`
	// Reward amount for a block creation
	BlockReward *NativeAmount `protobuf:"bytes,8,opt,name=block_reward,json=blockReward,proto3" json:"block_reward,omitempty"`
	// Price of a roll on the network
	RollPrice *NativeAmount `protobuf:"bytes,9,opt,name=roll_price,json=rollPrice,proto3" json:"roll_price,omitempty"`
	// Max total size of a block
	MaxBlockSize uint32 `protobuf:"varint,10,opt,name=max_block_size,json=maxBlockSize,proto3" json:"max_block_size,omitempty"`
}

func (x *CompactConfig) Reset() {
	*x = CompactConfig{}
	mi := &file_massa_model_v1_node_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactConfig) ProtoMessage() {}

func (x *CompactConfig) ProtoReflect() protoreflect.Message {
	mi := &file_massa_model_v1_node_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactConfig.ProtoReflect.Descriptor instead.
func (*CompactConfig) Descriptor() ([]byte, []int) {
	return file_massa_model_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *CompactConfig) GetGenesisTimestamp() *NativeTime {
	if x != nil {
		return x.GenesisTimestamp
	}
	return nil
}

func (x *CompactConfig) GetEndTimestamp() *NativeTime {
	if x != nil {
		return x.EndTimestamp
	}
	return nil
}

func (x *CompactConfig) GetThreadCount() uint32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

func (x *CompactConfig) GetT0() *NativeTime {
	if x != nil {
		return x.T0
	}
	return nil
}

func (x *CompactConfig) GetDeltaF0() uint64 {
	if x != nil {
		return x.DeltaF0
	}
	return 0
}

func (x *CompactConfig) GetOperationValidityPeriods() uint64 {
	if x != nil {
		return x.OperationValidityPeriods
	}
	return 0
}

func (x *CompactConfig) GetPeriodsPerCycle() uint64 {
	if x != nil {
		return x.PeriodsPerCycle
	}
	return 0
}

func (x *CompactConfig) GetBlockReward() *NativeAmount {
	if x != nil {
		return x.BlockReward
	}
	return nil
}

func (x *CompactConfig) GetRollPrice() *NativeAmount {
	if x != nil {
		return x.RollPrice
	}
	return nil
}

func (x *CompactConfig) GetMaxBlockSize() uint32 {
	if x != nil {
		return x.MaxBlockSize
	}
	return 0
}

// Public status
type PublicStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Our node id
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Now
	CurrentTime *NativeTime `protobuf:"bytes,4,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	// Current cycle
	CurrentCycle uint64 `protobuf:"varint,5,opt,name=current_cycle,json=currentCycle,proto3" json:"current_cycle,omitempty"`
	// Current cycle starting timestamp
	CurrentCycleTime *NativeTime `protobuf:"bytes,6,opt,name=current_cycle_time,json=currentCycleTime,proto3" json:"current_cycle_time,omitempty"`
	// Next cycle starting timestamp
	NextCycleTime *NativeTime `protobuf:"bytes,7,opt,name=next_cycle_time,json=nextCycleTime,proto3" json:"next_cycle_time,omitempty"`
	// Last executed final slot
	LastExecutedFinalSlot *Slot `protobuf:"bytes,8,opt,name=last_executed_final_slot,json=lastExecutedFinalSlot,proto3" json:"last_executed_final_slot,omitempty"`
	// Last executed speculative slot
	LastExecutedSpeculativeSlot *Slot `protobuf:"bytes,9,opt,name=last_executed_speculative_slot,json=lastExecutedSpeculativeSlot,proto3" json:"last_executed_speculative_slot,omitempty"`
	// The hash of the XOF final state hash
	FinalStateFingerprint string `protobuf:"bytes,10,opt,name=final_state_fingerprint,json=finalStateFingerprint,proto3" json:"final_state_fingerprint,omitempty"`
	// Compact configuration
	Config *CompactConfig `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	// Chain id
	ChainId uint64 `protobuf:"varint,12,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// minimal fees
	MinimalFees *NativeAmount `protobuf:"bytes,13,opt,name=minimal_fees,json=minimalFees,proto3" json:"minimal_fees,omitempty"`
	// current mip version
	CurrentMipVersion uint32 `protobuf:"varint,14,opt,name=current_mip_version,json=currentMipVersion,proto3" json:"current_mip_version,omitempty"`
}

func (x *PublicStatus) Reset() {
	*x = PublicStatus{}
	mi := &file_massa_model_v1_node_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStatus) ProtoMessage() {}

func (x *PublicStatus) ProtoReflect() protoreflect.Message {
	mi := &file_massa_model_v1_node_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStatus.ProtoReflect.Descriptor instead.
func (*PublicStatus) Descriptor() ([]byte, []int) {
	return file_massa_model_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *PublicStatus) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *PublicStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PublicStatus) GetCurrentTime() *NativeTime {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

func (x *PublicStatus) GetCurrentCycle() uint64 {
	if x != nil {
		return x.CurrentCycle
	}
	return 0
}

func (x *PublicStatus) GetCurrentCycleTime() *NativeTime {
	if x != nil {
		return x.CurrentCycleTime
	}
	return nil
}

func (x *PublicStatus) GetNextCycleTime() *NativeTime {
	if x != nil {
		return x.NextCycleTime
	}
	return nil
}

func (x *PublicStatus) GetLastExecutedFinalSlot() *Slot {
	if x != nil {
		return x.LastExecutedFinalSlot
	}
	return nil
}

func (x *PublicStatus) GetLastExecutedSpeculativeSlot() *Slot {
	if x != nil {
		return x.LastExecutedSpeculativeSlot
	}
	return nil
}

func (x *PublicStatus) GetFinalStateFingerprint() string {
	if x != nil {
		return x.FinalStateFingerprint
	}
	return ""
}

func (x *PublicStatus) GetConfig() *CompactConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *PublicStatus) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *PublicStatus) GetMinimalFees() *NativeAmount {
	if x != nil {
		return x.MinimalFees
	}
	return nil
}

func (x *PublicStatus) GetCurrentMipVersion() uint32 {
	if x != nil {
		return x.CurrentMipVersion
	}
	return 0
}

var File_massa_model_v1_node_proto protoreflect.FileDescriptor

var file_massa_model_v1_node_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x73,
	0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x6d, 0x61, 0x73,
	0x73, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x08, 0x0a, 0x0a, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61,
	0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x4d, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x59,
	0x0a, 0x1e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x1b, 0x6c, 0x61,
	0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x12, 0x47, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x73,
	0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61,
	0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x69,
	0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x70, 0x12, 0x47, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x91, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x47, 0x0a, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0c, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73,
	0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x02, 0x74, 0x30, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x5f, 0x66, 0x30, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x46, 0x30, 0x12, 0x3c, 0x0a, 0x1a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x50, 0x65, 0x72, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x3b, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0xd8, 0x05, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x18, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61,
	0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x59, 0x0a, 0x1e, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x1b, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61,
	0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x69, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4d, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x6d,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x89, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x73, 0x73,
	0x61, 0x67, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x06, 0x4d, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0xaa, 0x02,
	0x12, 0x43, 0x6f, 0x6d, 0x2e, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0xba, 0x02, 0x06, 0x4d, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0xca, 0x02, 0x12, 0x43,
	0x6f, 0x6d, 0x5c, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x3a, 0x3a, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_massa_model_v1_node_proto_rawDescOnce sync.Once
	file_massa_model_v1_node_proto_rawDescData = file_massa_model_v1_node_proto_rawDesc
)

func file_massa_model_v1_node_proto_rawDescGZIP() []byte {
	file_massa_model_v1_node_proto_rawDescOnce.Do(func() {
		file_massa_model_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_massa_model_v1_node_proto_rawDescData)
	})
	return file_massa_model_v1_node_proto_rawDescData
}

var file_massa_model_v1_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_massa_model_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_massa_model_v1_node_proto_goTypes = []any{
	(ConnectionType)(0),    // 0: massa.model.v1.ConnectionType
	(*NodeStatus)(nil),     // 1: massa.model.v1.NodeStatus
	(*ConnectedNode)(nil),  // 2: massa.model.v1.ConnectedNode
	(*CompactConfig)(nil),  // 3: massa.model.v1.CompactConfig
	(*PublicStatus)(nil),   // 4: massa.model.v1.PublicStatus
	(*NativeTime)(nil),     // 5: massa.model.v1.NativeTime
	(*Slot)(nil),           // 6: massa.model.v1.Slot
	(*ConsensusStats)(nil), // 7: massa.model.v1.ConsensusStats
	(*PoolStats)(nil),      // 8: massa.model.v1.PoolStats
	(*NetworkStats)(nil),   // 9: massa.model.v1.NetworkStats
	(*ExecutionStats)(nil), // 10: massa.model.v1.ExecutionStats
	(*NativeAmount)(nil),   // 11: massa.model.v1.NativeAmount
}
var file_massa_model_v1_node_proto_depIdxs = []int32{
	5,  // 0: massa.model.v1.NodeStatus.current_time:type_name -> massa.model.v1.NativeTime
	5,  // 1: massa.model.v1.NodeStatus.current_cycle_time:type_name -> massa.model.v1.NativeTime
	5,  // 2: massa.model.v1.NodeStatus.next_cycle_time:type_name -> massa.model.v1.NativeTime
	2,  // 3: massa.model.v1.NodeStatus.connected_nodes:type_name -> massa.model.v1.ConnectedNode
	6,  // 4: massa.model.v1.NodeStatus.last_executed_final_slot:type_name -> massa.model.v1.Slot
	6,  // 5: massa.model.v1.NodeStatus.last_executed_speculative_slot:type_name -> massa.model.v1.Slot
	7,  // 6: massa.model.v1.NodeStatus.consensus_stats:type_name -> massa.model.v1.ConsensusStats
	8,  // 7: massa.model.v1.NodeStatus.pool_stats:type_name -> massa.model.v1.PoolStats
	9,  // 8: massa.model.v1.NodeStatus.network_stats:type_name -> massa.model.v1.NetworkStats
	10, // 9: massa.model.v1.NodeStatus.execution_stats:type_name -> massa.model.v1.ExecutionStats
	3,  // 10: massa.model.v1.NodeStatus.config:type_name -> massa.model.v1.CompactConfig
	0,  // 11: massa.model.v1.ConnectedNode.connection_type:type_name -> massa.model.v1.ConnectionType
	5,  // 12: massa.model.v1.CompactConfig.genesis_timestamp:type_name -> massa.model.v1.NativeTime
	5,  // 13: massa.model.v1.CompactConfig.end_timestamp:type_name -> massa.model.v1.NativeTime
	5,  // 14: massa.model.v1.CompactConfig.t0:type_name -> massa.model.v1.NativeTime
	11, // 15: massa.model.v1.CompactConfig.block_reward:type_name -> massa.model.v1.NativeAmount
	11, // 16: massa.model.v1.CompactConfig.roll_price:type_name -> massa.model.v1.NativeAmount
	5,  // 17: massa.model.v1.PublicStatus.current_time:type_name -> massa.model.v1.NativeTime
	5,  // 18: massa.model.v1.PublicStatus.current_cycle_time:type_name -> massa.model.v1.NativeTime
	5,  // 19: massa.model.v1.PublicStatus.next_cycle_time:type_name -> massa.model.v1.NativeTime
	6,  // 20: massa.model.v1.PublicStatus.last_executed_final_slot:type_name -> massa.model.v1.Slot
	6,  // 21: massa.model.v1.PublicStatus.last_executed_speculative_slot:type_name -> massa.model.v1.Slot
	3,  // 22: massa.model.v1.PublicStatus.config:type_name -> massa.model.v1.CompactConfig
	11, // 23: massa.model.v1.PublicStatus.minimal_fees:type_name -> massa.model.v1.NativeAmount
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_massa_model_v1_node_proto_init() }
func file_massa_model_v1_node_proto_init() {
	if File_massa_model_v1_node_proto != nil {
		return
	}
	file_massa_model_v1_amount_proto_init()
	file_massa_model_v1_slot_proto_init()
	file_massa_model_v1_stats_proto_init()
	file_massa_model_v1_time_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_massa_model_v1_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_massa_model_v1_node_proto_goTypes,
		DependencyIndexes: file_massa_model_v1_node_proto_depIdxs,
		EnumInfos:         file_massa_model_v1_node_proto_enumTypes,
		MessageInfos:      file_massa_model_v1_node_proto_msgTypes,
	}.Build()
	File_massa_model_v1_node_proto = out.File
	file_massa_model_v1_node_proto_rawDesc = nil
	file_massa_model_v1_node_proto_goTypes = nil
	file_massa_model_v1_node_proto_depIdxs = nil
}
