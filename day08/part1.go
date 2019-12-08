package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	image := string(input)
	const w, h = 25, 6
	numLayers := len(image) / (w * h)
	var output, fewestZeros int
	for l := 0; l < numLayers; l++ {
		var numZeros, numOnes, numTwos int
		for _, c := range image[l*w*h : (l+1)*w*h] {
			switch c {
			case '0':
				numZeros++
			case '1':
				numOnes++
			case '2':
				numTwos++
			}
		}
		if numZeros < fewestZeros || l == 0 {
			fewestZeros = numZeros
			output = numOnes * numTwos
		}
	}
	fmt.Println(output)
}
