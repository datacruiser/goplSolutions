// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
/*
Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation
by computing the color value at several points within each pixel and taking the average.
The simplest method is to divide each pixel into four ‘‘subpixels.’’ Implement it.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xf4, 0x43, 0x36, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0x4f, 0xaf, 0x50, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0x22, 0x98, 0xf3, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0x00},
	color.RGBA{0x00, 0xff, 0x00, 0x00},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			z0 := complex(x0, y0)
			z1 := complex(x1, y0)
			z2 := complex(x0, y1)
			z3 := complex(x1, y1)

			color := GetAverageColor([]color.Color{
				mandelbrot(z0),
				mandelbrot(z1),
				mandelbrot(z2),
				mandelbrot(z3),
			})
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func GetAverageColor(colors []color.Color) color.Color {
	if len(colors) < 1 {
		return nil
	}

	var r, g, b, a float64

	for _, cl := range colors {
		dr, dg, db, da := cl.RGBA()
		r += float64(dr>>8) / float64(len(colors))
		g += float64(dg>>8) / float64(len(colors))
		b += float64(db>>8) / float64(len(colors))
		a += float64(da>>8) / float64(len(colors))
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
