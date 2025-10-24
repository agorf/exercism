package jedlik

import "fmt"

func (car *Car) Drive() {
	newBattery := car.battery - car.batteryDrain
	if newBattery < 0 {
		return
	}
	car.battery = newBattery
	car.distance += car.speed
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	newBattery := float64(car.battery) - (float64(trackDistance)/float64(car.speed))*float64(car.batteryDrain)
	return newBattery >= 0
}
