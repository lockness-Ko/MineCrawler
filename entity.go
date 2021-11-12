package main

import (
	"fmt"
	"math"
)

type Vec struct {
	x, y int
}

func getDist(e1 Entity, e2 Entity) float64 {
	return math.Sqrt(math.Pow(float64(e2.x-e1.x), 2) + math.Pow(float64(e2.y-e1.y), 2))
}

// func getDir(e1 Entity, e2 Entity) float64 {
// 	return math.Atan2(float64(e2.x-e1.x), float64(e2.y-e1.y))
// }

type Entity struct {
	x, y int
	char string
}

func (_this *Entity) tick() {
}

func (_this *Entity) draw(p Player) {
	if ((*_this).x-(p.x) > width || (*_this).x-(p.x) < 0) || ((*_this).y-(p.y*2) > height*2 || (*_this).y-(p.y*2) < 0) {

	} else {
		pos(_this.x-(p.x), _this.y-(p.y*2))
		fmt.Println(_this.char)
	}
}

type Player struct {
	Entity
	gold, health, fups, upg int
}

func (_this *Entity) collides(m Map, sx, sy float64, tems []Item) []bool {
	moves := []bool{true, true, true, true}
	switch m.getTileAtPos(*_this, 0, -1, sx, sy, tems) {
	case wall:
		moves[0] = false
	}

	switch m.getTileAtPos(*_this, 0, 1, sx, sy, tems) {
	case wall:
		moves[1] = false
	}

	switch m.getTileAtPos(*_this, -1, 0, sx, sy, tems) {
	case wall:
		moves[2] = false
	}

	switch m.getTileAtPos(*_this, 1, 0, sx, sy, tems) {
	case wall:
		moves[3] = false
	}
	return moves
}

func (_this *Player) tick(m Map, tems []Item) {
	if m.getTileAtPos((*_this).Entity, 0, 0, float64(height/2)-1, float64(width/2)-1, tems) == gold {
		(*_this).gold += 1
	}

	if m.getTileAtPos((*_this).Entity, 0, 0, float64(height/2)-1, float64(width/2)-1, tems) == item {
		(*_this).upg += 1
	}

	(*_this).health += (*_this).fups

	moves := (*_this).Entity.collides(m, float64(height/2)-1, float64(width/2)-1, tems)

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
	case "W":
		if moves[0] {
			(*_this).y -= 3
		}
	case "S":
		if moves[1] {
			(*_this).y += 3
		}
	case "A":
		if moves[2] {
			(*_this).x -= 3
		}
	case "D":
		if moves[3] {
			(*_this).x += 3
		}
	default:
	}
}

func (_this *Player) draw() {
	pos(int(width/2), int(height))
	fmt.Println(_this.char)
}
