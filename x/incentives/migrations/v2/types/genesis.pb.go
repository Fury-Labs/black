// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: black/V2Incentives/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// V2GenesisState defines the module's genesis state.
type V2GenesisState struct {
	// V2Params are the V2Incentives module parameters
	V2Params V2Params `protobuf:"bytes,1,opt,name=V2Params,proto3" json:"V2Params"`
	// V2Incentives is a slice of active V2Incentives
	V2Incentives []V2Incentive `protobuf:"bytes,2,rep,name=V2Incentives,proto3" json:"V2Incentives"`
	// gas_meters is a slice of active V2GasMeters
	V2GasMeters []V2GasMeter `protobuf:"bytes,3,rep,name=gas_meters,json=V2GasMeters,proto3" json:"gas_meters"`
}

func (m *V2GenesisState) Reset()         { *m = V2GenesisState{} }
func (m *V2GenesisState) String() string { return proto.CompactTextString(m) }
func (*V2GenesisState) ProtoMessage()    {}
func (*V2GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb1f7c7e8ad160b, []int{0}
}

func (m *V2GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2GenesisState.Merge(m, src)
}

func (m *V2GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *V2GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_V2GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_V2GenesisState proto.InternalMessageInfo

func (m *V2GenesisState) GetV2Params() V2Params {
	if m != nil {
		return m.V2Params
	}
	return V2Params{}
}

func (m *V2GenesisState) GetV2Incentives() []V2Incentive {
	if m != nil {
		return m.V2Incentives
	}
	return nil
}

func (m *V2GenesisState) GetV2GasMeters() []V2GasMeter {
	if m != nil {
		return m.V2GasMeters
	}
	return nil
}

// V2Params defines the V2Incentives module V2Params
type V2Params struct {
	// enable_V2Incentives is the parameter to enable V2Incentives
	EnableIncentives bool `protobuf:"varint,1,opt,name=enable_V2Incentives,json=EnableIncentives,proto3" json:"enable_V2Incentives,omitempty"`
	// allocation_limit is the maximum percentage an V2Incentive can allocate per denomination
	AllocationLimit github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=allocation_limit,json=allocationLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"allocation_limit"`
	// V2Incentives_epoch_identifier for the epochs module hooks
	IncentivesEpochIdentifier string `protobuf:"bytes,3,opt,name=V2Incentives_epoch_identifier,json=IncentivesEpochIdentifier,proto3" json:"V2Incentives_epoch_identifier,omitempty"`
	// reward_scaler is the scaling factor for capping rewards
	RewardScaler github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reward_scaler,json=rewardScaler,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_scaler"`
}

func (m *V2Params) Reset()         { *m = V2Params{} }
func (m *V2Params) String() string { return proto.CompactTextString(m) }
func (*V2Params) ProtoMessage()    {}
func (*V2Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb1f7c7e8ad160b, []int{1}
}

func (m *V2Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2Params.Merge(m, src)
}

func (m *V2Params) XXX_Size() int {
	return m.Size()
}

func (m *V2Params) XXX_DiscardUnknown() {
	xxx_messageInfo_V2Params.DiscardUnknown(m)
}

var xxx_messageInfo_V2Params proto.InternalMessageInfo

func (m *V2Params) GetEnableIncentives() bool {
	if m != nil {
		return m.EnableIncentives
	}
	return false
}

func (m *V2Params) GetIncentivesEpochIdentifier() string {
	if m != nil {
		return m.IncentivesEpochIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*V2GenesisState)(nil), "black.V2Incentives.v1.V2GenesisState")
	proto.RegisterType((*V2Params)(nil), "black.V2Incentives.v1.V2Params")
}

func init() {
	proto.RegisterFile("black/V2Incentives/v1/genesis.proto", fileDescriptor_7bb1f7c7e8ad160b)
}

