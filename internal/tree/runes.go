package tree

import (
	"unicode"
)

func stripFirstRune(s string) (rune, string) {
	if s == "" {
		return 0, s
	}

	var (
		i            int
		firstRune, r rune
		run          bool
	)

	for i, r = range s {
		if run {
			break
		}

		run = true
		firstRune = r
	}

	if i == 0 {
		return firstRune, ""
	}

	return firstRune, s[i:]
}

func countRunes(s string) int {
	var n int

	for _, r := range s {
		if r == notInPlace {
			continue
		}

		n++
	}

	return n
}

func wordIsCyrillic(word string) bool {
	for _, r := range word {
		if !unicode.Is(unicode.Cyrillic, r) {
			return false
		}
	}

	return true
}
