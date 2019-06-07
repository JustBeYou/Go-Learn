package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	term_1, term_2 := 0, 1
	return func () int {
		to_return := term_1
		
		term_1 = term_2
		term_2 = term_2 + to_return
	
		return to_return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

