package main

import "fmt"

func tempsi(temps []float64) map[int][]float64 {
	result := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/10) * 10
		result[key] = append(result[key], temp)
	}

	return result
}

func main() {
	temperatures := []float64{-25.4, -21.0, 13.0, 19.0, 24.5, 32.5, -5.0, 0.0, 15.5, -30.0, 27.8}

	fmt.Println(temperatures)

	grouped := tempsi(temperatures)

	for key, temps := range grouped {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
