package main

import (
	"fmt"
	"math"
)

func main() {
	var square, diameter, сircumference float64
	fmt.Print("Введите площаль круга: ")
	fmt.Scanln(&square)
	diameter = 2 * math.Sqrt(square/math.Pi)
	fmt.Printf("Диаметр: %f\n", diameter)
	сircumference = diameter * math.Pi
	fmt.Printf("Длина окружности: %f\n", сircumference)
}
