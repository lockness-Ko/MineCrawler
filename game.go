package main

import (
	"math/rand"
	"time"
)

var inp string = ""

type Game struct {
	robots []Robot
	player Player
	_map   Map
}

func newGame(robots int) Game {
	ents := []Robot{}
	for i := 0; i < robots; i++ {
		ents = append(ents, Robot{Entity{10, 10, char_robot}})
	}
	return Game{ents, Player{Entity{0, 0, char_player}}, Map{float64(width), float64(height), *NewPerlin(2, 2, 3, time.Now().Local().UnixNano())}}
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
	pos(0, height*2+1)
	inp = input(">")
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
		((*_this).robots[i]).draw()
	}
}
