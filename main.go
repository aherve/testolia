package main

import (
	"log"
	"net/http"
	"os"
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
