package router

import "obsidian_go/internal/handlers"

func CreateRoutes() *Router {
	router := &Router{}

	router.Route("GET", `/notes/(?P<id>\w+)`, handlers.CreateNoteHandler)

	return router
}
