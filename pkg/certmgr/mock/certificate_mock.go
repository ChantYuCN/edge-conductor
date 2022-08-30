//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/intel/edge-conductor/pkg/certmgr (interfaces: CertificateWrapper)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	certmgr "github.com/intel/edge-conductor/pkg/api/certmgr"
)

// MockCertificateWrapper is a mock of CertificateWrapper interface.
type MockCertificateWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateWrapperMockRecorder
}

// MockCertificateWrapperMockRecorder is the mock recorder for MockCertificateWrapper.
type MockCertificateWrapperMockRecorder struct {
	mock *MockCertificateWrapper
}

// NewMockCertificateWrapper creates a new mock instance.
func NewMockCertificateWrapper(ctrl *gomock.Controller) *MockCertificateWrapper {
	mock := &MockCertificateWrapper{ctrl: ctrl}
	mock.recorder = &MockCertificateWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateWrapper) EXPECT() *MockCertificateWrapperMockRecorder {
	return m.recorder
}

// GenCertAndConfig mocks base method.
func (m *MockCertificateWrapper) GenCertAndConfig(arg0 certmgr.Certificate, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenCertAndConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenCertAndConfig indicates an expected call of GenCertAndConfig.
func (mr *MockCertificateWrapperMockRecorder) GenCertAndConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenCertAndConfig", reflect.TypeOf((*MockCertificateWrapper)(nil).GenCertAndConfig), arg0, arg1)
}
