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
		lineA float64
		lineB float64
		lineC float64
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

func (rec *Rectangle) String() string {
	return fmt.Sprintf("%v (h = %v, l = %v)", rec.Type, rec.height, rec.lenght)
}

func (c *Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (cir *Circle) String() string {
	return fmt.Sprintf("%v (r = %v)", cir.Type, cir.radius)
}

func (t *Triangle) getPerimeter() float64 {
	return t.lineA+t.lineB+t.lineC
}

func (t *Triangle) getArea() float64 {
	halfPerim := t.getPerimeter()
	res := math.Sqrt(halfPerim*(halfPerim-t.lineA)*(halfPerim-t.lineB)*(halfPerim-t.lineC))
	return res
}

func (t *Triangle) String() string {
	return fmt.Sprintf("%v (a = %v, b = %v, c = %v)", t.Type, t.lineA, t.lineB, t.lineC)
}

func describe(shape Shapes) {
	fmt.Printf("%v \n", shape)
	fmt.Printf("Perimeter = %.2f \n", shape.getPerimeter())
	fmt.Printf("Area = %.2f \n ********** \n", shape.getArea())
}
