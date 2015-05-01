package metaphone

import "github.com/jamesturk/go-jellyfish"

func Fuzz(data []byte) int {
	jellyfish.Metaphone(string(data))
	return 0
}
