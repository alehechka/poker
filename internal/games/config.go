package games

import "github.com/alehechka/poker/pkg/cards"

// Config represents the initial configuration options for poker games
type Config struct {
	AppName string
	Version string

	Deck cards.Deck
}

// PokerHandler represents the playable handler for poker games
type PokerHandler struct {
	config *Config
}

func Init(config *Config) *PokerHandler {
	return &PokerHandler{
		config: config,
	}
}

func (p *PokerHandler) LogDeck() {
	p.config.Deck.Log()
}

func (p *PokerHandler) ShuffleDeck() {
	shuffled := p.config.Deck.Shuffle()
	p.config.Deck = shuffled
}

func (p *PokerHandler) DealCard(indexes ...int) cards.Card {
	dealt, card := p.config.Deck.DealCard(indexes...)
	p.config.Deck = dealt
	return card
}
