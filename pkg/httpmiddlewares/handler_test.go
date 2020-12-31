package httpmiddlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/skyaxl/synack/pkg/apierrors"
	"github.com/stretchr/testify/assert"
)

func TestErrorEncoder_WhenErrorIsAApiError(t *testing.T) {
	recorder := httptest.NewRecorder()
	err := apierrors.APIError{Msg: "test", Status: 400, Code: "test"}
	ErrorEncoder(context.TODO(), err, recorder)
	assert.Equal(t, 400, recorder.Code)
	nerr := apierrors.APIError{}
	json.NewDecoder(recorder.Body).Decode(&nerr)
	assert.Equal(t, err, nerr)
}

func TestErrorEncoder_WhenErrorNotIsAApiError(t *testing.T) {
	recorder := httptest.NewRecorder()
	err := apierrors.APIError{Msg: "test", Status: 500, Code: "unexpected_error"}
	ErrorEncoder(context.TODO(), errors.New("test"), recorder)
	assert.Equal(t, 500, recorder.Code)
	nerr := apierrors.APIError{}
	json.NewDecoder(recorder.Body).Decode(&nerr)
	assert.Equal(t, err, nerr)
}

func TestEncodeJSONResponse_NotPanicsWithValue(t *testing.T) {
	recorder := httptest.NewRecorder()
	type Foo struct {
		B string `json:"b,omitempty"`
	}
	in := Foo{"in"}
	EncodeJSONResponse(context.TODO(), recorder, in)
	out := Foo{}
	json.NewDecoder(recorder.Body).Decode(&out)
	assert.Equal(t, out, in)
}

func TestEncodeJSONResponse_NotPanicsWithoutValue(t *testing.T) {
	recorder := httptest.NewRecorder()
	type Foo struct {
		B string `json:"b,omitempty"`
	}
	in := Foo{"in"}
	EncodeJSONResponse(context.TODO(), recorder, nil)
	out := Foo{}
	json.NewDecoder(recorder.Body).Decode(&out)
	assert.NotEqual(t, out, in)
}
