// Code generated by protoc-gen-go.
// source: google/spanner/v1/keys.proto
// DO NOT EDIT!

/*
Package spanner is a generated protocol buffer package.

It is generated from these files:
	google/spanner/v1/keys.proto
	google/spanner/v1/mutation.proto
	google/spanner/v1/query_plan.proto
	google/spanner/v1/result_set.proto
	google/spanner/v1/spanner.proto
	google/spanner/v1/transaction.proto
	google/spanner/v1/type.proto

It has these top-level messages:
	KeyRange
	KeySet
	Mutation
	PlanNode
	QueryPlan
	ResultSet
	PartialResultSet
	ResultSetMetadata
	ResultSetStats
	CreateSessionRequest
	Session
	GetSessionRequest
	DeleteSessionRequest
	ExecuteSqlRequest
	ReadRequest
	BeginTransactionRequest
	CommitRequest
	CommitResponse
	RollbackRequest
	TransactionOptions
	Transaction
	TransactionSelector
	Type
	StructType
*/
package spanner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KeyRange represents a range of rows in a table or index.
//
// A range has a start key and an end key. These keys can be open or
// closed, indicating if the range includes rows with that key.
//
// Keys are represented by lists, where the ith value in the list
// corresponds to the ith component of the table or index primary key.
// Individual values are encoded as described [here][google.spanner.v1.TypeCode].
//
// For example, consider the following table definition:
//
//     CREATE TABLE UserEvents (
//       UserName STRING(MAX),
//       EventDate STRING(10)
//     ) PRIMARY KEY(UserName, EventDate);
//
// The following keys name rows in this table:
//
//     ["Bob", "2014-09-23"]
//     ["Alfred", "2015-06-12"]
//
// Since the `UserEvents` table's `PRIMARY KEY` clause names two
// columns, each `UserEvents` key has two elements; the first is the
// `UserName`, and the second is the `EventDate`.
//
// Key ranges with multiple components are interpreted
// lexicographically by component using the table or index key's declared
// sort order. For example, the following range returns all events for
// user `"Bob"` that occurred in the year 2015:
//
//     "start_closed": ["Bob", "2015-01-01"]
//     "end_closed": ["Bob", "2015-12-31"]
//
// Start and end keys can omit trailing key components. This affects the
// inclusion and exclusion of rows that exactly match the provided key
// components: if the key is closed, then rows that exactly match the
// provided components are included; if the key is open, then rows
// that exactly match are not included.
//
// For example, the following range includes all events for `"Bob"` that
// occurred during and after the year 2000:
//
//     "start_closed": ["Bob", "2000-01-01"]
//     "end_closed": ["Bob"]
//
// The next example retrieves all events for `"Bob"`:
//
//     "start_closed": ["Bob"]
//     "end_closed": ["Bob"]
//
// To retrieve events before the year 2000:
//
//     "start_closed": ["Bob"]
//     "end_open": ["Bob", "2000-01-01"]
//
// The following range includes all rows in the table:
//
//     "start_closed": []
//     "end_closed": []
//
// This range returns all users whose `UserName` begins with any
// character from A to C:
//
//     "start_closed": ["A"]
//     "end_open": ["D"]
//
// This range returns all users whose `UserName` begins with B:
//
//     "start_closed": ["B"]
//     "end_open": ["C"]
//
// Key ranges honor column sort order. For example, suppose a table is
// defined as follows:
//
//     CREATE TABLE DescendingSortedTable {
//       Key INT64,
//       ...
//     ) PRIMARY KEY(Key DESC);
//
// The following range retrieves all rows with key values between 1
// and 100 inclusive:
//
//     "start_closed": ["100"]
//     "end_closed": ["1"]
//
// Note that 100 is passed as the start, and 1 is passed as the end,
// because `Key` is a descending column in the schema.
type KeyRange struct {
	// The start key must be provided. It can be either closed or open.
	//
	// Types that are valid to be assigned to StartKeyType:
	//	*KeyRange_StartClosed
	//	*KeyRange_StartOpen
	StartKeyType isKeyRange_StartKeyType `protobuf_oneof:"start_key_type"`
	// The end key must be provided. It can be either closed or open.
	//
	// Types that are valid to be assigned to EndKeyType:
	//	*KeyRange_EndClosed
	//	*KeyRange_EndOpen
	EndKeyType isKeyRange_EndKeyType `protobuf_oneof:"end_key_type"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isKeyRange_StartKeyType interface {
	isKeyRange_StartKeyType()
}
type isKeyRange_EndKeyType interface {
	isKeyRange_EndKeyType()
}

type KeyRange_StartClosed struct {
	StartClosed *google_protobuf1.ListValue `protobuf:"bytes,1,opt,name=start_closed,json=startClosed,oneof"`
}
type KeyRange_StartOpen struct {
	StartOpen *google_protobuf1.ListValue `protobuf:"bytes,2,opt,name=start_open,json=startOpen,oneof"`
}
type KeyRange_EndClosed struct {
	EndClosed *google_protobuf1.ListValue `protobuf:"bytes,3,opt,name=end_closed,json=endClosed,oneof"`
}
type KeyRange_EndOpen struct {
	EndOpen *google_protobuf1.ListValue `protobuf:"bytes,4,opt,name=end_open,json=endOpen,oneof"`
}

func (*KeyRange_StartClosed) isKeyRange_StartKeyType() {}
func (*KeyRange_StartOpen) isKeyRange_StartKeyType()   {}
func (*KeyRange_EndClosed) isKeyRange_EndKeyType()     {}
func (*KeyRange_EndOpen) isKeyRange_EndKeyType()       {}

func (m *KeyRange) GetStartKeyType() isKeyRange_StartKeyType {
	if m != nil {
		return m.StartKeyType
	}
	return nil
}
func (m *KeyRange) GetEndKeyType() isKeyRange_EndKeyType {
	if m != nil {
		return m.EndKeyType
	}
	return nil
}

func (m *KeyRange) GetStartClosed() *google_protobuf1.ListValue {
	if x, ok := m.GetStartKeyType().(*KeyRange_StartClosed); ok {
		return x.StartClosed
	}
	return nil
}

func (m *KeyRange) GetStartOpen() *google_protobuf1.ListValue {
	if x, ok := m.GetStartKeyType().(*KeyRange_StartOpen); ok {
		return x.StartOpen
	}
	return nil
}

func (m *KeyRange) GetEndClosed() *google_protobuf1.ListValue {
	if x, ok := m.GetEndKeyType().(*KeyRange_EndClosed); ok {
		return x.EndClosed
	}
	return nil
}

func (m *KeyRange) GetEndOpen() *google_protobuf1.ListValue {
	if x, ok := m.GetEndKeyType().(*KeyRange_EndOpen); ok {
		return x.EndOpen
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeyRange) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeyRange_OneofMarshaler, _KeyRange_OneofUnmarshaler, _KeyRange_OneofSizer, []interface{}{
		(*KeyRange_StartClosed)(nil),
		(*KeyRange_StartOpen)(nil),
		(*KeyRange_EndClosed)(nil),
		(*KeyRange_EndOpen)(nil),
	}
}

func _KeyRange_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeyRange)
	// start_key_type
	switch x := m.StartKeyType.(type) {
	case *KeyRange_StartClosed:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartClosed); err != nil {
			return err
		}
	case *KeyRange_StartOpen:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartOpen); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KeyRange.StartKeyType has unexpected type %T", x)
	}
	// end_key_type
	switch x := m.EndKeyType.(type) {
	case *KeyRange_EndClosed:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EndClosed); err != nil {
			return err
		}
	case *KeyRange_EndOpen:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EndOpen); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KeyRange.EndKeyType has unexpected type %T", x)
	}
	return nil
}

func _KeyRange_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeyRange)
	switch tag {
	case 1: // start_key_type.start_closed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.ListValue)
		err := b.DecodeMessage(msg)
		m.StartKeyType = &KeyRange_StartClosed{msg}
		return true, err
	case 2: // start_key_type.start_open
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.ListValue)
		err := b.DecodeMessage(msg)
		m.StartKeyType = &KeyRange_StartOpen{msg}
		return true, err
	case 3: // end_key_type.end_closed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.ListValue)
		err := b.DecodeMessage(msg)
		m.EndKeyType = &KeyRange_EndClosed{msg}
		return true, err
	case 4: // end_key_type.end_open
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.ListValue)
		err := b.DecodeMessage(msg)
		m.EndKeyType = &KeyRange_EndOpen{msg}
		return true, err
	default:
		return false, nil
	}
}

func _KeyRange_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeyRange)
	// start_key_type
	switch x := m.StartKeyType.(type) {
	case *KeyRange_StartClosed:
		s := proto.Size(x.StartClosed)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeyRange_StartOpen:
		s := proto.Size(x.StartOpen)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// end_key_type
	switch x := m.EndKeyType.(type) {
	case *KeyRange_EndClosed:
		s := proto.Size(x.EndClosed)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeyRange_EndOpen:
		s := proto.Size(x.EndOpen)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// `KeySet` defines a collection of Cloud Spanner keys and/or key ranges. All
// the keys are expected to be in the same table or index. The keys need
// not be sorted in any particular way.
//
// If the same key is specified multiple times in the set (for example
// if two ranges, two keys, or a key and a range overlap), Cloud Spanner
// behaves as if the key were only specified once.
type KeySet struct {
	// A list of specific keys. Entries in `keys` should have exactly as
	// many elements as there are columns in the primary or index key
	// with which this `KeySet` is used.  Individual key values are
	// encoded as described [here][google.spanner.v1.TypeCode].
	Keys []*google_protobuf1.ListValue `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	// A list of key ranges. See [KeyRange][google.spanner.v1.KeyRange] for more information about
	// key range specifications.
	Ranges []*KeyRange `protobuf:"bytes,2,rep,name=ranges" json:"ranges,omitempty"`
	// For convenience `all` can be set to `true` to indicate that this
	// `KeySet` matches all keys in the table or index. Note that any keys
	// specified in `keys` or `ranges` are only yielded once.
	All bool `protobuf:"varint,3,opt,name=all" json:"all,omitempty"`
}

