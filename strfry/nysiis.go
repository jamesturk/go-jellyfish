package strfry

import "strings"

func Nysiis(s string) string {
	var key []rune
	runes := []rune(strings.ToUpper(s))
	rlen := len(runes)

	// step 1 - prefixes
	switch {
	case runes[0] == 'M' && runes[1] == 'A' && runes[2] == 'C':
		runes[1] = 'C'
	case runes[0] == 'K' && runes[2] == 'N':
		runes = runes[1:]
	case runes[0] == 'K':
		runes[0] = 'C'
	case runes[0] == 'P' && (runes[1] == 'H' || runes[1] == 'F'):
		runes[0] = 'F'
		runes[1] = 'F'
	case runes[0] == 'S' && runes[1] == 'C' && runes[2] == 'H':
		runes[1] = 'S'
		runes[2] = 'S'
	}

	// step 2 - suffixes
	switch {
	case (runes[rlen-2] == 'I' || runes[rlen-2] == 'E') && runes[rlen-1] == 'E':
		runes = append(runes[:rlen-2], 'Y')
	case runes[rlen-2] == 'D' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'R' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'R' && runes[rlen-1] == 'D',
		runes[rlen-2] == 'N' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'N' && runes[rlen-1] == 'D':
		runes = append(runes[:rlen-2], 'D')
	}

	// step 3 - first character from name
	key = append(key, runes[0])

	// step 4 - translate remaining
	//for i, ch := range runes[1:] {

	//}

	return string(key)
}
