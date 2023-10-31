package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(indice int) {
	if indice >= 0 && indice < len(a.notas) {
		a.notas = append(a.notas[:indice], a.notas[indice+1:]...)
	}
}

func (a *Aluno) calcularMedia() float64 {
	if len(a.notas) == 0 {
		return 0
	}

	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	return total / float64(len(a.notas))
}

func (a *Aluno) imprimirDetalhes() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Média das notas: %.2f\n", a.calcularMedia())
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
	}

	aluno.adicionarNota(8.5)
	aluno.adicionarNota(9.0)
	aluno.adicionarNota(7.5)

	aluno.imprimirDetalhes()

	aluno.removerNota(1)
	fmt.Println("Após remover a segunda nota:")
	aluno.imprimirDetalhes()
}
