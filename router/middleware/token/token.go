package token

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Debug("csrf2")
		c.Next()
	}
}
