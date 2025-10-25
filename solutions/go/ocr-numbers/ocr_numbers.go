package ocr

/*
     _  _     _  _  _  _  _  _  #
   | _| _||_||_ |_   ||_||_|| | # decimal numbers.
   ||_  _|  | _||_|  ||_| _||_| #
                                # fourth line is always blank
*/

import (
	"fmt"
	"slices"
	"strings"
)

// Needed for the tests.
const recognizeDigit = 0

var digits = []string{
	" _ | ||_|   ", // 0
	"     |  |   ", // 1
	" _  _||_    ", // 2
	" _  _| _|   ", // 3
	"   |_|  |   ", // 4
	" _ |_  _|   ", // 5
	" _ |_ |_|   ", // 6
	" _   |  |   ", // 7
	" _ |_||_|   ", // 8
	" _ |_| _|   ", // 9
}

// Recognize runs OCR.
func Recognize(in string) []string {
	lines := strings.Split(strings.Trim(in, "\n"), "\n")
	var out []string
	// For each row of digits,
	for i := 0; i < len(lines); i += 4 {
		var line strings.Builder
		// For each letter block,
		for j := 0; j < len(lines[i]); j += 3 {
			block := ""
			// For each row of the digit,
			for k := i; k < i+4; k++ {
				block += lines[k][j : j+3]
			}
			if idx := slices.Index(digits, block); idx == -1 {
				line.WriteString("?")
			} else {
				line.WriteString(fmt.Sprintf("%d", idx))
			}
		}
		out = append(out, line.String())
	}
	return out
}
