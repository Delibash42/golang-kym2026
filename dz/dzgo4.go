package main

import "fmt"

func removeAt(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
func main() {
	masiv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	masiv = removeAt(masiv, 3)
	fmt.Println(masiv)
}
