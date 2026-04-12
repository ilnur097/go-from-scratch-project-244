package main

import (
	"log"
	"os"
        "code" 
        "fmt"
	"github.com/urfave/cli/v2"
)
func main() {
	app := &cli.App{
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
		Action: func(c *cli.Context) error {
			
			if c.Args().Len() < 2 {
				return fmt.Errorf("two file paths are required")
			}

			path1 := c.Args().Get(0)
			path2 := c.Args().Get(1)

			
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
