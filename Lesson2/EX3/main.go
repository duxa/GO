package main

import (
	"fmt"
)

func main() {
	var digit, hundreds, dozens, units int64
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&digit)
	if !(100 <= digit && digit <= 999) {
		fmt.Println("Не верное число!")
		return;
	}

	hundreds = digit / 100
	dozens = (digit - hundreds*100) / 10
	units = digit - hundreds*100 - dozens*10

	fmt.Printf("Сотни: %v, Десятки: %v Единицы: %v \n", hundreds, dozens, units)
}
