package main

import "fmt"

func reverse(str string) string {
	stri := ""
	for _, v := range str {
		stri = string(v) + stri
	}
	return stri
}

func main() {
	fmt.Println(reverse("главрыба"))
}
