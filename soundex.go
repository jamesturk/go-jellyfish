package jellyfish

// Soundex is an algorithm to convert a word (typically a name) to a four
// digit code in the form 'A123' where 'A' is the first letter of the name
// and the digits represent similar sounds.
//
// For example:
//     soundex("Ann") == soundex("Anne")      // A500
//     soundex("Rupert") == soundex("Robert") // R163
//
// See the Soundex article at Wikipedia (http://en.wikipedia.org/wiki/Soundex) for more details.
func Soundex(str string) string {
	if str == "" {
		return ""
	}

	replacements := map[rune]rune{
		'B': '1', 'F': '1', 'P': '1', 'V': '1',
		'C': '2', 'G': '2', 'J': '2', 'K': '2',
		'Q': '2', 'S': '2', 'X': '2', 'Z': '2',
		'D': '3', 'T': '3',
		'L': '4',
		'M': '5', 'N': '5',
		'R': '6',
	}

	var result [4]rune
	count := 1

	// normalize and convert to runes
	runes := normalize(str)
	result[0] = runes[0]

	// find would-be replacement for first character
	last := replacements[runes[0]]

	for _, letter := range runes[1:] {
		sub := replacements[letter]
		if sub != 0 {
			if sub != last {
				result[count] = sub
				count++
			}
		}
		last = sub

		if count == 4 {
			break
		}
	}

	for i, ch := range result {
		if ch == 0 {
			result[i] = '0'
		}
	}

	return string(result[:])
}
