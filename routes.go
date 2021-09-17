package main

import (
	"fmt"
)

type RoutesStruct struct {
	urlRoute      []byte
	routeMethod   []byte
	routeCallback func()
}

var routes = [1]RoutesStruct{
	RoutesStruct{
		urlRoute:      []byte("/"),
		routeMethod:   []byte("GET"),
		routeCallback: test,
	},
}

func test() {
	fmt.Println("woah")
}
