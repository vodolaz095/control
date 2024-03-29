// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: control.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlClient interface {
	GetLine(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	SubscribeToOrders(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (Control_SubscribeToOrdersClient, error)
	GetTaskByName(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Task, error)
	ReportTelemetry(ctx context.Context, opts ...grpc.CallOption) (Control_ReportTelemetryClient, error)
	ReportTaskUpdate(ctx context.Context, opts ...grpc.CallOption) (Control_ReportTaskUpdateClient, error)
}

type controlClient struct {
	cc grpc.ClientConnInterface
}

func NewControlClient(cc grpc.ClientConnInterface) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) GetLine(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/Control/GetLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) SubscribeToOrders(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (Control_SubscribeToOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Control_ServiceDesc.Streams[0], "/Control/SubscribeToOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlSubscribeToOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Control_SubscribeToOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type controlSubscribeToOrdersClient struct {
	grpc.ClientStream
}

func (x *controlSubscribeToOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlClient) GetTaskByName(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/Control/GetTaskByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) ReportTelemetry(ctx context.Context, opts ...grpc.CallOption) (Control_ReportTelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Control_ServiceDesc.Streams[1], "/Control/ReportTelemetry", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlReportTelemetryClient{stream}
	return x, nil
}

type Control_ReportTelemetryClient interface {
	Send(*Telemetry) error
	CloseAndRecv() (*StringResponse, error)
	grpc.ClientStream
}

type controlReportTelemetryClient struct {
	grpc.ClientStream
}

func (x *controlReportTelemetryClient) Send(m *Telemetry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlReportTelemetryClient) CloseAndRecv() (*StringResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlClient) ReportTaskUpdate(ctx context.Context, opts ...grpc.CallOption) (Control_ReportTaskUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Control_ServiceDesc.Streams[2], "/Control/ReportTaskUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlReportTaskUpdateClient{stream}
	return x, nil
}

type Control_ReportTaskUpdateClient interface {
	Send(*TaskUpdate) error
	CloseAndRecv() (*StringResponse, error)
	grpc.ClientStream
}

type controlReportTaskUpdateClient struct {
	grpc.ClientStream
}

func (x *controlReportTaskUpdateClient) Send(m *TaskUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlReportTaskUpdateClient) CloseAndRecv() (*StringResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControlServer is the server API for Control service.
// All implementations must embed UnimplementedControlServer
// for forward compatibility
type ControlServer interface {
	GetLine(context.Context, *StringRequest) (*StringResponse, error)
	SubscribeToOrders(*StringRequest, Control_SubscribeToOrdersServer) error
	GetTaskByName(context.Context, *StringRequest) (*Task, error)
	ReportTelemetry(Control_ReportTelemetryServer) error
	ReportTaskUpdate(Control_ReportTaskUpdateServer) error
	mustEmbedUnimplementedControlServer()
}

// UnimplementedControlServer must be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (UnimplementedControlServer) GetLine(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLine not implemented")
}
func (UnimplementedControlServer) SubscribeToOrders(*StringRequest, Control_SubscribeToOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToOrders not implemented")
}
func (UnimplementedControlServer) GetTaskByName(context.Context, *StringRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskByName not implemented")
}
func (UnimplementedControlServer) ReportTelemetry(Control_ReportTelemetryServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportTelemetry not implemented")
}
func (UnimplementedControlServer) ReportTaskUpdate(Control_ReportTaskUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportTaskUpdate not implemented")
}
func (UnimplementedControlServer) mustEmbedUnimplementedControlServer() {}

// UnsafeControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServer will
// result in compilation errors.
type UnsafeControlServer interface {
	mustEmbedUnimplementedControlServer()
}

func RegisterControlServer(s grpc.ServiceRegistrar, srv ControlServer) {
	s.RegisterService(&Control_ServiceDesc, srv)
}

func _Control_GetLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Control/GetLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetLine(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_SubscribeToOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServer).SubscribeToOrders(m, &controlSubscribeToOrdersServer{stream})
}

type Control_SubscribeToOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type controlSubscribeToOrdersServer struct {
	grpc.ServerStream
}

func (x *controlSubscribeToOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _Control_GetTaskByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetTaskByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Control/GetTaskByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetTaskByName(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_ReportTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlServer).ReportTelemetry(&controlReportTelemetryServer{stream})
}

type Control_ReportTelemetryServer interface {
	SendAndClose(*StringResponse) error
	Recv() (*Telemetry, error)
	grpc.ServerStream
}

type controlReportTelemetryServer struct {
	grpc.ServerStream
}

func (x *controlReportTelemetryServer) SendAndClose(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlReportTelemetryServer) Recv() (*Telemetry, error) {
	m := new(Telemetry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Control_ReportTaskUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlServer).ReportTaskUpdate(&controlReportTaskUpdateServer{stream})
}

type Control_ReportTaskUpdateServer interface {
	SendAndClose(*StringResponse) error
	Recv() (*TaskUpdate, error)
	grpc.ServerStream
}

type controlReportTaskUpdateServer struct {
	grpc.ServerStream
}

func (x *controlReportTaskUpdateServer) SendAndClose(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlReportTaskUpdateServer) Recv() (*TaskUpdate, error) {
	m := new(TaskUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Control_ServiceDesc is the grpc.ServiceDesc for Control service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Control_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLine",
			Handler:    _Control_GetLine_Handler,
		},
		{
			MethodName: "GetTaskByName",
			Handler:    _Control_GetTaskByName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToOrders",
			Handler:       _Control_SubscribeToOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReportTelemetry",
			Handler:       _Control_ReportTelemetry_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReportTaskUpdate",
			Handler:       _Control_ReportTaskUpdate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "control.proto",
}
