package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, 0)
	fmt.Println(a.Distance(*b))
}
