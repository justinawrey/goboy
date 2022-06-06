package gb

const (
	lcdHeight = 144
	lcdWidth  = 160
	numPixels = lcdHeight * lcdWidth
)

type Pixel = int
type ppu struct {
	*memory
	scanCycles int
	pixels     []Pixel
}

func newPpu() *ppu {
	pixels := make([]Pixel, numPixels)
	return &ppu{pixels: pixels}
}

// invoked every 456 cycles (every scanline)
// do this 57 times
//
// Get tile -- calculate pointer to correct tile
// Get tile data low -- get actual tile data low
// Get tile data high -- get actual tile data high
// Sleep -- do nothing
// Push -- get em into the fifos
// Render -- RENDER!
func (ppu *ppu) tick() {}
