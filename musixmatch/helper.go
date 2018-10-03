package musixmatch

import "strings"

func formatURL(x string) string {
	return strings.Replace(x, " ", "-", -1)
}
