package __simple_shapes

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{3, 4, "Rectangle"}
	//fmt.Printf("Rectangle perimeter = %d \n", rect.perimeter())
	//fmt.Printf("Rectangle area = %d \n", rect.area())
	rect.say()

	circ := Circle{5, "CIRCLE"}
	fmt.Printf("\n Circle perimeter = %f \n", circ.perimeter())
	fmt.Printf("Circle area = %f \n", circ.area())
	circ.say()
}

//type Shape interface {
//	Area()
//}
//
//func (r *Rectangle) Area() {
//
//}

type Rectangle struct {
	height, lenght int
	Type           string
}
type Circle struct {
	radius float64
	Type   string
}

func (r *Rectangle) perimeter() int {
	return (r.height + r.lenght) * 2
}

func (r *Rectangle) area() int {
	return (r.height * r.lenght)
}

func (r *Rectangle) say() {

	fmt.Printf("Param of shape are:  'Type' = %s, 'height' = %v, 'length' = %v", r.Type, r.height, r.lenght)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) say() {

	fmt.Printf("Param of shape are:  'Type' = %s, 'redius' = %v", c.Type, c.radius)
}

func makeRectangle() {}
func makeTriangle()  {}
func makeCircle()    {}
