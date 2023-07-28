// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: OmniFlix/marketplace/v1beta1/listing.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Listing struct {
	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NftId       string            `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty" yaml:"nft_id"`
	DenomId     string            `protobuf:"bytes,3,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	Price       types.Coin        `protobuf:"bytes,4,opt,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"price"`
	Owner       string            `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	SplitShares []WeightedAddress `protobuf:"bytes,6,rep,name=split_shares,json=splitShares,proto3" json:"split_shares" yaml:"split_shares"`
}

func (m *Listing) Reset()         { *m = Listing{} }
func (m *Listing) String() string { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()    {}
func (*Listing) Descriptor() ([]byte, []int) {
	return fileDescriptor_300685922a0a2983, []int{0}
}
func (m *Listing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listing.Merge(m, src)
}
func (m *Listing) XXX_Size() int {
	return m.Size()
}
func (m *Listing) XXX_DiscardUnknown() {
	xxx_messageInfo_Listing.DiscardUnknown(m)
}

var xxx_messageInfo_Listing proto.InternalMessageInfo

type WeightedAddress struct {
	Address string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedAddress) Reset()         { *m = WeightedAddress{} }
func (m *WeightedAddress) String() string { return proto.CompactTextString(m) }
func (*WeightedAddress) ProtoMessage()    {}
func (*WeightedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_300685922a0a2983, []int{1}
}
func (m *WeightedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAddress.Merge(m, src)
}
func (m *WeightedAddress) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAddress proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Listing)(nil), "OmniFlix.marketplace.v1beta1.Listing")
	proto.RegisterType((*WeightedAddress)(nil), "OmniFlix.marketplace.v1beta1.WeightedAddress")
}

func init() {
	proto.RegisterFile("OmniFlix/marketplace/v1beta1/listing.proto", fileDescriptor_300685922a0a2983)
}

var fileDescriptor_300685922a0a2983 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x4e, 0xb6, 0xdd, 0x2d, 0x78, 0x4b, 0x0b, 0x2e, 0x87, 0x50, 0x50, 0xb2, 0xca, 0xa1, 0xac,
	0x10, 0xb5, 0xd5, 0x45, 0xe2, 0xd0, 0x0b, 0x22, 0x20, 0xa4, 0x4a, 0x48, 0x48, 0x01, 0xa9, 0x12,
	0x97, 0x92, 0xc4, 0xde, 0xac, 0xd5, 0xc4, 0x8e, 0x62, 0xf7, 0xef, 0x2d, 0x78, 0x02, 0xc4, 0x99,
	0x27, 0xd9, 0x63, 0x8f, 0x88, 0x43, 0x80, 0xdd, 0x0b, 0x07, 0x4e, 0x79, 0x02, 0x14, 0x3b, 0x4b,
	0x17, 0x0e, 0x88, 0x93, 0xc7, 0x33, 0xdf, 0x7c, 0xdf, 0xe7, 0xf1, 0x80, 0x07, 0xaf, 0x72, 0xce,
	0x5e, 0x64, 0xec, 0x1c, 0xe7, 0x51, 0x79, 0x4c, 0x55, 0x91, 0x45, 0x09, 0xc5, 0xa7, 0x7b, 0x31,
	0x55, 0xd1, 0x1e, 0xce, 0x98, 0x54, 0x8c, 0xa7, 0xa8, 0x28, 0x85, 0x12, 0xf0, 0xde, 0x02, 0x8b,
	0x96, 0xb0, 0xa8, 0xc5, 0x6e, 0xbb, 0x89, 0x90, 0xb9, 0x90, 0x38, 0x8e, 0xe4, 0x15, 0x41, 0x22,
	0x18, 0x37, 0xdd, 0xdb, 0xb7, 0x53, 0x91, 0x0a, 0x1d, 0xe2, 0x26, 0x32, 0x59, 0xff, 0x67, 0x07,
	0xac, 0xbd, 0x34, 0x2a, 0x70, 0x03, 0x74, 0x18, 0x71, 0xec, 0x81, 0x3d, 0xbc, 0x1e, 0x76, 0x18,
	0x81, 0x43, 0xd0, 0xe3, 0x63, 0x75, 0xc4, 0x88, 0xd3, 0x69, 0x72, 0xc1, 0xad, 0xba, 0xf2, 0x6e,
	0x5c, 0x44, 0x79, 0xb6, 0xef, 0x9b, 0xbc, 0x1f, 0x76, 0xf9, 0x58, 0x1d, 0x10, 0x88, 0xc0, 0x35,
	0x42, 0xb9, 0xc8, 0x1b, 0xec, 0x8a, 0xc6, 0x6e, 0xd5, 0x95, 0xb7, 0x69, 0xb0, 0x8b, 0x8a, 0x1f,
	0xae, 0xe9, 0xf0, 0x80, 0xc0, 0x77, 0xa0, 0x5b, 0x94, 0x2c, 0xa1, 0xce, 0xea, 0xc0, 0x1e, 0xf6,
	0x47, 0x77, 0x90, 0xf1, 0x8e, 0x1a, 0xef, 0x8b, 0x07, 0xa1, 0x67, 0x82, 0xf1, 0x00, 0x4f, 0x2b,
	0xcf, 0xfa, 0xf4, 0xd5, 0xbb, 0x9f, 0x32, 0x35, 0x39, 0x89, 0x51, 0x22, 0x72, 0xdc, 0x3e, 0xd4,
	0x1c, 0xbb, 0x92, 0x1c, 0x63, 0x75, 0x51, 0x50, 0xa9, 0x1b, 0x42, 0x43, 0x0c, 0x77, 0x40, 0x57,
	0x9c, 0x71, 0x5a, 0x3a, 0x5d, 0x6d, 0xe7, 0x66, 0x5d, 0x79, 0xeb, 0xc6, 0x8e, 0x4e, 0xfb, 0xa1,
	0x29, 0xc3, 0x1c, 0xac, 0xcb, 0x22, 0x63, 0xea, 0x48, 0x4e, 0xa2, 0x92, 0x4a, 0xa7, 0x37, 0x58,
	0x19, 0xf6, 0x47, 0xbb, 0xe8, 0x5f, 0xa3, 0x46, 0x87, 0x94, 0xa5, 0x13, 0x45, 0xc9, 0x53, 0x42,
	0x4a, 0x2a, 0x65, 0x70, 0xb7, 0x31, 0x59, 0x57, 0xde, 0x96, 0x51, 0x58, 0x26, 0xf4, 0xc3, 0xbe,
	0xbe, 0xbe, 0x36, 0xb7, 0x0f, 0x36, 0xd8, 0xfc, 0xab, 0x1b, 0x3e, 0x04, 0x6b, 0x91, 0x09, 0xcd,
	0xec, 0x03, 0x58, 0x57, 0xde, 0x86, 0xa1, 0x6a, 0x0b, 0x7e, 0xb8, 0x80, 0xc0, 0x43, 0xd0, 0x3b,
	0xd3, 0x04, 0xed, 0xa7, 0x3c, 0x69, 0xb4, 0xbf, 0x54, 0xde, 0xce, 0x7f, 0x0c, 0xe8, 0x39, 0x4d,
	0xae, 0xbe, 0xd0, 0xb0, 0xf8, 0x61, 0x4b, 0xb7, 0xbf, 0xfa, 0xe3, 0xa3, 0x67, 0x07, 0x6f, 0xa6,
	0xdf, 0x5d, 0x6b, 0x3a, 0x73, 0xed, 0xcb, 0x99, 0x6b, 0x7f, 0x9b, 0xb9, 0xf6, 0xfb, 0xb9, 0x6b,
	0x5d, 0xce, 0x5d, 0xeb, 0xf3, 0xdc, 0xb5, 0xde, 0x3e, 0x5e, 0x12, 0xf9, 0xbd, 0xb8, 0x22, 0xe7,
	0x6c, 0x9c, 0xb1, 0xf3, 0xc9, 0x49, 0x8c, 0x4f, 0x47, 0xf8, 0xcf, 0x4d, 0xd6, 0xc2, 0x71, 0x4f,
	0x2f, 0xdb, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xf5, 0xb7, 0xf1, 0xee, 0x02, 0x00,
	0x00,
}

func (this *WeightedAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WeightedAddress)
	if !ok {
		that2, ok := that.(WeightedAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !this.Weight.Equal(that1.Weight) {
		return false
	}
	return true
}
func (m *Listing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Listing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SplitShares) > 0 {
		for iNdEx := len(m.SplitShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SplitShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListing(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintListing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WeightedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintListing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListing(dAtA []byte, offset int, v uint64) int {
	offset -= sovListing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Listing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovListing(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	if len(m.SplitShares) > 0 {
		for _, e := range m.SplitShares {
			l = e.Size()
			n += 1 + l + sovListing(uint64(l))
		}
	}
	return n
}

func (m *WeightedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovListing(uint64(l))
	return n
}

func sovListing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListing(x uint64) (n int) {
	return sovListing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: Listing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SplitShares = append(m.SplitShares, WeightedAddress{})
			if err := m.SplitShares[len(m.SplitShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func (m *WeightedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: WeightedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func skipListing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
				return 0, ErrInvalidLengthListing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListing = fmt.Errorf("proto: unexpected end of group")
)