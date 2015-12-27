// Code generated by protoc-gen-go.
// source: github.com/micro/discovery-srv/proto/discovery/discovery.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/discovery-srv/proto/discovery/discovery.proto

It has these top-level messages:
	HeartbeatsRequest
	HeartbeatsResponse
*/
package discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import discovery1 "github.com/micro/go-platform/discovery/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HeartbeatsRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	After  uint64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
	Limit  uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *HeartbeatsRequest) Reset()                    { *m = HeartbeatsRequest{} }
func (m *HeartbeatsRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsRequest) ProtoMessage()               {}
func (*HeartbeatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeartbeatsResponse struct {
	Heartbeats []*discovery1.Heartbeat `protobuf:"bytes,1,rep,name=heartbeats" json:"heartbeats,omitempty"`
}

func (m *HeartbeatsResponse) Reset()                    { *m = HeartbeatsResponse{} }
func (m *HeartbeatsResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsResponse) ProtoMessage()               {}
func (*HeartbeatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HeartbeatsResponse) GetHeartbeats() []*discovery1.Heartbeat {
	if m != nil {
		return m.Heartbeats
	}
	return nil
}

func init() {
	proto.RegisterType((*HeartbeatsRequest)(nil), "HeartbeatsRequest")
	proto.RegisterType((*HeartbeatsResponse)(nil), "HeartbeatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Discovery service

type DiscoveryClient interface {
	Heartbeats(ctx context.Context, in *HeartbeatsRequest, opts ...client.CallOption) (*HeartbeatsResponse, error)
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
		serviceName = "discovery"
	}
	return &discoveryClient{
		c:           c,
		serviceName: serviceName,
	}
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

// Server API for Discovery service

type DiscoveryHandler interface {
	Heartbeats(context.Context, *HeartbeatsRequest, *HeartbeatsResponse) error
}

func RegisterDiscoveryHandler(s server.Server, hdlr DiscoveryHandler) {
	s.Handle(s.NewHandler(&Discovery{hdlr}))
}

type Discovery struct {
	DiscoveryHandler
}

func (h *Discovery) Heartbeats(ctx context.Context, in *HeartbeatsRequest, out *HeartbeatsResponse) error {
	return h.DiscoveryHandler.Heartbeats(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xc9, 0x2c,
	0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0x41,
	0x12, 0x43, 0xb0, 0xf4, 0xc0, 0x32, 0x52, 0x36, 0x18, 0xfa, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12,
	0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0x91, 0xf4, 0xa1, 0x99, 0x03, 0xd1, 0xad, 0xe4, 0xcb, 0x25, 0xe8,
	0x91, 0x9a, 0x58, 0x54, 0x92, 0x94, 0x9a, 0x58, 0x52, 0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0xc4, 0xc5, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x29, 0xc4, 0xcb, 0xc5,
	0x9a, 0x98, 0x56, 0x92, 0x5a, 0x24, 0xc1, 0x04, 0xe4, 0xb2, 0x80, 0xb8, 0x39, 0x99, 0xb9, 0x99,
	0x25, 0x12, 0xcc, 0x60, 0x2e, 0x1f, 0x17, 0x5b, 0x7e, 0x5a, 0x5a, 0x71, 0x6a, 0x89, 0x04, 0x0b,
	0x88, 0xaf, 0x64, 0xc2, 0x25, 0x84, 0x6c, 0x5c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1c,
	0x17, 0x57, 0x06, 0x5c, 0x14, 0x68, 0x2e, 0xb3, 0x06, 0xb7, 0x11, 0x97, 0x1e, 0x5c, 0xa1, 0x91,
	0x0b, 0x17, 0xa7, 0x0b, 0xcc, 0x5d, 0x42, 0xe6, 0x5c, 0x5c, 0x08, 0x23, 0x84, 0x84, 0xf4, 0x30,
	0x9c, 0x27, 0x25, 0xac, 0x87, 0x69, 0x87, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x47, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x68, 0xbe, 0x79, 0xce, 0x51, 0x01, 0x00, 0x00,
}