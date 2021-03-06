// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite/handler/token (interfaces: TokenStorage)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	core "github.com/ory-am/fosite/handler/core"
)

// Mock of TokenStorage interface
type MockTokenStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockTokenStorageRecorder
}

// Recorder for MockTokenStorage (not exported)
type _MockTokenStorageRecorder struct {
	mock *MockTokenStorage
}

func NewMockTokenStorage(ctrl *gomock.Controller) *MockTokenStorage {
	mock := &MockTokenStorage{ctrl: ctrl}
	mock.recorder = &_MockTokenStorageRecorder{mock}
	return mock
}

func (_m *MockTokenStorage) EXPECT() *_MockTokenStorageRecorder {
	return _m.recorder
}

func (_m *MockTokenStorage) CreateAccessTokenSession(_param0 string, _param1 fosite.AccessRequester, _param2 *core.TokenSession) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenStorage) CreateRefreshTokenSession(_param0 string, _param1 fosite.AccessRequester, _param2 *core.TokenSession) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenStorage) DeleteAccessTokenSession(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenStorageRecorder) DeleteAccessTokenSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0)
}

func (_m *MockTokenStorage) DeleteRefreshTokenSession(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenStorageRecorder) DeleteRefreshTokenSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0)
}

func (_m *MockTokenStorage) GetAccessTokenSession(_param0 string, _param1 *core.TokenSession) (fosite.AccessRequester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(fosite.AccessRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokenStorageRecorder) GetAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1)
}

func (_m *MockTokenStorage) GetRefreshTokenSession(_param0 string, _param1 *core.TokenSession) (fosite.AccessRequester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(fosite.AccessRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokenStorageRecorder) GetRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1)
}
