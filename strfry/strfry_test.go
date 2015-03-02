package strfry

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"testing"
)

func getTestdata(filename string, t *testing.T) [][]string {
	csvfile, err := os.Open(filename)
	if err != nil {
		t.Error("no test data file")
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	testdata, err := reader.ReadAll()
	if err != nil {
		t.Error("error reading test data")
	}

	return testdata
}

func TestLevenshtein(t *testing.T) {
	testdata := getTestdata("testdata/levenshtein.csv", t)

	for _, row := range testdata {
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
	testdata := getTestdata("testdata/damerau_levenshtein.csv", t)

	for _, row := range testdata {
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
	testdata := getTestdata("testdata/jaro_distance.csv", t)

	for _, row := range testdata {
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
	testdata := getTestdata("testdata/jaro_winkler.csv", t)

	for _, row := range testdata {
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
	testdata := getTestdata("testdata/match_rating_codex.csv", t)

	for _, row := range testdata {
		res := MatchRatingCodex(row[0])

		if res != row[1] {
			t.Errorf("MatchRatingCodex(%q) => %q, expected %q", row[0], res, row[1])
		}
	}
}
