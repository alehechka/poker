package cards

import (
	"math/rand"
	"time"
)

// Deck represents a deck of cards
type Deck []Card

// StandardDeck returns a standard deck in a pre-sorted order as a fresh deck would come
func StandardDeck() Deck {
	return Deck{
		AceOfHearts, TwoOfHearts, ThreeOfHearts, FourOfHearts, FiveOfHearts, SixOfHearts, SevenOfHearts, EightOfHearts, NineOfHearts, TenOfHearts, JackOfHearts, QueenOfHearts, KingOfHearts,
		AceOfSpades, TwoOfSpades, ThreeOfSpades, FourOfSpades, FiveOfSpades, SixOfSpades, SevenOfSpades, EightOfSpades, NineOfSpades, TenOfSpades, JackOfSpades, QueenOfSpades, KingOfSpades,
		AceOfDiamonds, TwoOfDiamonds, ThreeOfDiamonds, FourOfDiamonds, FiveOfDiamonds, SixOfDiamonds, SevenOfDiamonds, EightOfDiamonds, NineOfDiamonds, TenOfDiamonds, JackOfDiamonds, QueenOfDiamonds, KingOfDiamonds,
		AceOfClubs, TwoOfClubs, ThreeOfClubs, FourOfClubs, FiveOfClubs, SixOfClubs, SevenOfClubs, EightOfClubs, NineOfClubs, TenOfClubs, JackOfClubs, QueenOfClubs, KingOfClubs,
	}
}

func (d Deck) Log() {
	for _, card := range d {
		card.Log()
	}
}

func (d Deck) Shuffle() Deck {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	newDeck := make(Deck, 0)
	var card Card
	for len(d) > 0 {
		d, card = d.DealCard(r1.Intn(len(d)))
		newDeck = append(newDeck, card)
	}

	return newDeck
}

// DealCard by default will deal the "top" card of deck or the card at index 0. The card will be removed from the deck and returned from the function.
// If an index is provided, only the first index will be considered, and the card at that index will be removed from the deck and returned from the function.
func (d Deck) DealCard(indexes ...int) (Deck, Card) {
	index := 0

	if len(indexes) > 0 {
		index = indexes[0]
	}

	card := d[index]
	newDeck := append(d[:index], d[index+1:]...)

	return newDeck, card
}
