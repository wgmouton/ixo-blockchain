// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgAddDid struct {
	//Did
	Did    string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty" yaml:"did"`
	Pubkey string `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty" yaml:"pubKey"`
}

func (m *MsgAddDid) Reset()         { *m = MsgAddDid{} }
func (m *MsgAddDid) String() string { return proto.CompactTextString(m) }
func (*MsgAddDid) ProtoMessage()    {}
func (*MsgAddDid) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f788725073a497, []int{0}
}
func (m *MsgAddDid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddDid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddDid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddDid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddDid.Merge(m, src)
}
func (m *MsgAddDid) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddDid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddDid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddDid proto.InternalMessageInfo

func (m *MsgAddDid) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *MsgAddDid) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type MsgAddDidResponse struct {
}

func (m *MsgAddDidResponse) Reset()         { *m = MsgAddDidResponse{} }
func (m *MsgAddDidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddDidResponse) ProtoMessage()    {}
func (*MsgAddDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f788725073a497, []int{1}
}
func (m *MsgAddDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddDidResponse.Merge(m, src)
}
func (m *MsgAddDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddDidResponse proto.InternalMessageInfo

type MsgAddCredential struct {
	Didcredential *DidCredential `protobuf:"bytes,1,opt,name=didcredential,proto3" json:"didcredential,omitempty" yaml:"credential"`
}

func (m *MsgAddCredential) Reset()         { *m = MsgAddCredential{} }
func (m *MsgAddCredential) String() string { return proto.CompactTextString(m) }
func (*MsgAddCredential) ProtoMessage()    {}
func (*MsgAddCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f788725073a497, []int{2}
}
func (m *MsgAddCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddCredential.Merge(m, src)
}
func (m *MsgAddCredential) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddCredential.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddCredential proto.InternalMessageInfo

func (m *MsgAddCredential) GetDidcredential() *DidCredential {
	if m != nil {
		return m.Didcredential
	}
	return nil
}

type MsgAddCredentialResponse struct {
}

func (m *MsgAddCredentialResponse) Reset()         { *m = MsgAddCredentialResponse{} }
func (m *MsgAddCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddCredentialResponse) ProtoMessage()    {}
func (*MsgAddCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f788725073a497, []int{3}
}
func (m *MsgAddCredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddCredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddCredentialResponse.Merge(m, src)
}
func (m *MsgAddCredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddCredentialResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddDid)(nil), "did.MsgAddDid")
	proto.RegisterType((*MsgAddDidResponse)(nil), "did.MsgAddDidResponse")
	proto.RegisterType((*MsgAddCredential)(nil), "did.MsgAddCredential")
	proto.RegisterType((*MsgAddCredentialResponse)(nil), "did.MsgAddCredentialResponse")
}

func init() { proto.RegisterFile("did/tx.proto", fileDescriptor_75f788725073a497) }

