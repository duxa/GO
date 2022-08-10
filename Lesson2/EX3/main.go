package main

import (
	"fmt"
)

func main() {
	var number, hundreds, dozens, units int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&number)
	if !(100 <= number && number <= 999) {
		fmt.Println("Не верное число!")
		return
	}

	hundreds = number / 100
	dozens = (number - hundreds*100) / 10
	units = number - hundreds*100 - dozens*10

	fmt.Printf("Сотни: %d, Десятки: %d Единицы: %d \n", hundreds, dozens, units)
}
