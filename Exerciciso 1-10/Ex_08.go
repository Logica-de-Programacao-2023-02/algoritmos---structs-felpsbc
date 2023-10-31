package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    int
	preco   float64
}

func main() {
	viagens := []Viagem{
		{origem: "Brasilia", destino: "Paris", data: 19042000, preco: 2},
		{origem: "Brasilia", destino: "Guarulhos", data: 04102007, preco: 7000},
		{origem: "Brasilia", destino: "California", data: 30072004, preco: 8000},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Printf("A viagem mais cara é de %s para %s, com preço de R$%.2f\n", viagemMaisCara.origem, viagemMaisCara.destino, viagemMaisCara.preco)
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}
