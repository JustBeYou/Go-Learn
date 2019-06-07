package main

import (
	"fmt"
)

func main() {
	fmt.Println("buna")

	// this thing gets executed last
	defer fmt.Println("salut baieti") 

	fmt.Println("ceau")
}
