// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package accesscontrol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type AccessType int32

const (
	AccessType_UNKNOWN AccessType = 0
	AccessType_READ    AccessType = 1
	AccessType_WRITE   AccessType = 2
	AccessType_COMMIT  AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	3: "COMMIT",
}

var AccessType_value = map[string]int32{
	"UNKNOWN": 0,
	"READ":    1,
	"WRITE":   2,
	"COMMIT":  3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{0}
}

type ResourceType int32

const (
	ResourceType_ANY                       ResourceType = 0
	ResourceType_KV                        ResourceType = 1
	ResourceType_Mem                       ResourceType = 2
	ResourceType_DexMem                    ResourceType = 3
	ResourceType_KV_BANK                   ResourceType = 4
	ResourceType_KV_STAKING                ResourceType = 5
	ResourceType_KV_WASM                   ResourceType = 6
	ResourceType_KV_ORACLE                 ResourceType = 7
	ResourceType_KV_DEX                    ResourceType = 8
	ResourceType_KV_EPOCH                  ResourceType = 9
	ResourceType_KV_TOKENFACTORY           ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS    ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS         ResourceType = 13
	ResourceType_KV_STAKING_DELEGATION     ResourceType = 14
	ResourceType_KV_STAKING_VALIDATOR      ResourceType = 15
	ResourceType_KV_AUTH                   ResourceType = 16
)

var ResourceType_name = map[int32]string{
	0:  "ANY",
	1:  "KV",
	2:  "Mem",
	3:  "DexMem",
	4:  "KV_BANK",
	5:  "KV_STAKING",
	6:  "KV_WASM",
	7:  "KV_ORACLE",
	8:  "KV_DEX",
	9:  "KV_EPOCH",
	10: "KV_TOKENFACTORY",
	11: "KV_ORACLE_VOTE_TARGETS",
	12: "KV_ORACLE_AGGREGATE_VOTES",
	13: "KV_ORACLE_FEEDERS",
	14: "KV_STAKING_DELEGATION",
	15: "KV_STAKING_VALIDATOR",
	16: "KV_AUTH",
}

