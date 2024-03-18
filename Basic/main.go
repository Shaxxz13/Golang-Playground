package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

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
	response, err = http.Get(args[1])

	if err != nil {
		log.Fatal(err)

	}

	defer response.Body.close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)

	}

}
