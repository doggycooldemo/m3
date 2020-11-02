// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/wide (interfaces: CrossBlockIterator,QueryShardIterator,QuerySeriesIterator,BorrowedBuffer,FixedBufferManager)

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package wide is a generated GoMock package.
package wide

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/ts"
	time0 "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockCrossBlockIterator is a mock of CrossBlockIterator interface
type MockCrossBlockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockCrossBlockIteratorMockRecorder
}

// MockCrossBlockIteratorMockRecorder is the mock recorder for MockCrossBlockIterator
type MockCrossBlockIteratorMockRecorder struct {
	mock *MockCrossBlockIterator
}

// NewMockCrossBlockIterator creates a new mock instance
func NewMockCrossBlockIterator(ctrl *gomock.Controller) *MockCrossBlockIterator {
	mock := &MockCrossBlockIterator{ctrl: ctrl}
	mock.recorder = &MockCrossBlockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrossBlockIterator) EXPECT() *MockCrossBlockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCrossBlockIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCrossBlockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCrossBlockIterator)(nil).Close))
}

// Current mocks base method
func (m *MockCrossBlockIterator) Current() QuerySeriesIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(QuerySeriesIterator)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockCrossBlockIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockCrossBlockIterator)(nil).Current))
}

// Err mocks base method
func (m *MockCrossBlockIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockCrossBlockIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockCrossBlockIterator)(nil).Err))
}

// Next mocks base method
func (m *MockCrossBlockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockCrossBlockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCrossBlockIterator)(nil).Next))
}

// MockQueryShardIterator is a mock of QueryShardIterator interface
type MockQueryShardIterator struct {
	ctrl     *gomock.Controller
	recorder *MockQueryShardIteratorMockRecorder
}

// MockQueryShardIteratorMockRecorder is the mock recorder for MockQueryShardIterator
type MockQueryShardIteratorMockRecorder struct {
	mock *MockQueryShardIterator
}

// NewMockQueryShardIterator creates a new mock instance
func NewMockQueryShardIterator(ctrl *gomock.Controller) *MockQueryShardIterator {
	mock := &MockQueryShardIterator{ctrl: ctrl}
	mock.recorder = &MockQueryShardIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryShardIterator) EXPECT() *MockQueryShardIteratorMockRecorder {
	return m.recorder
}

// BlockStart mocks base method
func (m *MockQueryShardIterator) BlockStart() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockStart")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// BlockStart indicates an expected call of BlockStart
func (mr *MockQueryShardIteratorMockRecorder) BlockStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockStart", reflect.TypeOf((*MockQueryShardIterator)(nil).BlockStart))
}

// Close mocks base method
func (m *MockQueryShardIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockQueryShardIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQueryShardIterator)(nil).Close))
}

// Current mocks base method
func (m *MockQueryShardIterator) Current() QuerySeriesIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(QuerySeriesIterator)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockQueryShardIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockQueryShardIterator)(nil).Current))
}

// Err mocks base method
func (m *MockQueryShardIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockQueryShardIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockQueryShardIterator)(nil).Err))
}

// Next mocks base method
func (m *MockQueryShardIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockQueryShardIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockQueryShardIterator)(nil).Next))
}

// PushRecord mocks base method
func (m *MockQueryShardIterator) PushRecord(arg0 ShardIteratorRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushRecord", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushRecord indicates an expected call of PushRecord
func (mr *MockQueryShardIteratorMockRecorder) PushRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushRecord", reflect.TypeOf((*MockQueryShardIterator)(nil).PushRecord), arg0)
}

// Shard mocks base method
func (m *MockQueryShardIterator) Shard() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shard")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Shard indicates an expected call of Shard
func (mr *MockQueryShardIteratorMockRecorder) Shard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shard", reflect.TypeOf((*MockQueryShardIterator)(nil).Shard))
}

// MockQuerySeriesIterator is a mock of QuerySeriesIterator interface
type MockQuerySeriesIterator struct {
	ctrl     *gomock.Controller
	recorder *MockQuerySeriesIteratorMockRecorder
}

