package utils

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	// Not letter, not number
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func Analyze(text string) []string {
	tokens := Tokenize(text)
	tokens = LowercaseFilter(tokens)
	tokens = StopWordFilter(tokens)
	tokens = StemmerFilter(tokens)

	return tokens

}
