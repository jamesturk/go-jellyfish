package strfry

import "strings"

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func MatchRatingCodex(str string) string {
	str = strings.ToUpper(str)
	var prev rune
	var codex []rune

	for i, c := range str {
		if c != ' ' && (i == 0 && isVowel(c)) || (!isVowel(c) && c != prev) {
			codex = append(codex, c)
		}

		prev = c
	}

	// cut the middle if longer than 6
	if len(codex) > 6 {
		return string(codex[:3]) + string(codex[len(codex)-3:])
	} else {
		return string(codex)
	}
}
