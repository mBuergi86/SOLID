package main

type Square struct {
	length float64
}

func (s Square) area() float64 {
	return s.length * s.length
}
