package router

import (
	"known-anchors/api"
	"known-anchors/middleware/jwt"
	"known-anchors/service"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler gin.HandlerFunc
}

type Routes []Route

func NewRouter(s *service.ServiceContext) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("service", s)
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
		router.Handle(route.Method, route.Pattern, route.Handler)
	}
}

var authRoutes = Routes{
	{
		"AuthLoginPost",
		"POST",
		"/login",
		api.AuthLoginPost,
	},
	{
		"AuthRegisterPost",
		"POST",
		"/register",
		api.AuthRegisterPost,
	},
	{
		"AuthConfirmPost",
		"POST",
		"/confirm",
		api.AuthConfirmPost,
	},
}

var deckRoutes = Routes{
	{
		"CreateDeck",
		"POST",
		"/",
		api.DeckCreate,
	},
	{
		"GetDeck",
		"GET",
		"/:id",
		api.DeckGet,
	},
	{
		"UpdateDeck",
		"PUT",
		"/",
		api.DeckUpdate,
	},
	{
		"DeleteDeck",
		"DELETE",
		"/:id",
		api.DeckDelete,
	},
	{
		"ListDecks",
		"GET",
		"/list",
		api.DeckList,
	},
}

var routes = Routes{
	{
		"UserUpdate",
		"PUT",
		"/user/:userid",
		api.UserUpdate,
	},
}
