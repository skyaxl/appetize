package maze

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skyaxl/synack/pkg/maze/paths"
	"github.com/skyaxl/synack/pkg/maze/quadrants"
	"github.com/skyaxl/synack/pkg/maze/spots"
)

func HttpWarmUp() (handler http.Handler) {
	m := mux.NewRouter()
	spotsService := spots.NewSpotService()
	spotsEndpoints := spots.NewEndpoints(spotsService)
	spots.MapHandlers(m, spotsEndpoints)

	pathsService := paths.NewPathService()
	pathsEndpoints := paths.NewEndpoints(pathsService)
	paths.MapHandlers(m, pathsEndpoints)

	quadrantsService := quadrants.NewQuadrantService()
	quadrantsEndpoints := quadrants.NewEndpoints(quadrantsService)
	quadrants.MapHandlers(m, quadrantsEndpoints)

	return m
}
