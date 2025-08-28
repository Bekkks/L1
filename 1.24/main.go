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

func (p *Point) Distance(other *Point) float64 {
	dis1 := other.x - p.x
	dis2 := other.y - p.y
	return math.Sqrt(dis1*dis1 + dis2*dis2)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(5, 6)
	res := p1.Distance(p2)
	fmt.Println(res)
}
