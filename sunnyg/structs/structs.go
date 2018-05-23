package structs

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() 		float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Perimeter() (float64) {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() (float64) {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() (float64) {
	a2b2 := math.Pow(t.Base / 2, 2) + math.Pow(t.Height, 2)
	hypot := math.Sqrt(a2b2)
	return t.Base + (2 * hypot)
}

func (c Circle) Area() (float64) {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() (float64) {
	return r.Width * r.Height
}

func (t Triangle) Area() (float64) {
	return (t.Base * t.Height) / 2
}
