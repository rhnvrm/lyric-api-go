package genius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: Add mock responses and check other functions.

func TestScrape(t *testing.T) {
	type args struct {
		artist string
		song   string
		url    string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Linkin Park - Numb",
			args: args{
				artist: "Linkin Park",
				song:   "Numb",
				url:    "https://genius.com/Linkin-park-numb-lyrics",
			},
			want: `And every second I waste is more than I can take`,
		},
		{
			name: "John Lennon - Imagine",
			args: args{
				artist: "John Lennon",
				song:   "Imagine",
				url:    "https://genius.com/John-lennon-imagine-lyrics",
			},
			want: `No need for greed or hunger`,
		},
	}

	g := New("")
	for _, tt := range tests {
		result, err := g.Scrape(tt.args.url)
		require.Nil(t, err)
		require.Contains(t, result, tt.want)
	}
}
