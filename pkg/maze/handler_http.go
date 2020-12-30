package maze

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints Endpoints) http.Handler {
	m := mux.NewRouter()

	m.Use(func(h http.Handler) http.Handler {
		return LogHandler{h}
	})
	m.Handle("/mazes", httptransport.NewServer(
		endpoints.CreateEndpoint,
		DecodeCreateRequest,
		EncodeJSONResponse,
		httptransport.ServerErrorEncoder(ErrorEncoder),
	)).Methods(http.MethodPost)

	m.Handle("/mazes/{id}", httptransport.NewServer(
		endpoints.DeleteEndpoint,
		DecodeByIDRequest,
		EncodeJSONResponse,
		httptransport.ServerErrorEncoder(ErrorEncoder),
	)).Methods(http.MethodGet)

	m.Handle("/mazes/{id}", httptransport.NewServer(
		endpoints.DeleteEndpoint,
		DecodeByIDRequest,
		EncodeJSONResponse,
		httptransport.ServerErrorEncoder(ErrorEncoder),
	)).Methods(http.MethodDelete)

	m.Handle("/mazes/{id}", httptransport.NewServer(
		endpoints.UpdateEndpoint,
		DecodeUpdateRequest,
		EncodeJSONResponse,
		httptransport.ServerErrorEncoder(ErrorEncoder),
	)).Methods(http.MethodPut)

	return m
}

// DecodeCreateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeCreateRequest(_ context.Context, r *http.Request) (_ interface{}, err error) {
	m := Maze{}
	if err = json.NewDecoder(r.Body).Decode(&m); err != nil {
		return
	}
	if err = m.Validate(); err != nil {
		return
	}

	return m, err
}

// DecodeByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeByIDRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = ByIDRequest{
		ID: mux.Vars(r)["id"],
	}

	return req, err
}

//EncodeJSONResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeUpdateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := UpdateRequest{
		ID: mux.Vars(r)["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req.Maze)
	return req, err
}

//ErrorEncoder intercept errors to use them to response
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if apiError, ok := err.(ApiError); ok {
		w.WriteHeader(apiError.Status)
		_ = json.NewEncoder(w).Encode(apiError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&ApiError{
		err.Error(),
		http.StatusInternalServerError,
		"unexpected_error",
	})
}
