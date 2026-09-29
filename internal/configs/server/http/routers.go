package app_http_server

import "net/http"

type Router struct {
	ApiVersion string
	serverMux  *http.ServeMux
}

func NewRouter(apiVersion string, serverMux *http.ServeMux) Router {
	return Router{
		ApiVersion: apiVersion,
		serverMux:  serverMux,
	}
}

func (r *Router) RegisterRouters(routers ...Route) {
	for _, route := range routers {
		handler := http.Handler(route.Handler)

		for i := len(route.Middlewares) - 1; i >= 0; i-- {
			handler = route.Middlewares[i](handler)
		}

		fullPath := route.Method + " /api/" + r.ApiVersion + route.Path

		r.serverMux.Handle(fullPath, handler)
	}
}
