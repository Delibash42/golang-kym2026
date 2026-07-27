package main

import "fmt"

func quicksort(arr []int) []int {
	switch len(arr) {
	case 0:
		return []int{}
	case 1:
		return arr
	case 2:
		if arr[0] < arr[1] {
			return []int{arr[0], arr[1]}
		}
		return []int{arr[1], arr[0]}
	default:
		cntr := arr[int(len(arr)/2)]
		var left, right, equal []int
		for _, v := range arr {
			if v > cntr {
				right = append(right, v)
			} else if v < cntr {
				left = append(left, v)
			} else {
				equal = append(equal, v)
			}
		}
		result := append(quicksort(left), equal...)
		result = append(result, quicksort(right)...)
		return result
	}
}

func main() {
	fmt.Println(quicksort([]int{3, 6, 8, 10, 1, 2, 1}))
}
