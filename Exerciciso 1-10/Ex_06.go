package main

import "fmt"

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func (f *Funcionario) aumentarSalario(porcentagem float64) {
	f.salario *= (1 + porcentagem/100)
}
func (f *Funcionario) diminuirSalario(porcentagem float64) {
	f.salario *= (1 - porcentagem/100)
}
func (f *Funcionario) tempoServico() int {
	return f.idade - 18
}
func main() {
	felipe := Funcionario{
		nome:    "Felipe",
		salario: 1000,
		idade:   19,
	}
	felipe.aumentarSalario(20)

	felipe.diminuirSalario(20)

	fmt.Printf("Nome: %s\n", felipe.nome)
	fmt.Printf("Salário: R$%.2f\n", felipe.salario)
	fmt.Printf("Tempo de serviço: %d anos\n", felipe.tempoServico())
}
