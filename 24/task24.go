package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *point {
	return &point{x, y}
}

func (p point) Distance(pnt *point) float64 {
	return math.Pow(math.Pow(p.x-pnt.x, 2)+math.Pow(p.y-pnt.y, 2), 0.5)
}
func main() {
	p := NewPoint(4.2, 2.4)
	po := NewPoint(6.7, 7.6)
	fmt.Println(p.Distance(po))
}
