package poker

import (
	"cmp"
	"errors"
	"slices"
	"strconv"
	"strings"
)

type Rank int

// Ranks of hands.
const (
	royalFlush Rank = iota
	straightFlush
	fourOfAKind
	fullHouse
	flush
	straight
	threeOfAKind
	twoPair
	onePair
	highCard
)

var faceCards = map[string]int{
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

// Card represents a single card.
type Card struct {
	val  int
	suit rune
}

// Count represents a value and how many times it appears.
type Count struct {
	value      int
	occurances int
}

// Hand represents a hand of five cards.
type Hand struct {
	input string
	cards []Card
}

// HandCount represents the counts of all the cards in a hand.
type HandCount []Count

// Values returns the unique pip values in the count, ordered by occurances.
func (hc HandCount) Values() []int {
	var values []int
	for _, count := range hc {
		values = append(values, count.value)
	}
	return values
}

// Values returns the pip values of the cards in a hand.
func (h Hand) Values() []int {
	var out []int
	for _, c := range h.cards {
		out = append(out, c.val)
	}
	return out
}

// Count returns a slice of pairs containing the card value and how many times it appears.
func (h Hand) Count() HandCount {
	counts := map[int]int{}
	for _, c := range h.cards {
		counts[c.val]++
	}
	var out []Count
	for val, count := range counts {
		out = append(out, Count{value: val, occurances: count})
	}
	slices.SortFunc(out, func(a, b Count) int {
		if a.occurances != b.occurances {
			return -cmp.Compare(a.occurances, b.occurances)
		}
		return -cmp.Compare(a.value, b.value)
	})
	return out
}

func (h Hand) flush() bool {
	for _, c := range h.cards {
		if c.suit != h.cards[0].suit {
			return false
		}
	}
	return true
}

func (h Hand) royal() bool {
	return h.cards[0].val > 10
}

func (h Hand) straight() bool {
	vals := h.Values()
	// The first four cards must be in order.
	for i := range 3 {
		if vals[i] != vals[i+1]-1 {
			return false
		}
	}
	// The fifth card must be in order or the cards much be "2 ... A".
	return vals[3] == vals[4]-1 || (vals[0] == 2 && vals[4] == faceCards["A"])
}

// compare compares two hands and determines the better hand.
func (h Hand) compare(other Hand) int {
	rankA, valsA := h.Value()
	rankB, valsB := other.Value()
	// Compare ranks first.
	if rankA != rankB {
		return -cmp.Compare(rankA, rankB)
	}
	// For matching ranks, compare the card values.
	for i, v := range valsA {
		if v != valsB[i] {
			return cmp.Compare(v, valsB[i])
		}
	}
	return 0
}

// On a straight, we need special handling for "2 3 4 5 A" vs "2 3 4 5 6". Return the lowest "value".
func (h Hand) straightValues() []int {
	vals := h.Values()
	if vals[0] == 2 && vals[4] == 14 {
		return []int{1}
	}
	return vals[:1]
}

// Value returns the rank and relevant card values of a hand.
func (h Hand) Value() (Rank, []int) {
	count := h.Count()
	if h.royal() && h.flush() && h.straight() {
		return royalFlush, h.straightValues()
	}
	if h.flush() && h.straight() {
		return straightFlush, h.straightValues()
	}
	if count[0].occurances == 4 {
		return fourOfAKind, count.Values()
	}
	if count[0].occurances == 3 && count[1].occurances == 2 {
		return fullHouse, count.Values()
	}
	if h.flush() {
		return flush, count.Values()
	}
	if h.straight() {
		return straight, h.straightValues()
	}
	if count[0].occurances == 3 {
		return threeOfAKind, count.Values()
	}
	if count[0].occurances == 2 && count[1].occurances == 2 {
		return twoPair, count.Values()
	}
	if count[0].occurances == 2 {
		return onePair, count.Values()
	}
	return highCard, count.Values()
}

// NewHand returns a hand, validating the input.
func NewHand(s string) (Hand, error) {
	var cards []Card
	for c := range strings.FieldsSeq(s) {
		var face string
		var suit rune
		runes := []rune(c)
		for i, r := range runes {
			if r == '♤' || r == '♡' || r == '♢' || r == '♧' {
				if i != len(runes)-1 {
					return Hand{}, errors.New("invalid card")
				}
				suit = r
			} else {
				face += string(r)
			}
		}

		if suit == 0 {
			return Hand{}, errors.New("invalid suit")
		}
		v, ok := faceCards[face]
		if !ok {
			v, _ = strconv.Atoi(face)
			if v < 2 || v > 10 {
				return Hand{}, errors.New("invalid value")
			}
		}
		cards = append(cards, Card{v, suit})
	}
	if len(cards) != 5 {
		return Hand{}, errors.New("invalid card count")
	}
	slices.SortFunc(cards, func(a, b Card) int { return cmp.Compare(a.val, b.val) })
	return Hand{s, cards}, nil
}

// BestHand returns the best hands.
func BestHand(hands []string) ([]string, error) {
	var cardHands []Hand
	for _, h := range hands {
		hand, err := NewHand(h)
		if err != nil {
			return nil, err
		}
		cardHands = append(cardHands, hand)
	}
	ranked := slices.Clone(cardHands)
	slices.SortFunc(ranked, func(a, b Hand) int {
		return -a.compare(b)
	})

	var out []string
	for _, hand := range cardHands {
		if hand.compare(ranked[0]) == 0 {
			out = append(out, hand.input)
		}
	}
	return out, nil
}
