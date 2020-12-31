package httpmiddlewares

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogHandler_ServeHTTPWithBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	lh := LogHandler{Original: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		return
	})}
	req := httptest.NewRequest("POST", "/", strings.NewReader("{}"))

	assert.NotPanics(t, func() {
		lh.ServeHTTP(recorder, req)
	})
}

func TestLogHandler_ServeHTTPWithBody_AndResponseBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	lh := LogHandler{Original: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		return
	})}
	req := httptest.NewRequest("POST", "/", strings.NewReader("{}"))
	req.Response = &http.Response{
		Body: ioutil.NopCloser(strings.NewReader("{}")),
	}
	assert.NotPanics(t, func() {
		lh.ServeHTTP(recorder, req)
	})
}

func TestLogHandler_ServeHTTPWithBody_AndResponseWithoutBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	lh := LogHandler{Original: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		return
	})}
	req := httptest.NewRequest("POST", "/", strings.NewReader("{}"))
	req.Response = &http.Response{}
	assert.NotPanics(t, func() {
		lh.ServeHTTP(recorder, req)
	})
}
