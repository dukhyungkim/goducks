package hangul

import "testing"

func Test_extractLast(t *testing.T) {
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
			if got := extractLast(tt.args.s); got != tt.want {
				t.Errorf("extractLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasJongSung(t *testing.T) {
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
			if got := hasJongSung(tt.args.r); got != tt.want {
				t.Errorf("hasJongSung() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithJosa(t *testing.T) {
	type args struct {
		s string
		j Josa
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "다람쥐 은/는",
			args: args{
				s: "다람쥐",
				j: EunNun,
			},
			want: "다람쥐는",
		},
		{
			name: "다람쥐 이/가",
			args: args{
				s: "다람쥐",
				j: LeeGa,
			},
			want: "다람쥐가",
		},
		{
			name: "다람쥐 을/를",
			args: args{
				s: "다람쥐",
				j: EulLul,
			},
			want: "다람쥐를",
		},
		{
			name: "다람쥐 와/과",
			args: args{
				s: "다람쥐",
				j: WaGwa,
			},
			want: "다람쥐와",
		},
		{
			name: "사슴 은/는",
			args: args{
				s: "사슴",
				j: EunNun,
			},
			want: "사슴은",
		},
		{
			name: "사슴 이/가",
			args: args{
				s: "사슴",
				j: LeeGa,
			},
			want: "사슴이",
		},
		{
			name: "사슴 을/를",
			args: args{
				s: "사슴",
				j: EulLul,
			},
			want: "사슴을",
		},
		{
			name: "사슴 와/과",
			args: args{
				s: "사슴",
				j: WaGwa,
			},
			want: "사슴과",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithJosa(tt.args.s, tt.args.j); got != tt.want {
				t.Errorf("WithJosa() = %v, want %v", got, tt.want)
			}
		})
	}
}
