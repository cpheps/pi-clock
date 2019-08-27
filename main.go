package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
)

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

func degreeToRadians(degrees int) float64 {
	return float64(degrees) * (math.Pi / 180.0)
}

func drawTick(img draw.Image, x0, y0, r, length int, angle float64, c color.Color) {
	x := math.Round(float64(x0) + float64(r)*math.Cos(angle))
	y := math.Round(float64(y0) + float64(r)*math.Sin(angle))

	fmt.Println("x", x, "y", y)

	intX, intY := int(x), int(y)
	img.Set(intX, intY, c)
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	drawCircle(img, 16, 16, 16, color.RGBA{255, 255, 0, 255})

	drawTick(img, 16, 16, 14, 2, degreeToRadians(0), color.RGBA{0, 255, 0, 255})
	drawTick(img, 16, 16, 14, 2, degreeToRadians(90), color.RGBA{0, 255, 0, 255})
	drawTick(img, 16, 16, 14, 2, degreeToRadians(180), color.RGBA{0, 255, 0, 255})
	drawTick(img, 16, 16, 14, 2, degreeToRadians(270), color.RGBA{0, 255, 0, 255})

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, img); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("circle.png", buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}
