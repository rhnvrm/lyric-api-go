package songlyrics

import (
	"log"
	"net/http"
	"errors"

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
func (*SongLyrics) Fetch(artist, song string) (string, error) {
	url := "http://www.songlyrics.com/" + slug.Make(artist) + "/" + slug.Make(song) + "-lyrics/"

	res, err := http.Get(url)
	if err != nil {
		log.Println("error during http request while attempting songlyrics provider ", err)
		return "", err
	}
	defer res.Body.Close()
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("error in reading document body while attempting songlyrics provider ", err)
		return "", err
	}

	result := document.Find("#songLyricsDiv").First()
	output := goquery_helpers.RenderSelection(result, "")

	sorryLen := len("Sorry")
	if len(output) >= sorryLen && output[:sorryLen] == "Sorry" {
		return "", errors.New("sorry, no results found")
	}

	return output, err
}
