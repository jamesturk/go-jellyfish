package jellyfish

// Levenshtein computes the Levenshtein distance between two strings.
//
// Levenshtein distance represents the number of insertions, deletions,
// and subsititutions required to change one word to another.
//
// For example:
//    Levenshtein("berne", "born") == 2
// representing the transformation of the first e to o and the
// deletion of the second e.
//
// See the Levenshtein distance article at Wikipedia (http://en.wikipedia.org/wiki/Levenshtein_distance) for more details.
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
			cur[c] = min(min(edit, deletion), insertion)
		}
	}

	return cur[len(cur)-1]
}

// DamerauLevenshtein computes the Damerau-Levenshtein distance between two strings.
//
// A modification of Levenshtein distance, Damerau-Levenshtein distance counts
// the number of edits (insertions, deletions, and substitutions) but unlike
// Levenshtein, considers transpositions (such as ifhs for fish) a single edit.
//
// For example:
//    Levenshtein("fish", "ifsh") == 2          // one deletion, one insertion
//                                              // but...
//    DamerauLevenshtein("fish", "ifsh") == 1   // one transposition
//
// See the Damerau-Levenshtein distance article at Wikipedia
// (http://en.wikipedia.org/wiki/Damerau-Levenshtein_distance) for more details.
func DamerauLevenshtein(s1, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)
	len1 := len(r1)
	len2 := len(r2)
	infinite := len1 + len2

	da := make(map[rune]int)
	score := make([][]int, len1+2)
	for i := range score {
		score[i] = make([]int, len2+2)
	}
	// initialize scores
	score[0][0] = infinite
	for i := 0; i < len2+1; i++ {
		score[0][i+1] = infinite
		score[1][i+1] = i
	}
	for i := 0; i < len1+1; i++ {
		score[i+1][0] = infinite
		score[i+1][1] = i
	}

	for i := 1; i < len1+1; i++ {
		db := 0
		for j := 1; j < len2+1; j++ {
			i1 := da[r2[j-1]]
			j1 := db
			cost := 1
			if r1[i-1] == r2[j-1] {
				cost = 0
				db = j
			}

			score[i+1][j+1] = min(min(score[i][j]+cost, score[i+1][j]+1), min(score[i][j+1]+1, score[i1][j1]+(i-i1-1)+1+(j-j1-1)))
		}

		da[r1[i-1]] = i
	}
	return score[len1+1][len2+1]
}
