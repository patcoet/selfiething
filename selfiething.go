package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

type filter func(uint8, uint8, uint8) (uint8, uint8, uint8)

func filter1(r, g, b uint8) (uint8, uint8, uint8) {
	if (r+g+b)%3 == 0 {
		r = (g * b) % 255
		g = (r * b) % 255
		b = (r * g) % 255
	}

	return r, g, b
}

func filter2(r, g, b uint8) (uint8, uint8, uint8) {
	if r > g {
		temp := r
		r = g
		g = temp
	}
	return r, g, b
}

func main() {
	// applyFilter(filter2)
	arg := os.Args[1]

	filters := []filter{filter1, filter2}

	n, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Do the file opening and writing stuff in main and have
	// applyFilter only just apply the filter. Should probably use pointers
	// to avoid making multiple copies of the files, if that's how Go
	// works...?
	applyFilter(filters[n-1])
}

func applyFilter(f filter) {
	file, err := os.Open("img.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	img2 := image.NewRGBA(img.Bounds())

	for xPos := img.Bounds().Min.X; xPos < img.Bounds().Max.X; xPos++ {
		for yPos := img.Bounds().Min.Y; yPos < img.Bounds().Max.Y; yPos++ {
			rr, gg, bb, aa := img.At(xPos, yPos).RGBA()
			r, g, b, a := uint8(rr), uint8(gg), uint8(bb), uint8(aa)

			r, g, b = f(r, g, b)

			img2.SetRGBA(xPos, yPos, color.RGBA{r, g, b, a})
		}
	}

	file2, err := os.Create("img2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	err = png.Encode(file2, img2)
	if err != nil {
		log.Fatal(err)
	}
}
