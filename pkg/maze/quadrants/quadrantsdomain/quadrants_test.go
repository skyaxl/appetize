package quadrantsdomain

import (
	"net/http"
	"testing"

	"github.com/skyaxl/synack/pkg/apierrors"
	"github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"

	"github.com/stretchr/testify/assert"
)

func TestQuadrant_ValidateWithInvalidStartPoint(t *testing.T) {
	s := Quadrant{
		Start: spotsdomain.Spot{
			X: 2,
			Y: 3,
		},
	}
	expec := apierrors.APIError{Msg: "Start spot is not orthogonal a valid value its (0,0) (2,2)", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}

func TestQuadrant_ValidateWithInvalidHorizontalDirection(t *testing.T) {
	s := Quadrant{
		Start: spotsdomain.Spot{
			X: 2,
			Y: 2,
		},
		HorizontalDirection: HorizontalDirection("upset"),
	}
	expec := apierrors.APIError{Msg: "horizontal_direction must have a valid value 'left' or 'right'", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}

func TestQuadrant_ValidateWithInvalidVerticalDirection(t *testing.T) {
	s := Quadrant{
		Start: spotsdomain.Spot{
			X: 2,
			Y: 2,
		},
		HorizontalDirection: LeftHorizontalDirection,
		VerticalDirection:   VerticalDirection("don`t let me down"),
	}
	expec := apierrors.APIError{Msg: "vertical_direction must have a valid value 'top' or 'bottom'", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}
