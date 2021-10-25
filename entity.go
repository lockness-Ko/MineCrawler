package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func getDist(e1 Entity, e2 Entity) float64 {
	return math.Sqrt(math.Pow(2, float64(e2.x-e1.x)) + math.Pow(2, float64(e2.y-e1.y)))
}

func getDir(e1 Entity, e2 Entity) float64 {
	return math.Atan2(float64(e2.x-e1.x), float64(e2.y-e1.y))
}

type Entity struct {
	x, y int
	char string
}

func (_this *Entity) tick() {
}

func (_this *Entity) draw() {
	pos(_this.x, _this.y)
	fmt.Println(_this.char)
}

type Robot struct {
	Entity
}

func (_this *Robot) tick(p Player) {
	if getDist((*_this).Entity, p.Entity) < 10 {
		(*_this).x += rand.Intn(3) - 1
		(*_this).y += rand.Intn(3) - 1
	} else {
		(*_this).x += int(2 * math.Sin(getDir((*_this).Entity, p.Entity)))
		(*_this).y -= int(2 * math.Cos(getDir((*_this).Entity, p.Entity)))
	}
}

type Player struct {
	Entity
}

func (_this *Player) tick() {
	funcs := strings.SplitAfter(inp, ";")

	for _, func_ := range funcs {
		func_ = strings.ReplaceAll(func_, ";", "")
		switch func_ {
		case "w":
			(*_this).y -= 1
		case "s":
			(*_this).y += 1
		case "a":
			(*_this).x -= 1
		case "d":
			(*_this).x += 1
		default:

		}
	}
}

func (_this *Player) draw() {
	pos(int(width/2), int(height))
	fmt.Println(_this.char)
}

type Gold struct {
	Entity
}

func (_this *Gold) tick() {
}
