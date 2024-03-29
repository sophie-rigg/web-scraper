// Code generated by MockGen. DO NOT EDIT.
// Source: processor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	sync "sync"

	gomock "github.com/golang/mock/gomock"
	html "golang.org/x/net/html"
)

// MockProcessor is a mock of Processor interface.
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor.
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance.
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockProcessor) Process(node *html.Node, urlChannel chan string, wg *sync.WaitGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process", node, urlChannel, wg)
}

// Process indicates an expected call of Process.
func (mr *MockProcessorMockRecorder) Process(node, urlChannel, wg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockProcessor)(nil).Process), node, urlChannel, wg)
}
