package jellyfish

import "testing"

var int_result int
var float_result float64
var str_result string
var bool_result bool

func BenchmarkLevenshtein(b *testing.B) {
	var res int

	for n := 0; n < b.N; n++ {
		for _, row := range lev_testdata {
			res = Levenshtein(row[0], row[1])
		}
	}
	int_result = res
}

func BenchmarkDamerauLevenshtein(b *testing.B) {
	var res int

	for n := 0; n < b.N; n++ {
		for _, row := range dlev_testdata {
			res = DamerauLevenshtein(row[0], row[1])
		}
	}
	int_result = res
}

func BenchmarkJaro(b *testing.B) {
	var res float64

	for n := 0; n < b.N; n++ {
		for _, row := range jaro_testdata {
			res = Jaro(row[0], row[1])
		}
	}
	float_result = res
}

func BenchmarkJaroWinkler(b *testing.B) {
	var res float64

	for n := 0; n < b.N; n++ {
		for _, row := range jwink_testdata {
			res = JaroWinkler(row[0], row[1])
		}
	}
	float_result = res
}

func BenchmarkMatchRatingCodex(b *testing.B) {
	var res string

	for n := 0; n < b.N; n++ {
		for _, row := range mrcodex_testdata {
			res = MatchRatingCodex(row[0])
		}
	}
	str_result = res
}

func BenchmarkMatchRatingComparison(b *testing.B) {
	var res bool

	for n := 0; n < b.N; n++ {
		for _, row := range mrcmp_testdata {
			res = MatchRatingComparison(row[0], row[1])
		}
	}

	bool_result = res
}

func BenchmarkSoundex(b *testing.B) {
	var res string

	for n := 0; n < b.N; n++ {
		for _, row := range soundex_testdata {
			res = Soundex(row[0])
		}
	}

	str_result = res
}

func BenchmarkHamming(b *testing.B) {
	var res int

	for n := 0; n < b.N; n++ {
		for _, row := range hamming_testdata {
			res = Hamming(row[0], row[1])
		}
	}

	int_result = res
}

func BenchmarkNysiis(b *testing.B) {
	var res string

	for n := 0; n < b.N; n++ {
		for _, row := range nysiis_testdata {
			res = Nysiis(row[0])
		}
	}
	str_result = res
}

func BenchmarkMetaphone(b *testing.B) {
	var res string

	for n := 0; n < b.N; n++ {
		for _, row := range metaphone_testdata {
			res = Metaphone(row[0])
		}
	}
	str_result = res
}

func BenchmarkPorter(b *testing.B) {
	var res string

	for n := 0; n < b.N; n++ {
		for _, row := range porter_testdata {
			res = Porter(row[0])
		}
	}
	str_result = res
}
