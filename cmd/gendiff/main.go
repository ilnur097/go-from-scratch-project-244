package main

import (
	"context" 
	"fmt"
	"log"
	"os"
	"code"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 2 {
				return fmt.Errorf("two file paths are required")
			}

			path1 := cmd.Args().Get(0)
			path2 := cmd.Args().Get(1)

			data1, err := code.ParseFile(path1)
			if err != nil {
				return err
			}
			data2, err := code.ParseFile(path2)
			if err != nil {
				return err
			}

			result := code.GenDiff(data1, data2)
			fmt.Println(result)

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
