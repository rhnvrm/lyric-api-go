package genius

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO: Add mock responses and check other functions.

func TestScrape(t *testing.T) {
	type args struct {
		artist string
		song   string
		url    string
	}

	Convey("For each song in the test cases", t, func() {
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

		Convey("Want should be a substring of Got", func() {
			for _, tt := range tests {
				result, err := scrape(tt.args.url)
				So(err, ShouldBeNil)
				So(result, ShouldContainSubstring, tt.want)
			}
		})
	})
}
