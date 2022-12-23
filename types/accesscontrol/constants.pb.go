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

type AccessOperationSelectorType int32

const (
	AccessOperationSelectorType_NONE                           AccessOperationSelectorType = 0
	AccessOperationSelectorType_JQ                             AccessOperationSelectorType = 1
	AccessOperationSelectorType_JQ_BECH32_ADDRESS              AccessOperationSelectorType = 2
	AccessOperationSelectorType_JQ_LENGTH_PREFIXED_ADDRESS     AccessOperationSelectorType = 3
	AccessOperationSelectorType_SENDER_BECH32_ADDRESS          AccessOperationSelectorType = 4
	AccessOperationSelectorType_SENDER_LENGTH_PREFIXED_ADDRESS AccessOperationSelectorType = 5
	AccessOperationSelectorType_CONTRACT_ADDRESS               AccessOperationSelectorType = 6
	AccessOperationSelectorType_JQ_MESSAGE_CONDITIONAL         AccessOperationSelectorType = 7
	AccessOperationSelectorType_CONSTANT_STRING_TO_HEX         AccessOperationSelectorType = 8
)

var AccessOperationSelectorType_name = map[int32]string{
	0: "NONE",
	1: "JQ",
	2: "JQ_BECH32_ADDRESS",
	3: "JQ_LENGTH_PREFIXED_ADDRESS",
	4: "SENDER_BECH32_ADDRESS",
	5: "SENDER_LENGTH_PREFIXED_ADDRESS",
	6: "CONTRACT_ADDRESS",
	7: "JQ_MESSAGE_CONDITIONAL",
	8: "CONSTANT_STRING_TO_HEX",
}

var AccessOperationSelectorType_value = map[string]int32{
	"NONE":                           0,
	"JQ":                             1,
	"JQ_BECH32_ADDRESS":              2,
	"JQ_LENGTH_PREFIXED_ADDRESS":     3,
	"SENDER_BECH32_ADDRESS":          4,
	"SENDER_LENGTH_PREFIXED_ADDRESS": 5,
	"CONTRACT_ADDRESS":               6,
	"JQ_MESSAGE_CONDITIONAL":         7,
	"CONSTANT_STRING_TO_HEX":         8,
}

func (x AccessOperationSelectorType) String() string {
	return proto.EnumName(AccessOperationSelectorType_name, int32(x))
}

