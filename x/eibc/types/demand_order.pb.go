// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/eibc/demand_order.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/dymensionxyz/dymension/v3/x/common/types"
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

type DemandOrder struct {
	// id is a hash of the form generated by GetRollappPacketKey,
	// e.g status/rollappid/packetProofHeight/packetDestinationChannel-PacketSequence which gurantees uniqueness
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// tracking_packet_key is the key of the packet that is being tracked.
	// This key can change depends on the packet status.
	TrackingPacketKey    string                                   `protobuf:"bytes,2,opt,name=tracking_packet_key,json=trackingPacketKey,proto3" json:"tracking_packet_key,omitempty"`
	Price                github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	Fee                  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee"`
	Recipient            string                                   `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	IsFullfilled         bool                                     `protobuf:"varint,6,opt,name=is_fullfilled,json=isFullfilled,proto3" json:"is_fullfilled,omitempty"`
	RollappId            string                                   `protobuf:"bytes,7,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	Type                 types1.Type                              `protobuf:"varint,8,opt,name=type,proto3,enum=dymensionxyz.dymension.common.Type" json:"type,omitempty"`
	TrackingPacketStatus types1.Status                            `protobuf:"varint,9,opt,name=tracking_packet_status,json=trackingPacketStatus,proto3,enum=dymensionxyz.dymension.common.Status" json:"tracking_packet_status,omitempty"`
}

func (m *DemandOrder) Reset()         { *m = DemandOrder{} }
func (m *DemandOrder) String() string { return proto.CompactTextString(m) }
func (*DemandOrder) ProtoMessage()    {}
func (*DemandOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_3808f42eed32f331, []int{0}
}
func (m *DemandOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemandOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemandOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemandOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemandOrder.Merge(m, src)
}
func (m *DemandOrder) XXX_Size() int {
	return m.Size()
}
func (m *DemandOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_DemandOrder.DiscardUnknown(m)
}

var xxx_messageInfo_DemandOrder proto.InternalMessageInfo

func (m *DemandOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DemandOrder) GetTrackingPacketKey() string {
	if m != nil {
		return m.TrackingPacketKey
	}
	return ""
}

func (m *DemandOrder) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *DemandOrder) GetFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *DemandOrder) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *DemandOrder) GetIsFullfilled() bool {
	if m != nil {
		return m.IsFullfilled
	}
	return false
}

func (m *DemandOrder) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *DemandOrder) GetType() types1.Type {
	if m != nil {
		return m.Type
	}
	return types1.Type_ON_RECV
}

func (m *DemandOrder) GetTrackingPacketStatus() types1.Status {
	if m != nil {
		return m.TrackingPacketStatus
	}
	return types1.Status_PENDING
}

func init() {
	proto.RegisterType((*DemandOrder)(nil), "dymensionxyz.dymension.eibc.DemandOrder")
}

func init() { proto.RegisterFile("dymension/eibc/demand_order.proto", fileDescriptor_3808f42eed32f331) }

