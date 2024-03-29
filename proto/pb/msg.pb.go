// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	CPriceItem
	User
	MsgCfgBase
	MsgCfgTest
	CfgCfg
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CPriceItem struct {
	Mtype      int32  `protobuf:"zigzag32,1,opt,name=mtype" json:"mtype,omitempty"`
	Oriname    string `protobuf:"bytes,2,opt,name=oriname" json:"oriname,omitempty"`
	Count      int32  `protobuf:"zigzag32,3,opt,name=count" json:"count,omitempty"`
	InnerParam int32  `protobuf:"zigzag32,4,opt,name=inner_param,json=innerParam" json:"inner_param,omitempty"`
	Exinfo     int32  `protobuf:"zigzag32,5,opt,name=exinfo" json:"exinfo,omitempty"`
}

func (m *CPriceItem) Reset()                    { *m = CPriceItem{} }
func (m *CPriceItem) String() string            { return proto.CompactTextString(m) }
func (*CPriceItem) ProtoMessage()               {}
func (*CPriceItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CPriceItem) GetMtype() int32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *CPriceItem) GetOriname() string {
	if m != nil {
		return m.Oriname
	}
	return ""
}

func (m *CPriceItem) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CPriceItem) GetInnerParam() int32 {
	if m != nil {
		return m.InnerParam
	}
	return 0
}

func (m *CPriceItem) GetExinfo() int32 {
	if m != nil {
		return m.Exinfo
	}
	return 0
}

type User struct {
	Pwd    string        `protobuf:"bytes,1,opt,name=pwd" json:"pwd,omitempty"`
	Item   *CPriceItem   `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	Awards []*CPriceItem `protobuf:"bytes,3,rep,name=awards" json:"awards,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetItem() *CPriceItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *User) GetAwards() []*CPriceItem {
	if m != nil {
		return m.Awards
	}
	return nil
}

type MsgCfgBase struct {
	Name string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Rate float32 `protobuf:"fixed32,2,opt,name=rate" json:"rate,omitempty"`
}

func (m *MsgCfgBase) Reset()                    { *m = MsgCfgBase{} }
func (m *MsgCfgBase) String() string            { return proto.CompactTextString(m) }
func (*MsgCfgBase) ProtoMessage()               {}
func (*MsgCfgBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MsgCfgBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCfgBase) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type MsgCfgTest struct {
	Id     int32         `protobuf:"zigzag32,1,opt,name=id" json:"id,omitempty"`
	Base   *MsgCfgBase   `protobuf:"bytes,2,opt,name=base" json:"base,omitempty"`
	Open   bool          `protobuf:"varint,3,opt,name=open" json:"open,omitempty"`
	Awards []*CPriceItem `protobuf:"bytes,4,rep,name=awards" json:"awards,omitempty"`
	Params []int32       `protobuf:"zigzag32,5,rep,packed,name=params" json:"params,omitempty"`
}

func (m *MsgCfgTest) Reset()                    { *m = MsgCfgTest{} }
func (m *MsgCfgTest) String() string            { return proto.CompactTextString(m) }
func (*MsgCfgTest) ProtoMessage()               {}
func (*MsgCfgTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MsgCfgTest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgCfgTest) GetBase() *MsgCfgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *MsgCfgTest) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *MsgCfgTest) GetAwards() []*CPriceItem {
	if m != nil {
		return m.Awards
	}
	return nil
}

func (m *MsgCfgTest) GetParams() []int32 {
	if m != nil {
		return m.Params
	}
	return nil
}

type CfgCfg struct {
	A float32      `protobuf:"fixed32,1,opt,name=A" json:"A,omitempty"`
	B []float32    `protobuf:"fixed32,2,rep,packed,name=B" json:"B,omitempty"`
	C *CfgCfg_Cfgc `protobuf:"bytes,3,opt,name=C" json:"C,omitempty"`
}

func (m *CfgCfg) Reset()                    { *m = CfgCfg{} }
func (m *CfgCfg) String() string            { return proto.CompactTextString(m) }
func (*CfgCfg) ProtoMessage()               {}
func (*CfgCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CfgCfg) GetA() float32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CfgCfg) GetB() []float32 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *CfgCfg) GetC() *CfgCfg_Cfgc {
	if m != nil {
		return m.C
	}
	return nil
}

type CfgCfg_Cfgc struct {
	D float32 `protobuf:"fixed32,1,opt,name=D" json:"D,omitempty"`
	E float32 `protobuf:"fixed32,2,opt,name=E" json:"E,omitempty"`
}

func (m *CfgCfg_Cfgc) Reset()                    { *m = CfgCfg_Cfgc{} }
func (m *CfgCfg_Cfgc) String() string            { return proto.CompactTextString(m) }
func (*CfgCfg_Cfgc) ProtoMessage()               {}
func (*CfgCfg_Cfgc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *CfgCfg_Cfgc) GetD() float32 {
	if m != nil {
		return m.D
	}
	return 0
}

func (m *CfgCfg_Cfgc) GetE() float32 {
	if m != nil {
		return m.E
	}
	return 0
}

