// Code generated by MockGen. DO NOT EDIT.
// Source: ../../vendor/k8s.io/client-go/discovery/discovery_client.go

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	openapiv2 "github.com/googleapis/gnostic/openapiv2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	version "k8s.io/apimachinery/pkg/version"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockDiscoveryInterface is a mock of DiscoveryInterface interface
type MockDiscoveryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryInterfaceMockRecorder
}

// MockDiscoveryInterfaceMockRecorder is the mock recorder for MockDiscoveryInterface
type MockDiscoveryInterfaceMockRecorder struct {
	mock *MockDiscoveryInterface
}

// NewMockDiscoveryInterface creates a new mock instance
func NewMockDiscoveryInterface(ctrl *gomock.Controller) *MockDiscoveryInterface {
	mock := &MockDiscoveryInterface{ctrl: ctrl}
	mock.recorder = &MockDiscoveryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoveryInterface) EXPECT() *MockDiscoveryInterfaceMockRecorder {
	return m.recorder
}

// RESTClient mocks base method
func (m *MockDiscoveryInterface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockDiscoveryInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockDiscoveryInterface)(nil).RESTClient))
}

// ServerGroups mocks base method
func (m *MockDiscoveryInterface) ServerGroups() (*v1.APIGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroups")
	ret0, _ := ret[0].(*v1.APIGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerGroups indicates an expected call of ServerGroups
func (mr *MockDiscoveryInterfaceMockRecorder) ServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroups", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerGroups))
}

// ServerResourcesForGroupVersion mocks base method
func (m *MockDiscoveryInterface) ServerResourcesForGroupVersion(groupVersion string) (*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResourcesForGroupVersion", groupVersion)
	ret0, _ := ret[0].(*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResourcesForGroupVersion indicates an expected call of ServerResourcesForGroupVersion
func (mr *MockDiscoveryInterfaceMockRecorder) ServerResourcesForGroupVersion(groupVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResourcesForGroupVersion", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerResourcesForGroupVersion), groupVersion)
}

// ServerResources mocks base method
func (m *MockDiscoveryInterface) ServerResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResources indicates an expected call of ServerResources
func (mr *MockDiscoveryInterfaceMockRecorder) ServerResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResources", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerResources))
}

// ServerGroupsAndResources mocks base method
func (m *MockDiscoveryInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroupsAndResources")
	ret0, _ := ret[0].([]*v1.APIGroup)
	ret1, _ := ret[1].([]*v1.APIResourceList)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerGroupsAndResources indicates an expected call of ServerGroupsAndResources
func (mr *MockDiscoveryInterfaceMockRecorder) ServerGroupsAndResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroupsAndResources", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerGroupsAndResources))
}

// ServerPreferredResources mocks base method
func (m *MockDiscoveryInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredResources indicates an expected call of ServerPreferredResources
func (mr *MockDiscoveryInterfaceMockRecorder) ServerPreferredResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredResources", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerPreferredResources))
}

// ServerPreferredNamespacedResources mocks base method
func (m *MockDiscoveryInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredNamespacedResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredNamespacedResources indicates an expected call of ServerPreferredNamespacedResources
func (mr *MockDiscoveryInterfaceMockRecorder) ServerPreferredNamespacedResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredNamespacedResources", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerPreferredNamespacedResources))
}

// ServerVersion mocks base method
func (m *MockDiscoveryInterface) ServerVersion() (*version.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(*version.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerVersion indicates an expected call of ServerVersion
func (mr *MockDiscoveryInterfaceMockRecorder) ServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerVersion))
}

// OpenAPISchema mocks base method
func (m *MockDiscoveryInterface) OpenAPISchema() (*openapiv2.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAPISchema")
	ret0, _ := ret[0].(*openapiv2.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenAPISchema indicates an expected call of OpenAPISchema
func (mr *MockDiscoveryInterfaceMockRecorder) OpenAPISchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAPISchema", reflect.TypeOf((*MockDiscoveryInterface)(nil).OpenAPISchema))
}

// MockCachedDiscoveryInterface is a mock of CachedDiscoveryInterface interface
type MockCachedDiscoveryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCachedDiscoveryInterfaceMockRecorder
}

// MockCachedDiscoveryInterfaceMockRecorder is the mock recorder for MockCachedDiscoveryInterface
type MockCachedDiscoveryInterfaceMockRecorder struct {
	mock *MockCachedDiscoveryInterface
}

// NewMockCachedDiscoveryInterface creates a new mock instance
func NewMockCachedDiscoveryInterface(ctrl *gomock.Controller) *MockCachedDiscoveryInterface {
	mock := &MockCachedDiscoveryInterface{ctrl: ctrl}
	mock.recorder = &MockCachedDiscoveryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCachedDiscoveryInterface) EXPECT() *MockCachedDiscoveryInterfaceMockRecorder {
	return m.recorder
}

// RESTClient mocks base method
func (m *MockCachedDiscoveryInterface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockCachedDiscoveryInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).RESTClient))
}

