// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

/*
Package configpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	LogConfigSet
	LogConfig
*/
package configpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LogConfigSet is a set of LogConfig messages.
type LogConfigSet struct {
	Config []*LogConfig `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
}

func (m *LogConfigSet) Reset()                    { *m = LogConfigSet{} }
func (m *LogConfigSet) String() string            { return proto.CompactTextString(m) }
func (*LogConfigSet) ProtoMessage()               {}
func (*LogConfigSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogConfigSet) GetConfig() []*LogConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// LogConfig describes the configuration options for a log instance.
type LogConfig struct {
	LogId        int64                `protobuf:"varint,1,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Prefix       string               `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	RootsPemFile []string             `protobuf:"bytes,3,rep,name=roots_pem_file,json=rootsPemFile" json:"roots_pem_file,omitempty"`
	PrivateKey   *google_protobuf.Any `protobuf:"bytes,4,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// The public key is included for the convenience of test tools (and obviously
	// should match the private key above); it is not used by the CT personality.
	PublicKey     *keyspb.PublicKey `protobuf:"bytes,5,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	RejectExpired bool              `protobuf:"varint,6,opt,name=reject_expired,json=rejectExpired" json:"reject_expired,omitempty"`
	ExtKeyUsages  []string          `protobuf:"bytes,7,rep,name=ext_key_usages,json=extKeyUsages" json:"ext_key_usages,omitempty"`
	// not_after_start defines the start of the range of acceptable NotAfter
	// values, inclusive.
	// Leaving this unset implies no lower bound to the range.
	NotAfterStart *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=not_after_start,json=notAfterStart" json:"not_after_start,omitempty"`
	// not_after_limit defines the end of the range of acceptable NotAfter values,
	// exclusive.
	// Leaving this unset implies no upper bound to the range.
	NotAfterLimit *google_protobuf1.Timestamp `protobuf:"bytes,9,opt,name=not_after_limit,json=notAfterLimit" json:"not_after_limit,omitempty"`
	// accept_only_ca controls whether or not *only* certificates with the CA bit
	// set will be accepted.
	AcceptOnlyCa bool `protobuf:"varint,10,opt,name=accept_only_ca,json=acceptOnlyCa" json:"accept_only_ca,omitempty"`
}

func (m *LogConfig) Reset()                    { *m = LogConfig{} }
func (m *LogConfig) String() string            { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()               {}
func (*LogConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *LogConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *LogConfig) GetRootsPemFile() []string {
	if m != nil {
		return m.RootsPemFile
	}
	return nil
}

func (m *LogConfig) GetPrivateKey() *google_protobuf.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *LogConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *LogConfig) GetRejectExpired() bool {
	if m != nil {
		return m.RejectExpired
	}
	return false
}

func (m *LogConfig) GetExtKeyUsages() []string {
	if m != nil {
		return m.ExtKeyUsages
	}
	return nil
}

func (m *LogConfig) GetNotAfterStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.NotAfterStart
	}
	return nil
}

func (m *LogConfig) GetNotAfterLimit() *google_protobuf1.Timestamp {
	if m != nil {
		return m.NotAfterLimit
	}
	return nil
}

func (m *LogConfig) GetAcceptOnlyCa() bool {
	if m != nil {
		return m.AcceptOnlyCa
	}
	return false
}

