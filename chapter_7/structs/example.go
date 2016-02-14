package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r

}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	total := float64(0)
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {

	c1 := Circle{x: 0, y: 0, r: 5}
	c2 := Circle{0, 0, 5}
	c3 := &Circle{0, 0, 5}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(circleArea(&c1))
	fmt.Println(circleArea(&c2))
	fmt.Println(circleArea(c3))
	fmt.Println(c1.area())
	fmt.Println(c3.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(totalArea(&c1, &r))
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println(multiShape.area())
	fmt.Println(multiShape.perimeter())
}
