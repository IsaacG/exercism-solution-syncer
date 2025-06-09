// Package luhn checks if a number satisfies the Luhn requirements.
package luhn

import (
	"strconv"
	"strings"
)

// Valid Luhn number or not?
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}
	a, b := sumA(s), sumB(s)
	if a != b {
		panic("Error")
	}
	return a%10 == 0
}

func sumA(s string) int {
	var sum int
	if len(s)%2 == 1 {
		s = "0" + s
	}
	for i, c := range s {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return 1
		}
		if i%2 == 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
	}
	return sum
}

func sumB(s string) int {
	var sum int
	flip := len(s)%2 == 1
	for _, c := range s {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return 1
		}
		if flip = !flip; flip {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
	}
	return sum
}
