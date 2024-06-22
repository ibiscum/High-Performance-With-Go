package main

import "fmt"

func main() {

	var tmp = []int{1, 2, 3}
	b := tmp[1:]
	fmt.Println(b)
	for i := range tmp {
		fmt.Println(tmp[i])
	}

}
