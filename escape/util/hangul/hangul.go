package hangul

import (
	"unicode/utf8"
)

type Josa string

const (
	EunNun Josa = "은/는"
	LeeGa  Josa = "이/가"
	WaGwa  Josa = "와/과"
	EulLul Josa = "을/를"
)

var josaIndex = map[Josa]map[bool]string{
	EunNun: {true: "은", false: "는"},
	LeeGa:  {true: "이", false: "가"},
	WaGwa:  {true: "과", false: "와"},
	EulLul: {true: "을", false: "를"},
}

func extractLast(s string) rune {
	last, _ := utf8.DecodeLastRune([]byte(s))
	return last
}

func hasJongSung(r rune) bool {
	return (r-0xac00)%28 != 0
}

func WithJosa(s string, j Josa) string {
	return s + josaIndex[j][hasJongSung(extractLast(s))]
}
