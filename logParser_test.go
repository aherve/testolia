package main

import (
	"testing"
)

func TestFirstIndex(t *testing.T) {
	data := []parsed{
		{timestamp: "2015-08-01 00:03:46", search: "1"},
		{timestamp: "2016-08-01 00:03:48", search: "1"},
		{timestamp: "2016-08-01 01:03:48", search: "1"},
		{timestamp: "2018-08-01 00:03:49", search: "1"},
	}

	if r := firstIndex("2015", data); r != 0 {
		t.Error("expected firstIndex to equal 0, got ", r)
	}
	if r := firstIndex("2016-08-01", data); r != 1 {
		t.Error("expected firstIndex to equal 1, got ", r)
	}
	if r := firstIndex("2016-08-01 01:", data); r != 2 {
		t.Error("expected firstIndex to equal 2, got ", r)
	}
	if r := firstIndex("2018", data); r != 3 {
		t.Error("expected firstIndex to equal 3, got ", r)
	}
	if r := firstIndex("2019", data); r != -1 {
		t.Error("expected firstIndex to equal -1, got ", r)
	}
	if r := firstIndex("2018-08-01 00:03:49-wait-wat-please-stop", data); r != -1 {
		t.Error("expected firstIndex to equal -1, got ", r)
	}
	if r := firstIndex("2018", []parsed{}); r != -1 {
		t.Error("expected firstIndex to equal -1, got ", r)
	}
}

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
