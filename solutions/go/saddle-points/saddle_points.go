package matrix

import (
	"strconv"
	"strings"
)

// Pair stores a coordinate.
type Pair struct{ y, x int }

// Matrix is used to get a saddle point.
type Matrix struct {
	data          [][]int
	width, height int
}

// New returns a Matrix for a string.
func New(s string) (*Matrix, error) {
	var height, width int
	lines := strings.Split(s, "\n")
	height = len(lines)
	data := make([][]int, height)
	for idx, line := range lines {
		if line == "" {
			break
		}
		words := strings.Split(line, " ")
		width = len(words)
		data[idx] = make([]int, width)
		for wIdx, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			data[idx][wIdx] = num
		}
	}
	return &Matrix{data, width, height}, nil
}

func (m *Matrix) optimal(x, y int) bool {
	height := m.data[y][x]
	for otherX := range m.width {
		if m.data[y][otherX] > height {
			return false
		}
	}
	for otherY := range m.height {
		if m.data[otherY][x] < height {
			return false
		}
	}
	return true
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
