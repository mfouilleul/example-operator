// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/promotion_extension_discount_modifier.proto

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

// A promotion extension discount modifier.
type PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier int32

const (
	// Not specified.
	PromotionExtensionDiscountModifierEnum_UNSPECIFIED PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionExtensionDiscountModifierEnum_UNKNOWN PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 1
	// 'Up to'.
	PromotionExtensionDiscountModifierEnum_UP_TO PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 2
)

var PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UP_TO",
}
var PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UP_TO":       2,
}

func (x PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier) String() string {
	return proto.EnumName(PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name, int32(x))
}
func (PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_promotion_extension_discount_modifier_13658de280b60b6c, []int{0, 0}
}

// Container for enum describing possible a promotion extension
// discount modifier.
type PromotionExtensionDiscountModifierEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromotionExtensionDiscountModifierEnum) Reset() {
	*m = PromotionExtensionDiscountModifierEnum{}
}
func (m *PromotionExtensionDiscountModifierEnum) String() string { return proto.CompactTextString(m) }
func (*PromotionExtensionDiscountModifierEnum) ProtoMessage()    {}
func (*PromotionExtensionDiscountModifierEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_promotion_extension_discount_modifier_13658de280b60b6c, []int{0}
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Unmarshal(m, b)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Marshal(b, m, deterministic)
}
func (dst *PromotionExtensionDiscountModifierEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Merge(dst, src)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Size() int {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Size(m)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PromotionExtensionDiscountModifierEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PromotionExtensionDiscountModifierEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PromotionExtensionDiscountModifierEnum)(nil), "google.ads.googleads.v1.enums.PromotionExtensionDiscountModifierEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier", PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name, PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/promotion_extension_discount_modifier.proto", fileDescriptor_promotion_extension_discount_modifier_13658de280b60b6c)
}

var fileDescriptor_promotion_extension_discount_modifier_13658de280b60b6c = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x69, 0x10, 0x20, 0xdc, 0x81, 0x2a, 0x23, 0xa2, 0x43, 0x3b, 0xc0, 0x66, 0x2b, 0x62,
	0x33, 0x53, 0x4a, 0x43, 0x55, 0xa1, 0xa6, 0x91, 0xa0, 0x45, 0x42, 0x91, 0xaa, 0x50, 0x1b, 0xcb,
	0x52, 0xe3, 0x8b, 0xe2, 0xb4, 0xf0, 0x3c, 0x8c, 0x3c, 0x0a, 0x8f, 0xc2, 0x33, 0x30, 0x20, 0xdb,
	0x49, 0x36, 0xe8, 0x62, 0xfd, 0xf2, 0xdd, 0xfd, 0xdf, 0xdd, 0x8f, 0xa6, 0x02, 0x40, 0x6c, 0x38,
	0xc9, 0x98, 0x26, 0x4e, 0x1a, 0xb5, 0x0b, 0x08, 0x57, 0xdb, 0x5c, 0x93, 0xa2, 0x84, 0x1c, 0x2a,
	0x09, 0x6a, 0xc5, 0xdf, 0x2b, 0xae, 0xb4, 0x51, 0x4c, 0xea, 0x35, 0x6c, 0x55, 0xb5, 0xca, 0x81,
	0xc9, 0x57, 0xc9, 0x4b, 0x5c, 0x94, 0x50, 0x81, 0xdf, 0x77, 0xf3, 0x38, 0x63, 0x1a, 0xb7, 0x56,
	0x78, 0x17, 0x60, 0x6b, 0x75, 0x7e, 0xd1, 0x90, 0x0a, 0x49, 0x32, 0xa5, 0xa0, 0xca, 0x8c, 0xaf,
	0x76, 0xc3, 0xc3, 0x37, 0x74, 0x99, 0x34, 0xac, 0xa8, 0x41, 0x8d, 0x6b, 0xd2, 0xac, 0x06, 0x45,
	0x6a, 0x9b, 0x0f, 0x67, 0x68, 0xb8, 0xbf, 0xd3, 0x3f, 0x43, 0xdd, 0x45, 0xfc, 0x90, 0x44, 0xb7,
	0xd3, 0xbb, 0x69, 0x34, 0xee, 0x1d, 0xf8, 0x5d, 0x74, 0xb2, 0x88, 0xef, 0xe3, 0xf9, 0x53, 0xdc,
	0xeb, 0xf8, 0xa7, 0xe8, 0x68, 0x91, 0xac, 0x1e, 0xe7, 0x3d, 0x6f, 0xf4, 0xd3, 0x41, 0x83, 0x35,
	0xe4, 0xf8, 0xdf, 0xe5, 0x47, 0x57, 0xfb, 0x91, 0x89, 0xb9, 0x23, 0xe9, 0x3c, 0x8f, 0x6a, 0x27,
	0x01, 0x9b, 0x4c, 0x09, 0x0c, 0xa5, 0x20, 0x82, 0x2b, 0x7b, 0x65, 0x93, 0x70, 0x21, 0xf5, 0x1f,
	0x81, 0xdf, 0xd8, 0xf7, 0xc3, 0x3b, 0x9c, 0x84, 0xe1, 0xa7, 0xd7, 0x9f, 0x38, 0xab, 0x90, 0x69,
	0xec, 0xa4, 0x51, 0xcb, 0x00, 0x9b, 0x20, 0xf4, 0x57, 0x53, 0x4f, 0x43, 0xa6, 0xd3, 0xb6, 0x9e,
	0x2e, 0x83, 0xd4, 0xd6, 0xbf, 0xbd, 0x81, 0xfb, 0xa4, 0x34, 0x64, 0x9a, 0xd2, 0xb6, 0x83, 0xd2,
	0x65, 0x40, 0xa9, 0xed, 0x79, 0x39, 0xb6, 0x8b, 0x5d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x06, 0xad, 0xad, 0x08, 0x02, 0x00, 0x00,
}
