package router

import (
	"net/http"
	"github.com/caiwp/impala-web/server"
	"github.com/gin-gonic/gin"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.Default()

	e.Static("/public", "./public")
	e.LoadHTMLGlob("templates/*")

	e.Use(middleware...)

	e.NoRoute(server.Index)
	e.GET("/query", server.Query)

	return e
}
