package jellyfish

func consonant(str []rune, i int) bool {
	switch str[i] {
	case 'a', 'e', 'i', 'o', 'u':
		return false
	case 'y':
		if i == 0 {
			return true
		} else {
			return !consonant(str, i-1)
		}
	default:
		return true
	}
}

func measure(str []rune) int {
	count := 0
	if len(str) > 0 {
		consonant_state := consonant(str, 0)

		for i := 0; i < len(str); i++ {
			is_cons := consonant(str, i)
			if is_cons && !consonant_state {
				count++
			}
			consonant_state = is_cons
		}
	}
	return count
}

// true iff str contains vowel
func vowel_in_stem(str []rune) bool {
	for i := range str {
		if !consonant(str, i) {
			return true
		}
	}
	return false
}

func ends_with(str []rune, suffix string) bool {
	rs := []rune(suffix)
	offset := len(str) - len(rs)
	if offset < 0 {
		return false
	}
	for i, r := range suffix {
		if str[offset+i] != r {
			return false
		}
	}
	return true
}

// deal with sses, ies, ss, s
func one_a(str []rune) []rune {
	slen := len(str)
	if ends_with(str, "sses") || ends_with(str, "ies") {
		return str[:slen-2]
	} else if str[slen-1] == 's' && str[slen-2] != 's' {
		return str[:slen-1]
	}
	return str
}

// check ending of string for cons-vowel-cons suffix
func star_o(str []rune) bool {
	slen := len(str)
	if slen >= 3 && consonant(str, slen-3) && !consonant(str, slen-2) && consonant(str, slen-1) {
		return str[slen-1] != 'w' && str[slen-1] != 'x' && str[slen-1] != 'y'
	}
	return false
}

// handle (at, bl, iz) suffix and trim cons-cons suffix if it isn't (l, s, z)
func one_b_a(str []rune) []rune {
	slen := len(str)

	if ends_with(str, "at") || ends_with(str, "bl") || ends_with(str, "iz") {
		return append(str, 'e')
	} else if consonant(str, slen-1) && str[slen-1] == str[slen-2] {
		if str[slen-1] != 'l' && str[slen-1] != 's' && str[slen-1] != 'z' {
			return str[:slen-1]
		}
	} else if star_o(str) && measure(str) == 1 {
		return append(str, 'e')
	}

	return str
}

// handle eed, ed, ing
func one_b(str []rune) []rune {
	if ends_with(str, "eed") {
		if measure(str[:len(str)-3]) > 0 {
			return str[:len(str)-1]
		}
	} else if ends_with(str, "ed") {
		tmp := str[:len(str)-2]
		if vowel_in_stem(tmp) {
			return one_b_a(tmp)
		}
	} else if ends_with(str, "ing") {
		tmp := str[:len(str)-3]
		if vowel_in_stem(tmp) {
			return one_b_a(tmp)
		}
	}
	return str
}

// handle y ending
func one_c(str []rune) []rune {
	if str[len(str)-1] == 'y' && vowel_in_stem(str[:len(str)-1]) {
		str[len(str)-1] = 'i'
	}
	return str
}

// helper func for steps 2-4
func cond_replace(str []rune, replacements []replacement, min_measure int) []rune {
	for _, r := range replacements {
		if ends_with(str, r.from) {
			tmp := str[:len(str)-len(r.from)]
			if measure(tmp) > min_measure {
				// special case for [st]ion
				if r.from == "ion" {
					if len(str) > 4 && (str[len(str)-4] == 's' || str[len(str)-4] == 't') {
						return str[:len(str)-3]
					}
				} else if r.to != "" {
					return append(tmp, []rune(r.to)...)
				} else {
					return tmp
				}
			}
			return str
		}
	}
	return str
}

type replacement struct {
	from, to string
}

func two(str []rune) []rune {
	replacements := []replacement{
		{"ational", "ate"},
		{"tional", "tion"},
		{"enci", "ence"},
		{"anci", "ance"},
		{"izer", "ize"},
		{"abli", "able"},
		{"bli", "ble"},
		{"alli", "al"},
		{"entli", "ent"},
		{"eli", "e"},
		{"ousli", "ous"},
		{"ization", "ize"},
		{"ation", "ate"},
		{"ator", "ate"},
		{"alism", "al"},
		{"iveness", "ive"},
		{"fulness", "ful"},
		{"ousness", "ous"},
		{"aliti", "al"},
		{"iviti", "ive"},
		{"biliti", "ble"},
		{"logi", "log"},
	}
	return cond_replace(str, replacements, 0)
}

func three(str []rune) []rune {
	replacements := []replacement{
		{"icate", "ic"},
		{"ative", ""},
		{"alize", "al"},
		{"iciti", "ic"},
		{"ical", "ic"},
		{"ful", ""},
		{"ness", ""},
	}
	return cond_replace(str, replacements, 0)
}

func four(str []rune) []rune {
	replacements := []replacement{
		{"al", ""},
		{"ance", ""},
		{"ence", ""},
		{"er", ""},
		{"ic", ""},
		{"able", ""},
		{"ible", ""},
		{"ant", ""},
		{"ement", ""},
		{"ment", ""},
		{"ent", ""},
		{"ion", ""},
		{"ou", ""},
		{"ism", ""},
		{"ate", ""},
		{"iti", ""},
		{"ous", ""},
		{"ive", ""},
		{"ize", ""},
	}
	return cond_replace(str, replacements, 1)
}

// sometimes truncate a trailing e
func five_a(str []rune) []rune {
	last := len(str) - 1
	if str[last] == 'e' {
		tmp := str[:last]
		tmp_meas := measure(tmp)
		if tmp_meas > 1 || (tmp_meas == 1 && !star_o(tmp)) {
			return tmp
		}
	}
	return str
}

// handle double l at end of word
func five_b(str []rune) []rune {
	slen := len(str)
	if measure(str) > 1 && str[slen-1] == 'l' && str[slen-1] == str[slen-2] {
		return str[:slen-1]
	}
	return str
}

func Porter(str string) string {
	runes := []rune(str)
	if len(runes) > 2 {
		return string(five_b(five_a(four(three(two(one_c(one_b(one_a(runes)))))))))
	}
	return string(runes)
}
