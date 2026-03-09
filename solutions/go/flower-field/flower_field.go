package flowerfield

import "strings"

// Annotate returns an annotated board
func Annotate(board []string) []string {
	flowers := make(map[[2]int]bool)
	for y, row := range board {
		for x, r := range row {
			if r == '*' {
				flowers[[2]int{x, y}] = true
			}
		}
	}

	getCount := func(x, y int) int {
		var c int
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if flowers[[2]int{x + dx, y + dy}] && (dx != 0 || dy != 0) {
					c++
				}
			}
		}
		return c
	}

	out := make([]string, len(board))
	for y := range len(board) {
		var sb strings.Builder
		for x := range len(board[0]) {
			if board[y][x] == '*' {
				sb.WriteRune('*')
			} else if c := getCount(x, y); c > 0 {
				sb.WriteRune(rune('0' + c))
			} else {
				sb.WriteRune(rune(' ' + c))
			}
		}
		out[y] = sb.String()
	}
	return out
}
