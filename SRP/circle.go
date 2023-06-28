package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() {
	fmt.Printf("circle area: %0.3f\n", math.Pi*c.radius*c.radius)
}
