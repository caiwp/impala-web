package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Index(c *gin.Context) {
	s := sessions.Default(c)
	var cnt int
	v := s.Get("cnt")
	if v == nil {
		cnt = 0;
	} else {
		cnt = v.(int)
		cnt ++
	}
	s.Set("cnt", cnt)
	s.Save()

	logrus.Info(cnt)

	t := csrf.GetToken(c)
	c.Header("X-CSRF-TOKEN", t)
	logrus.Info(t)

	c.HTML(200, "index", gin.H{
		"title": "首页",
	})
}
