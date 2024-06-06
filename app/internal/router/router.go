package router

import (
	"net/http"
)

type Router struct{}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
