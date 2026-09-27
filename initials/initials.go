// Package initials turns a person's name into the letters an avatar shows in place of a picture.
package initials

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Of returns the first letters of the first and last words of name, upper-cased, or the first
// letter of fallback when name is blank.
func Of(name, fallback string) string {
	words := strings.Fields(name)

	switch len(words) {
	case 0:
		return firstLetter(fallback)
	case 1:
		return firstLetter(words[0])
	default:
		return firstLetter(words[0]) + firstLetter(words[len(words)-1])
	}
}

func firstLetter(word string) string {
	r, _ := utf8.DecodeRuneInString(strings.TrimSpace(word))
	if r == utf8.RuneError {
		return ""
	}

	return string(unicode.ToUpper(r))
}
