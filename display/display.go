package display

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Display struct {
	*glfw.Window
}

func init() {
	runtime.LockOSThread()
}

func New() *Display {
	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}

	window, err := glfw.CreateWindow(640, 480, "goboy", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()
	return &Display{window}
}

func (d *Display) Close() {
	glfw.Terminate()
}

func (d *Display) Render() {
	if d.ShouldClose() {
		d.Close()
	}

	// do stuff...
	d.SwapBuffers()
	glfw.PollEvents()
}
