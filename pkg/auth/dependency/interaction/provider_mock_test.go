// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go

// Package interaction_test is a generated GoMock package.
package interaction_test

import (
	gomock "github.com/golang/mock/gomock"
	oob "github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/oob"
	identity "github.com/skygeario/skygear-server/pkg/auth/dependency/identity"
	interaction "github.com/skygeario/skygear-server/pkg/auth/dependency/interaction"
	model "github.com/skygeario/skygear-server/pkg/auth/model"
	authn "github.com/skygeario/skygear-server/pkg/core/authn"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStore) Create(i *interaction.Interaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), i)
}

// Get mocks base method
func (m *MockStore) Get(token string) (*interaction.Interaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", token)
	ret0, _ := ret[0].(*interaction.Interaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), token)
}

// Update mocks base method
func (m *MockStore) Update(i *interaction.Interaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), i)
}

// Delete mocks base method
func (m *MockStore) Delete(i *interaction.Interaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), i)
}

// MockIdentityProvider is a mock of IdentityProvider interface
type MockIdentityProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityProviderMockRecorder
}

// MockIdentityProviderMockRecorder is the mock recorder for MockIdentityProvider
type MockIdentityProviderMockRecorder struct {
	mock *MockIdentityProvider
}

// NewMockIdentityProvider creates a new mock instance
func NewMockIdentityProvider(ctrl *gomock.Controller) *MockIdentityProvider {
	mock := &MockIdentityProvider{ctrl: ctrl}
	mock.recorder = &MockIdentityProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentityProvider) EXPECT() *MockIdentityProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockIdentityProvider) Get(userID string, typ authn.IdentityType, id string) (*identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID, typ, id)
	ret0, _ := ret[0].(*identity.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIdentityProviderMockRecorder) Get(userID, typ, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIdentityProvider)(nil).Get), userID, typ, id)
}

