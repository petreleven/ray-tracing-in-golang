package main

import (
	"fmt"
	helpers "raytracer/helpers"
)

const aspectRatio float64 = 16 / 9
const image_width int = 256

func main() {
	//calcilate image height based on aspect ratio and width , its atleast 1
	var image_height int = int(float64(image_width) / aspectRatio)
	if (image_height) < 1 {
		image_height = 1
	}

	//Set up viewport and Camera
	var focal_length int = 1
	var viewport_height int = 2
	var viewport_width int = viewport_height * int(image_width/image_height)
	var camera_center helpers.Vec3 = *helpers.CreateVec3(0, 0, 0)

	//Viewport vectors
	var viewport_u helpers.Vec3 = *helpers.CreateVec3(float64(viewport_width), 0, 0)
	var viewport_v helpers.Vec3 = *helpers.CreateVec3(0, float64(-viewport_height), 0)
	//pixel deltas
	pixel_delta_u := viewport_u.DividebyFloat(float64(image_width))
	pixel_delta_v := viewport_v.DividebyFloat(float64(image_height))
	//upper left pixel
	viewport_upper_left_pixel := camera_center.Substract(helpers.CreateVec3(0, 0, float64(focal_length)))
	viewport_upper_left_pixel = viewport_upper_left_pixel.Substract(viewport_u.DividebyFloat(2.0))
	viewport_upper_left_pixel = viewport_upper_left_pixel.Substract(viewport_v.DividebyFloat(2.0))
	//Add padding of 1/2 pixel in u and y
	pixel00_loc := viewport_upper_left_pixel.Add(pixel_delta_u.Add(pixel_delta_v).MultiplybyFloat(0.5))

	//Render the image
	fmt.Printf("P3\n %d %d \n 255\n", image_width, image_height)
	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			//get direction from pixel to camera
			pixel_center := pixel00_loc.Add(pixel_delta_u.MultiplybyFloat(float64(i)).Add(pixel_delta_v.MultiplybyFloat(float64(j))))
			ray_direction := pixel_center.Substract(&camera_center)
			ray := helpers.Ray{}
			ray.SetOrigin(camera_center)
			ray.SetDir(*ray_direction)
			colorOfRay := rayColor(&ray)
			pixelColor := helpers.CreateColor(colorOfRay.GetCordinateAtIndex(0), colorOfRay.GetCordinateAtIndex(1), colorOfRay.GetCordinateAtIndex(2))
			pixelColor.WriteColor()
		}
	}
}

func rayColor(ray *helpers.Ray) *helpers.Vec3 {
	unit_direction := ray.GetDir()
	unit_direction = unit_direction.UnitVector()
	//lerp in y direction
	a := 0.5 * (unit_direction.Y() + 1.0)
	startvalue := helpers.Color{Vec3: *helpers.CreateVec3(1, 1, 1)}
	endvalue := helpers.Color{Vec3: *helpers.CreateVec3(0.5, 0.7, 1)}
	blend := startvalue.MultiplybyFloat((1 - a)).Add(endvalue.MultiplybyFloat(a))
	return blend
}
