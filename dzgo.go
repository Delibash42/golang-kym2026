package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}
func (p Point) Distance(point Point) float64 {
	return math.Pow(math.Pow(p.x-point.x, 2)+math.Pow(p.y-point.y, 2), 0.5)
}
func main() {
	point1 := NewPoint(6.7, 6.7)
	point2 := NewPoint(7.6, 7.6)
	fmt.Println(point1.x)
	fmt.Println(point1.Distance(*point2))

}
