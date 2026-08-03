package main

import (
	"fmt"
	"strings"
)

func unic(str string) bool {
	str = strings.ToUpper(str)
	m := make(map[string]int)
	for _, v := range str {
		m[string(v)] += 1
		if m[string(v)] >= 2 {
			return false
		}
	}
	return true
}

func main() {
	s := "asdfghjkl"
	c := "asdfghjkla"
	fmt.Println(unic(s))
	fmt.Println(unic(c))

}
