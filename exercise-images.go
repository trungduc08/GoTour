package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}
func (img Image) At(x, y int) color.Color {
	return color.RGBA{(uint8)(x + y), (uint8)(x + y), 255, 255}
}
func main() {
	img := Image{2, 2}
	img.Bounds()
	s := make([][]color.Color, 2)
	for i := 0; i < 2; i++ {
		s[i] = make([]color.Color, 2)

	}
	for i := img.Bounds().Min.X; i < img.Bounds().Max.X; i++ {
		for j := img.Bounds().Min.Y; j < img.Bounds().Max.Y; j++ {
			img.At(i, j)
			// fmt.Println(img.At(i, j))
			s[i][j] = img.At(i, j)

		}
	}
	fmt.Println(s)

	// for _, v := range s {
	// 	fmt.Println(v)
	// }
}
