package gb

const (
	lcdHeight = 144
	lcdWidth  = 160
	numPixels = lcdHeight * lcdWidth
)

type Pixel = int
type ppu struct {
	*memory
	pixels []Pixel
}

// an 8x8 grouping of pixels
type tile struct {
	pixels []Pixel
}

// constructs a well-formed 'tile' from 16 bytes
func newTile(bytes []byte) tile {
	pixels := make([]Pixel, 64)

	for i := 0; i < 16; i += 2 {
		b1 := bytes[i]
		b2 := bytes[i+1]
		pixels = append(pixels, newPixelRow(b1, b2)...)
	}

	return tile{pixels}
}

func newPixelRow(b1 byte, b2 byte) []Pixel {
	pixels := make([]Pixel, 8)

	for i := 7; i >= 0; i-- {
		bit1 := getBit(b1, i)
		bit2 := getBit(b2, i)
		pixels = append(pixels, newPixel(bit1, bit2))
	}

	return pixels
}

func newPixel(bit1 bool, bit2 bool) (pixel Pixel) {
	if !bit1 && !bit2 {
		// 00
		pixel = 0
	} else if bit1 && !bit2 {
		// 10 (01 flipped)
		pixel = 1
	} else if !bit1 && bit2 {
		// 01 (10 flipped)
		pixel = 2
	} else {
		// 11
		pixel = 3
	}

	return pixel
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
// needs to populate pixels
func (ppu *ppu) drawScanline() {}
