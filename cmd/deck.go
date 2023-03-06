package cmd

import (
	"github.com/alehechka/poker/cmd/flags"
	"github.com/alehechka/poker/internal/games"
	"github.com/alehechka/poker/pkg/cards"
	"github.com/urfave/cli/v2"
)

var deckFlags = []cli.Flag{
	flags.LogLevelFlag,
	flags.LogFormatFlag,
}

func deck(ctx *cli.Context) (err error) {

	flags.PrepareLogger(ctx)

	games.Init(&games.Config{
		AppName: ctx.App.Name,
		Version: ctx.App.Version,
		Deck:    cards.StandardDeck(),
	}).Log()

	return nil
}

// DeckCommand will create and print out a deck.
var DeckCommand = &cli.Command{
	Name:   "deck",
	Usage:  "Will prepare a standard or configured deck and print them to the console",
	Action: deck,
	Flags:  deckFlags,
}
