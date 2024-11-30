package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		sleep(5)
		sleep(10)
		fmt.Fprintf(w, "Sleep Profiling Test")
	}
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sleep(sleepTime int) {
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("Slept for ", sleepTime, " Milliseconds")
}
