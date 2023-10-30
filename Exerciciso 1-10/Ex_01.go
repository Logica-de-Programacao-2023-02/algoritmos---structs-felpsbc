package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	c := Circulo{raio: 5.0}

	resultado1 := calculo1(c)

	fmt.Printf("A área do círculo com raio %.2f é %.2f \n", c.raio, resultado1)
}

func calculo1(c Circulo) float64 {
	return math.Pi * (c.raio * c.raio)
}
