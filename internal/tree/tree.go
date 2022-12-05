package tree

import "strings"

const (
	any        = '*'
	notInPlace = '!'
)

type Word struct {
	children map[int][]*wordNode
}

type wordNode struct {
	val      rune
	children []*wordNode
}

func NewWord() *Word {
	return &Word{
		children: map[int][]*wordNode{},
	}
}

func (w *Word) Insert(word string) {
	if word == "" {
		return
	}

	// only supports cyrillic languages for now
	if !wordIsCyrillic(word) {
		return
	}

	n := countRunes(word)

	cc, ok := w.children[n]
	if !ok {
		root := buildRootTree(word, n)
		w.children[n] = []*wordNode{root}

		return
	}

	r, stripped := stripFirstRune(word)

	for _, c := range cc {
		if c.val == r {
			if n == 1 {
				return
			}

			c.insert(stripped)
			return
		}

		if c.val > r {
			w.children[n] = insertChild(cc, buildRootTree(word, n), n)
			return
		}
	}

	w.children[n] = append(cc, buildRootTree(word, n))
}

func (w *Word) Find(word string, ignored []rune, guessed []rune) []string {
	n := countRunes(word)

	cc, ok := w.children[n]
	if !ok {
		return []string{}
	}

	res := []string{}

	ignoredMap := make(map[rune]struct{}, len(ignored))
	for _, r := range ignored {
		ignoredMap[r] = struct{}{}
	}

	for _, c := range cc {
		found := c.find(word, n, ignoredMap, true)
		for _, f := range found {
			if !wordContainsAllGuessed(f, guessed) {
				continue
			}

			res = append(res, f)
		}
	}

	return res
}

func wordContainsAllGuessed(word string, guessed []rune) bool {
	for _, r := range guessed {
		if !strings.ContainsRune(word, r) {
			return false
		}
	}

	return true
}

func (w *wordNode) insert(word string) {
	if word == "" {
		return
	}

	n := countRunes(word)
	r, stripped := stripFirstRune(word)

	for i, c := range w.children {
		if c.val == r {
			if n == 1 {
				return
			}

			c.insert(stripped)
			return
		}

		if c.val > r {
			w.children = insertChild(w.children, buildRootTree(word, n), i)
			return
		}
	}

	w.children = append(w.children, buildRootTree(word, n))
}

func (w *wordNode) find(word string, n int, ignored map[rune]struct{}, letterInPlace bool) []string {
	if _, ok := ignored[w.val]; ok {
		return []string{}
	}

	if n <= 0 {
		return []string{}
	}

	r, stripped := stripFirstRune(word)
	if r == notInPlace {
		return w.find(stripped, n, ignored, false)
	}

	if !letterInPlace {
		if r == w.val {
			return []string{}
		}

		r = any
	}

	if r != w.val && r != any && r != notInPlace {
		return []string{}
	}

	results := []string{}

	if len(w.children) == 0 && (r == w.val || r == any) {
		return []string{string(w.val)}
	}

	for _, c := range w.children {
		found := c.find(stripped, n-1, ignored, r != notInPlace)

		for _, f := range found {
			results = append(results, string(w.val)+f)
		}
	}

	return results
}

func insertChild(children []*wordNode, c *wordNode, pos int) []*wordNode {
	if len(children) < pos {
		return children
	}

	if pos < 0 {
		return children
	}

	children = append(children, nil)

	for i := len(children) - 1; i > pos; i-- {
		children[i], children[i-1] = children[i-1], children[i]
	}

	children[pos] = c

	return children
}

func buildRootTree(word string, n int) *wordNode {
	if n <= 0 {
		return nil
	}

	if n == 1 {
		r, _ := stripFirstRune(word)

		return &wordNode{
			val: r,
		}
	}

	r, w := stripFirstRune(word)
	child := buildRootTree(w, n-1)

	return &wordNode{
		val:      r,
		children: []*wordNode{child},
	}
}