func init() {
	proto.RegisterType((*CPriceItem)(nil), "pb.CPriceItem")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*MsgCfgBase)(nil), "pb.MsgCfgBase")
	proto.RegisterType((*MsgCfgTest)(nil), "pb.MsgCfgTest")
	proto.RegisterType((*CfgCfg)(nil), "pb.CfgCfg")
	proto.RegisterType((*CfgCfg_Cfgc)(nil), "pb.CfgCfg.Cfgc")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x65, 0x26, 0x69, 0xde, 0xeb, 0xed, 0xa3, 0x4f, 0x07, 0x91, 0x41, 0x10, 0x43, 0x16, 0xd2,
	0x55, 0x17, 0xd5, 0x3f, 0xd0, 0xa4, 0x05, 0x5d, 0x08, 0x65, 0xd0, 0xb5, 0x4c, 0x9a, 0x69, 0xc8,
	0x22, 0x1f, 0xcc, 0x8c, 0x54, 0x7f, 0x83, 0x5b, 0x7f, 0xb0, 0xdc, 0x9b, 0x94, 0x76, 0x21, 0xae,
	0x72, 0xce, 0x9c, 0x7b, 0xe7, 0x9c, 0x39, 0x04, 0xc6, 0xb5, 0x2b, 0xe7, 0x9d, 0x6d, 0x7d, 0x2b,
	0x78, 0x97, 0x27, 0x9f, 0x0c, 0x20, 0xdb, 0xd8, 0x6a, 0x6b, 0x1e, 0xbd, 0xa9, 0xc5, 0x05, 0x8c,
	0x6a, 0xff, 0xd1, 0x19, 0xc9, 0x62, 0x36, 0x3b, 0x57, 0x3d, 0x11, 0x12, 0xfe, 0xb4, 0xb6, 0x6a,
	0x74, 0x6d, 0x24, 0x8f, 0xd9, 0x6c, 0xac, 0x0e, 0x14, 0xe7, 0xb7, 0xed, 0x5b, 0xe3, 0x65, 0xd0,
	0xcf, 0x13, 0x11, 0x37, 0x30, 0xa9, 0x9a, 0xc6, 0xd8, 0xd7, 0x4e, 0x5b, 0x5d, 0xcb, 0x90, 0x34,
	0xa0, 0xa3, 0x0d, 0x9e, 0x88, 0x4b, 0x88, 0xcc, 0x7b, 0xd5, 0xec, 0x5a, 0x39, 0x22, 0x6d, 0x60,
	0x49, 0x01, 0xe1, 0x8b, 0x33, 0x56, 0x9c, 0x41, 0xd0, 0xed, 0x0b, 0x0a, 0x31, 0x56, 0x08, 0x45,
	0x02, 0x61, 0xe5, 0x4d, 0x4d, 0xfe, 0x93, 0xc5, 0x74, 0xde, 0xe5, 0xf3, 0x63, 0x6c, 0x45, 0x9a,
	0xb8, 0x85, 0x48, 0xef, 0xb5, 0x2d, 0x9c, 0x0c, 0xe2, 0xe0, 0x87, 0xa9, 0x41, 0x4d, 0xee, 0x01,
	0x9e, 0x5c, 0x99, 0xed, 0xca, 0x54, 0x3b, 0x23, 0x04, 0x84, 0xf4, 0xb2, 0xde, 0x8c, 0x30, 0x9e,
	0x59, 0xed, 0xfb, 0xd7, 0x72, 0x45, 0x38, 0xf9, 0x62, 0x87, 0xb5, 0x67, 0xe3, 0xbc, 0x98, 0x02,
	0xaf, 0x8a, 0xa1, 0x26, 0x5e, 0x51, 0xc0, 0x5c, 0x3b, 0x73, 0x1a, 0xf0, 0x68, 0xa2, 0x48, 0xc3,
	0x6b, 0xdb, 0xce, 0x34, 0x54, 0xd6, 0x5f, 0x45, 0xf8, 0x24, 0x74, 0xf8, 0x5b, 0x68, 0xac, 0x8c,
	0xda, 0x74, 0x72, 0x14, 0x07, 0x58, 0x59, 0xcf, 0x92, 0x12, 0xa2, 0x6c, 0x87, 0x3e, 0xe2, 0x1f,
	0xb0, 0x25, 0x05, 0xe2, 0x8a, 0x2d, 0x91, 0xa5, 0x92, 0xc7, 0x01, 0xb2, 0x54, 0x5c, 0x03, 0xcb,
	0xc8, 0x76, 0xb2, 0xf8, 0x4f, 0x06, 0xb4, 0x82, 0x9f, 0xad, 0x62, 0xd9, 0x55, 0x02, 0x21, 0x42,
	0x5c, 0x5a, 0x1d, 0xae, 0x58, 0x21, 0x5b, 0x0f, 0x15, 0xb0, 0x75, 0xca, 0x1f, 0x82, 0x3c, 0xa2,
	0x1f, 0xe7, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x73, 0xbd, 0xe0, 0xb7, 0x45, 0x02, 0x00, 0x00,
}
