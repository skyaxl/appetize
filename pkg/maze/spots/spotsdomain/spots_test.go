package spotsdomain

import (
	"net/http"
	"testing"

	"github.com/skyaxl/synack/pkg/apierrors"

	"github.com/stretchr/testify/assert"
)

func TestSpot_ValidateWithInvalidName(t *testing.T) {
	s := Spot{}
	expec := apierrors.APIError{Msg: "Name must have a valid value not empty", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}

func TestSpot_ValidateWithInvalidAmount(t *testing.T) {
	s := Spot{Name: "Entrance"}
	expec := apierrors.APIError{Msg: "Amount must have a valid value > 0 ", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}

func TestSpot_ValidateOk(t *testing.T) {
	s := Spot{Name: "Entrance", Amount: 1}

	assert.Nil(t, s.Validate())
}
