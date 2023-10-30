package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func calculo(triangulo2 Triangulo) float64 {
	return (triangulo2.base * triangulo2.altura) / 2
}

func main() {

	triangulo2 := Triangulo{
		base:   5.0,
		altura: 2.0,
	}

	area := calculo(triangulo2)
	fmt.Printf("A area do Triangulo é %.2f\n", area)
}
