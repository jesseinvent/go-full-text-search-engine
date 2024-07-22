package utils

import (
	"strings"

	"github.com/kljensen/snowball"
)

// Lowercase words
func LowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}

	return r
}

// Remove commonly used words
func StopWordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	r := make([]string, len(tokens))

	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}

	return r
}

func StemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i], _ = snowball.Stem(token, "english", false)

	}

	return r
}
