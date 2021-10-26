package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var char_player string = "▲"
var char_robot string = "§"

func mapVal(val float64, max float64, min float64) float64 {
	return math.Floor(val*(max-min)) + min
}

func cls() {
	fmt.Print("\033[H\033[2J")
}

func pos(x int, y int) {
	fmt.Printf("\033[%d;%dH", y/2, x)
}

func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// fmt.Println(scanner.Text())
		return scanner.Text()
	}
	return ""
}
