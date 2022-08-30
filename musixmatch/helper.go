package musixmatch

import (
	"regexp"
	"strings"
)

var (
	urlRegexp        = regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	whitespaceRegexp = regexp.MustCompile(`\s+`)
)

func formatURL(x string) string {
	x = strings.ReplaceAll(x, "'", " ")
	x = urlRegexp.ReplaceAllString(x, "")
	x = whitespaceRegexp.ReplaceAllString(x, " ")
	return strings.Replace(x, " ", "-", -1)
}
