package main

import (
	"fmt"
	"time"
)

func Sleepseconds1(d time.Duration) {
	<-time.After(d * time.Second)
}
func Sleepseconds2(d int) {
	start := time.Now()
	for time.Since(start) < time.Duration(d)*time.Second {
	}

}
func main() {
	fmt.Println("начало")
	Sleepseconds1(3)
	fmt.Println("1 все")
	fmt.Println("пошла вторая")
	Sleepseconds2(3)
	fmt.Println("2 все")
}