func (AccessOperationSelectorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

type ResourceType int32

const (
	ResourceType_ANY                                      ResourceType = 0
	ResourceType_KV                                       ResourceType = 1
	ResourceType_Mem                                      ResourceType = 2
	ResourceType_DexMem                                   ResourceType = 3
	ResourceType_KV_BANK                                  ResourceType = 4
	ResourceType_KV_STAKING                               ResourceType = 5
	ResourceType_KV_WASM                                  ResourceType = 6
	ResourceType_KV_ORACLE                                ResourceType = 7
	ResourceType_KV_DEX                                   ResourceType = 8
	ResourceType_KV_EPOCH                                 ResourceType = 9
	ResourceType_KV_TOKENFACTORY                          ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS                   ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES                ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS                        ResourceType = 13
	ResourceType_KV_STAKING_DELEGATION                    ResourceType = 14
	ResourceType_KV_STAKING_VALIDATOR                     ResourceType = 15
	ResourceType_KV_AUTH                                  ResourceType = 16
	ResourceType_KV_AUTH_ADDRESS_STORE                    ResourceType = 17
	ResourceType_KV_BANK_SUPPLY                           ResourceType = 18
	ResourceType_KV_BANK_DENOM                            ResourceType = 19
	ResourceType_KV_BANK_BALANCES                         ResourceType = 20
	ResourceType_KV_TOKENFACTORY_DENOM                    ResourceType = 21
	ResourceType_KV_TOKENFACTORY_METADATA                 ResourceType = 22
	ResourceType_KV_TOKENFACTORY_ADMIN                    ResourceType = 23
	ResourceType_KV_TOKENFACTORY_CREATOR                  ResourceType = 24
	ResourceType_KV_ORACLE_EXCHANGE_RATE                  ResourceType = 25
	ResourceType_KV_ORACLE_VOTE_PENALTY_COUNTER           ResourceType = 26
	ResourceType_KV_ORACLE_PRICE_SNAPSHOT                 ResourceType = 27
	ResourceType_KV_STAKING_VALIDATION_POWER              ResourceType = 28
	ResourceType_KV_STAKING_TOTAL_POWER                   ResourceType = 29
	ResourceType_KV_STAKING_VALIDATORS_CON_ADDR           ResourceType = 30
	ResourceType_KV_STAKING_UNBONDING_DELEGATION          ResourceType = 31
	ResourceType_KV_STAKING_UNBONDING_DELEGATION_VAL      ResourceType = 32
	ResourceType_KV_STAKING_REDELEGATION                  ResourceType = 33
	ResourceType_KV_STAKING_REDELEGATION_VAL_SRC          ResourceType = 34
	ResourceType_KV_STAKING_REDELEGATION_VAL_DST          ResourceType = 35
	ResourceType_KV_STAKING_REDELEGATION_QUEUE            ResourceType = 36
	ResourceType_KV_STAKING_VALIDATOR_QUEUE               ResourceType = 37
	ResourceType_KV_STAKING_HISTORICAL_INFO               ResourceType = 38
	ResourceType_KV_STAKING_UNBONDING                     ResourceType = 39
	ResourceType_KV_STAKING_VALIDATORS_BY_POWER           ResourceType = 41
	ResourceType_KV_DISTRIBUTION                          ResourceType = 40
	ResourceType_KV_DISTRIBUTION_FEE_POOL                 ResourceType = 42
	ResourceType_KV_DISTRIBUTION_PROPOSER_KEY             ResourceType = 43
	ResourceType_KV_DISTRIBUTION_OUTSTANDING_REWARDS      ResourceType = 44
	ResourceType_KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR  ResourceType = 45
	ResourceType_KV_DISTRIBUTION_DELEGATOR_STARTING_INFO  ResourceType = 46
	ResourceType_KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS   ResourceType = 47
	ResourceType_KV_DISTRIBUTION_VAL_CURRENT_REWARDS      ResourceType = 48
	ResourceType_KV_DISTRIBUTION_VAL_ACCUM_COMMISSION     ResourceType = 49
	ResourceType_KV_DISTRIBUTION_SLASH_EVENT              ResourceType = 50
	ResourceType_KV_DEX_CONTRACT_LONGBOOK                 ResourceType = 51
	ResourceType_KV_DEX_CONTRACT_SHORTBOOK                ResourceType = 52
	ResourceType_KV_DEX_SETTLEMENT                        ResourceType = 53
	ResourceType_KV_DEX_PAIR_PREFIX                       ResourceType = 54
	ResourceType_KV_DEX_TWAP                              ResourceType = 55
	ResourceType_KV_DEX_PRICE                             ResourceType = 56
	ResourceType_KV_DEX_SETTLEMENT_ENTRY                  ResourceType = 57
	ResourceType_KV_DEX_REGISTERED_PAIR                   ResourceType = 58
	ResourceType_KV_DEX_ORDER                             ResourceType = 60
	ResourceType_KV_DEX_CANCEL                            ResourceType = 61
	ResourceType_KV_DEX_ACCOUNT_ACTIVE_ORDERS             ResourceType = 62
	ResourceType_KV_DEX_ASSET_LIST                        ResourceType = 64
	ResourceType_KV_DEX_NEXT_ORDER_ID                     ResourceType = 65
	ResourceType_KV_DEX_NEXT_SETTLEMENT_ID                ResourceType = 66
	ResourceType_KV_DEX_MATCH_RESULT                      ResourceType = 67
	ResourceType_KV_DEX_SETTLEMENT_ORDER_ID               ResourceType = 68
	ResourceType_KV_DEX_ORDER_BOOK                        ResourceType = 69
	ResourceType_KV_ACCESSCONTROL                         ResourceType = 71
	ResourceType_KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING ResourceType = 72
	ResourceType_KV_WASM_CODE                             ResourceType = 73
	ResourceType_KV_WASM_CONTRACT_ADDRESS                 ResourceType = 74
	ResourceType_KV_WASM_CONTRACT_STORE                   ResourceType = 75
	ResourceType_KV_WASM_SEQUENCE_KEY                     ResourceType = 76
	ResourceType_KV_WASM_CONTRACT_CODE_HISTORY            ResourceType = 77
	ResourceType_KV_WASM_CONTRACT_BY_CODE_ID              ResourceType = 78
	ResourceType_KV_WASM_PINNED_CODE_INDEX                ResourceType = 79
	ResourceType_KV_AUTH_GLOBAL_ACCOUNT_NUMBER            ResourceType = 80
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
	17: "KV_AUTH_ADDRESS_STORE",
	18: "KV_BANK_SUPPLY",
	19: "KV_BANK_DENOM",
	20: "KV_BANK_BALANCES",
	21: "KV_TOKENFACTORY_DENOM",
	22: "KV_TOKENFACTORY_METADATA",
	23: "KV_TOKENFACTORY_ADMIN",
	24: "KV_TOKENFACTORY_CREATOR",
	25: "KV_ORACLE_EXCHANGE_RATE",
	26: "KV_ORACLE_VOTE_PENALTY_COUNTER",
	27: "KV_ORACLE_PRICE_SNAPSHOT",
	28: "KV_STAKING_VALIDATION_POWER",
	29: "KV_STAKING_TOTAL_POWER",
	30: "KV_STAKING_VALIDATORS_CON_ADDR",
	31: "KV_STAKING_UNBONDING_DELEGATION",
	32: "KV_STAKING_UNBONDING_DELEGATION_VAL",
	33: "KV_STAKING_REDELEGATION",
	34: "KV_STAKING_REDELEGATION_VAL_SRC",
	35: "KV_STAKING_REDELEGATION_VAL_DST",
	36: "KV_STAKING_REDELEGATION_QUEUE",
	37: "KV_STAKING_VALIDATOR_QUEUE",
	38: "KV_STAKING_HISTORICAL_INFO",
	39: "KV_STAKING_UNBONDING",
	41: "KV_STAKING_VALIDATORS_BY_POWER",
	40: "KV_DISTRIBUTION",
	42: "KV_DISTRIBUTION_FEE_POOL",
	43: "KV_DISTRIBUTION_PROPOSER_KEY",
	44: "KV_DISTRIBUTION_OUTSTANDING_REWARDS",
	45: "KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR",
	46: "KV_DISTRIBUTION_DELEGATOR_STARTING_INFO",
	47: "KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS",
	48: "KV_DISTRIBUTION_VAL_CURRENT_REWARDS",
	49: "KV_DISTRIBUTION_VAL_ACCUM_COMMISSION",
	50: "KV_DISTRIBUTION_SLASH_EVENT",
	51: "KV_DEX_CONTRACT_LONGBOOK",
	52: "KV_DEX_CONTRACT_SHORTBOOK",
	53: "KV_DEX_SETTLEMENT",
	54: "KV_DEX_PAIR_PREFIX",
	55: "KV_DEX_TWAP",
	56: "KV_DEX_PRICE",
	57: "KV_DEX_SETTLEMENT_ENTRY",
	58: "KV_DEX_REGISTERED_PAIR",
	60: "KV_DEX_ORDER",
	61: "KV_DEX_CANCEL",
	62: "KV_DEX_ACCOUNT_ACTIVE_ORDERS",
	64: "KV_DEX_ASSET_LIST",
	65: "KV_DEX_NEXT_ORDER_ID",
	66: "KV_DEX_NEXT_SETTLEMENT_ID",
	67: "KV_DEX_MATCH_RESULT",
	68: "KV_DEX_SETTLEMENT_ORDER_ID",
	69: "KV_DEX_ORDER_BOOK",
	71: "KV_ACCESSCONTROL",
	72: "KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING",
	73: "KV_WASM_CODE",
	74: "KV_WASM_CONTRACT_ADDRESS",
	75: "KV_WASM_CONTRACT_STORE",
	76: "KV_WASM_SEQUENCE_KEY",
	77: "KV_WASM_CONTRACT_CODE_HISTORY",
	78: "KV_WASM_CONTRACT_BY_CODE_ID",
	79: "KV_WASM_PINNED_CODE_INDEX",
	80: "KV_AUTH_GLOBAL_ACCOUNT_NUMBER",
}

var ResourceType_value = map[string]int32{
	"ANY":                                      0,
	"KV":                                       1,
	"Mem":                                      2,
	"DexMem":                                   3,
	"KV_BANK":                                  4,
	"KV_STAKING":                               5,
	"KV_WASM":                                  6,
	"KV_ORACLE":                                7,
	"KV_DEX":                                   8,
	"KV_EPOCH":                                 9,
	"KV_TOKENFACTORY":                          10,
	"KV_ORACLE_VOTE_TARGETS":                   11,
	"KV_ORACLE_AGGREGATE_VOTES":                12,
	"KV_ORACLE_FEEDERS":                        13,
	"KV_STAKING_DELEGATION":                    14,
	"KV_STAKING_VALIDATOR":                     15,
	"KV_AUTH":                                  16,
	"KV_AUTH_ADDRESS_STORE":                    17,
	"KV_BANK_SUPPLY":                           18,
	"KV_BANK_DENOM":                            19,
	"KV_BANK_BALANCES":                         20,
	"KV_TOKENFACTORY_DENOM":                    21,
	"KV_TOKENFACTORY_METADATA":                 22,
	"KV_TOKENFACTORY_ADMIN":                    23,
	"KV_TOKENFACTORY_CREATOR":                  24,
	"KV_ORACLE_EXCHANGE_RATE":                  25,
	"KV_ORACLE_VOTE_PENALTY_COUNTER":           26,
	"KV_ORACLE_PRICE_SNAPSHOT":                 27,
	"KV_STAKING_VALIDATION_POWER":              28,
	"KV_STAKING_TOTAL_POWER":                   29,
	"KV_STAKING_VALIDATORS_CON_ADDR":           30,
	"KV_STAKING_UNBONDING_DELEGATION":          31,
	"KV_STAKING_UNBONDING_DELEGATION_VAL":      32,
	"KV_STAKING_REDELEGATION":                  33,
	"KV_STAKING_REDELEGATION_VAL_SRC":          34,
	"KV_STAKING_REDELEGATION_VAL_DST":          35,
	"KV_STAKING_REDELEGATION_QUEUE":            36,
	"KV_STAKING_VALIDATOR_QUEUE":               37,
	"KV_STAKING_HISTORICAL_INFO":               38,
	"KV_STAKING_UNBONDING":                     39,
	"KV_STAKING_VALIDATORS_BY_POWER":           41,
	"KV_DISTRIBUTION":                          40,
	"KV_DISTRIBUTION_FEE_POOL":                 42,
	"KV_DISTRIBUTION_PROPOSER_KEY":             43,
	"KV_DISTRIBUTION_OUTSTANDING_REWARDS":      44,
	"KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR":  45,
	"KV_DISTRIBUTION_DELEGATOR_STARTING_INFO":  46,
	"KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS":   47,
	"KV_DISTRIBUTION_VAL_CURRENT_REWARDS":      48,
	"KV_DISTRIBUTION_VAL_ACCUM_COMMISSION":     49,
	"KV_DISTRIBUTION_SLASH_EVENT":              50,
	"KV_DEX_CONTRACT_LONGBOOK":                 51,
	"KV_DEX_CONTRACT_SHORTBOOK":                52,
	"KV_DEX_SETTLEMENT":                        53,
	"KV_DEX_PAIR_PREFIX":                       54,
	"KV_DEX_TWAP":                              55,
	"KV_DEX_PRICE":                             56,
	"KV_DEX_SETTLEMENT_ENTRY":                  57,
	"KV_DEX_REGISTERED_PAIR":                   58,
	"KV_DEX_ORDER":                             60,
	"KV_DEX_CANCEL":                            61,
	"KV_DEX_ACCOUNT_ACTIVE_ORDERS":             62,
	"KV_DEX_ASSET_LIST":                        64,
	"KV_DEX_NEXT_ORDER_ID":                     65,
	"KV_DEX_NEXT_SETTLEMENT_ID":                66,
	"KV_DEX_MATCH_RESULT":                      67,
	"KV_DEX_SETTLEMENT_ORDER_ID":               68,
	"KV_DEX_ORDER_BOOK":                        69,
	"KV_ACCESSCONTROL":                         71,
	"KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING": 72,
	"KV_WASM_CODE":                             73,
	"KV_WASM_CONTRACT_ADDRESS":                 74,
	"KV_WASM_CONTRACT_STORE":                   75,
	"KV_WASM_SEQUENCE_KEY":                     76,
	"KV_WASM_CONTRACT_CODE_HISTORY":            77,
	"KV_WASM_CONTRACT_BY_CODE_ID":              78,
	"KV_WASM_PINNED_CODE_INDEX":                79,
	"KV_AUTH_GLOBAL_ACCOUNT_NUMBER":            80,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{2}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessOperationSelectorType", AccessOperationSelectorType_name, AccessOperationSelectorType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 1265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x4d, 0x77, 0x13, 0x37,
	0x17, 0x8e, 0x89, 0x49, 0x8c, 0x08, 0x70, 0x11, 0xdf, 0x24, 0x18, 0x08, 0xbc, 0xc0, 0x1b, 0x20,
	0xe1, 0xe3, 0xfd, 0x68, 0xa1, 0x2d, 0x95, 0x47, 0x37, 0xf6, 0x64, 0x66, 0xa4, 0x89, 0xa4, 0xb1,
	0xe3, 0x6e, 0x74, 0x82, 0xeb, 0xd3, 0x72, 0x0a, 0x71, 0x4e, 0x6c, 0x7a, 0xca, 0xbf, 0xe8, 0xcf,
	0xea, 0x92, 0x65, 0x97, 0x3d, 0xf0, 0x13, 0xba, 0xe9, 0xb2, 0x47, 0x33, 0xb2, 0xb1, 0x87, 0x50,
	0x56, 0x89, 0xef, 0xf3, 0xdc, 0x2b, 0xdd, 0xe7, 0x7e, 0x68, 0xc8, 0xcd, 0xde, 0x60, 0xf8, 0x6a,
	0x30, 0xdc, 0xd8, 0xed, 0xf5, 0xfa, 0xc3, 0x61, 0x6f, 0xb0, 0x37, 0x3a, 0x18, 0xbc, 0xdc, 0xe8,
	0x0d, 0xf6, 0x86, 0xa3, 0xdd, 0xbd, 0xd1, 0x70, 0x7d, 0xff, 0x60, 0x30, 0x1a, 0xd0, 0x95, 0x82,
	0xb5, 0x3e, 0xc3, 0x5a, 0xff, 0xf9, 0xe1, 0xf3, 0xfe, 0x68, 0xf7, 0xe1, 0xda, 0x13, 0x42, 0x58,
	0x0e, 0x98, 0x37, 0xfb, 0x7d, 0x7a, 0x9c, 0x2c, 0x66, 0x22, 0x12, 0xb2, 0x23, 0x60, 0x8e, 0xd6,
	0x48, 0x55, 0x21, 0xe3, 0x50, 0xa1, 0xc7, 0xc8, 0xd1, 0x8e, 0x0a, 0x0d, 0xc2, 0x11, 0x4a, 0xc8,
	0x42, 0x20, 0x93, 0x24, 0x34, 0x30, 0xbf, 0xf6, 0x67, 0x85, 0x2c, 0x17, 0xce, 0x72, 0xbf, 0x7f,
	0xb0, 0x3b, 0x7a, 0x31, 0xd8, 0xd3, 0xfd, 0x97, 0xfd, 0xde, 0x68, 0x70, 0x90, 0x47, 0xab, 0x91,
	0xaa, 0x90, 0x02, 0x61, 0x8e, 0x2e, 0x90, 0x23, 0x5b, 0xdb, 0x50, 0xa1, 0xe7, 0xc8, 0xe9, 0xad,
	0x6d, 0xdb, 0xc0, 0xa0, 0xf5, 0xf8, 0x91, 0x65, 0x9c, 0x2b, 0xd4, 0x1a, 0x8e, 0xd0, 0x3a, 0xb9,
	0xbc, 0xb5, 0x6d, 0x63, 0x14, 0x4d, 0xd3, 0xb2, 0xa9, 0xc2, 0xcd, 0x70, 0x07, 0xf9, 0x04, 0x9f,
	0xa7, 0x97, 0xc8, 0x39, 0x8d, 0x82, 0xa3, 0x2a, 0xbb, 0x56, 0xe9, 0x2a, 0xa9, 0x7b, 0xe8, 0x53,
	0xee, 0x47, 0xe9, 0x59, 0x02, 0x81, 0x14, 0x46, 0xb1, 0xc0, 0x4c, 0xac, 0x0b, 0xf4, 0x32, 0x39,
	0xbf, 0xb5, 0x6d, 0x13, 0xd4, 0x9a, 0x35, 0xd1, 0x06, 0x52, 0xf0, 0xd0, 0x84, 0x52, 0xb0, 0x18,
	0x16, 0x1d, 0x16, 0x48, 0xa1, 0x0d, 0x13, 0xc6, 0x6a, 0xa3, 0x42, 0xd1, 0xb4, 0x46, 0xda, 0x16,
	0xee, 0x40, 0x6d, 0xed, 0x2f, 0x20, 0x4b, 0xaa, 0x3f, 0x1c, 0xbc, 0x3e, 0xe8, 0xf5, 0xf3, 0x34,
	0x17, 0xc9, 0x3c, 0x13, 0xdd, 0x22, 0xcb, 0xa8, 0x0d, 0x15, 0x67, 0x48, 0xfa, 0xaf, 0x0a, 0xb1,
	0x78, 0xff, 0x17, 0xf7, 0xff, 0xbc, 0x93, 0x36, 0x6a, 0xdb, 0x06, 0x13, 0x11, 0x54, 0xe9, 0x49,
	0x42, 0xa2, 0xb6, 0xd5, 0x86, 0x45, 0xa1, 0x68, 0xc2, 0x51, 0x0f, 0x76, 0x98, 0x4e, 0x60, 0x81,
	0x9e, 0x20, 0xc7, 0xa2, 0xb6, 0x95, 0x8a, 0x05, 0x31, 0xc2, 0xa2, 0x0b, 0x12, 0xb5, 0x2d, 0x77,
	0x67, 0xd3, 0x25, 0x52, 0x8b, 0xda, 0x16, 0x53, 0x19, 0xb4, 0xe0, 0x18, 0x3d, 0x43, 0x4e, 0x45,
	0x6d, 0x6b, 0x64, 0x84, 0x62, 0x93, 0x05, 0x46, 0xaa, 0x2e, 0x10, 0x77, 0xf5, 0x89, 0xb7, 0x6d,
	0x4b, 0x83, 0xd6, 0x30, 0xd5, 0x44, 0xa3, 0xe1, 0x38, 0xbd, 0x42, 0x2e, 0x7d, 0xc0, 0x58, 0xb3,
	0xa9, 0xb0, 0xc9, 0x4c, 0xc1, 0xd2, 0xb0, 0xe4, 0xaa, 0xf3, 0x01, 0xde, 0x44, 0xe4, 0xa8, 0x34,
	0x9c, 0x70, 0xea, 0x7f, 0xb8, 0xac, 0xe5, 0x18, 0x3b, 0xaf, 0x50, 0x0a, 0x38, 0x49, 0x2f, 0x92,
	0xb3, 0x53, 0x50, 0x9b, 0xc5, 0x21, 0x67, 0x46, 0x2a, 0x38, 0xe5, 0x33, 0x62, 0x99, 0x69, 0x01,
	0xf8, 0x08, 0xee, 0xc7, 0x58, 0x7f, 0xab, 0x8d, 0x54, 0x08, 0xa7, 0x29, 0x25, 0x27, 0xbd, 0x2c,
	0x56, 0x67, 0x69, 0x1a, 0x77, 0x81, 0xd2, 0xd3, 0xe4, 0xc4, 0xd8, 0xc6, 0x51, 0xc8, 0x04, 0xce,
	0xb8, 0x12, 0x8e, 0x4d, 0x0d, 0x16, 0x33, 0x11, 0xa0, 0x86, 0xb3, 0x3e, 0xee, 0xb4, 0x00, 0xde,
	0xe1, 0x1c, 0x5d, 0x21, 0x17, 0xcb, 0x50, 0x82, 0x86, 0x71, 0x66, 0x18, 0x9c, 0x3f, 0xcc, 0x91,
	0xf1, 0x24, 0x14, 0x70, 0x81, 0x2e, 0x93, 0x0b, 0x65, 0x28, 0x50, 0x98, 0x67, 0x75, 0xd1, 0x83,
	0x5e, 0x21, 0xdc, 0x09, 0x5a, 0x4c, 0x34, 0xd1, 0x2a, 0x66, 0x10, 0x2e, 0xb9, 0x56, 0x2c, 0x29,
	0x9f, 0xa2, 0x60, 0xb1, 0xe9, 0xda, 0x40, 0x66, 0xc2, 0xa0, 0x82, 0xcb, 0xfe, 0x5a, 0x9e, 0x93,
	0xaa, 0x30, 0x40, 0xab, 0x05, 0x4b, 0x75, 0x4b, 0x1a, 0x58, 0xa6, 0x57, 0xc9, 0xf2, 0xc7, 0x72,
	0x86, 0x52, 0xd8, 0x54, 0x76, 0x50, 0xc1, 0x8a, 0x2f, 0xee, 0x98, 0x60, 0xa4, 0x61, 0xb1, 0xc7,
	0xae, 0xf8, 0xe3, 0x3f, 0xaa, 0x85, 0x76, 0xad, 0x9d, 0xcb, 0x0e, 0x75, 0x7a, 0x83, 0x5c, 0x9d,
	0xe2, 0x64, 0xa2, 0xe1, 0xba, 0x7e, 0xb6, 0xa8, 0x57, 0xe9, 0x6d, 0x72, 0xe3, 0x33, 0x24, 0x17,
	0x1d, 0xae, 0x79, 0x35, 0xc6, 0x44, 0x85, 0x53, 0x51, 0xae, 0x97, 0x8e, 0x9a, 0x06, 0x9d, 0xb7,
	0xd5, 0x2a, 0x80, 0xd5, 0xcf, 0x91, 0xb8, 0x36, 0x70, 0x83, 0x5e, 0x27, 0x57, 0x3e, 0x45, 0xda,
	0xce, 0x30, 0x43, 0xb8, 0xe9, 0x16, 0xc8, 0x61, 0xb9, 0x7b, 0xfc, 0x5f, 0x25, 0xbc, 0x15, 0xba,
	0xee, 0x0b, 0x03, 0x16, 0xdb, 0x50, 0x6c, 0x4a, 0xb8, 0x55, 0xea, 0xe3, 0x49, 0xca, 0x70, 0xfb,
	0xd3, 0xaa, 0x36, 0xba, 0x5e, 0xf9, 0x7f, 0xfb, 0x39, 0xe4, 0xa1, 0xdb, 0x14, 0x8d, 0x2c, 0xcf,
	0xff, 0x8e, 0xaf, 0xf4, 0xb4, 0xd1, 0x8d, 0x94, 0x4d, 0xa5, 0x8c, 0x61, 0x8d, 0x5e, 0x23, 0x2b,
	0x65, 0x34, 0x55, 0x32, 0x95, 0x1a, 0x95, 0x8d, 0xb0, 0x0b, 0x77, 0x7d, 0x15, 0x66, 0x18, 0x32,
	0x33, 0x6e, 0x25, 0xf1, 0x42, 0x86, 0x0e, 0x53, 0x5c, 0xc3, 0x3d, 0x7a, 0x97, 0xdc, 0x2e, 0x13,
	0xbd, 0x42, 0x52, 0xd9, 0x4e, 0x68, 0x5a, 0x5c, 0xb1, 0x4e, 0xd1, 0x00, 0xf7, 0xff, 0x99, 0xac,
	0x0d, 0x53, 0xc6, 0x05, 0xcf, 0x55, 0x59, 0xa7, 0x6b, 0xe4, 0x56, 0x99, 0xec, 0xaa, 0x32, 0x25,
	0xdf, 0xf8, 0x16, 0x1b, 0x87, 0x5d, 0xd7, 0x71, 0x83, 0x4c, 0x29, 0x14, 0x66, 0x42, 0x7c, 0x40,
	0xef, 0x90, 0x9b, 0x87, 0x11, 0x59, 0x10, 0x64, 0x89, 0xcd, 0x9f, 0x16, 0xad, 0x9d, 0x82, 0x0f,
	0xfd, 0x34, 0xcc, 0x30, 0x75, 0xcc, 0x74, 0xcb, 0x62, 0x1b, 0x85, 0x81, 0x47, 0x63, 0x89, 0x71,
	0xc7, 0x4e, 0xd6, 0x7b, 0x2c, 0x45, 0xb3, 0x21, 0x65, 0x04, 0x8f, 0xfd, 0xb2, 0x9b, 0x41, 0x75,
	0x4b, 0x2a, 0x93, 0xc3, 0xff, 0xf1, 0xcb, 0xce, 0xc1, 0x1a, 0x8d, 0x89, 0x31, 0x71, 0x31, 0xff,
	0x4b, 0xcf, 0x13, 0xea, 0xcd, 0x29, 0x0b, 0x95, 0x7f, 0x4d, 0xe0, 0x7f, 0xf4, 0x14, 0x39, 0xee,
	0xed, 0xa6, 0xc3, 0x52, 0xf8, 0x3f, 0x05, 0xb2, 0x34, 0x26, 0xba, 0x31, 0x86, 0x2f, 0xfc, 0x38,
	0xcc, 0x46, 0xb4, 0x28, 0x8c, 0xea, 0xc2, 0x97, 0x7e, 0x72, 0x1d, 0xa8, 0xb0, 0x19, 0x6a, 0x83,
	0x0a, 0x79, 0x7e, 0x04, 0x3c, 0x99, 0x0a, 0x25, 0x15, 0x47, 0x05, 0x5f, 0xf9, 0x0d, 0x98, 0xdf,
	0xdd, 0xed, 0xba, 0x18, 0xbe, 0x1e, 0x77, 0x0c, 0xee, 0x38, 0xa9, 0xdc, 0x3e, 0xb1, 0x2c, 0x30,
	0x61, 0x1b, 0x0b, 0x1f, 0x0d, 0xdf, 0x4c, 0x65, 0xc4, 0xb4, 0x46, 0x63, 0xe3, 0x50, 0x1b, 0xf8,
	0xd6, 0xf7, 0xb6, 0x33, 0x0b, 0xdc, 0x31, 0x05, 0xdd, 0x86, 0x1c, 0xd8, 0x94, 0x42, 0x39, 0x32,
	0x75, 0xeb, 0x90, 0x43, 0x83, 0x5e, 0x20, 0x67, 0x3c, 0x9c, 0x30, 0x13, 0xb4, 0xac, 0x42, 0x9d,
	0xc5, 0x06, 0x02, 0x3f, 0x4d, 0xa5, 0x44, 0x27, 0x71, 0xf9, 0xd4, 0x45, 0x0a, 0x63, 0xae, 0x38,
	0xfa, 0x1d, 0xce, 0x82, 0x00, 0xb5, 0xce, 0x4b, 0x22, 0x63, 0x68, 0xd2, 0x7b, 0xe4, 0x4e, 0xd9,
	0x9a, 0x3f, 0x84, 0x96, 0x63, 0xea, 0x1e, 0x76, 0x11, 0x74, 0x6d, 0xc2, 0xd2, 0xd4, 0x8d, 0x63,
	0xcb, 0x4b, 0x95, 0xe3, 0x81, 0xe4, 0x08, 0xa1, 0x6f, 0x02, 0x6f, 0x29, 0x3d, 0xf2, 0x5b, 0x5e,
	0xf6, 0x59, 0xb4, 0x78, 0x7a, 0x22, 0x2f, 0x4c, 0x8e, 0x69, 0xdc, 0xce, 0x50, 0x04, 0x98, 0xcf,
	0x5e, 0xec, 0x37, 0xce, 0xac, 0x97, 0x3b, 0xce, 0xb7, 0x7e, 0x17, 0x12, 0xdf, 0x9c, 0xb3, 0x94,
	0x46, 0xb7, 0x60, 0x85, 0x1c, 0x84, 0x17, 0x37, 0x27, 0xa4, 0xa1, 0x10, 0xc8, 0x3d, 0x26, 0xdc,
	0x4b, 0x2e, 0xfd, 0x11, 0xf9, 0x93, 0xd8, 0x8c, 0x65, 0xa3, 0x98, 0x80, 0xbc, 0xac, 0x22, 0x4b,
	0x1a, 0xa8, 0x20, 0x5d, 0xad, 0xd6, 0x9e, 0xc2, 0xd3, 0xd5, 0x6a, 0xed, 0x19, 0x3c, 0x5b, 0xad,
	0xd6, 0x36, 0x61, 0xb3, 0xb1, 0xf5, 0xdb, 0xbb, 0x7a, 0xe5, 0xed, 0xbb, 0x7a, 0xe5, 0x8f, 0x77,
	0xf5, 0xca, 0xaf, 0xef, 0xeb, 0x73, 0x6f, 0xdf, 0xd7, 0xe7, 0x7e, 0x7f, 0x5f, 0x9f, 0xfb, 0xee,
	0xc1, 0x0f, 0x2f, 0x46, 0x3f, 0xbe, 0x7e, 0xbe, 0xde, 0x1b, 0xbc, 0xda, 0xf0, 0x5f, 0x85, 0xc5,
	0x9f, 0xfb, 0xc3, 0xef, 0x7f, 0xda, 0x18, 0xbd, 0xd9, 0xef, 0x97, 0x3e, 0x13, 0x9f, 0x2f, 0xe4,
	0x5f, 0x87, 0x8f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x6d, 0x96, 0x43, 0x45, 0x0a, 0x00,
	0x00,
}
