package router

import "obsidian_go/internal/handlers"

func CreateRoutes() *Router {
	router := &Router{}
	router.Use(LoggingMiddleware)
	router.Route("GET", `/notes/(?P<id>\w+)`, handlers.SelectNoteByIdHandler)
	router.Route("GET", `/notes`, handlers.SelectNotesHandler)
	router.Route("POST", `/notes`, handlers.CreateNoteHandler)

	return router
}
