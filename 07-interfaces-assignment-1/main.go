package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	side float64
}

func main() {

	t := triangle {
		height: 10,
		base: 10,
	}
	printArea(t)
	
	s := square {
		side: 4,
	}
	printArea(s)

}

func printArea(s shape) {
	fmt.Println("Area of shape: ", s, " is: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return  0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

