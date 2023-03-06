package cards

// Suit represents an enum for card suits
type Suit string

const (
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

// Color represents an enum for the color of suits
type Color string

const (
	Red     Color = "red"
	Black   Color = "black"
	Unknown Color = "unknown"
)

// Color returns the color of the suit
func (s Suit) Color() Color {
	switch s {
	case Hearts, Diamonds:
		return Red
	case Spades, Clubs:
		return Black
	default:
		return Unknown
	}
}
