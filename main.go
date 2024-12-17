package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	demoHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}

	http.HandleFunc("/hello", demoHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
