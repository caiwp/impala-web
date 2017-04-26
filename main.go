package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/router"
	"github.com/caiwp/impala-web/router/middleware"
	"github.com/gin-gonic/contrib/ginrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false),
		middleware.Version,
	)
	http.ListenAndServe(":3000", handler)
}
