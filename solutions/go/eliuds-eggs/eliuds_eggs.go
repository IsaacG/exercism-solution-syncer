package eliudseggs

func EggCount(displayValue int) int {
	count := 0
	for i := displayValue; i > 0; i >>= 1 {
		if i&1 == 1 {
			count++
		}
	}
	return count
}
