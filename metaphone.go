package jellyfish

// Metaphone calculates the metaphone code for a string.
//
// The Metaphone algorithm was designed as an improvement on Soundex.
// It transforms a word into a string consisting of '0BFHJKLMNPRSTWXY' where '0' is pronounced 'th' and 'X' is a '[sc]h' sound.
//
// For example:
//    Metaphone("Klumpz") == Metaphone("Clumps")    // KLMPS
//
// See the Metaphone article at Wikipedia (http://en.wikipedia.org/wiki/Metaphone) for more details.
func Metaphone(s string) string {
	r := normalize(s)
	rlen := len(r)

	// skip first character sometimes

	if rlen > 1 {
		switch {
		case r[0] == 'K' && r[1] == 'N',
			r[0] == 'G' && r[1] == 'N',
			r[0] == 'P' && r[1] == 'N',
			r[0] == 'A' && r[1] == 'C',
			r[0] == 'W' && r[1] == 'R',
			r[0] == 'A' && r[1] == 'E':
			r = r[1:]
			rlen -= 1
		}
	}

	var next rune
	var nextnext rune
	var result []rune

	for i := 0; i < rlen; i++ {
		// get current and next 2 runes
		c := r[i]
		if i+1 < rlen {
			next = r[i+1]
			if i+2 < rlen {
				nextnext = r[i+2]
			} else {
				nextnext = 0
			}
		} else {
			next = 0
		}

		// skip doubles except cc
		if c == next && c != 'C' {
			continue
		}

		// main rules
		switch c {
		case 'A', 'E', 'I', 'O', 'U':
			if i == 0 || r[i-1] == ' ' {
				result = append(result, c)
			}
		case 'B':
			if i == 0 || r[i-1] != 'M' || next != 0 {
				result = append(result, 'B')
			}
		case 'C':
			if next == 'I' && nextnext == 'A' || next == 'H' {
				result = append(result, 'X')
				i++
			} else if next == 'I' || next == 'E' || next == 'Y' {
				result = append(result, 'S')
				i++
			} else {
				result = append(result, 'K')
			}
		case 'D':
			if next == 'G' && (nextnext == 'I' || nextnext == 'E' || nextnext == 'Y') {
				result = append(result, 'J')
				i++
			} else {
				result = append(result, 'T')
			}
		case 'F', 'J', 'L', 'M', 'N', 'R':
			result = append(result, c)
		case 'G':
			if next == 'I' || next == 'E' || next == 'Y' {
				result = append(result, 'J')
			} else if next != 'H' && next != 'N' {
				result = append(result, 'K')
			} else if next == 'H' && !isVowel(nextnext) {
				i++
			}
		case 'H':
			if i == 0 || isVowel(next) || !isVowel(r[i-1]) {
				result = append(result, 'H')
			}
		case 'K':
			if i == 0 || r[i-1] != 'C' {
				result = append(result, 'K')
			}
		case 'P':
			if next == 'H' {
				result = append(result, 'F')
				i++
			} else {
				result = append(result, 'P')
			}
		case 'Q':
			result = append(result, 'K')
		case 'S':
			if next == 'H' {
				result = append(result, 'X')
				i++
			} else if next == 'I' && (nextnext == 'O' || nextnext == 'A') {
				result = append(result, 'X')
				i++
			} else {
				result = append(result, 'S')
			}
		case 'T':
			if next == 'I' && (nextnext == 'O' || nextnext == 'A') {
				result = append(result, 'X')
			} else if next == 'H' {
				result = append(result, '0')
				i++
			} else if next != 'C' || nextnext != 'H' {
				result = append(result, 'T')
			}
		case 'V':
			result = append(result, 'F')
		case 'W':
			if i == 0 && next == 'H' {
				i++
			}
			if isVowel(nextnext) || nextnext == 0 {
				result = append(result, 'W')
			}
		case 'X':
			if i == 0 {
				if next == 'H' || (next == 'I' && (nextnext == 'O' || nextnext == 'A')) {
					result = append(result, 'X')
				} else {
					result = append(result, 'S')
				}
			} else {
				result = append(result, 'K')
				result = append(result, 'S')
			}
		case 'Y':
			if isVowel(next) {
				result = append(result, 'Y')
			}
		case 'Z':
			result = append(result, 'S')
		case ' ':
			if len(result) > 0 && result[len(result)-1] != ' ' {
				result = append(result, ' ')
			}
		}
	}

	return string(result)
}
