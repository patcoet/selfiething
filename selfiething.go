package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
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
			r, g, b, a := img.At(xPos, yPos).RGBA()

			if (r+g+b)%3 == 0 {
				r = (g * b) % 65535
				g = (r * b) % 65535
				b = (r * g) % 65535
			}

			var col color.RGBA
			col.R = uint8(r)
			col.G = uint8(g)
			col.B = uint8(b)
			col.A = uint8(a)

			img2.SetRGBA(xPos, yPos, col)
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
