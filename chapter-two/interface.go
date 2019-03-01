package main

import "fmt"
import "math"

type geometer interface {
	area() float64
	perim() float64
}

type rect struct {
	Name          string
	width, height float64
}
type circle struct {
	Name   string
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometer) {
	fmt.Println(g)
	fmt.Println("area", g.area())
	fmt.Println("perim", g.perim())
	fmt.Println()
}

func main() {
	r := rect{Name: "myRect", width: 3, height: 4}
	c := circle{Name: "myCircle", radius: 5}

	measure(r)
	measure(c)
}
