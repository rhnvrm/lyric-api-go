package lyricswikia

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	type args struct {
		artist string
		song   string
	}
	provider := New()

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

	for _, tt := range tests {
		require.Contains(t, provider.Fetch(tt.args.artist, tt.args.song), tt.want)
	}
}
