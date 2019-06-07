package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	if x == 1 {
		return 1, 0	
	}
	
	const epsilon float64 = 1e-10
	var left, right float64 = 1, x
	var ans float64 = 0 

	var steps int = 0
	for math.Abs(ans * ans - x) > epsilon {
		ans = (left + right) / 2	
		//fmt.Printf("%f - %f - %f\n", left, ans, right)
		if ans * ans > x {
			right = ans - epsilon	
		} else if ans * ans < x {
			left = ans + epsilon
		}
		steps++
	}
	
	return ans, steps
}

func main() {
	const val = 999
	ans, steps := Sqrt(val)
	fmt.Printf("My SQRT: %f in %d steps\n", ans, steps)
	fmt.Printf("Math Lib SQRT: %f\n", math.Sqrt(val))
}

