package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func main() {
	animal := Animal{
		nome:    "Gato",
		especie: "Felino",
		idade:   8,
		som:     "Miau",
	}

	AnimalInfo(&animal)

	animal = MudarSom(animal, "AU")

	AnimalInfo(&animal)
}

func MudarSom(animal Animal, novoSom string) Animal {
	animal.som = novoSom
	return animal
}

func AnimalInfo(f *Animal) {
	fmt.Printf("O animal %s com idade %d da espécie %s faz o som %s\n", f.nome, f.idade, f.especie, f.som)
}
