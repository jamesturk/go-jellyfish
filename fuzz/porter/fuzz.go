package porter

import "github.com/jamesturk/go-jellyfish"

func Fuzz(data []byte) int {
	jellyfish.Porter(string(data))
	return 0
}
