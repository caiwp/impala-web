package router

import (
	"net/http"

	"github.com/caiwp/impala-web/server"
	"github.com/gin-contrib/multitemplate"
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

	return r
}
