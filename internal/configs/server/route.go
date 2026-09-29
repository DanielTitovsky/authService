package app_server

import (
	"authService/internal/midelware"
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []midelware.Middleware
}

func NewRoute(method string, path string, handler http.HandlerFunc, middlewares []midelware.Middleware) Route {
	return Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}
}
