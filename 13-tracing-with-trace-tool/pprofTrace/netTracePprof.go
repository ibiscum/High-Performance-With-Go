package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(5 * time.Second)
		_, err := io.WriteString(w, "Network Trace Profile Test")
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
