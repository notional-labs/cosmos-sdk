// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol_x/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	accesscontrol "github.com/cosmos/cosmos-sdk/types/accesscontrol"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RegisterWasmDependencyJSONFile struct {
	ContractAddress       string                              `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address" yaml:"contract_address"`
	WasmDependencyMapping accesscontrol.WasmDependencyMapping `protobuf:"bytes,2,opt,name=wasm_dependency_mapping,json=wasmDependencyMapping,proto3" json:"wasm_dependency_mapping" yaml:"wasm_dependency_mapping"`
}

func (m *RegisterWasmDependencyJSONFile) Reset()         { *m = RegisterWasmDependencyJSONFile{} }
func (m *RegisterWasmDependencyJSONFile) String() string { return proto.CompactTextString(m) }
func (*RegisterWasmDependencyJSONFile) ProtoMessage()    {}
func (*RegisterWasmDependencyJSONFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a7b777fc99c34b4, []int{0}
}
func (m *RegisterWasmDependencyJSONFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterWasmDependencyJSONFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterWasmDependencyJSONFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterWasmDependencyJSONFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterWasmDependencyJSONFile.Merge(m, src)
}
func (m *RegisterWasmDependencyJSONFile) XXX_Size() int {
	return m.Size()
}
func (m *RegisterWasmDependencyJSONFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterWasmDependencyJSONFile.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterWasmDependencyJSONFile proto.InternalMessageInfo

func (m *RegisterWasmDependencyJSONFile) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *RegisterWasmDependencyJSONFile) GetWasmDependencyMapping() accesscontrol.WasmDependencyMapping {
	if m != nil {
		return m.WasmDependencyMapping
	}
	return accesscontrol.WasmDependencyMapping{}
}

type MsgRegisterWasmDependency struct {
	FromAddress           string                              `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address" yaml:"from_address"`
	ContractAddress       string                              `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address" yaml:"contract_address"`
	WasmDependencyMapping accesscontrol.WasmDependencyMapping `protobuf:"bytes,3,opt,name=wasm_dependency_mapping,json=wasmDependencyMapping,proto3" json:"wasm_dependency_mapping" yaml:"wasm_dependency_mapping"`
}

func (m *MsgRegisterWasmDependency) Reset()         { *m = MsgRegisterWasmDependency{} }
func (m *MsgRegisterWasmDependency) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterWasmDependency) ProtoMessage()    {}
func (*MsgRegisterWasmDependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a7b777fc99c34b4, []int{1}
}
func (m *MsgRegisterWasmDependency) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterWasmDependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterWasmDependency.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterWasmDependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterWasmDependency.Merge(m, src)
}
func (m *MsgRegisterWasmDependency) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterWasmDependency) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterWasmDependency.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterWasmDependency proto.InternalMessageInfo

func (m *MsgRegisterWasmDependency) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgRegisterWasmDependency) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgRegisterWasmDependency) GetWasmDependencyMapping() accesscontrol.WasmDependencyMapping {
	if m != nil {
		return m.WasmDependencyMapping
	}
	return accesscontrol.WasmDependencyMapping{}
}

type MsgRegisterWasmDependencyResponse struct {
}

func (m *MsgRegisterWasmDependencyResponse) Reset()         { *m = MsgRegisterWasmDependencyResponse{} }
func (m *MsgRegisterWasmDependencyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterWasmDependencyResponse) ProtoMessage()    {}
func (*MsgRegisterWasmDependencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a7b777fc99c34b4, []int{2}
}
func (m *MsgRegisterWasmDependencyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterWasmDependencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterWasmDependencyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterWasmDependencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterWasmDependencyResponse.Merge(m, src)
}
func (m *MsgRegisterWasmDependencyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterWasmDependencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterWasmDependencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterWasmDependencyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterWasmDependencyJSONFile)(nil), "cosmos.accesscontrol_x.v1beta1.RegisterWasmDependencyJSONFile")
	proto.RegisterType((*MsgRegisterWasmDependency)(nil), "cosmos.accesscontrol_x.v1beta1.MsgRegisterWasmDependency")
	proto.RegisterType((*MsgRegisterWasmDependencyResponse)(nil), "cosmos.accesscontrol_x.v1beta1.MsgRegisterWasmDependencyResponse")
}

