package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	// Print HTTP status
	fmt.Println(resp.Status)

	// Read and display response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
