package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
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
			gc := new(debug.GCStats)
			fmt.Println("\nTime: ", time.Now())
			fmt.Println("Last Garbage Collection: ", gc.LastGC)
			fmt.Println("Number of Garbage Collections: ", gc.NumGC)
			fmt.Println("Total Pause for all Collections: ", gc.PauseTotal)
			fmt.Println("Pause: ", gc.Pause)
			fmt.Println("PauseEnd: ", gc.PauseEnd)
			fmt.Println("PauseQuantiles: ", gc.PauseQuantiles)
			debug.FreeOSMemory()
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
