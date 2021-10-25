package main

import "golang.org/x/term"

var width, height int = 0, 0

func main() {
	width, height, _ = term.GetSize(0)

	height -= 1
	// width /= 2
	game := newGame(1)

	game.start()
}
