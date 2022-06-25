package gb

import (
	"io"
	"log"
	"math"
	"os"
	"time"
)

const refreshHz = float64(cpuHz) / float64(cyclesPerFrame)

type Renderer interface {
	Render([]Pixel)
}

type Gb struct {
	renderer Renderer
	*memory
	*ppu
	*cpu
}

func NewGb() *Gb {
	gb := new(Gb)
	mem := newMemory()
	ppu := newPpu()
	cpu := new(cpu)

	gb.memory = mem
	gb.ppu = ppu
	gb.cpu = cpu

	cpu.memory = mem
	cpu.ppu = ppu
	ppu.memory = mem

	return gb
}

func (gb *Gb) LoadCartridge(path string) {
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

func (gb *Gb) ConnectDisplay(r Renderer) {
	gb.renderer = r
}

func (gb *Gb) boot() error {
	gb.cpu.pc = 0x0100
	return nil
}

func (gb *Gb) mainLoop() {
	// render at 59.7275 fps
	timePerFrame := time.Duration(math.Round(float64(time.Second) / refreshHz))
	c := time.Tick(timePerFrame)

	for range c {
		gb.cpu.tick()
		gb.renderer.Render(gb.ppu.visiblePixels())
	}
}

func (gb *Gb) Run() {
	if err := gb.boot(); err != nil {
		log.Fatal(err)
	}

	gb.mainLoop()
}
