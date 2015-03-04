package jellyfish

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"testing"
)

var lev_testdata [][]string
var dlev_testdata [][]string
var jaro_testdata [][]string
var jwink_testdata [][]string
var mrcodex_testdata [][]string
var mrcmp_testdata [][]string
var soundex_testdata [][]string
var hamming_testdata [][]string
var nysiis_testdata [][]string
var metaphone_testdata [][]string
var porter_testdata [][]string

func getTestdata(filename string) [][]string {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic("no test data file " + filename)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	testdata, err := reader.ReadAll()
	if err != nil {
		panic("error reading test data " + filename)
	}

	return testdata
}

func TestMain(m *testing.M) {
	lev_testdata = getTestdata("testdata/levenshtein.csv")
	dlev_testdata = getTestdata("testdata/damerau_levenshtein.csv")
	jaro_testdata = getTestdata("testdata/jaro_distance.csv")
	jwink_testdata = getTestdata("testdata/jaro_winkler.csv")
	mrcodex_testdata = getTestdata("testdata/match_rating_codex.csv")
	mrcmp_testdata = getTestdata("testdata/match_rating_comparison.csv")
	soundex_testdata = getTestdata("testdata/soundex.csv")
	hamming_testdata = getTestdata("testdata/hamming.csv")
	nysiis_testdata = getTestdata("testdata/nysiis.csv")
	metaphone_testdata = getTestdata("testdata/metaphone.csv")
	porter_testdata = getTestdata("testdata/porter.csv")
	os.Exit(m.Run())
}

func TestLevenshtein(t *testing.T) {
	for _, row := range lev_testdata {
		res := Levenshtein(row[0], row[1])
		expected, err := strconv.Atoi(row[2])
		if err != nil {
			t.Error("bad row in test data")
		}
		if res != expected {
			t.Errorf("Levenshtein(%q, %q) => %d, expected %d", row[0], row[1], res, expected)
		}
	}
}

func TestDamerauLevenshtein(t *testing.T) {
	for _, row := range dlev_testdata {
		res := DamerauLevenshtein(row[0], row[1])
		expected, err := strconv.Atoi(row[2])
		if err != nil {
			t.Error("bad row in test data")
		}
		if res != expected {
			t.Errorf("DamerauLevenshtein(%q, %q) => %d, expected %d", row[0], row[1], res, expected)
		}
	}
}

func TestJaro(t *testing.T) {
	for _, row := range jaro_testdata {
		res := Jaro(row[0], row[1])
		expected, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			t.Error("bad row in test data")
		}
		if math.Abs(res-expected) > 0.001 {
			t.Errorf("Jaro(%q, %q) => %.3f, expected %.3f", row[0], row[1], res, expected)
		}
	}
}

func TestJaroWinkler(t *testing.T) {
	for _, row := range jwink_testdata {
		res := JaroWinkler(row[0], row[1])
		expected, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			t.Error("bad row in test data")
		}
		if math.Abs(res-expected) > 0.001 {
			t.Errorf("JaroWinkler(%q, %q) => %.3f, expected %.3f", row[0], row[1], res, expected)
		}
	}
}

func TestMatchRatingCodex(t *testing.T) {
	for _, row := range mrcodex_testdata {
		res := MatchRatingCodex(row[0])

		if res != row[1] {
			t.Errorf("MatchRatingCodex(%q) => %q, expected %q", row[0], res, row[1])
		}
	}
}

func TestMatchRatingComparison(t *testing.T) {
	for _, row := range mrcmp_testdata {
		res := MatchRatingComparison(row[0], row[1])
		expected := (row[2] == "True")

		if res != expected {
			t.Errorf("MatchRatingCodex(%q, %q) => %t, expected %t", row[0], row[1], res, expected)
		}
	}
}

func TestSoundex(t *testing.T) {
	for _, row := range soundex_testdata {
		res := Soundex(row[0])

		if res != row[1] {
			t.Errorf("Soundex(%q) => %q, expected %q", row[0], res, row[1])
		}
	}
}

func TestHamming(t *testing.T) {
	for _, row := range hamming_testdata {
		res := Hamming(row[0], row[1])
		expected, err := strconv.Atoi(row[2])
		if err != nil {
			t.Error("bad row in test data")
		}
		if res != expected {
			t.Errorf("Hamming(%q, %q) => %d, expected %d", row[0], row[1], res, expected)
		}
	}
}

func TestNysiis(t *testing.T) {
	for _, row := range nysiis_testdata {
		res := Nysiis(row[0])
		if res != row[1] {
			t.Errorf("Nysiis(%q) => %q, expected %q", row[0], res, row[1])
		}
	}
}

func TestMetaphone(t *testing.T) {
	for _, row := range metaphone_testdata {
		res := Metaphone(row[0])
		if res != row[1] {
			t.Errorf("Metaphone(%q) => %q, expected %q", row[0], res, row[1])
		}
	}
}

func TestPorter(t *testing.T) {
	wrong := 0
	total := 0

	for _, row := range porter_testdata {
		res := Porter(row[0])
		if res != row[1] {
			t.Errorf("Porter(%q) => %q, expected %q", row[0], res, row[1])
			wrong++
		}
		total++
	}
	if wrong > 0 {
		t.Errorf("%d / %d incorrect", wrong, total)
	}
}
