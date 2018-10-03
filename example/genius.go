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

	l := lyrics.New(lyrics.WithoutProviders(), lyrics.WithGeniusLyrics("your_access_token_here"))
	// Use the following if you wish to just add Genius as a fallback
	// l := lyrics.New(lyrics.WithGeniusLyrics("your_access_token_here"))
	lyric, err := l.Search(artist, song)

	if err != nil {
		fmt.Printf("%v: Lyrics for %v-%v were not found", err, artist, song)
	}
	fmt.Println(lyric)
}
