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
	var priorLen int
	// Work in reverse to handle length padding - each line must be as long as the next line.
	for i := longest - 1; i >= 0; i-- {
		var result strings.Builder
		for _, line := range input {
			if i < len(line) {
				result.WriteByte(line[i])
			} else {
				result.WriteRune(' ')
			}
		}
		// Strip trailing spaces then add back as much as is needed.
		padded := strings.TrimRight(result.String(), " ")
		for j := len(padded); j < priorLen; j++ {
			padded += " "
		}
		priorLen = len(padded)
		out[i] = padded
	}
	return out
}
