// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	spotsdomain "github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"
	mock "github.com/stretchr/testify/mock"
)

// SpotsService is an autogenerated mock type for the SpotsService type
type SpotsService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, spt
func (_m *SpotsService) Create(ctx context.Context, spt spotsdomain.Spot) (spotsdomain.Spot, error) {
	ret := _m.Called(ctx, spt)

	var r0 spotsdomain.Spot
	if rf, ok := ret.Get(0).(func(context.Context, spotsdomain.Spot) spotsdomain.Spot); ok {
		r0 = rf(ctx, spt)
	} else {
		r0 = ret.Get(0).(spotsdomain.Spot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, spotsdomain.Spot) error); ok {
		r1 = rf(ctx, spt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, spotID
func (_m *SpotsService) Delete(ctx context.Context, spotID string) error {
	ret := _m.Called(ctx, spotID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, spotID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *SpotsService) Get(ctx context.Context, id string) (spotsdomain.Spot, error) {
	ret := _m.Called(ctx, id)

	var r0 spotsdomain.Spot
	if rf, ok := ret.Get(0).(func(context.Context, string) spotsdomain.Spot); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(spotsdomain.Spot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, pagination
func (_m *SpotsService) GetAll(ctx context.Context, pagination spotsdomain.Pagination) (spotsdomain.GetAllResponse, error) {
	ret := _m.Called(ctx, pagination)

	var r0 spotsdomain.GetAllResponse
	if rf, ok := ret.Get(0).(func(context.Context, spotsdomain.Pagination) spotsdomain.GetAllResponse); ok {
		r0 = rf(ctx, pagination)
	} else {
		r0 = ret.Get(0).(spotsdomain.GetAllResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, spotsdomain.Pagination) error); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, spot
func (_m *SpotsService) Update(ctx context.Context, spot spotsdomain.Spot) (spotsdomain.Spot, error) {
	ret := _m.Called(ctx, spot)

	var r0 spotsdomain.Spot
	if rf, ok := ret.Get(0).(func(context.Context, spotsdomain.Spot) spotsdomain.Spot); ok {
		r0 = rf(ctx, spot)
	} else {
		r0 = ret.Get(0).(spotsdomain.Spot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, spotsdomain.Spot) error); ok {
		r1 = rf(ctx, spot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}