package main

import (
	"math"
	"fmt"
)

type (
	Shapes interface {
		getArea() float64
		getPerimeter() float64
	}
	Rectangle struct {
		height, lenght int
		Type           string
	}
	Circle struct {
		radius float64
		Type   string
	}
	Triangle struct {
		sideA float64
		sideB float64
		sideC float64
		Type string
	}
)

func main() {
	rectangle := Rectangle{3, 4, "Rectangle"}
	circle := Circle{5, "CIRCLE"}
	triangle := Triangle{3, 4, 5, "Triangle"}

	shapes := []Shapes{
		&rectangle,
		&circle,
		&triangle,
	}

	for _, shape := range shapes {
		describe(shape)
	}
}

func (r *Rectangle) getPerimeter() float64 {
	return float64((r.height + r.lenght) * 2)
}

func (r *Rectangle) getArea() float64 {
	return float64(r.height * r.lenght)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("%v (h = %v, l = %v)", r.Type, r.height, r.lenght)
}

func (c *Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) String() string {
	return fmt.Sprintf("%v (r = %v)", c.Type, c.radius)
}

func (t *Triangle) getPerimeter() float64 {
	return t.sideA+t.sideB+t.sideC
}

func (t *Triangle) getArea() float64 {
	halfPerim := t.getPerimeter()/2
	res := math.Sqrt(halfPerim*(halfPerim-t.sideA)*(halfPerim-t.sideB)*(halfPerim-t.sideC))
	return res
}

func (t *Triangle) String() string {
	return fmt.Sprintf("%v (a = %v, b = %v, c = %v)", t.Type, t.sideA, t.sideB, t.sideC)
}

func describe(shape Shapes) {
	fmt.Printf("%v \n", shape)
	fmt.Printf("Perimeter = %.2f \n", shape.getPerimeter())
	fmt.Printf("Area = %.2f \n ********** \n", shape.getArea())
}
