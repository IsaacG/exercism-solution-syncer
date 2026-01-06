package transpose

// Transpose a slice of strings.
func Transpose(input []string) []string {
	// Compute the longest line length.
	var longest int
	inputLen := len(input)
	for _, line := range input {
		if len(line) > longest {
			longest = len(line)
		}
	}

	out := make([]string, longest)
	var prevLen int
	// Work in reverse to handle length padding - each line must be as long as the next line.
	for i := longest - 1; i >= 0; i-- {
		transposed := make([]byte, inputLen)
		for j, inputLine := range input {
			if i < len(inputLine) {
				transposed[j] = inputLine[i]
			} else {
				transposed[j] = ' '
			}
		}

		// Determine the length of the transposed line.
		// It must be at least as long as the prev line and long enough for all chars.
		transposedLen := prevLen
		for j := prevLen; j < inputLen; j++ {
			if transposed[j] != ' ' {
				transposedLen = j
			}
		}
		out[i] = string(transposed[:transposedLen+1])
		prevLen = transposedLen
	}
	return out
}
