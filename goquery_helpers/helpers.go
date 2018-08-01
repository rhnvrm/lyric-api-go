package goquery_helpers

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// RenderSelection is a helper function modified from
// github.com/PuerkitoBio/goquery - Text() function
// it expects and additional separator to add between each
// line.
func RenderSelection(s *goquery.Selection, seperator string) string {
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode && n.Data != "\n" {
			buf.WriteString(n.Data)
			buf.WriteString(seperator)
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	for _, n := range s.Nodes {
		f(n)
	}
	return buf.String()
}
