package main

import (
	"github.com/justinawrey/goboy/gb"
)

func main() {
	gb := gb.NewGb()
	gb.LoadCartridge("./rom/tetris.gb")
	gb.Run()
}
