// Code generated by protoc-gen-go.
// source: gopilosa_pbuf/public.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	gopilosa_pbuf/public.proto

It has these top-level messages:
	Row
	Pair
	ValCount
	Bit
	ColumnAttrSet
	Attr
	AttrMap
	QueryRequest
	QueryResponse
	QueryResult
	ImportRequest
	ImportValueRequest
*/
package internal

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

type Row struct {
	Columns []uint64 `protobuf:"varint,1,rep,packed,name=Columns" json:"Columns,omitempty"`
	Keys    []string `protobuf:"bytes,3,rep,name=Keys" json:"Keys,omitempty"`
	Attrs   []*Attr  `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Row) GetColumns() []uint64 {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Row) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Row) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Pair struct {
	ID    uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=Key" json:"Key,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Pair) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ValCount struct {
	Val   int64 `protobuf:"varint,1,opt,name=Val" json:"Val,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
}

func (m *ValCount) Reset()                    { *m = ValCount{} }
func (m *ValCount) String() string            { return proto.CompactTextString(m) }
func (*ValCount) ProtoMessage()               {}
func (*ValCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ValCount) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *ValCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Bit struct {
	RowID     uint64 `protobuf:"varint,1,opt,name=RowID" json:"RowID,omitempty"`
	ColumnID  uint64 `protobuf:"varint,2,opt,name=ColumnID" json:"ColumnID,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *Bit) Reset()                    { *m = Bit{} }
func (m *Bit) String() string            { return proto.CompactTextString(m) }
func (*Bit) ProtoMessage()               {}
func (*Bit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bit) GetRowID() uint64 {
	if m != nil {
		return m.RowID
	}
	return 0
}

func (m *Bit) GetColumnID() uint64 {
	if m != nil {
		return m.ColumnID
	}
	return 0
}

func (m *Bit) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ColumnAttrSet struct {
	ID    uint64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Key   string  `protobuf:"bytes,3,opt,name=Key" json:"Key,omitempty"`
	Attrs []*Attr `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *ColumnAttrSet) Reset()                    { *m = ColumnAttrSet{} }
