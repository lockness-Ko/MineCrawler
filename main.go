package main

import (
	"os/exec"

	"golang.org/x/term"
)

var width, height int = 0, 0

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	width, height, _ = term.GetSize(0)

	height -= 1
	// width /= 2
	game := newGame(10)

	game.start()
}
