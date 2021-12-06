package lyrics

import (
	"errors"
	"net/http"

	"github.com/rhnvrm/lyric-api-go/genius"
	"github.com/rhnvrm/lyric-api-go/lyricswikia"
	"github.com/rhnvrm/lyric-api-go/musixmatch"
	"github.com/rhnvrm/lyric-api-go/songlyrics"
)

type Provider interface {
	Fetch(artist, song string) string
}

// Supported Providers:
// Default
// - Lyrics Wikia	(github.com/rhnvrm/lyric-api-go/lyricswikia) (deprecated)
// - Song Lyrics	(github.com/rhnvrm/lyric-api-go/songlyrics)
// - MusixMatch 	(github.com/rhnvrm/lyric-api-go/musixmatch)
// Requires Setup
// - Genius 		(github.com/rhnvrm/lyric-api-go/genius)
var (
	defaultProviders = []Provider{
		songlyrics.New(),
		musixmatch.New(),
	}
)

// Lyric API.
type Lyric struct {
	providers []Provider
	http      *http.Client
}

// Option type describes Option Configuration Decorator return type.
type Option func(Lyric) Lyric

// WithAllProviders is an Option Configuration Decorator that sets
// Lyric to attempt fetching lyrics using all providers that do
// not require setup.
func WithAllProviders() Option {
	return func(l Lyric) Lyric {
		l.providers = defaultProviders
		return l
	}
}

// WithLyricsWikia is an Option Configuration Decorator that adds
// Lyrics Wikia Provider to the list of providers to attempt fetching
// lyrics from. **DEPRECATED**.
func WithLyricsWikia() Option {
	return func(l Lyric) Lyric {
		l.providers = append(l.providers, lyricswikia.New())
		return l
	}
}

// WithSongLyrics is an Option Configuration Decorator that adds
// Song Lyrics Provider to the list of providers to attempt fetching
// lyrics from.
func WithSongLyrics() Option {
	return func(l Lyric) Lyric {
		if l.http == nil {
			l.providers = append(l.providers, songlyrics.New())
		} else {
			l.providers = append(l.providers, songlyrics.NewWithHTTP(l.http))
		}
		return l
	}
}

// WithMusixMatch is an Option Configuration Decorator that adds
// Musixmatch Provider to the list of providers to attempt fetching
// lyrics from.
func WithMusixMatch() Option {
	return func(l Lyric) Lyric {
		if l.http == nil {
			l.providers = append(l.providers, musixmatch.New())
		} else {
			l.providers = append(l.providers, musixmatch.NewWithHTTP(l.http))
		}
		return l
	}
}

// WithGeniusLyrics is an Option Configuration Decorator that adds
// Genius Provider to the list of providers to attempt fetching
// lyrics from. It requires an additional access token which can
// be obtained using the developer portal (https://genius.com/developers)
func WithGeniusLyrics(accessToken string) Option {
	return func(l Lyric) Lyric {
		if l.http == nil {
			l.providers = append(l.providers, genius.New(accessToken))
		} else {
			l.providers = append(l.providers, genius.NewWithHTTP(accessToken, l.http))
		}
		return l
	}
}

// WithoutProviders is an Option Configuration Decorator that removes
// all providers from the list of providers to attempt fetching
// lyrics from. It can be used to remove the default providers and
// set a custom provider list.
func WithoutProviders() Option {
	return func(l Lyric) Lyric {
		l.providers = []Provider{}
		return l
	}
}

// WithHTTPClient is an Option Configuration Decorator that sets
// the HTTP Client to be used by the Lyric API. If you wish to use
// the same client across multiple Lyric API instances, make sure to
// set this option before setting up the other With options.
func WithHTTPClient(c *http.Client) Option {
	return func(l Lyric) Lyric {
		l.http = c
		return l
	}
}

// New creates a new Lyric API, which can be used to Search for Lyrics
// using various providers. The default behaviour is to use all
// providers available, although it can be explicitly set to the same
// using, eg.
// 		lyrics.New(WithAllProviders())
// In case your usecase requires using only specific providers,
// you can provide New() with WithoutProviders() followed by
// the specific WithXXXXProvider() as an optional parameter.
//
// Note: The providers are processed one by one so custom providers,
// can also be used to set the priority for your usecase.
//
// Eg. to attempt only with Lyrics Wikia:
// 		lyrics.New(WithoutProviders(), WithLyricsWikia())
//
// Eg. to attempt with both Lyrics Wikia and Song Lyrics:
// 		lyrics.New(WithoutProviders(), WithLyricsWikia(), WithSongLyrics())
func New(o ...Option) Lyric {
	if len(o) == 0 {
		return Lyric{
			providers: defaultProviders,
		}
	}

	l := Lyric{}
	for _, option := range o {
		l = option(l)
	}

	return l
}

// Search attempts to search for lyrics using artist and song
// by trying various lyrics providers one by one.
func (l *Lyric) Search(artist, song string) (string, error) {
	if len(l.providers) == 0 {
		return "", errors.New("No providers selected")
	}

	for _, p := range l.providers {
		lyric := p.Fetch(artist, song)
		if len(lyric) > 5 { // Arbitrary size to make sure not empty.
			return lyric, nil
		}
	}
	return "", errors.New("Not Found")
}
