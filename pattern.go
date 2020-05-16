package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

type Figure interface {
	Paint(*image.RGBA)
}

type Square struct {
	X, Y, R int
}

func (s Square) Paint(img *image.RGBA) {
	var cl [3]uint8
	for i := 0; i < 3; i++ {
		cl[i] = uint8(rand.Intn(256))
	}
	s.R *= 10
	for i := s.X - s.R; i < s.X+s.R; i++ {
		for j := s.Y - s.R; j < s.Y+s.R; j++ {
			img.Set(i, j, color.RGBA{cl[0], cl[1], cl[2], 255})
		}
	}
}

type Circle struct {
	X, Y, R int
}

func (c Circle) Paint(img *image.RGBA) {
	var cl [3]uint8
	for i := 0; i < 3; i++ {
		cl[i] = uint8(rand.Intn(256))
	}
	c.R *= 10
	for i := c.X - c.R; i < c.X+c.R; i++ {
		for j := c.Y - c.R; j < c.Y+c.R; j++ {
			if 1 > math.Sqrt(float64((c.X-i)*(c.X-i)+(c.Y-j)*(c.Y-j)))/float64(c.R) {
				img.Set(i, j, color.RGBA{cl[0], cl[1], cl[2], 255})
			}
		}
	}
}
