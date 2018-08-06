package lyricswikia

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFetch(t *testing.T) {
	type args struct {
		artist string
		song   string
	}
	provider := New()

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
				},
				want: `And every second I waste is more than I can take`,
			},
			{
				name: "John Lennon - Imagine",
				args: args{
					artist: "John Lennon",
					song:   "Imagine",
				},
				want: `No need for greed or hunger`,
			},
		}

		Convey("Want should be a substring of Got", func() {
			for _, tt := range tests {
				So(provider.Fetch(tt.args.artist, tt.args.song), ShouldContainSubstring, tt.want)
			}
		})
	})
}
