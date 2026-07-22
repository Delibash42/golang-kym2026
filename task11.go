package main

import "fmt"

func intersect(a, b []int) []int {
	c := []int{}
	for _, v := range a {
		for _, val := range b {
			if v == val {
				c = append(c, v)
			}
		}
	}
	return c
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	fmt.Println(intersect(a, b))
}
