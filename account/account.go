package account

import (
	"github.com/Golem-Base/golembase-demo-cli/account/balance"
	"github.com/Golem-Base/golembase-demo-cli/account/create"
	"github.com/Golem-Base/golembase-demo-cli/account/importkey"
	"github.com/urfave/cli/v2"
)

func Account() *cli.Command {
	return &cli.Command{
		Name:  "account",
		Usage: "Manage accounts",
		Subcommands: []*cli.Command{
			create.Create(),
			balance.AccountBalance(),
			importkey.ImportAccount(),
		},
	}
}
