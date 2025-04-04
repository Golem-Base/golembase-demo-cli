package main

import (
	"log"
	"os"

	"github.com/Golem-Base/golembase-demo-cli/account"
	"github.com/Golem-Base/golembase-demo-cli/cat"
	"github.com/Golem-Base/golembase-demo-cli/entity"
	"github.com/Golem-Base/golembase-demo-cli/query"
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
