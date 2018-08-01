package lyrics

import (
	"github.com/rhnvrm/lyric-api-go/lyricswikia"
)

// Fetch attempts to search for lyrics using artist and song
// by trying various lyrics providers.
// Supported Providers:
// - Lyrics Wikia (github.com/rhnvrm/lyric-api-go/lyricswikia)
//
// (support for other providers will be added in the future)
func Fetch(artist, song string) string {
	return lyricswikia.Fetch(artist, song)
}
