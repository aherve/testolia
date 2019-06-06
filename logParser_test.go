package main

import (
	"testing"
)

// Read an actual test file
func TestReadAndSort(t *testing.T) {
	res := readAndSort("datatest/test.tsv")

	if len(res) != 4 {
		t.Error("expected test.tsv to yield 4 values")
	}

	if r := res[0].search; r != "first" {
		t.Error("expected first key to eql 'first' but got ", r)
	}

	if r := res[3].search; r != "bar" {
		t.Error("expected first key to eql 'bar' but got ", r)
	}

}

func TestParse(t *testing.T) {
	line := "2015-08-01 00:03:46	nixos"

	parsed, err := parse(line)
	if err != nil {
		t.Error("expected", err, " to be nil")
	}
	if r := parsed.search; r != "nixos" {
		t.Error("expected parse to yield query = nixos, got ", r)
	}
	if r := parsed.timestamp; r != "2015-08-01 00:03:46" {
		t.Error("expected parse to yield query = 2015-08-01 00:03:46, got ", r)
	}
}
