package transpose

import (
	"strings"
)

// Transpose a slice of strings.
func Transpose(input []string) []string {
	// Compute the longest line length.
	var longest int
	for _, line := range input {
		if len(line) > longest {
			longest = len(line)
		}
	}

	out := make([]string, longest)
	var prevLen int
	// Work in reverse to handle length padding - each line must be as long as the next line.
	for i := longest - 1; i >= 0; i-- {
		transposed := make([]byte, 0, longest)
		for _, inputLine := range input {
			if i < len(inputLine) {
				transposed = append(transposed, inputLine[i])
			} else {
				transposed = append(transposed, ' ')
			}
		}
		// Strip trailing spaces then add back as much as is needed.
		padded := strings.TrimRight(string(transposed), " ")
		for range prevLen - len(padded) {
			padded += " "
		}
		prevLen = len(padded)
		out[i] = padded
	}
	return out
}
