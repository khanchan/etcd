// Code generated by protoc-gen-gogo.
// source: lease.proto
// DO NOT EDIT!

/*
	Package leasepb is a generated protocol buffer package.

	It is generated from these files:
		lease.proto

	It has these top-level messages:
		Lease
*/
package leasepb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Lease struct {
	ID  int64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	TTL int64 `protobuf:"varint,2,opt,name=TTL,json=tTL,proto3" json:"TTL,omitempty"`
}

func (m *Lease) Reset()                    { *m = Lease{} }
func (m *Lease) String() string            { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()               {}
func (*Lease) Descriptor() ([]byte, []int) { return fileDescriptorLease, []int{0} }

func init() {
	proto.RegisterType((*Lease)(nil), "leasepb.Lease")
}
func (m *Lease) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Lease) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintLease(data, i, uint64(m.ID))
	}
	if m.TTL != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintLease(data, i, uint64(m.TTL))
	}
	return i, nil
}

func encodeFixed64Lease(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Lease(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLease(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Lease) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLease(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovLease(uint64(m.TTL))
	}
	return n
}

func sovLease(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLease(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
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
func skipLease(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLease
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLease
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLease(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLease = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorLease = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x49, 0x4d, 0x2c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x73, 0x0a, 0x92, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x62, 0xfa, 0x20, 0x16, 0x44, 0x5a, 0x49, 0x93, 0x8b, 0xd5, 0x07, 0xa4,
	0x40, 0x88, 0x8f, 0x8b, 0xc9, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0xd3,
	0x45, 0x48, 0x80, 0x8b, 0x39, 0x24, 0xc4, 0x47, 0x82, 0x09, 0x2c, 0xc0, 0x5c, 0x12, 0xe2, 0xe3,
	0x24, 0x72, 0xe2, 0xa1, 0x1c, 0xc3, 0x85, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x12, 0x1b, 0xd8, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8c, 0x1d, 0x92, 0xdc, 0x75, 0x00, 0x00, 0x00,
}
