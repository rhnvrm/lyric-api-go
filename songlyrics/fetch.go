package songlyrics

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

// SongLyrics Provider.
type SongLyrics struct {
}

// New creates an instance of SongLyrics Provider.
func New() *SongLyrics {
	return &SongLyrics{}
}

// Fetch scrapes SongLyrics based on Artist and Song.
func (*SongLyrics) Fetch(artist, song string) string {
	url := "http://www.songlyrics.com/" + slug.Make(artist) + "/" + slug.Make(song) + "-lyrics/"

	res, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	result := document.Find("#songLyricsDiv").First()
	output := goquery_helpers.RenderSelection(result, "")
	if output[:len("Sorry")] == "Sorry" {
		return ""
	}
	return output
}
