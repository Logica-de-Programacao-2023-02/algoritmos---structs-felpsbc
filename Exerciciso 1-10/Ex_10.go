package main

import "fmt"

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []avaliacao
}

type avaliacao struct {
	nome string
	nota int
}

func main() {
	f := Filme{
		titulo:  "Para todos os garotos que já amei",
		diretor: "Julia Roberts",
		ano:     2017,
	}
	f = AdicionarAvaliacao(f, "Usuário1", 4)
	f = AdicionarAvaliacao(f, "Usuário2", 5)

	f = RemoverAvaliacao(f, 0)

	FilmeInfo(f)
}

func AdicionarAvaliacao(f Filme, nomeUsuario string, nota int) Filme {
	f.avaliacoes = append(f.avaliacoes, avaliacao{
		nome: nomeUsuario,
		nota: nota,
	})
	return f
}

func RemoverAvaliacao(f Filme, posicao int) Filme {
	if posicao >= 0 && posicao < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:posicao], f.avaliacoes[posicao+1:]...)
	}
	return f
}

func CalcularMedia(f Filme) float64 {
	var total int
	for _, avaliacao := range f.avaliacoes {
		total += avaliacao.nota
	}
	if len(f.avaliacoes) > 0 {
		return float64(total) / float64(len(f.avaliacoes))
	}
	return 0.0
}
func FilmeInfo(f Filme) {
	fmt.Printf("O filme %s (%d) do diretor %s tem a média de avaliações de %.2f\n",
		f.titulo, f.ano, f.diretor, CalcularMedia(f))
}
