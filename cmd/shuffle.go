package cmd

import (
	"github.com/alehechka/poker/cmd/flags"
	"github.com/alehechka/poker/internal/games"
	"github.com/alehechka/poker/pkg/cards"
	"github.com/urfave/cli/v2"
)

var shuffleFlags = []cli.Flag{
	flags.LogLevelFlag,
	flags.LogFormatFlag,
}

func shuffle(ctx *cli.Context) (err error) {

	flags.PrepareLogger(ctx)

	return games.Init(&games.Config{
		AppName: ctx.App.Name,
		Version: ctx.App.Version,
		Deck:    cards.StandardDeck(),
	}).GameShuffle()
}

// ShuffleCommand starts the application.
var ShuffleCommand = &cli.Command{
	Name:   "shuffle",
	Usage:  "Will prepare a standard or configured deck, shuffle the cards, and print them to the console",
	Action: shuffle,
	Flags:  shuffleFlags,
}
