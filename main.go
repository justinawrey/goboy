package main

import (
	"github.com/justinawrey/goboy/display"
	"github.com/justinawrey/goboy/gb"
)

func main() {
	display.Run(run)
}

func run() {
	gb := gb.NewGb()
	display := display.New()
	defer display.Destroy()

	gb.ConnectDisplay(display)
	gb.LoadCartridge("./rom/tetris.gb")

	gb.Run()
}
