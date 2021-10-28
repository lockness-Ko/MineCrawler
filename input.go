package main

import (
	"fmt"
	"math"
	"os"
)

var char_player string = color(80, 120, 200) + "▲"
var char_robot string = color(255, 220, 220) + "∩"

func mapVal(val float64, max float64, min float64) float64 {
	return math.Floor(val*(max-min)) + min
}

func color(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func cls() {
	fmt.Print("\033[H\033[2J")
}

func pos(x int, y int) {
	fmt.Printf("\033[%d;%dH", y/2, x)
}

func Rpos(x int, y int) string {
	return fmt.Sprintf("\033[%d;%dH", y/2, x)
}

func input(prompt string) string {
	fmt.Print(prompt)
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	return scanner.Text()
	// }
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b[:])
}
