package transpose

import (
	"slices"
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

	var out []string
	var priorLen int
	for idx := range longest {
		// Work in reverse to handle length padding - each line must be as long as the next line.
		idx = longest - idx - 1
		var result strings.Builder
		for _, line := range input {
			if idx < len(line) {
				result.WriteByte(line[idx])
			} else {
				result.WriteRune(' ')
			}
		}
		// Strip trailing spaces then add back as much as is needed.
		padded := strings.TrimRight(result.String(), " ")
		for i := len(padded); i < priorLen; i++ {
			padded += " "
		}
		priorLen = len(padded)
		out = append(out, padded)
	}
	slices.Reverse(out)
	return out
}
