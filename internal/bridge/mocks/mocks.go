// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/v2/internal/bridge (interfaces: TLSReporter,ProxyController,Autostarter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTLSReporter is a mock of TLSReporter interface.
type MockTLSReporter struct {
	ctrl     *gomock.Controller
	recorder *MockTLSReporterMockRecorder
}

// MockTLSReporterMockRecorder is the mock recorder for MockTLSReporter.
type MockTLSReporterMockRecorder struct {
	mock *MockTLSReporter
}

// NewMockTLSReporter creates a new mock instance.
func NewMockTLSReporter(ctrl *gomock.Controller) *MockTLSReporter {
	mock := &MockTLSReporter{ctrl: ctrl}
	mock.recorder = &MockTLSReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTLSReporter) EXPECT() *MockTLSReporterMockRecorder {
	return m.recorder
}

// GetTLSIssueCh mocks base method.
func (m *MockTLSReporter) GetTLSIssueCh() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTLSIssueCh")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// GetTLSIssueCh indicates an expected call of GetTLSIssueCh.
func (mr *MockTLSReporterMockRecorder) GetTLSIssueCh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTLSIssueCh", reflect.TypeOf((*MockTLSReporter)(nil).GetTLSIssueCh))
}

// MockProxyController is a mock of ProxyController interface.
type MockProxyController struct {
	ctrl     *gomock.Controller
	recorder *MockProxyControllerMockRecorder
}

// MockProxyControllerMockRecorder is the mock recorder for MockProxyController.
type MockProxyControllerMockRecorder struct {
	mock *MockProxyController
}

// NewMockProxyController creates a new mock instance.
func NewMockProxyController(ctrl *gomock.Controller) *MockProxyController {
	mock := &MockProxyController{ctrl: ctrl}
	mock.recorder = &MockProxyControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProxyController) EXPECT() *MockProxyControllerMockRecorder {
	return m.recorder
}

// AllowProxy mocks base method.
func (m *MockProxyController) AllowProxy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AllowProxy")
}

// AllowProxy indicates an expected call of AllowProxy.
func (mr *MockProxyControllerMockRecorder) AllowProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowProxy", reflect.TypeOf((*MockProxyController)(nil).AllowProxy))
}

// DisallowProxy mocks base method.
func (m *MockProxyController) DisallowProxy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisallowProxy")
}

// DisallowProxy indicates an expected call of DisallowProxy.
func (mr *MockProxyControllerMockRecorder) DisallowProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisallowProxy", reflect.TypeOf((*MockProxyController)(nil).DisallowProxy))
}

// MockAutostarter is a mock of Autostarter interface.
type MockAutostarter struct {
	ctrl     *gomock.Controller
	recorder *MockAutostarterMockRecorder
}

// MockAutostarterMockRecorder is the mock recorder for MockAutostarter.
type MockAutostarterMockRecorder struct {
	mock *MockAutostarter
}

// NewMockAutostarter creates a new mock instance.
func NewMockAutostarter(ctrl *gomock.Controller) *MockAutostarter {
	mock := &MockAutostarter{ctrl: ctrl}
	mock.recorder = &MockAutostarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutostarter) EXPECT() *MockAutostarterMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockAutostarter) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockAutostarterMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockAutostarter)(nil).Disable))
}

// Enable mocks base method.
func (m *MockAutostarter) Enable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockAutostarterMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockAutostarter)(nil).Enable))
}