// MockQuerySeriesIteratorMockRecorder is the mock recorder for MockQuerySeriesIterator
type MockQuerySeriesIteratorMockRecorder struct {
	mock *MockQuerySeriesIterator
}

// NewMockQuerySeriesIterator creates a new mock instance
func NewMockQuerySeriesIterator(ctrl *gomock.Controller) *MockQuerySeriesIterator {
	mock := &MockQuerySeriesIterator{ctrl: ctrl}
	mock.recorder = &MockQuerySeriesIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerySeriesIterator) EXPECT() *MockQuerySeriesIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockQuerySeriesIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockQuerySeriesIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQuerySeriesIterator)(nil).Close))
}

// Current mocks base method
func (m *MockQuerySeriesIterator) Current() (ts.Datapoint, time0.Unit, ts.Annotation) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(ts.Datapoint)
	ret1, _ := ret[1].(time0.Unit)
	ret2, _ := ret[2].(ts.Annotation)
	return ret0, ret1, ret2
}

// Current indicates an expected call of Current
func (mr *MockQuerySeriesIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockQuerySeriesIterator)(nil).Current))
}

// Err mocks base method
func (m *MockQuerySeriesIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockQuerySeriesIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockQuerySeriesIterator)(nil).Err))
}

// Next mocks base method
func (m *MockQuerySeriesIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockQuerySeriesIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockQuerySeriesIterator)(nil).Next))
}

// SeriesMetadata mocks base method
func (m *MockQuerySeriesIterator) SeriesMetadata() SeriesMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMetadata")
	ret0, _ := ret[0].(SeriesMetadata)
	return ret0
}

// SeriesMetadata indicates an expected call of SeriesMetadata
func (mr *MockQuerySeriesIteratorMockRecorder) SeriesMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMetadata", reflect.TypeOf((*MockQuerySeriesIterator)(nil).SeriesMetadata))
}

// MockBorrowedBuffer is a mock of BorrowedBuffer interface
type MockBorrowedBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockBorrowedBufferMockRecorder
}

// MockBorrowedBufferMockRecorder is the mock recorder for MockBorrowedBuffer
type MockBorrowedBufferMockRecorder struct {
	mock *MockBorrowedBuffer
}

// NewMockBorrowedBuffer creates a new mock instance
func NewMockBorrowedBuffer(ctrl *gomock.Controller) *MockBorrowedBuffer {
	mock := &MockBorrowedBuffer{ctrl: ctrl}
	mock.recorder = &MockBorrowedBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBorrowedBuffer) EXPECT() *MockBorrowedBufferMockRecorder {
	return m.recorder
}

// Finalize mocks base method
func (m *MockBorrowedBuffer) Finalize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockBorrowedBufferMockRecorder) Finalize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockBorrowedBuffer)(nil).Finalize))
}

// MockFixedBufferManager is a mock of FixedBufferManager interface
type MockFixedBufferManager struct {
	ctrl     *gomock.Controller
	recorder *MockFixedBufferManagerMockRecorder
}

// MockFixedBufferManagerMockRecorder is the mock recorder for MockFixedBufferManager
type MockFixedBufferManagerMockRecorder struct {
	mock *MockFixedBufferManager
}

// NewMockFixedBufferManager creates a new mock instance
func NewMockFixedBufferManager(ctrl *gomock.Controller) *MockFixedBufferManager {
	mock := &MockFixedBufferManager{ctrl: ctrl}
	mock.recorder = &MockFixedBufferManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFixedBufferManager) EXPECT() *MockFixedBufferManagerMockRecorder {
	return m.recorder
}

// Copy mocks base method
func (m *MockFixedBufferManager) Copy(arg0 []byte) ([]byte, BorrowedBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(BorrowedBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Copy indicates an expected call of Copy
func (mr *MockFixedBufferManagerMockRecorder) Copy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockFixedBufferManager)(nil).Copy), arg0)
}
