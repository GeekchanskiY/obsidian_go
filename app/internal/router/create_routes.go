package router

import "obsidian_go/internal/handlers"

func CreateRoutes() *Router {
	router := &Router{}
	router.Use(LoggingMiddleware)

	// Note Routes
	router.Route("GET", `/notes/(?P<id>\w+)`, handlers.SelectNoteByIdHandler)
	router.Route("GET", `/notes`, handlers.SelectNotesHandler)
	router.Route("POST", `/notes`, handlers.CreateNoteHandler)
	router.Route("DELETE", `/notes/(?P<id>\w+)`, handlers.DeleteNoteHandler)

	// Topic Routes
	router.Route("POST", `/topics`, handlers.CreateTopicHandler)
	router.Route("GET", `/topics`, handlers.SelectTopicsHandler)
	router.Route("GET", `/topics/(?P<id>\w+)`, handlers.SelectTopicByIdHandler)
	router.Route("DELETE", `/topics/(?P<id>\w+)`, handlers.DeleteTopicHandler)

	// Question Routes
	router.Route("POST", `/questions`, handlers.CreateQuestionHandler)
	router.Route("GET", `/questions`, handlers.SelectQuestionsHandler)
	router.Route("GET", `/questions/(?P<id>\w+)`, handlers.SelectQuestionByIdHandler)
	router.Route("DELETE", `/questions/(?P<id>\w+)`, handlers.DeleteQuestionHandler)

	return router
}
