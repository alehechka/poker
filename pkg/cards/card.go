package cards

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Card represents a single instance of a playing card
type Card struct {
	Name  Name
	Suit  Suit
	Value int
}

// FullName returns the full name of a card
func (c *Card) FullName() string {
	return fmt.Sprintf("%s of %s", c.Name, c.Suit)
}

// Log will log all important information about a given card
func (c *Card) Log() {
	logrus.
		WithField("value", c.Value).
		WithField("is_face", c.Name.IsFace()).
		WithField("is_ace", c.Name.IsAce()).
		WithField("color", c.Suit.Color()).
		Info(c.FullName())
}
