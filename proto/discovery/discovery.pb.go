// Code generated by protoc-gen-go.
// source: github.com/micro/discovery-srv/proto/discovery/discovery.proto
// DO NOT EDIT!

/*
Package go_micro_srv_discovery_discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/discovery-srv/proto/discovery/discovery.proto

It has these top-level messages:
	ServiceEndpoint
	EndpointsRequest
	EndpointsResponse
	HeartbeatsRequest
	HeartbeatsResponse
	WatchResultsRequest
	WatchResultsResponse
*/
package go_micro_srv_discovery_discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_os_discovery "github.com/micro/go-os/discovery/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type ServiceEndpoint struct {
	Service  string                          `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Version  string                          `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Endpoint *go_micro_os_discovery.Endpoint `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *ServiceEndpoint) Reset()                    { *m = ServiceEndpoint{} }
func (m *ServiceEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ServiceEndpoint) ProtoMessage()               {}
func (*ServiceEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceEndpoint) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceEndpoint) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceEndpoint) GetEndpoint() *go_micro_os_discovery.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type EndpointsRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Limit   uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *EndpointsRequest) Reset()                    { *m = EndpointsRequest{} }
func (m *EndpointsRequest) String() string            { return proto.CompactTextString(m) }
func (*EndpointsRequest) ProtoMessage()               {}
func (*EndpointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EndpointsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *EndpointsRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *EndpointsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *EndpointsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type EndpointsResponse struct {
	Endpoints []*ServiceEndpoint `protobuf:"bytes,1,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *EndpointsResponse) Reset()                    { *m = EndpointsResponse{} }
func (m *EndpointsResponse) String() string            { return proto.CompactTextString(m) }
func (*EndpointsResponse) ProtoMessage()               {}
func (*EndpointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EndpointsResponse) GetEndpoints() []*ServiceEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type HeartbeatsRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	After  uint64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
	Limit  uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *HeartbeatsRequest) Reset()                    { *m = HeartbeatsRequest{} }
func (m *HeartbeatsRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsRequest) ProtoMessage()               {}
func (*HeartbeatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HeartbeatsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HeartbeatsRequest) GetAfter() uint64 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *HeartbeatsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *HeartbeatsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type HeartbeatsResponse struct {
	Heartbeats []*go_micro_os_discovery.Heartbeat `protobuf:"bytes,1,rep,name=heartbeats" json:"heartbeats,omitempty"`
}

func (m *HeartbeatsResponse) Reset()                    { *m = HeartbeatsResponse{} }
func (m *HeartbeatsResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsResponse) ProtoMessage()               {}
func (*HeartbeatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HeartbeatsResponse) GetHeartbeats() []*go_micro_os_discovery.Heartbeat {
	if m != nil {
		return m.Heartbeats
	}
	return nil
}

type WatchResultsRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	After   uint64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
	Limit   uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *WatchResultsRequest) Reset()                    { *m = WatchResultsRequest{} }
func (m *WatchResultsRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchResultsRequest) ProtoMessage()               {}
func (*WatchResultsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WatchResultsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *WatchResultsRequest) GetAfter() uint64 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *WatchResultsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *WatchResultsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type WatchResultsResponse struct {
	Results []*go_micro_os_discovery.Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *WatchResultsResponse) Reset()                    { *m = WatchResultsResponse{} }
func (m *WatchResultsResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResultsResponse) ProtoMessage()               {}
func (*WatchResultsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WatchResultsResponse) GetResults() []*go_micro_os_discovery.Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceEndpoint)(nil), "go.micro.srv.discovery.discovery.ServiceEndpoint")
	proto.RegisterType((*EndpointsRequest)(nil), "go.micro.srv.discovery.discovery.EndpointsRequest")
	proto.RegisterType((*EndpointsResponse)(nil), "go.micro.srv.discovery.discovery.EndpointsResponse")
	proto.RegisterType((*HeartbeatsRequest)(nil), "go.micro.srv.discovery.discovery.HeartbeatsRequest")
	proto.RegisterType((*HeartbeatsResponse)(nil), "go.micro.srv.discovery.discovery.HeartbeatsResponse")
	proto.RegisterType((*WatchResultsRequest)(nil), "go.micro.srv.discovery.discovery.WatchResultsRequest")
	proto.RegisterType((*WatchResultsResponse)(nil), "go.micro.srv.discovery.discovery.WatchResultsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Publisher API

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

type publisher struct {
	c     client.Client
	topic string
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	return p.c.Publish(ctx, p.c.NewPublication(p.topic, msg), opts...)
}

func NewPublisher(topic string, c client.Client) Publisher {
	if c == nil {
		c = client.NewClient()
	}
	return &publisher{c, topic}
}

// Subscriber API

func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}

