package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

type Words struct {
	Page  string   `json:"page"` //`json:"page"` this will parse input json format into struct type
	Input string   `json:"input"`
	Words []string `json:"words"`
}

// Checking HTTP Requests using input arguments

func main() {

	// Print OS Info
	osInfo := runtime.GOOS
	fmt.Printf("Runtime OS: %v\n", osInfo)

	args := os.Args

	if len(args) < 2 {

		fmt.Printf("Usage: ./main <url>\n")
		os.Exit(1)
	}

	// myurl, err := url.ParseRequestURI(args[1])   :::: Could be like this if we want the URL value returned

	if _, err := url.ParseRequestURI(args[1]); err != nil {

		fmt.Printf("The URL is invalid: %s\n", err)

	}
	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)

	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)

	}

	if response.StatusCode != 200 {

		fmt.Printf("HTTP Status Code: %d\n Whole Body: %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("json parsed\nPage: %s\nWords: %v\n", words.Page, words.Words)
}
