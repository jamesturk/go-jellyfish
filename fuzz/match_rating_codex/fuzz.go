package match_rating_codex

import "github.com/jamesturk/go-jellyfish"

func Fuzz(data []byte) int {
	jellyfish.MatchRatingCodex(string(data))
	return 0
}
