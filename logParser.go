package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"log"
	"os"
	"sort"
	"strings"
)

// parsed line structure with sort by timestamp interface
// ISO dates provided => string sort is equivalent to a date sort
type parsed struct {
	timestamp string
	search    string
}

type byTimestamp []parsed

func (a byTimestamp) Len() int           { return len(a) }
func (a byTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byTimestamp) Less(i, j int) bool { return a[i].timestamp < a[j].timestamp }

// Parse line to struct
func parse(line string) (parsed, error) {
	splitted := strings.Split(line, "\t")

	if len(splitted) == 2 {
		return parsed{splitted[0], splitted[1]}, nil
	} else {
		return parsed{}, errors.New("could not parse line")
	}

}

// Slurp file, parse lines and sort result by timestamp
// This function can either read a gzip, or plain tsv file
func readAndSort(filename string) []parsed {
	var scanner *bufio.Scanner

	file, err := os.Open(filename)
	handleError(err)
	defer file.Close()

	// handle gzip
	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(file)
		defer gz.Close()

		handleError(err)
		scanner = bufio.NewScanner(gz)
	} else {
		// handle regular file
		scanner = bufio.NewScanner(file)
	}

	res := []parsed{}
	for scanner.Scan() {
		parsed, err := parse(scanner.Text())
		handleError(err)
		res = append(res, parsed)
	}

	// sort by date
	sort.Sort(byTimestamp(res))
	return res
}

// firstIndex uses binary sort to find the first index that matches a search in a sorted dataset
func firstIndex(query string, data []parsed) int {
	// safety
	if len(data) == 0 {
		return -1
	}

	// matching function
	f := func(i int) bool {
		ts := data[i].timestamp
		return len(query) <= len(ts) && ts[0:len(query)] >= query
	}

	// binary search
	i := sort.Search(len(data), f)

	// it's up to the caller to check wether `sort.Search` actually found an element
	if i < len(data) && f(i) {
		return i
	} else {
		return -1
	}
}

func matchQuery(ts, query string) bool {
	return strings.HasPrefix(ts, query)
}

// log & panic if anything unexpected happens
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
