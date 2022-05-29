package main

import (
	"github.com/justinawrey/goboy/display"
	"github.com/justinawrey/goboy/gb"
)

func main() {
	gb := gb.NewGb()
	display := display.New()
	defer display.Close()

	gb.ConnectDisplay(display)
	gb.LoadCartridge("./rom/tetris.gb")

	gb.Run()
}