func (m *KeySet) Reset()                    { *m = KeySet{} }
func (m *KeySet) String() string            { return proto.CompactTextString(m) }
func (*KeySet) ProtoMessage()               {}
func (*KeySet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeySet) GetKeys() []*google_protobuf1.ListValue {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *KeySet) GetRanges() []*KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func (m *KeySet) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "google.spanner.v1.KeyRange")
	proto.RegisterType((*KeySet)(nil), "google.spanner.v1.KeySet")
}

func init() { proto.RegisterFile("google/spanner/v1/keys.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xed, 0x36, 0xe6, 0x96, 0x8d, 0x31, 0x0b, 0x42, 0x99, 0x1e, 0xc6, 0x4e, 0x9e, 0x12,
	0xe6, 0x0e, 0x0a, 0x3b, 0x08, 0xf5, 0x22, 0x4c, 0x70, 0x54, 0xf0, 0xe0, 0xc1, 0x91, 0xad, 0xcf,
	0x50, 0x56, 0x5f, 0x42, 0x93, 0x0d, 0x7a, 0xf2, 0x5f, 0xf1, 0x4f, 0x95, 0xa4, 0xa9, 0x08, 0x03,
	0xdd, 0x2d, 0x8f, 0xef, 0xfb, 0xbd, 0xef, 0xf5, 0xbd, 0x92, 0x4b, 0x21, 0xa5, 0xc8, 0x81, 0x69,
	0xc5, 0x11, 0xa1, 0x60, 0xfb, 0x29, 0xdb, 0x42, 0xa9, 0xa9, 0x2a, 0xa4, 0x91, 0xe1, 0x59, 0xa5,
	0x52, 0xaf, 0xd2, 0xfd, 0x74, 0x54, 0x03, 0x5c, 0x65, 0x8c, 0x23, 0x4a, 0xc3, 0x4d, 0x26, 0xd1,
	0x03, 0x3f, 0xaa, 0xab, 0xd6, 0xbb, 0x77, 0xa6, 0x4d, 0xb1, 0xdb, 0x98, 0x4a, 0x9d, 0x7c, 0x35,
	0x48, 0x67, 0x01, 0x65, 0xc2, 0x51, 0x40, 0x78, 0x47, 0xfa, 0xda, 0xf0, 0xc2, 0xac, 0x36, 0xb9,
	0xd4, 0x90, 0x46, 0xc1, 0x38, 0xb8, 0xea, 0x5d, 0x8f, 0xa8, 0x8f, 0xac, 0x3b, 0xd0, 0xc7, 0x4c,
	0x9b, 0x17, 0x9e, 0xef, 0xe0, 0xe1, 0x24, 0xe9, 0x39, 0xe2, 0xde, 0x01, 0xe1, 0x9c, 0x90, 0xaa,
	0x81, 0x54, 0x80, 0x51, 0xe3, 0x08, 0xbc, 0xeb, 0xfc, 0x4f, 0x0a, 0xd0, 0xc2, 0x80, 0x69, 0x9d,
	0xdd, 0xfc, 0x17, 0x0e, 0x92, 0x2e, 0x60, 0xea, 0x93, 0x6f, 0x48, 0xc7, 0xc2, 0x2e, 0xb7, 0x75,
	0x04, 0x7a, 0x0a, 0x98, 0xda, 0xd4, 0x78, 0x48, 0x06, 0xd5, 0xc8, 0x5b, 0x28, 0x57, 0xa6, 0x54,
	0x10, 0x0f, 0x48, 0xdf, 0xb6, 0xaa, 0xeb, 0xc9, 0x27, 0x69, 0x2f, 0xa0, 0x7c, 0x06, 0x13, 0x52,
	0xd2, 0xb2, 0x97, 0x88, 0x82, 0x71, 0xf3, 0xef, 0x80, 0xc4, 0xf9, 0xc2, 0x19, 0x69, 0x17, 0x76,
	0xb1, 0x3a, 0x6a, 0x38, 0xe2, 0x82, 0x1e, 0x1c, 0x8f, 0xd6, 0xcb, 0x4f, 0xbc, 0x35, 0x1c, 0x92,
	0x26, 0xcf, 0x73, 0xf7, 0xfd, 0x9d, 0xc4, 0x3e, 0xe3, 0x37, 0x72, 0xbe, 0x91, 0x1f, 0x87, 0x6c,
	0xdc, 0x5d, 0x40, 0xa9, 0x97, 0x36, 0x7d, 0x19, 0xbc, 0xde, 0x7a, 0x5d, 0xc8, 0x9c, 0xa3, 0xa0,
	0xb2, 0x10, 0x4c, 0x00, 0xba, 0xd9, 0x58, 0x25, 0x71, 0x95, 0xe9, 0x5f, 0x7f, 0xd5, 0xdc, 0x3f,
	0xd7, 0x6d, 0x67, 0x9a, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x83, 0x58, 0x82, 0xb7, 0x79, 0x02,
	0x00, 0x00,
}
