package router

import (
	"net/http"

	"github.com/caiwp/impala-web/router/middleware"
	"github.com/caiwp/impala-web/server"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Load(middlewares ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())

	//store := sessions.NewCookieStore([]byte("secret"))
	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	e.Use(sessions.Sessions("session", store))
	e.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
	}))

	e.Static("/public", "./public")
	e.StaticFile("/favicon.ico", "./public/favicon.ico")

	e.HTMLRender = server.CreateMyRender()

	e.Use(middlewares...)
	e.Use(middleware.SetUser())
	//e.Use(token.Refresh())

	e.GET("/", middleware.MustUser, server.Index)
	e.GET("/login", server.Login)
	e.POST("/login", server.LoginPost)
	e.GET("/query", server.Query)
	e.POST("/query", server.PostQuery)
	e.GET("/content", server.Content)
	e.GET("/json", server.Json)

	return e
}