func (m *ColumnAttrSet) String() string            { return proto.CompactTextString(m) }
func (*ColumnAttrSet) ProtoMessage()               {}
func (*ColumnAttrSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ColumnAttrSet) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ColumnAttrSet) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ColumnAttrSet) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Attr struct {
	Key         string  `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Type        uint64  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	StringValue string  `protobuf:"bytes,3,opt,name=StringValue" json:"StringValue,omitempty"`
	IntValue    int64   `protobuf:"varint,4,opt,name=IntValue" json:"IntValue,omitempty"`
	BoolValue   bool    `protobuf:"varint,5,opt,name=BoolValue" json:"BoolValue,omitempty"`
	FloatValue  float64 `protobuf:"fixed64,6,opt,name=FloatValue" json:"FloatValue,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Attr) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attr) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Attr) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Attr) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *Attr) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Attr) GetFloatValue() float64 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

type AttrMap struct {
	Attrs []*Attr `protobuf:"bytes,1,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *AttrMap) Reset()                    { *m = AttrMap{} }
func (m *AttrMap) String() string            { return proto.CompactTextString(m) }
func (*AttrMap) ProtoMessage()               {}
func (*AttrMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AttrMap) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type QueryRequest struct {
	Query           string   `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	Shards          []uint64 `protobuf:"varint,2,rep,packed,name=Shards" json:"Shards,omitempty"`
	ColumnAttrs     bool     `protobuf:"varint,3,opt,name=ColumnAttrs" json:"ColumnAttrs,omitempty"`
	Remote          bool     `protobuf:"varint,5,opt,name=Remote" json:"Remote,omitempty"`
	ExcludeRowAttrs bool     `protobuf:"varint,6,opt,name=ExcludeRowAttrs" json:"ExcludeRowAttrs,omitempty"`
	ExcludeColumns  bool     `protobuf:"varint,7,opt,name=ExcludeColumns" json:"ExcludeColumns,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *QueryRequest) GetShards() []uint64 {
	if m != nil {
		return m.Shards
	}
	return nil
}

func (m *QueryRequest) GetColumnAttrs() bool {
	if m != nil {
		return m.ColumnAttrs
	}
	return false
}

func (m *QueryRequest) GetRemote() bool {
	if m != nil {
		return m.Remote
	}
	return false
}

func (m *QueryRequest) GetExcludeRowAttrs() bool {
	if m != nil {
		return m.ExcludeRowAttrs
	}
	return false
}

func (m *QueryRequest) GetExcludeColumns() bool {
	if m != nil {
		return m.ExcludeColumns
	}
	return false
}

type QueryResponse struct {
	Err            string           `protobuf:"bytes,1,opt,name=Err" json:"Err,omitempty"`
	Results        []*QueryResult   `protobuf:"bytes,2,rep,name=Results" json:"Results,omitempty"`
	ColumnAttrSets []*ColumnAttrSet `protobuf:"bytes,3,rep,name=ColumnAttrSets" json:"ColumnAttrSets,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *QueryResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetColumnAttrSets() []*ColumnAttrSet {
	if m != nil {
		return m.ColumnAttrSets
	}
	return nil
}

type QueryResult struct {
	Type     uint32    `protobuf:"varint,6,opt,name=Type" json:"Type,omitempty"`
	Row      *Row      `protobuf:"bytes,1,opt,name=Row" json:"Row,omitempty"`
	N        uint64    `protobuf:"varint,2,opt,name=N" json:"N,omitempty"`
	Pairs    []*Pair   `protobuf:"bytes,3,rep,name=Pairs" json:"Pairs,omitempty"`
	ValCount *ValCount `protobuf:"bytes,5,opt,name=ValCount" json:"ValCount,omitempty"`
	Changed  bool      `protobuf:"varint,4,opt,name=Changed" json:"Changed,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryResult) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *QueryResult) GetRow() *Row {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *QueryResult) GetN() uint64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *QueryResult) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func (m *QueryResult) GetValCount() *ValCount {
	if m != nil {
		return m.ValCount
	}
	return nil
}

func (m *QueryResult) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
}

type ImportRequest struct {
	Index      string   `protobuf:"bytes,1,opt,name=Index" json:"Index,omitempty"`
	Field      string   `protobuf:"bytes,2,opt,name=Field" json:"Field,omitempty"`
	Shard      uint64   `protobuf:"varint,3,opt,name=Shard" json:"Shard,omitempty"`
	RowIDs     []uint64 `protobuf:"varint,4,rep,packed,name=RowIDs" json:"RowIDs,omitempty"`
	ColumnIDs  []uint64 `protobuf:"varint,5,rep,packed,name=ColumnIDs" json:"ColumnIDs,omitempty"`
	RowKeys    []string `protobuf:"bytes,7,rep,name=RowKeys" json:"RowKeys,omitempty"`
	ColumnKeys []string `protobuf:"bytes,8,rep,name=ColumnKeys" json:"ColumnKeys,omitempty"`
	Timestamps []int64  `protobuf:"varint,6,rep,packed,name=Timestamps" json:"Timestamps,omitempty"`
}

func (m *ImportRequest) Reset()                    { *m = ImportRequest{} }
func (m *ImportRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()               {}
func (*ImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ImportRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ImportRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ImportRequest) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *ImportRequest) GetRowIDs() []uint64 {
	if m != nil {
		return m.RowIDs
	}
	return nil
}

func (m *ImportRequest) GetColumnIDs() []uint64 {
	if m != nil {
		return m.ColumnIDs
	}
	return nil
}

func (m *ImportRequest) GetRowKeys() []string {
	if m != nil {
		return m.RowKeys
	}
	return nil
}

func (m *ImportRequest) GetColumnKeys() []string {
	if m != nil {
		return m.ColumnKeys
	}
	return nil
}

func (m *ImportRequest) GetTimestamps() []int64 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type ImportValueRequest struct {
	Index      string   `protobuf:"bytes,1,opt,name=Index" json:"Index,omitempty"`
	Field      string   `protobuf:"bytes,2,opt,name=Field" json:"Field,omitempty"`
	Shard      uint64   `protobuf:"varint,3,opt,name=Shard" json:"Shard,omitempty"`
	ColumnIDs  []uint64 `protobuf:"varint,5,rep,packed,name=ColumnIDs" json:"ColumnIDs,omitempty"`
	ColumnKeys []string `protobuf:"bytes,7,rep,name=ColumnKeys" json:"ColumnKeys,omitempty"`
	Values     []int64  `protobuf:"varint,6,rep,packed,name=Values" json:"Values,omitempty"`
}

func (m *ImportValueRequest) Reset()                    { *m = ImportValueRequest{} }
func (m *ImportValueRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportValueRequest) ProtoMessage()               {}
func (*ImportValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ImportValueRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ImportValueRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ImportValueRequest) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *ImportValueRequest) GetColumnIDs() []uint64 {
	if m != nil {
		return m.ColumnIDs
	}
	return nil
}

func (m *ImportValueRequest) GetColumnKeys() []string {
	if m != nil {
		return m.ColumnKeys
	}
	return nil
}

func (m *ImportValueRequest) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Row)(nil), "internal.Row")
	proto.RegisterType((*Pair)(nil), "internal.Pair")
	proto.RegisterType((*ValCount)(nil), "internal.ValCount")
	proto.RegisterType((*Bit)(nil), "internal.Bit")
	proto.RegisterType((*ColumnAttrSet)(nil), "internal.ColumnAttrSet")
	proto.RegisterType((*Attr)(nil), "internal.Attr")
	proto.RegisterType((*AttrMap)(nil), "internal.AttrMap")
	proto.RegisterType((*QueryRequest)(nil), "internal.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "internal.QueryResponse")
	proto.RegisterType((*QueryResult)(nil), "internal.QueryResult")
	proto.RegisterType((*ImportRequest)(nil), "internal.ImportRequest")
	proto.RegisterType((*ImportValueRequest)(nil), "internal.ImportValueRequest")
}

func init() { proto.RegisterFile("gopilosa_pbuf/public.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0x4e, 0xe2, 0x4c, 0x9a, 0x80, 0x56, 0x50, 0xac, 0x0a, 0x81, 0x65, 0x21, 0xe4,
	0x53, 0x2a, 0x85, 0x3b, 0x88, 0xfe, 0x49, 0x51, 0x45, 0x05, 0xd3, 0x52, 0xc4, 0x09, 0xb9, 0xcd,
	0xd2, 0x5a, 0x72, 0xbc, 0xc6, 0x5e, 0x2b, 0xcd, 0x73, 0xf0, 0x12, 0x1c, 0x78, 0x08, 0x8e, 0xbc,
	0x06, 0x6f, 0x82, 0x76, 0xd6, 0x1b, 0x3b, 0x45, 0xaa, 0x7a, 0xe0, 0xb6, 0xdf, 0x37, 0x33, 0xeb,
	0xf9, 0x76, 0xbf, 0x59, 0xc3, 0xce, 0x95, 0xcc, 0x93, 0x54, 0x96, 0xf1, 0x97, 0xfc, 0xa2, 0xfa,
	0xba, 0x9b, 0x57, 0x17, 0x69, 0x72, 0x39, 0xc9, 0x0b, 0xa9, 0x24, 0xf7, 0x92, 0x4c, 0x89, 0x22,
	0x8b, 0xd3, 0xf0, 0x33, 0x38, 0x28, 0x97, 0xdc, 0x87, 0xfe, 0xbe, 0x4c, 0xab, 0x45, 0x56, 0xfa,
	0x2c, 0x70, 0x22, 0x17, 0x2d, 0xe4, 0x1c, 0xdc, 0x63, 0xb1, 0x2a, 0x7d, 0x27, 0x70, 0xa2, 0x01,
	0xd2, 0x9a, 0xbf, 0x80, 0xee, 0x5b, 0xa5, 0x8a, 0xd2, 0xef, 0x04, 0x4e, 0x34, 0x9c, 0x8e, 0x27,
	0x76, 0xbb, 0x89, 0xa6, 0xd1, 0x04, 0xc3, 0xd7, 0xe0, 0xbe, 0x8f, 0x93, 0x82, 0x8f, 0xa1, 0x33,
	0x3b, 0xf0, 0x59, 0xc0, 0x22, 0x17, 0x3b, 0xb3, 0x03, 0xfe, 0x10, 0x9c, 0x63, 0xb1, 0xf2, 0x9d,
	0x80, 0x45, 0x03, 0xd4, 0x4b, 0xfe, 0x08, 0xba, 0xfb, 0xb2, 0xca, 0x94, 0xdf, 0xa1, 0x24, 0x03,
	0xc2, 0x29, 0x78, 0xe7, 0x71, 0x4a, 0x6b, 0x5d, 0x73, 0x1e, 0xa7, 0xb4, 0x89, 0x83, 0x7a, 0xb9,
	0x59, 0xe3, 0xd8, 0x9a, 0x8f, 0xe0, 0xec, 0x25, 0x4a, 0x07, 0x51, 0x2e, 0xd7, 0x5f, 0x35, 0x80,
	0xef, 0x80, 0x67, 0x54, 0xcd, 0x0e, 0xea, 0x2f, 0xad, 0x31, 0x7f, 0x0a, 0x83, 0xb3, 0x64, 0x21,
	0x4a, 0x15, 0x2f, 0x72, 0x6a, 0xcd, 0xc1, 0x86, 0x08, 0x3f, 0xc1, 0xc8, 0x64, 0x6a, 0x65, 0xa7,
	0x42, 0xdd, 0x43, 0xd3, 0xfd, 0xce, 0xe8, 0x07, 0x03, 0x57, 0xaf, 0xec, 0x06, 0xac, 0xd9, 0x80,
	0x83, 0x7b, 0xb6, 0xca, 0x45, 0xdd, 0x29, 0xad, 0x79, 0x00, 0xc3, 0x53, 0x55, 0x24, 0xd9, 0xd5,
	0x79, 0x9c, 0x56, 0xa2, 0xfe, 0x5c, 0x9b, 0xd2, 0x1a, 0x67, 0x99, 0x32, 0x61, 0x97, 0x64, 0xac,
	0xb1, 0xd6, 0xb8, 0x27, 0x65, 0x6a, 0x82, 0xdd, 0x80, 0x45, 0x1e, 0x36, 0x04, 0x7f, 0x06, 0x70,
	0x94, 0xca, 0xb8, 0xae, 0xed, 0x05, 0x2c, 0x62, 0xd8, 0x62, 0xc2, 0x5d, 0xe8, 0xeb, 0x4e, 0xdf,
	0xc5, 0x79, 0xa3, 0x8d, 0xdd, 0xa5, 0xed, 0x37, 0x83, 0xad, 0x0f, 0x95, 0x28, 0x56, 0x28, 0xbe,
	0x55, 0xa2, 0xa4, 0x5b, 0x21, 0x5c, 0xab, 0x34, 0x80, 0x6f, 0x43, 0xef, 0xf4, 0x3a, 0x2e, 0xe6,
	0xe6, 0xa4, 0x5c, 0xac, 0x91, 0xd6, 0xda, 0x9c, 0x79, 0x49, 0x5a, 0x3d, 0x6c, 0x53, 0xba, 0x12,
	0xc5, 0x42, 0x2a, 0x2b, 0xa6, 0x46, 0x3c, 0x82, 0x07, 0x87, 0x37, 0x97, 0x69, 0x35, 0x17, 0x28,
	0x97, 0xa6, 0xba, 0x47, 0x09, 0xb7, 0x69, 0xfe, 0x12, 0xc6, 0x35, 0x65, 0xdd, 0xdf, 0xa7, 0xc4,
	0x5b, 0x6c, 0xf8, 0x9d, 0xc1, 0xa8, 0x96, 0x52, 0xe6, 0x32, 0x2b, 0x85, 0xbe, 0xaf, 0xc3, 0xa2,
	0xb0, 0xf7, 0x75, 0x58, 0x14, 0x7c, 0x17, 0xfa, 0x28, 0xca, 0x2a, 0x55, 0xf6, 0xca, 0x1f, 0x37,
	0xc7, 0x62, 0x6b, 0xab, 0x54, 0xa1, 0xcd, 0xe2, 0x6f, 0x60, 0xbc, 0x61, 0x2a, 0x33, 0x63, 0xc3,
	0xe9, 0x93, 0xa6, 0x6e, 0x23, 0x8e, 0xb7, 0xd2, 0xc3, 0x5f, 0x0c, 0x86, 0xad, 0x9d, 0xd7, 0x8e,
	0xd1, 0x62, 0x47, 0xb5, 0x63, 0x9e, 0xd3, 0x7c, 0x53, 0x9f, 0xc3, 0xe9, 0xa8, 0xd9, 0x19, 0xe5,
	0x12, 0x69, 0xf2, 0xb7, 0x80, 0x9d, 0xd4, 0x1e, 0x63, 0x27, 0xfa, 0x66, 0xf5, 0xcc, 0xda, 0x56,
	0x5a, 0x37, 0xab, 0x69, 0x34, 0x41, 0x3e, 0x69, 0x26, 0x93, 0x8e, 0x7e, 0x38, 0xe5, 0x4d, 0xa2,
	0x8d, 0x60, 0x33, 0xbd, 0xfa, 0x75, 0xb9, 0x8e, 0xb3, 0x2b, 0x31, 0x27, 0x4f, 0x7a, 0x68, 0x61,
	0xf8, 0x87, 0xc1, 0x68, 0xb6, 0xc8, 0x65, 0xa1, 0x5a, 0x26, 0x99, 0x65, 0x73, 0x71, 0x63, 0x4d,
	0x42, 0x40, 0xb3, 0x47, 0x89, 0x48, 0xe7, 0xd4, 0xe9, 0x00, 0x0d, 0xd0, 0x2c, 0x99, 0x85, 0xcc,
	0xe1, 0xa2, 0x01, 0x64, 0x0b, 0x3d, 0xef, 0xa5, 0xef, 0x1a, 0x43, 0x19, 0xa4, 0xed, 0x6f, 0xc7,
	0xbd, 0xf4, 0xbb, 0x14, 0x6a, 0x08, 0xdd, 0x23, 0xca, 0x25, 0x3d, 0x75, 0x7d, 0x7a, 0xea, 0x2c,
	0xd4, 0x83, 0x61, 0xd2, 0x28, 0xe8, 0x51, 0xb0, 0xc5, 0xe8, 0xf8, 0xfa, 0xa5, 0xd0, 0x4e, 0x73,
	0x22, 0x07, 0x5b, 0x4c, 0xf8, 0x93, 0x01, 0x37, 0x1a, 0x69, 0x90, 0xfe, 0x9f, 0xd0, 0xbb, 0x05,
	0x6d, 0xb6, 0xdd, 0xff, 0xa7, 0xed, 0x6d, 0xe8, 0x51, 0x3f, 0xb6, 0xe5, 0x1a, 0x5d, 0xf4, 0xe8,
	0x17, 0xf1, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x72, 0x01, 0x85, 0x40, 0x06, 0x00,
	0x00,
}