var fileDescriptor_3808f42eed32f331 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xda, 0x6e, 0xac, 0x1e, 0x4c, 0xc2, 0x4c, 0xc8, 0x6c, 0x2c, 0x2b, 0x4c, 0x48, 0xb9,
	0x60, 0xd3, 0xed, 0xc0, 0x7d, 0x20, 0x24, 0xb4, 0x03, 0x28, 0x70, 0x02, 0xa1, 0x28, 0xb1, 0x5f,
	0x8b, 0xd5, 0x24, 0xb6, 0x62, 0x77, 0x5a, 0xf8, 0x03, 0x5c, 0xf9, 0x1d, 0xfc, 0x92, 0x1d, 0x77,
	0xe4, 0x04, 0xa8, 0xfd, 0x23, 0x28, 0x76, 0x68, 0x0b, 0x08, 0x71, 0xe1, 0x94, 0xf8, 0x7d, 0xdf,
	0xfb, 0xbe, 0x67, 0xbf, 0x0f, 0xdd, 0x13, 0x75, 0x01, 0xa5, 0x91, 0xaa, 0x64, 0x20, 0x33, 0xce,
	0x04, 0x14, 0x69, 0x29, 0x12, 0x55, 0x09, 0xa8, 0xa8, 0xae, 0x94, 0x55, 0x78, 0x7f, 0x49, 0xb9,
	0xa8, 0x3f, 0xd0, 0xe5, 0x81, 0x36, 0xfc, 0xbd, 0xdd, 0x89, 0x9a, 0x28, 0xc7, 0x63, 0xcd, 0x9f,
	0x6f, 0xd9, 0x3b, 0x58, 0xa9, 0x72, 0x55, 0x14, 0xaa, 0x64, 0xc6, 0xa6, 0x76, 0x66, 0x5a, 0x78,
	0xff, 0x0f, 0xd8, 0xd6, 0x1a, 0x5a, 0x30, 0xe4, 0xca, 0x14, 0xca, 0xb0, 0x2c, 0x35, 0xc0, 0xce,
	0x47, 0x19, 0xd8, 0x74, 0xc4, 0xb8, 0x92, 0xa5, 0xc7, 0xef, 0x7f, 0xec, 0xa3, 0xed, 0xa7, 0x6e,
	0xca, 0x17, 0xcd, 0x90, 0x78, 0x07, 0x75, 0xa5, 0x20, 0xc1, 0x30, 0x88, 0x06, 0x71, 0x57, 0x0a,
	0x4c, 0xd1, 0x2d, 0x5b, 0xa5, 0x7c, 0x2a, 0xcb, 0x49, 0xa2, 0x53, 0x3e, 0x05, 0x9b, 0x4c, 0xa1,
	0x26, 0x5d, 0x47, 0xb8, 0xf9, 0x13, 0x7a, 0xe9, 0x90, 0x33, 0xa8, 0x71, 0x8a, 0x36, 0x74, 0x25,
	0x39, 0x90, 0xde, 0xb0, 0x17, 0x6d, 0x1f, 0xdf, 0xa1, 0xde, 0x9f, 0x36, 0xfe, 0xb4, 0xf5, 0xa7,
	0x4f, 0x94, 0x2c, 0x4f, 0x1f, 0x5d, 0x7e, 0x3d, 0xec, 0x7c, 0xfe, 0x76, 0x18, 0x4d, 0xa4, 0x7d,
	0x3f, 0xcb, 0x28, 0x57, 0x05, 0x6b, 0x87, 0xf5, 0x9f, 0x87, 0x46, 0x4c, 0xdd, 0x55, 0x8c, 0x6b,
	0x30, 0xb1, 0x57, 0xc6, 0xef, 0x50, 0x6f, 0x0c, 0x40, 0xfa, 0xff, 0xdf, 0xa0, 0xd1, 0xc5, 0x77,
	0xd1, 0xa0, 0x02, 0x2e, 0xb5, 0x84, 0xd2, 0x92, 0x0d, 0x77, 0xcf, 0x55, 0x01, 0x1f, 0xa1, 0x1b,
	0xd2, 0x24, 0xe3, 0x59, 0x9e, 0x8f, 0x65, 0x9e, 0x83, 0x20, 0x9b, 0xc3, 0x20, 0xda, 0x8a, 0xaf,
	0x4b, 0xf3, 0x6c, 0x59, 0xc3, 0x07, 0x08, 0x55, 0x2a, 0xcf, 0x53, 0xad, 0x13, 0x29, 0xc8, 0xb5,
	0x56, 0xc3, 0x57, 0x9e, 0x0b, 0xfc, 0x18, 0xf5, 0x1b, 0x57, 0xb2, 0x35, 0x0c, 0xa2, 0x9d, 0xe3,
	0x23, 0xfa, 0x97, 0x44, 0xf8, 0x65, 0xd2, 0xd7, 0xb5, 0x86, 0xd8, 0x35, 0xe0, 0xb7, 0xe8, 0xf6,
	0xef, 0xcb, 0xf0, 0x49, 0x20, 0x03, 0x27, 0xf5, 0xe0, 0x1f, 0x52, 0xaf, 0x1c, 0x39, 0xde, 0xfd,
	0x75, 0x6d, 0xbe, 0x7a, 0x7a, 0x76, 0x39, 0x0f, 0x83, 0xab, 0x79, 0x18, 0x7c, 0x9f, 0x87, 0xc1,
	0xa7, 0x45, 0xd8, 0xb9, 0x5a, 0x84, 0x9d, 0x2f, 0x8b, 0xb0, 0xf3, 0x66, 0xb4, 0xf6, 0x80, 0xeb,
	0x06, 0xab, 0x03, 0x3b, 0x3f, 0x61, 0x17, 0x3e, 0xf2, 0xee, 0x3d, 0xb3, 0x4d, 0x97, 0xae, 0x93,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0xa7, 0x9f, 0x00, 0x11, 0x03, 0x00, 0x00,
}

func (m *DemandOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemandOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemandOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TrackingPacketStatus != 0 {
		i = encodeVarintDemandOrder(dAtA, i, uint64(m.TrackingPacketStatus))
		i--
		dAtA[i] = 0x48
	}
	if m.Type != 0 {
		i = encodeVarintDemandOrder(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x40
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x3a
	}
	if m.IsFullfilled {
		i--
		if m.IsFullfilled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemandOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemandOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TrackingPacketKey) > 0 {
		i -= len(m.TrackingPacketKey)
		copy(dAtA[i:], m.TrackingPacketKey)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.TrackingPacketKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDemandOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovDemandOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DemandOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	l = len(m.TrackingPacketKey)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovDemandOrder(uint64(l))
		}
	}
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovDemandOrder(uint64(l))
		}
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	if m.IsFullfilled {
		n += 2
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovDemandOrder(uint64(m.Type))
	}
	if m.TrackingPacketStatus != 0 {
		n += 1 + sovDemandOrder(uint64(m.TrackingPacketStatus))
	}
	return n
}

func sovDemandOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDemandOrder(x uint64) (n int) {
	return sovDemandOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DemandOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemandOrder
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
			return fmt.Errorf("proto: DemandOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemandOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackingPacketKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackingPacketKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFullfilled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
			m.IsFullfilled = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= types1.Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackingPacketStatus", wireType)
			}
			m.TrackingPacketStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrackingPacketStatus |= types1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDemandOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDemandOrder
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
func skipDemandOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemandOrder
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
					return 0, ErrIntOverflowDemandOrder
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
					return 0, ErrIntOverflowDemandOrder
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
				return 0, ErrInvalidLengthDemandOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDemandOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDemandOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDemandOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemandOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDemandOrder = fmt.Errorf("proto: unexpected end of group")
)
