package main

import "fmt"

func helloGo() {
	fmt.Println("Hello Go from function helloGo")
}

func main() {
	helloGo()

	func() {
		fmt.Println("Hello Go from an anonymous function")
	}()

	var hello func() = func() {
		fmt.Println("Hello Go from an anonymous function variable")
	}

	hello()
}