// GetByClaims mocks base method
func (m *MockIdentityProvider) GetByClaims(typ authn.IdentityType, claims map[string]interface{}) (string, *identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByClaims", typ, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*identity.Info)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByClaims indicates an expected call of GetByClaims
func (mr *MockIdentityProviderMockRecorder) GetByClaims(typ, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByClaims", reflect.TypeOf((*MockIdentityProvider)(nil).GetByClaims), typ, claims)
}

// GetByUserAndClaims mocks base method
func (m *MockIdentityProvider) GetByUserAndClaims(typ authn.IdentityType, userID string, claims map[string]interface{}) (*identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserAndClaims", typ, userID, claims)
	ret0, _ := ret[0].(*identity.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserAndClaims indicates an expected call of GetByUserAndClaims
func (mr *MockIdentityProviderMockRecorder) GetByUserAndClaims(typ, userID, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserAndClaims", reflect.TypeOf((*MockIdentityProvider)(nil).GetByUserAndClaims), typ, userID, claims)
}

// ListByClaims mocks base method
func (m *MockIdentityProvider) ListByClaims(claims map[string]string) ([]*identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByClaims", claims)
	ret0, _ := ret[0].([]*identity.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByClaims indicates an expected call of ListByClaims
func (mr *MockIdentityProviderMockRecorder) ListByClaims(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByClaims", reflect.TypeOf((*MockIdentityProvider)(nil).ListByClaims), claims)
}

// ListByUser mocks base method
func (m *MockIdentityProvider) ListByUser(userID string) ([]*identity.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUser", userID)
	ret0, _ := ret[0].([]*identity.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUser indicates an expected call of ListByUser
func (mr *MockIdentityProviderMockRecorder) ListByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUser", reflect.TypeOf((*MockIdentityProvider)(nil).ListByUser), userID)
}

// New mocks base method
func (m *MockIdentityProvider) New(userID string, typ authn.IdentityType, claims map[string]interface{}) *identity.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", userID, typ, claims)
	ret0, _ := ret[0].(*identity.Info)
	return ret0
}

// New indicates an expected call of New
func (mr *MockIdentityProviderMockRecorder) New(userID, typ, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIdentityProvider)(nil).New), userID, typ, claims)
}

// WithClaims mocks base method
func (m *MockIdentityProvider) WithClaims(userID string, ii *identity.Info, claims map[string]interface{}) *identity.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithClaims", userID, ii, claims)
	ret0, _ := ret[0].(*identity.Info)
	return ret0
}

// WithClaims indicates an expected call of WithClaims
func (mr *MockIdentityProviderMockRecorder) WithClaims(userID, ii, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithClaims", reflect.TypeOf((*MockIdentityProvider)(nil).WithClaims), userID, ii, claims)
}

// CreateAll mocks base method
func (m *MockIdentityProvider) CreateAll(userID string, is []*identity.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAll", userID, is)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAll indicates an expected call of CreateAll
func (mr *MockIdentityProviderMockRecorder) CreateAll(userID, is interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAll", reflect.TypeOf((*MockIdentityProvider)(nil).CreateAll), userID, is)
}

// UpdateAll mocks base method
func (m *MockIdentityProvider) UpdateAll(userID string, is []*identity.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", userID, is)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockIdentityProviderMockRecorder) UpdateAll(userID, is interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockIdentityProvider)(nil).UpdateAll), userID, is)
}

// DeleteAll mocks base method
func (m *MockIdentityProvider) DeleteAll(userID string, is []*identity.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", userID, is)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockIdentityProviderMockRecorder) DeleteAll(userID, is interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockIdentityProvider)(nil).DeleteAll), userID, is)
}

// Validate mocks base method
func (m *MockIdentityProvider) Validate(is []*identity.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", is)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockIdentityProviderMockRecorder) Validate(is interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockIdentityProvider)(nil).Validate), is)
}

// RelateIdentityToAuthenticator mocks base method
func (m *MockIdentityProvider) RelateIdentityToAuthenticator(identitySpec identity.Spec, authenticatorSpec *interaction.AuthenticatorSpec) *interaction.AuthenticatorSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelateIdentityToAuthenticator", identitySpec, authenticatorSpec)
	ret0, _ := ret[0].(*interaction.AuthenticatorSpec)
	return ret0
}

// RelateIdentityToAuthenticator indicates an expected call of RelateIdentityToAuthenticator
func (mr *MockIdentityProviderMockRecorder) RelateIdentityToAuthenticator(identitySpec, authenticatorSpec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelateIdentityToAuthenticator", reflect.TypeOf((*MockIdentityProvider)(nil).RelateIdentityToAuthenticator), identitySpec, authenticatorSpec)
}

// MockAuthenticatorProvider is a mock of AuthenticatorProvider interface
type MockAuthenticatorProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorProviderMockRecorder
}

// MockAuthenticatorProviderMockRecorder is the mock recorder for MockAuthenticatorProvider
type MockAuthenticatorProviderMockRecorder struct {
	mock *MockAuthenticatorProvider
}

// NewMockAuthenticatorProvider creates a new mock instance
func NewMockAuthenticatorProvider(ctrl *gomock.Controller) *MockAuthenticatorProvider {
	mock := &MockAuthenticatorProvider{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticatorProvider) EXPECT() *MockAuthenticatorProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAuthenticatorProvider) Get(userID string, typ authn.AuthenticatorType, id string) (*interaction.AuthenticatorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID, typ, id)
	ret0, _ := ret[0].(*interaction.AuthenticatorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAuthenticatorProviderMockRecorder) Get(userID, typ, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAuthenticatorProvider)(nil).Get), userID, typ, id)
}

// List mocks base method
func (m *MockAuthenticatorProvider) List(userID string, typ authn.AuthenticatorType) ([]*interaction.AuthenticatorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userID, typ)
	ret0, _ := ret[0].([]*interaction.AuthenticatorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAuthenticatorProviderMockRecorder) List(userID, typ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthenticatorProvider)(nil).List), userID, typ)
}

