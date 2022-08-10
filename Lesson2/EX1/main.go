package main

import "fmt"

func main() {
	var a, b int32
	fmt.Print("Введите сторону пямоугольника: ")
	fmt.Scanln(&a)
	fmt.Print("Введите вторую сторону прямоугольника: ")
	fmt.Scanln(&b)
	fmt.Printf("Площадь: %v\n", a*b)
}