// ServerGroups mocks base method
func (m *MockCachedDiscoveryInterface) ServerGroups() (*v1.APIGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroups")
	ret0, _ := ret[0].(*v1.APIGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerGroups indicates an expected call of ServerGroups
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroups", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerGroups))
}

// ServerResourcesForGroupVersion mocks base method
func (m *MockCachedDiscoveryInterface) ServerResourcesForGroupVersion(groupVersion string) (*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResourcesForGroupVersion", groupVersion)
	ret0, _ := ret[0].(*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResourcesForGroupVersion indicates an expected call of ServerResourcesForGroupVersion
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerResourcesForGroupVersion(groupVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResourcesForGroupVersion", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerResourcesForGroupVersion), groupVersion)
}

// ServerResources mocks base method
func (m *MockCachedDiscoveryInterface) ServerResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResources indicates an expected call of ServerResources
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResources", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerResources))
}

// ServerGroupsAndResources mocks base method
func (m *MockCachedDiscoveryInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroupsAndResources")
	ret0, _ := ret[0].([]*v1.APIGroup)
	ret1, _ := ret[1].([]*v1.APIResourceList)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerGroupsAndResources indicates an expected call of ServerGroupsAndResources
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerGroupsAndResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroupsAndResources", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerGroupsAndResources))
}

// ServerPreferredResources mocks base method
func (m *MockCachedDiscoveryInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredResources indicates an expected call of ServerPreferredResources
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerPreferredResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredResources", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerPreferredResources))
}

// ServerPreferredNamespacedResources mocks base method
func (m *MockCachedDiscoveryInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredNamespacedResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredNamespacedResources indicates an expected call of ServerPreferredNamespacedResources
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerPreferredNamespacedResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredNamespacedResources", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerPreferredNamespacedResources))
}

// ServerVersion mocks base method
func (m *MockCachedDiscoveryInterface) ServerVersion() (*version.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(*version.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerVersion indicates an expected call of ServerVersion
func (mr *MockCachedDiscoveryInterfaceMockRecorder) ServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).ServerVersion))
}

// OpenAPISchema mocks base method
func (m *MockCachedDiscoveryInterface) OpenAPISchema() (*openapiv2.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAPISchema")
	ret0, _ := ret[0].(*openapiv2.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenAPISchema indicates an expected call of OpenAPISchema
func (mr *MockCachedDiscoveryInterfaceMockRecorder) OpenAPISchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAPISchema", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).OpenAPISchema))
}

// Fresh mocks base method
func (m *MockCachedDiscoveryInterface) Fresh() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fresh")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Fresh indicates an expected call of Fresh
func (mr *MockCachedDiscoveryInterfaceMockRecorder) Fresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fresh", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).Fresh))
}

// Invalidate mocks base method
func (m *MockCachedDiscoveryInterface) Invalidate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Invalidate")
}

// Invalidate indicates an expected call of Invalidate
func (mr *MockCachedDiscoveryInterfaceMockRecorder) Invalidate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockCachedDiscoveryInterface)(nil).Invalidate))
}

// MockServerGroupsInterface is a mock of ServerGroupsInterface interface
type MockServerGroupsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServerGroupsInterfaceMockRecorder
}

// MockServerGroupsInterfaceMockRecorder is the mock recorder for MockServerGroupsInterface
type MockServerGroupsInterfaceMockRecorder struct {
	mock *MockServerGroupsInterface
}

// NewMockServerGroupsInterface creates a new mock instance
func NewMockServerGroupsInterface(ctrl *gomock.Controller) *MockServerGroupsInterface {
	mock := &MockServerGroupsInterface{ctrl: ctrl}
	mock.recorder = &MockServerGroupsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerGroupsInterface) EXPECT() *MockServerGroupsInterfaceMockRecorder {
	return m.recorder
}

// ServerGroups mocks base method
func (m *MockServerGroupsInterface) ServerGroups() (*v1.APIGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroups")
	ret0, _ := ret[0].(*v1.APIGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerGroups indicates an expected call of ServerGroups
func (mr *MockServerGroupsInterfaceMockRecorder) ServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroups", reflect.TypeOf((*MockServerGroupsInterface)(nil).ServerGroups))
}

// MockServerResourcesInterface is a mock of ServerResourcesInterface interface
type MockServerResourcesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServerResourcesInterfaceMockRecorder
}

// MockServerResourcesInterfaceMockRecorder is the mock recorder for MockServerResourcesInterface
type MockServerResourcesInterfaceMockRecorder struct {
	mock *MockServerResourcesInterface
}

