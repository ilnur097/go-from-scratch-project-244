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
		
		HideHelpCommand: true,
		UsageText: "gendiff [global options]",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
