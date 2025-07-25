package funcs

import "testing"

func TestMaximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			name: "1",
			args: args{
				s: "aabbaaxybbaabb",
				x: 5,
				y: 4,
			},
			want: 20,
		},
		{
			name: "2",
			args: args{
				s: "cdbcbbaaabab",
				x: 4,
				y: 5,
			},
			want: 19,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumGain(tt.args.s, tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("Abba() got = %v, want %v", got, tt.want)
			}

		})
	}
}
