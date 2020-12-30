package maze

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//ApiError Define a error
type ApiError struct {
	Msg    string
	Status int
	Code   string
}

func (e ApiError) Error() string {
	return e.Msg
}

// Endpoints collects all of the endpoints that compose the maze  It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

// ByIDRequest collects the request parameters for the Delete method.
type ByIDRequest struct {
	ID string
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	ID   string
	Maze Maze
}

// UpdateResponse collects the response values for the Update method.
type UpdateResponse struct {
	M0 Maze
	E1 error
}

// NewEndpoints return a new instance of the endpoint that wraps the provided
func NewEndpoints(svc MazeService) (ep Endpoints) {
	ep.CreateEndpoint = MakeCreateEndpoint(svc)
	ep.DeleteEndpoint = MakeDeleteEndpoint(svc)
	ep.GetEndpoint = MakeGetEndpoint(svc)
	ep.UpdateEndpoint = MakeUpdateEndpoint(svc)
	return ep
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the
// Primarily useful in a server.
func MakeCreateEndpoint(svc MazeService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Maze)
		return svc.Create(ctx, req)
	}
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the
func MakeDeleteEndpoint(svc MazeService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ByIDRequest)
		return nil, svc.Delete(ctx, req.ID)
	}
}

// MakeGetEndpoint returns an endpoint that invokes get on the
func MakeGetEndpoint(svc MazeService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ByIDRequest)
		return svc.Get(ctx, req.ID)
	}
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the
func MakeUpdateEndpoint(svc MazeService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		return svc.Update(ctx, req.Maze)
	}
}
