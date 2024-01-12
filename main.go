package main

import (
	"fmt"
	helpers "raytracer/helpers"
)

func main() {
	var image_width int = 256
	var image_height int = 256
	//Render the image
	fmt.Printf("P3\n %d %d \n 255\n", image_width, image_height)
	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			//Red increases top to bottom and green from left to right
			pixelColor := helpers.CreateColor(float64(i)/float64(image_height-1), float64(j)/float64(image_width-1), 0.0)
			pixelColor.WriteColor()
		}
	}
}
