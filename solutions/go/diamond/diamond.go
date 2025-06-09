package diamond

import (
	"fmt"
	"strings"
)

// Gen returns a diamond pattern string.
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid char %b not in range [A, Z]", char)
	}

	first := int('A')
	last := int(char)
	length := last - first
	grid := make([]string, length*2+1)

	// Top row
	grid[0] = strings.Repeat(" ", length) + "A" + strings.Repeat(" ", length)
	// B through char.
	for i := 1; i <= length; i++ {
		edge := strings.Repeat(" ", length-i)
		center := strings.Repeat(" ", i*2-1)
		grid[i] = edge + string(byte(first+i)) + center + string(byte(first+i)) + edge
	}
	// Mirror the bottom half.
	for i := length; i >= 0; i-- {
		grid[2*length-i] = grid[i]
	}

	return strings.Join(grid, "\n"), nil
}
