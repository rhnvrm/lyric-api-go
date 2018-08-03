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

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	result := document.Find("#songLyricsDiv").First()
	return goquery_helpers.RenderSelection(result, "")
}
