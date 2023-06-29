package main

import "fmt"

func main() {
	shapes := []Area{
		Circle{radius: 5.5},
		Square{length: 7.5},
	}

	for _, shape := range shapes {
		switch shape := shape.(type) {
		case Circle:
			fmt.Printf("circle area: %0.3f\n", shape.Area())
		case Square:
			fmt.Printf("square area: %0.3f\n", shape.Area())
		}
	}
}
