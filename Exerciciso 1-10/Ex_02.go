package main

import "fmt"

type Endereço struct {
	rua    string
	número int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereço Endereço
}

func Imprimir(p Pessoa) {
	fmt.Printf("Rua: %s \n", p.endereço.rua)
	fmt.Printf("Número: %d\n", p.endereço.número)
	fmt.Printf("Cidade: %s\n", p.endereço.cidade)
	fmt.Printf("Estado: %s\n", p.endereço.estado)
}

func main() {

	p := Pessoa{
		nome:  "Felipe",
		idade: 19,
		endereço: Endereço{
			rua:    "Barcelos",
			número: 17,
			cidade: "Brasilia",
			estado: "DF",
		},
	}
	Imprimir(p)
}
