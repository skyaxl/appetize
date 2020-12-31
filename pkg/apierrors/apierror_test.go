package apierrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	assert.Equal(t, "test", APIError{Msg: "test"}.Error())
}
