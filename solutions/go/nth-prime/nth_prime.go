// Package prime
package prime

var primes []int

func isPrime(i int) bool {
	for _, n := range primes {
		if i % n == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes[n - 1], true
}
