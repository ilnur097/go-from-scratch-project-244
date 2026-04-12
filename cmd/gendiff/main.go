package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",

		UsageText: "gendiff [global options]",

		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{

				Name:        "format",
				Aliases:     []string{"f"},
				Value:       "stylish",
				Usage:       "output format",
				TakesFile:   true,
				DefaultText: `"stylish"`,
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
