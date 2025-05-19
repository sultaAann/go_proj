package main

import (
	"fmt"
	"os"
)

type Shape interface {
	Perimeter() float64
	Square() float64
}

type Rectangle struct {
	a float64
	b float64
}

type Circle struct {
	r float64
}

func (r *Rectangle) Perimeter() float64 {
	return (r.a + r.b) * 2
}

func (r *Rectangle) Square() float64 {
	return r.a * r.b
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.r
}

func (c *Circle) Square() float64 {
	return 3.14 * c.r
}

func DetectType(s Shape) {
	if _, ok := s.(*Rectangle); ok {
		fmt.Println("Its rectangle!")
	} else if _, ok := s.(*Circle); ok {
		fmt.Println("Its Circle!")
	} else {
		fmt.Println("Unknown type!")
	}
}

func main() {
	r := Rectangle{4, 5}
	fmt.Println("Rectangle:", r)	

	fmt.Println(r.Perimeter())
	fmt.Println(r.Square())

	c := Circle{3}
	fmt.Println("Circle r:", c)

	fmt.Println(c.Perimeter())
	fmt.Println(c.Square())

	DetectType(&r)

	DetectType(&c)

	os.Exit(0)
}
