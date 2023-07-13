package allrecipes

import "testing"

func TestToMinutes(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{
			arg:  "10 mins",
			want: 10,
		},
		{
			arg:  "6 hrs 10 mins",
			want: 370,
		},
		{
			arg:  "12 hrs",
			want: 720,
		},
		{
			arg:  "1 min",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := toMinutes(tt.arg); got != tt.want {
				t.Errorf("toMinutes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{
			arg:  "10",
			want: 10,
		},
		{
			arg:  "610",
			want: 610,
		},
		{
			arg:  "12,123",
			want: 12123,
		},
		{
			arg:  "1",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := toInt(tt.arg); got != tt.want {
				t.Errorf("toInt(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
