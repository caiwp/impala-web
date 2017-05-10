package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/model"
	"github.com/drone/drone/shared/httputil"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Login(c *gin.Context) {
	t := csrf.GetToken(c)
	c.HTML(200, tplLogin, gin.H{
		"title": "登录",
		"token": t,
	})
}

func LoginPost(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	u, err := model.SignIn(email, password)

	if err != nil {
		logrus.Warn(err)
		c.Redirect(303, "/login")
		return
	}

	ses := sessions.Default(c)
	ses.Set("email", u.Email)
	e, _ := ses.Get("email").(string)
	logrus.Info(e)

	httputil.SetCookie(c.Writer, c.Request, "user_sess", "123123")

	c.Redirect(303, "/")
}
