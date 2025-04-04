package main

import (
	"log"
	"os"

	"github.com/Golem-Base/golembase-demo/account"
	"github.com/Golem-Base/golembase-demo/cat"
	"github.com/Golem-Base/golembase-demo/entity"
	"github.com/Golem-Base/golembase-demo/query"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "golembase CLI",
		Usage: "Golem Base",

		Commands: []*cli.Command{
			account.Account(),
			entity.Entity(),
			cat.Cat(),
			query.Query(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
