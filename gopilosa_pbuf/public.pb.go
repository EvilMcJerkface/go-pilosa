// Code generated by protoc-gen-go.
// source: gopilosa_pbuf/public.proto
// DO NOT EDIT!

/*
Package gopilosa_pbuf is a generated protocol buffer package.

It is generated from these files:
	gopilosa_pbuf/public.proto

It has these top-level messages:
	Row
	RowIdentifiers
	Pair
	FieldRow
	GroupCount
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
	TranslateKeysRequest
	TranslateKeysResponse
	ImportRoaringRequestView
	ImportRoaringRequest
*/
package gopilosa_pbuf

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

type RowIdentifiers struct {
	Rows []uint64 `protobuf:"varint,1,rep,packed,name=Rows" json:"Rows,omitempty"`
	Keys []string `protobuf:"bytes,2,rep,name=Keys" json:"Keys,omitempty"`
}

func (m *RowIdentifiers) Reset()                    { *m = RowIdentifiers{} }
func (m *RowIdentifiers) String() string            { return proto.CompactTextString(m) }
func (*RowIdentifiers) ProtoMessage()               {}
func (*RowIdentifiers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RowIdentifiers) GetRows() []uint64 {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *RowIdentifiers) GetKeys() []string {
	if m != nil {
		return m.Keys
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
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

type FieldRow struct {
	Field string `protobuf:"bytes,1,opt,name=Field" json:"Field,omitempty"`
	RowID uint64 `protobuf:"varint,2,opt,name=RowID" json:"RowID,omitempty"`
}

func (m *FieldRow) Reset()                    { *m = FieldRow{} }
func (m *FieldRow) String() string            { return proto.CompactTextString(m) }
func (*FieldRow) ProtoMessage()               {}
func (*FieldRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FieldRow) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *FieldRow) GetRowID() uint64 {
	if m != nil {
		return m.RowID
	}
	return 0
}

type GroupCount struct {
	Group []*FieldRow `protobuf:"bytes,1,rep,name=Group" json:"Group,omitempty"`
	Count uint64      `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
}

func (m *GroupCount) Reset()                    { *m = GroupCount{} }
func (m *GroupCount) String() string            { return proto.CompactTextString(m) }
func (*GroupCount) ProtoMessage()               {}
func (*GroupCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GroupCount) GetGroup() []*FieldRow {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupCount) GetCount() uint64 {
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
func (*ValCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*Bit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*ColumnAttrSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*AttrMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
	Type           uint32          `protobuf:"varint,6,opt,name=Type" json:"Type,omitempty"`
	Row            *Row            `protobuf:"bytes,1,opt,name=Row" json:"Row,omitempty"`
	N              uint64          `protobuf:"varint,2,opt,name=N" json:"N,omitempty"`
	Pairs          []*Pair         `protobuf:"bytes,3,rep,name=Pairs" json:"Pairs,omitempty"`
	Changed        bool            `protobuf:"varint,4,opt,name=Changed" json:"Changed,omitempty"`
	ValCount       *ValCount       `protobuf:"bytes,5,opt,name=ValCount" json:"ValCount,omitempty"`
	RowIDs         []uint64        `protobuf:"varint,7,rep,packed,name=RowIDs" json:"RowIDs,omitempty"`
	GroupCounts    []*GroupCount   `protobuf:"bytes,8,rep,name=GroupCounts" json:"GroupCounts,omitempty"`
	RowIdentifiers *RowIdentifiers `protobuf:"bytes,9,opt,name=RowIdentifiers" json:"RowIdentifiers,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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

func (m *QueryResult) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
}

func (m *QueryResult) GetValCount() *ValCount {
	if m != nil {
		return m.ValCount
	}
	return nil
}

func (m *QueryResult) GetRowIDs() []uint64 {
	if m != nil {
		return m.RowIDs
	}
	return nil
}

func (m *QueryResult) GetGroupCounts() []*GroupCount {
	if m != nil {
		return m.GroupCounts
	}
	return nil
}

func (m *QueryResult) GetRowIdentifiers() *RowIdentifiers {
	if m != nil {
		return m.RowIdentifiers
	}
	return nil
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
func (*ImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

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
func (*ImportValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

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

type TranslateKeysRequest struct {
	Index string   `protobuf:"bytes,1,opt,name=Index" json:"Index,omitempty"`
	Field string   `protobuf:"bytes,2,opt,name=Field" json:"Field,omitempty"`
	Keys  []string `protobuf:"bytes,3,rep,name=Keys" json:"Keys,omitempty"`
}

func (m *TranslateKeysRequest) Reset()                    { *m = TranslateKeysRequest{} }
func (m *TranslateKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*TranslateKeysRequest) ProtoMessage()               {}
func (*TranslateKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *TranslateKeysRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TranslateKeysRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *TranslateKeysRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type TranslateKeysResponse struct {
	IDs []uint64 `protobuf:"varint,3,rep,packed,name=IDs" json:"IDs,omitempty"`
}

func (m *TranslateKeysResponse) Reset()                    { *m = TranslateKeysResponse{} }
func (m *TranslateKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*TranslateKeysResponse) ProtoMessage()               {}
func (*TranslateKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *TranslateKeysResponse) GetIDs() []uint64 {
	if m != nil {
		return m.IDs
	}
	return nil
}

type ImportRoaringRequestView struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ImportRoaringRequestView) Reset()                    { *m = ImportRoaringRequestView{} }
func (m *ImportRoaringRequestView) String() string            { return proto.CompactTextString(m) }
func (*ImportRoaringRequestView) ProtoMessage()               {}
func (*ImportRoaringRequestView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ImportRoaringRequestView) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImportRoaringRequestView) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImportRoaringRequest struct {
	Clear bool                        `protobuf:"varint,1,opt,name=Clear" json:"Clear,omitempty"`
	Views []*ImportRoaringRequestView `protobuf:"bytes,2,rep,name=views" json:"views,omitempty"`
}

func (m *ImportRoaringRequest) Reset()                    { *m = ImportRoaringRequest{} }
func (m *ImportRoaringRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportRoaringRequest) ProtoMessage()               {}
func (*ImportRoaringRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ImportRoaringRequest) GetClear() bool {
	if m != nil {
		return m.Clear
	}
	return false
}

func (m *ImportRoaringRequest) GetViews() []*ImportRoaringRequestView {
	if m != nil {
		return m.Views
	}
	return nil
}

func init() {
	proto.RegisterType((*Row)(nil), "gopilosa_pbuf.Row")
	proto.RegisterType((*RowIdentifiers)(nil), "gopilosa_pbuf.RowIdentifiers")
	proto.RegisterType((*Pair)(nil), "gopilosa_pbuf.Pair")
	proto.RegisterType((*FieldRow)(nil), "gopilosa_pbuf.FieldRow")
	proto.RegisterType((*GroupCount)(nil), "gopilosa_pbuf.GroupCount")
	proto.RegisterType((*ValCount)(nil), "gopilosa_pbuf.ValCount")
	proto.RegisterType((*Bit)(nil), "gopilosa_pbuf.Bit")
	proto.RegisterType((*ColumnAttrSet)(nil), "gopilosa_pbuf.ColumnAttrSet")
	proto.RegisterType((*Attr)(nil), "gopilosa_pbuf.Attr")
	proto.RegisterType((*AttrMap)(nil), "gopilosa_pbuf.AttrMap")
	proto.RegisterType((*QueryRequest)(nil), "gopilosa_pbuf.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "gopilosa_pbuf.QueryResponse")
	proto.RegisterType((*QueryResult)(nil), "gopilosa_pbuf.QueryResult")
	proto.RegisterType((*ImportRequest)(nil), "gopilosa_pbuf.ImportRequest")
	proto.RegisterType((*ImportValueRequest)(nil), "gopilosa_pbuf.ImportValueRequest")
	proto.RegisterType((*TranslateKeysRequest)(nil), "gopilosa_pbuf.TranslateKeysRequest")
	proto.RegisterType((*TranslateKeysResponse)(nil), "gopilosa_pbuf.TranslateKeysResponse")
	proto.RegisterType((*ImportRoaringRequestView)(nil), "gopilosa_pbuf.ImportRoaringRequestView")
	proto.RegisterType((*ImportRoaringRequest)(nil), "gopilosa_pbuf.ImportRoaringRequest")
}

func init() { proto.RegisterFile("gopilosa_pbuf/public.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xd6, 0xc4, 0x4e, 0xe2, 0x9c, 0x6c, 0x96, 0x6a, 0x58, 0x8a, 0x59, 0x2d, 0x28, 0xb2, 0x10,
	0xb8, 0x17, 0x2c, 0x52, 0x5a, 0x21, 0x24, 0x04, 0x12, 0xdb, 0x6c, 0x51, 0x54, 0xb1, 0xa2, 0x67,
	0x97, 0x5c, 0x21, 0xd0, 0x6c, 0x33, 0xdd, 0x5a, 0x38, 0x1e, 0x63, 0x8f, 0x49, 0xf3, 0x30, 0xdc,
	0x73, 0xc1, 0x83, 0xf0, 0x1a, 0x5c, 0xf3, 0x12, 0x68, 0xce, 0x78, 0x62, 0xc7, 0xdb, 0xf2, 0xa7,
	0xde, 0x9d, 0xff, 0xf9, 0xbe, 0x39, 0x67, 0x8e, 0x0d, 0xc7, 0x37, 0x2a, 0x4f, 0x52, 0x55, 0x8a,
	0x1f, 0xf2, 0xeb, 0xea, 0xd9, 0xc7, 0x79, 0x75, 0x9d, 0x26, 0x4f, 0x4f, 0xf3, 0x42, 0x69, 0xc5,
	0x27, 0x7b, 0xbe, 0xe8, 0x7b, 0xf0, 0x50, 0x6d, 0x78, 0x08, 0xc3, 0x87, 0x2a, 0xad, 0xd6, 0x59,
	0x19, 0xb2, 0xa9, 0x17, 0xfb, 0xe8, 0x54, 0xce, 0xc1, 0x7f, 0x2c, 0xb7, 0x65, 0xe8, 0x4d, 0xbd,
	0x78, 0x84, 0x24, 0xf3, 0x7b, 0xd0, 0xff, 0x52, 0xeb, 0xa2, 0x0c, 0x7b, 0x53, 0x2f, 0x1e, 0xcf,
	0xde, 0x3c, 0xdd, 0xab, 0x79, 0x6a, 0x7c, 0x68, 0x23, 0xa2, 0x4f, 0xe1, 0x10, 0xd5, 0x66, 0xb1,
	0x92, 0x99, 0x4e, 0x9e, 0x25, 0xb2, 0xa0, 0x82, 0xa8, 0x36, 0xee, 0x1c, 0x92, 0x77, 0x87, 0xf4,
	0x9a, 0x43, 0xa2, 0x2f, 0xc0, 0xff, 0x46, 0x24, 0x05, 0x3f, 0x84, 0xde, 0x62, 0x1e, 0xb2, 0x29,
	0x8b, 0x7d, 0xec, 0x2d, 0xe6, 0xfc, 0x0e, 0x78, 0x8f, 0xe5, 0x36, 0xf4, 0xa6, 0x2c, 0x1e, 0xa1,
	0x11, 0xf9, 0x11, 0xf4, 0x1f, 0xaa, 0x2a, 0xd3, 0x61, 0x8f, 0x82, 0xac, 0x12, 0x7d, 0x02, 0xc1,
	0xa3, 0x44, 0xa6, 0x2b, 0x43, 0xef, 0x08, 0xfa, 0x24, 0x53, 0x99, 0x11, 0x5a, 0xc5, 0x58, 0x0d,
	0xb6, 0xb9, 0xcb, 0x23, 0x25, 0x7a, 0x02, 0xf0, 0x55, 0xa1, 0xaa, 0x9c, 0xaa, 0xf0, 0x8f, 0xa0,
	0x4f, 0x1a, 0xc1, 0x1d, 0xcf, 0xde, 0xee, 0x50, 0x75, 0x27, 0xa0, 0x8d, 0x7a, 0x05, 0x94, 0x19,
	0x04, 0x4b, 0x91, 0xda, 0x82, 0x77, 0xc0, 0x5b, 0x8a, 0x94, 0x80, 0x78, 0x68, 0xc4, 0xfd, 0x1c,
	0xcf, 0xe5, 0x7c, 0x0b, 0xde, 0x59, 0xa2, 0x1b, 0x8c, 0xac, 0x85, 0x91, 0x1f, 0x43, 0x60, 0xfb,
	0xb3, 0x03, 0xbf, 0xd3, 0xf9, 0x09, 0x8c, 0xae, 0x92, 0xb5, 0x2c, 0xb5, 0x58, 0xe7, 0x74, 0x4b,
	0x1e, 0x36, 0x86, 0xe8, 0x3b, 0x98, 0xd8, 0x48, 0xd3, 0x9e, 0x4b, 0xa9, 0xff, 0xc5, 0xf5, 0xfe,
	0x87, 0x6e, 0xff, 0xca, 0xc0, 0x37, 0x92, 0xab, 0xc2, 0x9a, 0x2a, 0x1c, 0xfc, 0xab, 0x6d, 0x2e,
	0x6b, 0xb8, 0x24, 0xf3, 0x29, 0x8c, 0x2f, 0x75, 0x91, 0x64, 0x37, 0x4b, 0x91, 0x56, 0xb2, 0x3e,
	0xb3, 0x6d, 0x32, 0x44, 0x17, 0x99, 0xb6, 0x6e, 0x9f, 0xb8, 0xec, 0x74, 0x43, 0xf4, 0x4c, 0xa9,
	0xd4, 0x3a, 0xfb, 0x53, 0x16, 0x07, 0xd8, 0x18, 0xf8, 0x7b, 0x00, 0x8f, 0x52, 0x25, 0xea, 0xdc,
	0xc1, 0x94, 0xc5, 0x0c, 0x5b, 0x96, 0xe8, 0x01, 0x0c, 0x0d, 0xd2, 0xaf, 0x45, 0xde, 0x10, 0x64,
	0xff, 0x48, 0xf0, 0x77, 0x06, 0x07, 0x4f, 0x2a, 0x59, 0x6c, 0x51, 0xfe, 0x54, 0xc9, 0x92, 0xfa,
	0x43, 0xba, 0x9b, 0x2c, 0x52, 0xf8, 0x5d, 0x18, 0x5c, 0x3e, 0x17, 0xc5, 0xca, 0xde, 0x99, 0x8f,
	0xb5, 0x66, 0x08, 0x37, 0xb7, 0x5f, 0x12, 0xe1, 0x00, 0xdb, 0x26, 0x93, 0x89, 0x72, 0xad, 0xb4,
	0x63, 0x54, 0x6b, 0x3c, 0x86, 0x37, 0xce, 0x5f, 0x3c, 0x4d, 0xab, 0x95, 0x44, 0xb5, 0xb1, 0xd9,
	0x03, 0x0a, 0xe8, 0x9a, 0xf9, 0x07, 0x70, 0x58, 0x9b, 0xdc, 0x8b, 0x1e, 0x52, 0x60, 0xc7, 0x1a,
	0xfd, 0xc2, 0x60, 0x52, 0x53, 0x29, 0x73, 0x95, 0x95, 0xd2, 0x34, 0xed, 0xbc, 0x28, 0x5c, 0xd3,
	0xce, 0x8b, 0x82, 0x3f, 0x80, 0x21, 0xca, 0xb2, 0x4a, 0xb5, 0x6b, 0xfe, 0x71, 0xe7, 0x6e, 0x5c,
	0x81, 0x2a, 0xd5, 0xe8, 0x42, 0xf9, 0x1c, 0x0e, 0xf7, 0x66, 0xcc, 0x2e, 0x8f, 0xf1, 0xec, 0xa4,
	0x93, 0xbc, 0x17, 0x84, 0x9d, 0x9c, 0xe8, 0xcf, 0x1e, 0x8c, 0x5b, 0xe5, 0x77, 0x03, 0x64, 0x68,
	0x4f, 0xea, 0x01, 0x7a, 0x9f, 0xb6, 0x17, 0x21, 0x1e, 0xcf, 0x78, 0xa7, 0xbc, 0x79, 0x96, 0xb4,
	0xdc, 0x0e, 0x80, 0x5d, 0xd4, 0x73, 0xc7, 0x2e, 0x4c, 0xb7, 0xcd, 0x5e, 0x71, 0xa0, 0xba, 0xdd,
	0x36, 0x3e, 0xb4, 0x11, 0xb4, 0x15, 0x9f, 0x8b, 0xec, 0x46, 0xae, 0x68, 0xf8, 0x02, 0x74, 0x2a,
	0xbf, 0xdf, 0xbc, 0x68, 0x6a, 0xd4, 0xed, 0xcd, 0xe0, 0xdc, 0xd8, 0x3c, 0x7d, 0xd3, 0x5b, 0xf3,
	0x7c, 0x4d, 0x47, 0x68, 0x2a, 0xac, 0xc6, 0x3f, 0x83, 0x71, 0xb3, 0x71, 0xca, 0x30, 0x20, 0x5c,
	0xef, 0x74, 0xea, 0x35, 0x11, 0xd8, 0x8e, 0xe6, 0xe7, 0xdd, 0x05, 0x1b, 0x8e, 0x08, 0xcf, 0xbb,
	0xb7, 0x6f, 0xa3, 0x15, 0x84, 0x9d, 0xa4, 0xe8, 0x0f, 0x06, 0x93, 0xc5, 0x3a, 0x57, 0x85, 0x6e,
	0x4d, 0xf6, 0x22, 0x5b, 0xc9, 0x17, 0x6e, 0xb2, 0x49, 0x69, 0x36, 0x69, 0xaf, 0xb3, 0x49, 0x69,
	0xc2, 0x69, 0xa2, 0x7d, 0xb4, 0x4a, 0x8b, 0xaf, 0xbf, 0xc7, 0xf7, 0x04, 0x46, 0x6e, 0x5b, 0x95,
	0x61, 0x9f, 0x5c, 0x8d, 0xc1, 0x5c, 0x3a, 0xaa, 0x0d, 0x7d, 0x0e, 0x86, 0xf4, 0x39, 0x70, 0xaa,
	0x79, 0xd2, 0x36, 0x8c, 0x9c, 0x01, 0x39, 0x5b, 0x16, 0xe3, 0xdf, 0x2d, 0x3a, 0xf3, 0x3c, 0xbc,
	0xd8, 0xc3, 0x96, 0x25, 0xfa, 0x8d, 0x01, 0xb7, 0x1c, 0x69, 0x05, 0xbc, 0x3e, 0xa2, 0x7f, 0x4f,
	0x68, 0x1f, 0xf6, 0xf0, 0x16, 0xec, 0xbb, 0x30, 0x20, 0x3c, 0x0e, 0x72, 0xad, 0x45, 0x4b, 0x38,
	0xba, 0x2a, 0x44, 0x56, 0xa6, 0x42, 0x4b, 0x13, 0xf8, 0x7f, 0xf0, 0xbe, 0xe4, 0xeb, 0x1d, 0xdd,
	0x83, 0xb7, 0x3a, 0x75, 0x9b, 0xf7, 0x6f, 0x08, 0x78, 0x44, 0xc0, 0x88, 0xd1, 0x19, 0x84, 0xf5,
	0x50, 0x28, 0x61, 0x96, 0x72, 0x0d, 0x61, 0x99, 0xc8, 0x8d, 0x29, 0x7d, 0x21, 0xd6, 0xb2, 0x46,
	0x41, 0xb2, 0xb1, 0xcd, 0x85, 0x16, 0x84, 0xe1, 0x00, 0x49, 0x8e, 0x7e, 0x84, 0xa3, 0x97, 0xd5,
	0xa0, 0xcf, 0x5e, 0x2a, 0x85, 0xdd, 0x37, 0x01, 0x5a, 0x85, 0x7f, 0x0e, 0xfd, 0x9f, 0x13, 0xb9,
	0x71, 0xfb, 0xe6, 0xc3, 0xce, 0x14, 0xbf, 0x0a, 0x0d, 0xda, 0xac, 0xeb, 0x01, 0xfd, 0xe4, 0xdc,
	0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x91, 0x59, 0x0b, 0xff, 0x02, 0x09, 0x00, 0x00,
}
