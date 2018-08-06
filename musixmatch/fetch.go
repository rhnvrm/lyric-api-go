package musixmatch

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

// MusixMatch Provider.
type MusixMatch struct {
}

// New creates an instance of MusixMatch Provider.
func New() *MusixMatch {
	return &MusixMatch{}
}

// Fetch scrapes MusixMatch based on Artist and Song.
func (*MusixMatch) Fetch(artist, song string) string {
	url := "https://www.musixmatch.com/lyrics/linkin-park/numb"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "a1686462-7043-6bf1-a401-543c7ab9746b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	fmt.Println(document.Text(), url, res.StatusCode)

	result := document.Find(".mxm-lyrics__content").First()
	return goquery_helpers.RenderSelection(result, "")
}
