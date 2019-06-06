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
	split := strings.Split(line, "\t")

	if len(split) == 2 {
		return parsed{split[0], split[1]}, nil
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

// log & panic if anything unexpected happens
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
