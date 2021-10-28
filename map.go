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
var floor string = color(30, 30, 30) + "@"
var wall string = color(170, 170, 170) + "#"
var wall1 string = color(130, 130, 130) + "$"
var wall2 string = color(110, 110, 110) + "@"
var wall3 string = color(100, 100, 100) + "@"
var gold string = color(220, 200, 0) + "¢"
var chars = []string{floor, gold, floor, floor, floor, wall, wall1, wall2, wall3}

func (_this *Map) draw(p Player) {
	pos(0, 0)
	scale := 2.
	for i := 0.; i < (*_this).x; i++ {
		for j := 0.; j < (*_this).y; j++ {
			aa := int(mapVal((*_this).noise.Noise2D(float64(math.Round((i+float64(p.y))/scale))/10, float64(math.Round((j+float64(p.x))/scale))/10)+1.1, 4, 0))
			fmt.Print(chars[aa])
		}
		fmt.Print("\r\n")
	}
}

func (_this *Map) getTileAtPos(p Entity, x, y int, sx, sy float64) string {
	scale := 2.
	return chars[int(mapVal((*_this).noise.Noise2D(float64(math.Round((float64(y)+float64(p.y)+sx)/scale))/10, float64(math.Round((float64(x)+float64(p.x)+sy)/scale))/10)+1.1, 4, 0))]
}
