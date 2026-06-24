package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	radius float64
	pi     float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return c.pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * c.pi * c.radius
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
func Describe(s Shape) string {
	return fmt.Sprintf("Area: %.2f, Perimeter: %.2f", s.Area(), s.Perimeter())

}
