package soundex

import "github.com/jamesturk/go-jellyfish"

func Fuzz(data []byte) int {
	jellyfish.Soundex(string(data))
	return 0
}
