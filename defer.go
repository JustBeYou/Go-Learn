package main

import (
	"fmt"
)

func main() {
	fmt.Println("buna")

	// this thing gets executed last
	defer fmt.Println("salut baieti") 

	fmt.Println("ceau")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
