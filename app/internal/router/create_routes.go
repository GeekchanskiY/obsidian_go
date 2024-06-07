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
	router.Route("PUT", `/notes/(?P<id>\w+)`, handlers.UpdateNoteHandler)

	// Topic Routes
	router.Route("POST", `/topics`, handlers.CreateTopicHandler)
	router.Route("GET", `/topics`, handlers.SelectTopicsHandler)
	router.Route("GET", `/topics/(?P<id>\w+)`, handlers.SelectTopicByIdHandler)
	router.Route("DELETE", `/topics/(?P<id>\w+)`, handlers.DeleteTopicHandler)
	router.Route("PUT", `/topics/(?P<id>\w+)`, handlers.UpdateTopicHandler)

	// Question Routes
	router.Route("POST", `/questions`, handlers.CreateQuestionHandler)
	router.Route("GET", `/questions`, handlers.SelectQuestionsHandler)
	router.Route("GET", `/questions/(?P<id>\w+)`, handlers.SelectQuestionByIdHandler)
	router.Route("DELETE", `/questions/(?P<id>\w+)`, handlers.DeleteQuestionHandler)
	router.Route("PUT", `/questions/(?P<id>\w+)`, handlers.UpdateQuestionHandler)

	// Answer Routes
	router.Route("POST", `/answers`, handlers.CreateAnswerHandler)
	router.Route("GET", `/answers`, handlers.SelectAnswersHandler)
	router.Route("GET", `/answers/(?P<id>\w+)`, handlers.SelectAnswerByIdHandler)
	router.Route("PUT", `/answers/(?P<id>\w+)`, handlers.UpdateAnswerHandler)
	router.Route("DELETE", `/answers/(?P<id>\w+)`, handlers.DeleteAnswerHandler)

	return router
}
