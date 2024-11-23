package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	zapLogger := zap.NewExample()
	defer func() {
		err := zapLogger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}()
	zapLogger.Debug("Hi Gophers - from our Zap Logger")
}
