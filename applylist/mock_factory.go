// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/box/kube-applier/applylist

package applylist

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of FactoryInterface interface
type MockFactoryInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockFactoryInterfaceRecorder
}

// Recorder for MockFactoryInterface (not exported)
type _MockFactoryInterfaceRecorder struct {
	mock *MockFactoryInterface
}

func NewMockFactoryInterface(ctrl *gomock.Controller) *MockFactoryInterface {
	mock := &MockFactoryInterface{ctrl: ctrl}
	mock.recorder = &_MockFactoryInterfaceRecorder{mock}
	return mock
}

func (_m *MockFactoryInterface) EXPECT() *_MockFactoryInterfaceRecorder {
	return _m.recorder
}

func (_m *MockFactoryInterface) Create() ([]string, []string, []string, error) {
	ret := _m.ctrl.Call(_m, "Create")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

func (_mr *_MockFactoryInterfaceRecorder) Create() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create")
}
