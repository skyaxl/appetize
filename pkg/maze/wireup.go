package maze

import (
	"net/http"
)

func HttpWarmUp() (handler http.Handler) {
	mazeService := NewMazeService()
	ends := NewEndpoints(mazeService)
	handler = NewHTTPHandler(ends)
	return
}
