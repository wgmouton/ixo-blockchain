// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/entity/v1beta1/entity.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Params struct {
	NftContractAddress string `protobuf:"bytes,1,opt,name=NftContractAddress,proto3" json:"NftContractAddress" yaml:"NftContractAddress"`
	NftContractMinter  string `protobuf:"bytes,2,opt,name=NftContractMinter,proto3" json:"NftContractMinter" yaml:"NftContractMinter"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9631845bd4f69820, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNftContractAddress() string {
	if m != nil {
		return m.NftContractAddress
	}
	return ""
}

func (m *Params) GetNftContractMinter() string {
	if m != nil {
		return m.NftContractMinter
	}
	return ""
}

// // ProjectDoc defines a project (or entity) type with all of its parameters.
type EntityDoc struct {
}

func (m *EntityDoc) Reset()         { *m = EntityDoc{} }
func (m *EntityDoc) String() string { return proto.CompactTextString(m) }
func (*EntityDoc) ProtoMessage()    {}
func (*EntityDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_9631845bd4f69820, []int{1}
}
func (m *EntityDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityDoc.Merge(m, src)
}
func (m *EntityDoc) XXX_Size() int {
	return m.Size()
}
func (m *EntityDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityDoc.DiscardUnknown(m)
}

var xxx_messageInfo_EntityDoc proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "ixo.entity.v1beta1.Params")
	proto.RegisterType((*EntityDoc)(nil), "ixo.entity.v1beta1.EntityDoc")
}

func init() { proto.RegisterFile("ixo/entity/v1beta1/entity.proto", fileDescriptor_9631845bd4f69820) }

var fileDescriptor_9631845bd4f69820 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0xc3, 0x40,
	0x18, 0xc5, 0x73, 0x0e, 0x85, 0xc6, 0xa9, 0xc1, 0x21, 0x76, 0xb8, 0x48, 0x26, 0x17, 0x73, 0x84,
	0x82, 0x83, 0x9b, 0x55, 0x47, 0xa5, 0x74, 0x74, 0x91, 0xcb, 0x25, 0x4d, 0x0f, 0x9b, 0xfb, 0xca,
	0xdd, 0x57, 0x49, 0xfe, 0x0b, 0xff, 0x2c, 0x27, 0xe9, 0xe8, 0x14, 0x24, 0xd9, 0x3a, 0xf6, 0x2f,
	0x90, 0x26, 0x11, 0x0a, 0xe9, 0xf6, 0xbd, 0xf7, 0x7b, 0x1f, 0x0f, 0x9e, 0xed, 0xc9, 0x1c, 0x58,
	0xa2, 0x50, 0x62, 0xc1, 0x3e, 0xc2, 0x28, 0x41, 0x1e, 0x76, 0x32, 0x58, 0x6b, 0x40, 0x70, 0x1c,
	0x99, 0x43, 0xd0, 0x39, 0x5d, 0x60, 0x7c, 0x91, 0x42, 0x0a, 0x0d, 0x66, 0x87, 0xab, 0x4d, 0x8e,
	0x47, 0x02, 0x4c, 0x06, 0x86, 0x09, 0x90, 0xaa, 0xb5, 0xfc, 0x6f, 0x62, 0x0f, 0x66, 0x5c, 0xf3,
	0xcc, 0x38, 0xc2, 0x76, 0x5e, 0x16, 0xf8, 0x00, 0x0a, 0x35, 0x17, 0x78, 0x1f, 0xc7, 0x3a, 0x31,
	0xc6, 0x25, 0x57, 0xe4, 0x7a, 0x38, 0x9d, 0xec, 0x4a, 0xef, 0x04, 0xdd, 0x97, 0xde, 0x65, 0xc1,
	0xb3, 0xd5, 0x9d, 0xdf, 0x67, 0xfe, 0xfc, 0xc4, 0x83, 0xf3, 0x66, 0x8f, 0x8e, 0xdc, 0x67, 0xa9,
	0x30, 0xd1, 0xee, 0x59, 0xd3, 0x11, 0xee, 0x4a, 0xaf, 0x0f, 0xf7, 0xa5, 0xe7, 0xf6, 0x2a, 0x5a,
	0xe4, 0xcf, 0xfb, 0x71, 0xff, 0xdc, 0x1e, 0x3e, 0x35, 0x5b, 0x3c, 0x82, 0x98, 0xce, 0xbe, 0x2a,
	0x4a, 0xb6, 0x15, 0x25, 0xbf, 0x15, 0x25, 0x9f, 0x35, 0xb5, 0xb6, 0x35, 0xb5, 0x7e, 0x6a, 0x6a,
	0xbd, 0xde, 0xa6, 0x12, 0x97, 0x9b, 0x28, 0x10, 0x90, 0x31, 0x99, 0xc3, 0x02, 0x36, 0x2a, 0xe6,
	0x28, 0x41, 0x1d, 0xd4, 0x4d, 0xb4, 0x02, 0xf1, 0x2e, 0x96, 0x5c, 0x2a, 0x96, 0xff, 0x6f, 0x8f,
	0xc5, 0x3a, 0x31, 0xd1, 0xa0, 0x99, 0x6d, 0xf2, 0x17, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x71, 0x6b,
	0x43, 0x96, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftContractMinter) > 0 {
		i -= len(m.NftContractMinter)
		copy(dAtA[i:], m.NftContractMinter)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.NftContractMinter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftContractAddress) > 0 {
		i -= len(m.NftContractAddress)
		copy(dAtA[i:], m.NftContractAddress)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.NftContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintEntity(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftContractAddress)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	l = len(m.NftContractMinter)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	return n
}

func (m *EntityDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovEntity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntity(x uint64) (n int) {
	return sovEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftContractMinter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftContractMinter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntity
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
func (m *EntityDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
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
			return fmt.Errorf("proto: EntityDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntity
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
func skipEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
				return 0, ErrInvalidLengthEntity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntity = fmt.Errorf("proto: unexpected end of group")
)