func init() { proto.RegisterFile("cosmos/accesscontrol_x/tx.proto", fileDescriptor_2a7b777fc99c34b4) }

var fileDescriptor_2a7b777fc99c34b4 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x5b, 0x11, 0x9c, 0x0a, 0x4a, 0xfc, 0xd1, 0xba, 0xc8, 0xa4, 0x46, 0xd1, 0x45,
	0x30, 0x43, 0xb7, 0x27, 0x7b, 0xeb, 0x22, 0x1e, 0x8a, 0xab, 0x10, 0x0f, 0x42, 0x2f, 0xcb, 0x6c,
	0xf2, 0x1c, 0x83, 0x9b, 0x99, 0x31, 0x6f, 0x6a, 0x77, 0xff, 0x0b, 0x41, 0x10, 0x4f, 0xe2, 0x9f,
	0xd3, 0x63, 0x8f, 0x9e, 0xa2, 0xec, 0x5e, 0xa4, 0xc7, 0xfd, 0x0b, 0xa4, 0x9b, 0xac, 0xb8, 0x31,
	0xf1, 0xe0, 0x41, 0x4f, 0xc9, 0xbc, 0xf7, 0x7d, 0xf9, 0x7e, 0xf3, 0x19, 0x1e, 0xf5, 0x22, 0x8d,
	0xa9, 0x46, 0x2e, 0xa2, 0x08, 0x10, 0x23, 0xad, 0x6c, 0xa6, 0x47, 0x83, 0x31, 0xb7, 0xe3, 0xc0,
	0x64, 0xda, 0x6a, 0x97, 0x15, 0x82, 0xa0, 0x22, 0x08, 0xde, 0x6e, 0x0f, 0xc1, 0x8a, 0xed, 0xf6,
	0x55, 0xa9, 0xa5, 0x5e, 0x48, 0xf9, 0xd9, 0x5b, 0x31, 0xd5, 0xbe, 0x29, 0xb5, 0x96, 0x23, 0xe0,
	0xc2, 0x24, 0x5c, 0x28, 0xa5, 0xad, 0xb0, 0x89, 0x56, 0x58, 0x76, 0xef, 0x97, 0xa6, 0x43, 0x81,
	0xc0, 0xdf, 0x1c, 0x42, 0x36, 0xe1, 0xe5, 0xe7, 0xb8, 0x11, 0x32, 0x51, 0x0b, 0x71, 0xa9, 0xed,
	0xd4, 0x05, 0x5c, 0x3d, 0x95, 0xca, 0x3b, 0x0d, 0xbf, 0x22, 0x41, 0x01, 0x26, 0xa5, 0xb7, 0xff,
	0xb1, 0x45, 0x59, 0x08, 0x32, 0x41, 0x0b, 0xd9, 0x0b, 0x81, 0xe9, 0x23, 0x30, 0xa0, 0x62, 0x50,
	0xd1, 0x64, 0xff, 0xf9, 0xb3, 0xa7, 0x8f, 0x93, 0x11, 0xb8, 0x07, 0xf4, 0xf2, 0x62, 0x5a, 0x44,
	0x76, 0x20, 0xe2, 0x38, 0x03, 0xc4, 0x4d, 0xb2, 0x45, 0x3a, 0x17, 0x7a, 0xfc, 0x34, 0xf7, 0x7e,
	0xeb, 0xcd, 0x73, 0x6f, 0x63, 0x22, 0xd2, 0xd1, 0xae, 0x5f, 0xed, 0xf8, 0xe1, 0xa5, 0x65, 0x69,
	0xaf, 0xa8, 0xb8, 0xef, 0x09, 0xdd, 0x38, 0x12, 0x98, 0x0e, 0xe2, 0x9f, 0xbe, 0x83, 0x54, 0x18,
	0x93, 0x28, 0xb9, 0xd9, 0xda, 0x22, 0x9d, 0xf5, 0xee, 0x4e, 0x50, 0x47, 0x7c, 0xc9, 0x3b, 0x58,
	0xcd, 0xdc, 0x2f, 0x46, 0x7b, 0x77, 0x8f, 0x73, 0xcf, 0x99, 0xe7, 0x1e, 0x2b, 0x82, 0x34, 0x38,
	0xf8, 0xe1, 0xb5, 0xa3, 0xba, 0xf1, 0xdd, 0x73, 0xdf, 0x3f, 0x7b, 0x8e, 0xff, 0xb5, 0x45, 0x6f,
	0xf4, 0x51, 0xd6, 0xd3, 0x71, 0xf7, 0xe9, 0xc5, 0x97, 0x99, 0x4e, 0x2b, 0x44, 0xee, 0x9d, 0xe6,
	0xde, 0x4a, 0x7d, 0x9e, 0x7b, 0x57, 0x8a, 0x10, 0xbf, 0x56, 0xfd, 0x70, 0xfd, 0xec, 0xb8, 0xa4,
	0x50, 0x47, 0xb8, 0xf5, 0x0f, 0x08, 0xaf, 0xfd, 0x67, 0xc2, 0xb7, 0xe9, 0xad, 0x46, 0xc0, 0x21,
	0xa0, 0xd1, 0x0a, 0xa1, 0xfb, 0x89, 0xd0, 0xb5, 0x3e, 0x4a, 0xf7, 0x03, 0xa1, 0xd7, 0x1b, 0xee,
	0xe2, 0x61, 0xf0, 0xe7, 0xad, 0x0c, 0x1a, 0x5d, 0xda, 0x7b, 0x7f, 0x3d, 0xba, 0x0c, 0xd8, 0x7b,
	0x72, 0x3c, 0x65, 0xe4, 0x64, 0xca, 0xc8, 0xb7, 0x29, 0x23, 0xef, 0x66, 0xcc, 0x39, 0x99, 0x31,
	0xe7, 0xcb, 0x8c, 0x39, 0x07, 0x5d, 0x99, 0xd8, 0x57, 0x87, 0xc3, 0x20, 0xd2, 0x29, 0x2f, 0xb7,
	0xb1, 0x78, 0x3c, 0xc0, 0xf8, 0x35, 0x1f, 0x57, 0x96, 0xd8, 0x4e, 0x0c, 0xe0, 0xf0, 0xfc, 0x62,
	0x2f, 0x77, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x44, 0x21, 0x43, 0x6a, 0x8a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterWasmDependency(ctx context.Context, in *MsgRegisterWasmDependency, opts ...grpc.CallOption) (*MsgRegisterWasmDependencyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterWasmDependency(ctx context.Context, in *MsgRegisterWasmDependency, opts ...grpc.CallOption) (*MsgRegisterWasmDependencyResponse, error) {
	out := new(MsgRegisterWasmDependencyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Msg/RegisterWasmDependency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterWasmDependency(context.Context, *MsgRegisterWasmDependency) (*MsgRegisterWasmDependencyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterWasmDependency(ctx context.Context, req *MsgRegisterWasmDependency) (*MsgRegisterWasmDependencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWasmDependency not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterWasmDependency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterWasmDependency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterWasmDependency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Msg/RegisterWasmDependency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterWasmDependency(ctx, req.(*MsgRegisterWasmDependency))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.accesscontrol_x.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWasmDependency",
			Handler:    _Msg_RegisterWasmDependency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/accesscontrol_x/tx.proto",
}

func (m *RegisterWasmDependencyJSONFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterWasmDependencyJSONFile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterWasmDependencyJSONFile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.WasmDependencyMapping.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterWasmDependency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterWasmDependency) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterWasmDependency) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.WasmDependencyMapping.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterWasmDependencyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterWasmDependencyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterWasmDependencyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterWasmDependencyJSONFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.WasmDependencyMapping.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRegisterWasmDependency) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.WasmDependencyMapping.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRegisterWasmDependencyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterWasmDependencyJSONFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RegisterWasmDependencyJSONFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterWasmDependencyJSONFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmDependencyMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmDependencyMapping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterWasmDependency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterWasmDependency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterWasmDependency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmDependencyMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmDependencyMapping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterWasmDependencyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterWasmDependencyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterWasmDependencyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
