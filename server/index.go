package server

import (
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Index(c *gin.Context) {
	t := csrf.GetToken(c)
	c.Header("X-CSRF-TOKEN", t)
	c.HTML(200, tplIndex, gin.H{
		"title": "首页",
	})
}
