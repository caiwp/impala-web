package router

import (
	"net/http"

	"github.com/caiwp/impala-web/server"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())

	store := sessions.NewCookieStore([]byte("secret"))
	e.Use(sessions.Sessions("session", store))
	e.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
	}))

	e.Static("/public", "./public")
	e.StaticFile("/favicon.ico", "./public/favicon.ico")

	e.HTMLRender = server.CreateMyRender()

	e.Use(middleware...)

	e.NoRoute(server.Index)
	e.GET("/login", server.Login)
	e.POST("/login", server.LoginPost)
	e.GET("/query", server.Query)
	e.POST("/query", server.PostQuery)

	return e
}