var fileDescriptor_75f788725073a497 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xa9, 0x24, 0x24, 0x8c, 0x42, 0x64, 0x14, 0x43, 0x9a, 0x58, 0xc8, 0xac, 0x74, 0x21,
	0x35, 0xb8, 0x33, 0x6e, 0x44, 0x5c, 0x19, 0x62, 0xd2, 0x85, 0x31, 0xee, 0xda, 0xde, 0xb1, 0xdc,
	0x50, 0x66, 0x1a, 0x3a, 0x4d, 0xe8, 0xc6, 0x67, 0xf0, 0xb1, 0x5c, 0xb2, 0x74, 0x45, 0x0c, 0xbc,
	0x01, 0x4f, 0x60, 0xa6, 0xe5, 0xd7, 0xb8, 0xeb, 0x39, 0xf3, 0xdd, 0x7b, 0x4e, 0x73, 0xc9, 0x11,
	0x20, 0xd8, 0x6a, 0xd2, 0x8e, 0xc6, 0x52, 0x49, 0x5a, 0x04, 0x04, 0xf3, 0x34, 0x90, 0x81, 0xcc,
	0xb4, 0xad, 0xbf, 0xf2, 0x27, 0xb3, 0xa2, 0x41, 0x40, 0xc8, 0x25, 0x7b, 0x25, 0xe5, 0x7e, 0x1c,
	0xdc, 0x03, 0xf4, 0x10, 0x68, 0x8b, 0xe8, 0xc1, 0x86, 0xd1, 0x32, 0x2e, 0xca, 0xdd, 0xea, 0x72,
	0xd6, 0x24, 0xa9, 0x3b, 0x0a, 0x6f, 0x19, 0x20, 0x30, 0x47, 0x3f, 0xd1, 0x4b, 0x52, 0x8a, 0x12,
	0x6f, 0xc8, 0xd3, 0xc6, 0x41, 0x06, 0xd5, 0x96, 0xb3, 0x66, 0x25, 0x87, 0xa2, 0xc4, 0x7b, 0xe2,
	0x29, 0x73, 0x56, 0x00, 0x3b, 0x21, 0xb5, 0xcd, 0x66, 0x87, 0xc7, 0x91, 0x14, 0x31, 0x67, 0x3e,
	0x39, 0xce, 0xcd, 0x87, 0x31, 0x07, 0x2e, 0x14, 0xba, 0x21, 0x7d, 0x26, 0xba, 0x93, 0xbf, 0x31,
	0xb2, 0xfc, 0xc3, 0x0e, 0x6d, 0xeb, 0x96, 0x3d, 0xdc, 0x41, 0xbb, 0xf5, 0xe5, 0xac, 0x59, 0xcb,
	0xe3, 0xb6, 0x3c, 0x73, 0xf6, 0xe7, 0x99, 0x49, 0x1a, 0x7f, 0x43, 0xd6, 0x05, 0x3a, 0x1f, 0xa4,
	0xd8, 0x8f, 0x03, 0x7a, 0x4d, 0x4a, 0xab, 0x7f, 0xae, 0x66, 0x31, 0x9b, 0xa6, 0xe6, 0xd9, 0xbe,
	0x5e, 0x0f, 0xd2, 0x47, 0x52, 0xd9, 0xaf, 0x5d, 0xdf, 0x01, 0xb7, 0xb6, 0x79, 0xfe, 0xaf, 0xbd,
	0x5e, 0xd3, 0x7d, 0xf9, 0x9a, 0x5b, 0xc6, 0x74, 0x6e, 0x19, 0x3f, 0x73, 0xcb, 0xf8, 0x5c, 0x58,
	0x85, 0xe9, 0xc2, 0x2a, 0x7c, 0x2f, 0xac, 0xc2, 0xdb, 0x5d, 0x80, 0x6a, 0x90, 0x78, 0x6d, 0x5f,
	0x8e, 0x6c, 0x9c, 0xc8, 0x77, 0x99, 0x08, 0x70, 0x15, 0x4a, 0xa1, 0xd5, 0x95, 0x17, 0x4a, 0x7f,
	0xe8, 0x0f, 0x5c, 0x14, 0xf6, 0x44, 0x9f, 0xcf, 0x46, 0xa1, 0xf8, 0x58, 0xb8, 0xa1, 0xad, 0xd2,
	0x88, 0xc7, 0x5e, 0x29, 0x3b, 0xe7, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xcb, 0xb3,
	0xc0, 0x08, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddDid(ctx context.Context, in *MsgAddDid, opts ...grpc.CallOption) (*MsgAddDidResponse, error)
	AddCredential(ctx context.Context, in *MsgAddCredential, opts ...grpc.CallOption) (*MsgAddCredentialResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddDid(ctx context.Context, in *MsgAddDid, opts ...grpc.CallOption) (*MsgAddDidResponse, error) {
	out := new(MsgAddDidResponse)
	err := c.cc.Invoke(ctx, "/did.Msg/AddDid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddCredential(ctx context.Context, in *MsgAddCredential, opts ...grpc.CallOption) (*MsgAddCredentialResponse, error) {
	out := new(MsgAddCredentialResponse)
	err := c.cc.Invoke(ctx, "/did.Msg/AddCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddDid(context.Context, *MsgAddDid) (*MsgAddDidResponse, error)
	AddCredential(context.Context, *MsgAddCredential) (*MsgAddCredentialResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddDid(ctx context.Context, req *MsgAddDid) (*MsgAddDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDid not implemented")
}
func (*UnimplementedMsgServer) AddCredential(ctx context.Context, req *MsgAddCredential) (*MsgAddCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCredential not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddDid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/did.Msg/AddDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddDid(ctx, req.(*MsgAddDid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/did.Msg/AddCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddCredential(ctx, req.(*MsgAddCredential))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "did.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDid",
			Handler:    _Msg_AddDid_Handler,
		},
		{
			MethodName: "AddCredential",
			Handler:    _Msg_AddCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/tx.proto",
}

func (m *MsgAddDid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddDid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddDid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Didcredential != nil {
		{
			size, err := m.Didcredential.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddCredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddCredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddCredentialResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddDid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Didcredential != nil {
		l = m.Didcredential.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddCredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddDid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddDid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddDid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Didcredential", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Didcredential == nil {
				m.Didcredential = &DidCredential{}
			}
			if err := m.Didcredential.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddCredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddCredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddCredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)