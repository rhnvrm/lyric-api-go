package lyricswikia

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

// LyricsWikia Provider.
type LyricsWikia struct {
}

// New creates an instance of LyricsWikia Provider.
func New() *LyricsWikia {
	return &LyricsWikia{}
}

// Fetch scrapes Lyrics Wikia based on Artist and Song.
func (l *LyricsWikia) Fetch(artist, song string) string {
	url := "http://lyrics.wikia.com/wiki/" + artist + ":" + song

	// Make HTTP request
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

	result := document.Find("div.lyricbox").First()
	return goquery_helpers.RenderSelection(result, "\n")
}
