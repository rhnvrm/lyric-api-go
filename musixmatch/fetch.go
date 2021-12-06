package musixmatch

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

// MusixMatch Provider.
type MusixMatch struct {
	http *http.Client
}

// New creates an instance of MusixMatch Provider.
func New() *MusixMatch {
	return &MusixMatch{http: http.DefaultClient}
}

// NewWithHTTP creates an instance of MusixMatch Provider with a custom http client.
func NewWithHTTP(http *http.Client) *MusixMatch {
	return &MusixMatch{http: http}
}

// Fetch scrapes MusixMatch based on Artist and Song.
func (m *MusixMatch) Fetch(artist, song string) string {
	url := "https://www.musixmatch.com/lyrics/" + formatURL(artist) + "/" + formatURL(song)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:48.0) Gecko/20100101 Firefox/48.0")
	res, err := m.http.Do(req)
	if err != nil {
		log.Println("error during http request while attempting musixmatch package ", err)
		return ""
	}
	defer res.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("error in reading document body while attempting musixmatch package ", err)
		return ""
	}

	result := document.Find(".mxm-lyrics__content")

	return goquery_helpers.RenderSelection(result, "")
}

func formatURL(x string) string {
	return strings.Replace(x, " ", "-", -1)
}
