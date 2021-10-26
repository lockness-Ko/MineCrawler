package main

import (
	"fmt"
	"math"
)

type Map struct {
	y, x  float64
	noise Perlin
}

// var chars = []string{".", ",", "-", "~", ":", ";", "=", "!", "*", "#", "$", "@"}
var chars = []string{".", ".", ".", ".", ".", "#", " ", " ", " "}

func (_this *Map) draw(p Player) {
	pos(0, 0)
	scale := 2.
	for i := 0.; i < (*_this).x; i++ {
		// rand.Seed(time.Now().Local().UnixNano())
		for j := 0.; j < (*_this).y; j++ {
			fmt.Print(chars[int(mapVal((*_this).noise.Noise2D(float64(math.Round((i+float64(p.y))/scale))/10, float64(math.Round((j+float64(p.x))/scale))/10)+1.1, 4, 0))])
		}
		fmt.Println()
	}
}

func (_this *Map) getTileAtPos(p Player, x, y int) string {
	sy, sx := float64(width/2)-1, float64(height/2)-1
	scale := 2.
	return chars[int(mapVal((*_this).noise.Noise2D(float64(math.Round((float64(y)+float64(p.y)+sx)/scale))/10, float64(math.Round((float64(x)+float64(p.x)+sy)/scale))/10)+1.1, 4, 0))]
}
