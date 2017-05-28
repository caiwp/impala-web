package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/caiwp/impala-web/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bgyadmin"
	app.Commands = []cli.Command{
		cmd.Server,
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("Run app failed with %s: %v", os.Args, err)
	}
}
