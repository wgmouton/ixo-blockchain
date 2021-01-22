// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/did.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type DidCredential struct {
	Credtype []string `protobuf:"bytes,1,rep,name=credtype,proto3" json:"credtype,omitempty"`
	//Did
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Issued string `protobuf:"bytes,3,opt,name=issued,proto3" json:"issued,omitempty"`
	Claim  *Claim `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *DidCredential) Reset()      { *m = DidCredential{} }
func (*DidCredential) ProtoMessage() {}
func (*DidCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{0}
}
func (m *DidCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidCredential.Merge(m, src)
}
func (m *DidCredential) XXX_Size() int {
	return m.Size()
}
func (m *DidCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_DidCredential.DiscardUnknown(m)
}

var xxx_messageInfo_DidCredential proto.InternalMessageInfo

func (m *DidCredential) GetCredtype() []string {
	if m != nil {
		return m.Credtype
	}
	return nil
}

func (m *DidCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *DidCredential) GetIssued() string {
	if m != nil {
		return m.Issued
	}
	return ""
}

func (m *DidCredential) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

type Claim struct {
	//Did
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KYCvalidated bool   `protobuf:"varint,2,opt,name=KYCvalidated,proto3" json:"KYCvalidated,omitempty"`
}

func (m *Claim) Reset()      { *m = Claim{} }
func (*Claim) ProtoMessage() {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claim) GetKYCvalidated() bool {
	if m != nil {
		return m.KYCvalidated
	}
	return false
}

func init() {
	proto.RegisterType((*DidCredential)(nil), "did.DidCredential")
	proto.RegisterType((*Claim)(nil), "did.Claim")
}

func init() { proto.RegisterFile("did/did.proto", fileDescriptor_a95066e7350e77d5) }

var fileDescriptor_a95066e7350e77d5 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xfe, 0xa9, 0x35, 0x94, 0xc1, 0x42, 0x28, 0x74, 0x30, 0x51, 0xa7, 0x2e, 0xd4,
	0x12, 0x6c, 0x88, 0x05, 0xca, 0xc6, 0x96, 0x01, 0x09, 0x36, 0x27, 0xd7, 0x4d, 0xaf, 0x70, 0xed,
	0x2a, 0x75, 0x50, 0x3b, 0xf3, 0x02, 0x8c, 0x8c, 0x3c, 0x0e, 0x63, 0x47, 0x46, 0xd4, 0xbe, 0x08,
	0xb2, 0x83, 0x2a, 0xd8, 0xfc, 0x7d, 0x47, 0x3a, 0xb6, 0x0f, 0xed, 0x03, 0x82, 0x00, 0x84, 0xf1,
	0xa2, 0xb4, 0xce, 0xb2, 0x26, 0x20, 0x0c, 0x8e, 0x0b, 0x5b, 0xd8, 0xc0, 0xc2, 0x9f, 0xea, 0x68,
	0x70, 0x5a, 0x58, 0x5b, 0x68, 0x25, 0x02, 0x65, 0xd5, 0x54, 0x48, 0xb3, 0xae, 0xa3, 0xe1, 0x2b,
	0xa1, 0xfd, 0x3b, 0x84, 0x49, 0xa9, 0x40, 0x19, 0x87, 0x52, 0xb3, 0x01, 0xed, 0xe6, 0xa5, 0x02,
	0xb7, 0x5e, 0xa8, 0x98, 0x24, 0xcd, 0x51, 0x2f, 0xdd, 0x33, 0x3b, 0xa1, 0x1d, 0x5c, 0x2e, 0x2b,
	0x55, 0xc6, 0x8d, 0x84, 0x8c, 0x7a, 0xe9, 0x2f, 0xed, 0x3d, 0xc4, 0xcd, 0x3f, 0x1e, 0x58, 0x42,
	0xdb, 0xb9, 0x96, 0x38, 0x8f, 0x5b, 0x09, 0x19, 0x1d, 0x5c, 0xd0, 0xb1, 0x7f, 0xee, 0xc4, 0x9b,
	0xb4, 0x0e, 0xae, 0x5a, 0xef, 0x1f, 0x67, 0xd1, 0xf0, 0x86, 0xb6, 0x83, 0x65, 0x47, 0xb4, 0x81,
	0x10, 0x93, 0x50, 0xd2, 0x40, 0x60, 0x43, 0x7a, 0x78, 0xff, 0x38, 0x79, 0x91, 0x1a, 0x41, 0x3a,
	0x05, 0xe1, 0xda, 0x6e, 0xfa, 0xcf, 0xd5, 0x15, 0xb7, 0x0f, 0x9f, 0x5b, 0x4e, 0x36, 0x5b, 0x4e,
	0xbe, 0xb7, 0x9c, 0xbc, 0xed, 0x78, 0xb4, 0xd9, 0xf1, 0xe8, 0x6b, 0xc7, 0xa3, 0xa7, 0xeb, 0x02,
	0xdd, 0xac, 0xca, 0xc6, 0xb9, 0x9d, 0x0b, 0x5c, 0xd9, 0xa9, 0xad, 0x0c, 0x48, 0x87, 0xd6, 0x78,
	0x3a, 0xcf, 0xb4, 0xcd, 0x9f, 0xf3, 0x99, 0x44, 0x23, 0x56, 0x7e, 0x4d, 0x81, 0xc6, 0xa9, 0xd2,
	0x48, 0x2d, 0xfc, 0x8f, 0x97, 0x59, 0x27, 0xec, 0x74, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2f,
	0xcb, 0xad, 0x56, 0x6e, 0x01, 0x00, 0x00,
}

func (m *DidCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issued) > 0 {
		i -= len(m.Issued)
		copy(dAtA[i:], m.Issued)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issued)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Credtype) > 0 {
		for iNdEx := len(m.Credtype) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Credtype[iNdEx])
			copy(dAtA[i:], m.Credtype[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Credtype[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KYCvalidated {
		i--
		if m.KYCvalidated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Credtype) > 0 {
		for _, s := range m.Credtype {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Issued)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.KYCvalidated {
		n += 2
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: DidCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credtype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credtype = append(m.Credtype, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
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
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KYCvalidated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.KYCvalidated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
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
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)