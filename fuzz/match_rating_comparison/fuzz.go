package match_rating_comparison

import "github.com/jamesturk/go-jellyfish"

//import "math/rand"
//import "time"

func Fuzz(data []byte) int {
	jellyfish.MatchRatingComparison(string(data), string(data))
	/*if len(data) > 2 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(data))
		jellyfish.Jaro(string(data[:n]), string(data[n:]))
	}*/
	return 0
}
