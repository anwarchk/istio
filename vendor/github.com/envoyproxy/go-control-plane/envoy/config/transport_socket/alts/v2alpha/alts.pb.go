// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/transport_socket/alts/v2alpha/alts.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration for ALTS transport socket. This provides Google's ALTS protocol to Envoy.
// https://cloud.google.com/security/encryption-in-transit/application-layer-transport-security/
type Alts struct {
	// The location of a handshaker service, this is usually 169.254.169.254:8080
	// on GCE.
	HandshakerService string `protobuf:"bytes,1,opt,name=handshaker_service,json=handshakerService,proto3" json:"handshaker_service,omitempty"`
	// The acceptable service accounts from peer, peers not in the list will be rejected in the
	// handshake validation step. If empty, no validation will be performed.
	PeerServiceAccounts  []string `protobuf:"bytes,2,rep,name=peer_service_accounts,json=peerServiceAccounts" json:"peer_service_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alts) Reset()         { *m = Alts{} }
func (m *Alts) String() string { return proto.CompactTextString(m) }
func (*Alts) ProtoMessage()    {}
func (*Alts) Descriptor() ([]byte, []int) {
	return fileDescriptor_alts_71af6698b5273eb9, []int{0}
}
func (m *Alts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Alts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Alts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Alts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alts.Merge(dst, src)
}
func (m *Alts) XXX_Size() int {
	return m.Size()
}
func (m *Alts) XXX_DiscardUnknown() {
	xxx_messageInfo_Alts.DiscardUnknown(m)
}

var xxx_messageInfo_Alts proto.InternalMessageInfo

func (m *Alts) GetHandshakerService() string {
	if m != nil {
		return m.HandshakerService
	}
	return ""
}

func (m *Alts) GetPeerServiceAccounts() []string {
	if m != nil {
		return m.PeerServiceAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Alts)(nil), "envoy.config.transport_socket.alts.v2alpha.Alts")
}
func (m *Alts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Alts) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HandshakerService) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAlts(dAtA, i, uint64(len(m.HandshakerService)))
		i += copy(dAtA[i:], m.HandshakerService)
	}
	if len(m.PeerServiceAccounts) > 0 {
		for _, s := range m.PeerServiceAccounts {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAlts(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Alts) Size() (n int) {
	var l int
	_ = l
	l = len(m.HandshakerService)
	if l > 0 {
		n += 1 + l + sovAlts(uint64(l))
	}
	if len(m.PeerServiceAccounts) > 0 {
		for _, s := range m.PeerServiceAccounts {
			l = len(s)
			n += 1 + l + sovAlts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAlts(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAlts(x uint64) (n int) {
	return sovAlts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Alts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlts
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Alts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Alts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HandshakerService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HandshakerService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerServiceAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerServiceAccounts = append(m.PeerServiceAccounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlts
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAlts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlts
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
					return 0, ErrIntOverflowAlts
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowAlts
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAlts
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAlts
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipAlts(dAtA[start:])
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
	ErrInvalidLengthAlts = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlts   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/transport_socket/alts/v2alpha/alts.proto", fileDescriptor_alts_71af6698b5273eb9)
}

var fileDescriptor_alts_71af6698b5273eb9 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x4f, 0xcc, 0x29, 0x29, 0xd6, 0x2f,
	0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x04, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4,
	0xc0, 0xda, 0xf4, 0x20, 0xda, 0xf4, 0xd0, 0xb5, 0xe9, 0x81, 0x55, 0x42, 0xb5, 0x49, 0x89, 0x97,
	0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x43, 0x94, 0x4a, 0xb8, 0x58,
	0x1c, 0x73, 0x4a, 0x8a, 0x85, 0x2c, 0xb8, 0x84, 0x32, 0x12, 0xf3, 0x52, 0x8a, 0x33, 0x12, 0xb3,
	0x53, 0x8b, 0xe2, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x9d, 0x38, 0x77, 0xbd, 0x3c, 0xc0, 0xcc, 0x52, 0xc4, 0xa4, 0xc0, 0x18, 0x24, 0x88, 0x50, 0x14,
	0x0c, 0x51, 0x23, 0x64, 0xc4, 0x25, 0x5a, 0x90, 0x8a, 0xd0, 0x13, 0x9f, 0x98, 0x9c, 0x9c, 0x5f,
	0x9a, 0x57, 0x52, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x19, 0x24, 0x0c, 0x92, 0x84, 0xaa, 0x75,
	0x84, 0x4a, 0x39, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x51, 0x4c, 0x65, 0x46, 0x49, 0x6c, 0x60, 0xe7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0xbb, 0xeb, 0xaf, 0x0c, 0x01, 0x00, 0x00,
}