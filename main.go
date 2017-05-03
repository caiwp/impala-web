package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/router"
	"github.com/caiwp/impala-web/router/middleware"
	"github.com/gin-gonic/contrib/ginrus"
	"flag"
	"fmt"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "listen port")
}

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false),
		middleware.Version,
	)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
