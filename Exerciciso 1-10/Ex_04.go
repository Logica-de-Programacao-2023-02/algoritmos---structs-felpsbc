package main

import "fmt"

type Playlist struct {
	Nome    string
	Musicas []Musica
}
type Musica struct {
	titulo  string
	artista string
	duracao int
}

func main() {

	p := Playlist{
		Nome: "Funk",
		Musicas: []Musica{
			{
				titulo:  "funk 1",
				artista: "dj rato",
				duracao: 180,
			},
			{
				titulo:  "funk 2",
				artista: "dj bola",
				duracao: 180,
			},
			{
				titulo:  "funk 3",
				artista: "dj rato",
				duracao: 180,
			},
		},
	}
	PlaylistInfo(p)
}
func PlaylistInfo(p Playlist) {
	total := 0
	for _, musica := range p.Musicas {
		total = total + musica.duracao
	}
	fmt.Printf("A playlist %s tem %d:%d minutos\n", p.Nome, total/60, total%60)
	for _, musica := range p.Musicas {
		fmt.Printf("A música %s do artista %s tem %d:%d minutos\n", musica.titulo, musica.artista, musica.duracao/60, musica.duracao%60)
	}
}
