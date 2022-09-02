// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type QueryEntityListRequest struct {
	EntityType   string `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty" yaml:"entity_type"`
	EntityStatus string `protobuf:"bytes,2,opt,name=entity_status,json=entityStatus,proto3" json:"entity_status,omitempty" yaml:"entity_status"`
}

func (m *QueryEntityListRequest) Reset()         { *m = QueryEntityListRequest{} }
func (m *QueryEntityListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEntityListRequest) ProtoMessage()    {}
func (*QueryEntityListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b58e82145e4ece, []int{0}
}
func (m *QueryEntityListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEntityListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEntityListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEntityListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEntityListRequest.Merge(m, src)
}
func (m *QueryEntityListRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEntityListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEntityListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEntityListRequest proto.InternalMessageInfo

func (m *QueryEntityListRequest) GetEntityType() string {
	if m != nil {
		return m.EntityType
	}
	return ""
}

func (m *QueryEntityListRequest) GetEntityStatus() string {
	if m != nil {
		return m.EntityStatus
	}
	return ""
}

// // QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.
type QueryEntityListResponse struct {
}

func (m *QueryEntityListResponse) Reset()         { *m = QueryEntityListResponse{} }
func (m *QueryEntityListResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEntityListResponse) ProtoMessage()    {}
func (*QueryEntityListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b58e82145e4ece, []int{1}
}
func (m *QueryEntityListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEntityListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEntityListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEntityListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEntityListResponse.Merge(m, src)
}
func (m *QueryEntityListResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEntityListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEntityListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEntityListResponse proto.InternalMessageInfo

// QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC method.
type QueryEntityDocRequest struct {
	EntityDid string `protobuf:"bytes,1,opt,name=entity_did,json=entityDid,proto3" json:"entity_did,omitempty" yaml:"entity_did"`
}

func (m *QueryEntityDocRequest) Reset()         { *m = QueryEntityDocRequest{} }
func (m *QueryEntityDocRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEntityDocRequest) ProtoMessage()    {}
func (*QueryEntityDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b58e82145e4ece, []int{2}
}
func (m *QueryEntityDocRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEntityDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEntityDocRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEntityDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEntityDocRequest.Merge(m, src)
}
func (m *QueryEntityDocRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEntityDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEntityDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEntityDocRequest proto.InternalMessageInfo

func (m *QueryEntityDocRequest) GetEntityDid() string {
	if m != nil {
		return m.EntityDid
	}
	return ""
}

// // QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.
type QueryEntityDocResponse struct {
}

func (m *QueryEntityDocResponse) Reset()         { *m = QueryEntityDocResponse{} }
func (m *QueryEntityDocResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEntityDocResponse) ProtoMessage()    {}
func (*QueryEntityDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b58e82145e4ece, []int{3}
}
func (m *QueryEntityDocResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEntityDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEntityDocResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEntityDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEntityDocResponse.Merge(m, src)
}
func (m *QueryEntityDocResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEntityDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEntityDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEntityDocResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryEntityListRequest)(nil), "entity.QueryEntityListRequest")
	proto.RegisterType((*QueryEntityListResponse)(nil), "entity.QueryEntityListResponse")
	proto.RegisterType((*QueryEntityDocRequest)(nil), "entity.QueryEntityDocRequest")
	proto.RegisterType((*QueryEntityDocResponse)(nil), "entity.QueryEntityDocResponse")
}

func init() { proto.RegisterFile("entity/query.proto", fileDescriptor_f9b58e82145e4ece) }

var fileDescriptor_f9b58e82145e4ece = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x6d, 0x1e, 0xf8, 0xa0, 0xf3, 0x7c, 0x8b, 0x37, 0xfd, 0x30, 0x06, 0x9d, 0x94, 0xac, 0xdc,
	0xd8, 0x01, 0x15, 0x05, 0xc1, 0x4d, 0xa9, 0x3b, 0x05, 0x8d, 0xae, 0xdc, 0x94, 0x7c, 0x8c, 0xe9,
	0x60, 0x3a, 0x37, 0xed, 0x4c, 0xa0, 0x41, 0xdc, 0xf8, 0x0b, 0x0a, 0xfe, 0x29, 0x97, 0x05, 0x37,
	0xba, 0x29, 0xd2, 0xfa, 0x0b, 0xfa, 0x0b, 0x24, 0x33, 0x53, 0xdb, 0x62, 0xbb, 0x4a, 0xce, 0x39,
	0x73, 0xef, 0x39, 0xf7, 0xce, 0x20, 0xcc, 0x84, 0xe2, 0xaa, 0xa2, 0xd3, 0x92, 0xcd, 0xaa, 0x7e,
	0x31, 0x03, 0x05, 0xf8, 0xd2, 0x70, 0x5e, 0x3b, 0x83, 0x0c, 0x34, 0x45, 0xeb, 0x3f, 0xa3, 0x7a,
	0xf7, 0x32, 0x80, 0x2c, 0x67, 0x34, 0x2a, 0x38, 0x8d, 0x84, 0x00, 0x15, 0x29, 0x0e, 0x42, 0x5a,
	0xf5, 0x26, 0x01, 0x39, 0x01, 0x49, 0x13, 0xe0, 0xc2, 0x52, 0x2d, 0x6b, 0x61, 0x3e, 0x86, 0x0c,
	0x16, 0x0e, 0xea, 0xbe, 0xad, 0x3d, 0x5f, 0x6a, 0xf6, 0x15, 0x97, 0x2a, 0x64, 0xd3, 0x92, 0x49,
	0x85, 0x9f, 0xa1, 0x2b, 0x73, 0x74, 0xa4, 0xaa, 0x82, 0xb9, 0x4e, 0xcf, 0x79, 0xd0, 0x1c, 0x74,
	0xb7, 0x2b, 0x1f, 0x57, 0xd1, 0x24, 0x7f, 0x1e, 0x1c, 0x88, 0x41, 0x88, 0x0c, 0x7a, 0x5f, 0x15,
	0x0c, 0xbf, 0x40, 0xd7, 0x56, 0x93, 0x2a, 0x52, 0xa5, 0x74, 0x2f, 0x74, 0xa9, 0xbb, 0x5d, 0xf9,
	0xed, 0xa3, 0x52, 0x23, 0x07, 0xe1, 0x6d, 0x83, 0xdf, 0x19, 0x78, 0x17, 0xdd, 0xf9, 0x2f, 0x91,
	0x2c, 0x40, 0x48, 0x16, 0xbc, 0x46, 0x9d, 0x03, 0x69, 0x08, 0xc9, 0x2e, 0xeb, 0x13, 0x64, 0x03,
	0x8c, 0x52, 0x9e, 0xda, 0xa8, 0x9d, 0xed, 0xca, 0xbf, 0x39, 0xf2, 0x4b, 0x79, 0x1a, 0x84, 0x4d,
	0x03, 0x86, 0x3c, 0x0d, 0xdc, 0xa3, 0xd9, 0x75, 0x3b, 0x63, 0xf4, 0xe8, 0x97, 0x83, 0x6e, 0x69,
	0x09, 0xc7, 0x08, 0xed, 0x83, 0x60, 0xd2, 0xb7, 0xdb, 0x3b, 0xbd, 0x33, 0xcf, 0x3f, 0xab, 0xdb,
	0x09, 0x5a, 0x5f, 0x7f, 0xfc, 0xf9, 0x76, 0x71, 0x8d, 0xaf, 0x28, 0x9f, 0x83, 0xbd, 0x0a, 0x9c,
	0xa3, 0xe6, 0xbf, 0x08, 0xf8, 0xfe, 0x89, 0x16, 0xfb, 0x49, 0x3d, 0x72, 0x4e, 0xb6, 0x06, 0x3d,
	0x6d, 0xe0, 0x61, 0xf7, 0xc0, 0x80, 0x7e, 0xde, 0x2f, 0xe0, 0xcb, 0xe0, 0xcd, 0xf7, 0x35, 0x71,
	0x96, 0x6b, 0xe2, 0xfc, 0x5e, 0x13, 0x67, 0xb1, 0x21, 0x8d, 0xe5, 0x86, 0x34, 0x7e, 0x6e, 0x48,
	0xe3, 0xc3, 0xd3, 0x8c, 0xab, 0x71, 0x19, 0xf7, 0x13, 0x98, 0xd4, 0xd5, 0x1f, 0xa1, 0x14, 0xa9,
	0x7e, 0x55, 0x35, 0x7a, 0x18, 0xe7, 0x90, 0x7c, 0x4a, 0xc6, 0x11, 0x17, 0x74, 0xbe, 0x6b, 0x5c,
	0x5f, 0xbe, 0x8c, 0x2f, 0xf5, 0x5b, 0x7a, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x50, 0xf4, 0x8e,
	0x43, 0xc5, 0x02, 0x00, 0x00,
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
	EntityList(ctx context.Context, in *QueryEntityListRequest, opts ...grpc.CallOption) (*QueryEntityListResponse, error)
	EntityDoc(ctx context.Context, in *QueryEntityDocRequest, opts ...grpc.CallOption) (*QueryEntityDocResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EntityList(ctx context.Context, in *QueryEntityListRequest, opts ...grpc.CallOption) (*QueryEntityListResponse, error) {
	out := new(QueryEntityListResponse)
	err := c.cc.Invoke(ctx, "/entity.Query/EntityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EntityDoc(ctx context.Context, in *QueryEntityDocRequest, opts ...grpc.CallOption) (*QueryEntityDocResponse, error) {
	out := new(QueryEntityDocResponse)
	err := c.cc.Invoke(ctx, "/entity.Query/EntityDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	EntityList(context.Context, *QueryEntityListRequest) (*QueryEntityListResponse, error)
	EntityDoc(context.Context, *QueryEntityDocRequest) (*QueryEntityDocResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EntityList(ctx context.Context, req *QueryEntityListRequest) (*QueryEntityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityList not implemented")
}
func (*UnimplementedQueryServer) EntityDoc(ctx context.Context, req *QueryEntityDocRequest) (*QueryEntityDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityDoc not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EntityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEntityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EntityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.Query/EntityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EntityList(ctx, req.(*QueryEntityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EntityDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEntityDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EntityDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.Query/EntityDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EntityDoc(ctx, req.(*QueryEntityDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EntityList",
			Handler:    _Query_EntityList_Handler,
		},
		{
			MethodName: "EntityDoc",
			Handler:    _Query_EntityDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/query.proto",
}

func (m *QueryEntityListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEntityListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEntityListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EntityStatus) > 0 {
		i -= len(m.EntityStatus)
		copy(dAtA[i:], m.EntityStatus)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EntityStatus)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EntityType) > 0 {
		i -= len(m.EntityType)
		copy(dAtA[i:], m.EntityType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EntityType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEntityListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEntityListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEntityListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEntityDocRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEntityDocRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEntityDocRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EntityDid) > 0 {
		i -= len(m.EntityDid)
		copy(dAtA[i:], m.EntityDid)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EntityDid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEntityDocResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEntityDocResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEntityDocResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *QueryEntityListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EntityType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EntityStatus)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEntityListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEntityDocRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EntityDid)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEntityDocResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEntityListRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEntityListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEntityListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityStatus = string(dAtA[iNdEx:postIndex])
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
func (m *QueryEntityListResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEntityListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEntityListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryEntityDocRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEntityDocRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEntityDocRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityDid = string(dAtA[iNdEx:postIndex])
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
func (m *QueryEntityDocResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEntityDocResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEntityDocResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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