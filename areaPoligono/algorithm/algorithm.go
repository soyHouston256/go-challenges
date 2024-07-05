package algorithm

import "fmt"

type Polygon interface {
	Area() float64
	PrintArea()
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) PrintArea() {
	fmt.Printf("The triangle area is: %f\n", t.Area())
}

type Square struct {
	Base float64
}

func (s Square) Area() float64 {
	return s.Base * s.Base
}

func (s Square) PrintArea() {
	fmt.Printf("The square ares is %f\n", s.Area())
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

func (r Rectangle) PrintArea() {
	fmt.Printf("The rectangle Area is %f\n", r.Area())
}

func Area(p Polygon) {
	p.PrintArea()
}
