package cards_test

import (
	"testing"

	"github.com/alehechka/poker/pkg/cards"
	"github.com/stretchr/testify/assert"
)

func Test_DealCard(t *testing.T) {
	deck := cards.StandardDeck()
	assert.Equal(t, 52, len(deck))

	deck, card := deck.DealCard(9)
	assert.Equal(t, 51, len(deck))
	assert.Equal(t, card, cards.TenOfHearts)

	deck, card = deck.DealCard()
	assert.Equal(t, cards.AceOfHearts, card)
	assert.Equal(t, 50, len(deck))
}
