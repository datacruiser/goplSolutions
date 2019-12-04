package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		x := parseFirstFloat64OrDefault(r.Form["x"], 0)
		y := parseFirstFloat64OrDefault(r.Form["y"], 0)
		zoom := parseFirstFloat64OrDefault(r.Form["zoom"], 0)

		renderPNG(w, x, y, zoom)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func parseFirstFloat64OrDefault(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}

	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}

	return value
}

func renderPNG(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, acos(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}
