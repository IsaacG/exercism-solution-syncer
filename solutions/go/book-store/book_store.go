package bookstore

import (
	"slices"
)

// Cost for n books with the discounts applied.
var cost = map[int]int{
	1: 800,  // Full price
	2: 1520, // 5% off
	3: 2160, // 10% off
	4: 2560, // 20% off
	5: 3000, // 25% off
}

type cart []int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// count returns a slice how many of each title there is.
// count([1, 1, 1, 2, 3]) => 3, 1, 1  // 3x1 + 1x2 + 1x3
func count(vals []int) []int {
	counts := map[int]int{}
	for _, val := range vals {
		counts[val]++
	}
	var got []int
	for _, v := range counts {
		got = append(got, v)
	}
	slices.Sort(got)
	slices.Reverse(got)
	return got
}

// price sums up the price of each group with discounts applied.
func (c cart) price() int {
	var sum int
	for _, group := range c {
		sum += cost[group]
	}
	return sum
}

// min brute forces the minimum price.
func (c cart) min(toAdd int, minGroup int, additional []int) int {
	// Once all the books are in the cart, return the cart price.
	if toAdd == 0 && len(additional) == 0 {
		return c.price()
	}
	// If we are done with one title but have additional titles,
	// move on to the next title.
	if toAdd == 0 {
		toAdd, minGroup, additional = additional[0], 0, additional[1:]
	}
	// If the number of books matches the number of groups, each group gets a book.
	if toAdd == len(c)-minGroup {
		newCart := slices.Clone(c)
		for i := minGroup; i < len(c); i++ {
			newCart[i]++
		}
		return newCart.min(0, 0, additional)
	}
	// Take the min of with the book in the min group and without it.
	c[minGroup]++
	with := c.min(toAdd-1, minGroup+1, additional)
	c[minGroup]--
	without := c.min(toAdd, minGroup+1, additional)
	return min(with, without)
}

// Cost computes the minimum cost with all possible discount combos.
func Cost(books []int) int {
	counts := count(books)
	combo := cart{}
	// Distribute the most common book into one-per-group.
	if len(counts) > 0 {
		for range counts[0] {
			combo = append(combo, 1)
		}
	}
	// Distribute the second most common book into the groups.
	// The other here doesn't matter.
	if len(counts) > 1 {
		for i := range counts[1] {
			combo[i]++
		}
	}
	if len(counts) > 2 {
		return combo.min(0, 0, counts[2:])
	}
	return combo.price()
}
