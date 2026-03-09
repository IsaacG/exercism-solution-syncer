package intergalactictransmission

import (
	"errors"
	"math/big"
	"math/bits"
	"slices"
)

// Encode the message.
func Transmit(message []byte) []byte {
	if len(message) == 0 {
		return message
	}

	// Convert the message to a number by shifting and adding 8 bits at a time.
	total := big.NewInt(0)
	for _, b := range message {
		total.Lsh(total, 8)
		total.Add(total, big.NewInt(int64(b)))
	}

	// Right pad (left shift) the message so it has n*7 bits.
	size := len(message) * 8
	shift := (7 - size%7) % 7
	total.Lsh(total, uint(shift))

	zero := big.NewInt(0)
	mod := big.NewInt(0)
	mask := big.NewInt(128)
	// Extract 7 bits at a time to add to the output by doing a shift-and-mod.
	var out []byte
	for total.Cmp(zero) != 0 {
		b := byte(mod.Mod(total, mask).Int64()) << 1
		b += byte(bits.OnesCount8(b) % 2)
		out = append(out, b)
		total.Rsh(total, 7)
	}
	slices.Reverse(out)
	// The message must be at least two bytes.
	for len(out) < 2 {
		out = append(out, 0)
	}

	return out
}

func Decode(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return message, nil
	}

	// Convert the message to an Int by doing repeated shift-and-add.
	total := big.NewInt(0)
	for _, b := range message {
		// Parity check.
		if bits.OnesCount8(b)%2 == 1 {
			return nil, errors.New("wrong parity")
		}
		// Shift and add.
		total.Lsh(total, 7)
		total.Add(total, big.NewInt(int64(b>>1)))
	}

	// Remove excess bits to make the message n*8 bits.
	dataBytes := len(message) * 7
	extra := dataBytes % 8
	total.Rsh(total, uint(extra))
	got := total.Bytes()
	// The message must be at least one byte.
	if len(got) == 0 {
		got = append(got, 0)
	}
	return got, nil
}
