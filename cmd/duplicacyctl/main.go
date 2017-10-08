package main

import (
	"fmt"
	"github.com/joharohl/duplicacyctl/src/config"
	"gopkg.in/urfave/cli.v1"
	"os"
	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()

	app := cli.NewApp()
	app.Name = "duplicacyctl"
	app.Usage = "Control your backups"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "./duplicacyctl.yaml",
			Usage: "Config to use.",
		},
	}

	app.Action = func(c *cli.Context) error {
		conf := config.New()
		conf.Load(fs, c.String("config"))
		fmt.Printf(conf.String())
		return nil
	}

	app.Run(os.Args)
}
