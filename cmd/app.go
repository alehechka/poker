package cmd

import (
	"github.com/urfave/cli/v2"
)

// App represents the CLI application
func App(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "Poker Player 3000"
	app.Version = version
	app.Usage = "Will play different poker games depending on selected subcommand"
	app.Commands = []*cli.Command{
		ShuffleCommand,
		DeckCommand,
	}

	return app
}
