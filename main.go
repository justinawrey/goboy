package main

import (
	"log"
	"os"

	"github.com/justinawrey/goboy/audit"
	"github.com/justinawrey/goboy/display"
	"github.com/justinawrey/goboy/gb"
)

// goboy run -- runs goboy
// goboy audit -- generates cpu opcode completion chart
func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fail()
	}

	switch cmd := args[0]; cmd {
	case "run":
		display.Run(run)
	case "audit":
		audit.Generate()
	default:
		fail()
	}
}

func run() {
	gb := gb.NewGb()
	display := display.New()
	defer display.Destroy()

	gb.ConnectDisplay(display)
	gb.LoadCartridge("./rom/tetris.gb")

	gb.Run()
}

func fail() {
	log.Fatal("Usage: goboy <run|audit>")
}
