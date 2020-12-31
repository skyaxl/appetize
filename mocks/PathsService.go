// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	pathsdomain "github.com/skyaxl/synack/pkg/maze/paths/pathsdomain"
	mock "github.com/stretchr/testify/mock"
)

// PathsService is an autogenerated mock type for the PathsService type
type PathsService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, spt
func (_m *PathsService) Create(ctx context.Context, spt pathsdomain.Path) (pathsdomain.Path, error) {
	ret := _m.Called(ctx, spt)

	var r0 pathsdomain.Path
	if rf, ok := ret.Get(0).(func(context.Context, pathsdomain.Path) pathsdomain.Path); ok {
		r0 = rf(ctx, spt)
	} else {
		r0 = ret.Get(0).(pathsdomain.Path)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pathsdomain.Path) error); ok {
		r1 = rf(ctx, spt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, pathID
func (_m *PathsService) Delete(ctx context.Context, pathID string) error {
	ret := _m.Called(ctx, pathID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, pathID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *PathsService) Get(ctx context.Context, id string) (pathsdomain.Path, error) {
	ret := _m.Called(ctx, id)

	var r0 pathsdomain.Path
	if rf, ok := ret.Get(0).(func(context.Context, string) pathsdomain.Path); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pathsdomain.Path)
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
func (_m *PathsService) GetAll(ctx context.Context, pagination pathsdomain.Pagination) (pathsdomain.GetAllResponse, error) {
	ret := _m.Called(ctx, pagination)

	var r0 pathsdomain.GetAllResponse
	if rf, ok := ret.Get(0).(func(context.Context, pathsdomain.Pagination) pathsdomain.GetAllResponse); ok {
		r0 = rf(ctx, pagination)
	} else {
		r0 = ret.Get(0).(pathsdomain.GetAllResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pathsdomain.Pagination) error); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, path
func (_m *PathsService) Update(ctx context.Context, path pathsdomain.Path) (pathsdomain.Path, error) {
	ret := _m.Called(ctx, path)

	var r0 pathsdomain.Path
	if rf, ok := ret.Get(0).(func(context.Context, pathsdomain.Path) pathsdomain.Path); ok {
		r0 = rf(ctx, path)
	} else {
		r0 = ret.Get(0).(pathsdomain.Path)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pathsdomain.Path) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}