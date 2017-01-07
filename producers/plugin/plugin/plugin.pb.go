// Code generated by protoc-gen-go.
// source: plugin.proto
// DO NOT EDIT!

/*
Package plugin is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	MetricsCollectorType
	MetricsMessage
	Datapoint
	Dimensions
*/
package plugin

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

// Specifies the type of metrics message to recieve. This is based on the
// collectors available. Currently we have `node` and `framework` collectors.
type MetricsCollectorType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MetricsCollectorType) Reset()                    { *m = MetricsCollectorType{} }
func (m *MetricsCollectorType) String() string            { return proto.CompactTextString(m) }
func (*MetricsCollectorType) ProtoMessage()               {}
func (*MetricsCollectorType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricsCollectorType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type MetricsMessage struct {
	Name       string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Datapoints []*Datapoint `protobuf:"bytes,2,rep,name=datapoints" json:"datapoints,omitempty"`
	Dimensions *Dimensions  `protobuf:"bytes,3,opt,name=dimensions" json:"dimensions,omitempty"`
	Timestamp  int64        `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *MetricsMessage) Reset()                    { *m = MetricsMessage{} }
func (m *MetricsMessage) String() string            { return proto.CompactTextString(m) }
func (*MetricsMessage) ProtoMessage()               {}
func (*MetricsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetricsMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsMessage) GetDatapoints() []*Datapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

func (m *MetricsMessage) GetDimensions() *Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *MetricsMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Datapoint struct {
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *Datapoint) Reset()                    { *m = Datapoint{} }
func (m *Datapoint) String() string            { return proto.CompactTextString(m) }
func (*Datapoint) ProtoMessage()               {}
func (*Datapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Datapoint) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Datapoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Datapoint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Dimensions struct {
	MesosID            string            `protobuf:"bytes,1,opt,name=mesosID" json:"mesosID,omitempty"`
	ClusterID          string            `protobuf:"bytes,2,opt,name=clusterID" json:"clusterID,omitempty"`
	ContainerID        string            `protobuf:"bytes,3,opt,name=containerID" json:"containerID,omitempty"`
	ExecutorID         string            `protobuf:"bytes,4,opt,name=executorID" json:"executorID,omitempty"`
	FrameworkName      string            `protobuf:"bytes,5,opt,name=frameworkName" json:"frameworkName,omitempty"`
	FrameworkID        string            `protobuf:"bytes,6,opt,name=frameworkID" json:"frameworkID,omitempty"`
	FrameworkRole      string            `protobuf:"bytes,7,opt,name=frameworkRole" json:"frameworkRole,omitempty"`
	FrameworkPrincipal string            `protobuf:"bytes,8,opt,name=frameworkPrincipal" json:"frameworkPrincipal,omitempty"`
	Hostname           string            `protobuf:"bytes,9,opt,name=hostname" json:"hostname,omitempty"`
	Labels             map[string]string `protobuf:"bytes,10,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Dimensions) Reset()                    { *m = Dimensions{} }
func (m *Dimensions) String() string            { return proto.CompactTextString(m) }
func (*Dimensions) ProtoMessage()               {}
func (*Dimensions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Dimensions) GetMesosID() string {
	if m != nil {
		return m.MesosID
	}
	return ""
}

func (m *Dimensions) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *Dimensions) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *Dimensions) GetExecutorID() string {
	if m != nil {
		return m.ExecutorID
	}
	return ""
}

func (m *Dimensions) GetFrameworkName() string {
	if m != nil {
		return m.FrameworkName
	}
	return ""
}

func (m *Dimensions) GetFrameworkID() string {
	if m != nil {
		return m.FrameworkID
	}
	return ""
}

func (m *Dimensions) GetFrameworkRole() string {
	if m != nil {
		return m.FrameworkRole
	}
	return ""
}

func (m *Dimensions) GetFrameworkPrincipal() string {
	if m != nil {
		return m.FrameworkPrincipal
	}
	return ""
}

