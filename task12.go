package main

import "fmt"

func unic(arr []string) []string {
	m := make(map[string]int)
	for i, v := range arr {
		m[v] = i
	}
	un := []string{}
	for i, _ := range m {
		un = append(un, i)
	}
	return un
}
func main() {
	arr := []string{"cat", "dog", "cat", "tree"}
	fmt.Println(unic(arr))
}
