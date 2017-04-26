package router

import (
	"net/http"

	"github.com/caiwp/impala-web/server"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())

	e.Static("/public", "./public")
	e.HTMLRender = createMyRender()

	e.Use(middleware...)

	e.NoRoute(server.Index)
	e.GET("/query", server.Query)
	e.POST("/query", server.PostQuery)

	return e
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/base.html", "templates/layout.html", "templates/index.html")
	r.AddFromFiles("query", "templates/base.html", "templates/layout.html", "templates/query.html")

	// r.AddFromFiles("login", "base.html", "login.html")
	// r.AddFromFiles("dashboard", "base.html", "dashboard.html")

	return r
}
