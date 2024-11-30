package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, r *http.Request) {
		sleepDuration := r.URL.Query().Get("time")
		sleepDurationInt, err := strconv.Atoi(sleepDuration)
		if err != nil {
			fmt.Println("Incorrect value passed as a query string for time")
			return
		}
		sleep(sleepDurationInt)
		fmt.Fprintf(w, "Slept for %v Milliseconds", sleepDuration)
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
