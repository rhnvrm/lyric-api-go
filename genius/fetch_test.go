package genius

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		artist string
		song   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test",
			args: args{
				artist: "John Lennon",
				song:   "Imagine",
			},
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.artist, tt.args.song); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
