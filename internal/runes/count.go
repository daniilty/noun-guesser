package runes

func Count(s string, ignored []rune) int {
	n := 0

	ignoredMap := make(map[rune]struct{}, len(ignored))
	for _, r := range ignored {
		ignoredMap[r] = struct{}{}
	}

	for _, r := range s {
		if _, ok := ignoredMap[r]; ok {
			continue
		}

		n++
	}

	return n
}
