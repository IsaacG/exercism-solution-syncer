package alphametics

import (
	"errors"
	"strings"
)

type Puzzle struct {
	lhs     map[rune]int
	rhs     map[rune]int
	chars   []rune
	nonZero map[rune]bool
	used    map[int]bool
}

func charWeight(words []string) map[rune]int {
	w := map[rune]int{}
	for _, word := range words {
		runes := []rune(word)
		base := 1
		for i := len(runes) - 1; i >= 0; i-- {
			w[runes[i]] += base
			base *= 10
		}
	}
	return w
}

func newPuzzle(input string) (*Puzzle, error) {
	// Split into LHS and RHS words.
	parts := strings.Split(input, "==")
	if len(parts) != 2 {
		return nil, errors.New("puzzle must have exactly one '=='")
	}
	words := make([]string, 0)
	for _, word := range strings.Split(parts[0], "+") {
		words = append(words, strings.Trim(word, " "))
	}
	words = append(words, strings.Trim(parts[1], " "))

	// Compute chars and nonZero chars.
	nonZero := make(map[rune]bool, len(words))
	charMap := make(map[rune]struct{}, len(input))
	for _, word := range words {
		nonZero[rune(word[0])] = true
		for _, c := range word {
			charMap[c] = struct{}{}
		}
	}

	chars := make([]rune, 0, len(charMap))
	for c := range charMap {
		chars = append(chars, c)
	}

	return &Puzzle{
		lhs:     charWeight(words[:len(words)-1]),
		rhs:     charWeight(words[len(words)-1:]),
		chars:   chars,
		nonZero: nonZero,
		used:    map[int]bool{},
	}, nil
}

func (p Puzzle) recursiveSolve(balance int, chars []rune) (map[string]int, bool) {
	c, rest := chars[0], chars[1:]
	weight := p.lhs[c] - p.rhs[c]
	start := 0
	if p.nonZero[c] {
		start = 1
	}
	for i := start; i < 10; i++ {
		if p.used[i] {
			continue
		}

		p.used[i] = true
		if len(rest) == 0 {
			if balance+weight*i == 0 {
				return map[string]int{string(c): i}, true
			}
		} else {
			vals, ok := p.recursiveSolve(balance+weight*i, rest)
			if ok {
				vals[string(c)] = i
				return vals, true
			}
		}
		p.used[i] = false
	}
	return nil, false
}

func Solve(puzzle string) (map[string]int, error) {
	p, err := newPuzzle(puzzle)
	if err != nil {
		return nil, err
	}
	solution, ok := p.recursiveSolve(0, p.chars)
	if !ok {
		return nil, errors.New("no solution found")
	}
	return solution, nil
}
