package pathsdomain

import (
	"net/http"
	"testing"

	"github.com/skyaxl/synack/pkg/apierrors"

	"github.com/stretchr/testify/assert"
)

func TestPath_ValidateWithInvalidName(t *testing.T) {
	s := Path{Distance: -1}
	expec := apierrors.APIError{Msg: "Distance must be greather than 0", Status: http.StatusBadRequest}
	assert.Equal(t, expec, s.Validate())
}
