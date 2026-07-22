package main

import (
	"fmt"
)

func rplc(num int, i int, bit int) int {
	if bit == 1 {
		return num | (1<<i - 1)
	} else {
		return num &^ (1<<i - 1)
	}
}
func main() {
	fmt.Println(rplc(5, 1, 0))
}
