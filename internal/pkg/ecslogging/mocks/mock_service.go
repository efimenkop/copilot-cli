// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/ecslogging/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	cloudwatchlogs "github.com/aws/copilot-cli/internal/pkg/aws/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MocklogGetter is a mock of logGetter interface
type MocklogGetter struct {
	ctrl     *gomock.Controller
	recorder *MocklogGetterMockRecorder
}

// MocklogGetterMockRecorder is the mock recorder for MocklogGetter
type MocklogGetterMockRecorder struct {
	mock *MocklogGetter
}

// NewMocklogGetter creates a new mock instance
func NewMocklogGetter(ctrl *gomock.Controller) *MocklogGetter {
	mock := &MocklogGetter{ctrl: ctrl}
	mock.recorder = &MocklogGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocklogGetter) EXPECT() *MocklogGetterMockRecorder {
	return m.recorder
}

// LogEvents mocks base method
func (m *MocklogGetter) LogEvents(logGroupName string, streamLastEventTime map[string]int64, opts ...cloudwatchlogs.GetLogEventsOpts) (*cloudwatchlogs.LogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{logGroupName, streamLastEventTime}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LogEvents", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.LogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogEvents indicates an expected call of LogEvents
func (mr *MocklogGetterMockRecorder) LogEvents(logGroupName, streamLastEventTime interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{logGroupName, streamLastEventTime}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogEvents", reflect.TypeOf((*MocklogGetter)(nil).LogEvents), varargs...)
}
