package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type Route struct {
	urlRoute      []byte
	routeMethod   []byte
	routeCallback func(req *fasthttp.RequestCtx)
}

var routes = []Route{
	{
		urlRoute:      []byte("/"),
		routeMethod:   []byte("GET"),
		routeCallback: test,
	},
}

func test(req *fasthttp.RequestCtx) {
	fmt.Fprintf(req, "Saying hi through route %q", string(req.Path()))
}
