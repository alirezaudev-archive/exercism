package poker

import (
	"errors"
	"sort"
	"strings"
)

var ranks = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"10": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

type hand struct {
	original string
	cards    []int
	score    int
}

func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return nil, errors.New("no hands")
	}

	parsed := make([]hand, len(hands))
	for i, h := range hands {
		ph, err := parseHand(h)
		if err != nil {
			return nil, err
		}
		parsed[i] = ph
	}

	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i].score > parsed[j].score
	})

	best := parsed[0].score
	var winners []string
	for _, h := range parsed {
		if h.score == best {
			winners = append(winners, h.original)
		}
	}

	return winners, nil
}

func parseHand(s string) (hand, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 5 {
		return hand{}, errors.New("invalid hand")
	}

	cards := make([]int, 5)
	cardSuits := make([]rune, 5)

	for i, p := range parts {
		runes := []rune(p)
		if len(runes) < 2 {
			return hand{}, errors.New("invalid card")
		}

		suit := runes[len(runes)-1]
		if suit != '♤' && suit != '♡' && suit != '♢' && suit != '♧' {
			return hand{}, errors.New("invalid suit")
		}

		rankStr := string(runes[:len(runes)-1])
		rank, ok := ranks[rankStr]
		if !ok {
			return hand{}, errors.New("invalid rank")
		}

		cards[i] = rank
		cardSuits[i] = suit
	}

	sort.Ints(cards)

	counts := make([]int, 15)
	for _, c := range cards {
		counts[c]++
	}

	multiples := make([]int, 6)
	for _, count := range counts {
		if count > 0 {
			multiples[count]++
		}
	}

	isFlush := cardSuits[0] == cardSuits[1] && cardSuits[1] == cardSuits[2] &&
		cardSuits[2] == cardSuits[3] && cardSuits[3] == cardSuits[4]

	isStraight := cards[0]+1 == cards[1] && cards[1]+1 == cards[2] &&
		cards[2]+1 == cards[3] && cards[3]+1 == cards[4]

	isLowStraight := cards[0] == 2 && cards[1] == 3 && cards[2] == 4 &&
		cards[3] == 5 && cards[4] == 14

	allCards := cards[0] + cards[1]<<4 + cards[2]<<8 + cards[3]<<12 + cards[4]<<16

	score := 0
	switch {
	case isStraight && isFlush && isLowStraight:
		score = 8<<24 + 5<<20
	case isStraight && isFlush:
		score = 8<<24 + cards[4]<<20
	case multiples[4] == 1:
		score = 7<<24 + cards[2]<<20 + allCards
	case multiples[3] == 1 && multiples[2] == 1:
		score = 6<<24 + cards[2]<<20 + allCards
	case isFlush:
		score = 5<<24 + allCards
	case isLowStraight:
		score = 4<<24 + 5<<20
	case isStraight:
		score = 4<<24 + cards[4]<<20
	case multiples[3] == 1:
		score = 3<<24 + cards[2]<<20 + allCards
	case multiples[2] == 2:
		score = 2<<24 + allCards
	case multiples[2] == 1:
		score = 1<<24 + allCards
	default:
		score = allCards
	}

	return hand{original: s, cards: cards, score: score}, nil
}
