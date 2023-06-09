package hangul

import (
	"unicode/utf8"
)

func Extract(s string) rune {
	last, _ := utf8.DecodeLastRune([]byte(s))
	return last
}

func HasJongSung(r rune) bool {
	return (r-0xac00)%28 != 0
}
