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
		var lastChar int
		for j, inputLine := range input {
			if i < len(inputLine) {
				transposed[j] = inputLine[i]
				lastChar = j
			} else {
				transposed[j] = ' '
			}
		}

		// Determine the length of the transposed line.
		// It must be the max of the prev line and the last char.
		if prevLen < lastChar {
			prevLen = lastChar
		}
		out[i] = string(transposed[:prevLen+1])
	}
	return out
}
