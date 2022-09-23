package mockups

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockUIDGen is a mock of UIDGen interface
type MockUIDGen struct {
	ctrl     *gomock.Controller
	recorder *MockUIDGenMockRecorder
}

// MockUIDGenMockRecorder is the mock recorder for MockUIDGen
type MockUIDGenMockRecorder struct {
	mock *MockUIDGen
}

// NewMockUIDGen creates a new mock instance
func NewMockUIDGen(ctrl *gomock.Controller) *MockUIDGen {
	mock := &MockUIDGen{ctrl: ctrl}
	mock.recorder = &MockUIDGenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUIDGen) EXPECT() *MockUIDGenMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockUIDGen) NewUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// New indicates an expected call of New
func (mr *MockUIDGenMockRecorder) NewUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUID", reflect.TypeOf((*MockUIDGen)(nil).NewUUID))
}
