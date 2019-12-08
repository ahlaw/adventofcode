package main

import (
	"fmt"
	"io/ioutil"
)

const w, h = 25, 6

func render(image string) [h][w]byte {
	var decoded [h][w]byte
	numLayers := len(image) / (w * h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for l := 0; l < numLayers; l++ {
				if c := image[l*w*h+w*i+j]; c != '2' {
					decoded[i][j] = c
					break
				}
			}
		}
	}
	return decoded
}

func draw(image [h][w]byte) {
	for _, row := range image {
		for _, c := range row {
			switch c {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	image := string(input)
	renderedImage := render(image)
	draw(renderedImage)
}
