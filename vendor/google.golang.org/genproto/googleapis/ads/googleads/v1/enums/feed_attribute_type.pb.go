// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/feed_attribute_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible data types for a feed attribute.
type FeedAttributeTypeEnum_FeedAttributeType int32

const (
	// Not specified.
	FeedAttributeTypeEnum_UNSPECIFIED FeedAttributeTypeEnum_FeedAttributeType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedAttributeTypeEnum_UNKNOWN FeedAttributeTypeEnum_FeedAttributeType = 1
	// Int64.
	FeedAttributeTypeEnum_INT64 FeedAttributeTypeEnum_FeedAttributeType = 2
	// Double.
	FeedAttributeTypeEnum_DOUBLE FeedAttributeTypeEnum_FeedAttributeType = 3
	// String.
	FeedAttributeTypeEnum_STRING FeedAttributeTypeEnum_FeedAttributeType = 4
	// Boolean.
	FeedAttributeTypeEnum_BOOLEAN FeedAttributeTypeEnum_FeedAttributeType = 5
	// Url.
	FeedAttributeTypeEnum_URL FeedAttributeTypeEnum_FeedAttributeType = 6
	// Datetime.
	FeedAttributeTypeEnum_DATE_TIME FeedAttributeTypeEnum_FeedAttributeType = 7
	// Int64 list.
	FeedAttributeTypeEnum_INT64_LIST FeedAttributeTypeEnum_FeedAttributeType = 8
	// Double (8 bytes) list.
	FeedAttributeTypeEnum_DOUBLE_LIST FeedAttributeTypeEnum_FeedAttributeType = 9
	// String list.
	FeedAttributeTypeEnum_STRING_LIST FeedAttributeTypeEnum_FeedAttributeType = 10
	// Boolean list.
	FeedAttributeTypeEnum_BOOLEAN_LIST FeedAttributeTypeEnum_FeedAttributeType = 11
	// Url list.
	FeedAttributeTypeEnum_URL_LIST FeedAttributeTypeEnum_FeedAttributeType = 12
	// Datetime list.
	FeedAttributeTypeEnum_DATE_TIME_LIST FeedAttributeTypeEnum_FeedAttributeType = 13
	// Price.
	FeedAttributeTypeEnum_PRICE FeedAttributeTypeEnum_FeedAttributeType = 14
)

var FeedAttributeTypeEnum_FeedAttributeType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INT64",
	3:  "DOUBLE",
	4:  "STRING",
	5:  "BOOLEAN",
	6:  "URL",
	7:  "DATE_TIME",
	8:  "INT64_LIST",
	9:  "DOUBLE_LIST",
	10: "STRING_LIST",
	11: "BOOLEAN_LIST",
	12: "URL_LIST",
	13: "DATE_TIME_LIST",
	14: "PRICE",
}
var FeedAttributeTypeEnum_FeedAttributeType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"INT64":          2,
	"DOUBLE":         3,
	"STRING":         4,
	"BOOLEAN":        5,
	"URL":            6,
	"DATE_TIME":      7,
	"INT64_LIST":     8,
	"DOUBLE_LIST":    9,
	"STRING_LIST":    10,
	"BOOLEAN_LIST":   11,
	"URL_LIST":       12,
	"DATE_TIME_LIST": 13,
	"PRICE":          14,
}

func (x FeedAttributeTypeEnum_FeedAttributeType) String() string {
	return proto.EnumName(FeedAttributeTypeEnum_FeedAttributeType_name, int32(x))
}
func (FeedAttributeTypeEnum_FeedAttributeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_attribute_type_8f52ef3e8b3a5d7e, []int{0, 0}
}

// Container for enum describing possible data types for a feed attribute.
type FeedAttributeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeTypeEnum) Reset()         { *m = FeedAttributeTypeEnum{} }
func (m *FeedAttributeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeTypeEnum) ProtoMessage()    {}
func (*FeedAttributeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_attribute_type_8f52ef3e8b3a5d7e, []int{0}
}
func (m *FeedAttributeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeTypeEnum.Unmarshal(m, b)
}
func (m *FeedAttributeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeTypeEnum.Marshal(b, m, deterministic)
}
func (dst *FeedAttributeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeTypeEnum.Merge(dst, src)
}
func (m *FeedAttributeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeTypeEnum.Size(m)
}
func (m *FeedAttributeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedAttributeTypeEnum)(nil), "google.ads.googleads.v1.enums.FeedAttributeTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.FeedAttributeTypeEnum_FeedAttributeType", FeedAttributeTypeEnum_FeedAttributeType_name, FeedAttributeTypeEnum_FeedAttributeType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/feed_attribute_type.proto", fileDescriptor_feed_attribute_type_8f52ef3e8b3a5d7e)
}

