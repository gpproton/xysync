package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	_ "github.com/kr/pretty"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"

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

	//var testcard []string
	//db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5432 user=postgres dbname=testDatabase password=")
	//if err != nil {
	//	panic(err)
	//}
	//db.LogMode(true)
	//// db.Raw("SELECT * FROM tc_positions where sync != 1").Find(&testcard)
	//
	//query := db.Raw("SELECT * FROM tc_positions where sync != 1").Find(&testcard)
	//m := query.Value
	//fmt.Printf("%v", m)
	//
	//fmt.Println("---- TEST OUTPUT ----")
	//pretty.Print(testcard)
	//defer db.Close()

}
