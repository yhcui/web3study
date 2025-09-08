package main

import "fmt"

func main() {
	r := Rectangle{height: 10, width: 10}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	c := Circle{radius: 2}
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}

type Shap interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float32
	width  float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) Area() float64 {
	return float64(r.width * r.height)
}
func (r Rectangle) Perimeter() float64 {
	return float64(2 * (r.width + r.height))
}

func (r Circle) Area() float64 {
	return float64(3.14 * r.radius * r.radius)
}
func (r Circle) Perimeter() float64 {
	return float64(2 * 3.14 * r.radius)
}
