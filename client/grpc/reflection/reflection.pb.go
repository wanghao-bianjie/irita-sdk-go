// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/reflection/v1beta1/reflection.proto

package reflection

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ListAllInterfacesRequest is the request type of the ListAllInterfaces RPC.
type ListAllInterfacesRequest struct {
}

func (m *ListAllInterfacesRequest) Reset()         { *m = ListAllInterfacesRequest{} }
func (m *ListAllInterfacesRequest) String() string { return proto.CompactTextString(m) }
func (*ListAllInterfacesRequest) ProtoMessage()    {}
func (*ListAllInterfacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48c054165687f5c, []int{0}
}
func (m *ListAllInterfacesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAllInterfacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAllInterfacesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAllInterfacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllInterfacesRequest.Merge(m, src)
}
func (m *ListAllInterfacesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListAllInterfacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllInterfacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllInterfacesRequest proto.InternalMessageInfo

// ListAllInterfacesResponse is the response type of the ListAllInterfaces RPC.
type ListAllInterfacesResponse struct {
	// interface_names is an array of all the registered interfaces.
	InterfaceNames []string `protobuf:"bytes,1,rep,name=interface_names,json=interfaceNames,proto3" json:"interface_names,omitempty"`
}

func (m *ListAllInterfacesResponse) Reset()         { *m = ListAllInterfacesResponse{} }
func (m *ListAllInterfacesResponse) String() string { return proto.CompactTextString(m) }
func (*ListAllInterfacesResponse) ProtoMessage()    {}
func (*ListAllInterfacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48c054165687f5c, []int{1}
}
func (m *ListAllInterfacesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAllInterfacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAllInterfacesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAllInterfacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllInterfacesResponse.Merge(m, src)
}
func (m *ListAllInterfacesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListAllInterfacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllInterfacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllInterfacesResponse proto.InternalMessageInfo

func (m *ListAllInterfacesResponse) GetInterfaceNames() []string {
	if m != nil {
		return m.InterfaceNames
	}
	return nil
}

// ListImplementationsRequest is the request type of the ListImplementations
// RPC.
type ListImplementationsRequest struct {
	// interface_name defines the interface to query the implementations for.
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (m *ListImplementationsRequest) Reset()         { *m = ListImplementationsRequest{} }
func (m *ListImplementationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListImplementationsRequest) ProtoMessage()    {}
func (*ListImplementationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48c054165687f5c, []int{2}
}
func (m *ListImplementationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListImplementationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListImplementationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListImplementationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImplementationsRequest.Merge(m, src)
}
func (m *ListImplementationsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListImplementationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImplementationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListImplementationsRequest proto.InternalMessageInfo

func (m *ListImplementationsRequest) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

// ListImplementationsResponse is the response type of the ListImplementations
// RPC.
type ListImplementationsResponse struct {
	ImplementationMessageNames []string `protobuf:"bytes,1,rep,name=implementation_message_names,json=implementationMessageNames,proto3" json:"implementation_message_names,omitempty"`
}

func (m *ListImplementationsResponse) Reset()         { *m = ListImplementationsResponse{} }
func (m *ListImplementationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListImplementationsResponse) ProtoMessage()    {}
func (*ListImplementationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48c054165687f5c, []int{3}
}
func (m *ListImplementationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListImplementationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListImplementationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListImplementationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImplementationsResponse.Merge(m, src)
}
func (m *ListImplementationsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListImplementationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImplementationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListImplementationsResponse proto.InternalMessageInfo

func (m *ListImplementationsResponse) GetImplementationMessageNames() []string {
	if m != nil {
		return m.ImplementationMessageNames
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAllInterfacesRequest)(nil), "cosmos.base.reflection.v1beta1.ListAllInterfacesRequest")
	proto.RegisterType((*ListAllInterfacesResponse)(nil), "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse")
	proto.RegisterType((*ListImplementationsRequest)(nil), "cosmos.base.reflection.v1beta1.ListImplementationsRequest")
	proto.RegisterType((*ListImplementationsResponse)(nil), "cosmos.base.reflection.v1beta1.ListImplementationsResponse")
}

func init() {
	proto.RegisterFile("cosmos/base/reflection/v1beta1/reflection.proto", fileDescriptor_d48c054165687f5c)
}

var fileDescriptor_d48c054165687f5c = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x6e, 0xdb, 0x30,
	0x18, 0x34, 0x6b, 0xa0, 0x80, 0x09, 0xd4, 0x85, 0xd9, 0xc5, 0x55, 0x0d, 0xc1, 0x10, 0x50, 0xd4,
	0x28, 0x6a, 0x11, 0x76, 0x17, 0xb7, 0x5d, 0xfa, 0xb7, 0x18, 0xad, 0x3b, 0xc8, 0x5b, 0x17, 0x83,
	0x52, 0x3f, 0xab, 0x6c, 0x25, 0x52, 0x11, 0x69, 0x2f, 0x41, 0x96, 0x3c, 0x41, 0x80, 0xbc, 0x4c,
	0x86, 0x3c, 0x40, 0x46, 0x03, 0x59, 0x32, 0x06, 0x56, 0x1e, 0x24, 0x90, 0xa5, 0x38, 0x16, 0xa2,
	0x24, 0x86, 0xd7, 0xef, 0x78, 0xc7, 0xbb, 0xe3, 0x47, 0x4c, 0x3d, 0xa9, 0x42, 0xa9, 0xa8, 0xcb,
	0x14, 0xd0, 0x18, 0xa6, 0x01, 0x78, 0x9a, 0x4b, 0x41, 0xe7, 0x3d, 0x17, 0x34, 0xeb, 0x6d, 0x8c,
	0xec, 0x28, 0x96, 0x5a, 0x12, 0x33, 0x23, 0xd8, 0x29, 0xc1, 0xde, 0x40, 0x73, 0x82, 0xd1, 0xf2,
	0xa5, 0xf4, 0x03, 0xa0, 0x2c, 0xe2, 0x94, 0x09, 0x21, 0x35, 0x4b, 0x61, 0x95, 0xb1, 0x2d, 0x03,
	0x37, 0x7f, 0x72, 0xa5, 0xbf, 0x04, 0xc1, 0x50, 0x68, 0x88, 0xa7, 0xcc, 0x03, 0xe5, 0xc0, 0xde,
	0x0c, 0x94, 0xb6, 0xbe, 0xe3, 0x97, 0x25, 0x98, 0x8a, 0xa4, 0x50, 0x40, 0xde, 0xe0, 0xe7, 0xfc,
	0x66, 0x3a, 0x11, 0x2c, 0x04, 0xd5, 0x44, 0xed, 0x6a, 0xa7, 0xe6, 0xd4, 0xd7, 0xe3, 0x5f, 0xe9,
	0xd4, 0xfa, 0x86, 0x8d, 0x54, 0x65, 0x18, 0x46, 0x01, 0x84, 0x20, 0xf2, 0xeb, 0xf3, 0x3b, 0xc8,
	0x6b, 0x5c, 0x2f, 0xca, 0x34, 0x51, 0x1b, 0x75, 0x6a, 0xce, 0xb3, 0x82, 0x8a, 0x35, 0xc1, 0xaf,
	0x4a, 0x45, 0x72, 0x33, 0x9f, 0x71, 0x8b, 0x17, 0xa0, 0x49, 0x08, 0x4a, 0x31, 0xbf, 0xe8, 0xcc,
	0x28, 0x9e, 0x19, 0x65, 0x47, 0x56, 0x2e, 0xfb, 0x27, 0x55, 0xdc, 0x70, 0xd6, 0xe5, 0x8d, 0x21,
	0x9e, 0x73, 0x0f, 0xc8, 0x29, 0xc2, 0x8d, 0x3b, 0x15, 0x90, 0x81, 0xfd, 0x70, 0xe5, 0xf6, 0x7d,
	0x8d, 0x1a, 0x1f, 0x76, 0x60, 0x66, 0x11, 0xad, 0xfe, 0xe1, 0xf9, 0xd5, 0xf1, 0x93, 0x77, 0xe4,
	0xed, 0x63, 0x0b, 0xc2, 0x6f, 0x8d, 0x26, 0x08, 0xbf, 0x28, 0xa9, 0x8d, 0x7c, 0xdc, 0xc6, 0x46,
	0xf9, 0x83, 0x19, 0x9f, 0x76, 0xe2, 0xe6, 0x21, 0xc6, 0xab, 0x10, 0x23, 0xf2, 0x63, 0xfb, 0x10,
	0x74, 0xbf, 0xb8, 0x1f, 0x07, 0xb4, 0xf8, 0x8a, 0xea, 0xab, 0x73, 0xb6, 0x34, 0xd1, 0x62, 0x69,
	0xa2, 0xcb, 0xa5, 0x89, 0x8e, 0x12, 0xb3, 0xb2, 0x48, 0xcc, 0xca, 0x45, 0x62, 0x56, 0x7e, 0x0f,
	0x7c, 0xae, 0xff, 0xce, 0x5c, 0xdb, 0x93, 0x21, 0x75, 0x39, 0x13, 0xff, 0x38, 0x30, 0x4e, 0x79,
	0xcc, 0x35, 0xeb, 0xaa, 0x3f, 0xff, 0xbb, 0xbe, 0xa4, 0x5e, 0xc0, 0x41, 0x68, 0xea, 0xc7, 0x91,
	0xb7, 0xe1, 0xc3, 0x7d, 0xba, 0xfa, 0x1d, 0xef, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf1,
	0x52, 0x13, 0x8e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReflectionServiceClient interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error)
}

type reflectionServiceClient struct {
	cc grpc1.ClientConn
}

func NewReflectionServiceClient(cc grpc1.ClientConn) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error) {
	out := new(ListAllInterfacesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v1beta1.ReflectionService/ListAllInterfaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error) {
	out := new(ListImplementationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v1beta1.ReflectionService/ListImplementations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
type ReflectionServiceServer interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(context.Context, *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(context.Context, *ListImplementationsRequest) (*ListImplementationsResponse, error)
}

// UnimplementedReflectionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReflectionServiceServer struct {
}

func (*UnimplementedReflectionServiceServer) ListAllInterfaces(ctx context.Context, req *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllInterfaces not implemented")
}
func (*UnimplementedReflectionServiceServer) ListImplementations(ctx context.Context, req *ListImplementationsRequest) (*ListImplementationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImplementations not implemented")
}

func RegisterReflectionServiceServer(s grpc1.Server, srv ReflectionServiceServer) {
	s.RegisterService(&_ReflectionService_serviceDesc, srv)
}

func _ReflectionService_ListAllInterfaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllInterfacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v1beta1.ReflectionService/ListAllInterfaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, req.(*ListAllInterfacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_ListImplementations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImplementationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v1beta1.ReflectionService/ListImplementations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, req.(*ListImplementationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReflectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v1beta1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllInterfaces",
			Handler:    _ReflectionService_ListAllInterfaces_Handler,
		},
		{
			MethodName: "ListImplementations",
			Handler:    _ReflectionService_ListImplementations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/reflection/v1beta1/reflection.proto",
}

func (m *ListAllInterfacesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAllInterfacesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAllInterfacesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListAllInterfacesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAllInterfacesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAllInterfacesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterfaceNames) > 0 {
		for iNdEx := len(m.InterfaceNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InterfaceNames[iNdEx])
			copy(dAtA[i:], m.InterfaceNames[iNdEx])
			i = encodeVarintReflection(dAtA, i, uint64(len(m.InterfaceNames[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListImplementationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListImplementationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListImplementationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterfaceName) > 0 {
		i -= len(m.InterfaceName)
		copy(dAtA[i:], m.InterfaceName)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.InterfaceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListImplementationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListImplementationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListImplementationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ImplementationMessageNames) > 0 {
		for iNdEx := len(m.ImplementationMessageNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ImplementationMessageNames[iNdEx])
			copy(dAtA[i:], m.ImplementationMessageNames[iNdEx])
			i = encodeVarintReflection(dAtA, i, uint64(len(m.ImplementationMessageNames[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintReflection(dAtA []byte, offset int, v uint64) int {
	offset -= sovReflection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListAllInterfacesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListAllInterfacesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.InterfaceNames) > 0 {
		for _, s := range m.InterfaceNames {
			l = len(s)
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *ListImplementationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterfaceName)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *ListImplementationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ImplementationMessageNames) > 0 {
		for _, s := range m.ImplementationMessageNames {
			l = len(s)
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func sovReflection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReflection(x uint64) (n int) {
	return sovReflection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListAllInterfacesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAllInterfacesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAllInterfacesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListAllInterfacesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAllInterfacesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAllInterfacesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceNames = append(m.InterfaceNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListImplementationsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListImplementationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListImplementationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListImplementationsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListImplementationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListImplementationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImplementationMessageNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImplementationMessageNames = append(m.ImplementationMessageNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReflection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReflection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReflection
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthReflection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReflection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReflection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReflection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReflection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReflection = fmt.Errorf("proto: unexpected end of group")
)
