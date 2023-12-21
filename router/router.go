package router

import (
	"known-anchors/api"
	"known-anchors/middleware/jwt"
	"known-anchors/middleware/mailproducer"
	"known-anchors/service"
	"known-anchors/util/close"
	"known-anchors/util/mail"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name     string
	Method   string
	Pattern  string
	Handlers []gin.HandlerFunc
}

type Routes []Route

func NewRouter(s *service.ServiceContext, cl *close.Close) *gin.Engine {
	router := gin.Default()
	mailconsumer := mail.NewMailConsumer()
	go mailconsumer.Consume()
	cl.AddCloseable(mailconsumer)
	router.Use(func(c *gin.Context) {
		c.Set("service", s)
		c.Set("closechan", cl)
		c.Next()
	})
	authGroup := router.Group("/api/auth")
	addRoutes(authGroup, authRoutes)
	router.Use(jwt.AuthMiddleware())
	deckGroup := router.Group("/api/decks")
	addRoutes(deckGroup, deckRoutes)
	addRoutes(router.Group("/api"), routes)
	return router
}

func addRoutes(router *gin.RouterGroup, routes Routes) {
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handlers...)
	}
}

var authRoutes = Routes{
	{
		"AuthLoginPost",
		"POST",
		"/login",
		[]gin.HandlerFunc{api.AuthLoginPost},
	},
	{
		"AuthRegisterPost",
		"POST",
		"/register",
		[]gin.HandlerFunc{mailproducer.MailProducerMiddleWare(), api.AuthRegisterPost},
	},
	{
		"AuthConfirmPost",
		"POST",
		"/confirm",
		[]gin.HandlerFunc{api.AuthConfirmPost},
	},
}

var deckRoutes = Routes{
	{
		"CreateDeck",
		"POST",
		"/",
		[]gin.HandlerFunc{api.DeckCreate},
	},
	{
		"GetDeck",
		"GET",
		"/:id",
		[]gin.HandlerFunc{api.DeckGet},
	},
	{
		"UpdateDeck",
		"PUT",
		"/",
		[]gin.HandlerFunc{api.DeckUpdate},
	},
	{
		"DeleteDeck",
		"DELETE",
		"/:id",
		[]gin.HandlerFunc{api.DeckDelete},
	},
	{
		"ListDecks",
		"GET",
		"/list",
		[]gin.HandlerFunc{api.DeckList},
	},
}

var routes = Routes{
	{
		"UserUpdate",
		"PUT",
		"/user/:userid",
		[]gin.HandlerFunc{api.UserUpdate},
	},
}
