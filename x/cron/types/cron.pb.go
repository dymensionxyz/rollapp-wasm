// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasmrollapp/cron/v1beta1/cron.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type CronJob struct {
	// id is the unique identifier for the cron job
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the name of the cron job
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// description is the description of the cron job
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Msgs that will be executed every period amount of time
	MsgContractCron []MsgContractCron `protobuf:"bytes,4,rep,name=msg_contract_cron,json=msgContractCron,proto3" json:"msg_contract_cron"`
}

func (m *CronJob) Reset()         { *m = CronJob{} }
func (m *CronJob) String() string { return proto.CompactTextString(m) }
func (*CronJob) ProtoMessage()    {}
func (*CronJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6aacd87fd409ed, []int{0}
}
func (m *CronJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CronJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CronJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CronJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJob.Merge(m, src)
}
func (m *CronJob) XXX_Size() int {
	return m.Size()
}
func (m *CronJob) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJob.DiscardUnknown(m)
}

var xxx_messageInfo_CronJob proto.InternalMessageInfo

func (m *CronJob) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CronJob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CronJob) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CronJob) GetMsgContractCron() []MsgContractCron {
	if m != nil {
		return m.MsgContractCron
	}
	return nil
}

type MsgContractCron struct {
	// Contract is the address of the smart contract
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// Msg is json encoded message to be passed to the contract
	JsonMsg string `protobuf:"bytes,2,opt,name=json_msg,json=jsonMsg,proto3" json:"json_msg,omitempty"`
}

func (m *MsgContractCron) Reset()         { *m = MsgContractCron{} }
func (m *MsgContractCron) String() string { return proto.CompactTextString(m) }
func (*MsgContractCron) ProtoMessage()    {}
func (*MsgContractCron) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6aacd87fd409ed, []int{1}
}
func (m *MsgContractCron) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgContractCron) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgContractCron.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgContractCron) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgContractCron.Merge(m, src)
}
func (m *MsgContractCron) XXX_Size() int {
	return m.Size()
}
func (m *MsgContractCron) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgContractCron.DiscardUnknown(m)
}

var xxx_messageInfo_MsgContractCron proto.InternalMessageInfo

func (m *MsgContractCron) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgContractCron) GetJsonMsg() string {
	if m != nil {
		return m.JsonMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*CronJob)(nil), "wasmrollapp.cron.v1beta1.CronJob")
	proto.RegisterType((*MsgContractCron)(nil), "wasmrollapp.cron.v1beta1.MsgContractCron")
}

func init() {
	proto.RegisterFile("wasmrollapp/cron/v1beta1/cron.proto", fileDescriptor_ad6aacd87fd409ed)
}

