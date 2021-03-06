// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_criterion_simulation.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An ad group criterion simulation. Supported combinations of advertising
// channel type, criterion types, simulation type, and simulation modification
// method are detailed below respectively.
//
// SEARCH  KEYWORDS  CPC_BID  UNIFORM
type AdGroupCriterionSimulation struct {
	// The resource name of the ad group criterion simulation.
	// Ad group criterion simulation resource names have the form:
	//
	//
	// `customers/{customer_id}/adGroupCriterionSimulations/{ad_group_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// AdGroup ID of the simulation.
	AdGroupId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
	// Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v1.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*AdGroupCriterionSimulation_CpcBidPointList
	PointList            isAdGroupCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *AdGroupCriterionSimulation) Reset()         { *m = AdGroupCriterionSimulation{} }
func (m *AdGroupCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionSimulation) ProtoMessage()    {}
func (*AdGroupCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_criterion_simulation_a6c4c667ef018570, []int{0}
}
func (m *AdGroupCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionSimulation.Unmarshal(m, b)
}
func (m *AdGroupCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionSimulation.Marshal(b, m, deterministic)
}
func (dst *AdGroupCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionSimulation.Merge(dst, src)
}
func (m *AdGroupCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionSimulation.Size(m)
}
func (m *AdGroupCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionSimulation proto.InternalMessageInfo

func (m *AdGroupCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupCriterionSimulation) GetAdGroupId() *wrappers.Int64Value {
	if m != nil {
		return m.AdGroupId
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *AdGroupCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *AdGroupCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isAdGroupCriterionSimulation_PointList interface {
	isAdGroupCriterionSimulation_PointList()
}

type AdGroupCriterionSimulation_CpcBidPointList struct {
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

func (*AdGroupCriterionSimulation_CpcBidPointList) isAdGroupCriterionSimulation_PointList() {}

func (m *AdGroupCriterionSimulation) GetPointList() isAdGroupCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupCriterionSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdGroupCriterionSimulation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdGroupCriterionSimulation_OneofMarshaler, _AdGroupCriterionSimulation_OneofUnmarshaler, _AdGroupCriterionSimulation_OneofSizer, []interface{}{
		(*AdGroupCriterionSimulation_CpcBidPointList)(nil),
	}
}

func _AdGroupCriterionSimulation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdGroupCriterionSimulation)
	// point_list
	switch x := m.PointList.(type) {
	case *AdGroupCriterionSimulation_CpcBidPointList:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CpcBidPointList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdGroupCriterionSimulation.PointList has unexpected type %T", x)
	}
	return nil
}

func _AdGroupCriterionSimulation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdGroupCriterionSimulation)
	switch tag {
	case 8: // point_list.cpc_bid_point_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CpcBidSimulationPointList)
		err := b.DecodeMessage(msg)
		m.PointList = &AdGroupCriterionSimulation_CpcBidPointList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdGroupCriterionSimulation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdGroupCriterionSimulation)
	// point_list
	switch x := m.PointList.(type) {
	case *AdGroupCriterionSimulation_CpcBidPointList:
		s := proto.Size(x.CpcBidPointList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AdGroupCriterionSimulation)(nil), "google.ads.googleads.v1.resources.AdGroupCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_criterion_simulation.proto", fileDescriptor_ad_group_criterion_simulation_a6c4c667ef018570)
}

