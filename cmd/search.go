package main

import (
	"fmt"

	search "github.com/ryomak/search-scraping"
)

func main() {
	//  app := cli.NewApp()
	//
	//  app.Flags = []cli.Flag {
	//    cli.StringFlag{
	//      Name: "lang, l",
	//      Value: "english",
	//      Usage: "language for the greeting",
	//    },
	//  }
	//
	//  err := app.Run(os.Args)
	//  if err != nil {
	//    log.Fatal(err)
	//  }
	conf := search.LoadConfig("config.toml")
  fmt.Println(conf)
	fmt.Println(conf.AllSearch())
}
