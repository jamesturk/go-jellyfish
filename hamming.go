package jellyfish

// Hamming computes the Hamming distance between s1 and s2.
//
// Hamming distance is the number of characters that differ between two strings.
//
// Typically Hamming distance is undefined when strings are of different lengths,
// this implementation considers extra characters as differing.  Thus Hamming("abc", "abcd") == 1
//
// See the Hamming distance article at Wikipedia (http://en.wikipedia.org/wiki/Hamming_distance) for more details.
func Hamming(s1, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r2) > len(r1) {
		r1, r2 = r2, r1
	}

	distance := len(r1) - len(r2)

	for i, c := range r2 {
		if c != r1[i] {
			distance += 1
		}
	}

	return distance
}
