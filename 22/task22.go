package main

import (
	"fmt"
	"math/big"
)

func operacii(a, b *big.Int, op string) *big.Int {
	result := new(big.Int)
	switch op {
	case "+":
		return result.Add(a, b)
	case "-":
		return result.Sub(a, b)
	case "*":
		return result.Mul(a, b)
	case "/":
		return result.Div(a, b)
	default:
		return nil
	}
}

func main() {
	fmt.Println(operacii(big.NewInt(42424242), big.NewInt(24242424), "+"))
	fmt.Println(operacii(big.NewInt(42424242), big.NewInt(24242424), "/"))
	fmt.Println(operacii(big.NewInt(42424242), big.NewInt(24242424), "-"))
	fmt.Println(operacii(big.NewInt(42424242), big.NewInt(24242424), "*"))
}
