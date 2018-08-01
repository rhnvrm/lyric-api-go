package main

import (
	"fmt"

	"github.com/rhnvrm/lyric-api-go"
)

func main() {
	var (
		artist = "John Lennon"
		song   = "Imagine"
	)

	fmt.Println(lyrics.Fetch(artist, song))
}
