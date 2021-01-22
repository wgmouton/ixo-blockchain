// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

// Request/response types from old x/did/client/cli/query.go
type QueryDidDocRequest struct {
	//Did
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *QueryDidDocRequest) Reset()         { *m = QueryDidDocRequest{} }
func (m *QueryDidDocRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocRequest) ProtoMessage()    {}
func (*QueryDidDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{0}
}
func (m *QueryDidDocRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocRequest.Merge(m, src)
}
func (m *QueryDidDocRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocRequest proto.InternalMessageInfo

func (m *QueryDidDocRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type QueryDidDocResponse struct {
	// DidDoc corresponding to given DID
	Diddoc string `protobuf:"bytes,1,opt,name=diddoc,proto3" json:"diddoc,omitempty"`
}

func (m *QueryDidDocResponse) Reset()         { *m = QueryDidDocResponse{} }
func (m *QueryDidDocResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocResponse) ProtoMessage()    {}
func (*QueryDidDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{1}
}
func (m *QueryDidDocResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocResponse.Merge(m, src)
}
func (m *QueryDidDocResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocResponse proto.InternalMessageInfo

func (m *QueryDidDocResponse) GetDiddoc() string {
	if m != nil {
		return m.Diddoc
	}
	return ""
}

type QueryAllDidsRequest struct {
}

func (m *QueryAllDidsRequest) Reset()         { *m = QueryAllDidsRequest{} }
func (m *QueryAllDidsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidsRequest) ProtoMessage()    {}
func (*QueryAllDidsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{2}
}
func (m *QueryAllDidsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidsRequest.Merge(m, src)
}
func (m *QueryAllDidsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidsRequest proto.InternalMessageInfo

type QueryAllDidsResponse struct {
	// List of all DIDs
	Dids string `protobuf:"bytes,1,opt,name=dids,proto3" json:"dids,omitempty"`
}

func (m *QueryAllDidsResponse) Reset()         { *m = QueryAllDidsResponse{} }
func (m *QueryAllDidsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidsResponse) ProtoMessage()    {}
func (*QueryAllDidsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{3}
}
func (m *QueryAllDidsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidsResponse.Merge(m, src)
}
func (m *QueryAllDidsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidsResponse proto.InternalMessageInfo

func (m *QueryAllDidsResponse) GetDids() string {
	if m != nil {
		return m.Dids
	}
	return ""
}

type QueryAllDidDocsRequest struct {
}

func (m *QueryAllDidDocsRequest) Reset()         { *m = QueryAllDidDocsRequest{} }
func (m *QueryAllDidDocsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidDocsRequest) ProtoMessage()    {}
func (*QueryAllDidDocsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{4}
}
func (m *QueryAllDidDocsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidDocsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidDocsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidDocsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidDocsRequest.Merge(m, src)
}
func (m *QueryAllDidDocsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidDocsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidDocsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidDocsRequest proto.InternalMessageInfo

type QueryAllDidDocsResponse struct {
	// List of all DidDocs
	Diddocs string `protobuf:"bytes,1,opt,name=diddocs,proto3" json:"diddocs,omitempty"`
}

func (m *QueryAllDidDocsResponse) Reset()         { *m = QueryAllDidDocsResponse{} }
func (m *QueryAllDidDocsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidDocsResponse) ProtoMessage()    {}
func (*QueryAllDidDocsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{5}
}
func (m *QueryAllDidDocsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidDocsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidDocsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidDocsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidDocsResponse.Merge(m, src)
}
func (m *QueryAllDidDocsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidDocsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidDocsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidDocsResponse proto.InternalMessageInfo

func (m *QueryAllDidDocsResponse) GetDiddocs() string {
	if m != nil {
		return m.Diddocs
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryDidDocRequest)(nil), "did.QueryDidDocRequest")
	proto.RegisterType((*QueryDidDocResponse)(nil), "did.QueryDidDocResponse")
	proto.RegisterType((*QueryAllDidsRequest)(nil), "did.QueryAllDidsRequest")
	proto.RegisterType((*QueryAllDidsResponse)(nil), "did.QueryAllDidsResponse")
	proto.RegisterType((*QueryAllDidDocsRequest)(nil), "did.QueryAllDidDocsRequest")
	proto.RegisterType((*QueryAllDidDocsResponse)(nil), "did.QueryAllDidDocsResponse")
}

func init() { proto.RegisterFile("did/query.proto", fileDescriptor_31228b4ee4821623) }

var fileDescriptor_31228b4ee4821623 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0xc7, 0x29, 0x62, 0x89, 0x4f, 0x88, 0x66, 0x40, 0xa8, 0x95, 0x34, 0xa6, 0x07, 0x63, 0x4c,
	0x60, 0x12, 0xb9, 0x7a, 0xd1, 0xf4, 0x68, 0x62, 0xe4, 0xe0, 0xc1, 0x78, 0x29, 0x7d, 0xb5, 0x4c,
	0xac, 0xf3, 0x0a, 0x6d, 0x13, 0x88, 0xf1, 0x62, 0xfc, 0x00, 0x26, 0x7e, 0xa9, 0x3d, 0x92, 0xec,
	0x65, 0x8f, 0x1b, 0xd8, 0x0f, 0xb2, 0x99, 0xe9, 0xc0, 0xc2, 0xc2, 0xed, 0xfd, 0xfb, 0x7f, 0xef,
	0xd7, 0xf9, 0xbf, 0x3c, 0x78, 0x82, 0x02, 0xf9, 0xbc, 0x8c, 0x17, 0xab, 0x51, 0xb6, 0xa0, 0x82,
	0xd8, 0x03, 0x14, 0xe8, 0x76, 0x13, 0x4a, 0x48, 0x6b, 0xae, 0xaa, 0xca, 0x72, 0x07, 0x09, 0x51,
	0x92, 0xc6, 0x3c, 0xcc, 0x04, 0x0f, 0xa5, 0xa4, 0x22, 0x2c, 0x04, 0xc9, 0xdc, 0xb8, 0x6d, 0x45,
	0x42, 0x81, 0x95, 0xf4, 0x5f, 0x01, 0xfb, 0xac, 0xb0, 0x81, 0xc0, 0x80, 0xa2, 0x49, 0x3c, 0x2f,
	0xe3, 0xbc, 0x60, 0x4f, 0x41, 0xf1, 0x1d, 0xeb, 0xa5, 0xf5, 0xfa, 0xd1, 0x44, 0x95, 0xfe, 0x10,
	0x3a, 0x47, 0x7d, 0x79, 0x46, 0x32, 0x8f, 0x59, 0x0f, 0x6c, 0x14, 0x88, 0x14, 0x99, 0x5e, 0xa3,
	0xfc, 0x67, 0xa6, 0xfd, 0x7d, 0x9a, 0x06, 0x02, 0x73, 0xc3, 0xf5, 0xdf, 0x40, 0xf7, 0xf8, 0xb3,
	0xc1, 0x30, 0x68, 0xa0, 0xc0, 0xdc, 0x40, 0x74, 0xed, 0x3b, 0xd0, 0x3b, 0xe8, 0x0d, 0x28, 0xda,
	0x53, 0xc6, 0xd0, 0x3f, 0x71, 0x0c, 0xc8, 0x81, 0x66, 0xf5, 0x82, 0x1d, 0x6b, 0x27, 0xdf, 0xfe,
	0xad, 0xc3, 0x43, 0x3d, 0xc5, 0x3e, 0x81, 0x5d, 0x8d, 0xb1, 0xfe, 0x48, 0x2d, 0xe2, 0x34, 0xbf,
	0xeb, 0x9c, 0x1a, 0xd5, 0x0f, 0x7c, 0xf6, 0xe7, 0xf2, 0xe6, 0x7f, 0xbd, 0xc5, 0x40, 0xed, 0x90,
	0xff, 0x42, 0x81, 0xbf, 0xd9, 0x47, 0x68, 0x9a, 0x40, 0xec, 0x60, 0xf0, 0x38, 0xba, 0xfb, 0xfc,
	0x8c, 0x63, 0x98, 0x2d, 0xcd, 0xb4, 0x59, 0x43, 0x31, 0xd9, 0x37, 0x80, 0xbb, 0x60, 0xec, 0xc5,
	0xfd, 0xb1, 0x83, 0x45, 0xb8, 0x83, 0xf3, 0xa6, 0xc1, 0x76, 0x34, 0xb6, 0xcd, 0x1e, 0xf3, 0x70,
	0x6f, 0x7e, 0xf8, 0x72, 0xb1, 0xf1, 0xac, 0xf5, 0xc6, 0xb3, 0xae, 0x37, 0x9e, 0xf5, 0x6f, 0xeb,
	0xd5, 0xd6, 0x5b, 0xaf, 0x76, 0xb5, 0xf5, 0x6a, 0x5f, 0xdf, 0x25, 0xa2, 0x98, 0x95, 0xd3, 0x51,
	0x44, 0x3f, 0xb9, 0x58, 0xd2, 0x77, 0x2a, 0x25, 0xea, 0xcb, 0x51, 0x6a, 0x38, 0x4d, 0x29, 0xfa,
	0x11, 0xcd, 0x42, 0x21, 0xf9, 0x52, 0x47, 0x17, 0xb2, 0x88, 0x17, 0x32, 0x4c, 0x79, 0xb1, 0xca,
	0xe2, 0x7c, 0x6a, 0xeb, 0x73, 0x1a, 0xdf, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x36, 0xa7, 0x9c,
	0xa9, 0x02, 0x00, 0x00,
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
	DidDoc(ctx context.Context, in *QueryDidDocRequest, opts ...grpc.CallOption) (*QueryDidDocResponse, error)
	AllDids(ctx context.Context, in *QueryAllDidsRequest, opts ...grpc.CallOption) (*QueryAllDidsResponse, error)
	AllDidDocs(ctx context.Context, in *QueryAllDidDocsRequest, opts ...grpc.CallOption) (*QueryAllDidDocsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDoc(ctx context.Context, in *QueryDidDocRequest, opts ...grpc.CallOption) (*QueryDidDocResponse, error) {
	out := new(QueryDidDocResponse)
	err := c.cc.Invoke(ctx, "/did.Query/DidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllDids(ctx context.Context, in *QueryAllDidsRequest, opts ...grpc.CallOption) (*QueryAllDidsResponse, error) {
	out := new(QueryAllDidsResponse)
	err := c.cc.Invoke(ctx, "/did.Query/AllDids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllDidDocs(ctx context.Context, in *QueryAllDidDocsRequest, opts ...grpc.CallOption) (*QueryAllDidDocsResponse, error) {
	out := new(QueryAllDidDocsResponse)
	err := c.cc.Invoke(ctx, "/did.Query/AllDidDocs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	DidDoc(context.Context, *QueryDidDocRequest) (*QueryDidDocResponse, error)
	AllDids(context.Context, *QueryAllDidsRequest) (*QueryAllDidsResponse, error)
	AllDidDocs(context.Context, *QueryAllDidDocsRequest) (*QueryAllDidDocsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDoc(ctx context.Context, req *QueryDidDocRequest) (*QueryDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDoc not implemented")
}
func (*UnimplementedQueryServer) AllDids(ctx context.Context, req *QueryAllDidsRequest) (*QueryAllDidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllDids not implemented")
}
func (*UnimplementedQueryServer) AllDidDocs(ctx context.Context, req *QueryAllDidDocsRequest) (*QueryAllDidDocsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllDidDocs not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/did.Query/DidDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDoc(ctx, req.(*QueryDidDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllDids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllDids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/did.Query/AllDids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllDids(ctx, req.(*QueryAllDidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllDidDocs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDidDocsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllDidDocs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/did.Query/AllDidDocs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllDidDocs(ctx, req.(*QueryAllDidDocsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "did.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDoc",
			Handler:    _Query_DidDoc_Handler,
		},
		{
			MethodName: "AllDids",
			Handler:    _Query_AllDids_Handler,
		},
		{
			MethodName: "AllDidDocs",
			Handler:    _Query_AllDidDocs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/query.proto",
}

func (m *QueryDidDocRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Diddoc) > 0 {
		i -= len(m.Diddoc)
		copy(dAtA[i:], m.Diddoc)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Diddoc)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllDidsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllDidsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Dids) > 0 {
		i -= len(m.Dids)
		copy(dAtA[i:], m.Dids)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Dids)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllDidDocsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidDocsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidDocsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllDidDocsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidDocsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidDocsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Diddocs) > 0 {
		i -= len(m.Diddocs)
		copy(dAtA[i:], m.Diddocs)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Diddocs)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryDidDocRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Diddoc)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllDidsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllDidsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dids)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllDidDocsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllDidDocsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Diddocs)
	if l > 0 {
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
func (m *QueryDidDocRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
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
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryDidDocResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Diddoc", wireType)
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
			m.Diddoc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAllDidsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDidsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAllDidsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDidsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dids", wireType)
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
			m.Dids = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAllDidDocsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDidDocsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidDocsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAllDidDocsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDidDocsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidDocsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Diddocs", wireType)
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
			m.Diddocs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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