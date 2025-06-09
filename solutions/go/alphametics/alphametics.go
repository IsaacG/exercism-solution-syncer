package alphametics

import (
	"errors"
	"strings"
)

type Puzzle struct {
	lhs     []string
	rhs     string
	chars   []rune
	nonZero map[rune]bool
	maxVal  int
}

type values struct {
	m       map[rune]int
	chars   []rune
	nonZero map[rune]bool
}

func newValues(chars []rune, nonZero map[rune]bool) values {
	v := make(map[rune]int, len(chars))
	for _, c := range chars {
		v[c] = 0
	}
	return values{v, chars, nonZero}
}

func (v values) stringMap() map[string]int {
	sm := make(map[string]int, len(v.m))
	for r, i := range v.m {
		sm[string(r)] = i
	}
	return sm
}

func (v values) increment() {
	for _, c := range v.chars {
		v.m[c] += 1
		if v.m[c] != 10 {
			break
		}
		if v.nonZero[c] {
			v.m[c] = 1
		} else {
			v.m[c] = 0
		}
	}
}

func (v values) translate(s string) int {
	sum := 0
	for _, b := range s {
		sum *= 10
		sum += v.m[b]
	}
	return sum
}

func (p Puzzle) valid(v values) bool {
	// Disqualify repeating numbers.
	vals := make(map[int]struct{}, len(v.m))
	for _, val := range v.m {
		vals[val] = struct{}{}
	}
	if len(vals) != len(v.m) {
		return false
	}
	// Check the equallity.
	sum := 0
	for _, w := range p.lhs {
		sum += v.translate(w)
	}
	return sum == v.translate(p.rhs)
}

func newPuzzle(input string) (*Puzzle, error) {
	// Split into LHS and RHS words.
	parts := strings.Split(input, "==")
	if len(parts) != 2 {
		return nil, errors.New("puzzle must have exactly one '=='")
	}
	lhs := make([]string, 0)
	for _, word := range strings.Split(parts[0], "+") {
		lhs = append(lhs, strings.Trim(word, " "))
	}
	rhs := strings.Trim(parts[1], " ")

	// Collect all the words to compute chars and nonZero chars.
	words := make([]string, len(lhs), len(lhs)+1)
	copy(words, lhs)
	words = append(words, rhs)

	nonZero := make(map[rune]bool, len(lhs)+1)
	charMap := make(map[rune]struct{}, len(input))
	for _, word := range words {
		nonZero[rune(word[0])] = true
		for _, c := range word {
			charMap[c] = struct{}{}
		}
	}

	chars := make([]rune, 0, len(charMap))
	maxVal := 1
	for c := range charMap {
		chars = append(chars, c)
		maxVal *= 10
	}

	return &Puzzle{
		lhs:     lhs,
		rhs:     rhs,
		chars:   chars,
		nonZero: nonZero,
		maxVal:  maxVal,
	}, nil
}

func (p Puzzle) solve() (map[string]int, error) {
	v := newValues(p.chars, p.nonZero)
	for i := 0; i < p.maxVal; i++ {
		if p.valid(v) {
			println("Solved", i)
			return v.stringMap(), nil
		}
		v.increment()
	}
	return nil, errors.New("no solution found")
}

func Solve(puzzle string) (map[string]int, error) {
	p, err := newPuzzle(puzzle)
	if err != nil {
		return nil, err
	}
	return p.solve()

}
