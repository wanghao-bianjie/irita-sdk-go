// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/record/types.proto

package record

import (
	bytes "bytes"
	fmt "fmt"
	github_com_bianjieai_irita_sdk_go_types "github.com/bianjieai/irita-sdk-go/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
	io "io"
	math "math"
	math_bits "math/bits"
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

// MsgCreateRecord defines an SDK message for creating a new record.
type MsgCreateRecord struct {
	Contents []Content                                          `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents"`
	Creator  github_com_bianjieai_irita_sdk_go_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/bianjieai/irita-sdk-go/types.AccAddress" json:"creator,omitempty"`
}

func (m *MsgCreateRecord) Reset()         { *m = MsgCreateRecord{} }
func (m *MsgCreateRecord) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRecord) ProtoMessage()    {}
func (*MsgCreateRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_7144fc6efa27d236, []int{0}
}
func (m *MsgCreateRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRecord.Merge(m, src)
}
func (m *MsgCreateRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRecord proto.InternalMessageInfo

func (m *MsgCreateRecord) GetContents() []Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *MsgCreateRecord) GetCreator() github_com_bianjieai_irita_sdk_go_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

// Content defines a detail information for a record.
type Content struct {
	Digest     string `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	DigestAlgo string `protobuf:"bytes,2,opt,name=digest_algo,json=digestAlgo,proto3" json:"digest_algo,omitempty" yaml:"digest_algo"`
	URI        string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty" yaml:"uri"`
	Meta       string `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_7144fc6efa27d236, []int{1}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return m.Size()
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Content) GetDigestAlgo() string {
	if m != nil {
		return m.DigestAlgo
	}
	return ""
}

func (m *Content) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *Content) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type Record struct {
	TxHash   github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"tx_hash,omitempty" yaml:"tx_hash"`
	Contents []Content                                            `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents"`
	Creator  github_com_bianjieai_irita_sdk_go_types.AccAddress   `protobuf:"bytes,3,opt,name=creator,proto3,casttype=github.com/bianjieai/irita-sdk-go/types.AccAddress" json:"creator,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_7144fc6efa27d236, []int{2}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTxHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Record) GetContents() []Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Record) GetCreator() github_com_bianjieai_irita_sdk_go_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateRecord)(nil), "irita.modules.record.MsgCreateRecord")
	proto.RegisterType((*Content)(nil), "irita.modules.record.Content")
	proto.RegisterType((*Record)(nil), "irita.modules.record.Record")
}

func init() { proto.RegisterFile("modules/record/types.proto", fileDescriptor_7144fc6efa27d236) }

var fileDescriptor_7144fc6efa27d236 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3f, 0x8b, 0xd4, 0x40,
	0x14, 0xdf, 0xd9, 0x2c, 0xbb, 0xde, 0xdc, 0xa1, 0x30, 0xe8, 0x11, 0x0e, 0x4c, 0x96, 0x14, 0xba,
	0xcd, 0x25, 0x72, 0x8a, 0xc2, 0x35, 0xb2, 0xb9, 0xe6, 0x3c, 0x10, 0x24, 0x60, 0x63, 0xb3, 0x4c,
	0x32, 0xc3, 0x64, 0x34, 0xc9, 0x2c, 0x33, 0x6f, 0x61, 0xf7, 0x0b, 0x58, 0xfb, 0x05, 0x04, 0x6b,
	0x3f, 0xc9, 0x95, 0x57, 0x5a, 0x05, 0xc9, 0x36, 0xd6, 0x5b, 0x5e, 0x25, 0x3b, 0x89, 0xba, 0x07,
	0x82, 0x85, 0xd7, 0xfd, 0x86, 0xf7, 0xfb, 0xf3, 0xe6, 0xbd, 0x87, 0x8f, 0x4a, 0xc5, 0x16, 0x05,
	0x37, 0x91, 0xe6, 0x99, 0xd2, 0x2c, 0x82, 0xd5, 0x9c, 0x9b, 0x70, 0xae, 0x15, 0x28, 0x72, 0x5f,
	0x6a, 0x09, 0x34, 0xec, 0x18, 0x61, 0xcb, 0x38, 0x7a, 0x04, 0xb9, 0xd4, 0x6c, 0x36, 0xa7, 0x1a,
	0x56, 0x91, 0x25, 0x46, 0x42, 0x09, 0xf5, 0x07, 0xb5, 0xea, 0xe0, 0x2b, 0xc2, 0xf7, 0x5e, 0x1b,
	0x71, 0xa6, 0x39, 0x05, 0x9e, 0x58, 0x2d, 0x79, 0x89, 0xef, 0x64, 0xaa, 0x02, 0x5e, 0x81, 0x71,
	0xd1, 0xd8, 0x99, 0xec, 0x9f, 0x3c, 0x0c, 0xff, 0x16, 0x12, 0x9e, 0xb5, 0xac, 0x78, 0x70, 0x59,
	0xfb, 0xbd, 0xe4, 0xb7, 0x88, 0xbc, 0xc1, 0xa3, 0x6c, 0x6b, 0xa8, 0xb4, 0xdb, 0x1f, 0xa3, 0xc9,
	0x41, 0xfc, 0xfc, 0xba, 0xf6, 0x4f, 0x84, 0x84, 0x7c, 0x91, 0x86, 0x99, 0x2a, 0xa3, 0x54, 0xd2,
	0xea, 0xbd, 0xe4, 0x54, 0x46, 0xd6, 0xf7, 0xd8, 0xb0, 0x0f, 0xc7, 0x42, 0x75, 0xdf, 0x9a, 0x66,
	0xd9, 0x94, 0x31, 0xcd, 0x8d, 0x49, 0x7e, 0xd9, 0x9c, 0x0e, 0x7e, 0x7c, 0xf1, 0x51, 0xf0, 0x19,
	0xe1, 0x51, 0x97, 0x49, 0x0e, 0xf1, 0x90, 0x49, 0xc1, 0x0d, 0xb8, 0x68, 0x8c, 0x26, 0x7b, 0x49,
	0xf7, 0x22, 0x2f, 0xf0, 0x7e, 0x8b, 0x66, 0xb4, 0x10, 0xca, 0xe6, 0xef, 0xc5, 0x87, 0x9b, 0xda,
	0x27, 0x2b, 0x5a, 0x16, 0xa7, 0xc1, 0x4e, 0x31, 0x48, 0x70, 0xfb, 0x9a, 0x16, 0x42, 0x91, 0xc7,
	0xd8, 0x59, 0x68, 0xe9, 0x3a, 0x56, 0xf0, 0xa0, 0xa9, 0x7d, 0xe7, 0x6d, 0xf2, 0x6a, 0x53, 0xfb,
	0xb8, 0xd5, 0x2d, 0xb4, 0x0c, 0x92, 0x2d, 0x83, 0x10, 0x3c, 0x28, 0x39, 0x50, 0x77, 0x60, 0x73,
	0x2d, 0xee, 0xfa, 0xfb, 0xd8, 0xc7, 0xc3, 0x6e, 0x86, 0x19, 0x1e, 0xc1, 0x72, 0x96, 0x53, 0x93,
	0xdb, 0xfe, 0x0e, 0xe2, 0x8b, 0x4d, 0xed, 0xdf, 0x6d, 0xad, 0xba, 0x42, 0x70, 0x5d, 0xfb, 0xcf,
	0x76, 0x86, 0x02, 0xbc, 0x62, 0x5c, 0x97, 0xb2, 0x82, 0x5d, 0x58, 0xc8, 0xd4, 0x44, 0xe9, 0x0a,
	0xb8, 0x09, 0xcf, 0xf9, 0x32, 0xde, 0x82, 0x64, 0x08, 0xcb, 0x73, 0x6a, 0xf2, 0x1b, 0x8b, 0xea,
	0xff, 0xe7, 0xa2, 0x9c, 0x5b, 0x5c, 0x54, 0x7c, 0x71, 0xd9, 0x78, 0xe8, 0xaa, 0xf1, 0xd0, 0xf7,
	0xc6, 0x43, 0x9f, 0xd6, 0x5e, 0xef, 0x6a, 0xed, 0xf5, 0xbe, 0xad, 0xbd, 0xde, 0xbb, 0x27, 0xff,
	0x36, 0xbf, 0x79, 0xeb, 0xe9, 0xd0, 0x1e, 0xea, 0xd3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x2a, 0x5f, 0x8d, 0x04, 0x03, 0x00, 0x00,
}

func (this *MsgCreateRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateRecord)
	if !ok {
		that2, ok := that.(MsgCreateRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Contents) != len(that1.Contents) {
		return false
	}
	for i := range this.Contents {
		if !this.Contents[i].Equal(&that1.Contents[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	return true
}
func (this *Content) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Content)
	if !ok {
		that2, ok := that.(Content)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Digest != that1.Digest {
		return false
	}
	if this.DigestAlgo != that1.DigestAlgo {
		return false
	}
	if this.URI != that1.URI {
		return false
	}
	if this.Meta != that1.Meta {
		return false
	}
	return true
}
func (this *Record) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Record)
	if !ok {
		that2, ok := that.(Record)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	if len(this.Contents) != len(that1.Contents) {
		return false
	}
	for i := range this.Contents {
		if !this.Contents[i].Equal(&that1.Contents[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	return true
}
func (m *MsgCreateRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contents) > 0 {
		for iNdEx := len(m.Contents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DigestAlgo) > 0 {
		i -= len(m.DigestAlgo)
		copy(dAtA[i:], m.DigestAlgo)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DigestAlgo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contents) > 0 {
		for iNdEx := len(m.Contents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DigestAlgo)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, Content{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DigestAlgo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, Content{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)