package blackjack

var cards = map[string]int{
	"ace":   11,
	"eight": 8,
	"two":   2,
	"nine":  9,
	"three": 3,
	"ten":   10,
	"four":  4,
	"jack":  10,
	"five":  5,
	"queen": 10,
	"six":   6,
	"king":  10,
	"seven": 7,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, exists := cards[card]
	if !exists {
		value = cards["other"]
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == "ace" && card1 == card2:
		return "P"
	case sum == 21 && !dealerViolenceBlackJack(dealerCard):
		return "W"
	case 12 <= sum && sum <= 16 && ParseCard(dealerCard) >= 7:
		return "H"
	case sum <= 11:
		return "H"
	}

	return "S"
}

func dealerViolenceBlackJack(card string) bool {
	switch card {
	case "jack", "queen", "king", "ten", "ace":
		return true
	}
	return false
}
