package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var n string

	fmt.Print("Введите массив: ")
	fmt.Scanln(&n)

	sArr := strings.Split(n, ",") 
	
    ary := make([]int, len(sArr))
    for i := range ary {
        ary[i], _ = strconv.Atoi(sArr[i])
    }

	for i := 0; i < len(ary); i++ {
		temp := ary[i]
		for j := i-1; j>=0; j-- {
			if(ary[j] < temp){
				break;
			}
			ary[j + 1] = ary[j];
			ary[j] = temp;
		}		
	}
	fmt.Println(ary)
}
