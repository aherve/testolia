package main

import (
	"testing"
)

func TestPopular(t *testing.T) {
	data := []parsed{
		{timestamp: "2016", query: "a"},
		{timestamp: "2016", query: "b"},
		{timestamp: "2016", query: "c"},
		{timestamp: "2016", query: "a"},
		{timestamp: "2016", query: "b"},
		{timestamp: "2016", query: "d"},
		{timestamp: "2016", query: "a"},
		{timestamp: "2018", query: "e"},
		{timestamp: "2018", query: "e"},
		{timestamp: "2018", query: "e"},
		{timestamp: "2018", query: "e"},
		{timestamp: "2018", query: "e"},
	}

	// find 0 value
	p := popular(1, "2017", data)
	if r := len(p); r != 0 {
		t.Error("expected r to have length of 0, got ", r)
	}

	// find 1 value
	p = popular(1, "2016", data)
	if r := len(p); r != 1 {
		t.Error("expected r to have length of 1, got ", r)
	}
	if r := p[0]; r.Count != 3 {
		t.Error("expected r[0] to have count of 3, got ", r.Count)
	}
	if r := p[0]; r.Query != "a" {
		t.Error("expected r[0] to have query 'a', got ", r.Query)
	}

	// find 2 values
	p = popular(2, "2016", data)
	if r := len(p); r != 2 {
		t.Error("expected r to have length of 2, got ", r)
	}
	if r := p[0]; r.Count != 3 {
		t.Error("expected r[0] to have count of 3, got ", r.Count)
	}
	if r := p[0]; r.Query != "a" {
		t.Error("expected r[0] to have query 'a', got ", r.Query)
	}
	if r := p[1]; r.Count != 2 {
		t.Error("expected r[0] to have count of 3, got ", r.Count)
	}
	if r := p[1]; r.Query != "b" {
		t.Error("expected r[0] to have query 'b', got ", r.Query)
	}

	// find 2 values with different filter
	p = popular(2, "2", data)
	if r := len(p); r != 2 {
		t.Error("expected r to have length of 2, got ", r)
	}
	if r := p[1]; r.Count != 3 {
		t.Error("expected r[1] to have count of 3, got ", r.Count)
	}
	if r := p[1]; r.Query != "a" {
		t.Error("expected r[1] to have query 'a', got ", r.Query)
	}
	if r := p[0]; r.Count != 5 {
		t.Error("expected r[0] to have count of 5, got ", r.Count)
	}
	if r := p[0]; r.Query != "e" {
		t.Error("expected r[0] to have query 'e', got ", r.Query)
	}

	// try to find too much values
	p = popular(20, "2", data)
	if r := len(p); r != 5 {
		t.Error("expected r to have length of 5, got ", r)
	}

	// empty data
	p = popular(2, "2", []parsed{})
	if r := len(p); r != 0 {
		t.Error("expected r to have length of 0, got ", r)
	}

}

func TestDistinct(t *testing.T) {
	data := []parsed{
		{timestamp: "2015-08-01 00:03:46", query: "a"},
		{timestamp: "2016-08-01 00:03:48", query: "b"},
		{timestamp: "2016-08-01 01:03:48", query: "a"},
		{timestamp: "2018-08-01 00:03:49", query: "c"},
	}

	if r := distinct("20", data); r != 3 {
		t.Error("expected distinct(20) to return 3, got ", r)
	}

	if r := distinct("2016", data); r != 2 {
		t.Error("expected distinct(2016) to return 2, got ", r)
	}

	if r := distinct("", data); r != 3 {
		t.Error("expected distinct('') to return 3, got", r)
	}

	// empty data
	if r := distinct("", []parsed{}); r != 0 {
		t.Error("expected distinct('', []) to return 0, got", r)
	}
}

func TestFirstIndex(t *testing.T) {
	data := []parsed{
		{timestamp: "2015-08-01 00:03:46", query: "1"},
		{timestamp: "2016-08-01 00:03:48", query: "1"},
		{timestamp: "2016-08-01 01:03:48", query: "1"},
		{timestamp: "2018-08-01 00:03:49", query: "1"},
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
