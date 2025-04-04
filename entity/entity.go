package entity

import (
	"github.com/Golem-Base/golembase-demo-cli/entity/create"
	"github.com/Golem-Base/golembase-demo-cli/entity/delete"
	"github.com/Golem-Base/golembase-demo-cli/entity/update"
	"github.com/urfave/cli/v2"
)

func Entity() *cli.Command {
	return &cli.Command{
		Name:  "entity",
		Usage: "Manage entities",
		Subcommands: []*cli.Command{
			create.Create(),
			delete.Delete(),
			update.Update(),
		},
	}
}
