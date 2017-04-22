package server

import "github.com/gin-gonic/gin"

func Query(c *gin.Context) {
	c.HTML(200, "query.tmpl", gin.H{
		"title": "请求",
	})
}
