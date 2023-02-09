package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 150
	height = 15
)

type Universe [][]bool

func (Uni Universe) Show() {
	for y := 0; y < height; y++ {
		print("\n")
		for x := 0; x < width; x++ {
			if Uni[y][x] {
				print("*")
			} else {
				print(" ")
			}
		}
	}
}

func (Uni Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		Uni.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

func (Uni Universe) Set(x, y int, b bool) {
	Uni[y][x] = b
}

func (Uni Universe) Alive(x, y int) bool {

	if x < 0 {
		x += width
	}

	if x > width-1 {
		x = x % width
	}

	if y < 0 {
		y += height
	}

	if y > height-1 {
		y = y % height
	}

	if Uni[y][x] {
		return true
	} else {
		return false
	}
}

func (Uni Universe) Neighbors(x, y int) int {
	var count int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && Uni.Alive(x+j, y+i) {
				count++
			}
		}
	}
	return count
}

func (Uni Universe) NextGen(x, y int) bool {

	n := Uni.Neighbors(x, y)

	if Uni.Alive(x, y) {
		if n == 2 || n == 3 {
			return true
		} else {
			return false
		}
	} else {
		if n == 3 {
			return true
		}
	}
	return false
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.NextGen(x, y))
		}
	}
}
func NewUniverse() Universe {
	Uni := make(Universe, height)
	for i := range Uni {
		Uni[i] = make([]bool, width)
	}

	return Uni
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 150; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
		fmt.Print("\033[H\033[2J")
	}

}
