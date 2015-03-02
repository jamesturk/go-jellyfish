package strfry

func Jaro(s1, s2 string) float64 {
	return jaroWinkler(s1, s2, false, false)
}

func JaroWinkler(s1, s2 string) float64 {
	return jaroWinkler(s1, s2, false, true)
}

func jaroWinkler(s1, s2 string, long_tolerance, winklerize bool) float64 {
	r1 := []rune(s1)
	r2 := []rune(s2)
	len1 := len(r1)
	len2 := len(r2)

	if len1 == 0 || len2 == 0 {
		return 0
	}

	min_len := max(len1, len2)
	search_range := (min_len / 2) - 1
	if search_range < 0 {
		search_range = 0
	}

	flags1 := make([]bool, len1)
	flags2 := make([]bool, len2)

	// look within search range for matched pairs
	common_chars := 0
	for i, ch := range s1 {
		low := 0
		if i > search_range {
			low = i - search_range
		}
		high := len2 - 1
		if i+search_range < len2 {
			high = i + search_range
		}
		for j := low; j <= high; j++ {
			if !flags2[j] && r2[j] == ch {
				flags1[i] = true
				flags2[j] = true
				common_chars++
				break
			}
		}
	}

	// short circuit if no characters match
	if common_chars == 0 {
		return 0
	}

	// count transpositions
	k := 0
	trans_count := 0
	for i, f1 := range flags1 {
		if f1 {
			j := k
			for ; j < len2; j++ {
				if flags2[j] {
					k = j + 1
					break
				}
			}
			if r1[i] != r2[j] {
				trans_count++
			}
		}
	}
	trans_count /= 2

	// adjust for similarities in nonmatched characters
	ccf := float64(common_chars)
	weight := (ccf/float64(len1) + ccf/float64(len2) + (ccf-float64(trans_count))/ccf) / 3

	// winkler modification: boost if strings are similar
	if winklerize && weight > 0.7 && len1 > 3 && len2 > 3 {
		j := min(min_len, 4)
		i := 0

		for i < j && r1[i] == r2[i] {
			i++
		}

		if i != 0 {
			weight += float64(i) * 0.1 * (1 - weight)
		}

		// TODO: add long_tolerance? optionally adjust for long strings
	}

	return weight
}
