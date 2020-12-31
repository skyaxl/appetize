package spots

import (
	"context"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"

	"github.com/stretchr/testify/mock"
)

func TestMakeCreateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.SpotsService)
	spot := spotsdomain.Spot{}
	m.On("Create", mock.Anything, spot).Return(spot, nil)
	MakeCreateEndpoint(m)(context.TODO(), spot)
}

func TestMakeDeleteEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.SpotsService)
	req := ByIDRequest{ID: "test"}
	m.On("Delete", mock.Anything, "test").Return(nil)
	MakeDeleteEndpoint(m)(context.TODO(), req)
}

func TestMakeGetEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.SpotsService)
	req := ByIDRequest{ID: "test"}
	spot := spotsdomain.Spot{}
	m.On("Get", mock.Anything, "test").Return(spot, nil)
	MakeGetEndpoint(m)(context.TODO(), req)
}

func TestMakeUpdateEndpoint_NotPanics(t *testing.T) {
	m := new(mocks.SpotsService)
	spot := spotsdomain.Spot{}
	req := UpdateRequest{ID: "test", Spot: spot}
	m.On("Update", mock.Anything, spot).Return(spot, nil)
	MakeUpdateEndpoint(m)(context.TODO(), req)
}
