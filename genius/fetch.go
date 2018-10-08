package genius

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

// Genius Provider.
type Genius struct {
	accessToken string
}

// New creates an instance of genius provider.
func New(accessToken string) *Genius {
	return &Genius{
		accessToken: accessToken,
	}
}

func search(artist, song, accessToken string) (string, error) {
	url := "http://api.genius.com/search?access_token=" + accessToken + "&q=" + url.PathEscape(artist) + "-" + url.PathEscape(song)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return "", err
	}
	defer resp.Body.Close()
	defer io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode != 200 {
		return "", errors.New("non 200 error code from API, got " + string(resp.StatusCode) + " : " + resp.Status)
	}

	return parse(resp.Body)
}

func parse(data io.Reader) (string, error) {
	var res map[string]interface{}

	if err := json.NewDecoder(data).Decode(&res); err != nil {
		return "", err
	}
	hits := res["response"].(map[string]interface{})["hits"].([]interface{})
	for _, v := range hits {
		h := v.(map[string]interface{})
		if h["type"] == "song" {
			url := h["result"].(map[string]interface{})["url"].(string)
			return url, nil
		}
	}
	return "", errors.New("no song found")
}

func scrape(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	result := document.Find(".lyrics").First()
	return strings.TrimSpace(goquery_helpers.RenderSelection(result, "\n")), nil
}

// Fetch Searches Genius API based on Artist and Song. Then parses the result,
// to get a song and obtaines the url and scrapes it to return the lyrics.
func (g *Genius) Fetch(artist, song string) string {
	u, err := search(artist, song, g.accessToken)
	if err != nil {
		log.Println("error in genius provider during search while attempting genius provider ", err)
		return ""
	}
	lyric, err := scrape(u)
	if err != nil {
		log.Println("error in genius provider during scraping while attempting genius provider ", err)
		return ""
	}
	return lyric
}
