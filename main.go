package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

// Averages color around a certain point
func getAverageWithin(img image.Image, x, y int, size int) RGBa {

}

// Gets most common colors in image based on averaging
func GetMostCommonWithin() {

}

type RGBa struct {
	R uint32
	G uint32
	B uint32
	a uint32
}

func main() {
	file, err := os.Open("pic.jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			fmt.Printf("RGB: %d %d %d\n", r, g, b)
		}
	}
}
