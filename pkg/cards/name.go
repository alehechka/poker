package cards

import "fmt"

// Name represents the front showing Name of the card
type Name string

const (
	Ace   Name = "Ace"
	Two   Name = "Two"
	Three Name = "Three"
	Four  Name = "Four"
	Five  Name = "Five"
	Six   Name = "Six"
	Seven Name = "Seven"
	Eight Name = "Eight"
	Nine  Name = "Nine"
	Ten   Name = "Ten"
	Jack  Name = "Jack"
	Queen Name = "Queen"
	King  Name = "King"
)

// Plural returns the plural format of the given name
func (n Name) Plural() string {
	switch n {
	case Six:
		return "Sixes"
	default:
		return fmt.Sprintf("%ss", n)
	}
}

// IsFace returns the truthiness of the given name representing a face card
func (n Name) IsFace() bool {
	switch n {
	case Jack, Queen, King:
		return true
	default:
		return false
	}
}

// IsAce returns the truthiness of the given name representing an Ace
func (n Name) IsAce() bool {
	return n == Ace
}
