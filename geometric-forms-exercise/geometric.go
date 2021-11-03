package main

// Definir 3 structs
// Circulo, cuadrado y triangulo

import (
	"fmt"
	"math"
)

type Circle struct{ radio float64 }
type Square struct{ side float64 }
type Triangle struct{ base, height float64 }

func main() {
	c := Circle{2}
	s := Square{2}
	t := Triangle{2, 4}
	fmt.Printf("Area de circulo: %v", c.area())
	fmt.Printf("\nArea de cuadrado: %v", s.area())
	fmt.Printf("\nArea de triangulo: %v", t.area())

}

func (circle Circle) area() float64 {
	return math.Pow(circle.radio*math.Pi, 2)
}

func (square Square) area() float64 {
	return math.Pow(square.side, 2)
}

func (triangle Triangle) area() float64 {
	return (triangle.base * triangle.height) / 2
}
