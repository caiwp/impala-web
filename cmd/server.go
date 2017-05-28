package cmd

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/router"
	"github.com/caiwp/impala-web/setting"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/urfave/cli"
)

var Server = cli.Command{
	Name:   "server",
	Action: runServer,
	Before: func(c *cli.Context) error {
		setting.NewContext(c)
		return nil
	},
	After: func(c *cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 3000,
		},
	},
}

func runServer(c *cli.Context) error {
	logrus.Info("Run server")

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), setting.TimestampFormat, false),
	)

	http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), handler)
	return nil
}
