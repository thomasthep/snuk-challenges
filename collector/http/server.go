package http

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var router fasthttprouter.Router

func init() {
	router = fasthttprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}

func AddRoute(method string, path string, handle fasthttp.RequestHandler) {
	router.Handle(method, path, handle)
}

func Start() {
	go func() {
		fasthttp.ListenAndServe(":3000", router.Handler)
	}()
}
