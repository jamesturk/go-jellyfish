package strfry

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
