package vec3

import (
	"fmt"
	"math"
	"strconv"
)

type Vec3 struct {
	e [3]float64
}

func CreateVec3(e0, e1, e2 float64) *Vec3 {
	var e [3]float64
	e[0] = e0
	e[1] = e1
	e[2] = e2
	return &Vec3{e: e}
}

func (v *Vec3) X() float64 {
	return v.e[0]
}
func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) CreateNewNegatedVec3() *Vec3 {
	return CreateVec3(-v.e[0], -v.e[1], -v.e[2])
}

func (v *Vec3) GetCordinateAtIndex(i int) float64 {
	return v.e[i]
}

func (e *Vec3) AddtoOriginalVec3(v *Vec3) {
	e.e[0] += v.e[0]
	e.e[1] += v.e[1]
	e.e[2] += v.e[2]
}

func (e *Vec3) MultipyToOriginalVec3(v *Vec3) {
	e.e[0] *= v.e[0]
	e.e[1] *= v.e[1]
	e.e[2] *= v.e[2]
}

func (e *Vec3) DivideOriginalVec3ByFloat(t float64) {
	e.e[0] /= t
	e.e[1] /= t
	e.e[2] /= t
}

func (e *Vec3) Length() float64 {
	return math.Sqrt(lengthSquared(*e))
}

func lengthSquared(v Vec3) float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

// Utility functions
func (v *Vec3) String() string {
	stringCordinates := strconv.FormatFloat(v.e[0], 'f', -1, 64) + " " + strconv.FormatFloat(v.e[1], 'f', -1, 64) + strconv.FormatFloat(v.e[2], 'f', -1, 64)
	return stringCordinates
}

func (v *Vec3) Add(u *Vec3) *Vec3 {
	return CreateVec3(v.e[0]+u.e[0], v.e[1]+u.e[1], v.e[2]+u.e[2])
}

func (v *Vec3) Substract(u *Vec3) *Vec3 {
	return CreateVec3(v.e[0]-u.e[0], v.e[1]-u.e[1], v.e[2]-u.e[2])
}
func (v *Vec3) Multiply(u *Vec3) *Vec3 {
	return CreateVec3(v.e[0]*u.e[0], v.e[1]*u.e[1], v.e[2]*u.e[2])
}

func (v *Vec3) MultiplybyFloat(t float64) *Vec3 {
	return CreateVec3(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func (v *Vec3) DividebyFloat(t float64) *Vec3 {
	return CreateVec3(v.e[0]/t, v.e[1]/t, v.e[2]/t)
}

func (v *Vec3) Dot(u *Vec3) float64 {
	return v.e[0]*u.e[0] + v.e[1]*u.e[1] + v.e[2]*u.e[2]
}

func (u *Vec3) Cross(v *Vec3) Vec3 {
	x := u.e[1]*v.e[2] - u.e[2]*v.e[1]
	y := u.e[2]*v.e[0] - u.e[0]*v.e[2]
	z := u.e[0]*v.e[1] - u.e[1]*v.e[0]
	return *CreateVec3(x, y, z)
}

func (v *Vec3) UnitVector() Vec3 {
	return *CreateVec3(v.e[0]/v.Length(), v.e[1]/v.Length(), v.e[2]/v.Length())
}

type Color struct {
	Vec3
}

func (pixelColor Color) WriteColor() {
	fmt.Printf("%d %d %d \n", int(pixelColor.X()*255.99), int(pixelColor.Y()*255.99), int(pixelColor.Z()*255.99))
}
func CreateColor(r, g, b float64) *Color {
	return &Color{Vec3: *CreateVec3(r, g, b)}
}
