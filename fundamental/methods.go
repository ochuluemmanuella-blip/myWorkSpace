package main
import (
    "fmt"
)
type Rectangle struct {
    Width  float64
    Height float64
}

// value receiver — works on a copy
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// pointer receiver — can mutate the original
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

var rect = Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // 50
rect.Scale(2)
fmt.Println(rect.Area())  // 200