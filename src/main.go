package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{"c"},
				Usage: "Load configuration from `FILE`",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}


}
