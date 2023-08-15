package interfacelearn

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
}

type Shape interface{
	Area() float64
	Perimeter() float64
}

// SHAPE METHODS
// Circle Methods
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//  Rectangle Methods
func (r Rectangle) Area() float64 {
	return r.Length * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 4 * r.Length
}

type ShapeInfo struct {
	area float64
	perimeter float64
}

func getShapeInfo(s Shape) ShapeInfo {
	println("Area", s.Area())
	println("Perimeter", s.Perimeter())
	return ShapeInfo{s.Area(), s.Perimeter()}
}