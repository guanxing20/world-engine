// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockSocialConfigFacebookInstantGame is an autogenerated mock type for the SocialConfigFacebookInstantGame type
type MockSocialConfigFacebookInstantGame struct {
	mock.Mock
}

type MockSocialConfigFacebookInstantGame_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSocialConfigFacebookInstantGame) EXPECT() *MockSocialConfigFacebookInstantGame_Expecter {
	return &MockSocialConfigFacebookInstantGame_Expecter{mock: &_m.Mock}
}

// GetAppSecret provides a mock function with no fields
func (_m *MockSocialConfigFacebookInstantGame) GetAppSecret() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAppSecret")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSocialConfigFacebookInstantGame_GetAppSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAppSecret'
type MockSocialConfigFacebookInstantGame_GetAppSecret_Call struct {
	*mock.Call
}

// GetAppSecret is a helper method to define mock.On call
func (_e *MockSocialConfigFacebookInstantGame_Expecter) GetAppSecret() *MockSocialConfigFacebookInstantGame_GetAppSecret_Call {
	return &MockSocialConfigFacebookInstantGame_GetAppSecret_Call{Call: _e.mock.On("GetAppSecret")}
}

func (_c *MockSocialConfigFacebookInstantGame_GetAppSecret_Call) Run(run func()) *MockSocialConfigFacebookInstantGame_GetAppSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSocialConfigFacebookInstantGame_GetAppSecret_Call) Return(_a0 string) *MockSocialConfigFacebookInstantGame_GetAppSecret_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSocialConfigFacebookInstantGame_GetAppSecret_Call) RunAndReturn(run func() string) *MockSocialConfigFacebookInstantGame_GetAppSecret_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSocialConfigFacebookInstantGame creates a new instance of MockSocialConfigFacebookInstantGame. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSocialConfigFacebookInstantGame(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSocialConfigFacebookInstantGame {
	mock := &MockSocialConfigFacebookInstantGame{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
