// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/nftables/firewall.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/metal-stack/firewall-controller/api/v1"
)

// MockFQDNCache is a mock of FQDNCache interface.
type MockFQDNCache struct {
	ctrl     *gomock.Controller
	recorder *MockFQDNCacheMockRecorder
}

// MockFQDNCacheMockRecorder is the mock recorder for MockFQDNCache.
type MockFQDNCacheMockRecorder struct {
	mock *MockFQDNCache
}

// NewMockFQDNCache creates a new mock instance.
func NewMockFQDNCache(ctrl *gomock.Controller) *MockFQDNCache {
	mock := &MockFQDNCache{ctrl: ctrl}
	mock.recorder = &MockFQDNCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFQDNCache) EXPECT() *MockFQDNCacheMockRecorder {
	return m.recorder
}

// GetSetsForFQDN mocks base method.
func (m *MockFQDNCache) GetSetsForFQDN(fqdn v1.FQDNSelector, update bool) []v1.IPSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetsForFQDN", fqdn, update)
	ret0, _ := ret[0].([]v1.IPSet)
	return ret0
}

// GetSetsForFQDN indicates an expected call of GetSetsForFQDN.
func (mr *MockFQDNCacheMockRecorder) GetSetsForFQDN(fqdn, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetsForFQDN", reflect.TypeOf((*MockFQDNCache)(nil).GetSetsForFQDN), fqdn, update)
}

// GetSetsForRendering mocks base method.
func (m *MockFQDNCache) GetSetsForRendering(fqdns []v1.FQDNSelector) []v1.IPSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetsForRendering", fqdns)
	ret0, _ := ret[0].([]v1.IPSet)
	return ret0
}

// GetSetsForRendering indicates an expected call of GetSetsForRendering.
func (mr *MockFQDNCacheMockRecorder) GetSetsForRendering(fqdns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetsForRendering", reflect.TypeOf((*MockFQDNCache)(nil).GetSetsForRendering), fqdns)
}
