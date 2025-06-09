package diamond

import (
	"bytes"
	"fmt"
	"strings"
)
var Space = []byte{' '}

// Gen returns a diamond pattern string.
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid char %b not in range [A, Z]", char)
	}

	first := int('A')
	last := int(char)
	length := last - first

	var sb strings.Builder

	for i := 0; i <= length; i++ {
		sb.Write(bytes.Repeat(Space, length - i))
		sb.WriteByte(byte(first+i))
		if i > 0 {
			sb.Write(bytes.Repeat(Space, i*2-1))
			sb.WriteByte(byte(first+i))
		}
		sb.Write(bytes.Repeat(Space, length - i))
		if char != 'A' {
			sb.WriteByte('\n')
		}
	}
	// Mirror the bottom half.
	for i := length - 1; i >= 0; i-- {
		sb.Write(bytes.Repeat(Space, length - i))
		sb.WriteByte(byte(first+i))
		if i > 0 {
			sb.Write(bytes.Repeat(Space, i*2-1))
			sb.WriteByte(byte(first+i))
		}
		sb.Write(bytes.Repeat(Space, length - i))
		if i != 0 {
			sb.WriteByte('\n')
		}
	}

	return sb.String(), nil
}
