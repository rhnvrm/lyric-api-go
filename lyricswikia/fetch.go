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
		log.Println("error during http request while attempting lyricswikia provider ", err)
		return ""
	}
	defer res.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("error in reading document body while attempting lyricswikia provider ", err)
		return ""
	}

	result := document.Find("div.lyricbox").First()
	return goquery_helpers.RenderSelection(result, "\n")
}