func (m *Dimensions) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Dimensions) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsCollectorType)(nil), "plugin.MetricsCollectorType")
	proto.RegisterType((*MetricsMessage)(nil), "plugin.MetricsMessage")
	proto.RegisterType((*Datapoint)(nil), "plugin.Datapoint")
	proto.RegisterType((*Dimensions)(nil), "plugin.Dimensions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metrics service

type MetricsClient interface {
	AttachOutputStream(ctx context.Context, in *MetricsCollectorType, opts ...grpc.CallOption) (*MetricsMessage, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) AttachOutputStream(ctx context.Context, in *MetricsCollectorType, opts ...grpc.CallOption) (*MetricsMessage, error) {
	out := new(MetricsMessage)
	err := grpc.Invoke(ctx, "/plugin.Metrics/AttachOutputStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metrics service

type MetricsServer interface {
	AttachOutputStream(context.Context, *MetricsCollectorType) (*MetricsMessage, error)
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_AttachOutputStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsCollectorType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).AttachOutputStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Metrics/AttachOutputStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).AttachOutputStream(ctx, req.(*MetricsCollectorType))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttachOutputStream",
			Handler:    _Metrics_AttachOutputStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x9c, 0x26, 0xf5, 0x04, 0x10, 0x8c, 0x2a, 0x64, 0x45, 0x55, 0x65, 0x59, 0x1c,
	0x22, 0x0e, 0x91, 0x08, 0x12, 0x02, 0x6e, 0x88, 0x70, 0x88, 0xd4, 0x02, 0x72, 0x91, 0x38, 0x6f,
	0xdd, 0xa1, 0x5d, 0x75, 0xbd, 0x6b, 0xed, 0x8e, 0x81, 0x3c, 0x12, 0x67, 0x5e, 0x10, 0x79, 0xed,
	0xd8, 0x4e, 0x9b, 0xdb, 0xce, 0x3f, 0x9f, 0xe7, 0x1f, 0xef, 0xcc, 0xc2, 0xe3, 0x52, 0x55, 0x37,
	0x52, 0x2f, 0x4b, 0x6b, 0xd8, 0xe0, 0xa4, 0x89, 0xd2, 0x57, 0x70, 0x72, 0x41, 0x6c, 0x65, 0xee,
	0x3e, 0x19, 0xa5, 0x28, 0x67, 0x63, 0xbf, 0x6f, 0x4b, 0x42, 0x84, 0x31, 0x6f, 0x4b, 0x8a, 0x83,
	0x24, 0x58, 0x44, 0x99, 0x3f, 0xa7, 0x7f, 0x03, 0x78, 0xda, 0xc2, 0x17, 0xe4, 0x9c, 0xb8, 0xf1,
	0x98, 0x16, 0x45, 0x87, 0xd5, 0x67, 0x7c, 0x0d, 0x70, 0x2d, 0x58, 0x94, 0x46, 0x6a, 0x76, 0xf1,
	0x28, 0x09, 0x17, 0xb3, 0xd5, 0xf3, 0x65, 0xeb, 0xbe, 0xde, 0x65, 0xb2, 0x01, 0x84, 0x2b, 0x80,
	0x6b, 0x59, 0x90, 0x76, 0xd2, 0x68, 0x17, 0x87, 0x49, 0xb0, 0x98, 0xad, 0xb0, 0xfb, 0xa4, 0xcb,
	0x64, 0x03, 0x0a, 0x4f, 0x21, 0x62, 0x59, 0x90, 0x63, 0x51, 0x94, 0xf1, 0x38, 0x09, 0x16, 0x61,
	0xd6, 0x0b, 0xe9, 0x25, 0x44, 0x9d, 0xd5, 0x3e, 0xda, 0xb4, 0xda, 0x0b, 0xdd, 0x3f, 0x8c, 0x06,
	0xff, 0x70, 0x02, 0x47, 0xbf, 0x84, 0xaa, 0xc8, 0xf7, 0x12, 0x65, 0x4d, 0x90, 0xfe, 0x0b, 0x01,
	0xfa, 0x6e, 0x30, 0x86, 0x69, 0x41, 0xce, 0xb8, 0xcd, 0xba, 0x2d, 0xba, 0x0b, 0x6b, 0xc3, 0x5c,
	0x55, 0x8e, 0xc9, 0x6e, 0xd6, 0x6d, 0xdd, 0x5e, 0xc0, 0x04, 0x66, 0xb9, 0xd1, 0x2c, 0xa4, 0xf6,
	0xf9, 0xc6, 0x62, 0x28, 0xe1, 0x19, 0x00, 0xfd, 0xa1, 0xbc, 0x62, 0x53, 0x03, 0x63, 0x0f, 0x0c,
	0x14, 0x7c, 0x09, 0x4f, 0x7e, 0x5a, 0x51, 0xd0, 0x6f, 0x63, 0xef, 0xbe, 0xd4, 0xbd, 0x1f, 0x79,
	0x64, 0x5f, 0xac, 0x7d, 0x3a, 0x61, 0xb3, 0x8e, 0x27, 0x8d, 0xcf, 0x40, 0xda, 0xab, 0x93, 0x19,
	0x45, 0xf1, 0xf4, 0x5e, 0x9d, 0x5a, 0xc4, 0x25, 0x60, 0x27, 0x7c, 0xb3, 0x52, 0xe7, 0xb2, 0x14,
	0x2a, 0x3e, 0xf6, 0xe8, 0x81, 0x0c, 0xce, 0xe1, 0xf8, 0xd6, 0x38, 0xf6, 0x97, 0x1a, 0x79, 0xaa,
	0x8b, 0xf1, 0x2d, 0x4c, 0x94, 0xb8, 0x22, 0xe5, 0x62, 0xf0, 0x8b, 0x71, 0xf6, 0x70, 0xca, 0xcb,
	0x73, 0x0f, 0x7c, 0xd6, 0x6c, 0xb7, 0x59, 0x4b, 0xcf, 0xdf, 0xc3, 0x6c, 0x20, 0xe3, 0x33, 0x08,
	0xef, 0x68, 0xdb, 0x5e, 0x7b, 0x7d, 0xec, 0x27, 0x36, 0x1a, 0x4c, 0xec, 0xc3, 0xe8, 0x5d, 0xb0,
	0xfa, 0x01, 0xd3, 0x76, 0x6b, 0xf1, 0x1c, 0xf0, 0x23, 0xb3, 0xc8, 0x6f, 0xbf, 0x56, 0x5c, 0x56,
	0x7c, 0xc9, 0x96, 0x44, 0x81, 0xa7, 0xbb, 0x1e, 0x0e, 0xbd, 0x84, 0xf9, 0x8b, 0x7b, 0xd9, 0x76,
	0xf5, 0xd3, 0x47, 0x57, 0x13, 0xff, 0x94, 0xde, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x36, 0x47,
	0x43, 0xbe, 0x5a, 0x03, 0x00, 0x00,
}
