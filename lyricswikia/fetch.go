// Package lyricswikia implements LyricsWikia provider. THIS IS NOW DEPRECATED
// as the lyricswikia wiki is no longer hosted. You can manually set the Host in
// case a new host supporting a similar API is hosted somewhere.
package lyricswikia

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

const (
	DefaultHost = "http://lyrics.wikia.com/wiki/"
)

// LyricsWikia Provider.
type LyricsWikia struct {
	http *http.Client
	host string
}

// New creates an instance of LyricsWikia Provider.
func New() *LyricsWikia {
	return &LyricsWikia{http: http.DefaultClient}
}

// NewWithHTTP creates an instance of LyricsWikia Provider with a custom http client.
func NewWithHTTP(http *http.Client) *LyricsWikia {
	return &LyricsWikia{http: http}
}

func (l *LyricsWikia) SetHost(host string) {
	l.host = host
}

// Fetch scrapes Lyrics Wikia based on Artist and Song.
func (l *LyricsWikia) Fetch(artist, song string) string {
	url := DefaultHost
	if l.host != "" {
		url = l.host
	}

	url = url + artist + ":" + song

	// Make HTTP request
	res, err := l.http.Get(url)
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
