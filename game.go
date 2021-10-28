package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inp string = ""
var toDraw string = ""

type Game struct {
	robots []Robot
	player Player
	_map   Map
}

func stats(p Player) {
	// pos(0, 0)
	fmt.Print("\x1b[0m")
	pos(width-15, 2)
	fmt.Print(color(0, 160, 255))
	fmt.Print("----------------")
	pos(width-15, 4)
	fmt.Printf("| Gold:   %d    |", p.gold)
	pos(width-15, 6)
	fmt.Printf("| Health: %s %d %s | ", color(255-int(mapVal(float64(p.health), 3, 0)), 160, int(mapVal(float64(p.health), 2, 0))), p.health, color(0, 160, 255))
	pos(width-15, 8)
	fmt.Printf("| X: %d Y: %d |", p.x, p.y)
	pos(width-15, 10)
	fmt.Print("----------------")
}

func newGame(robots int) Game {
	ents := []Robot{}
	for i := 0; i < robots; i++ { //int(mapVal(rand.Float64(), float64(width)/2-60, float64(width)/2) + 30), int(mapVal(rand.Float64(), float64(height)-30, float64(height)+30))
		ents = append(ents, Robot{Entity{int(mapVal(rand.Float64(), float64(width)/2-60, float64(width)/2) + 30), int(mapVal(rand.Float64(), float64(height)-30, float64(height)+30)), char_robot}})
	}
	return Game{ents, Player{Entity{0, 0, char_player}, 0, 100}, Map{float64(width), float64(height), *NewPerlin(2, 2, 3, time.Now().Local().UnixNano())}}
}

func (_this *Game) start() {
	for {
		(*_this).loop()
	}
}

func (_this *Game) loop() {
	// cls()
	rand.Seed(time.Now().Local().UnixNano())
	(*_this).tick()
	(*_this).draw()
	stats((*_this).player)
	pos(0, height*2+1)
	// fmt.Println(inp)
	inp = input("")
}

func (_this *Game) tick() {
	switch inp {
	case "\t":
		toDraw = Rpos(width/2, height/2) + color(255, 255, 255) + "We do a little testing"
		fmt.Print(toDraw)
		toDraw = ""
		for inp != "\n" {
			inp = input("")
		}

	default:
		((*_this).player).tick((*_this)._map, (*_this).robots)
	}
	for i := 0; i < len((*_this).robots); i++ {
		((*_this).robots[i]).tick((*_this)._map, ((*_this).player))
	}
}

func (_this *Game) draw() {
	((*_this)._map).draw((*_this).player)

	((*_this).player).draw()
	for i := 0; i < len((*_this).robots); i++ {
		((*_this).robots[i]).draw((*_this).player)
	}
}
