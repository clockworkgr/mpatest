// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mpa/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// this line is used by starport scaffolding # 3
type QueryGetPositionRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetPositionRequest) Reset()         { *m = QueryGetPositionRequest{} }
func (m *QueryGetPositionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetPositionRequest) ProtoMessage()    {}
func (*QueryGetPositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc97cf48d3cc800, []int{0}
}
func (m *QueryGetPositionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPositionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPositionRequest.Merge(m, src)
}
func (m *QueryGetPositionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPositionRequest proto.InternalMessageInfo

func (m *QueryGetPositionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetPositionResponse struct {
	Position *Position `protobuf:"bytes,1,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (m *QueryGetPositionResponse) Reset()         { *m = QueryGetPositionResponse{} }
func (m *QueryGetPositionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetPositionResponse) ProtoMessage()    {}
func (*QueryGetPositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc97cf48d3cc800, []int{1}
}
func (m *QueryGetPositionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPositionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPositionResponse.Merge(m, src)
}
func (m *QueryGetPositionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPositionResponse proto.InternalMessageInfo

func (m *QueryGetPositionResponse) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type QueryAllPositionRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPositionRequest) Reset()         { *m = QueryAllPositionRequest{} }
func (m *QueryAllPositionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPositionRequest) ProtoMessage()    {}
func (*QueryAllPositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc97cf48d3cc800, []int{2}
}
func (m *QueryAllPositionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPositionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPositionRequest.Merge(m, src)
}
func (m *QueryAllPositionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPositionRequest proto.InternalMessageInfo

func (m *QueryAllPositionRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPositionResponse struct {
	Position   []*Position         `protobuf:"bytes,1,rep,name=Position,proto3" json:"Position,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPositionResponse) Reset()         { *m = QueryAllPositionResponse{} }
func (m *QueryAllPositionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPositionResponse) ProtoMessage()    {}
func (*QueryAllPositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc97cf48d3cc800, []int{3}
}
func (m *QueryAllPositionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPositionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPositionResponse.Merge(m, src)
}
func (m *QueryAllPositionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPositionResponse proto.InternalMessageInfo

func (m *QueryAllPositionResponse) GetPosition() []*Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *QueryAllPositionResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetPositionRequest)(nil), "chain.mpatest.mpa.QueryGetPositionRequest")
	proto.RegisterType((*QueryGetPositionResponse)(nil), "chain.mpatest.mpa.QueryGetPositionResponse")
	proto.RegisterType((*QueryAllPositionRequest)(nil), "chain.mpatest.mpa.QueryAllPositionRequest")
	proto.RegisterType((*QueryAllPositionResponse)(nil), "chain.mpatest.mpa.QueryAllPositionResponse")
}

func init() { proto.RegisterFile("mpa/query.proto", fileDescriptor_7dc97cf48d3cc800) }

var fileDescriptor_7dc97cf48d3cc800 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0x23, 0x31,
	0x1c, 0xc5, 0x9b, 0xd9, 0x1f, 0x2c, 0x29, 0xec, 0xb2, 0xb9, 0x6c, 0x69, 0x77, 0x87, 0x32, 0xbb,
	0xec, 0x76, 0x2b, 0x24, 0xb4, 0x1e, 0xbc, 0x09, 0xf5, 0x60, 0xaf, 0xb5, 0xde, 0xbc, 0x65, 0xda,
	0x30, 0x0d, 0xcc, 0x4c, 0xd2, 0x26, 0x15, 0x8b, 0x78, 0xf1, 0x2a, 0x82, 0xe0, 0xd5, 0xab, 0xff,
	0x8b, 0xc7, 0x82, 0x17, 0x8f, 0xd2, 0xfa, 0x87, 0xc8, 0x64, 0xd2, 0xdf, 0xad, 0xf6, 0x34, 0x90,
	0x79, 0xef, 0xfb, 0x3e, 0x2f, 0xdf, 0xc0, 0x6f, 0x91, 0xa4, 0xa4, 0xdb, 0x67, 0xbd, 0x01, 0x96,
	0x3d, 0xa1, 0x05, 0xfa, 0xde, 0xea, 0x50, 0x1e, 0xe3, 0x48, 0x52, 0xcd, 0x94, 0x4e, 0xbe, 0xf9,
	0x9f, 0x81, 0x10, 0x41, 0xc8, 0x08, 0x95, 0x9c, 0xd0, 0x38, 0x16, 0x9a, 0x6a, 0x2e, 0x62, 0x95,
	0x1a, 0xf2, 0xe5, 0x96, 0x50, 0x91, 0x50, 0xc4, 0xa7, 0x8a, 0xa5, 0x93, 0xc8, 0x69, 0xc5, 0x67,
	0x9a, 0x56, 0x88, 0xa4, 0x01, 0x8f, 0x8d, 0xd8, 0x6a, 0x51, 0x92, 0x26, 0x85, 0xe2, 0xb3, 0x33,
	0xef, 0x3f, 0xfc, 0x71, 0x94, 0xb8, 0xea, 0x4c, 0x37, 0xec, 0x9f, 0x26, 0xeb, 0xf6, 0x99, 0xd2,
	0xe8, 0x2b, 0x74, 0x78, 0x3b, 0x07, 0x8a, 0xa0, 0xf4, 0xb1, 0xe9, 0xf0, 0xb6, 0x77, 0x0c, 0x73,
	0xab, 0x52, 0x25, 0x45, 0xac, 0x18, 0xda, 0x83, 0x5f, 0x26, 0x67, 0xc6, 0x91, 0xad, 0x16, 0xf0,
	0x4a, 0x15, 0x3c, 0xb5, 0x4d, 0xc5, 0x1e, 0xb5, 0xf9, 0xb5, 0x30, 0x5c, 0xce, 0x3f, 0x84, 0x70,
	0x56, 0xc1, 0x4e, 0xfd, 0x8b, 0xd3, 0xbe, 0x38, 0xe9, 0x8b, 0xd3, 0x9b, 0xb3, 0x7d, 0x71, 0x83,
	0x06, 0xcc, 0x7a, 0x9b, 0x73, 0x4e, 0xef, 0x0e, 0x58, 0xf0, 0x85, 0x8c, 0xb5, 0xe0, 0x1f, 0xb6,
	0x06, 0x47, 0xf5, 0x05, 0x3a, 0xc7, 0xd0, 0xfd, 0x7b, 0x97, 0x2e, 0x4d, 0x9d, 0xc7, 0xab, 0xde,
	0x3b, 0xf0, 0x93, 0xc1, 0x43, 0xd7, 0x60, 0x06, 0x83, 0xca, 0x6b, 0x30, 0x36, 0x6c, 0x2a, 0xbf,
	0xb3, 0x95, 0x36, 0xcd, 0xf6, 0x4a, 0x97, 0x8f, 0x2f, 0xb7, 0x8e, 0x87, 0x8a, 0xc4, 0x98, 0x88,
	0x35, 0x91, 0xf9, 0xc7, 0x41, 0xce, 0x79, 0xfb, 0x02, 0x5d, 0x01, 0x98, 0x9d, 0xd8, 0x6b, 0x61,
	0xb8, 0x19, 0x69, 0x75, 0x79, 0x9b, 0x91, 0xd6, 0x2c, 0xc1, 0xfb, 0x6d, 0x90, 0x7e, 0xa1, 0xc2,
	0x1b, 0x48, 0x07, 0xfb, 0x0f, 0x23, 0x17, 0x0c, 0x47, 0x2e, 0x78, 0x1e, 0xb9, 0xe0, 0x66, 0xec,
	0x66, 0x86, 0x63, 0x37, 0xf3, 0x34, 0x76, 0x33, 0x27, 0x7f, 0x02, 0xae, 0x3b, 0x7d, 0x1f, 0xb7,
	0x44, 0xb4, 0x34, 0xe0, 0xcc, 0x8c, 0xd0, 0x03, 0xc9, 0x94, 0xff, 0xd9, 0x3c, 0xf8, 0xdd, 0xd7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x44, 0x76, 0x3d, 0x74, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	Position(ctx context.Context, in *QueryGetPositionRequest, opts ...grpc.CallOption) (*QueryGetPositionResponse, error)
	PositionAll(ctx context.Context, in *QueryAllPositionRequest, opts ...grpc.CallOption) (*QueryAllPositionResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Position(ctx context.Context, in *QueryGetPositionRequest, opts ...grpc.CallOption) (*QueryGetPositionResponse, error) {
	out := new(QueryGetPositionResponse)
	err := c.cc.Invoke(ctx, "/chain.mpatest.mpa.Query/Position", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionAll(ctx context.Context, in *QueryAllPositionRequest, opts ...grpc.CallOption) (*QueryAllPositionResponse, error) {
	out := new(QueryAllPositionResponse)
	err := c.cc.Invoke(ctx, "/chain.mpatest.mpa.Query/PositionAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Position(context.Context, *QueryGetPositionRequest) (*QueryGetPositionResponse, error)
	PositionAll(context.Context, *QueryAllPositionRequest) (*QueryAllPositionResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Position(ctx context.Context, req *QueryGetPositionRequest) (*QueryGetPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (*UnimplementedQueryServer) PositionAll(ctx context.Context, req *QueryAllPositionRequest) (*QueryAllPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.mpatest.mpa.Query/Position",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Position(ctx, req.(*QueryGetPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.mpatest.mpa.Query/PositionAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionAll(ctx, req.(*QueryAllPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chain.mpatest.mpa.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Position",
			Handler:    _Query_Position_Handler,
		},
		{
			MethodName: "PositionAll",
			Handler:    _Query_PositionAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mpa/query.proto",
}

func (m *QueryGetPositionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPositionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPositionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetPositionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPositionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPositionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Position != nil {
		{
			size, err := m.Position.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPositionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPositionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPositionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPositionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPositionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPositionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Position) > 0 {
		for iNdEx := len(m.Position) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Position[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetPositionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetPositionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPositionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPositionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Position) > 0 {
		for _, e := range m.Position {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetPositionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetPositionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPositionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetPositionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetPositionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPositionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &Position{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllPositionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPositionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPositionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllPositionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPositionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPositionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = append(m.Position, &Position{})
			if err := m.Position[len(m.Position)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
