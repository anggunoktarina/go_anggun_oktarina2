package main

import "fmt"

type Car struct {
	Type   string
	FuelIn int
}

func (car *Car) EstimateDestination() float64 {
	return float64(car.FuelIn) / 1.5
}

func main() {
	car1 := Car{
		Type:   "sedan",
		FuelIn: 15,
	}

	distance := car1.EstimateDestination()
	fmt.Printf("Jarak yang bisa ditempuh oleh mobil %v adalah %.2f km\n", car1.Type, distance)
}
