package middleware

import "github.com/gin-gonic/gin"

const VERSION = "0.0.1+dev"

func Version(c *gin.Context) {
	c.Header("X-IMPALA-WEB-VERSION", VERSION)
}
