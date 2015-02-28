package strfry

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	csvfile, err := os.Open("testdata/levenshtein.csv")
	if err != nil {
		t.Error("no test data file")
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	testdata, err := reader.ReadAll()
	if err != nil {
		t.Error("error reading test data")
	}

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
