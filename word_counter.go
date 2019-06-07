package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map [string] int);
	words := strings.Split(s, " ");
	for _, v := range words {
		count[v]++;	
	}
	
	return count
}

func main() {
	wc.Test(WordCount)
}

