package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot square negative number %f", float64(e))	
}

func Sqrt(x float64) (float64, error){
	if x == 1 || x == 0 {
		return x, nil
	} else if x < 0 {
		return 0, ErrNegativeSqrt(x)	
	}
	
	const epsilon float64 = 1e-10
	var left, right float64 = 1, x
	var ans float64 = 0 

	for math.Abs(ans * ans - x) > epsilon {
		ans = (left + right) / 2	
		//fmt.Printf("%f - %f - %f\n", left, ans, right)
		if ans * ans > x {
			right = ans - epsilon	
		} else if ans * ans < x {
			left = ans + epsilon
		}
	}
	
	return ans, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
