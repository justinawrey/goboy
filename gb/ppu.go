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

// lcd control
func (ppu *ppu) lcdc() byte {
	return ppu.memory.readByte(0xff40)
}

func (ppu *ppu) lcdEnable() bool {
	return getBit(ppu.lcdc(), 7)
}

func (ppu *ppu) windowTileMapControl() bool {
	return getBit(ppu.lcdc(), 6)
}

func (ppu *ppu) windowEnable() bool {
	return getBit(ppu.lcdc(), 5)
}

func (ppu *ppu) bgAndWindowTileData() bool {
	return getBit(ppu.lcdc(), 4)
}

func (ppu *ppu) bgTileMapControl() bool {
	return getBit(ppu.lcdc(), 3)
}

func (ppu *ppu) objSize() bool {
	return getBit(ppu.lcdc(), 2)
}

func (ppu *ppu) objEnable() bool {
	return getBit(ppu.lcdc(), 1)
}

func (ppu *ppu) bgAndWindowPriority() bool {
	return getBit(ppu.lcdc(), 0)
}

// lcd status
func (ppu *ppu) lcds() byte {
	return ppu.memory.readByte(0xff41)
}

// scrolling related memory locations
func (ppu *ppu) scy() byte {
	return ppu.memory.readByte(0xff42)
}

func (ppu *ppu) scx() byte {
	return ppu.memory.readByte(0xff43)
}

func (ppu *ppu) ly() byte {
	return ppu.memory.readByte(0xff44)
}

func (ppu *ppu) lyc() byte {
	return ppu.memory.readByte(0xff45)
}

func (ppu *ppu) wy() byte {
	return ppu.memory.readByte(0xff4a)
}

func (ppu *ppu) wx() byte {
	return ppu.memory.readByte(0xff4b)
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

// cycles is 0 - 456 (cycles elapsed in given scanline)
func (ppu *ppu) updateLcdStatus(cycles int) {

	// TODO: set mode

	// TODO: lyc==ly

	// TODO: set STAT interrupt source
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
func (ppu *ppu) drawScanline() {

}
