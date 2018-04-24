package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func main() {

	earth := createEarth()
	newEarth := createEarth()
	earthLength := len(earth)
	printEarth(earth)
	for true {
		time.Sleep(500000000)
		for y := 0; y < earthLength; y++ {
			for x := 0; x < earthLength; x++ {
				s := computeNeighbourState(earth, x, y)
				var cellLife int8
				if s == 3 || (earth[y][x] == 1 && s == 2) {
					cellLife = 1
				}
				newEarth[y][x] = cellLife
			}
		}

		printEarth(newEarth)
		earth, newEarth = newEarth, earth
	}
}

func createEarth() [][]int8 {
	return [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
}

func printEarth(earth [][]int8) {
	//printEarthText(earth)
	printEarthImage(earth)
}

func printEarthText(earth [][]int8) {
	earthLength := len(earth)
	fmt.Println("------------")
	for y := 0; y < earthLength; y++ {
		fmt.Printf("%+v\n", earth[y])
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func printEarthImage(earth [][]int8) {
	img := image.NewRGBA(image.Rect(0, 0, 5, 5))
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	earthLength := len(earth)
	for y := 0; y < earthLength; y++ {
		for x := 0; x < earthLength; x++ {
			if earth[y][x] == 1 {
				img.Set(x, y, black)
			} else {
				img.Set(x, y, white)
			}
		}
	}
	f, _ := os.OpenFile(fmt.Sprintf("/tmp/out/out-%d.png", makeTimestamp()), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func computeNeighbourState(earth [][]int8, x int, y int) int8 {
	var s int8
	earthLength := len(earth)
	if x > 0 {
		if y > 0 {
			s += earth[y-1][x-1]
		}
		s += earth[y][x-1]
		if y < (earthLength - 1) {
			s += earth[y+1][x-1]
		}
	}
	if y > 0 {
		s += earth[y-1][x]
	}
	if y < (earthLength - 1) {
		s += earth[y+1][x]
	}
	if x < (earthLength - 1) {
		if y > 0 {
			s += earth[y-1][x+1]
		}
		s += earth[y][x+1]
		if y < (earthLength - 1) {
			s += earth[y+1][x+1]
		}
	}
	return s
}
