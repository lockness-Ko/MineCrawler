package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
)

func getDist(e1 Entity, e2 Entity) float64 {
	return math.Sqrt(math.Pow(float64(e2.x-e1.x), 2) + math.Pow(float64(e2.y-e1.y), 2))
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

func (_this *Entity) draw(p Player) {
	if ((*_this).x-(p.x) > width || (*_this).x-(p.x) < 0) || ((*_this).y-(p.y) > height*2 || (*_this).y-(p.y) < 0) {

	} else {
		pos(_this.x-(p.x), _this.y-(p.y))
		fmt.Println(_this.char)
	}
}

type Robot struct {
	Entity
}

func (_this *Robot) tick() {
	(*_this).x += rand.Intn(3) - 1
	(*_this).y += rand.Intn(3) - 1
}

type Player struct {
	Entity
	gold, health, kills int
}

func (_this *Player) collides(m Map) []bool {
	moves := []bool{true, true, true, true}
	switch m.getTileAtPos(*_this, 0, -1) {
	case wall:
		moves[0] = false
	}

	switch m.getTileAtPos(*_this, 0, 1) {
	case wall:
		moves[1] = false
	}

	switch m.getTileAtPos(*_this, -1, 0) {
	case wall:
		moves[2] = false
	}

	switch m.getTileAtPos(*_this, 1, 0) {
	case wall:
		moves[3] = false
	}
	return moves
}

func getNumber(str string) int {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	o, _ := strconv.Atoi(re.FindAllString(str, -1)[0])
	return o
}

func (_this *Player) tick(m Map) {
	moves := (*_this).collides(m)

	switch inp {
	case "w":
		if moves[0] {
			(*_this).y -= 1
		}
	case "s":
		if moves[1] {
			(*_this).y += 1
		}
	case "a":
		if moves[2] {
			(*_this).x -= 1
		}
	case "d":
		if moves[3] {
			(*_this).x += 1
		}
	default:
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
