package musixmatch

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
			name: "Michael Jackson - Remember the Time",
			args: args{
				artist: "Michael Jackson",
				song:   "Remember the Time",
			},
			want: `Do you remember`,
		},
		{
			name: "Linkin Park - Numb",
			args: args{
				artist: "Linkin Park",
				song:   "Numb",
			},
			want: `And every second I waste is more than I can take`,
		},
	}

	for _, tt := range tests {
		require.Contains(t, provider.Fetch(tt.args.artist, tt.args.song), tt.want)
	}
}
