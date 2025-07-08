package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Number int
	Factorizations [][2]int
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	l := len(str) - 1
	for i := 0; i <= l / 2; i++ {
		if str[i] != str[l - i] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (Product, Product, error) {
	var low, high Product
	if fmin > fmax {
		return low, high, errors.New("fmin > fmax")
	}
	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			product := a * b
			if isPalindrome(product) {
				if low.Number == 0 || product < low.Number {
					low = Product{product, [][2]int{{a, b}}}
				} else if product == low.Number {
					low.Factorizations = append(low.Factorizations, [2]int{a, b})
				}
				if product > high.Number {
					high = Product{product, [][2]int{{a, b}}}
				} else if product == high.Number {
					high.Factorizations = append(high.Factorizations, [2]int{a, b})
				}
			}
		}
	}
	if low.Number == 0 {
		return low, high, errors.New("no palindromes")
	}
	return low, high, nil
}
