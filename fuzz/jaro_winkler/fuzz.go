package jaro_winkler

import "github.com/jamesturk/go-jellyfish"

//import "math/rand"
//import "time"

func Fuzz(data []byte) int {
	jellyfish.JaroWinkler(string(data), string(data))
	/*if len(data) > 2 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(data))
		jellyfish.Jaro(string(data[:n]), string(data[n:]))
	}*/
	return 0
}
