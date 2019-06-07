package main

import "golang.org/x/tour/pic"

func Graph(dx, dy int, f func(int, int) int) [][]uint8 {
	grayscale := make([][]uint8, dy)
	
	for y := 0; y < dy; y++ {
		grayscale[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			grayscale[y][x]	= uint8(f(int(x), int(y)))
		}
	}
	
	return grayscale
}

func Fancy(dx, dy int) [][]uint8 {
	return Graph(dx, dy, func (x, y int) (r int) {
		r = (x * y - x ^ y) / 2
		return 
	})	
}

func main() {
	pic.Show(Fancy)
}
