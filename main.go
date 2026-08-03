package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifyslice(s)
	fmt.Println(s)
}
func modifyslice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
