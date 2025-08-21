package main

import "fmt"

type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}

func (c *circle) area() float64 {
	return (3.14 * c.radius * c.radius)

}

func (r *rectangle) area() float64 {
	return (r.width * r.height)
}

type shape interface {
	area() float64
}

func main() {

	c := &circle{radius: 5}
	r := &rectangle{width: 4, height: 6}

	shapes := []shape{c, r}
	for _, s := range shapes {
		fmt.Println("Area:", s.area())
	}

}
