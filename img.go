package main

import (
	"image"
	"image/color"
	"math/rand"
)

func getRandomImg() *image.RGBA {
	var rgb [3]uint8
	for i := 0; i < 3; i++ {
		rgb[i] = uint8(rand.Intn(256))
	}
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			img.Set(i, j, color.RGBA{rgb[0], rgb[1], rgb[2], 255})
		}
	}
	for i := 0; i < 3; {
		Circle{rand.Intn(500), rand.Intn(500), rand.Intn(20)}.Paint(img)
		Square{rand.Intn(500), rand.Intn(500), rand.Intn(20)}.Paint(img)
		i++
	}
	return img
}
