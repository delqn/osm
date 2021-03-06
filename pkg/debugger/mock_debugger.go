// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/debugger (interfaces: CertificateManagerDebugger,MeshCatalogDebugger,XDSDebugger,DebugServer)

// Package debugger is a generated GoMock package.
package debugger

import (
	http "net/http"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	v1alpha20 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"

	certificate "github.com/openservicemesh/osm/pkg/certificate"
	envoy "github.com/openservicemesh/osm/pkg/envoy"
	service "github.com/openservicemesh/osm/pkg/service"
)

// MockCertificateManagerDebugger is a mock of CertificateManagerDebugger interface
type MockCertificateManagerDebugger struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateManagerDebuggerMockRecorder
}

// MockCertificateManagerDebuggerMockRecorder is the mock recorder for MockCertificateManagerDebugger
type MockCertificateManagerDebuggerMockRecorder struct {
	mock *MockCertificateManagerDebugger
}

// NewMockCertificateManagerDebugger creates a new mock instance
func NewMockCertificateManagerDebugger(ctrl *gomock.Controller) *MockCertificateManagerDebugger {
	mock := &MockCertificateManagerDebugger{ctrl: ctrl}
	mock.recorder = &MockCertificateManagerDebuggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateManagerDebugger) EXPECT() *MockCertificateManagerDebuggerMockRecorder {
	return m.recorder
}

// ListIssuedCertificates mocks base method
func (m *MockCertificateManagerDebugger) ListIssuedCertificates() []certificate.Certificater {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIssuedCertificates")
	ret0, _ := ret[0].([]certificate.Certificater)
	return ret0
}

// ListIssuedCertificates indicates an expected call of ListIssuedCertificates
func (mr *MockCertificateManagerDebuggerMockRecorder) ListIssuedCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssuedCertificates", reflect.TypeOf((*MockCertificateManagerDebugger)(nil).ListIssuedCertificates))
}

// MockMeshCatalogDebugger is a mock of MeshCatalogDebugger interface
type MockMeshCatalogDebugger struct {
	ctrl     *gomock.Controller
	recorder *MockMeshCatalogDebuggerMockRecorder
}

// MockMeshCatalogDebuggerMockRecorder is the mock recorder for MockMeshCatalogDebugger
type MockMeshCatalogDebuggerMockRecorder struct {
	mock *MockMeshCatalogDebugger
}

// NewMockMeshCatalogDebugger creates a new mock instance
func NewMockMeshCatalogDebugger(ctrl *gomock.Controller) *MockMeshCatalogDebugger {
	mock := &MockMeshCatalogDebugger{ctrl: ctrl}
	mock.recorder = &MockMeshCatalogDebuggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshCatalogDebugger) EXPECT() *MockMeshCatalogDebuggerMockRecorder {
	return m.recorder
}

// ListConnectedProxies mocks base method
func (m *MockMeshCatalogDebugger) ListConnectedProxies() map[certificate.CommonName]*envoy.Proxy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnectedProxies")
	ret0, _ := ret[0].(map[certificate.CommonName]*envoy.Proxy)
	return ret0
}

// ListConnectedProxies indicates an expected call of ListConnectedProxies
func (mr *MockMeshCatalogDebuggerMockRecorder) ListConnectedProxies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectedProxies", reflect.TypeOf((*MockMeshCatalogDebugger)(nil).ListConnectedProxies))
}

// ListDisconnectedProxies mocks base method
func (m *MockMeshCatalogDebugger) ListDisconnectedProxies() map[certificate.CommonName]time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDisconnectedProxies")
	ret0, _ := ret[0].(map[certificate.CommonName]time.Time)
	return ret0
}

// ListDisconnectedProxies indicates an expected call of ListDisconnectedProxies
func (mr *MockMeshCatalogDebuggerMockRecorder) ListDisconnectedProxies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDisconnectedProxies", reflect.TypeOf((*MockMeshCatalogDebugger)(nil).ListDisconnectedProxies))
}

// ListExpectedProxies mocks base method
func (m *MockMeshCatalogDebugger) ListExpectedProxies() map[certificate.CommonName]time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExpectedProxies")
	ret0, _ := ret[0].(map[certificate.CommonName]time.Time)
	return ret0
}

