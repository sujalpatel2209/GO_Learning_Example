package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*func measure(g geometry)  {

	fmt.Println(g.area())
	fmt.Println(g.perim())
}*/

func main()  {

	r := rect{width:25.5, height:20.25}
	c := circle{radius:25.5}

/*	measure(r)
	measure(c)*/

	fmt.Println(r.area())
	fmt.Println(c.perim())

}