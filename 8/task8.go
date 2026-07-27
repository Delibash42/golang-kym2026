package main

import (
	"fmt"
	"math"
)

func rplc(exp int, j int, b string) int {
	var dv string
	for i := exp; i > 0; {
		dv = fmt.Sprint(i%2) + dv
		i = int(i / 2)
	}
	dv = dv[:j-1] + b + dv[j:]
	var ch int
	for i, v := range dv {
		ch += int(v-'0') * int(math.Pow(2, float64(len(dv)-1-i)))
	}
	return ch
}
func main() {
	fmt.Println(rplc(5, 1, "0"))
}
