// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commitment/types.proto

package commitment

import (
	fmt "fmt"
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

type Sequence struct {
	Value     uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Sequence) Reset()         { *m = Sequence{} }
func (m *Sequence) String() string { return proto.CompactTextString(m) }
func (*Sequence) ProtoMessage()    {}
func (*Sequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_921e8bf76465e275, []int{0}
}
func (m *Sequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequence.Merge(m, src)
}
func (m *Sequence) XXX_Size() int {
	return m.Size()
}
func (m *Sequence) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequence.DiscardUnknown(m)
}

var xxx_messageInfo_Sequence proto.InternalMessageInfo

func (m *Sequence) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sequence) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Sequence)(nil), "cosmos_sdk.x.ibc.commitment.v1.Sequence")
}

func init() { proto.RegisterFile("commitment/types.proto", fileDescriptor_921e8bf76465e275) }

var fileDescriptor_921e8bf76465e275 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0xcd,
	0xcd, 0x2c, 0xc9, 0x4d, 0xcd, 0x2b, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x2f, 0x4e, 0xc9, 0xd6, 0xab,
	0xd0, 0xcb, 0x4c, 0x4a, 0xd6, 0x43, 0x28, 0xd4, 0x2b, 0x33, 0x54, 0xb2, 0xe3, 0xe2, 0x08, 0x4e,
	0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x70, 0x84, 0x64, 0xb8, 0x38, 0x4b, 0x32, 0x73, 0x53,
	0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x10, 0x02, 0x4e, 0xee,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0x04, 0xb2, 0x4f, 0x3f, 0x25, 0xb1, 0x24, 0x31, 0x39, 0x23, 0x31, 0x33, 0x2f, 0x27, 0x31,
	0x49, 0x3f, 0x2d, 0x31, 0xa9, 0x28, 0x33, 0x59, 0x37, 0x33, 0x29, 0x59, 0x1f, 0xe1, 0x96, 0x24,
	0x36, 0xb0, 0x7b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x9e, 0x1a, 0x3e, 0xc9, 0x00,
	0x00, 0x00,
}

func (m *Sequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovTypes(uint64(m.Value))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTypes(uint64(m.Timestamp))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Sequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