var fileDescriptor_7bb1f7c7e8ad160b = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0xee, 0xd2, 0x40,
	0x10, 0x87, 0x5b, 0xf8, 0x87, 0xc8, 0x82, 0x11, 0xab, 0x87, 0x0a, 0xb1, 0x20, 0x31, 0x86, 0xc4,
	0xb0, 0x15, 0x3c, 0x79, 0xf1, 0xd0, 0x60, 0x08, 0x89, 0x26, 0xa6, 0x9c, 0xf4, 0xd2, 0x2c, 0x65,
	0x2c, 0x1b, 0xdb, 0x6e, 0xd3, 0x5d, 0xab, 0xbe, 0x85, 0x2f, 0xe1, 0xbb, 0x70, 0x24, 0xf1, 0x62,
	0x3c, 0x10, 0x03, 0x2f, 0x62, 0x76, 0xb7, 0xda, 0x1e, 0x7a, 0xf2, 0xd2, 0xee, 0xb4, 0xdf, 0x7c,
	0xbf, 0x69, 0x33, 0xe8, 0x11, 0x14, 0x09, 0xe3, 0x2e, 0x4d, 0x43, 0x48, 0x05, 0x2d, 0x80, 0xbb,
	0xc5, 0xc2, 0x8d, 0x20, 0x05, 0x4e, 0x39, 0xce, 0x72, 0x26, 0x98, 0x75, 0x4f, 0x21, 0xb8, 0x42,
	0x70, 0xb1, 0x18, 0x3e, 0x6e, 0xea, 0xab, 0x21, 0xaa, 0x75, 0x78, 0x3f, 0x62, 0x11, 0x53, 0x47,
	0x57, 0x9e, 0xf4, 0xd3, 0xe9, 0x0f, 0x13, 0xf5, 0xd7, 0x3a, 0x62, 0x2b, 0x88, 0x00, 0xeb, 0x05,
	0xea, 0x64, 0x24, 0x27, 0x09, 0xb7, 0xcd, 0x89, 0x39, 0xeb, 0x2d, 0x47, 0xb8, 0x21, 0x12, 0xbf,
	0x55, 0x88, 0x77, 0x73, 0x3c, 0x8f, 0x0d, 0xbf, 0x6c, 0xb0, 0x56, 0x08, 0x55, 0x94, 0xdd, 0x9a,
	0xb4, 0x67, 0xbd, 0xa5, 0xd3, 0xd8, 0xbe, 0xf9, 0x5b, 0x95, 0x86, 0x5a, 0x9f, 0xe5, 0x21, 0x14,
	0x11, 0x1e, 0x24, 0x20, 0x20, 0xe7, 0x76, 0x5b, 0x59, 0x1e, 0x36, 0x5a, 0xd6, 0x84, 0xbf, 0x91,
	0x54, 0x29, 0xe9, 0x46, 0x65, 0xcd, 0xa7, 0xdf, 0x5b, 0xa8, 0xa3, 0x47, 0xb4, 0x9e, 0xa2, 0xbb,
	0x90, 0x92, 0x5d, 0x0c, 0x41, 0x6d, 0x36, 0xf9, 0x69, 0xb7, 0xfc, 0x81, 0x7e, 0xb1, 0xa9, 0xb2,
	0xdf, 0xa1, 0x01, 0x89, 0x63, 0x16, 0x12, 0x41, 0x59, 0x1a, 0xc4, 0x34, 0xa1, 0xc2, 0x6e, 0x4d,
	0xcc, 0x59, 0xd7, 0xc3, 0x32, 0xe2, 0xd7, 0x79, 0xfc, 0x24, 0xa2, 0xe2, 0xf0, 0x69, 0x87, 0x43,
	0x96, 0xb8, 0x21, 0xe3, 0xf2, 0xbf, 0xeb, 0xdb, 0x9c, 0xef, 0x3f, 0xba, 0xe2, 0x6b, 0x06, 0x1c,
	0xaf, 0x20, 0xf4, 0xef, 0x54, 0x9e, 0xd7, 0x52, 0x63, 0xbd, 0x44, 0xa3, 0x6a, 0x80, 0x00, 0x32,
	0x16, 0x1e, 0x02, 0xba, 0x97, 0xf5, 0x07, 0x0a, 0xb9, 0xdd, 0x96, 0x29, 0xfe, 0x83, 0x0a, 0x79,
	0x25, 0x89, 0xcd, 0x3f, 0xc0, 0xda, 0xa2, 0xdb, 0x39, 0x7c, 0x26, 0xf9, 0x3e, 0xe0, 0x21, 0x89,
	0x21, 0xb7, 0x6f, 0xfe, 0x6b, 0xae, 0xbe, 0x96, 0x6c, 0x95, 0xc3, 0x5b, 0x1f, 0x2f, 0x8e, 0x79,
	0xba, 0x38, 0xe6, 0xef, 0x8b, 0x63, 0x7e, 0xbb, 0x3a, 0xc6, 0xe9, 0xea, 0x18, 0x3f, 0xaf, 0x8e,
	0xf1, 0x7e, 0x5e, 0xf3, 0xe9, 0xf5, 0xd2, 0xd7, 0x62, 0xf1, 0xcc, 0xfd, 0x52, 0x5f, 0x35, 0xa5,
	0xde, 0x75, 0xd4, 0x36, 0x3d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xce, 0xce, 0x10, 0xc3,
	0x02, 0x00, 0x00,
}

func (m *V2GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.V2GasMeters) > 0 {
		for iNdEx := len(m.V2GasMeters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.V2GasMeters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.V2Incentives) > 0 {
		for iNdEx := len(m.V2Incentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.V2Incentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.V2Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *V2Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RewardScaler.Size()
		i -= size
		if _, err := m.RewardScaler.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.IncentivesEpochIdentifier) > 0 {
		i -= len(m.IncentivesEpochIdentifier)
		copy(dAtA[i:], m.IncentivesEpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IncentivesEpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.AllocationLimit.Size()
		i -= size
		if _, err := m.AllocationLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableIncentives {
		i--
		if m.EnableIncentives {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *V2GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.V2Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.V2Incentives) > 0 {
		for _, e := range m.V2Incentives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.V2GasMeters) > 0 {
		for _, e := range m.V2GasMeters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *V2Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableIncentives {
		n += 2
	}
	l = m.AllocationLimit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.IncentivesEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.RewardScaler.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V2GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V2GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.V2Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2Incentives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V2Incentives = append(m.V2Incentives, V2Incentive{})
			if err := m.V2Incentives[len(m.V2Incentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2GasMeters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V2GasMeters = append(m.V2GasMeters, V2GasMeter{})
			if err := m.V2GasMeters[len(m.V2GasMeters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *V2Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V2Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableIncentives", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableIncentives = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AllocationLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentivesEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardScaler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardScaler.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
