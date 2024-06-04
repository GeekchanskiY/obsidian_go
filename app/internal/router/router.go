package router

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Handler     http.Handler
	Methods     []string
	Middlewares []http.Handler
}

type SimpleJsonResponse struct {
	Message string `json:"message"`
}

func NewRoute(handler http.Handler, methods ...string) *Route {
	return &Route{
		Handler:     handler,
		Methods:     methods,
		Middlewares: []http.Handler{},
	}
}

func MethodNotFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(SimpleJsonResponse{
		Message: "Method not found"})

}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for _, method := range r.Methods {
		if req.Method == method {
			r.Handler.ServeHTTP(w, req)
			return
		}
	}

	MethodNotFoundHandler(w, req)
}
