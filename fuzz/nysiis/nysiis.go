package nysiis

import "github.com/jamesturk/go-jellyfish"

func Fuzz(data []byte) int {
	jellyfish.Nysiis(string(data))
	return 0
}
