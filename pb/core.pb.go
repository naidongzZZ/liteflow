// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: pb/core.proto

package __

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

// operator type
type OpType int32

const (
	OpType_OpUnknown OpType = 0
	OpType_Map       OpType = 1
	OpType_Reduce    OpType = 2
)

// Enum value maps for OpType.
var (
	OpType_name = map[int32]string{
		0: "OpUnknown",
		1: "Map",
		2: "Reduce",
	}
	OpType_value = map[string]int32{
		"OpUnknown": 0,
		"Map":       1,
		"Reduce":    2,
	}
)

func (x OpType) Enum() *OpType {
	p := new(OpType)
	*p = x
	return p
}

func (x OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_core_proto_enumTypes[0].Descriptor()
}

func (OpType) Type() protoreflect.EnumType {
	return &file_pb_core_proto_enumTypes[0]
}

func (x OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpType.Descriptor instead.
func (OpType) EnumDescriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{0}
}

// task status
type TaskStatus int32

const (
	TaskStatus_TsUnknown TaskStatus = 0
	TaskStatus_Ready     TaskStatus = 1
	TaskStatus_Running   TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "TsUnknown",
		1: "Ready",
		2: "Running",
	}
	TaskStatus_value = map[string]int32{
		"TsUnknown": 0,
		"Ready":     1,
		"Running":   2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_core_proto_enumTypes[1].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_pb_core_proto_enumTypes[1]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{1}
}

// directed graph
type Digraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adj []*OperatorTask `protobuf:"bytes,1,rep,name=adj,proto3" json:"adj,omitempty"`
}

func (x *Digraph) Reset() {
	*x = Digraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Digraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Digraph) ProtoMessage() {}

func (x *Digraph) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Digraph.ProtoReflect.Descriptor instead.
func (*Digraph) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{0}
}

func (x *Digraph) GetAdj() []*OperatorTask {
	if x != nil {
		return x.Adj
	}
	return nil
}

type OperatorTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task identity
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// client identity
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// operator type
	OpType OpType `protobuf:"varint,3,opt,name=op_type,json=opType,proto3,enum=pb.OpType" json:"op_type,omitempty"`
	// operator identity
	OpId string `protobuf:"bytes,4,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// task manager identity
	TaskManagerId string `protobuf:"bytes,5,opt,name=task_manager_id,json=taskManagerId,proto3" json:"task_manager_id,omitempty"`
	// task status
	State TaskStatus `protobuf:"varint,6,opt,name=state,proto3,enum=pb.TaskStatus" json:"state,omitempty"`
	// many instances are executed in parallel
	Parallelism int32 `protobuf:"varint,7,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	// upstream task
	Upstream []*OperatorTask `protobuf:"bytes,8,rep,name=upstream,proto3" json:"upstream,omitempty"`
	// downstream task
	Downstream []*OperatorTask `protobuf:"bytes,9,rep,name=downstream,proto3" json:"downstream,omitempty"`
}

func (x *OperatorTask) Reset() {
	*x = OperatorTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorTask) ProtoMessage() {}

func (x *OperatorTask) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorTask.ProtoReflect.Descriptor instead.
func (*OperatorTask) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OperatorTask) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OperatorTask) GetOpType() OpType {
	if x != nil {
		return x.OpType
	}
	return OpType_OpUnknown
}

func (x *OperatorTask) GetOpId() string {
	if x != nil {
		return x.OpId
	}
	return ""
}

func (x *OperatorTask) GetTaskManagerId() string {
	if x != nil {
		return x.TaskManagerId
	}
	return ""
}

func (x *OperatorTask) GetState() TaskStatus {
	if x != nil {
		return x.State
	}
	return TaskStatus_TsUnknown
}

func (x *OperatorTask) GetParallelism() int32 {
	if x != nil {
		return x.Parallelism
	}
	return 0
}

func (x *OperatorTask) GetUpstream() []*OperatorTask {
	if x != nil {
		return x.Upstream
	}
	return nil
}

func (x *OperatorTask) GetDownstream() []*OperatorTask {
	if x != nil {
		return x.Downstream
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventTime int64  `protobuf:"varint,2,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// next task id
	TaskId string            `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Data   map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

func (x *Event) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Event) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventChannelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventChannelReq) Reset() {
	*x = EventChannelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventChannelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventChannelReq) ProtoMessage() {}

func (x *EventChannelReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventChannelReq.ProtoReflect.Descriptor instead.
func (*EventChannelReq) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{3}
}

func (x *EventChannelReq) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type EventChannelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventChannelResp) Reset() {
	*x = EventChannelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventChannelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventChannelResp) ProtoMessage() {}

func (x *EventChannelResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventChannelResp.ProtoReflect.Descriptor instead.
func (*EventChannelResp) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{4}
}

type HealthCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskManagerId string `protobuf:"bytes,1,opt,name=task_manager_id,json=taskManagerId,proto3" json:"task_manager_id,omitempty"`
	ServiceStatus int32  `protobuf:"varint,2,opt,name=service_status,json=serviceStatus,proto3" json:"service_status,omitempty"`
}

func (x *HealthCheckReq) Reset() {
	*x = HealthCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReq) ProtoMessage() {}

func (x *HealthCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReq.ProtoReflect.Descriptor instead.
func (*HealthCheckReq) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheckReq) GetTaskManagerId() string {
	if x != nil {
		return x.TaskManagerId
	}
	return ""
}

func (x *HealthCheckReq) GetServiceStatus() int32 {
	if x != nil {
		return x.ServiceStatus
	}
	return 0
}

type HealthCheckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResp) Reset() {
	*x = HealthCheckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResp) ProtoMessage() {}

func (x *HealthCheckResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResp.ProtoReflect.Descriptor instead.
func (*HealthCheckResp) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{6}
}

type SubmitOpTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digraph  *Digraph `protobuf:"bytes,1,opt,name=digraph,proto3" json:"digraph,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *SubmitOpTaskReq) Reset() {
	*x = SubmitOpTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitOpTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitOpTaskReq) ProtoMessage() {}

func (x *SubmitOpTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitOpTaskReq.ProtoReflect.Descriptor instead.
func (*SubmitOpTaskReq) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitOpTaskReq) GetDigraph() *Digraph {
	if x != nil {
		return x.Digraph
	}
	return nil
}

func (x *SubmitOpTaskReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubmitOpTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitOpTaskResp) Reset() {
	*x = SubmitOpTaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitOpTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitOpTaskResp) ProtoMessage() {}

func (x *SubmitOpTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitOpTaskResp.ProtoReflect.Descriptor instead.
func (*SubmitOpTaskResp) Descriptor() ([]byte, []int) {
	return file_pb_core_proto_rawDescGZIP(), []int{8}
}

var File_pb_core_proto protoreflect.FileDescriptor

var file_pb_core_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x2d, 0x0a, 0x07, 0x44, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x22,
	0x0a, 0x03, 0x61, 0x64, 0x6a, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x03, 0x61,
	0x64, 0x6a, 0x22, 0xc5, 0x02, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x12, 0x2c, 0x0a, 0x08, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x08,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xb1, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x34,
	0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x21, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5f, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x55, 0x0a, 0x0f,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12,
	0x25, 0x0a, 0x07, 0x64, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x07, 0x64,
	0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x70, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x2c, 0x0a, 0x06, 0x4f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x70, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x32, 0xc0, 0x01, 0x0a, 0x04, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x70, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_core_proto_rawDescOnce sync.Once
	file_pb_core_proto_rawDescData = file_pb_core_proto_rawDesc
)

func file_pb_core_proto_rawDescGZIP() []byte {
	file_pb_core_proto_rawDescOnce.Do(func() {
		file_pb_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_core_proto_rawDescData)
	})
	return file_pb_core_proto_rawDescData
}

var file_pb_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_core_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_core_proto_goTypes = []interface{}{
	(OpType)(0),              // 0: pb.OpType
	(TaskStatus)(0),          // 1: pb.TaskStatus
	(*Digraph)(nil),          // 2: pb.Digraph
	(*OperatorTask)(nil),     // 3: pb.OperatorTask
	(*Event)(nil),            // 4: pb.Event
	(*EventChannelReq)(nil),  // 5: pb.EventChannelReq
	(*EventChannelResp)(nil), // 6: pb.EventChannelResp
	(*HealthCheckReq)(nil),   // 7: pb.HealthCheckReq
	(*HealthCheckResp)(nil),  // 8: pb.HealthCheckResp
	(*SubmitOpTaskReq)(nil),  // 9: pb.SubmitOpTaskReq
	(*SubmitOpTaskResp)(nil), // 10: pb.SubmitOpTaskResp
	nil,                      // 11: pb.Event.DataEntry
}
var file_pb_core_proto_depIdxs = []int32{
	3,  // 0: pb.Digraph.adj:type_name -> pb.OperatorTask
	0,  // 1: pb.OperatorTask.op_type:type_name -> pb.OpType
	1,  // 2: pb.OperatorTask.state:type_name -> pb.TaskStatus
	3,  // 3: pb.OperatorTask.upstream:type_name -> pb.OperatorTask
	3,  // 4: pb.OperatorTask.downstream:type_name -> pb.OperatorTask
	11, // 5: pb.Event.data:type_name -> pb.Event.DataEntry
	4,  // 6: pb.EventChannelReq.events:type_name -> pb.Event
	2,  // 7: pb.SubmitOpTaskReq.digraph:type_name -> pb.Digraph
	5,  // 8: pb.core.EventChannel:input_type -> pb.EventChannelReq
	7,  // 9: pb.core.SendHeartBeat:input_type -> pb.HealthCheckReq
	9,  // 10: pb.core.SubmitOpTask:input_type -> pb.SubmitOpTaskReq
	6,  // 11: pb.core.EventChannel:output_type -> pb.EventChannelResp
	8,  // 12: pb.core.SendHeartBeat:output_type -> pb.HealthCheckResp
	10, // 13: pb.core.SubmitOpTask:output_type -> pb.SubmitOpTaskResp
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pb_core_proto_init() }
func file_pb_core_proto_init() {
	if File_pb_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Digraph); i {
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
		file_pb_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorTask); i {
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
		file_pb_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_pb_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventChannelReq); i {
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
		file_pb_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventChannelResp); i {
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
		file_pb_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckReq); i {
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
		file_pb_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResp); i {
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
		file_pb_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitOpTaskReq); i {
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
		file_pb_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitOpTaskResp); i {
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
			RawDescriptor: file_pb_core_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_core_proto_goTypes,
		DependencyIndexes: file_pb_core_proto_depIdxs,
		EnumInfos:         file_pb_core_proto_enumTypes,
		MessageInfos:      file_pb_core_proto_msgTypes,
	}.Build()
	File_pb_core_proto = out.File
	file_pb_core_proto_rawDesc = nil
	file_pb_core_proto_goTypes = nil
	file_pb_core_proto_depIdxs = nil
}
