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
	authGroup.POST("/login", api.AuthLoginPost)
	authGroup.POST("/register", api.AuthRegisterPost)
	authGroup.POST("/activate", api.AuthActivatePost)
	authGroup.POST("/confitm", api.AuthConfirmPost)

	router.Use(jwt.AuthMiddleware())
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handler)
	}
	return router
}

var routes = Routes{
	{
		"UserUpdate",
		"PUT",
		"/api/user/:userid",
		api.UserUpdate,
	},
}
