package main

import (
	"fmt"
	"os"

	"github.com/ibiscum/High-Performance-With-Go/7-template-programming-in-go/clitooling/cmd"
)

func main() {

	if err := cmd.DateCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