// NewMockServerResourcesInterface creates a new mock instance
func NewMockServerResourcesInterface(ctrl *gomock.Controller) *MockServerResourcesInterface {
	mock := &MockServerResourcesInterface{ctrl: ctrl}
	mock.recorder = &MockServerResourcesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerResourcesInterface) EXPECT() *MockServerResourcesInterfaceMockRecorder {
	return m.recorder
}

// ServerResourcesForGroupVersion mocks base method
func (m *MockServerResourcesInterface) ServerResourcesForGroupVersion(groupVersion string) (*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResourcesForGroupVersion", groupVersion)
	ret0, _ := ret[0].(*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResourcesForGroupVersion indicates an expected call of ServerResourcesForGroupVersion
func (mr *MockServerResourcesInterfaceMockRecorder) ServerResourcesForGroupVersion(groupVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResourcesForGroupVersion", reflect.TypeOf((*MockServerResourcesInterface)(nil).ServerResourcesForGroupVersion), groupVersion)
}

// ServerResources mocks base method
func (m *MockServerResourcesInterface) ServerResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerResources indicates an expected call of ServerResources
func (mr *MockServerResourcesInterfaceMockRecorder) ServerResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerResources", reflect.TypeOf((*MockServerResourcesInterface)(nil).ServerResources))
}

// ServerGroupsAndResources mocks base method
func (m *MockServerResourcesInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGroupsAndResources")
	ret0, _ := ret[0].([]*v1.APIGroup)
	ret1, _ := ret[1].([]*v1.APIResourceList)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerGroupsAndResources indicates an expected call of ServerGroupsAndResources
func (mr *MockServerResourcesInterfaceMockRecorder) ServerGroupsAndResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGroupsAndResources", reflect.TypeOf((*MockServerResourcesInterface)(nil).ServerGroupsAndResources))
}

// ServerPreferredResources mocks base method
func (m *MockServerResourcesInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredResources indicates an expected call of ServerPreferredResources
func (mr *MockServerResourcesInterfaceMockRecorder) ServerPreferredResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredResources", reflect.TypeOf((*MockServerResourcesInterface)(nil).ServerPreferredResources))
}

// ServerPreferredNamespacedResources mocks base method
func (m *MockServerResourcesInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPreferredNamespacedResources")
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPreferredNamespacedResources indicates an expected call of ServerPreferredNamespacedResources
func (mr *MockServerResourcesInterfaceMockRecorder) ServerPreferredNamespacedResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPreferredNamespacedResources", reflect.TypeOf((*MockServerResourcesInterface)(nil).ServerPreferredNamespacedResources))
}

// MockServerVersionInterface is a mock of ServerVersionInterface interface
type MockServerVersionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServerVersionInterfaceMockRecorder
}

// MockServerVersionInterfaceMockRecorder is the mock recorder for MockServerVersionInterface
type MockServerVersionInterfaceMockRecorder struct {
	mock *MockServerVersionInterface
}

// NewMockServerVersionInterface creates a new mock instance
func NewMockServerVersionInterface(ctrl *gomock.Controller) *MockServerVersionInterface {
	mock := &MockServerVersionInterface{ctrl: ctrl}
	mock.recorder = &MockServerVersionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerVersionInterface) EXPECT() *MockServerVersionInterfaceMockRecorder {
	return m.recorder
}

// ServerVersion mocks base method
func (m *MockServerVersionInterface) ServerVersion() (*version.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(*version.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerVersion indicates an expected call of ServerVersion
func (mr *MockServerVersionInterfaceMockRecorder) ServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockServerVersionInterface)(nil).ServerVersion))
}

// MockOpenAPISchemaInterface is a mock of OpenAPISchemaInterface interface
type MockOpenAPISchemaInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOpenAPISchemaInterfaceMockRecorder
}

// MockOpenAPISchemaInterfaceMockRecorder is the mock recorder for MockOpenAPISchemaInterface
type MockOpenAPISchemaInterfaceMockRecorder struct {
	mock *MockOpenAPISchemaInterface
}

// NewMockOpenAPISchemaInterface creates a new mock instance
func NewMockOpenAPISchemaInterface(ctrl *gomock.Controller) *MockOpenAPISchemaInterface {
	mock := &MockOpenAPISchemaInterface{ctrl: ctrl}
	mock.recorder = &MockOpenAPISchemaInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenAPISchemaInterface) EXPECT() *MockOpenAPISchemaInterfaceMockRecorder {
	return m.recorder
}

// OpenAPISchema mocks base method
func (m *MockOpenAPISchemaInterface) OpenAPISchema() (*openapiv2.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAPISchema")
	ret0, _ := ret[0].(*openapiv2.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenAPISchema indicates an expected call of OpenAPISchema
func (mr *MockOpenAPISchemaInterfaceMockRecorder) OpenAPISchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAPISchema", reflect.TypeOf((*MockOpenAPISchemaInterface)(nil).OpenAPISchema))
}
