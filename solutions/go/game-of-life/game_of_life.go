package gameoflife

func Tick(matrix [][]int) [][]int {
	alive := make(map[[2]int]bool)
	for y, row := range matrix {
		for x, i := range row {
			if i == 1 {
				alive[[2]int{x, y}] = true
			}
		}
	}

	getCount := func(x, y int) int {
		var c int
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if alive[[2]int{x + dx, y + dy}] && (dx != 0 || dy != 0) {
					c++
				}
			}
		}
		return c
	}

	out := make([][]int, len(matrix))
	for y := range len(matrix) {
		out[y] = make([]int, len(matrix[0]))
		for x := range len(matrix[0]) {
			c := getCount(x, y)
			if c == 3 || c+matrix[y][x] == 3 {
				out[y][x] = 1
			} else {
				out[y][x] = 0
			}
		}
	}
	return out
}
