package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/leak", leakyAbstraction)
	err := http.ListenAndServe("localhost:6060", nil)
	log.Fatal(err)
}

func leakyAbstraction(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string, 15)
	for {
		fmt.Fprintln(w, "Number of Goroutines: ", runtime.NumGoroutine())
		go func() { ch <- wait() }()
	}
}

func wait() string {
	time.Sleep(5 * time.Microsecond)
	return "Hello Gophers!"
}
