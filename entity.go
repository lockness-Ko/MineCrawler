package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
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

func (_this *Player) collides(m Map) []bool {
	moves := []bool{true, true, true, true}
	switch m.getTileAtPos(*_this, 0, -1) {
	case "#":
		moves[0] = false
	}

	switch m.getTileAtPos(*_this, 0, 1) {
	case "#":
		moves[1] = false
	}

	switch m.getTileAtPos(*_this, -1, 0) {
	case "#":
		moves[2] = false
	}

	switch m.getTileAtPos(*_this, 1, 0) {
	case "#":
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
	funcs := strings.SplitAfter(inp, ";")

	for _, func_ := range funcs {
		moves := (*_this).collides(m)

		func_ = strings.ReplaceAll(func_, ";", "")

		switch func_ {
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
			if strings.Contains(func_, "f") {
				num := getNumber(strings.Split(func_, ",")[1])

				for i := 0; i < num; i++ {
					switch strings.Split(func_, ",")[2] {
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
			}
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