// Client API for Discovery service

type DiscoveryClient interface {
	Endpoints(ctx context.Context, in *EndpointsRequest, opts ...client.CallOption) (*EndpointsResponse, error)
	Heartbeats(ctx context.Context, in *HeartbeatsRequest, opts ...client.CallOption) (*HeartbeatsResponse, error)
	WatchResults(ctx context.Context, in *WatchResultsRequest, opts ...client.CallOption) (*WatchResultsResponse, error)
}

type discoveryClient struct {
	c           client.Client
	serviceName string
}

func NewDiscoveryClient(serviceName string, c client.Client) DiscoveryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.discovery.discovery"
	}
	return &discoveryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *discoveryClient) Endpoints(ctx context.Context, in *EndpointsRequest, opts ...client.CallOption) (*EndpointsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.Endpoints", in)
	out := new(EndpointsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Heartbeats(ctx context.Context, in *HeartbeatsRequest, opts ...client.CallOption) (*HeartbeatsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.Heartbeats", in)
	out := new(HeartbeatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) WatchResults(ctx context.Context, in *WatchResultsRequest, opts ...client.CallOption) (*WatchResultsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.WatchResults", in)
	out := new(WatchResultsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryHandler interface {
	Endpoints(context.Context, *EndpointsRequest, *EndpointsResponse) error
	Heartbeats(context.Context, *HeartbeatsRequest, *HeartbeatsResponse) error
	WatchResults(context.Context, *WatchResultsRequest, *WatchResultsResponse) error
}

func RegisterDiscoveryHandler(s server.Server, hdlr DiscoveryHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Discovery{hdlr}, opts...))
}

type Discovery struct {
	DiscoveryHandler
}

func (h *Discovery) Endpoints(ctx context.Context, in *EndpointsRequest, out *EndpointsResponse) error {
	return h.DiscoveryHandler.Endpoints(ctx, in, out)
}

func (h *Discovery) Heartbeats(ctx context.Context, in *HeartbeatsRequest, out *HeartbeatsResponse) error {
	return h.DiscoveryHandler.Heartbeats(ctx, in, out)
}

func (h *Discovery) WatchResults(ctx context.Context, in *WatchResultsRequest, out *WatchResultsResponse) error {
	return h.DiscoveryHandler.WatchResults(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/micro/discovery-srv/proto/discovery/discovery.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xdf, 0xcb, 0xd3, 0x30,
	0x14, 0xb5, 0xdd, 0xdc, 0xec, 0x9d, 0xa8, 0x8b, 0x43, 0x4a, 0x41, 0x2c, 0x7d, 0xda, 0xcb, 0x32,
	0xec, 0x74, 0x3e, 0x08, 0xe2, 0x83, 0x82, 0x6f, 0x83, 0x08, 0xfa, 0xdc, 0xb5, 0x59, 0x17, 0xd8,
	0x9a, 0x99, 0x64, 0x85, 0x3d, 0x09, 0xfe, 0xe3, 0xca, 0x9a, 0xf4, 0x87, 0xdb, 0x37, 0xb6, 0x7e,
	0x7c, 0x6f, 0xb9, 0xf7, 0xf6, 0xdc, 0x73, 0x4e, 0x4e, 0x28, 0x7c, 0x4a, 0x99, 0x5a, 0xef, 0x97,
	0x38, 0xe6, 0xdb, 0xe9, 0x96, 0xc5, 0x82, 0x4f, 0x13, 0x26, 0x63, 0x9e, 0x53, 0x71, 0x98, 0x48,
	0x91, 0x4f, 0x77, 0x82, 0xab, 0x46, 0xaf, 0x3e, 0xe1, 0x62, 0x82, 0xfc, 0x94, 0xe3, 0x02, 0x87,
	0xa5, 0xc8, 0x71, 0x3d, 0xad, 0x4e, 0xde, 0xfc, 0x8c, 0x21, 0xe5, 0x13, 0x2e, 0x1b, 0x3b, 0x4f,
	0x38, 0xf4, 0xe6, 0xe0, 0x8f, 0x05, 0xcf, 0xbf, 0x53, 0x91, 0xb3, 0x98, 0x7e, 0xcd, 0x92, 0x1d,
	0x67, 0x99, 0x42, 0x2e, 0xf4, 0xa5, 0x6e, 0xb9, 0x96, 0x6f, 0x8d, 0x1d, 0x52, 0x96, 0xc7, 0x49,
	0x4e, 0x85, 0x64, 0x3c, 0x73, 0x6d, 0x3d, 0x31, 0x25, 0xfa, 0x08, 0x4f, 0xa8, 0xc1, 0xbb, 0x1d,
	0xdf, 0x1a, 0x0f, 0xc2, 0x37, 0xb8, 0x12, 0xcd, 0x65, 0x43, 0x73, 0x49, 0x43, 0x2a, 0x40, 0xa0,
	0xe0, 0x45, 0xd9, 0x95, 0x84, 0xfe, 0xda, 0x53, 0x79, 0x3f, 0x11, 0x23, 0x78, 0xbc, 0x61, 0x5b,
	0xa6, 0x15, 0x74, 0x89, 0x2e, 0xd0, 0x2b, 0xe8, 0xf1, 0xd5, 0x4a, 0x52, 0xe5, 0x76, 0x8b, 0xb6,
	0xa9, 0x82, 0x04, 0x86, 0x0d, 0x56, 0xb9, 0xe3, 0x99, 0xa4, 0x68, 0x01, 0x4e, 0x29, 0x4b, 0xba,
	0x96, 0xdf, 0x19, 0x0f, 0xc2, 0xb7, 0xf8, 0xda, 0xed, 0xe3, 0x93, 0x1b, 0x24, 0xf5, 0x8e, 0x20,
	0x85, 0xe1, 0x37, 0x1a, 0x09, 0xb5, 0xa4, 0x51, 0x6d, 0xee, 0x19, 0xd8, 0x2c, 0x31, 0xbe, 0x6c,
	0x96, 0x1c, 0x85, 0x47, 0x2b, 0x45, 0x45, 0x61, 0xa8, 0x4b, 0x74, 0xd1, 0xd2, 0xce, 0x0f, 0x40,
	0x4d, 0x22, 0xe3, 0xe7, 0x33, 0xc0, 0xba, 0xea, 0x1a, 0x43, 0xfe, 0x85, 0x64, 0x2a, 0x38, 0x69,
	0x60, 0x02, 0x09, 0x2f, 0x7f, 0x46, 0x2a, 0x5e, 0x13, 0x2a, 0xf7, 0x9b, 0x5b, 0xf2, 0x79, 0x08,
	0x33, 0x0b, 0x18, 0xfd, 0x4f, 0x6a, 0xec, 0x7c, 0x80, 0xbe, 0xd0, 0x2d, 0xe3, 0xe5, 0xf5, 0x05,
	0x2f, 0x1a, 0x48, 0xca, 0xaf, 0xc3, 0xbf, 0x36, 0x38, 0x5f, 0xca, 0x29, 0xca, 0xc1, 0xa9, 0xa2,
	0x47, 0xe1, 0xf5, 0x7c, 0x4f, 0x5f, 0xa7, 0x37, 0x6b, 0x85, 0xd1, 0xe2, 0x83, 0x47, 0xe8, 0x00,
	0x50, 0x67, 0x84, 0x6e, 0x58, 0x72, 0xf6, 0x74, 0xbc, 0x77, 0xed, 0x40, 0x15, 0xf5, 0x6f, 0x78,
	0xda, 0xbc, 0x51, 0xf4, 0xfe, 0xfa, 0x9e, 0x3b, 0x62, 0xf7, 0xe6, 0x6d, 0x61, 0xa5, 0x80, 0x65,
	0xaf, 0xf8, 0xe1, 0xcc, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x40, 0x3f, 0x27, 0x7e, 0x0c, 0x05,
	0x00, 0x00,
}
