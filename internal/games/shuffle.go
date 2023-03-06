package games

import "github.com/sirupsen/logrus"

// Shuffle starts the shuffle application
func (p *PokerHandler) GameShuffle() error {
	logrus.Debug("shuffling...")

	p.ShuffleDeck()

	p.config.Deck.Log()
	return nil
}
