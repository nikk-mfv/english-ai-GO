package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var jsonDataMyWords = `
{
	"words": [
		{
			"word": "hello",
			"definition": "a greeting"
		},
		{
			"word": "world",
			"definition": "the planet we live on"
		}
	]
}
`

func JSON(w http.ResponseWriter, req *http.Request) {
	// Allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, jsonDataMyWords)
}

func main() {
	http.HandleFunc("/api/v1/my-words", JSON)

	http.ListenAndServe(":8090", nil)
}