func init() {
	proto.RegisterType((*LogConfigSet)(nil), "configpb.LogConfigSet")
	proto.RegisterType((*LogConfig)(nil), "configpb.LogConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0x89, 0xb1, 0x71, 0x33, 0xbb, 0xad, 0x38, 0xfe, 0x61, 0xdc, 0x17, 0x43, 0x51, 0x08,
	0x08, 0x89, 0x54, 0xfa, 0xe4, 0x53, 0x2d, 0x0a, 0xb2, 0x05, 0x4b, 0xaa, 0xcf, 0xc3, 0x24, 0x7b,
	0x13, 0xc7, 0x4e, 0x32, 0xc3, 0xe4, 0xae, 0xec, 0x7c, 0x3f, 0x3f, 0x98, 0x64, 0x26, 0x5b, 0x50,
	0x5f, 0xfa, 0x94, 0xdc, 0x73, 0xcf, 0x39, 0xfc, 0x6e, 0x42, 0x56, 0x8d, 0x1e, 0x5a, 0xd9, 0x15,
	0xc6, 0x6a, 0xd4, 0x74, 0x11, 0x26, 0x53, 0xaf, 0xcf, 0x3b, 0x89, 0x3f, 0x76, 0x75, 0xd1, 0xe8,
	0xbe, 0xec, 0xb4, 0xee, 0x14, 0x94, 0x68, 0xa5, 0x52, 0x52, 0x0c, 0x65, 0x63, 0x9d, 0x41, 0x5d,
	0xde, 0x82, 0x1b, 0x4d, 0x3d, 0x3f, 0x42, 0xc1, 0xfa, 0xe5, 0xec, 0xf5, 0x53, 0xbd, 0x6b, 0x4b,
	0x31, 0xb8, 0x79, 0xf5, 0xea, 0xdf, 0x15, 0xca, 0x1e, 0x46, 0x14, 0xbd, 0x09, 0x86, 0xd3, 0x0f,
	0x64, 0x75, 0xa5, 0xbb, 0x4b, 0x4f, 0x70, 0x03, 0x48, 0xdf, 0x92, 0x24, 0xe0, 0xb0, 0x28, 0x8b,
	0xf3, 0xe5, 0xd9, 0xd3, 0xe2, 0x40, 0x57, 0xdc, 0xf9, 0xaa, 0xd9, 0x72, 0xfa, 0x3b, 0x26, 0xe9,
	0x9d, 0x4a, 0x9f, 0x93, 0x44, 0xe9, 0x8e, 0xcb, 0x2d, 0x8b, 0xb2, 0x28, 0x8f, 0xab, 0x23, 0xa5,
	0xbb, 0x2f, 0x5b, 0xfa, 0x82, 0x24, 0xc6, 0x42, 0x2b, 0xf7, 0xec, 0x41, 0x16, 0xe5, 0x69, 0x35,
	0x4f, 0xf4, 0x35, 0x39, 0xb1, 0x5a, 0xe3, 0xc8, 0x0d, 0xf4, 0xbc, 0x95, 0x0a, 0x58, 0x9c, 0xc5,
	0x79, 0x5a, 0xad, 0xbc, 0x7a, 0x0d, 0xfd, 0x67, 0xa9, 0x80, 0x9e, 0x93, 0xa5, 0xb1, 0xf2, 0x97,
	0x40, 0xe0, 0xb7, 0xe0, 0xd8, 0xc3, 0x2c, 0xca, 0x97, 0x67, 0xcf, 0x8a, 0x70, 0x56, 0x71, 0x38,
	0xab, 0xb8, 0x18, 0x5c, 0x45, 0x66, 0xe3, 0x06, 0x1c, 0x7d, 0x47, 0x88, 0xd9, 0xd5, 0x4a, 0x36,
	0x3e, 0x75, 0xe4, 0x53, 0x4f, 0x8a, 0xf9, 0xab, 0x5d, 0xfb, 0xcd, 0x06, 0x5c, 0x95, 0x9a, 0xc3,
	0x2b, 0x7d, 0x43, 0x4e, 0x2c, 0xfc, 0x84, 0x06, 0x39, 0xec, 0x8d, 0xb4, 0xb0, 0x65, 0x49, 0x16,
	0xe5, 0x8b, 0xea, 0x38, 0xa8, 0x9f, 0x82, 0x38, 0x51, 0xc3, 0x1e, 0xa7, 0x56, 0xbe, 0x1b, 0x45,
	0x07, 0x23, 0x7b, 0x14, 0xa8, 0x61, 0x8f, 0x1b, 0x70, 0xdf, 0xbd, 0x46, 0x3f, 0x92, 0xc7, 0x83,
	0x46, 0x2e, 0x5a, 0x04, 0xcb, 0x47, 0x14, 0x16, 0xd9, 0xc2, 0x33, 0xac, 0xff, 0x23, 0xff, 0x76,
	0xf8, 0x21, 0xd5, 0xf1, 0xa0, 0xf1, 0x62, 0x4a, 0xdc, 0x4c, 0x81, 0xbf, 0x3b, 0x94, 0xec, 0x25,
	0xb2, 0xf4, 0xfe, 0x1d, 0x57, 0x53, 0x60, 0xa2, 0x15, 0x4d, 0x03, 0x06, 0xb9, 0x1e, 0x94, 0xe3,
	0x8d, 0x60, 0xc4, 0x1f, 0xb5, 0x0a, 0xea, 0xd7, 0x41, 0xb9, 0x4b, 0x51, 0x27, 0xbe, 0xe8, 0xfd,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xb1, 0xab, 0xaf, 0x97, 0x02, 0x00, 0x00,
}