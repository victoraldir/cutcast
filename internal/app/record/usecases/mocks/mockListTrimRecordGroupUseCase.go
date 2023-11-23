// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/victoraldir/cutcast/internal/app/record/usecases (interfaces: ListTrimRecordGroupUseCase)

// Package usecases is a generated GoMock package.
package usecases

import (
	reflect "reflect"

	usecases "github.com/victoraldir/cutcast/internal/app/record/usecases"
	gomock "go.uber.org/mock/gomock"
)

// MockListTrimRecordGroupUseCase is a mock of ListTrimRecordGroupUseCase interface.
type MockListTrimRecordGroupUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockListTrimRecordGroupUseCaseMockRecorder
}

// MockListTrimRecordGroupUseCaseMockRecorder is the mock recorder for MockListTrimRecordGroupUseCase.
type MockListTrimRecordGroupUseCaseMockRecorder struct {
	mock *MockListTrimRecordGroupUseCase
}

// NewMockListTrimRecordGroupUseCase creates a new mock instance.
func NewMockListTrimRecordGroupUseCase(ctrl *gomock.Controller) *MockListTrimRecordGroupUseCase {
	mock := &MockListTrimRecordGroupUseCase{ctrl: ctrl}
	mock.recorder = &MockListTrimRecordGroupUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListTrimRecordGroupUseCase) EXPECT() *MockListTrimRecordGroupUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockListTrimRecordGroupUseCase) Execute(arg0 string) ([]usecases.TrimRecordGroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].([]usecases.TrimRecordGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockListTrimRecordGroupUseCaseMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockListTrimRecordGroupUseCase)(nil).Execute), arg0)
}