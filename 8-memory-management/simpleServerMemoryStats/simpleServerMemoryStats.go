package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Memory Management Test")
		if err != nil {
			log.Fatal(err)
		}
	}

	go func() {
		for {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Println("\nTime: ", time.Now())
			fmt.Println("Runtime MemStats Sys: ", r.Sys)
			fmt.Println("Runtime Heap Allocation: ", r.HeapAlloc)
			fmt.Println("Runtime Heap Idle: ", r.HeapIdle)
			fmt.Println("Runtime Heap In Use: ", r.HeapInuse)
			fmt.Println("Runtime Heap HeapObjects: ", r.HeapObjects)
			fmt.Println("Runtime Heap Released: ", r.HeapReleased)

			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
