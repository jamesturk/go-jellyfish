package jellyfish

import (
	"strings"
)

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

func MatchRatingComparison(s1, s2 string) bool {
	codex1 := []rune(MatchRatingCodex(s1))
	codex2 := []rune(MatchRatingCodex(s2))
	len1 := len(codex1)
	len2 := len(codex2)
	var min_rating int

	// TODO: should refuse to do comparison.. use an error?
	// hacky way to get abs(len1-len2)
	if max(len1, len2)-min(len1, len2) >= 3 {
		return false
	}

	// minimum rating based on sums of codexes
	lensum := len1 + len2
	switch {
	case lensum <= 4:
		min_rating = 5
	case lensum < 7:
		min_rating = 4
	case lensum < 11:
		min_rating = 3
	default:
		min_rating = 2
	}

	// strip off common prefixes
	var res_longer, res_shorter, longer, shorter []rune
	if len1 > len2 {
		longer = codex1
		shorter = codex2
	} else {
		longer = codex2
		shorter = codex1
	}
	for i, c1 := range longer {
		if i >= len(shorter) {
			res_longer = append(res_longer, c1)
		} else if c1 != shorter[i] {
			res_longer = append(res_longer, c1)
			res_shorter = append(res_shorter, shorter[i])
		}
	}

	unmatched_count1 := 0
	unmatched_count2 := 0

	for i, c1 := range res_longer {
		if i >= len(res_shorter) {
			unmatched_count2++
		} else if c1 != shorter[i] {
			unmatched_count1++
			unmatched_count2++
		}
	}

	return (6 - max(unmatched_count1, unmatched_count2)) >= min_rating
}
