package paths

import (
	"context"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/paths/pathsdomain"

	"github.com/stretchr/testify/mock"
)

func TestMakeCreateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.PathsService)
	path := pathsdomain.Path{}
	m.On("Create", mock.Anything, path).Return(path, nil)
	MakeCreateEndpoint(m)(context.TODO(), path)
}

func TestMakeDeleteEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.PathsService)
	req := ByIDRequest{ID: "test"}
	m.On("Delete", mock.Anything, "test").Return(nil)
	MakeDeleteEndpoint(m)(context.TODO(), req)
}

func TestMakeGetEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.PathsService)
	req := ByIDRequest{ID: "test"}
	path := pathsdomain.Path{}
	m.On("Get", mock.Anything, "test").Return(path, nil)
	MakeGetEndpoint(m)(context.TODO(), req)
}

func TestMakeUpdateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.PathsService)
	path := pathsdomain.Path{}
	req := UpdateRequest{ID: "test", Path: path}
	m.On("Update", mock.Anything, path).Return(path, nil)
	MakeUpdateEndpoint(m)(context.TODO(), req)
}