var fileDescriptor_ad_group_criterion_simulation_a6c4c667ef018570 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x3e,
	0x14, 0xff, 0xa7, 0xfb, 0xf6, 0xfa, 0x07, 0x29, 0xdc, 0x44, 0x65, 0x82, 0x0e, 0x34, 0xa9, 0x57,
	0x8e, 0xba, 0x21, 0x10, 0xa9, 0x84, 0x48, 0xc7, 0x54, 0x8a, 0x18, 0xaa, 0xb2, 0xa9, 0x17, 0xa8,
	0x52, 0xe4, 0xc6, 0x5e, 0x66, 0xa9, 0xb1, 0x2d, 0xdb, 0x19, 0xda, 0x3b, 0xc0, 0x0d, 0x8f, 0xc0,
	0x25, 0x8f, 0xc2, 0xa3, 0xf0, 0x14, 0x28, 0x4e, 0xe2, 0x4e, 0x2d, 0x65, 0xbd, 0x3b, 0x39, 0xe7,
	0xf7, 0xe1, 0xf3, 0x73, 0x5d, 0x70, 0x96, 0x72, 0x9e, 0xce, 0x88, 0x8f, 0xb0, 0xf2, 0xcb, 0xb2,
	0xa8, 0x6e, 0xba, 0xbe, 0x24, 0x8a, 0xe7, 0x32, 0x21, 0xca, 0x47, 0x38, 0x4e, 0x25, 0xcf, 0x45,
	0x9c, 0x48, 0xaa, 0x89, 0xa4, 0x9c, 0xc5, 0x8a, 0x66, 0xf9, 0x0c, 0x69, 0xca, 0x19, 0x14, 0x92,
	0x6b, 0xee, 0x1e, 0x96, 0x5c, 0x88, 0xb0, 0x82, 0x56, 0x06, 0xde, 0x74, 0xa1, 0x95, 0x69, 0xf9,
	0xab, 0x9c, 0x12, 0x9e, 0x65, 0x9c, 0xf9, 0x8b, 0x9a, 0xad, 0xfe, 0x2a, 0x02, 0x61, 0x79, 0xa6,
	0xee, 0xe0, 0xe3, 0x8c, 0x63, 0x7a, 0x45, 0x93, 0xea, 0x83, 0xe8, 0x6b, 0x8e, 0x2b, 0x8d, 0x93,
	0xb5, 0x35, 0xf4, 0xad, 0x20, 0x15, 0xe9, 0x49, 0x45, 0x32, 0x5f, 0xd3, 0xfc, 0xca, 0xff, 0x22,
	0x91, 0x10, 0x44, 0xaa, 0x6a, 0x7e, 0x50, 0x8b, 0x0a, 0xea, 0x23, 0xc6, 0xb8, 0x36, 0x0a, 0xd5,
	0xf4, 0xd9, 0xf7, 0x2d, 0xd0, 0x0a, 0xf1, 0xa0, 0x48, 0xec, 0xb4, 0x0e, 0xec, 0xc2, 0xfa, 0xb8,
	0xcf, 0xc1, 0xff, 0x75, 0x26, 0x31, 0x43, 0x19, 0xf1, 0x9c, 0xb6, 0xd3, 0xd9, 0x8b, 0x9a, 0x75,
	0xf3, 0x13, 0xca, 0x88, 0xdb, 0x03, 0xfb, 0x36, 0x75, 0x8a, 0xbd, 0x46, 0xdb, 0xe9, 0xec, 0x1f,
	0x3f, 0xae, 0x92, 0x85, 0xf5, 0xb9, 0xe0, 0x90, 0xe9, 0x97, 0x2f, 0xc6, 0x68, 0x96, 0x93, 0x68,
	0x0f, 0x95, 0x96, 0x43, 0xec, 0xbe, 0x01, 0xcd, 0xf9, 0x4d, 0x51, 0xec, 0x6d, 0xdc, 0xcf, 0xde,
	0xb7, 0x84, 0x21, 0x76, 0x2f, 0xc1, 0x66, 0x11, 0x86, 0xb7, 0xd9, 0x76, 0x3a, 0x0f, 0x8e, 0xdf,
	0xc2, 0x55, 0x57, 0x6b, 0x22, 0x84, 0xf3, 0xd5, 0x2e, 0x6f, 0x05, 0x39, 0x63, 0x79, 0xb6, 0xd0,
	0x8a, 0x8c, 0x9a, 0xfb, 0xcd, 0x01, 0x8f, 0xfe, 0x72, 0x4f, 0xde, 0x96, 0x71, 0x99, 0xac, 0xed,
	0x72, 0x7e, 0x47, 0xe3, 0xdc, 0x48, 0x2c, 0x78, 0x2e, 0x03, 0x22, 0x37, 0x5b, 0xea, 0xb9, 0x3d,
	0x00, 0x94, 0x46, 0x52, 0xc7, 0x18, 0x69, 0xe2, 0x6d, 0x9b, 0x8c, 0x0e, 0x96, 0x32, 0xba, 0xd0,
	0x92, 0xb2, 0xb4, 0x8a, 0xd8, 0xe0, 0xdf, 0x21, 0x4d, 0xdc, 0x57, 0x60, 0x97, 0x30, 0x5c, 0x52,
	0x77, 0xd6, 0xa0, 0xee, 0x10, 0x86, 0x0d, 0xf1, 0x1a, 0xb8, 0x89, 0x48, 0xe2, 0x29, 0xc5, 0xb1,
	0xe0, 0x94, 0xe9, 0x78, 0x46, 0x95, 0xf6, 0x76, 0x8d, 0xc4, 0xeb, 0x95, 0x19, 0x94, 0x2f, 0x04,
	0x9e, 0x8a, 0xa4, 0x4f, 0xf1, 0x7c, 0xd3, 0x51, 0xa1, 0xf0, 0x91, 0x2a, 0xfd, 0xfe, 0xbf, 0xe8,
	0x61, 0x62, 0x86, 0xb6, 0xd5, 0x6f, 0x02, 0x30, 0x77, 0xe8, 0x7f, 0x6d, 0x80, 0xa3, 0x84, 0x67,
	0xf0, 0xde, 0x67, 0xda, 0x7f, 0xba, 0xfa, 0xb7, 0x3b, 0x2a, 0x56, 0x1b, 0x39, 0x9f, 0x3f, 0x54,
	0x2a, 0x29, 0x9f, 0x21, 0x96, 0x42, 0x2e, 0x53, 0x3f, 0x25, 0xcc, 0x2c, 0x5e, 0x3f, 0x32, 0x41,
	0xd5, 0x3f, 0xfe, 0x52, 0x7a, 0xb6, 0xfa, 0xd1, 0xd8, 0x18, 0x84, 0xe1, 0xcf, 0xc6, 0xe1, 0xa0,
	0x94, 0x0c, 0xb1, 0x82, 0x65, 0x59, 0x54, 0xe3, 0x2e, 0x8c, 0x6a, 0xe4, 0xaf, 0x1a, 0x33, 0x09,
	0xb1, 0x9a, 0x58, 0xcc, 0x64, 0xdc, 0x9d, 0x58, 0xcc, 0xef, 0xc6, 0x51, 0x39, 0x08, 0x82, 0x10,
	0xab, 0x20, 0xb0, 0xa8, 0x20, 0x18, 0x77, 0x83, 0xc0, 0xe2, 0xa6, 0xdb, 0xe6, 0xb0, 0x27, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0x19, 0xf6, 0x17, 0xfe, 0x04, 0x00, 0x00,
}
