// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/castillofranciscodaniel/golang-beers/provider (interfaces: CurrencyClient)

// Package provider is a generated GoMock package.
package provider

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCurrencyClient is a mock of CurrencyClient interface.
type MockCurrencyClient struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyClientMockRecorder
}

// MockCurrencyClientMockRecorder is the mock recorder for MockCurrencyClient.
type MockCurrencyClientMockRecorder struct {
	mock *MockCurrencyClient
}

// NewMockCurrencyClient creates a new mock instance.
func NewMockCurrencyClient(ctrl *gomock.Controller) *MockCurrencyClient {
	mock := &MockCurrencyClient{ctrl: ctrl}
	mock.recorder = &MockCurrencyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyClient) EXPECT() *MockCurrencyClientMockRecorder {
	return m.recorder
}

// GetCurrencies mocks base method.
func (m *MockCurrencyClient) GetCurrencies() (map[string]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencies")
	ret0, _ := ret[0].(map[string]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockCurrencyClientMockRecorder) GetCurrencies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockCurrencyClient)(nil).GetCurrencies))
}