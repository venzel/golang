package interfaces

import (
	"fmt"
	"math"
)

type retangle struct {
	base float64
	height float64
}

func (r retangle) area() float64 {
	return r.base * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type model interface {
	area() float64
}

func calculate_area(m model) {
	fmt.Printf("A area eh %0.2f\n", m.area())
}

func generic(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func Execute(print bool) {

	if print {
		r := retangle{10, 20}
		calculate_area(r)

		c := circle{2}
		calculate_area(c)

		generic(c, r)
	}
}