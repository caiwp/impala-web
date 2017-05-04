package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/router"
	"github.com/caiwp/impala-web/router/middleware"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/lestrrat/go-file-rotatelogs"
)

var (
	port int

	logPath string
)

func init() {
	flag.IntVar(&port, "port", 3000, "listen port")
	flag.StringVar(&logPath, "log", "./logs", "log base path")
}

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false),
		middleware.Version,
	)

	rl, err := rotatelogs.New(logPath+"/log.%Y%m%d")

	if err != nil {
		panic(err)
	}

	logrus.SetOutput(rl)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
