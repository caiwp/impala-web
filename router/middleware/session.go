package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/model"
	"github.com/drone/drone/shared/httputil"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := httputil.GetCookie(c.Request, "user_token")
		if tk == "" {
			return
		}
		user := model.GetUserByToken(tk)
		if user.ID == 0 {
			httputil.DelCookie(c.Writer, c.Request, "user_token")
			return
		}
		sess := sessions.Default(c)
		sess.Set("user", user)

		c.Next()
	}
}

func User(c *gin.Context) *model.User {
	sess := sessions.Default(c)
	v := sess.Get("user")
	u, ok := v.(*model.User)
	if !ok {
		logrus.Debugf("sesseion get user %v", v)
		return nil
	}

	return u
}

func MustUser(c *gin.Context) {
	user := User(c)
	switch {
	case user == nil:
		c.Redirect(303, "/login")
	default:
		c.Next()
	}
}
