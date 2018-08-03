package lyrics

import (
	"errors"

	"github.com/rhnvrm/lyric-api-go/genius"
	"github.com/rhnvrm/lyric-api-go/lyricswikia"
	"github.com/rhnvrm/lyric-api-go/songlyrics"
)

type provider interface {
	Fetch(artist, song string) string
}

// Supported Providers:
// Default
// - Lyrics Wikia	(github.com/rhnvrm/lyric-api-go/lyricswikia)
// - Song Lyrics	(github.com/rhnvrm/lyric-api-go/songlyrics)
// Require Setup
// - Genius 		(github.com/rhnvrm/lyric-api-go/genius)
var defaultProviders = []provider{
	lyricswikia.New(),
	songlyrics.New(),
}

// Lyric API.
type Lyric struct {
	providers []provider
}

// Option type describes Option Configuration Decorator return type.
type Option func(Lyric) Lyric

// WithAllProviders is an Option Configuration Decorator that sets
// Lyric to attempt fetching lyrics using all providers.
func WithAllProviders() Option {
	return func(l Lyric) Lyric {
		l.providers = defaultProviders
		return l
	}
}

// WithLyricsWikia is an Option Configuration Decorator that adds
// Lyrics Wikia Provider to the list of providers to attempt fetching
// lyrics from.
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
		l.providers = append(l.providers, songlyrics.New())
		return l
	}
}

// WithGenius is an Option Configuration Decorator that adds
// Genius Provider to the list of providers to attempt fetching
// lyrics from. It requires an additional access token which can
// be obtained using the developer portal (https://genius.com/developers)
func WithGeniusLyrics(accessToken string) Option {
	return func(l Lyric) Lyric {
		l.providers = append(l.providers, genius.New(accessToken))
		return l
	}
}

// New creates a new Lyric API, which can be used to Search for Lyrics
// using various providers. The default behaviour is to use all
// providers available, although it can be explicitly set to the same
// using, eg.
// 		lyrics.New(WithAllProviders())
// In case your usecase requires using only specific providers,
// you can provide New() with
// the specific WithXXXXProvider() as an optional parameter.
//
// Eg. to attempt only with Lyrics Wikia:
// 		lyrics.New(WithLyricsWikia())
//
// Eg. to attempt with both Lyrics Wikia and Song Lyrics:
// 		lyrics.New(WithLyricsWikia(), WithSongLyrics())
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
	for _, p := range l.providers {
		lyric := p.Fetch(artist, song)
		if len(lyric) > 5 { // Arbitrary size to make sure not empty.
			return lyric, nil
		}
	}
	return "", errors.New("Not Found")
}
