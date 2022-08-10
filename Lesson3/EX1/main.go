package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, n!): ")
	fmt.Scanln(&op)

	if op != "n!" {
		fmt.Print("Введите второе число(аргумент): ")
		fmt.Scanln(&b)
	}

	switch op {
	case "+":
		res = a + b

	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на 0!")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "n!":
		res = float64(factorial(uint(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
