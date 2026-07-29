package main

import "fmt"

func reversewords(str string) string {
	st := []string{}
	temp := ""
	for _, v := range str {
		if string(v) != " " {
			temp = temp + string(v)
		} else {
			st = append(st, temp)
			temp = ""
		}
	}
	st = append(st, temp)
	result := ""
	for _, v := range st {
		result = v + " " + result
	}
	return result
}

func main() {
	fmt.Println(reversewords("go lang 2026 enstain"))
}
