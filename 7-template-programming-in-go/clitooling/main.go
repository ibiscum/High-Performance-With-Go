package main

import (
	"fmt"
	"github/ibiscum/High-Performance-With-Go/7-template-programming-in-go/clitooling/cmd"
	"os"
)

func main() {

	if err := cmd.DateCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
