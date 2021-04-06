// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kanisterio/kanister/pkg/kube/snapshot (interfaces: Snapshotter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	snapshot "github.com/kanisterio/kanister/pkg/kube/snapshot"
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	reflect "reflect"
)

// MockSnapshotter is a mock of Snapshotter interface
type MockSnapshotter struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotterMockRecorder
}

// MockSnapshotterMockRecorder is the mock recorder for MockSnapshotter
type MockSnapshotterMockRecorder struct {
	mock *MockSnapshotter
}

// NewMockSnapshotter creates a new mock instance
func NewMockSnapshotter(ctrl *gomock.Controller) *MockSnapshotter {
	mock := &MockSnapshotter{ctrl: ctrl}
	mock.recorder = &MockSnapshotterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshotter) EXPECT() *MockSnapshotterMockRecorder {
	return m.recorder
}

// Clone mocks base method
func (m *MockSnapshotter) Clone(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockSnapshotterMockRecorder) Clone(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockSnapshotter)(nil).Clone), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CloneVolumeSnapshotClass mocks base method
func (m *MockSnapshotter) CloneVolumeSnapshotClass(arg0, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneVolumeSnapshotClass", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloneVolumeSnapshotClass indicates an expected call of CloneVolumeSnapshotClass
func (mr *MockSnapshotterMockRecorder) CloneVolumeSnapshotClass(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneVolumeSnapshotClass", reflect.TypeOf((*MockSnapshotter)(nil).CloneVolumeSnapshotClass), arg0, arg1, arg2, arg3)
}

// Create mocks base method
func (m *MockSnapshotter) Create(arg0 context.Context, arg1, arg2, arg3 string, arg4 *string, arg5 bool, arg6 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSnapshotterMockRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSnapshotter)(nil).Create), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// CreateContentFromSource mocks base method
func (m *MockSnapshotter) CreateContentFromSource(arg0 context.Context, arg1 *snapshot.Source, arg2, arg3, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContentFromSource", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateContentFromSource indicates an expected call of CreateContentFromSource
func (mr *MockSnapshotterMockRecorder) CreateContentFromSource(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContentFromSource", reflect.TypeOf((*MockSnapshotter)(nil).CreateContentFromSource), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreateFromSource mocks base method
func (m *MockSnapshotter) CreateFromSource(arg0 context.Context, arg1 *snapshot.Source, arg2, arg3 string, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromSource", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFromSource indicates an expected call of CreateFromSource
func (mr *MockSnapshotterMockRecorder) CreateFromSource(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromSource", reflect.TypeOf((*MockSnapshotter)(nil).CreateFromSource), arg0, arg1, arg2, arg3, arg4)
}

// Delete mocks base method
func (m *MockSnapshotter) Delete(arg0 context.Context, arg1, arg2 string) (*v1.VolumeSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockSnapshotterMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSnapshotter)(nil).Delete), arg0, arg1, arg2)
}

// DeleteContent mocks base method
func (m *MockSnapshotter) DeleteContent(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContent indicates an expected call of DeleteContent
func (mr *MockSnapshotterMockRecorder) DeleteContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContent", reflect.TypeOf((*MockSnapshotter)(nil).DeleteContent), arg0, arg1)
}

// Get mocks base method
func (m *MockSnapshotter) Get(arg0 context.Context, arg1, arg2 string) (*v1.VolumeSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSnapshotterMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSnapshotter)(nil).Get), arg0, arg1, arg2)
}

// GetSource mocks base method
func (m *MockSnapshotter) GetSource(arg0 context.Context, arg1, arg2 string) (*snapshot.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource", arg0, arg1, arg2)
	ret0, _ := ret[0].(*snapshot.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSource indicates an expected call of GetSource
func (mr *MockSnapshotterMockRecorder) GetSource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockSnapshotter)(nil).GetSource), arg0, arg1, arg2)
}

// GetVolumeSnapshotClass mocks base method
func (m *MockSnapshotter) GetVolumeSnapshotClass(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeSnapshotClass", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeSnapshotClass indicates an expected call of GetVolumeSnapshotClass
func (mr *MockSnapshotterMockRecorder) GetVolumeSnapshotClass(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeSnapshotClass", reflect.TypeOf((*MockSnapshotter)(nil).GetVolumeSnapshotClass), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockSnapshotter) List(arg0 context.Context, arg1 string, arg2 map[string]string) (*v1.VolumeSnapshotList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeSnapshotList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSnapshotterMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSnapshotter)(nil).List), arg0, arg1, arg2)
}

// WaitOnReadyToUse mocks base method
func (m *MockSnapshotter) WaitOnReadyToUse(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitOnReadyToUse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitOnReadyToUse indicates an expected call of WaitOnReadyToUse
func (mr *MockSnapshotterMockRecorder) WaitOnReadyToUse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitOnReadyToUse", reflect.TypeOf((*MockSnapshotter)(nil).WaitOnReadyToUse), arg0, arg1, arg2)
}