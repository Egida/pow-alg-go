// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// QuoteStorage is an autogenerated mock type for the quoteStorage type
type QuoteStorage struct {
	mock.Mock
}

type QuoteStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *QuoteStorage) EXPECT() *QuoteStorage_Expecter {
	return &QuoteStorage_Expecter{mock: &_m.Mock}
}

// GetRandomQuote provides a mock function with given fields: ctx
func (_m *QuoteStorage) GetRandomQuote(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteStorage_GetRandomQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRandomQuote'
type QuoteStorage_GetRandomQuote_Call struct {
	*mock.Call
}

// GetRandomQuote is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QuoteStorage_Expecter) GetRandomQuote(ctx interface{}) *QuoteStorage_GetRandomQuote_Call {
	return &QuoteStorage_GetRandomQuote_Call{Call: _e.mock.On("GetRandomQuote", ctx)}
}

func (_c *QuoteStorage_GetRandomQuote_Call) Run(run func(ctx context.Context)) *QuoteStorage_GetRandomQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QuoteStorage_GetRandomQuote_Call) Return(_a0 string, _a1 error) *QuoteStorage_GetRandomQuote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteStorage_GetRandomQuote_Call) RunAndReturn(run func(context.Context) (string, error)) *QuoteStorage_GetRandomQuote_Call {
	_c.Call.Return(run)
	return _c
}

// NewQuoteStorage creates a new instance of QuoteStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuoteStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuoteStorage {
	mock := &QuoteStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}