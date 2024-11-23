package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Memory Management Test")
		if err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}