// ListByIdentity mocks base method
func (m *MockAuthenticatorProvider) ListByIdentity(userID string, ii *identity.Info) ([]*interaction.AuthenticatorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIdentity", userID, ii)
	ret0, _ := ret[0].([]*interaction.AuthenticatorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIdentity indicates an expected call of ListByIdentity
func (mr *MockAuthenticatorProviderMockRecorder) ListByIdentity(userID, ii interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIdentity", reflect.TypeOf((*MockAuthenticatorProvider)(nil).ListByIdentity), userID, ii)
}

// New mocks base method
func (m *MockAuthenticatorProvider) New(userID string, spec interaction.AuthenticatorSpec, secret string) ([]*interaction.AuthenticatorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", userID, spec, secret)
	ret0, _ := ret[0].([]*interaction.AuthenticatorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockAuthenticatorProviderMockRecorder) New(userID, spec, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockAuthenticatorProvider)(nil).New), userID, spec, secret)
}

// CreateAll mocks base method
func (m *MockAuthenticatorProvider) CreateAll(userID string, ais []*interaction.AuthenticatorInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAll", userID, ais)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAll indicates an expected call of CreateAll
func (mr *MockAuthenticatorProviderMockRecorder) CreateAll(userID, ais interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAll", reflect.TypeOf((*MockAuthenticatorProvider)(nil).CreateAll), userID, ais)
}

// DeleteAll mocks base method
func (m *MockAuthenticatorProvider) DeleteAll(userID string, ais []*interaction.AuthenticatorInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", userID, ais)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockAuthenticatorProviderMockRecorder) DeleteAll(userID, ais interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAuthenticatorProvider)(nil).DeleteAll), userID, ais)
}

// Authenticate mocks base method
func (m *MockAuthenticatorProvider) Authenticate(userID string, spec interaction.AuthenticatorSpec, state *map[string]string, secret string) (*interaction.AuthenticatorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", userID, spec, state, secret)
	ret0, _ := ret[0].(*interaction.AuthenticatorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticatorProviderMockRecorder) Authenticate(userID, spec, state, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticatorProvider)(nil).Authenticate), userID, spec, state, secret)
}

// MockUserProvider is a mock of UserProvider interface
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserProvider) Create(userID string, metadata map[string]interface{}, identities []*identity.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, metadata, identities)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserProviderMockRecorder) Create(userID, metadata, identities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserProvider)(nil).Create), userID, metadata, identities)
}

// Get mocks base method
func (m *MockUserProvider) Get(userID string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserProviderMockRecorder) Get(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserProvider)(nil).Get), userID)
}

// MockOOBProvider is a mock of OOBProvider interface
type MockOOBProvider struct {
	ctrl     *gomock.Controller
	recorder *MockOOBProviderMockRecorder
}

// MockOOBProviderMockRecorder is the mock recorder for MockOOBProvider
type MockOOBProviderMockRecorder struct {
	mock *MockOOBProvider
}

// NewMockOOBProvider creates a new mock instance
func NewMockOOBProvider(ctrl *gomock.Controller) *MockOOBProvider {
	mock := &MockOOBProvider{ctrl: ctrl}
	mock.recorder = &MockOOBProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOOBProvider) EXPECT() *MockOOBProviderMockRecorder {
	return m.recorder
}

// GenerateCode mocks base method
func (m *MockOOBProvider) GenerateCode() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCode")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateCode indicates an expected call of GenerateCode
func (mr *MockOOBProviderMockRecorder) GenerateCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCode", reflect.TypeOf((*MockOOBProvider)(nil).GenerateCode))
}

// SendCode mocks base method
func (m *MockOOBProvider) SendCode(opts oob.SendCodeOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCode", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCode indicates an expected call of SendCode
func (mr *MockOOBProviderMockRecorder) SendCode(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCode", reflect.TypeOf((*MockOOBProvider)(nil).SendCode), opts)
}
