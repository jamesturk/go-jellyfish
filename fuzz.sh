go get github.com/dvyukov/go-fuzz/...
cd fuzz
go-fuzz-build github.com/jamesturk/go-jellyfish/fuzz/nysiis
go-fuzz -bin=./nysiis-fuzz -corpus=./nysiis-corpus -workdir=./nysiis-workdir
