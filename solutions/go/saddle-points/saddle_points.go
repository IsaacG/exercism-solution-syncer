package matrix

import (
	"strconv"
	"strings"
)

// Pair stores a coordinate.
type Pair struct{ y, x int }

// Matrix is used to get a saddle point.
type Matrix struct {
	data           [][]int
	width, height  int
	rowMax, colMin []int
}

// New returns a Matrix for a string.
func New(s string) (*Matrix, error) {
	var height, width int
	lines := strings.Split(s, "\n")
	height = len(lines)
	data := make([][]int, height)
	rowMax := make([]int, height)
	for idx, line := range lines {
		if line == "" {
			break
		}
		words := strings.Fields(line)
		width = len(words)
		max := 0
		data[idx] = make([]int, width)
		for wIdx, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			data[idx][wIdx] = num
			if num > max {
				max = num
			}
		}
		rowMax[idx] = max
	}
	columns := 0
	if len(data) != 0 {
		columns = len(data[0])
	}
	colMin := make([]int, columns)
	for col := range columns {
		min := 0
		for row := range data {
			if row == 0 || data[row][col] < min {
				min = data[row][col]
			}
		}
		colMin[col] = min
	}

	return &Matrix{data, width, height, rowMax, colMin}, nil
}

func (m *Matrix) optimal(x, y int) bool {
	height := m.data[y][x]
	return height == m.rowMax[y] && height == m.colMin[x]
}

// Saddle returns saddle points for a Matrix.
func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for y := range m.height {
		for x := range m.width {
			if m.optimal(x, y) {
				pairs = append(pairs, Pair{y + 1, x + 1})
			}
		}
	}
	return pairs
}