// ListExpectedProxies indicates an expected call of ListExpectedProxies
func (mr *MockMeshCatalogDebuggerMockRecorder) ListExpectedProxies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExpectedProxies", reflect.TypeOf((*MockMeshCatalogDebugger)(nil).ListExpectedProxies))
}

// ListMonitoredNamespaces mocks base method
func (m *MockMeshCatalogDebugger) ListMonitoredNamespaces() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMonitoredNamespaces")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListMonitoredNamespaces indicates an expected call of ListMonitoredNamespaces
func (mr *MockMeshCatalogDebuggerMockRecorder) ListMonitoredNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitoredNamespaces", reflect.TypeOf((*MockMeshCatalogDebugger)(nil).ListMonitoredNamespaces))
}

// ListSMIPolicies mocks base method
func (m *MockMeshCatalogDebugger) ListSMIPolicies() ([]*v1alpha20.TrafficSplit, []service.WeightedService, []service.K8sServiceAccount, []*v1alpha3.HTTPRouteGroup, []*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSMIPolicies")
	ret0, _ := ret[0].([]*v1alpha20.TrafficSplit)
	ret1, _ := ret[1].([]service.WeightedService)
	ret2, _ := ret[2].([]service.K8sServiceAccount)
	ret3, _ := ret[3].([]*v1alpha3.HTTPRouteGroup)
	ret4, _ := ret[4].([]*v1alpha2.TrafficTarget)
	return ret0, ret1, ret2, ret3, ret4
}

// ListSMIPolicies indicates an expected call of ListSMIPolicies
func (mr *MockMeshCatalogDebuggerMockRecorder) ListSMIPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSMIPolicies", reflect.TypeOf((*MockMeshCatalogDebugger)(nil).ListSMIPolicies))
}

// MockXDSDebugger is a mock of XDSDebugger interface
type MockXDSDebugger struct {
	ctrl     *gomock.Controller
	recorder *MockXDSDebuggerMockRecorder
}

// MockXDSDebuggerMockRecorder is the mock recorder for MockXDSDebugger
type MockXDSDebuggerMockRecorder struct {
	mock *MockXDSDebugger
}

// NewMockXDSDebugger creates a new mock instance
func NewMockXDSDebugger(ctrl *gomock.Controller) *MockXDSDebugger {
	mock := &MockXDSDebugger{ctrl: ctrl}
	mock.recorder = &MockXDSDebuggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXDSDebugger) EXPECT() *MockXDSDebuggerMockRecorder {
	return m.recorder
}

// GetXDSLog mocks base method
func (m *MockXDSDebugger) GetXDSLog() *map[certificate.CommonName]map[envoy.TypeURI][]time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetXDSLog")
	ret0, _ := ret[0].(*map[certificate.CommonName]map[envoy.TypeURI][]time.Time)
	return ret0
}

// GetXDSLog indicates an expected call of GetXDSLog
func (mr *MockXDSDebuggerMockRecorder) GetXDSLog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXDSLog", reflect.TypeOf((*MockXDSDebugger)(nil).GetXDSLog))
}

// MockDebugServer is a mock of DebugServer interface
type MockDebugServer struct {
	ctrl     *gomock.Controller
	recorder *MockDebugServerMockRecorder
}

// MockDebugServerMockRecorder is the mock recorder for MockDebugServer
type MockDebugServerMockRecorder struct {
	mock *MockDebugServer
}

// NewMockDebugServer creates a new mock instance
func NewMockDebugServer(ctrl *gomock.Controller) *MockDebugServer {
	mock := &MockDebugServer{ctrl: ctrl}
	mock.recorder = &MockDebugServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDebugServer) EXPECT() *MockDebugServerMockRecorder {
	return m.recorder
}

// GetHandlers mocks base method
func (m *MockDebugServer) GetHandlers() map[string]http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandlers")
	ret0, _ := ret[0].(map[string]http.Handler)
	return ret0
}

// GetHandlers indicates an expected call of GetHandlers
func (mr *MockDebugServerMockRecorder) GetHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandlers", reflect.TypeOf((*MockDebugServer)(nil).GetHandlers))
}
