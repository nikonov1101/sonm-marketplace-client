// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type ListReply struct {
	Info map[string]*ListReply_ListValue `protobuf:"bytes,1,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ListReply) GetInfo() map[string]*ListReply_ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListReply_ListValue struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListReply_ListValue) Reset()                    { *m = ListReply_ListValue{} }
func (m *ListReply_ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListReply_ListValue) ProtoMessage()               {}
func (*ListReply_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func (m *ListReply_ListValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type HubInfoRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubInfoRequest) Reset()                    { *m = HubInfoRequest{} }
func (m *HubInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*HubInfoRequest) ProtoMessage()               {}
func (*HubInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *HubInfoRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type TaskRequirements struct {
	// How much resources consumes a task. This is used both for scheduling and for cgroups configuration.
	Resources *TaskResourceRequirements `protobuf:"bytes,1,opt,name=resources" json:"resources,omitempty"`
	// Optional miner ID restrictions (currently IP:port), that are allowed to start a task.
	Miners []string `protobuf:"bytes,2,rep,name=miners" json:"miners,omitempty"`
}

func (m *TaskRequirements) Reset()                    { *m = TaskRequirements{} }
func (m *TaskRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskRequirements) ProtoMessage()               {}
func (*TaskRequirements) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TaskRequirements) GetResources() *TaskResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskRequirements) GetMiners() []string {
	if m != nil {
		return m.Miners
	}
	return nil
}

type HubStartTaskRequest struct {
	Requirements  *TaskRequirements `protobuf:"bytes,1,opt,name=requirements" json:"requirements,omitempty"`
	Registry      string            `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string            `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	PublicKeyData string            `protobuf:"bytes,5,opt,name=PublicKeyData,json=publicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool              `protobuf:"varint,6,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string `protobuf:"bytes,7,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HubStartTaskRequest) Reset()                    { *m = HubStartTaskRequest{} }
func (m *HubStartTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskRequest) ProtoMessage()               {}
func (*HubStartTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *HubStartTaskRequest) GetRequirements() *TaskRequirements {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *HubStartTaskRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *HubStartTaskRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *HubStartTaskRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *HubStartTaskRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *HubStartTaskRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *HubStartTaskRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type HubStartTaskReply struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Endpoint []string `protobuf:"bytes,2,rep,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HubStartTaskReply) Reset()                    { *m = HubStartTaskReply{} }
func (m *HubStartTaskReply) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskReply) ProtoMessage()               {}
func (*HubStartTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *HubStartTaskReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HubStartTaskReply) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type HubStatusMapRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubStatusMapRequest) Reset()                    { *m = HubStatusMapRequest{} }
func (m *HubStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusMapRequest) ProtoMessage()               {}
func (*HubStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *HubStatusMapRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type HubStatusRequest struct {
}

func (m *HubStatusRequest) Reset()                    { *m = HubStatusRequest{} }
func (m *HubStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusRequest) ProtoMessage()               {}
func (*HubStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type HubStatusReply struct {
	MinerCount uint64 `protobuf:"varint,1,opt,name=minerCount" json:"minerCount,omitempty"`
	Uptime     uint64 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Platform   string `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
	EthAddr    string `protobuf:"bytes,5,opt,name=ethAddr" json:"ethAddr,omitempty"`
}

func (m *HubStatusReply) Reset()                    { *m = HubStatusReply{} }
func (m *HubStatusReply) String() string            { return proto.CompactTextString(m) }
func (*HubStatusReply) ProtoMessage()               {}
func (*HubStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *HubStatusReply) GetMinerCount() uint64 {
	if m != nil {
		return m.MinerCount
	}
	return 0
}

func (m *HubStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *HubStatusReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HubStatusReply) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *HubStatusReply) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

type DealRequest struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *DealRequest) Reset()                    { *m = DealRequest{} }
func (m *DealRequest) String() string            { return proto.CompactTextString(m) }
func (*DealRequest) ProtoMessage()               {}
func (*DealRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *DealRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type DealReply struct {
}

func (m *DealReply) Reset()                    { *m = DealReply{} }
func (m *DealReply) String() string            { return proto.CompactTextString(m) }
func (*DealReply) ProtoMessage()               {}
func (*DealReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func init() {
	proto.RegisterType((*ListRequest)(nil), "sonm.ListRequest")
	proto.RegisterType((*ListReply)(nil), "sonm.ListReply")
	proto.RegisterType((*ListReply_ListValue)(nil), "sonm.ListReply.ListValue")
	proto.RegisterType((*HubInfoRequest)(nil), "sonm.HubInfoRequest")
	proto.RegisterType((*TaskRequirements)(nil), "sonm.TaskRequirements")
	proto.RegisterType((*HubStartTaskRequest)(nil), "sonm.HubStartTaskRequest")
	proto.RegisterType((*HubStartTaskReply)(nil), "sonm.HubStartTaskReply")
	proto.RegisterType((*HubStatusMapRequest)(nil), "sonm.HubStatusMapRequest")
	proto.RegisterType((*HubStatusRequest)(nil), "sonm.HubStatusRequest")
	proto.RegisterType((*HubStatusReply)(nil), "sonm.HubStatusReply")
	proto.RegisterType((*DealRequest)(nil), "sonm.DealRequest")
	proto.RegisterType((*DealReply)(nil), "sonm.DealReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hub service

type HubClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error)
	StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error)
	ProposeDeal(ctx context.Context, in *DealRequest, opts ...grpc.CallOption) (*DealReply, error)
}

type hubClient struct {
	cc *grpc.ClientConn
}

func NewHubClient(cc *grpc.ClientConn) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error) {
	out := new(HubStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error) {
	out := new(HubStartTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StartTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StopTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/TaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error) {
	out := new(StatusMapReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/MinerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hub_serviceDesc.Streams[0], c.cc, "/sonm.Hub/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type hubTaskLogsClient struct {
	grpc.ClientStream
}

func (x *hubTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hubClient) ProposeDeal(ctx context.Context, in *DealRequest, opts ...grpc.CallOption) (*DealReply, error) {
	out := new(DealReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/ProposeDeal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hub service

type HubServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Status(context.Context, *HubStatusRequest) (*HubStatusReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Info(context.Context, *HubInfoRequest) (*InfoReply, error)
	StartTask(context.Context, *HubStartTaskRequest) (*HubStartTaskReply, error)
	StopTask(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TaskStatus(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	MinerStatus(context.Context, *HubStatusMapRequest) (*StatusMapReply, error)
	TaskLogs(*TaskLogsRequest, Hub_TaskLogsServer) error
	ProposeDeal(context.Context, *DealRequest) (*DealReply, error)
}

func RegisterHubServer(s *grpc.Server, srv HubServer) {
	s.RegisterService(&_Hub_serviceDesc, srv)
}

func _Hub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Status(ctx, req.(*HubStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Info(ctx, req.(*HubInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StartTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StartTask(ctx, req.(*HubStartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StopTask(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).TaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/TaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).TaskStatus(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_MinerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).MinerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/MinerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).MinerStatus(ctx, req.(*HubStatusMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).TaskLogs(m, &hubTaskLogsServer{stream})
}

type Hub_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type hubTaskLogsServer struct {
	grpc.ServerStream
}

func (x *hubTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Hub_ProposeDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).ProposeDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/ProposeDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).ProposeDeal(ctx, req.(*DealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hub_Ping_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Hub_Status_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Hub_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Hub_Info_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _Hub_StartTask_Handler,
		},
		{
			MethodName: "StopTask",
			Handler:    _Hub_StopTask_Handler,
		},
		{
			MethodName: "TaskStatus",
			Handler:    _Hub_TaskStatus_Handler,
		},
		{
			MethodName: "MinerStatus",
			Handler:    _Hub_MinerStatus_Handler,
		},
		{
			MethodName: "ProposeDeal",
			Handler:    _Hub_ProposeDeal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskLogs",
			Handler:       _Hub_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hub.proto",
}

func init() { proto.RegisterFile("hub.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x55, 0xed, 0x4e, 0xdb, 0x3c,
	0x14, 0x6e, 0xda, 0xb4, 0x34, 0x27, 0x7c, 0x9a, 0xaf, 0x90, 0x1f, 0xa8, 0x6f, 0xde, 0x57, 0xaf,
	0x90, 0xc6, 0x0a, 0x82, 0x69, 0x42, 0x68, 0xda, 0x86, 0x00, 0x89, 0x69, 0x20, 0x50, 0x98, 0xf6,
	0x3f, 0xa5, 0xa6, 0xb5, 0x68, 0xec, 0xcc, 0x71, 0x2a, 0xf5, 0x4e, 0x76, 0x1f, 0xfb, 0xbb, 0x7b,
	0xd8, 0x2d, 0x4d, 0xb6, 0xe3, 0xd4, 0x85, 0x6e, 0xff, 0xfc, 0x3c, 0x3e, 0xcf, 0x39, 0x39, 0xe7,
	0x3c, 0x6e, 0xc1, 0x1b, 0x16, 0xbd, 0x6e, 0xc6, 0x99, 0x60, 0xc8, 0xcd, 0x19, 0x4d, 0x43, 0xaf,
	0x47, 0xfa, 0x9a, 0x08, 0x57, 0x08, 0x95, 0x14, 0x25, 0x89, 0x26, 0xa2, 0x25, 0xf0, 0xaf, 0x49,
	0x2e, 0x62, 0xfc, 0xad, 0xc0, 0xb9, 0x88, 0x7e, 0x38, 0xe0, 0x69, 0x9c, 0x8d, 0x26, 0xe8, 0x35,
	0xb8, 0x84, 0x3e, 0xb2, 0xc0, 0xe9, 0x34, 0xf6, 0xfc, 0xa3, 0x9d, 0xae, 0x94, 0x76, 0xab, 0xeb,
	0xee, 0x27, 0xfa, 0xc8, 0x2e, 0xa9, 0xe0, 0x93, 0x58, 0x85, 0x85, 0xff, 0x6a, 0xed, 0xd7, 0x64,
	0x54, 0x60, 0xb4, 0x05, 0xad, 0xb1, 0x3c, 0xe4, 0x4a, 0xed, 0xc5, 0x25, 0x0a, 0x63, 0xf0, 0x2a,
	0x1d, 0x5a, 0x85, 0xc6, 0x13, 0x9e, 0x04, 0x4e, 0xc7, 0xd9, 0xf3, 0x62, 0x79, 0x44, 0x07, 0xd0,
	0x54, 0x81, 0x41, 0xbd, 0xe3, 0xcc, 0xab, 0x59, 0x15, 0x88, 0x75, 0xdc, 0x69, 0xfd, 0xc4, 0x89,
	0xfe, 0x87, 0xe5, 0xab, 0xa2, 0x27, 0xd3, 0x96, 0x7d, 0xa0, 0x0d, 0x68, 0xa6, 0x84, 0x62, 0x5e,
	0xa6, 0xd6, 0x20, 0x1a, 0xc2, 0xea, 0x97, 0x24, 0x7f, 0x92, 0x41, 0x84, 0xe3, 0x14, 0x53, 0x91,
	0xa3, 0x77, 0xe0, 0x71, 0x9c, 0xb3, 0x82, 0x3f, 0xa8, 0x4f, 0x95, 0x45, 0x77, 0x75, 0x51, 0x1d,
	0xaa, 0xaf, 0x6c, 0x49, 0x3c, 0x15, 0xc8, 0x2e, 0x55, 0xea, 0x3c, 0xa8, 0xeb, 0x2e, 0x35, 0x8a,
	0x7e, 0xd5, 0x61, 0xfd, 0xaa, 0xe8, 0xdd, 0x8b, 0x84, 0x0b, 0x53, 0x52, 0x7e, 0xd7, 0x29, 0x2c,
	0x72, 0x2b, 0x55, 0x59, 0x70, 0xcb, 0x2e, 0x68, 0x15, 0x9a, 0x89, 0x45, 0x21, 0xb4, 0x39, 0x1e,
	0x90, 0x5c, 0xf0, 0x89, 0x9a, 0x8e, 0x17, 0x57, 0x58, 0xf6, 0x4b, 0xd2, 0x64, 0x80, 0x83, 0x86,
	0xee, 0x57, 0x01, 0x84, 0xc0, 0x4d, 0x0a, 0x31, 0x0c, 0x5c, 0x45, 0xaa, 0x33, 0xfa, 0x0f, 0x96,
	0xee, 0x8a, 0xde, 0x88, 0x3c, 0x7c, 0xc6, 0x93, 0x8b, 0x44, 0x24, 0x41, 0x53, 0x5d, 0x2e, 0x65,
	0x36, 0x89, 0x22, 0x58, 0x7c, 0x60, 0x69, 0x4a, 0xc4, 0x2d, 0xbd, 0x17, 0x2c, 0x0b, 0x5a, 0x1d,
	0x67, 0xaf, 0x1d, 0xcf, 0x70, 0xe8, 0x0d, 0x34, 0x30, 0x1d, 0x07, 0x0b, 0xca, 0x1c, 0x91, 0x6e,
	0x61, 0x4e, 0xcf, 0xdd, 0x4b, 0x3a, 0xd6, 0x2e, 0x91, 0xe1, 0xe1, 0x5b, 0x68, 0x1b, 0x62, 0xce,
	0xfa, 0x37, 0xec, 0xf5, 0x7b, 0xf6, 0x8e, 0x3f, 0xc0, 0xda, 0x6c, 0x72, 0x69, 0xd0, 0x65, 0xa8,
	0x93, 0x7e, 0xa9, 0xaf, 0x93, 0xbe, 0x1c, 0x11, 0xa6, 0xfd, 0x8c, 0x11, 0x2a, 0xca, 0x85, 0x54,
	0x38, 0x7a, 0x65, 0x36, 0x22, 0x8a, 0xfc, 0x26, 0xc9, 0xfe, 0xee, 0x14, 0x04, 0xab, 0x55, 0xb0,
	0x79, 0x1b, 0xdf, 0x1d, 0x65, 0x33, 0x43, 0xca, 0xfa, 0xbb, 0x00, 0x2a, 0xfe, 0x9c, 0x15, 0x54,
	0xa8, 0x0c, 0x6e, 0x6c, 0x31, 0xd2, 0x1e, 0x45, 0x26, 0x48, 0xaa, 0xfb, 0x71, 0xe3, 0x12, 0xa1,
	0x00, 0x16, 0xc6, 0x98, 0xe7, 0x84, 0xd1, 0x72, 0x61, 0x06, 0xca, 0x0e, 0xb2, 0x51, 0x22, 0x1e,
	0x19, 0x4f, 0xcb, 0xb5, 0x55, 0x58, 0xaa, 0xb0, 0x18, 0x9e, 0xf5, 0xfb, 0xbc, 0x5c, 0x9a, 0x81,
	0xd1, 0x21, 0xf8, 0x17, 0x38, 0x19, 0x99, 0x9e, 0xfe, 0x81, 0x26, 0xe3, 0xfd, 0xb2, 0x27, 0xff,
	0xc8, 0xd7, 0xbb, 0xb9, 0x95, 0x54, 0xac, 0x6f, 0x22, 0x1f, 0x3c, 0xad, 0xc8, 0x46, 0x93, 0xa3,
	0x9f, 0x2e, 0x34, 0xae, 0x8a, 0x1e, 0xda, 0x07, 0xf7, 0x8e, 0xd0, 0x01, 0x5a, 0xd3, 0x02, 0x79,
	0x2e, 0x53, 0x86, 0x2b, 0x36, 0x95, 0x8d, 0x26, 0x51, 0x0d, 0x9d, 0x40, 0x4b, 0xcf, 0x02, 0x6d,
	0xd9, 0xcb, 0x9f, 0x4e, 0x2c, 0xdc, 0x78, 0xc1, 0x6b, 0xe5, 0x3e, 0xb8, 0xf2, 0x1d, 0x9b, 0x3a,
	0xd6, 0x0f, 0x90, 0xa9, 0x53, 0x3d, 0xf8, 0xa8, 0x86, 0x0e, 0xc0, 0x95, 0x4f, 0x1b, 0x4d, 0xb3,
	0x59, 0x2f, 0xdd, 0x08, 0x34, 0xa5, 0x05, 0x67, 0xe0, 0x55, 0x3e, 0x41, 0x3b, 0x7f, 0x34, 0x66,
	0xb8, 0x3d, 0xef, 0xca, 0xf4, 0xd6, 0x96, 0x1e, 0x57, 0x19, 0x36, 0x75, 0x98, 0xc1, 0x46, 0xbd,
	0xfe, 0x9c, 0xd6, 0xca, 0xf7, 0x00, 0x12, 0x96, 0x93, 0xd9, 0x9e, 0xbe, 0xec, 0xd9, 0xd1, 0x6c,
	0xbe, 0xbc, 0xd0, 0xfa, 0x8f, 0xe0, 0xdf, 0x48, 0x03, 0x95, 0x09, 0x76, 0x9e, 0x8d, 0x70, 0xea,
	0x5c, 0x33, 0x5d, 0x8b, 0xd7, 0x19, 0x4e, 0xa1, 0x2d, 0xd3, 0x5e, 0xb3, 0x41, 0x8e, 0xac, 0x32,
	0x12, 0x3f, 0xfb, 0x76, 0x43, 0x9f, 0x0f, 0x0b, 0xfa, 0x14, 0xd5, 0x0e, 0x1d, 0x74, 0x0c, 0xfe,
	0x1d, 0x67, 0x19, 0xcb, 0xb1, 0x74, 0x87, 0x59, 0x90, 0xe5, 0x2d, 0x33, 0xef, 0xca, 0x3c, 0x51,
	0xad, 0xd7, 0x52, 0x7f, 0x25, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x4e, 0x52, 0x4e,
	0x79, 0x06, 0x00, 0x00,
}