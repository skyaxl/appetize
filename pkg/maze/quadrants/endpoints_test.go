package quadrants

import (
	"context"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/quadrants/quadrantsdomain"

	"github.com/stretchr/testify/mock"
)

func TestMakeCreateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.QuadrantsService)
	quadrant := quadrantsdomain.Quadrant{}
	m.On("Create", mock.Anything, quadrant).Return(quadrant, nil)
	MakeCreateEndpoint(m)(context.TODO(), quadrant)
}

func TestMakeDeleteEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.QuadrantsService)
	req := ByIDRequest{ID: "test"}
	m.On("Delete", mock.Anything, "test").Return(nil)
	MakeDeleteEndpoint(m)(context.TODO(), req)
}

func TestMakeGetEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.QuadrantsService)
	req := ByIDRequest{ID: "test"}
	quadrant := quadrantsdomain.Quadrant{}
	m.On("Get", mock.Anything, "test").Return(quadrant, nil)
	MakeGetEndpoint(m)(context.TODO(), req)
}

func TestMakeUpdateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.QuadrantsService)
	quadrant := quadrantsdomain.Quadrant{}
	req := UpdateRequest{ID: "test", Quadrant: quadrant}
	m.On("Update", mock.Anything, quadrant).Return(quadrant, nil)
	MakeUpdateEndpoint(m)(context.TODO(), req)
}
