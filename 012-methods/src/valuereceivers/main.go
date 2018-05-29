package main

import "fmt"

type Car struct {
	gasPedal      uint16 // 0 to 65535
	brakePedal    uint16
	steeringWheel int16 // -32K +32K
	topSpeedKmph  float64
}

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60935

func (c Car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmph / usixteenbitmax)
}

func main() {
	car1 := Car{
		gasPedal:      23456,
		brakePedal:    0,
		steeringWheel: -12345,
		topSpeedKmph:  225.0,
	}

	fmt.Println(car1.kmh())

}
