package games

import "github.com/sirupsen/logrus"

// ShuffleAndLogDeck starts the shuffle application
func (p *PokerHandler) ShuffleAndLogDeck() error {
	logrus.Debug("shuffling...")

	p.ShuffleDeck()

	p.LogDeck()
	return nil
}
