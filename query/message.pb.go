// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query/message.proto

package query

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	execution "go.temporal.io/temporal-proto/execution"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type WorkflowQuery struct {
	QueryType string `protobuf:"bytes,1,opt,name=queryType,proto3" json:"queryType,omitempty"`
	QueryArgs []byte `protobuf:"bytes,2,opt,name=queryArgs,proto3" json:"queryArgs,omitempty"`
}

func (m *WorkflowQuery) Reset()      { *m = WorkflowQuery{} }
func (*WorkflowQuery) ProtoMessage() {}
func (*WorkflowQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ad58129da10f31, []int{0}
}
func (m *WorkflowQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowQuery.Merge(m, src)
}
func (m *WorkflowQuery) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowQuery.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowQuery proto.InternalMessageInfo

func (m *WorkflowQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *WorkflowQuery) GetQueryArgs() []byte {
	if m != nil {
		return m.QueryArgs
	}
	return nil
}

type WorkflowQueryResult struct {
	ResultType   QueryResultType `protobuf:"varint,1,opt,name=resultType,proto3,enum=query.QueryResultType" json:"resultType,omitempty"`
	Answer       []byte          `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	ErrorMessage string          `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (m *WorkflowQueryResult) Reset()      { *m = WorkflowQueryResult{} }
func (*WorkflowQueryResult) ProtoMessage() {}
func (*WorkflowQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ad58129da10f31, []int{1}
}
func (m *WorkflowQueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowQueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowQueryResult.Merge(m, src)
}
func (m *WorkflowQueryResult) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowQueryResult proto.InternalMessageInfo

func (m *WorkflowQueryResult) GetResultType() QueryResultType {
	if m != nil {
		return m.ResultType
	}
	return QueryResultType_Answered
}

func (m *WorkflowQueryResult) GetAnswer() []byte {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *WorkflowQueryResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type QueryRejected struct {
	Status execution.WorkflowExecutionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=execution.WorkflowExecutionStatus" json:"status,omitempty"`
}

func (m *QueryRejected) Reset()      { *m = QueryRejected{} }
func (*QueryRejected) ProtoMessage() {}
func (*QueryRejected) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ad58129da10f31, []int{2}
}
func (m *QueryRejected) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRejected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRejected.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRejected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRejected.Merge(m, src)
}
func (m *QueryRejected) XXX_Size() int {
	return m.Size()
}
func (m *QueryRejected) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRejected.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRejected proto.InternalMessageInfo

func (m *QueryRejected) GetStatus() execution.WorkflowExecutionStatus {
	if m != nil {
		return m.Status
	}
	return execution.WorkflowExecutionStatus_Unknown
}

func init() {
	proto.RegisterType((*WorkflowQuery)(nil), "query.WorkflowQuery")
	proto.RegisterType((*WorkflowQueryResult)(nil), "query.WorkflowQueryResult")
	proto.RegisterType((*QueryRejected)(nil), "query.QueryRejected")
}

func init() { proto.RegisterFile("query/message.proto", fileDescriptor_08ad58129da10f31) }

var fileDescriptor_08ad58129da10f31 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0x3a, 0x41,
	0x10, 0xc6, 0x77, 0xfe, 0xff, 0x48, 0xc2, 0x06, 0x8d, 0x59, 0x0c, 0x12, 0x62, 0x26, 0xe4, 0x6c,
	0x68, 0x5c, 0x12, 0x4d, 0x2c, 0xec, 0x34, 0xb1, 0x32, 0x26, 0x7a, 0x9a, 0x98, 0xd8, 0x9d, 0x38,
	0x12, 0x14, 0x58, 0xdc, 0xdd, 0x0b, 0xd2, 0xd9, 0xda, 0xf9, 0x18, 0x3e, 0x8a, 0x25, 0x25, 0xa5,
	0x2c, 0x8d, 0x25, 0x8f, 0x60, 0xd8, 0x3b, 0x8e, 0xa3, 0xdb, 0xf9, 0x7d, 0xdf, 0x37, 0x3b, 0x33,
	0xbc, 0xfc, 0x1a, 0x93, 0x1e, 0x35, 0x7b, 0x64, 0x4c, 0xd4, 0x26, 0x39, 0xd0, 0xca, 0x2a, 0xb1,
	0xe1, 0x61, 0x6d, 0x3b, 0xd1, 0xa8, 0x1f, 0xf7, 0x12, 0xa1, 0xb6, 0x43, 0x6f, 0xd4, 0x8a, 0x6d,
	0x47, 0xf5, 0x73, 0x34, 0xb8, 0xe0, 0x9b, 0x77, 0x4a, 0xbf, 0x3c, 0x75, 0xd5, 0xf0, 0x7a, 0x91,
	0x10, 0x7b, 0xbc, 0xe8, 0xa3, 0xb7, 0xa3, 0x01, 0x55, 0xa1, 0x0e, 0x8d, 0x62, 0xb8, 0x02, 0x99,
	0x7a, 0xaa, 0xdb, 0xa6, 0xfa, 0xaf, 0x0e, 0x8d, 0x52, 0xb8, 0x02, 0xc1, 0x07, 0xf0, 0xf2, 0x5a,
	0xb7, 0x90, 0x4c, 0xdc, 0xb5, 0xe2, 0x98, 0x73, 0xed, 0x5f, 0x59, 0xd3, 0xad, 0xc3, 0x8a, 0xf4,
	0x39, 0x99, 0xf3, 0x2d, 0xd4, 0x30, 0xe7, 0x14, 0x15, 0x5e, 0x88, 0xfa, 0x66, 0x48, 0x3a, 0xfd,
	0x2a, 0xad, 0x44, 0xc0, 0x4b, 0xa4, 0xb5, 0xd2, 0x97, 0xc9, 0xe6, 0xd5, 0xff, 0x7e, 0xcc, 0x35,
	0xb6, 0x58, 0x2c, 0x6d, 0xfd, 0x4c, 0x2d, 0x4b, 0x8f, 0xe2, 0x84, 0x17, 0x8c, 0x8d, 0x6c, 0x6c,
	0xd2, 0x01, 0x02, 0x99, 0x1d, 0x44, 0x2e, 0x87, 0x3e, 0x5f, 0x92, 0x1b, 0xef, 0x0c, 0xd3, 0xc4,
	0x99, 0x1d, 0x4f, 0x91, 0x4d, 0xa6, 0xc8, 0xe6, 0x53, 0x84, 0x77, 0x87, 0xf0, 0xe5, 0x10, 0xbe,
	0x1d, 0xc2, 0xd8, 0x21, 0xfc, 0x38, 0x84, 0x5f, 0x87, 0x6c, 0xee, 0x10, 0x3e, 0x67, 0xc8, 0xc6,
	0x33, 0x64, 0x93, 0x19, 0x32, 0xbe, 0xdb, 0x51, 0xd2, 0x52, 0x6f, 0xa0, 0x74, 0xd4, 0x4d, 0xce,
	0x9d, 0xec, 0x7c, 0x05, 0xf7, 0xfb, 0xed, 0x9c, 0xd4, 0x51, 0xcd, 0xe5, 0xfb, 0xc0, 0xdb, 0x9a,
	0xde, 0xf6, 0x50, 0xf0, 0xc5, 0xd1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x2c, 0xa0, 0xb0,
	0xe8, 0x01, 0x00, 0x00,
}

