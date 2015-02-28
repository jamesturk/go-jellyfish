package strfry

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func Levenshtein(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}

	rows := len(s1) + 1
	cols := len(s2) + 1

	if s1 == "" {
		return cols - 1
	}
	if s2 == "" {
		return rows - 1
	}

	// intiialize the first row
	prev := make([]int, cols)
	cur := make([]int, cols)
	for i := range cur {
		cur[i] = i
	}

	for r := 1; r < rows; r++ {
		prev[0] = cur[0]
		cur[0] = r
		for c := 1; c < cols; c++ {
			prev[c] = cur[c]
			deletion := prev[c] + 1
			insertion := cur[c-1] + 1
			edit := prev[c-1]
			if s1[r-1] != s2[c-1] {
				edit += 1
			}
			cur[c] = min3(edit, deletion, insertion)
		}
	}

	return cur[len(cur)-1]
}
