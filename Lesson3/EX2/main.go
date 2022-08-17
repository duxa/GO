package main

import (
	"fmt"
	"os"
)

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	if n == 1 || n == 0 {
		fmt.Println("Недопустимое число")
		os.Exit(1)
	}
	var isPrime bool = true
	for i := 2; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
			isPrime = true
		}

		if isPrime {
			fmt.Println(i)
		} else {
			continue
		}
	}
}
