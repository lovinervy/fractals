package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"time"
)

type Fractal func(z complex128, x1 complex128, x2 complex128) complex128

func main() {
	var g2 complex128 = 0
	var g3 complex128 = 1

	// Качество изображения
	var size int = 1000

	// N оттенков серого
	var iterations uint8 = 32

	img := createCanvas(size)

	Calculate(poincare, g2, g3, img, size, iterations)
	
	file, err := createFile("out", true)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, img)
}


func poincare(z complex128, g2 complex128, g3 complex128) (complex128) {
	res := (cmplx.Pow(z, 4) + g2*cmplx.Pow(z,2)/2 + 2*g3*z + cmplx.Pow(g2, 2)/16) / (4*cmplx.Pow(z, 3) - g2*z - g3)
	return res
}

func createCanvas(size int) (*image.Gray) {
	img := image.NewGray(image.Rectangle{Max: image.Point{X: size, Y: size}})
	return img
}

func setGrayPoint(img *image.Gray, x int, y int, grayTone uint8) {
	img.SetGray(x, y, color.Gray{Y: uint8(grayTone)})
}

func fixDisplay(orientation_point int, size int) (float64) {
	point := 4 * float64(orientation_point) / (float64(size) - 1) - 2
	return point
}

func createFile(name string, timestamp bool)(*os.File, error){
	if timestamp{
		name = fmt.Sprintf("%s-%d.png", name, time.Now().Unix())
	} else {
		name = fmt.Sprintf("%s.png", name)
	}
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return file, err
}

func Calculate(f Fractal, x1 complex128, x2 complex128, img *image.Gray, size int, gray uint8){
	var z complex128
	var zx, zy float64
	var setTone bool = false

	var gray_offset int = 256 / int(gray)

	for y:=0; y < size; y++ {
		zy = fixDisplay(y, size)
		for x:=0; x < size; x++ {
			zx = fixDisplay(x, size)
			z = complex(zx, zy)
			for i:=0; i < int(gray); i++ {
				if cmplx.Abs(z) > 2 {
					setGrayPoint(img, x, y, uint8(gray_offset * i))
					setTone = true
					break
				}
				z = f(z, x1, x2)
			}
			if setTone == false {
				setGrayPoint(img, x, y, 255)
			}
			setTone = false
		}
	}
}