package spiralmatrix

// SpiralMatrix returns a spiral matrix of a given size.
func SpiralMatrix(size int) [][]int {
	// Create an empty matrix.
	matrix := make([][]int, size)
	for i := range size {
		matrix[i] = make([]int, size)
	}

	// Find the initial starting position, ie the "center" location with the highest number.
	posX, posY := -1, 0
	for i := range size {
		if i%2 == 1 {
			posY++
		} else {
			posX++
		}
	}

	// Start either going up or down.
	dX, dY := 0, 1
	if size%2 == 1 {
		dY = -1
	}

	for num := size * size; num > 0; num-- {
		// Record the number to the current position.
		matrix[posY][posX] = num
		if num == 1 {
			break
		}
		// Rotate left if the left position was not yet filled.
		leftDirX, leftDirY := dY, -dX
		if matrix[posY+leftDirY][posX+leftDirX] == 0 {
			dX, dY = leftDirX, leftDirY
		}
		// Advance one square.
		posX += dX
		posY += dY
	}

	return matrix
}
