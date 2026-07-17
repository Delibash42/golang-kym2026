package main

import "fmt"

type User struct {
	id        string
	date      string
	balance   float64
	nickkname string
}

func swap(a, b *User) {
	*a, *b = *b, *a
}
func main() {
	a := User{"1234", "01.01.2026", 6.7, "delibash"}
	b := User{"56789", "02.02.2026", 7.6, "valentin"}
	fmt.Println(a.nickkname, b.nickkname)
	swap(&a, &b)
	fmt.Println(a.nickkname, b.nickkname)

}
