package httpmiddlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/skyaxl/synack/pkg/apierrors"
)

//ErrorEncoder intercept errors to use them to response
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if apiError, ok := err.(apierrors.APIError); ok {
		w.WriteHeader(apiError.Status)
		_ = json.NewEncoder(w).Encode(apiError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&apierrors.APIError{
		Msg:    err.Error(),
		Status: http.StatusInternalServerError,
		Code:   "unexpected_error",
	})
}

//EncodeJSONResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}
