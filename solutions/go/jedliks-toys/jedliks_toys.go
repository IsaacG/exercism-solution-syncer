package jedlik

import "fmt"

// Drive drives the car one time.
func (c *Car) Drive() {
	if c.batteryDrain > c.battery {
		return
	}
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

// DisplayDistance displays the distance the car is driven.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery displays the battery level.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (c *Car) CanFinish(trackDistance int) bool {
	drives := int(c.battery / c.batteryDrain)
	distance := drives * c.speed
	return distance >= trackDistance
}
