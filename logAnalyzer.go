package main

import (
	"sort"
)

// Implement countResult, with a sort by count interface
type countResult struct {
	Query string
	Count int
}
type byCount []countResult

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

// Main popular algo
func popular(size int, query string, data []parsed) []countResult {
	res := []countResult{}
	if size < 1 {
		return res
	}

	start := firstIndex(query, data)
	if start < 0 {
		return res
	}

	// count
	mapRes := make(map[string]int)
	for i := start; i < len(data) && matchQuery(data[i].timestamp, query); i++ {
		mapRes[data[i].search]++
	}

	// convert to slice
	for key, val := range mapRes {
		res = append(res, countResult{Query: key, Count: val})
	}

	// keep top n only
	sort.Sort(byCount(res))
	if len(res) > size {
		res = res[0:size]
	}

	return res
}

// Main distinct algo
func distinct(query string, data []parsed) int {

	// get first matching index
	start := firstIndex(query, data)
	if start < 0 {
		return 0
	}

	// accumulate in map and return final length
	keys := map[string]bool{}
	for i := start; i < len(data) && matchQuery(data[i].timestamp, query); i++ {
		keys[data[i].search] = true
	}

	return len(keys)
}
