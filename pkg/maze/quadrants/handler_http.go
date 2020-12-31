package quadrants

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/skyaxl/synack/pkg/httpmiddlewares"
	"github.com/skyaxl/synack/pkg/maze/quadrants/quadrantsdomain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapHandlers returns a handler that makes a set of endpoints available on
// predefined paths.
func MapHandlers(m *mux.Router, endpoints Endpoints) {

	m.Use(func(h http.Handler) http.Handler {
		return httpmiddlewares.LogHandler{h}
	})
	m.Handle("/quadrants", httptransport.NewServer(
		endpoints.CreateEndpoint,
		DecodeCreateRequest,
		httpmiddlewares.EncodeJSONResponse,
		httptransport.ServerErrorEncoder(httpmiddlewares.ErrorEncoder),
	)).Methods(http.MethodPost)

	m.Handle("/quadrants", httptransport.NewServer(
		endpoints.GetAllEndpoint,
		DecodeGetAllRequest,
		httpmiddlewares.EncodeJSONResponse,
		httptransport.ServerErrorEncoder(httpmiddlewares.ErrorEncoder),
	)).Methods(http.MethodGet)

	m.Handle("/quadrants/{id}", httptransport.NewServer(
		endpoints.GetEndpoint,
		DecodeByIDRequest,
		httpmiddlewares.EncodeJSONResponse,
		httptransport.ServerErrorEncoder(httpmiddlewares.ErrorEncoder),
	)).Methods(http.MethodGet)

	m.Handle("/quadrants/{id}", httptransport.NewServer(
		endpoints.DeleteEndpoint,
		DecodeByIDRequest,
		httpmiddlewares.EncodeJSONResponse,
		httptransport.ServerErrorEncoder(httpmiddlewares.ErrorEncoder),
	)).Methods(http.MethodDelete)

	m.Handle("/quadrants/{id}", httptransport.NewServer(
		endpoints.UpdateEndpoint,
		DecodeUpdateRequest,
		httpmiddlewares.EncodeJSONResponse,
		httptransport.ServerErrorEncoder(httpmiddlewares.ErrorEncoder),
	)).Methods(http.MethodPut)

	return
}

// DecodeCreateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeCreateRequest(_ context.Context, r *http.Request) (_ interface{}, err error) {
	m := quadrantsdomain.Quadrant{}
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

// DecodeGetAllRequest is a transport/http.DecodeRequestFunc that decodes a
func DecodeGetAllRequest(_ context.Context, r *http.Request) (_ interface{}, err error) {
	pagi := quadrantsdomain.Pagination{}
	err = schema.NewDecoder().Decode(&pagi, r.URL.Query())
	if err != nil {
		return nil, err
	}

	return pagi, err
}

// DecodeUpdateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := UpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.Quadrant)
	if err != nil {
		return nil, err
	}
	if err = req.Quadrant.Validate(); err != nil {
		return nil, err
	}
	req.Quadrant.ID, _ = primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	return req, err
}
