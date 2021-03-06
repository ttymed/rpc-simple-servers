// Code generated by protoc-gen-go.
// source: route_guide.proto
// DO NOT EDIT!

/*
Package rpcTest is a generated protocol buffer package.

It is generated from these files:
	route_guide.proto

It has these top-level messages:
	Point
	Rectangle
	Feature
	RouteNote
	RouteSummary
*/
package rpcTest

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Rectangle struct {
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

type Feature struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type RouteNote struct {
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RouteNote) Reset()                    { *m = RouteNote{} }
func (m *RouteNote) String() string            { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()               {}
func (*RouteNote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type RouteSummary struct {
	PointCount   int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	Distance     int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	ElapsedTime  int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
}

func (m *RouteSummary) Reset()                    { *m = RouteSummary{} }
func (m *RouteSummary) String() string            { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()               {}
func (*RouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Point)(nil), "rpcTest.Point")
	proto.RegisterType((*Rectangle)(nil), "rpcTest.Rectangle")
	proto.RegisterType((*Feature)(nil), "rpcTest.Feature")
	proto.RegisterType((*RouteNote)(nil), "rpcTest.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "rpcTest.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RouteGuide service

type RouteGuideClient interface {
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := grpc.Invoke(ctx, "/rpcTest.RouteGuide/GetFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[0], c.cc, "/rpcTest.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[1], c.cc, "/rpcTest.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[2], c.cc, "/rpcTest.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	GetFeature(context.Context, *Point) (*Feature, error)
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	RecordRoute(RouteGuide_RecordRouteServer) error
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcTest.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcTest.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0xd6, 0xd6, 0x9a, 0x49, 0x14, 0x1d, 0x10, 0x4a, 0x11, 0x3f, 0xe2, 0xa5, 0x78, 0x28,
	0xa5, 0x5e, 0xf4, 0x28, 0x05, 0x8b, 0x28, 0xa2, 0xb1, 0xf7, 0xb2, 0x26, 0x6b, 0x1a, 0x48, 0xb2,
	0x21, 0xd9, 0x1c, 0xfc, 0x1d, 0xfe, 0x4e, 0xff, 0x83, 0x9b, 0xdd, 0x4d, 0x5a, 0xb5, 0x7a, 0xcb,
	0xbe, 0x37, 0x6f, 0xde, 0xcc, 0x9b, 0xc0, 0x41, 0xce, 0x4b, 0xc1, 0x16, 0x61, 0x19, 0x05, 0x6c,
	0x94, 0xe5, 0x5c, 0x70, 0xec, 0xe5, 0x99, 0x3f, 0x67, 0x85, 0x70, 0x6f, 0xa0, 0xfb, 0xc4, 0xa3,
	0x54, 0xe0, 0x00, 0x76, 0x62, 0x2a, 0x22, 0x51, 0x06, 0xac, 0x4f, 0x4e, 0xc9, 0xb0, 0xeb, 0x35,
	0x6f, 0x3c, 0x02, 0x2b, 0xe6, 0x69, 0xa8, 0xc9, 0xb6, 0x22, 0x57, 0x80, 0x7b, 0x0f, 0x96, 0xc7,
	0x7c, 0x41, 0xd3, 0x30, 0x66, 0x78, 0x0c, 0xed, 0x98, 0xab, 0x06, 0xf6, 0x64, 0x6f, 0x64, 0x5c,
	0x46, 0xca, 0xc2, 0x93, 0x4c, 0xc5, 0x2f, 0x23, 0xd5, 0x63, 0x03, 0xbf, 0x8c, 0xdc, 0x3b, 0xe8,
	0xdd, 0x32, 0x2a, 0xca, 0x9c, 0x21, 0x42, 0x27, 0xa5, 0x89, 0x9e, 0xc6, 0xf2, 0xd4, 0x37, 0x5e,
	0xc8, 0x29, 0xb9, 0x2f, 0xe7, 0xe2, 0xe9, 0x1f, 0x4d, 0x1a, 0xde, 0x7d, 0x96, 0x73, 0x55, 0x8b,
	0x3f, 0x72, 0xf1, 0x5d, 0x48, 0xfe, 0x17, 0x62, 0x1f, 0x7a, 0x09, 0x2b, 0x0a, 0x1a, 0xea, 0x65,
	0x2d, 0xaf, 0x7e, 0xba, 0x1f, 0x04, 0x1c, 0xd5, 0xf3, 0xa5, 0x4c, 0x12, 0x9a, 0xbf, 0xe3, 0x09,
	0xd8, 0x59, 0xa5, 0x5e, 0xf8, 0xbc, 0x4c, 0x85, 0x09, 0x0e, 0x14, 0x34, 0xad, 0x10, 0x3c, 0x87,
	0xdd, 0x37, 0xbd, 0x8f, 0x29, 0xd1, 0xf1, 0x39, 0x06, 0xd4, 0x45, 0x32, 0xfb, 0x20, 0x2a, 0x64,
	0x82, 0x3e, 0xeb, 0x6f, 0xe9, 0xec, 0xeb, 0x37, 0x9e, 0x81, 0xc3, 0x62, 0x9a, 0x15, 0x2c, 0x58,
	0x88, 0x48, 0xa6, 0xd1, 0x51, 0xbc, 0x6d, 0xb0, 0xb9, 0x84, 0x26, 0x9f, 0x04, 0x40, 0x4d, 0x35,
	0xab, 0x2e, 0x8c, 0x63, 0x80, 0x19, 0x13, 0x75, 0x8a, 0x3f, 0xd6, 0x1c, 0xec, 0x37, 0x6f, 0x53,
	0xe1, 0xb6, 0xf0, 0x0a, 0x9c, 0x07, 0xe9, 0x67, 0x80, 0x02, 0xb1, 0xa9, 0x69, 0x0e, 0xbb, 0x49,
	0x37, 0x26, 0x52, 0x69, 0xcb, 0x12, 0x9e, 0x07, 0xca, 0xff, 0x97, 0xd9, 0xe1, 0xaa, 0xd1, 0x5a,
	0x6a, 0x6e, 0x6b, 0x48, 0xf0, 0xda, 0x5c, 0x67, 0xba, 0xa4, 0x62, 0xdd, 0xb0, 0xbe, 0xd8, 0x60,
	0x03, 0x56, 0x09, 0xc7, 0xe4, 0x75, 0x5b, 0xfd, 0xc3, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x89, 0x58, 0x83, 0xcf, 0xd8, 0x02, 0x00, 0x00,
}
