package elon

import "fmt"

// Drive method drives the car if the car's battery is greater
// than its batteryDrain. no return
func (car *Car) Drive() {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance method returns the car's driven distance as a string
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery method returns the current car's battery percentage as a string
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%",car.battery)
}

// CanFinish method return if the car can complete the race.
func (car Car) CanFinish(trackDistance int) bool {
	var maxDistance int =  car.speed * car.battery / car.batteryDrain

	return maxDistance >= trackDistance
}
