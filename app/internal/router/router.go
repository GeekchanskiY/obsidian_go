package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

type Router struct {
	routes []RouteEntry
}

type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:  method,
		Path:    exactPath,
		Handler: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}

	// Create a map to store URL parameters in
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}

func URLParam(r *http.Request, name string) string {
	ctx := r.Context()

	params := ctx.Value("params").(map[string]string)
	return params[name]
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r) // Log the error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue // No match found
		}

		// Create new request with params stored in context
		ctx := context.WithValue(r.Context(), "params", params)
		e.Handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	http.NotFound(w, r)
}
