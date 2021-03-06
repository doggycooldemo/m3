// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/x/pool (interfaces: CheckedBytesPool,BytesPool)

// Copyright (c) 2019 Uber Technologies, Inc.
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

// Package pool is a generated GoMock package.
package pool

import (
	"reflect"

	"github.com/m3db/m3/src/x/checked"

	"github.com/golang/mock/gomock"
)

// MockCheckedBytesPool is a mock of CheckedBytesPool interface
type MockCheckedBytesPool struct {
	ctrl     *gomock.Controller
	recorder *MockCheckedBytesPoolMockRecorder
}

// MockCheckedBytesPoolMockRecorder is the mock recorder for MockCheckedBytesPool
type MockCheckedBytesPoolMockRecorder struct {
	mock *MockCheckedBytesPool
}

// NewMockCheckedBytesPool creates a new mock instance
func NewMockCheckedBytesPool(ctrl *gomock.Controller) *MockCheckedBytesPool {
	mock := &MockCheckedBytesPool{ctrl: ctrl}
	mock.recorder = &MockCheckedBytesPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckedBytesPool) EXPECT() *MockCheckedBytesPoolMockRecorder {
	return m.recorder
}

// BytesPool mocks base method
func (m *MockCheckedBytesPool) BytesPool() BytesPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BytesPool")
	ret0, _ := ret[0].(BytesPool)
	return ret0
}

// BytesPool indicates an expected call of BytesPool
func (mr *MockCheckedBytesPoolMockRecorder) BytesPool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BytesPool", reflect.TypeOf((*MockCheckedBytesPool)(nil).BytesPool))
}

// Get mocks base method
func (m *MockCheckedBytesPool) Get(arg0 int) checked.Bytes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(checked.Bytes)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCheckedBytesPoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCheckedBytesPool)(nil).Get), arg0)
}

// Init mocks base method
func (m *MockCheckedBytesPool) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockCheckedBytesPoolMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockCheckedBytesPool)(nil).Init))
}

// MockBytesPool is a mock of BytesPool interface
type MockBytesPool struct {
	ctrl     *gomock.Controller
	recorder *MockBytesPoolMockRecorder
}

// MockBytesPoolMockRecorder is the mock recorder for MockBytesPool
type MockBytesPoolMockRecorder struct {
	mock *MockBytesPool
}

// NewMockBytesPool creates a new mock instance
func NewMockBytesPool(ctrl *gomock.Controller) *MockBytesPool {
	mock := &MockBytesPool{ctrl: ctrl}
	mock.recorder = &MockBytesPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBytesPool) EXPECT() *MockBytesPoolMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBytesPool) Get(arg0 int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockBytesPoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBytesPool)(nil).Get), arg0)
}

// Init mocks base method
func (m *MockBytesPool) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockBytesPoolMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockBytesPool)(nil).Init))
}

// Put mocks base method
func (m *MockBytesPool) Put(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0)
}

// Put indicates an expected call of Put
func (mr *MockBytesPoolMockRecorder) Put(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockBytesPool)(nil).Put), arg0)
}
