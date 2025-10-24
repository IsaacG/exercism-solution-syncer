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

func updateProduct(p Product, product, a, b int, compare func(int, int) bool) Product {
    if p.Number == 0 || compare(product, p.Number) {
        return Product{product, [][2]int{{a, b}}}
    }
    if product == p.Number {
        p.Factorizations = append(p.Factorizations, [2]int{a, b})
    }
    return p
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
				low = updateProduct(low, product, a, b, func(x, y int) bool { return x < y })
				high = updateProduct(high, product, a, b, func(x, y int) bool { return x > y })
			}
		}
	}
	if low.Number == 0 {
		return low, high, errors.New("no palindromes")
	}
	return low, high, nil
}
