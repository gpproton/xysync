package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}


	done := make(chan bool)

	func() {
		c := cron.New()
		c.AddFunc("@every 5s", func() {fmt.Println("First", 1) })
		c.AddFunc("@every 2s", func() {fmt.Println("First", 2) })
		c.Start()
		done <- true
	}()


}


