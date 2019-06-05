package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a filename to read as first argument")
	}
	filename := os.Args[1]
	data := readAndSort(filename)
	log.Println("read and sorted ", len(data), " lines")

	// define http handlers using closure
	http.HandleFunc("/1/queries/popular/", handlePopular(data))
	http.HandleFunc("/1/queries/count/", handleDistinct(data))

	// start server
	log.Println("starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

type distinctHttpResp struct {
	Count int
}

func handleDistinct(data []parsed) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Println(r.Method, r.URL) // basic logging
		// get query
		query := getQuery(r)

		// get results
		distinct := distinct(query, data)

		// format & send
		js, err := json.Marshal(distinctHttpResp{distinct})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func handlePopular(data []parsed) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Println(r.Method, r.URL) // basic logging

		// size parameter is mandatory
		size, err := getSize(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		// get query
		query := getQuery(r)

		// get data
		popular := popular(size, query, data)

		// format and send
		js, err := json.Marshal(popular)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// We choose to default to "" (i.e. all) if no proper query is found
func getQuery(r *http.Request) string {
	splitted := strings.Split(r.URL.Path, "/")
	if len(splitted) >= 5 {
		return splitted[4]
	} else {
		return ""
	}
}

// find size or error
func getSize(r *http.Request) (int, error) {
	if s, ok := r.URL.Query()["size"]; ok && len(s) == 1 {
		size, err := strconv.Atoi(s[0])
		return size, err
	} else {
		return -1, errors.New("missing size parameter")
	}
}
