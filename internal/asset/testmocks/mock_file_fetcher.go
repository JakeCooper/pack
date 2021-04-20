// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/pack/internal/asset (interfaces: FileFetcher)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	context "context"
	oci "github.com/buildpacks/pack/internal/oci"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFileFetcher is a mock of FileFetcher interface
type MockFileFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFileFetcherMockRecorder
}

// MockFileFetcherMockRecorder is the mock recorder for MockFileFetcher
type MockFileFetcherMockRecorder struct {
	mock *MockFileFetcher
}

// NewMockFileFetcher creates a new mock instance
func NewMockFileFetcher(ctrl *gomock.Controller) *MockFileFetcher {
	mock := &MockFileFetcher{ctrl: ctrl}
	mock.recorder = &MockFileFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileFetcher) EXPECT() *MockFileFetcherMockRecorder {
	return m.recorder
}

// FetchFileAssets mocks base method
func (m *MockFileFetcher) FetchFileAssets(arg0 context.Context, arg1 string, arg2 ...string) ([]*oci.LayoutPackage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchFileAssets", varargs...)
	ret0, _ := ret[0].([]*oci.LayoutPackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFileAssets indicates an expected call of FetchFileAssets
func (mr *MockFileFetcherMockRecorder) FetchFileAssets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFileAssets", reflect.TypeOf((*MockFileFetcher)(nil).FetchFileAssets), varargs...)
}
