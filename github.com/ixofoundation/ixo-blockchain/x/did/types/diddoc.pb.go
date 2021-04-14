// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/diddoc.proto

package types

import (
	fmt "fmt"
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

type BaseDidDoc struct {
	Did         string           `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty" json:"did" yaml:"did"`
	PubKey      string           `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty" json:"pubKey" yaml:"pubKey"`
	Credentials []*DidCredential `protobuf:"bytes,3,rep,name=credentials,proto3" json:"credentials,omitempty" json:"credentials" yaml:"credentials"`
}

func (m *BaseDidDoc) Reset()      { *m = BaseDidDoc{} }
func (*BaseDidDoc) ProtoMessage() {}
func (*BaseDidDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_df34300e393a57f6, []int{0}
}
func (m *BaseDidDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseDidDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseDidDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseDidDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseDidDoc.Merge(m, src)
}
func (m *BaseDidDoc) XXX_Size() int {
	return m.Size()
}
func (m *BaseDidDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseDidDoc.DiscardUnknown(m)
}

var xxx_messageInfo_BaseDidDoc proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseDidDoc)(nil), "did.BaseDidDoc")
}

func init() { proto.RegisterFile("did/diddoc.proto", fileDescriptor_df34300e393a57f6) }

var fileDescriptor_df34300e393a57f6 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xc9, 0x4c, 0xd1,
	0x4f, 0xc9, 0x4c, 0x49, 0xc9, 0x4f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xc9,
	0x4c, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x14, 0x2f,
	0x54, 0x31, 0x84, 0xab, 0xf4, 0x90, 0x91, 0x8b, 0xcb, 0x29, 0xb1, 0x38, 0xd5, 0x25, 0x33, 0xc5,
	0x25, 0x3f, 0x59, 0x48, 0x9b, 0x0b, 0xa4, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xf2,
	0xd3, 0x3d, 0x79, 0xd1, 0xac, 0xe2, 0xfc, 0x3c, 0x2b, 0xa5, 0x94, 0xcc, 0x14, 0x25, 0x85, 0xca,
	0xc4, 0xdc, 0x1c, 0x08, 0x33, 0x08, 0xa4, 0x4a, 0xc8, 0x9c, 0x8b, 0xad, 0xa0, 0x34, 0xc9, 0x3b,
	0xb5, 0x52, 0x82, 0x09, 0xac, 0x5e, 0xfe, 0xd3, 0x3d, 0x79, 0x69, 0x88, 0x7a, 0x88, 0x38, 0x4c,
	0x0b, 0x94, 0x17, 0x04, 0x55, 0x2e, 0x14, 0xcf, 0xc5, 0x9d, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57,
	0x92, 0x99, 0x98, 0x53, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa4, 0x07, 0x72, 0x95,
	0x4b, 0x66, 0x8a, 0x33, 0x5c, 0xca, 0x49, 0xf3, 0xd3, 0x3d, 0x79, 0x55, 0x88, 0x89, 0x48, 0x1a,
	0x60, 0xc6, 0x22, 0x0b, 0x05, 0x21, 0x9b, 0x68, 0xc5, 0xd3, 0xb1, 0x40, 0x9e, 0x61, 0xc6, 0x02,
	0x79, 0x86, 0x17, 0x0b, 0xe4, 0x19, 0x9c, 0x7c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0xca, 0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xb3,
	0x22, 0x3f, 0x2d, 0xbf, 0x34, 0x2f, 0x25, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xc4, 0xd3, 0x4d, 0xca,
	0xc9, 0x4f, 0xce, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xaf, 0x00, 0x05, 0x99, 0x7e, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xe4, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xd9,
	0xef, 0xb9, 0x77, 0x01, 0x00, 0x00,
}

func (m *BaseDidDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseDidDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseDidDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Credentials) > 0 {
		for iNdEx := len(m.Credentials) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Credentials[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDiddoc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintDiddoc(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDiddoc(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDiddoc(dAtA []byte, offset int, v uint64) int {
	offset -= sovDiddoc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseDidDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDiddoc(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovDiddoc(uint64(l))
	}
	if len(m.Credentials) > 0 {
		for _, e := range m.Credentials {
			l = e.Size()
			n += 1 + l + sovDiddoc(uint64(l))
		}
	}
	return n
}

func sovDiddoc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDiddoc(x uint64) (n int) {
	return sovDiddoc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseDidDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiddoc
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
			return fmt.Errorf("proto: BaseDidDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseDidDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiddoc
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
				return ErrInvalidLengthDiddoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiddoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiddoc
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
				return ErrInvalidLengthDiddoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiddoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiddoc
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
				return ErrInvalidLengthDiddoc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiddoc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = append(m.Credentials, &DidCredential{})
			if err := m.Credentials[len(m.Credentials)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiddoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiddoc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiddoc
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
func skipDiddoc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiddoc
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
					return 0, ErrIntOverflowDiddoc
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
					return 0, ErrIntOverflowDiddoc
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
				return 0, ErrInvalidLengthDiddoc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDiddoc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDiddoc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDiddoc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiddoc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDiddoc = fmt.Errorf("proto: unexpected end of group")
)