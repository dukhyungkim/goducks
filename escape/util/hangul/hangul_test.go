package hangul

import "testing"

func TestExtract(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "너구리 -> 리",
			args: args{s: "너구리"},
			want: '리',
		},
		{
			name: "신라면 -> 면",
			args: args{s: "신라면"},
			want: '면',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.s); got != tt.want {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasJongSung(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "리 -> true",
			args: args{r: '리'},
			want: false,
		},
		{
			name: "면 -> true",
			args: args{r: '면'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasJongSung(tt.args.r); got != tt.want {
				t.Errorf("HasJongSung() = %v, want %v", got, tt.want)
			}
		})
	}
}
