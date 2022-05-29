package gb

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/justinawrey/goboy/display"
)

const refreshHz = 60

type Gb struct {
	*memory
	*ppu
	cpu
}

func NewGb() *Gb {
	gb := new(Gb)
	mem := newMemory()
	ppu := new(ppu)

	gb.memory = mem
	gb.ppu = ppu
	gb.cpu.memory = mem
	gb.cpu.ppu = ppu
	ppu.memory = mem

	return gb
}

func (gb *Gb) LoadCartridge (path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = io.Copy(gb.memory, f)
	if err != nil {
		log.Fatal(err)
	}
}

func (gb *Gb) boot () error {
	gb.cpu.pc = 0x0100
	return nil
}

func (gb *Gb) mainLoop () {
	c := time.Tick(time.Second / refreshHz)

	for range c {
		gb.cpu.tick()
		display.Draw()
	}
}

func (gb *Gb) Run () {
	if err := gb.boot(); err != nil {
		log.Fatal(err)
	}

	gb.mainLoop()
}

