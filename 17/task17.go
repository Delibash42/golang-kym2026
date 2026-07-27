package main

import "fmt"

func binpoisk(arr []int, cel int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	cntr := arr[mid]
	switch {
	case cntr == cel:
		return mid
	case cntr > cel:
		return binpoisk(arr[:mid], cel)
	case cntr < cel:
		res := binpoisk(arr[mid+1:], cel)
		if res == -1 {
			return -1
		}
		return mid + 1 + res
	default:
		return -1
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binpoisk(arr, 5))
	fmt.Println(binpoisk(arr, 1))
	fmt.Println(binpoisk(arr, 10))
	fmt.Println(binpoisk(arr, 99))
}
