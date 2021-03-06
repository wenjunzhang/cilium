// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/regex.proto

package envoy_type_matcher

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A regex matcher designed for safety when used with untrusted input.
type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	// The regex match string. The string must be supported by the configured engine.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ba289439e2572f3, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

// Google's `RE2 <https://github.com/google/re2>`_ regex engine. The regex string must adhere to
// the documented `syntax <https://github.com/google/re2/wiki/Syntax>`_. The engine is designed
// to complete execution in linear time as well as limit the amount of memory used.
type RegexMatcher_GoogleRE2 struct {
	// This field controls the RE2 "program size" which is a rough estimate of how complex a
	// compiled regex is to evaluate. A regex that has a program size greater than the configured
	// value will fail to compile. In this case, the configured max program size can be increased
	// or the regex can be simplified. If not specified, the default is 100.
	MaxProgramSize       *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ba289439e2572f3, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.RegexMatcher.GoogleRE2")
}

func init() { proto.RegisterFile("envoy/type/matcher/regex.proto", fileDescriptor_0ba289439e2572f3) }

var fileDescriptor_0ba289439e2572f3 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x4d, 0x75, 0x6a, 0x33, 0x91, 0x91, 0x8b, 0xa3, 0xe8, 0x28, 0x9e, 0x86, 0x87, 0x44,
	0xba, 0x6f, 0x10, 0xf0, 0xdf, 0x41, 0x28, 0x19, 0xf3, 0x5a, 0x32, 0x7d, 0xad, 0x85, 0xb6, 0x09,
	0x59, 0x36, 0xdb, 0x7d, 0x04, 0xbf, 0xab, 0x77, 0xd9, 0x49, 0x9a, 0x6c, 0x22, 0xe8, 0xad, 0xf4,
	0xc9, 0xf3, 0x7b, 0xde, 0x1f, 0x1e, 0x41, 0xbd, 0x52, 0x2d, 0xb3, 0xad, 0x06, 0x56, 0x49, 0xfb,
	0xfc, 0x06, 0x86, 0x19, 0xc8, 0xa1, 0xa1, 0xda, 0x28, 0xab, 0x08, 0x71, 0x39, 0xed, 0x72, 0xba,
	0xcd, 0xa3, 0x51, 0xae, 0x54, 0x5e, 0x02, 0x73, 0x2f, 0xe6, 0xcb, 0x57, 0xf6, 0x6e, 0xa4, 0xd6,
	0x60, 0x16, 0xbe, 0x13, 0x9d, 0xad, 0x64, 0x59, 0xbc, 0x48, 0x0b, 0x6c, 0xf7, 0xe1, 0x83, 0xcb,
	0x4f, 0x84, 0x4f, 0x44, 0x07, 0x7f, 0xf4, 0x24, 0x32, 0xc3, 0xd8, 0xb3, 0x32, 0x03, 0xc9, 0x10,
	0xc5, 0x68, 0xdc, 0x4f, 0xae, 0xe8, 0xdf, 0x49, 0xfa, 0xbb, 0x45, 0xef, 0x5c, 0x45, 0xdc, 0x24,
	0xfc, 0x78, 0xc3, 0x7b, 0x1f, 0x28, 0x18, 0xa0, 0xfb, 0x3d, 0x11, 0x7a, 0x92, 0x80, 0x84, 0x5c,
	0xe0, 0x9e, 0x73, 0x18, 0x06, 0x31, 0x1a, 0x87, 0xfc, 0x68, 0xc3, 0x0f, 0x4c, 0x10, 0x23, 0xe1,
	0xff, 0x46, 0x53, 0x1c, 0xfe, 0x20, 0xc8, 0x2d, 0x1e, 0x54, 0xb2, 0xc9, 0xb4, 0x51, 0xb9, 0x91,
	0x55, 0xb6, 0x28, 0xd6, 0xb0, 0x3d, 0xe4, 0x9c, 0x7a, 0x22, 0xdd, 0x79, 0xd2, 0xd9, 0x43, 0x6d,
	0x27, 0xc9, 0x93, 0x2c, 0x97, 0x20, 0x4e, 0x2b, 0xd9, 0xa4, 0xbe, 0x34, 0x2d, 0xd6, 0xc0, 0x09,
	0xee, 0x43, 0x9d, 0x17, 0x35, 0x64, 0xdd, 0xe1, 0x64, 0xff, 0x8b, 0x23, 0x7e, 0x8d, 0xe3, 0x42,
	0x79, 0x1d, 0x6d, 0x54, 0xd3, 0xfe, 0x63, 0xc6, 0xb1, 0x53, 0x4b, 0xbb, 0x89, 0x14, 0xcd, 0x0f,
	0xdd, 0xd6, 0xe4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x94, 0xab, 0x1e, 0x0d, 0x97, 0x01, 0x00, 0x00,
}
