package main

import (
	"bytes"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		for _, route := range routes {
			if bytes.Compare(ctx.Path(), route.urlRoute) == 0 && bytes.Compare(ctx.Method(), route.routeMethod) == 0 {
				route.routeCallback()
			}
		}
	}

	s := &fasthttp.Server{
		Handler: requestHandler,
		Name:    "Server wot",
	}

	if err := s.ListenAndServe("127.0.0.1:8080"); err != nil {
		log.Fatalf("error in ListenAndServer: %s", err)
	}
}
