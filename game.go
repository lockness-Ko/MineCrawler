package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inp string = ""

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
	fmt.Printf("| Health: %d  |", p.health)
	pos(width-15, 8)
	fmt.Printf("| Kills:  %d    |", p.kills)
	pos(width-15, 10)
	fmt.Print("----------------")
}

func newGame(robots int) Game {
	ents := []Robot{}
	for i := 0; i < robots; i++ {
		ents = append(ents, Robot{Entity{int(mapVal(rand.Float64(), float64(width)/2-60, float64(width)/2) + 30), int(mapVal(rand.Float64(), float64(height)-30, float64(height)+30)), char_robot}})
	}
	return Game{ents, Player{Entity{0, 0, char_player}, 0, 100, 0}, Map{float64(width), float64(height), *NewPerlin(2, 2, 3, time.Now().Local().UnixNano())}}
}

func (_this *Game) start() {
	for {
		(*_this).loop()
	}
}

func (_this *Game) loop() {
	cls()
	rand.Seed(time.Now().Local().UnixNano())
	(*_this).tick()
	(*_this).draw()
	stats((*_this).player)
	pos(0, height*2+1)
	inp = input(color(220, 180, 50) + ">")
	// time.Sleep(time.Millisecond * 300)
}

func (_this *Game) tick() {
	((*_this).player).tick((*_this)._map)
	for i := 0; i < len((*_this).robots); i++ {
		((*_this).robots[i]).tick((*_this).player)
	}
}

func (_this *Game) draw() {
	((*_this)._map).draw((*_this).player)

	((*_this).player).draw()
	for i := 0; i < len((*_this).robots); i++ {
		((*_this).robots[i]).draw((*_this).player)
	}

}
