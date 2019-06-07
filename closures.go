package main

import "fmt"

// closures = something like static variables I think
func adderFactory() func (int) int {
	var value int = 0
	return func (x int) int {
		value += x
		return value
	}
}

func main() {
	adder1, adder2 := adderFactory(), adderFactory()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %d\n", adder1(1), adder2(-1));
	}
}
