package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

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

type Image struct{
	width, height int
	bitmap [][]uint8
}

func newImage(width int, height int, f func(int, int) [][]uint8) *Image {
	var ret *Image = new(Image)
	ret.width = width
	ret.height = height
	ret.bitmap = f(width, height)
	
	return ret
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.width, self.height)
}

func (self Image) At(x, y int) color.Color {
	val := self.bitmap[x][y]
	return color.RGBA{(val * 2) % 255, (val / 2), val + 25, (val + 50) % 255}
}

func main() {
	m := newImage(200, 200, Fancy)
	pic.ShowImage(m)
}
