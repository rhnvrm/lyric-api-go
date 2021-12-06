package genius

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

var (
	ErrInvalidResponse = errors.New("invalid response")
	ErrLyricsNotFound  = errors.New("no lyrics found")
	ErrRequestFailed   = errors.New("request failed")
)

// Genius Provider.
type Genius struct {
	accessToken string
	http        *http.Client
}

// New creates an instance of genius provider.
func New(accessToken string) *Genius {
	return &Genius{
		accessToken: accessToken,
		http:        http.DefaultClient,
	}
}

// NewWithHTTP creates an instance of genius provider with a custom http client.
func NewWithHTTP(accessToken string, http *http.Client) *Genius {
	return &Genius{
		accessToken: accessToken,
		http:        http,
	}
}

func (g *Genius) search(artist, song string) (string, error) {
	url := "http://api.genius.com/search?" +
		"access_token=" + g.accessToken +
		"&q=" + url.PathEscape(artist) + "-" + url.PathEscape(song)

	resp, err := g.http.Get(url)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return "", err
	}
	defer resp.Body.Close()
	defer io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response (%d `%s`): %w", resp.StatusCode, resp.Status, ErrRequestFailed)
	}

	return parse(resp.Body)
}

func parse(data io.Reader) (string, error) {
	var res map[string]interface{}

	if err := json.NewDecoder(data).Decode(&res); err != nil {
		return "", err
	}

	resp, ok := res["response"].(map[string]interface{})
	if !ok {
		return "", ErrInvalidResponse
	}

	hits, ok := resp["hits"].([]interface{})
	if !ok {
		return "", ErrInvalidResponse
	}

	for _, v := range hits {
		h, ok := v.(map[string]interface{})
		if !ok {
			return "", ErrInvalidResponse
		}

		if h["type"] == "song" {
			res, ok := h["result"].(map[string]interface{})
			if !ok {
				return "", ErrInvalidResponse
			}

			url, ok := res["url"].(string)
			if !ok {
				return "", ErrInvalidResponse
			}

			return url, nil
		}
	}
	return "", ErrLyricsNotFound
}

// Scrape scrapes the given Genius url to get the lyrics.
func (g *Genius) Scrape(url string) (string, error) {
	res, err := g.http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response (%d `%s`): %w", res.StatusCode, res.Status, ErrRequestFailed)
	}

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
	u, err := g.search(artist, song)
	if err != nil {
		log.Println("error in genius provider during search while attempting genius provider ", err)
		return ""
	}

	lyric, err := g.Scrape(u)
	if err != nil {
		log.Println("error in genius provider during scraping while attempting genius provider ", err)
		return ""
	}
	return lyric
}
