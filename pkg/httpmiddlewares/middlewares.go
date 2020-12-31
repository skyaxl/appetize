package httpmiddlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//LogHandler to logger request
type LogHandler struct {
	Original http.Handler
}

func (lh LogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var bts []byte
	if req.Body != nil {
		bts, _ = ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(bts))
	}

	fmt.Printf("[Request][Path:'%s'][Body: %s]", req.URL.String(), string(bts))
	lh.Original.ServeHTTP(res, req)
	if req.Response != nil {
		if req.Response.Body != nil {
			bts, _ = ioutil.ReadAll(req.Body)
			req.Response.Body = ioutil.NopCloser(bytes.NewReader(bts))
		}
		fmt.Printf("[Response][Path: '%s'][Status: %s][Body: %s][ResponseBody: %s]", req.URL.String(), req.Response.Status, string(bts), string(bts))
	}

}
