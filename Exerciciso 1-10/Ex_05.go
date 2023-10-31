package main

import "fmt"

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

func main() {
	p := []Playlist{
		{
			Nome: "funkzao",
			Musicas: []Musica{
				{
					Titulo:  "fonk",
					Artista: "dj rato",
					Duracao: 180,
				},
				{
					Titulo:  "funk 2",
					Artista: "dj bola",
					Duracao: 180,
				},
				{
					Titulo:  "funk 1",
					Artista: "dj kvl",
					Duracao: 180,
				},
			},
		},
		{
			Nome: "funkzao 2",
			Musicas: []Musica{
				{
					Titulo:  "bate",
					Artista: "eren",
					Duracao: 180,
				},
				{
					Titulo:  "bola",
					Artista: "eren",
					Duracao: 180,
				},
			},
		},
	}

	fmt.Print(FindPlaylists(p, "bate"))
}

func FindPlaylists(p []Playlist, nomeMusica string) []Playlist {
	var found []Playlist

	for _, playlist := range p {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == nomeMusica {
				found = append(found, playlist)
			}
		}
	}

	return found
}
