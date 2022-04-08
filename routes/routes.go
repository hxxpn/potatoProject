package routes

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func newRoute(method string, pattern string, handler http.HandlerFunc) Route {
	return Route{
		Name:        "",
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handler,
	}
}