var fileDescriptor_ad6aacd87fd409ed = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xe3, 0x36, 0xa2, 0xd4, 0x95, 0x28, 0x58, 0x0c, 0x81, 0x21, 0x44, 0x65, 0x69, 0x07,
	0x12, 0x15, 0x4e, 0x40, 0xbb, 0x81, 0xba, 0x64, 0x41, 0x82, 0xa1, 0xca, 0x1f, 0xcb, 0x18, 0xd5,
	0x7e, 0x91, 0x6d, 0xa0, 0xe5, 0x14, 0x1c, 0x84, 0x83, 0x74, 0xec, 0xc8, 0x84, 0x50, 0x7b, 0x11,
	0x14, 0x27, 0xa0, 0x82, 0xc4, 0xf6, 0xde, 0x2f, 0xbf, 0xc8, 0xdf, 0x67, 0xe3, 0xd3, 0xe7, 0x44,
	0x0b, 0x05, 0xb3, 0x59, 0x52, 0x14, 0x51, 0xa6, 0x40, 0x46, 0x4f, 0xc3, 0x94, 0x9a, 0x64, 0x68,
	0x97, 0xb0, 0x50, 0x60, 0x80, 0x78, 0x5b, 0x52, 0x68, 0x79, 0x2d, 0x1d, 0x1f, 0x32, 0x60, 0x60,
	0xa5, 0xa8, 0x9c, 0x2a, 0xbf, 0xf7, 0x86, 0x70, 0x6b, 0xac, 0x40, 0x5e, 0x41, 0x4a, 0xf6, 0x70,
	0x83, 0xe7, 0x1e, 0x0a, 0x50, 0xdf, 0x8d, 0x1b, 0x3c, 0x27, 0x04, 0xbb, 0x32, 0x11, 0xd4, 0x6b,
	0x04, 0xa8, 0xdf, 0x8e, 0xed, 0x4c, 0x02, 0xdc, 0xc9, 0xa9, 0xce, 0x14, 0x2f, 0x0c, 0x07, 0xe9,
	0x35, 0xed, 0xa7, 0x6d, 0x44, 0xee, 0xf0, 0x81, 0xd0, 0x6c, 0x9a, 0x81, 0x34, 0x2a, 0xc9, 0xcc,
	0xb4, 0x0c, 0xe1, 0xb9, 0x41, 0xb3, 0xdf, 0x39, 0x1f, 0x84, 0xff, 0xa5, 0x0b, 0x27, 0x9a, 0x8d,
	0xeb, 0x3f, 0xca, 0x38, 0x23, 0x77, 0xf9, 0x71, 0xe2, 0xc4, 0x5d, 0xf1, 0x1b, 0xf7, 0x6e, 0x70,
	0xf7, 0x8f, 0x49, 0x06, 0x78, 0xff, 0xe7, 0xac, 0x24, 0xcf, 0x15, 0xd5, 0xda, 0x76, 0x68, 0xc7,
	0xdd, 0x6f, 0x7e, 0x59, 0x61, 0x72, 0x84, 0x77, 0x1f, 0x34, 0xc8, 0xa9, 0xd0, 0xac, 0x2e, 0xd5,
	0x2a, 0xf7, 0x89, 0x66, 0xa3, 0xeb, 0xe5, 0xda, 0x47, 0xab, 0xb5, 0x8f, 0x3e, 0xd7, 0x3e, 0x7a,
	0xdd, 0xf8, 0xce, 0x6a, 0xe3, 0x3b, 0xef, 0x1b, 0xdf, 0xb9, 0x1d, 0x32, 0x6e, 0xee, 0x1f, 0xd3,
	0x30, 0x03, 0x11, 0xe5, 0x0b, 0x41, 0xa5, 0xe6, 0x20, 0xe7, 0x8b, 0x97, 0xa8, 0xee, 0x71, 0x56,
	0x76, 0x8a, 0xe6, 0xd5, 0x8b, 0x98, 0x45, 0x41, 0x75, 0xba, 0x63, 0xef, 0xf6, 0xe2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x9c, 0xf2, 0x95, 0xb2, 0x01, 0x00, 0x00,
}

func (m *CronJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CronJob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgContractCron) > 0 {
		for iNdEx := len(m.MsgContractCron) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MsgContractCron[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCron(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCron(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCron(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCron(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgContractCron) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgContractCron) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgContractCron) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JsonMsg) > 0 {
		i -= len(m.JsonMsg)
		copy(dAtA[i:], m.JsonMsg)
		i = encodeVarintCron(dAtA, i, uint64(len(m.JsonMsg)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCron(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCron(dAtA []byte, offset int, v uint64) int {
	offset -= sovCron(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CronJob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCron(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCron(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCron(uint64(l))
	}
	if len(m.MsgContractCron) > 0 {
		for _, e := range m.MsgContractCron {
			l = e.Size()
			n += 1 + l + sovCron(uint64(l))
		}
	}
	return n
}

func (m *MsgContractCron) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCron(uint64(l))
	}
	l = len(m.JsonMsg)
	if l > 0 {
		n += 1 + l + sovCron(uint64(l))
	}
	return n
}

func sovCron(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCron(x uint64) (n int) {
	return sovCron(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CronJob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCron
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
			return fmt.Errorf("proto: CronJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CronJob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgContractCron", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgContractCron = append(m.MsgContractCron, MsgContractCron{})
			if err := m.MsgContractCron[len(m.MsgContractCron)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCron(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCron
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
func (m *MsgContractCron) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCron
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
			return fmt.Errorf("proto: MsgContractCron: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgContractCron: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsonMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCron(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCron
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
func skipCron(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCron
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
					return 0, ErrIntOverflowCron
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
					return 0, ErrIntOverflowCron
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
				return 0, ErrInvalidLengthCron
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCron
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCron
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCron        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCron          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCron = fmt.Errorf("proto: unexpected end of group")
)
