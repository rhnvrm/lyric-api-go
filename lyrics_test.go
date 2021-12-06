package lyrics_test

import (
	"fmt"

	lyrics "github.com/rhnvrm/lyric-api-go"
)

func Example_genius() {
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

func Example_search() {
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
