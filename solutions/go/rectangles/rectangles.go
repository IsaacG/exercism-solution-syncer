package rectangles

type Point struct { x, y int }

func Count(diagram []string) int {
	corners := map[Point]bool{}
	horz := map[Point]bool{}
	vert := map[Point]bool{}

	for y, line := range diagram {
		for x, char := range line {
			if char == ' ' {
				continue
			}
			p := Point{x, y}
			switch char {
			case '+':
				corners[p] = true
				horz[p] = true
				vert[p] = true
			case '|':
				vert[p] = true
			case '-':
				horz[p] = true
			}
		}
	}

	count := 0
	// Iterate through all potential rectangle corners.
	for topLeft, _ := range corners {
		for bottomRight, _ := range corners {
			// topLeft must be above and left of bottomRight.
			if topLeft.x >= bottomRight.x || topLeft.y >= bottomRight.y {
				continue
			}
			// Check the other two corners exist.
			topRight := Point{bottomRight.x, topLeft.y}
			bottomLeft := Point{topLeft.x, bottomRight.y}
			if !corners[topRight] || !corners[bottomLeft] {
				continue
			}

			// Check the top and bottom edges exist.
			valid := true
			for x := topLeft.x; x <= topRight.x; x++ {
				if !horz[Point{x, topLeft.y}] || !horz[Point{x, bottomLeft.y}] {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			// Check the right and left edges exist.
			for y := topLeft.y; y <= bottomLeft.y; y++ {
				if !vert[Point{topLeft.x, y}] || !vert[Point{topRight.x, y}] {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			count++
		}
	}
	return count
	
	panic("Please implement the Count function")
}
