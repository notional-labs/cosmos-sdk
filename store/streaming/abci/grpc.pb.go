// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/store/streaming/abci/grpc.proto

package abci

import (
	context "context"
	types "cosmossdk.io/store/types"
	fmt "fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// ListenEndBlockRequest is the request type for the ListenEndBlock RPC method
//
// Deprecated: Store v1 is deprecated as of v0.50.x, please use Store v2 types
// instead.
//
// Deprecated: Do not use.
type ListenFinalizeBlockRequest struct {
	Req *v1.FinalizeBlockRequest  `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	Res *v1.FinalizeBlockResponse `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
}

func (m *ListenFinalizeBlockRequest) Reset()         { *m = ListenFinalizeBlockRequest{} }
func (m *ListenFinalizeBlockRequest) String() string { return proto.CompactTextString(m) }
func (*ListenFinalizeBlockRequest) ProtoMessage()    {}
func (*ListenFinalizeBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b98083eb9315fb6, []int{0}
}
func (m *ListenFinalizeBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenFinalizeBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenFinalizeBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenFinalizeBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenFinalizeBlockRequest.Merge(m, src)
}
func (m *ListenFinalizeBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListenFinalizeBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenFinalizeBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenFinalizeBlockRequest proto.InternalMessageInfo

func (m *ListenFinalizeBlockRequest) GetReq() *v1.FinalizeBlockRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *ListenFinalizeBlockRequest) GetRes() *v1.FinalizeBlockResponse {
	if m != nil {
		return m.Res
	}
	return nil
}

// ListenEndBlockResponse is the response type for the ListenEndBlock RPC method
//
// Deprecated: Store v1 is deprecated as of v0.50.x, please use Store v2 types
// instead.
//
// Deprecated: Do not use.
type ListenFinalizeBlockResponse struct {
}

func (m *ListenFinalizeBlockResponse) Reset()         { *m = ListenFinalizeBlockResponse{} }
func (m *ListenFinalizeBlockResponse) String() string { return proto.CompactTextString(m) }
func (*ListenFinalizeBlockResponse) ProtoMessage()    {}
func (*ListenFinalizeBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b98083eb9315fb6, []int{1}
}
func (m *ListenFinalizeBlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenFinalizeBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenFinalizeBlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenFinalizeBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenFinalizeBlockResponse.Merge(m, src)
}
func (m *ListenFinalizeBlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListenFinalizeBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenFinalizeBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenFinalizeBlockResponse proto.InternalMessageInfo

// ListenCommitRequest is the request type for the ListenCommit RPC method
//
// Deprecated: Store v1 is deprecated as of v0.50.x, please use Store v2 types
// instead.
//
// Deprecated: Do not use.
type ListenCommitRequest struct {
	// explicitly pass in block height as ResponseCommit does not contain this
	// info
	BlockHeight int64                `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Res         *v1.CommitResponse   `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
	ChangeSet   []*types.StoreKVPair `protobuf:"bytes,3,rep,name=change_set,json=changeSet,proto3" json:"change_set,omitempty"`
}

func (m *ListenCommitRequest) Reset()         { *m = ListenCommitRequest{} }
func (m *ListenCommitRequest) String() string { return proto.CompactTextString(m) }
func (*ListenCommitRequest) ProtoMessage()    {}
func (*ListenCommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b98083eb9315fb6, []int{2}
}
func (m *ListenCommitRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenCommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenCommitRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenCommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenCommitRequest.Merge(m, src)
}
func (m *ListenCommitRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListenCommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenCommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenCommitRequest proto.InternalMessageInfo

func (m *ListenCommitRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ListenCommitRequest) GetRes() *v1.CommitResponse {
	if m != nil {
		return m.Res
	}
	return nil
}

func (m *ListenCommitRequest) GetChangeSet() []*types.StoreKVPair {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

// ListenCommitResponse is the response type for the ListenCommit RPC method
//
// Deprecated: Store v1 is deprecated as of v0.50.x, please use Store v2 types
// instead.
//
// Deprecated: Do not use.
type ListenCommitResponse struct {
}

func (m *ListenCommitResponse) Reset()         { *m = ListenCommitResponse{} }
func (m *ListenCommitResponse) String() string { return proto.CompactTextString(m) }
func (*ListenCommitResponse) ProtoMessage()    {}
func (*ListenCommitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b98083eb9315fb6, []int{3}
}
func (m *ListenCommitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenCommitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenCommitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenCommitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenCommitResponse.Merge(m, src)
}
func (m *ListenCommitResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListenCommitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenCommitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenCommitResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListenFinalizeBlockRequest)(nil), "cosmos.store.streaming.abci.ListenFinalizeBlockRequest")
	proto.RegisterType((*ListenFinalizeBlockResponse)(nil), "cosmos.store.streaming.abci.ListenFinalizeBlockResponse")
	proto.RegisterType((*ListenCommitRequest)(nil), "cosmos.store.streaming.abci.ListenCommitRequest")
	proto.RegisterType((*ListenCommitResponse)(nil), "cosmos.store.streaming.abci.ListenCommitResponse")
}

func init() {
	proto.RegisterFile("cosmos/store/streaming/abci/grpc.proto", fileDescriptor_7b98083eb9315fb6)
}

var fileDescriptor_7b98083eb9315fb6 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0x87, 0xeb, 0x44, 0x42, 0xc2, 0xbd, 0x93, 0x2f, 0x43, 0x94, 0x8b, 0xa2, 0xa6, 0x42, 0xa5,
	0x93, 0x43, 0xc2, 0x40, 0x81, 0x05, 0x5a, 0x09, 0x81, 0x60, 0x40, 0xa9, 0xc4, 0xc0, 0x52, 0x25,
	0xe1, 0x90, 0x5a, 0x6d, 0xe2, 0xd4, 0x36, 0x91, 0xe0, 0x09, 0x3a, 0xb2, 0xb0, 0xf0, 0x18, 0x3c,
	0x05, 0x63, 0x47, 0x46, 0xd4, 0xbe, 0x08, 0x8a, 0x4d, 0x4b, 0x23, 0xca, 0x9f, 0x8e, 0x39, 0xfa,
	0x7d, 0xe7, 0x7c, 0xf6, 0x71, 0xf0, 0x20, 0xe3, 0xb2, 0xe0, 0x32, 0x90, 0x8a, 0x0b, 0x08, 0xa4,
	0x12, 0x90, 0x14, 0xac, 0xcc, 0x83, 0x24, 0xcd, 0x58, 0x90, 0x8b, 0x2a, 0xa3, 0x95, 0xe0, 0x8a,
	0x93, 0x2b, 0x93, 0xa3, 0x3a, 0x47, 0x0f, 0x39, 0xda, 0xe4, 0xdc, 0x9b, 0x19, 0x2f, 0x40, 0xa5,
	0x6f, 0x95, 0xc1, 0xea, 0x30, 0x50, 0xef, 0x2b, 0x90, 0x06, 0x75, 0x6f, 0xb5, 0x46, 0xd4, 0x61,
	0x0a, 0x2a, 0x09, 0x83, 0x25, 0x93, 0x0a, 0xca, 0xa6, 0x85, 0x4e, 0xf5, 0x3f, 0x21, 0xec, 0xbe,
	0xd0, 0xb5, 0x27, 0xac, 0x4c, 0x96, 0xec, 0x03, 0x8c, 0x97, 0x3c, 0x5b, 0xc4, 0xb0, 0x7a, 0x07,
	0x52, 0x91, 0x11, 0xb6, 0x05, 0xac, 0x1c, 0xd4, 0x43, 0xc3, 0x6e, 0x34, 0xa0, 0xfb, 0x81, 0x7a,
	0x3e, 0xad, 0x43, 0x7a, 0x0a, 0x8a, 0x1b, 0x84, 0xdc, 0x6f, 0x48, 0xe9, 0x58, 0x9a, 0xbc, 0xfd,
	0x4f, 0x52, 0x56, 0xbc, 0x94, 0xd0, 0xa0, 0xf2, 0x81, 0xe5, 0xa0, 0xbe, 0x8f, 0xaf, 0x4e, 0x6a,
	0x99, 0x9c, 0x8e, 0x7c, 0x41, 0xf8, 0xd2, 0x64, 0x26, 0xbc, 0x28, 0x98, 0xda, 0x3b, 0xfb, 0xf8,
	0x22, 0x6d, 0xc2, 0xb3, 0x39, 0xb0, 0x7c, 0xae, 0xb4, 0xbc, 0x1d, 0x77, 0x75, 0xed, 0xa9, 0x2e,
	0x91, 0xe8, 0x58, 0xae, 0xf7, 0xbb, 0xdc, 0xbe, 0xe1, 0x91, 0x15, 0x79, 0x84, 0x71, 0x36, 0x4f,
	0xca, 0x1c, 0x66, 0x12, 0x94, 0x63, 0xf7, 0xec, 0x61, 0x37, 0xf2, 0x69, 0x6b, 0x3f, 0x3f, 0x2f,
	0x99, 0x4e, 0x9b, 0xaf, 0xe7, 0xaf, 0x5e, 0x26, 0x4c, 0xc4, 0xd7, 0x0d, 0x34, 0x05, 0xa5, 0xa5,
	0x5d, 0x7c, 0xa3, 0xed, 0xfc, 0xeb, 0x40, 0xd1, 0x67, 0x0b, 0x5f, 0x3e, 0x1e, 0x4f, 0x9e, 0x99,
	0x00, 0x88, 0x29, 0x88, 0x9a, 0x65, 0x40, 0xd6, 0x87, 0x83, 0xb6, 0x2e, 0x83, 0xdc, 0xa3, 0x7f,
	0x79, 0x1d, 0xf4, 0xcf, 0x5b, 0x75, 0x47, 0xe7, 0x83, 0x46, 0x93, 0x48, 0x7c, 0x71, 0xac, 0x4f,
	0xee, 0xfc, 0x47, 0xa7, 0xd6, 0x76, 0xdc, 0xf0, 0x0c, 0xc2, 0x0c, 0x75, 0xed, 0xb5, 0x85, 0xc6,
	0x0f, 0xbf, 0x6e, 0x3d, 0xb4, 0xd9, 0x7a, 0xe8, 0xfb, 0xd6, 0x43, 0x1f, 0x77, 0x5e, 0x67, 0xb3,
	0xf3, 0x3a, 0xdf, 0x76, 0x5e, 0xe7, 0xb5, 0x6f, 0x1a, 0xca, 0x37, 0x0b, 0xca, 0xf8, 0xc9, 0x3f,
	0x2a, 0xbd, 0xa6, 0x1f, 0xfb, 0xdd, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x60, 0x15, 0x1c,
	0x77, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ABCIListenerServiceClient is the client API for ABCIListenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type ABCIListenerServiceClient interface {
	// ListenFinalizeBlock is the corresponding endpoint for
	// ABCIListener.ListenEndBlock
	ListenFinalizeBlock(ctx context.Context, in *ListenFinalizeBlockRequest, opts ...grpc.CallOption) (*ListenFinalizeBlockResponse, error)
	// ListenCommit is the corresponding endpoint for ABCIListener.ListenCommit
	ListenCommit(ctx context.Context, in *ListenCommitRequest, opts ...grpc.CallOption) (*ListenCommitResponse, error)
}

type aBCIListenerServiceClient struct {
	cc grpc1.ClientConn
}

// Deprecated: Do not use.
func NewABCIListenerServiceClient(cc grpc1.ClientConn) ABCIListenerServiceClient {
	return &aBCIListenerServiceClient{cc}
}

func (c *aBCIListenerServiceClient) ListenFinalizeBlock(ctx context.Context, in *ListenFinalizeBlockRequest, opts ...grpc.CallOption) (*ListenFinalizeBlockResponse, error) {
	out := new(ListenFinalizeBlockResponse)
	err := c.cc.Invoke(ctx, "/cosmos.store.streaming.abci.ABCIListenerService/ListenFinalizeBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIListenerServiceClient) ListenCommit(ctx context.Context, in *ListenCommitRequest, opts ...grpc.CallOption) (*ListenCommitResponse, error) {
	out := new(ListenCommitResponse)
	err := c.cc.Invoke(ctx, "/cosmos.store.streaming.abci.ABCIListenerService/ListenCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABCIListenerServiceServer is the server API for ABCIListenerService service.
//
// Deprecated: Do not use.
type ABCIListenerServiceServer interface {
	// ListenFinalizeBlock is the corresponding endpoint for
	// ABCIListener.ListenEndBlock
	ListenFinalizeBlock(context.Context, *ListenFinalizeBlockRequest) (*ListenFinalizeBlockResponse, error)
	// ListenCommit is the corresponding endpoint for ABCIListener.ListenCommit
	ListenCommit(context.Context, *ListenCommitRequest) (*ListenCommitResponse, error)
}

// Deprecated: Do not use.
// UnimplementedABCIListenerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedABCIListenerServiceServer struct {
}

func (*UnimplementedABCIListenerServiceServer) ListenFinalizeBlock(ctx context.Context, req *ListenFinalizeBlockRequest) (*ListenFinalizeBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenFinalizeBlock not implemented")
}
func (*UnimplementedABCIListenerServiceServer) ListenCommit(ctx context.Context, req *ListenCommitRequest) (*ListenCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenCommit not implemented")
}

// Deprecated: Do not use.
func RegisterABCIListenerServiceServer(s grpc1.Server, srv ABCIListenerServiceServer) {
	s.RegisterService(&_ABCIListenerService_serviceDesc, srv)
}

func _ABCIListenerService_ListenFinalizeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenFinalizeBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenFinalizeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.store.streaming.abci.ABCIListenerService/ListenFinalizeBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenFinalizeBlock(ctx, req.(*ListenFinalizeBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIListenerService_ListenCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.store.streaming.abci.ABCIListenerService/ListenCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenCommit(ctx, req.(*ListenCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABCIListenerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.store.streaming.abci.ABCIListenerService",
	HandlerType: (*ABCIListenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListenFinalizeBlock",
			Handler:    _ABCIListenerService_ListenFinalizeBlock_Handler,
		},
		{
			MethodName: "ListenCommit",
			Handler:    _ABCIListenerService_ListenCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/store/streaming/abci/grpc.proto",
}

func (m *ListenFinalizeBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenFinalizeBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenFinalizeBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Res != nil {
		{
			size, err := m.Res.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Req != nil {
		{
			size, err := m.Req.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListenFinalizeBlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenFinalizeBlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenFinalizeBlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListenCommitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenCommitRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenCommitRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChangeSet) > 0 {
		for iNdEx := len(m.ChangeSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChangeSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGrpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Res != nil {
		{
			size, err := m.Res.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListenCommitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenCommitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenCommitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintGrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListenFinalizeBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Req != nil {
		l = m.Req.Size()
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.Res != nil {
		l = m.Res.Size()
		n += 1 + l + sovGrpc(uint64(l))
	}
	return n
}

func (m *ListenFinalizeBlockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListenCommitRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovGrpc(uint64(m.BlockHeight))
	}
	if m.Res != nil {
		l = m.Res.Size()
		n += 1 + l + sovGrpc(uint64(l))
	}
	if len(m.ChangeSet) > 0 {
		for _, e := range m.ChangeSet {
			l = e.Size()
			n += 1 + l + sovGrpc(uint64(l))
		}
	}
	return n
}

func (m *ListenCommitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpc(x uint64) (n int) {
	return sovGrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListenFinalizeBlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ListenFinalizeBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenFinalizeBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Req == nil {
				m.Req = &v1.FinalizeBlockRequest{}
			}
			if err := m.Req.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Res == nil {
				m.Res = &v1.FinalizeBlockResponse{}
			}
			if err := m.Res.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
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
func (m *ListenFinalizeBlockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ListenFinalizeBlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenFinalizeBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
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
func (m *ListenCommitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ListenCommitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenCommitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Res == nil {
				m.Res = &v1.CommitResponse{}
			}
			if err := m.Res.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChangeSet = append(m.ChangeSet, &types.StoreKVPair{})
			if err := m.ChangeSet[len(m.ChangeSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
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
func (m *ListenCommitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ListenCommitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenCommitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
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
func skipGrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpc
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
					return 0, ErrIntOverflowGrpc
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
					return 0, ErrIntOverflowGrpc
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
				return 0, ErrInvalidLengthGrpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrpc = fmt.Errorf("proto: unexpected end of group")
)
