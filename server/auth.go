package server

import (
	"github.com/caiwp/impala-web/model"
	"github.com/drone/drone/shared/httputil"
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

	u := model.LoginIn(email, password)
	if u.ID == 0 || u.Token == "" {
		c.Redirect(303, "/login?err=invalid_user")
		return
	}

	httputil.SetCookie(c.Writer, c.Request, "user_token", u.Token)

	c.Redirect(303, "/")
	return
}