func (this *WorkflowQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowQuery)
	if !ok {
		that2, ok := that.(WorkflowQuery)
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
	if this.QueryType != that1.QueryType {
		return false
	}
	if !bytes.Equal(this.QueryArgs, that1.QueryArgs) {
		return false
	}
	return true
}
func (this *WorkflowQueryResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowQueryResult)
	if !ok {
		that2, ok := that.(WorkflowQueryResult)
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
	if this.ResultType != that1.ResultType {
		return false
	}
	if !bytes.Equal(this.Answer, that1.Answer) {
		return false
	}
	if this.ErrorMessage != that1.ErrorMessage {
		return false
	}
	return true
}
func (this *QueryRejected) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryRejected)
	if !ok {
		that2, ok := that.(QueryRejected)
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
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *WorkflowQuery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&query.WorkflowQuery{")
	s = append(s, "QueryType: "+fmt.Sprintf("%#v", this.QueryType)+",\n")
	s = append(s, "QueryArgs: "+fmt.Sprintf("%#v", this.QueryArgs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WorkflowQueryResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&query.WorkflowQueryResult{")
	s = append(s, "ResultType: "+fmt.Sprintf("%#v", this.ResultType)+",\n")
	s = append(s, "Answer: "+fmt.Sprintf("%#v", this.Answer)+",\n")
	s = append(s, "ErrorMessage: "+fmt.Sprintf("%#v", this.ErrorMessage)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryRejected) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&query.QueryRejected{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *WorkflowQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryArgs) > 0 {
		i -= len(m.QueryArgs)
		copy(dAtA[i:], m.QueryArgs)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.QueryArgs)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkflowQueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowQueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowQueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Answer) > 0 {
		i -= len(m.Answer)
		copy(dAtA[i:], m.Answer)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Answer)))
		i--
		dAtA[i] = 0x12
	}
	if m.ResultType != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ResultType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRejected) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRejected) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRejected) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.QueryArgs)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *WorkflowQueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResultType != 0 {
		n += 1 + sovMessage(uint64(m.ResultType))
	}
	l = len(m.Answer)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *QueryRejected) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovMessage(uint64(m.Status))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WorkflowQuery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowQuery{`,
		`QueryType:` + fmt.Sprintf("%v", this.QueryType) + `,`,
		`QueryArgs:` + fmt.Sprintf("%v", this.QueryArgs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WorkflowQueryResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowQueryResult{`,
		`ResultType:` + fmt.Sprintf("%v", this.ResultType) + `,`,
		`Answer:` + fmt.Sprintf("%v", this.Answer) + `,`,
		`ErrorMessage:` + fmt.Sprintf("%v", this.ErrorMessage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryRejected) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryRejected{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WorkflowQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: WorkflowQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryArgs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryArgs = append(m.QueryArgs[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryArgs == nil {
				m.QueryArgs = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *WorkflowQueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: WorkflowQueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowQueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultType", wireType)
			}
			m.ResultType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultType |= QueryResultType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answer = append(m.Answer[:0], dAtA[iNdEx:postIndex]...)
			if m.Answer == nil {
				m.Answer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *QueryRejected) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: QueryRejected: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRejected: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= execution.WorkflowExecutionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)