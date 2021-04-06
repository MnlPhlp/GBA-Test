package main

import (
	"GBA-Test/constants"
	"GBA-Test/floors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func getRandomLevel() floors.FloorList {
	floorList := floors.NewFloors()
	for i := 0; i < 10; i++ {
		x1 := rand.Intn(constants.WBg)
		floorList = floorList.AddFloor(
			floors.Floor{
				X1: int16(x1),
				X2: int16(rand.Intn(constants.W) + x1),
				Y:  int16(rand.Intn(constants.H-20) + 10),
			},
		)
	}
	return floorList
}

func loadImage(path string) image.Image {
	src, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	img, err := png.Decode(src)
	if err != nil {
		log.Fatalln(err)
	}
	return img
}

func levelFromPng(path string) floors.FloorList {
	floorList := floors.NewFloors()
	black := color.RGBA{A: 255}
	img := loadImage(path)
	// store currently running lines
	current := make(map[int]int)
	// loop over all image pixels
	for x := 0; x <= img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			// check if pixel is filled or border is reached
			if x < img.Bounds().Dx() && img.At(x, y) == black {
				// add to current
				y += 2
				current[y]++
			} else {
				// check if finished a line
				if width, ok := current[y]; ok {
					floorList = floorList.AddFloor(
						floors.Floor{
							X1: int16(x - width),
							X2: int16(x - 1),
							Y:  int16(y),
						},
					)
					delete(current, y)
				}
			}
		}
	}
	return floorList
}

func main() {
	var (
		random bool
		input  string
	)
	flag.BoolVar(&random, "r", false, "random level")
	flag.StringVar(&input, "i", "", "input file")
	flag.Parse()
	var floorList floors.FloorList
	if random {
		floorList = getRandomLevel()
	} else {
		floorList = levelFromPng(input)
	}
	out := "var Floors FloorList = []Floor{"
	for _, floor := range floorList {
		out += fmt.Sprintf("{%v,%v,%v},", floor.X1, floor.X2, floor.Y)
	}
	out += "}"
	println(out)
}
