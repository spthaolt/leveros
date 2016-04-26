// Code generated by protoc-gen-go.
// source: fleettracker.proto
// DO NOT EDIT!

/*
Package fleettracker is a generated protocol buffer package.

It is generated from these files:
	fleettracker.proto

It has these top-level messages:
	RPCEvent
	HandleReply
*/
package fleettracker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type RPCEvent struct {
	// The sessionID of the resource.
	SessionID   string `protobuf:"bytes,1,opt,name=sessionID" json:"sessionID,omitempty"`
	Environment string `protobuf:"bytes,2,opt,name=environment" json:"environment,omitempty"`
	Service     string `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	ServingID   string `protobuf:"bytes,4,opt,name=servingID" json:"servingID,omitempty"`
	CodeVersion int64  `protobuf:"varint,5,opt,name=codeVersion" json:"codeVersion,omitempty"`
	IsAdmin     bool   `protobuf:"varint,6,opt,name=isAdmin" json:"isAdmin,omitempty"`
	RpcNanos    uint64 `protobuf:"varint,7,opt,name=rpcNanos" json:"rpcNanos,omitempty"`
}

func (m *RPCEvent) Reset()                    { *m = RPCEvent{} }
func (m *RPCEvent) String() string            { return proto.CompactTextString(m) }
func (*RPCEvent) ProtoMessage()               {}
func (*RPCEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HandleReply struct {
}

func (m *HandleReply) Reset()                    { *m = HandleReply{} }
func (m *HandleReply) String() string            { return proto.CompactTextString(m) }
func (*HandleReply) ProtoMessage()               {}
func (*HandleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*RPCEvent)(nil), "fleettracker.RPCEvent")
	proto.RegisterType((*HandleReply)(nil), "fleettracker.HandleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for FleetTracker service

type FleetTrackerClient interface {
	HandleOnRPC(ctx context.Context, in *RPCEvent, opts ...grpc.CallOption) (*HandleReply, error)
}

type fleetTrackerClient struct {
	cc *grpc.ClientConn
}

func NewFleetTrackerClient(cc *grpc.ClientConn) FleetTrackerClient {
	return &fleetTrackerClient{cc}
}

func (c *fleetTrackerClient) HandleOnRPC(ctx context.Context, in *RPCEvent, opts ...grpc.CallOption) (*HandleReply, error) {
	out := new(HandleReply)
	err := grpc.Invoke(ctx, "/fleettracker.FleetTracker/HandleOnRPC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FleetTracker service

type FleetTrackerServer interface {
	HandleOnRPC(context.Context, *RPCEvent) (*HandleReply, error)
}

func RegisterFleetTrackerServer(s *grpc.Server, srv FleetTrackerServer) {
	s.RegisterService(&_FleetTracker_serviceDesc, srv)
}

func _FleetTracker_HandleOnRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetTrackerServer).HandleOnRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleettracker.FleetTracker/HandleOnRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetTrackerServer).HandleOnRPC(ctx, req.(*RPCEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _FleetTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fleettracker.FleetTracker",
	HandlerType: (*FleetTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleOnRPC",
			Handler:    _FleetTracker_HandleOnRPC_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x09, 0x2d, 0x6d, 0x7a, 0x5b, 0x96, 0x3b, 0x20, 0x53, 0x31, 0x54, 0x99, 0x98, 0x3a,
	0xc0, 0x13, 0xf0, 0x2b, 0x58, 0x00, 0x59, 0x88, 0x3d, 0x38, 0x17, 0x64, 0x91, 0x5e, 0x47, 0xb6,
	0x15, 0x89, 0x57, 0xe5, 0x69, 0xb0, 0x53, 0x4c, 0xdd, 0x2d, 0xe7, 0x7c, 0xd1, 0x67, 0xdd, 0x03,
	0xf8, 0xd1, 0x12, 0x79, 0x6f, 0x6b, 0xf5, 0x45, 0x76, 0xdd, 0x59, 0xe3, 0x0d, 0x2e, 0xf2, 0xae,
	0xfa, 0x29, 0xa0, 0x94, 0x2f, 0x37, 0x77, 0x3d, 0xb1, 0xc7, 0x33, 0x98, 0x39, 0x72, 0x4e, 0x1b,
	0x7e, 0xbc, 0x15, 0xc5, 0xaa, 0x38, 0x9f, 0xc9, 0x5d, 0x81, 0x2b, 0x98, 0x13, 0xf7, 0xda, 0x1a,
	0xde, 0x84, 0x9f, 0xc5, 0xe1, 0xc0, 0xf3, 0x0a, 0x05, 0x4c, 0x1d, 0xd9, 0x5e, 0x2b, 0x12, 0xa3,
	0x81, 0xa6, 0xb8, 0x35, 0x87, 0x4f, 0xfe, 0x0c, 0xe6, 0x71, 0x32, 0xff, 0x15, 0xd1, 0xac, 0x4c,
	0x43, 0x6f, 0x64, 0xe3, 0x53, 0xe2, 0x28, 0xf0, 0x91, 0xcc, 0xab, 0x68, 0xd6, 0xee, 0xaa, 0xd9,
	0x68, 0x16, 0x93, 0x40, 0x4b, 0x99, 0x22, 0x2e, 0xa1, 0xb4, 0x9d, 0x7a, 0xaa, 0xd9, 0x38, 0x31,
	0x0d, 0x68, 0x2c, 0xff, 0x73, 0x75, 0x0c, 0xf3, 0x87, 0x9a, 0x9b, 0x96, 0x24, 0x75, 0xed, 0xf7,
	0x85, 0x84, 0xc5, 0x7d, 0xbc, 0xfd, 0x75, 0x7b, 0x3b, 0x5e, 0x27, 0xfc, 0xcc, 0x61, 0x02, 0x3c,
	0x59, 0xef, 0xad, 0x95, 0x56, 0x59, 0x9e, 0xee, 0xf7, 0x99, 0xb1, 0x3a, 0x78, 0x9f, 0x0c, 0xa3,
	0x5e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x76, 0x51, 0x46, 0x36, 0x6a, 0x01, 0x00, 0x00,
}
