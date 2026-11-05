// Package lasagna does baking time.
package lasagna

const (
    // OvenTime is total time in the oven.
    OvenTime = 40
	// TimePerLayer is the prep time per layer.
	TimePerLayer = 2
)

// RemainingOvenTime returns time left in the oven.
func RemainingOvenTime(elapsed int) int {
	return OvenTime - elapsed
}

// PreparationTime returns time needed to prep.
func PreparationTime(layers int) int {
	return layers * TimePerLayer
}

// ElapsedTime returns time already elapsed.
func ElapsedTime(layers, elapsed int) int {
	return PreparationTime(layers) + elapsed
}
