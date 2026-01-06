// Package alphametics solves alphametics puzzles.
package alphametics

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// runeWeight returns a weight of each rune, used to evaluate equations.
// "ABA + BC" has weights {'A': 101, 'B': 20, 'C': 1}.
// By multiplying weights be rune value, equations can be evaluated quickly.
// For example, given values {'A': 2, 'B': 3, 'C': 4}, the above equation
// has a value of 2 * 101 + 3 * 20 + 4 * 1.
func runeWeight(words []string) map[rune]int {
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

// solver solves puzzles.
type solver struct {
	// Weight of runes for equation LHS - RHS = 0
	runeWeights map[rune]int
	// List of runes to solve.
	runes []rune
	// Map indicating which runes cannot be zero.
	nonZero map[rune]bool
}

// newSolver builds and returns a new solver.
func newSolver(input string) (*solver, error) {
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

	// Compute runes and nonZero runes.
	nonZero := make(map[rune]bool, len(words))
	runeMap := make(map[rune]struct{}, len(input))
	for _, word := range words {
		r, _ := utf8.DecodeRuneInString(word)
		nonZero[r] = true
		for _, r := range word {
			runeMap[r] = struct{}{}
		}
	}

	runes := make([]rune, 0, len(runeMap))
	for r := range runeMap {
		runes = append(runes, r)
	}

	weights := map[rune]int{}
	// LHS weights are positive.
	for r, w := range runeWeight(words[:len(words)-1]) {
		weights[r] += w
	}
	// RHS weights are negative..
	for r, w := range runeWeight(words[len(words)-1:]) {
		weights[r] -= w
	}
	return &solver{
		runes:       runes,
		nonZero:     nonZero,
		runeWeights: weights,
	}, nil
}

// recursiveSolve solves a puzzle recursively, one rune at a time.
// An equation is valid when LHS == RHS. Alternatively, LHS - RHS == 0.
// Equations can be solved recursively by replacing one rune at a time and tracking
// the computed "balance" of LHS - RHS. When all runes are evaluated, LHS - RHS = 0
// must be true.
func (p solver) recursiveSolve(balance int, runes []rune, used map[int]bool) (map[string]int, bool) {
	// Split the runes into the current rune to replace and the rest to be recursively handled.
	r, rest := runes[0], runes[1:]
	weight := p.runeWeights[r]
	// Try using values [0..9] or [1..9] as a value for r.
	start := 0
	if p.nonZero[r] {
		start = 1
	}
	for i := start; i < 10; i++ {
		if used[i] {
			continue
		}

		used[i] = true
		// If there are no remaining runes, check if the equation balances.
		if len(rest) == 0 {
			// If the equation balances, we found a solution.
			if balance+weight*i == 0 {
				return map[string]int{string(r): i}, true
			}
		} else {
			// Check if we can find a solution with this value for r.
			vals, ok := p.recursiveSolve(balance+weight*i, rest, used)
			if ok {
				vals[string(r)] = i
				return vals, true
			}
		}
		used[i] = false
	}
	return nil, false
}

// Solve solves an alphametics puzzle.
func Solve(puzzle string) (map[string]int, error) {
	p, err := newSolver(puzzle)
	if err != nil {
		return nil, err
	}
	solution, ok := p.recursiveSolve(0, p.runes, map[int]bool{})
	if !ok {
		return nil, errors.New("no solution found")
	}
	return solution, nil
}
