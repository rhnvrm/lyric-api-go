package genius

import (
	"encoding/json"
	"errors"
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

func search(artist, song, accessToken string) ([]byte, error) {
	url := "http://api.genius.com/search?access_token=" + accessToken + "&q=" + url.PathEscape(artist) + "-" + url.PathEscape(song)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("non 200 error code from API, got " + string(resp.StatusCode) + " : " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func parse(data []byte) (string, error) {
	var res map[string]interface{}

	if err := json.Unmarshal(data, &res); err != nil {
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

	result := document.Find(".lyrics").First()
	return strings.TrimSpace(goquery_helpers.RenderSelection(result, "\n")), nil
}

// Fetch Searches Genius API based on Artist and Song. Then parses the result,
// to get a song and obtaines the url and scrapes it to return the lyrics.
func (g *Genius) Fetch(artist, song string) string {
	d, err := search("John Lennon", "imagine", g.accessToken)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	u, err := parse(d)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	lyric, err := scrape(u)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	return lyric
}
