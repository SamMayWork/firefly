// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package datamocks

import (
	context "context"

	data "github.com/kaleido-io/firefly/internal/data"
	fftypes "github.com/kaleido-io/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CheckDatatype provides a mock function with given fields: ctx, datatype
func (_m *Manager) CheckDatatype(ctx context.Context, datatype *fftypes.Datatype) error {
	ret := _m.Called(ctx, datatype)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Datatype) error); ok {
		r0 = rf(ctx, datatype)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessageData provides a mock function with given fields: ctx, msg
func (_m *Manager) GetMessageData(ctx context.Context, msg *fftypes.Message) ([]*fftypes.Data, bool, error) {
	ret := _m.Called(ctx, msg)

	var r0 []*fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Message) []*fftypes.Data); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Data)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Message) bool); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *fftypes.Message) error); ok {
		r2 = rf(ctx, msg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetValidator provides a mock function with given fields: ctx, _a1
func (_m *Manager) GetValidator(ctx context.Context, _a1 *fftypes.Data) (data.Validator, error) {
	ret := _m.Called(ctx, _a1)

	var r0 data.Validator
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Data) data.Validator); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(data.Validator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Data) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}