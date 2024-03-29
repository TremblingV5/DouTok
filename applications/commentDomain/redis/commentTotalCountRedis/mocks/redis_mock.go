// Code generated by MockGen. DO NOT EDIT.
// Source: redis.go

// Package ComentTotalCountRedisRepositoryMocks is a generated GoMock package.
package ComentTotalCountRedisRepositoryMocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIClient is a mock of IClient interface.
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient.
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance.
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIClient) Delete(ctx context.Context, videoId ...int64) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range videoId {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIClientMockRecorder) Delete(ctx interface{}, videoId ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, videoId...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIClient)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockIClient) Get(ctx context.Context, videoId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, videoId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIClientMockRecorder) Get(ctx, videoId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIClient)(nil).Get), ctx, videoId)
}