var ResourceType_value = map[string]int32{
	"ANY":                       0,
	"KV":                        1,
	"Mem":                       2,
	"DexMem":                    3,
	"KV_BANK":                   4,
	"KV_STAKING":                5,
	"KV_WASM":                   6,
	"KV_ORACLE":                 7,
	"KV_DEX":                    8,
	"KV_EPOCH":                  9,
	"KV_TOKENFACTORY":           10,
	"KV_ORACLE_VOTE_TARGETS":    11,
	"KV_ORACLE_AGGREGATE_VOTES": 12,
	"KV_ORACLE_FEEDERS":         13,
	"KV_STAKING_DELEGATION":     14,
	"KV_STAKING_VALIDATOR":      15,
	"KV_AUTH":                   16,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x93, 0x76, 0xeb, 0x9f, 0x77, 0xdd, 0xf6, 0x62, 0x18, 0x62, 0x08, 0x72, 0xe2, 0x34,
	0x89, 0x96, 0x89, 0x1b, 0x37, 0xaf, 0xf1, 0xb2, 0xe0, 0x36, 0x46, 0x8e, 0xe7, 0x32, 0x2e, 0x56,
	0x1b, 0x22, 0x40, 0xd0, 0xba, 0xaa, 0x33, 0xc4, 0xbe, 0x05, 0x1f, 0x85, 0x8f, 0xc1, 0x71, 0x47,
	0x8e, 0xa8, 0xfd, 0x22, 0x28, 0x4d, 0x45, 0x61, 0x27, 0x5b, 0xfe, 0x3d, 0x3f, 0xdb, 0xaf, 0x1e,
	0x78, 0x96, 0x59, 0x37, 0xb5, 0xae, 0x37, 0xce, 0xb2, 0xdc, 0xb9, 0xcc, 0xce, 0x8a, 0x85, 0xfd,
	0xd2, 0xcb, 0xec, 0xcc, 0x15, 0xe3, 0x59, 0xe1, 0xba, 0xf3, 0x85, 0x2d, 0x2c, 0x79, 0x52, 0xa5,
	0xba, 0xff, 0xa5, 0xba, 0x5f, 0x4f, 0x27, 0x79, 0x31, 0x3e, 0x3d, 0x79, 0x05, 0x40, 0xd7, 0x40,
	0xdd, 0xcc, 0x73, 0xb2, 0x07, 0xcd, 0xcb, 0x84, 0x27, 0x62, 0x94, 0xa0, 0x47, 0x5a, 0xb0, 0x23,
	0x19, 0x0d, 0xd1, 0x27, 0x6d, 0xd8, 0x1d, 0xc9, 0x58, 0x31, 0xac, 0x11, 0x80, 0x46, 0x5f, 0x0c,
	0x87, 0xb1, 0xc2, 0xfa, 0xc9, 0x8f, 0x1a, 0x74, 0x64, 0xee, 0xec, 0xf5, 0x22, 0xcb, 0xd7, 0x7a,
	0x13, 0xea, 0x34, 0xb9, 0x42, 0x8f, 0x34, 0xa0, 0xc6, 0x35, 0xfa, 0xe5, 0xc1, 0x30, 0x9f, 0x56,
	0x5a, 0x98, 0x7f, 0x2b, 0xf7, 0xf5, 0xf2, 0x11, 0xae, 0xcd, 0x19, 0x4d, 0x38, 0xee, 0x90, 0x03,
	0x00, 0xae, 0x4d, 0xaa, 0x28, 0x8f, 0x93, 0x08, 0x77, 0x37, 0x70, 0x44, 0xd3, 0x21, 0x36, 0xc8,
	0x3e, 0xb4, 0xb9, 0x36, 0x42, 0xd2, 0xfe, 0x80, 0x61, 0xb3, 0xbc, 0x84, 0x6b, 0x13, 0xb2, 0xb7,
	0xd8, 0x22, 0x1d, 0x68, 0x71, 0x6d, 0xd8, 0x1b, 0xd1, 0xbf, 0xc0, 0x36, 0xb9, 0x0f, 0x87, 0x5c,
	0x1b, 0x25, 0x38, 0x4b, 0xce, 0x69, 0x5f, 0x09, 0x79, 0x85, 0x40, 0x1e, 0xc3, 0xc3, 0xbf, 0xb6,
	0xd1, 0x42, 0x31, 0xa3, 0xa8, 0x8c, 0x98, 0x4a, 0x71, 0x8f, 0x3c, 0x85, 0xe3, 0x2d, 0xa3, 0x51,
	0x24, 0x59, 0x44, 0x55, 0x95, 0x4a, 0xb1, 0x43, 0x8e, 0xe0, 0xde, 0x16, 0x9f, 0x33, 0x16, 0x32,
	0x99, 0xe2, 0x3e, 0x39, 0x86, 0xa3, 0xed, 0x67, 0x4d, 0xc8, 0x06, 0xa5, 0x15, 0x8b, 0x04, 0x0f,
	0xc8, 0x23, 0x78, 0xf0, 0x0f, 0xd2, 0x74, 0x10, 0x87, 0x54, 0x09, 0x89, 0x87, 0x9b, 0x89, 0xe8,
	0xa5, 0xba, 0x40, 0x3c, 0x7b, 0xfd, 0x73, 0x19, 0xf8, 0xb7, 0xcb, 0xc0, 0xff, 0xbd, 0x0c, 0xfc,
	0xef, 0xab, 0xc0, 0xbb, 0x5d, 0x05, 0xde, 0xaf, 0x55, 0xe0, 0xbd, 0x7b, 0xf1, 0xe1, 0x53, 0xf1,
	0xf1, 0x7a, 0xd2, 0xcd, 0xec, 0xb4, 0xb7, 0xe9, 0xb5, 0x5a, 0x9e, 0xbb, 0xf7, 0x9f, 0x7b, 0xc5,
	0xcd, 0x3c, 0xbf, 0x53, 0xf4, 0xa4, 0xb1, 0xee, 0xf7, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x99, 0x12, 0x45, 0xdb, 0x07, 0x02, 0x00, 0x00,
}
