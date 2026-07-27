package main

func swp(a, b int) (int, int) {
	return b, a
}
func main() {
	a := 42
	b := 24
	a, b = b, a
	println(a, b)
	//если это числа
	a = a + b
	b = a - b
	a = a - b
	println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	println(a, b)
	a, b = swp(a, b)
	println(a, b)
}
