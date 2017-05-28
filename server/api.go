package server

import "github.com/gin-gonic/gin"

func List(c *gin.Context) {
	data := map[string]string{
		"id":      "id1",
		"title":   "title1",
		"content": "content1",
	}
	sl := make([]map[string]string, 0)
	sl = append(sl, data)
	sl = append(sl, data)
	c.JSON(200, gin.H{
		"data": sl,
	})
}
