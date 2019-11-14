package main

import (
	"fmt"

	"github.com/vladcomp/lyric-api-go"
)

func main() {
	var (
		artist = "John Lennon"
		song   = "Imagine"
	)

	l := lyrics.New()
	lyric, err := l.Search(artist, song)

	if err != nil {
		fmt.Printf("Lyrics for %v-%v were not found", artist, song)
	}
	fmt.Println(lyric)
}
