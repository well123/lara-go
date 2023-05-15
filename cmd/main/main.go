package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"goApi/internal/app"
	"os"
)
import "goApi/pkg/logger"

var VERSION = "1.0.0"

func main() {
	ctx := logger.NewTagCtx(context.Background(), "__main__")

	cliApp := cli.NewApp()
	cliApp.Name = "go-api"
	cliApp.Version = VERSION
	cliApp.Commands = []*cli.Command{
		newCommand(ctx),
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newCommand(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "create a new project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./configs/config.toml",
				Usage:   "App configuration file(.json,.yaml,.toml)",
			},
			&cli.StringFlag{
				Name:    "model",
				Aliases: []string{"m"},
				Value:   "./configs/model.conf",
				Usage:   "Casbin model configuration(.conf)",
			},
			&cli.StringFlag{
				Name:    "www",
				Aliases: []string{"w"},
				Value:   "./dist/static",
				Usage:   "Static file directory",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("config")),
				app.SetModelFile(c.String("model")),
				app.SetVersion(VERSION),
				app.SetWWWDir(c.String("www")),
			)
		},
	}
}
