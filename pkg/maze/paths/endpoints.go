package paths

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/skyaxl/synack/pkg/maze/paths/pathsdomain"
)

// Endpoints collects all of the endpoints that compose the path  It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetAllEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

// ByIDRequest collects the request parameters for the Delete method.
type ByIDRequest struct {
	ID string
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	ID   string
	Path pathsdomain.Path
}

// UpdateResponse collects the response values for the Update method.
type UpdateResponse struct {
	M0 pathsdomain.Path
	E1 error
}

// NewEndpoints return a new instance of the endpoint that wraps the provided
func NewEndpoints(svc pathsdomain.PathsService) (ep Endpoints) {
	ep.CreateEndpoint = MakeCreateEndpoint(svc)
	ep.DeleteEndpoint = MakeDeleteEndpoint(svc)
	ep.GetEndpoint = MakeGetEndpoint(svc)
	ep.GetAllEndpoint = MakeGetAllEndpoint(svc)
	ep.UpdateEndpoint = MakeUpdateEndpoint(svc)
	return ep
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the
// Primarily useful in a server.
func MakeCreateEndpoint(svc pathsdomain.PathsService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pathsdomain.Path)
		return svc.Create(ctx, req)
	}
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the
func MakeDeleteEndpoint(svc pathsdomain.PathsService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ByIDRequest)
		return nil, svc.Delete(ctx, req.ID)
	}
}

// MakeGetEndpoint returns an endpoint that invokes get on the
func MakeGetEndpoint(svc pathsdomain.PathsService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ByIDRequest)
		return svc.Get(ctx, req.ID)
	}
}

// MakeGetAllEndpoint returns an endpoint that invokes get on the
func MakeGetAllEndpoint(svc pathsdomain.PathsService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pathsdomain.Pagination)
		return svc.GetAll(ctx, req)
	}
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the
func MakeUpdateEndpoint(svc pathsdomain.PathsService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		return svc.Update(ctx, req.Path)
	}
}
