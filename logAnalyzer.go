package main

import (
	"sort"
	"strings"
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

// firstIndex uses binary sort to find the first index that matches a search in a sorted dataset
func firstIndex(query string, data []parsed) int {
	// safety
	if len(data) == 0 {
		return -1
	}

	// comparison function
	f := func(i int) bool {
		ts := data[i].timestamp
		return len(query) <= len(ts) && ts[0:len(query)] >= query
	}

	// binary search
	i := sort.Search(len(data), f)

	// it's up to the caller to check wether `sort.Search` actually found an element
	if i < len(data) && matchQuery(data[i].timestamp, query) {
		return i
	} else {
		return -1
	}
}

func matchQuery(ts, query string) bool {
	return strings.HasPrefix(ts, query)
}
