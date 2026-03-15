package affinecipher

import (
	"errors"
	"strings"
)

const (
	letters = 26
)

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func runeOp(r rune, op func(int) int) rune {
	if r >= '0' && r <= '9' {
		return r
	}
	v := op(int(r - 'a'))
	for v < 0 {
		v += letters
	}
	return rune(v%letters) + 'a'
}

func Encode(text string, a, b int) (string, error) {
	if GCD(a, letters) != 1 {
		return "", errors.New("a and m must be coprime")
	}
	var written int
	var sb strings.Builder
	for _, r := range []rune(strings.ToLower(text)) {
		if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'z') {
			continue
		}
		sb.WriteRune(runeOp(r, func(c int) int { return a*c + b }))
		written++
		if written%5 == 0 {
			sb.WriteRune(' ')
		}
	}
	got := sb.String()
	if written%5 == 0 {
		got = got[:len(got)-1]
	}
	return got, nil
}

func mmi(a int) int {
	for i := 1; i < letters; i++ {
		if a*i%letters == 1 {
			return i
		}
	}
	panic("")
}

func Decode(text string, a, b int) (string, error) {
	if GCD(a, letters) != 1 {
		return "", errors.New("a and m must be coprime")
	}
	m := mmi(a)
	var sb strings.Builder
	for _, r := range []rune(strings.ToLower(text)) {
		if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'z') {
			continue
		}
		sb.WriteRune(runeOp(r, func(c int) int { return m * (c - b) }))
	}
	return sb.String(), nil
}
