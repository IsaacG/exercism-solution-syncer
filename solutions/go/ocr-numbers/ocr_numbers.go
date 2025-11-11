package ocr

import (
	"fmt"
	"slices"
	"strings"
)

// Needed for the tests. Not used.
const recognizeDigit = 0

// digits contains the characters which make up a digit, with the rows concatenated.
// Populated by init().
var digits = []string{}

// Recognize runs OCR.
func Recognize(display string) []string {

	lines := strings.Split(strings.Trim(display, "\n"), "\n")
	var out []string
	var block strings.Builder
	// For each row of digits,
	for i := 0; i < len(lines); i += 4 {
		var line strings.Builder
		// For each letter block,
		for j := 0; j < len(lines[i]); j += 3 {
			block.Reset()
			// For each row of the digit,
			for k := i; k < i+4; k++ {
				// Concatenate the rows making up the digit.
				block.WriteString(lines[k][j : j+3])
			}
			// Output either the digit or "?" if it is not recognized.
			if idx := slices.Index(digits, block.String()); idx == -1 {
				line.WriteString("?")
			} else {
				line.WriteString(fmt.Sprintf("%d", idx))
			}
		}
		out = append(out, line.String())
	}
	return out
}

func init() {
	screen := `
    _  _     _  _  _  _  _  _  #
  | _| _||_||_ |_   ||_||_|| | # decimal numbers.
  ||_  _|  | _||_|  ||_| _||_| #
                               # fourth line is always blank
`
	var block strings.Builder
	lines := strings.Split(strings.Trim(screen, "\n"), "\n")
	for i := range 10 {
		i = (i + 9) % 10
		block.Reset()
		for j := 0; j < 4; j++ {
			block.WriteString(lines[j][i*3 : (i+1)*3])
		}
		digits = append(digits, block.String())
	}
}
