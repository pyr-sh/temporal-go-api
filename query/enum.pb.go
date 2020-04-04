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
// source: query/enum.proto

package query

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type QueryResultType int32

const (
	QueryResultType_Answered QueryResultType = 0
	QueryResultType_Failed   QueryResultType = 1
)

var QueryResultType_name = map[int32]string{
	0: "Answered",
	1: "Failed",
}

var QueryResultType_value = map[string]int32{
	"Answered": 0,
	"Failed":   1,
}

func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{0}
}

type QueryRejectCondition int32

const (
	// None indicates that query should not be rejected.
	QueryRejectCondition_None QueryRejectCondition = 0
	// NotOpen indicates that query should be rejected if workflow is not open.
	QueryRejectCondition_NotOpen QueryRejectCondition = 1
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QueryRejectCondition_NotCompletedCleanly QueryRejectCondition = 2
)

var QueryRejectCondition_name = map[int32]string{
	0: "None",
	1: "NotOpen",
	2: "NotCompletedCleanly",
}

var QueryRejectCondition_value = map[string]int32{
	"None":                0,
	"NotOpen":             1,
	"NotCompletedCleanly": 2,
}

func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{1}
}

type QueryConsistencyLevel int32

const (
	// Eventual indicates that query should be eventually consistent.
	QueryConsistencyLevel_Eventual QueryConsistencyLevel = 0
	// Strong indicates that any events that came before query should be reflected in workflow state before running query.
	QueryConsistencyLevel_Strong QueryConsistencyLevel = 1
)

var QueryConsistencyLevel_name = map[int32]string{
	0: "Eventual",
	1: "Strong",
}

var QueryConsistencyLevel_value = map[string]int32{
	"Eventual": 0,
	"Strong":   1,
}

func (QueryConsistencyLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{2}
}

func init() {
	proto.RegisterEnum("query.QueryResultType", QueryResultType_name, QueryResultType_value)
	proto.RegisterEnum("query.QueryRejectCondition", QueryRejectCondition_name, QueryRejectCondition_value)
	proto.RegisterEnum("query.QueryConsistencyLevel", QueryConsistencyLevel_name, QueryConsistencyLevel_value)
}

func init() { proto.RegisterFile("query/enum.proto", fileDescriptor_4b1c9789806fe1e4) }

var fileDescriptor_4b1c9789806fe1e4 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0x77, 0x10, 0x84, 0x68, 0x41, 0x62, 0x65, 0x82, 0x22, 0x51, 0x4c, 0x43, 0x17, 0x44,
	0x22, 0xc4, 0x09, 0xc0, 0x22, 0x15, 0x0a, 0xbf, 0x15, 0x9d, 0x89, 0x47, 0x96, 0xd1, 0x66, 0xc7,
	0xd8, 0xe3, 0x20, 0x77, 0x1c, 0x81, 0x63, 0x70, 0x14, 0xca, 0x94, 0x29, 0xc9, 0xa6, 0xa1, 0xcc,
	0x11, 0x50, 0x6c, 0x90, 0xe8, 0xde, 0xbe, 0xef, 0x69, 0x9f, 0xe6, 0x69, 0xf3, 0x52, 0x52, 0x5e,
	0x0d, 0xc8, 0x95, 0x93, 0x7e, 0x96, 0xb3, 0x70, 0xb0, 0x55, 0x3b, 0x87, 0x9d, 0x84, 0x13, 0xae,
	0x9d, 0xc1, 0x5a, 0x35, 0xb0, 0x77, 0xac, 0xf7, 0x6e, 0xd7, 0xf8, 0x8e, 0x8a, 0xd2, 0xca, 0x43,
	0x95, 0x51, 0xb0, 0xab, 0xdb, 0xe7, 0xae, 0x78, 0xa5, 0x9c, 0x62, 0xa3, 0x02, 0xad, 0x5b, 0xc3,
	0x28, 0xb5, 0x14, 0x1b, 0xe8, 0x0d, 0x75, 0xe7, 0x37, 0xfc, 0x4c, 0x63, 0x09, 0xd9, 0xc5, 0xa9,
	0xa4, 0xec, 0x82, 0xb6, 0xde, 0x1c, 0xb1, 0x23, 0xa3, 0x82, 0x1d, 0xbd, 0x3d, 0x62, 0xb9, 0xce,
	0xc8, 0x19, 0x08, 0xba, 0x7a, 0x7f, 0xc4, 0x12, 0xf2, 0x24, 0xb3, 0x24, 0x14, 0x87, 0x96, 0x22,
	0x67, 0x2b, 0xb3, 0xd1, 0x3b, 0xd5, 0x07, 0xf5, 0x3f, 0x21, 0xbb, 0x22, 0x2d, 0x84, 0xdc, 0xb8,
	0xba, 0xa2, 0x29, 0xd9, 0x75, 0xf5, 0xe5, 0x94, 0x9c, 0x94, 0x91, 0x6d, 0xaa, 0xef, 0x25, 0x67,
	0x97, 0x18, 0xb8, 0x90, 0xd9, 0x02, 0x61, 0xbe, 0x40, 0xb5, 0x5a, 0x20, 0xbc, 0x79, 0x84, 0x0f,
	0x8f, 0xf0, 0xe9, 0x11, 0x66, 0x1e, 0xe1, 0xcb, 0x23, 0x7c, 0x7b, 0x54, 0x2b, 0x8f, 0xf0, 0xbe,
	0x44, 0x35, 0x5b, 0xa2, 0x9a, 0x2f, 0x51, 0xe9, 0x6e, 0xca, 0x7d, 0xa1, 0x49, 0xc6, 0x79, 0x64,
	0x9b, 0x83, 0xfb, 0xf5, 0x18, 0x37, 0xf0, 0x78, 0x94, 0xfc, 0x43, 0x29, 0x0f, 0xfe, 0xf4, 0x49,
	0x33, 0x51, 0x1d, 0x7b, 0x6a, 0xd5, 0x8f, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x1b,
	0x88, 0xd6, 0x55, 0x01, 0x00, 0x00,
}

func (x QueryResultType) String() string {
	s, ok := QueryResultType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryRejectCondition) String() string {
	s, ok := QueryRejectCondition_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryConsistencyLevel) String() string {
	s, ok := QueryConsistencyLevel_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}