var fileDescriptor_feed_attribute_type_8f52ef3e8b3a5d7e = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6a, 0xdb, 0x40,
	0x14, 0xac, 0xe5, 0xc6, 0x8e, 0x9f, 0x1d, 0x77, 0xbb, 0xd0, 0x1e, 0x4a, 0x73, 0x48, 0x3e, 0x60,
	0x85, 0x68, 0x69, 0x61, 0x7b, 0x5a, 0xc5, 0x1b, 0x23, 0xaa, 0xca, 0x46, 0x96, 0x5c, 0x28, 0x02,
	0xa1, 0x54, 0x5b, 0x21, 0x88, 0xb5, 0xc2, 0x2b, 0x07, 0x72, 0xef, 0x97, 0xf4, 0xd8, 0x4f, 0xe9,
	0xa7, 0xa4, 0x3f, 0x51, 0x56, 0x1b, 0xe9, 0x62, 0x9a, 0x8b, 0x18, 0xcd, 0x9b, 0x99, 0x7d, 0xcc,
	0x83, 0x8f, 0x85, 0x94, 0xc5, 0xad, 0xb0, 0xb3, 0x5c, 0xd9, 0x06, 0x6a, 0x74, 0xe7, 0xd8, 0xa2,
	0x3a, 0xec, 0x94, 0xfd, 0x43, 0x88, 0x3c, 0xcd, 0x9a, 0x66, 0x5f, 0xde, 0x1c, 0x1a, 0x91, 0x36,
	0xf7, 0xb5, 0x20, 0xf5, 0x5e, 0x36, 0x12, 0x9f, 0x1b, 0x35, 0xc9, 0x72, 0x45, 0x7a, 0x23, 0xb9,
	0x73, 0x48, 0x6b, 0x7c, 0xf3, 0xb6, 0xcb, 0xad, 0x4b, 0x3b, 0xab, 0x2a, 0xd9, 0x64, 0x4d, 0x29,
	0x2b, 0x65, 0xcc, 0x97, 0x3f, 0x2d, 0x78, 0x75, 0x2d, 0x44, 0xce, 0xba, 0xe4, 0xe8, 0xbe, 0x16,
	0xbc, 0x3a, 0xec, 0x2e, 0x1f, 0x06, 0xf0, 0xf2, 0x68, 0x82, 0x5f, 0xc0, 0x34, 0x0e, 0x36, 0x6b,
	0x7e, 0xe5, 0x5d, 0x7b, 0x7c, 0x81, 0x9e, 0xe1, 0x29, 0x8c, 0xe3, 0xe0, 0x73, 0xb0, 0xfa, 0x1a,
	0xa0, 0x01, 0x9e, 0xc0, 0x89, 0x17, 0x44, 0x1f, 0xde, 0x23, 0x0b, 0x03, 0x8c, 0x16, 0xab, 0xd8,
	0xf5, 0x39, 0x1a, 0x6a, 0xbc, 0x89, 0x42, 0x2f, 0x58, 0xa2, 0xe7, 0x5a, 0xef, 0xae, 0x56, 0x3e,
	0x67, 0x01, 0x3a, 0xc1, 0x63, 0x18, 0xc6, 0xa1, 0x8f, 0x46, 0xf8, 0x0c, 0x26, 0x0b, 0x16, 0xf1,
	0x34, 0xf2, 0xbe, 0x70, 0x34, 0xc6, 0x73, 0x80, 0x36, 0x27, 0xf5, 0xbd, 0x4d, 0x84, 0x4e, 0xf5,
	0xab, 0x26, 0xcc, 0x10, 0x13, 0x4d, 0x98, 0x44, 0x43, 0x00, 0x46, 0x30, 0x7b, 0x8c, 0x35, 0xcc,
	0x14, 0xcf, 0xe0, 0x34, 0x0e, 0x7d, 0xf3, 0x37, 0xc3, 0x18, 0xe6, 0xfd, 0x03, 0x86, 0x3b, 0xd3,
	0xdb, 0xae, 0x43, 0xef, 0x8a, 0xa3, 0xb9, 0xfb, 0x77, 0x00, 0x17, 0xdf, 0xe5, 0x8e, 0x3c, 0x59,
	0xa5, 0xfb, 0xfa, 0xa8, 0x8f, 0xb5, 0x2e, 0x71, 0x3d, 0xf8, 0xe6, 0x3e, 0x1a, 0x0b, 0x79, 0x9b,
	0x55, 0x05, 0x91, 0xfb, 0xc2, 0x2e, 0x44, 0xd5, 0x56, 0xdc, 0x1d, 0xb3, 0x2e, 0xd5, 0x7f, 0x6e,
	0xfb, 0xa9, 0xfd, 0xfe, 0xb2, 0x86, 0x4b, 0xc6, 0x7e, 0x5b, 0xe7, 0x4b, 0x13, 0xc5, 0x72, 0x45,
	0x0c, 0xd4, 0x68, 0xeb, 0x10, 0x7d, 0x15, 0xf5, 0xa7, 0x9b, 0x27, 0x2c, 0x57, 0x49, 0x3f, 0x4f,
	0xb6, 0x4e, 0xd2, 0xce, 0x1f, 0xac, 0x0b, 0x43, 0x52, 0xca, 0x72, 0x45, 0x69, 0xaf, 0xa0, 0x74,
	0xeb, 0x50, 0xda, 0x6a, 0x6e, 0x46, 0xed, 0x62, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0xfe, 0x62, 0xd5, 0x73, 0x02, 0x00, 0x00,
}
