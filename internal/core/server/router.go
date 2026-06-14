package core_server

import (
	"fmt"
	"net/http"
)

type ApiRouter struct {
	*http.ServeMux
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{
		http.NewServeMux(),
	}
}

func (r *ApiRouter) RegisterRoutes(routers ...Route) {
	for _, route := range routers {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.Handler)
	}
}
