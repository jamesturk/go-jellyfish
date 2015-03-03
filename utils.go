package jellyfish

import (
	"golang.org/x/text/unicode/norm"
	"strings"
)

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func normalize(str string) []rune {
	str = string(norm.NFKD.Bytes([]byte(str)))
	return []rune(strings.ToUpper(str))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
