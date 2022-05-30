package display

import (
	"image/color"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width  = 160
	height = 144
	pxSize = 4
)

var (
	lighter = color.RGBA{155, 188, 15, 1}
	light   = color.RGBA{139, 172, 15, 1}
	dark    = color.RGBA{48, 98, 48, 1}
	darker  = color.RGBA{15, 56, 15, 1}
	colors  = map[int]color.RGBA{
		0: lighter,
		1: light,
		2: dark,
		3: darker,
	}
)

// display.Display implements gb.Renderer
type Display struct {
	*pixelgl.Window
	*imdraw.IMDraw
}

func Run(run func()) {
	pixelgl.Run(run)
}

func New() *Display {
	cfg := pixelgl.WindowConfig{
		Title:  "goboy",
		Bounds: pixel.R(0, 0, width*pxSize, height*pxSize),
	}

	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	imd := imdraw.New(nil)
	return &Display{window, imd}
}

func (d *Display) Render(pixels []int) {
	if d.Window.Closed() {
		d.Window.Destroy()
		os.Exit(0)
	}

	d.IMDraw.Clear()
	d.IMDraw.Reset()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			d.drawPx(pixels, x, y)
		}
	}

	d.Window.Clear(color.White)
	d.IMDraw.Draw(d.Window)
	d.Window.Update()
}

func (d *Display) drawPx(pixels []int, x int, y int) {
	px := pixels[(width*y)+x]
	color := colors[px]

	lowerX := float64((x % width) * pxSize)
	lowerY := float64((height - y - 1) * pxSize)

	d.IMDraw.Color = color
	d.IMDraw.Push(pixel.V(lowerX, lowerY))
	d.IMDraw.Push(pixel.V(lowerX+pxSize, lowerY+pxSize))
	d.IMDraw.Rectangle(0)
}
