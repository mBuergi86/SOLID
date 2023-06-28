package main

import "fmt"

type square struct {
	length float64
}

func (s square) area() {
	fmt.Printf("squaree area: %0.3f\n", s.length*s.length)
}
