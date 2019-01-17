package main

import (
	"math"
	"fmt"
)

func main() {
	rectangle := Rectangle{3, 4, "Rectangle"}
	circle := Circle{5, "CIRCLE"}

	shapes := []Shapes{
		&rectangle,
		&circle,
	}

	for _, shape := range shapes {
		describe(shape)
	}
}

type Shapes interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	height, lenght int
	Type           string
}

type Circle struct {
	radius float64
	Type   string
}

func (r *Rectangle) getPerimeter() float64 {
	return float64((r.height + r.lenght) * 2)
}

func (r *Rectangle) getArea() float64 {
	return float64(r.height * r.lenght)
}

func (rec Rectangle) String() string {
	return fmt.Sprintf("%v (h = %v, l = %v)", rec.Type, rec.height, rec.lenght)
}

func (c *Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (cir Circle) String() string {
	return fmt.Sprintf("%v (r = %v)", cir.Type, cir.radius)
}

func describe(shape Shapes) {
	fmt.Printf("%v \n", shape)
	fmt.Printf("Perimeter = %v \n", shape.getPerimeter())
	fmt.Printf("Area = %v \n ********** \n", shape.getArea())
}
