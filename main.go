package main

import (
	"os/exec"
	"time"

	"golang.org/x/term"
)

var width, height int = 0, 0
var init_ int64 = time.Now().Local().UnixNano()

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("tput", "civis").Run()

	width, height, _ = term.GetSize(0)

	height -= 1
	width /= 2
	game := newGame()

	game.start()
}
