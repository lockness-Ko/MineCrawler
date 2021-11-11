package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var inp string = ""
var toDraw string = ""

type Game struct {
	player Player
	_map   Map
	item   []Item
	prinps []string
}

func newGame() Game {
	return Game{Player{Entity{0, 0, char_player}, 0, 100, -1}, Map{float64(width), float64(height), *NewPerlin(2, 2, 3, time.Now().Local().UnixNano())}, []Item{Item{Vec{0, 0}, "", ""}}, []string{}}
}

func (_this *Game) start() {
	for {
		(*_this).loop()
	}
}

func (_this *Game) menu(p Player) {
	off := 1
	// pos(0, 0)
	fmt.Print("\x1b[0m")
	pos(width+off, 2)
	fmt.Print(color(0, 160, 255))
	fmt.Print("----------------")
	pos(width+off, 4)
	fmt.Printf("| 🪙   %d    |", p.gold)
	pos(width+off, 6)
	fmt.Printf("| ⛽: %s %d %s | ", color(255-int(mapVal(float64(p.health), 3, 0)), 160, int(mapVal(float64(p.health), 2, 0))), p.health, color(0, 160, 255))
	pos(width+off, 8)
	fmt.Printf("| X: %d Y: %d |", p.x, p.y)
	pos(width+off, 10)
	fmt.Print("----------------")

	off = 5
	pos(width+off, 15)
	fmt.Print(color(0, 240, 0) + "Buy upgrades")
	pos(width+off, 17)
	fmt.Print(color(0, 240, 0) + "Craft ships")

}

func (_this *Game) loop() {
	// cls()
	// rand.Seed(time.Now().Local().UnixNano())
	(*_this).tick()
	(*_this).draw()
	(*_this).menu((*_this).player)
	// pos(0, height*2+1)
	pos(width+4, height+23)
	inp = input("")
}

func (_this *Game) printMenu() {
	pos(width+3, height)
	fmt.Print("--------------------------------")
	leng := len((*_this).prinps)
	for i := 0; i < 10; i++ {
		j := leng - i
		if j >= leng || j <= 0 {
			continue
		}
		pos(width+3, height+(i*2))
		if (*_this).prinps[j] == "w" || (*_this).prinps[j] == "a" || (*_this).prinps[j] == "s" || (*_this).prinps[j] == "d" || (*_this).prinps[j] == "W" || (*_this).prinps[j] == "A" || (*_this).prinps[j] == "S" || (*_this).prinps[j] == "D" {
			fmt.Print(color(100, 240, 100) + "Movement Command: " + (*_this).prinps[j])
		} else {
			fmt.Print(color(240, 100, 100) + (*_this).prinps[j])
		}
	}
	fmt.Print(color(100, 240, 100))
	pos(width+3, height+21)
	fmt.Print("--------------------------------")
	pos(width+3, height+23)
	fmt.Print(color(130, 130, 250))
	fmt.Print(">")
}

func (_this *Game) tick() {

	if rand.Intn(10) >= 2 {
		rand.Seed(time.Now().Local().UnixNano())
		mx := 1000
		(*_this).item = append((*_this).item, Item{Vec{rand.Intn(mx) - (mx / 2), rand.Intn(mx) - (mx / 2)}, "test", "test"})
	}

	if (*_this).player.health%20 == 0 {
		(*_this).prinps = append((*_this).prinps, "WARNING: Fuel at "+strconv.Itoa((*_this).player.health))
	}
	(*_this).prinps = append((*_this).prinps, inp)
	switch inp {
	case "\t":
		pos(width+3, 15)
		curry := 15
		toDraw = ""
		for inp != "`" {
			pos(width+3, curry)
			fmt.Print(color(0, 200, 150) + ">")
			inp = input("")
			switch inp {
			case "w":
				pos(width+3, curry)
				fmt.Print(color(0, 200, 150) + "  ")
				curry -= 2
			case "s":
				pos(width+3, curry)
				fmt.Print(color(0, 200, 150) + "  ")
				curry += 2
			}
		}

	default:
		((*_this).player).tick((*_this)._map, (*_this).item)
		_this.printMenu()
	}
}

func (_this *Game) draw() {
	((*_this)._map).draw((*_this).player, (*_this).item)

	((*_this).player).draw()
}
