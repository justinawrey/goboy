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

	lcdc memReg
	lcds memReg
	scy  memReg
	scx  memReg
	ly   memReg
	lyc  memReg
	wy   memReg
	wx   memReg
}

type memReg struct {
	*ppu
	index uint16
}

func (mr *memReg) get() byte {
	return mr.ppu.memory.readByte(mr.index)
}

func (mr *memReg) set(b byte) {
	mr.ppu.memory.writeByte(mr.index, b)
}

func (ppu *ppu) lcdEnable() bool {
	return getBit(ppu.lcdc.get(), 7)
}

func (ppu *ppu) windowTileMapSelect() bool {
	return getBit(ppu.lcdc.get(), 6)
}

func (ppu *ppu) windowEnable() bool {
	return getBit(ppu.lcdc.get(), 5)
}

func (ppu *ppu) bgAndWindowTileDataSelect() bool {
	return getBit(ppu.lcdc.get(), 4)
}

func (ppu *ppu) bgTileMapSelect() bool {
	return getBit(ppu.lcdc.get(), 3)
}

func (ppu *ppu) objSize() bool {
	return getBit(ppu.lcdc.get(), 2)
}

func (ppu *ppu) objEnable() bool {
	return getBit(ppu.lcdc.get(), 1)
}

func (ppu *ppu) bgAndWindowEnable() bool {
	return getBit(ppu.lcdc.get(), 0)
}

// 0x9800 - 0x9bff
func (ppu *ppu) tileMap0() []byte {
	return ppu.memory.Bytes()[0x9800:0x9c00]
}

// 0x9c00 - 0x9fff
func (ppu *ppu) tileMap1() []byte {
	return ppu.memory.Bytes()[0x9c00:0xa000]
}

// 0x8000 - 0x87ff
func (ppu *ppu) tileData0() []byte {
	return ppu.memory.Bytes()[0x8000:0x8800]
}

// 0x8800 - 0x8fff
func (ppu *ppu) tileData1() []byte {
	return ppu.memory.Bytes()[0x8800:0x9000]
}

// 0x9000 - 0x97ff
func (ppu *ppu) tileData2() []byte {
	return ppu.memory.Bytes()[0x9000:0x9800]
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

	ppu := ppu{pixels: pixels}
	ppu.lcdc = memReg{&ppu, 0xff40}
	ppu.lcds = memReg{&ppu, 0xff41}
	ppu.scy = memReg{&ppu, 0xff42}
	ppu.scx = memReg{&ppu, 0xff43}
	ppu.ly = memReg{&ppu, 0xff44}
	ppu.lyc = memReg{&ppu, 0xff45}
	ppu.wy = memReg{&ppu, 0xff4a}
	ppu.wx = memReg{&ppu, 0xff4b}

	return &ppu
}

func (ppu *ppu) getMode(scanline byte, cycles int) (mode byte) {
	// v-blank
	if scanline >= 144 {
		return 1
	}

	// oam scan
	if cycles <= 80 {
		return 2
	}

	// drawing pixels
	// TODO: does this need to be more precise?
	if cycles <= 172 {
		return 3
	}

	// h-blank
	return 0
}

// cycles is 0 - 456 (cycles elapsed in given scanline)
func (ppu *ppu) updateLcdStatus(cycles int) {
	scanline := ppu.ly.get()
	status := ppu.lcds.get()
	mode := ppu.getMode(scanline, cycles)

	// bits 0, 1: set mode
	status = (status & 0b11111100) | mode

	// bit 2: lyc == ly flag
	if scanline == ppu.lyc.get() {
		status = setBit(status, 2, true)
	} else {
		status = setBit(status, 2, false)
	}

	// TODO: set STAT interrupt source
	ppu.lcds.set(status)
}

func (ppu *ppu) incrementScanline() {
	ly := ppu.ly.get()
	maxScanlines := byte(153)

	if ly >= maxScanlines {
		ppu.ly.set(0)
		return
	}

	ppu.ly.set(ly + 1)
}

func (ppu *ppu) drawScanline(scanline byte) {
	if ppu.bgAndWindowEnable() {
		ppu.drawBG(scanline)

		if ppu.windowEnable() {
			ppu.drawWindow(scanline)
		}
	} else {
		ppu.drawWhiteBgAndWindow(scanline)
	}

	if ppu.objEnable() {
		ppu.drawObjs(scanline)
	}
}

func (ppu *ppu) drawBG(scanline byte)     {}
func (ppu *ppu) drawWindow(scanline byte) {}
func (ppu *ppu) drawObjs(scanline byte)   {}

// TODO: this is a terrible name
func (ppu *ppu) drawWhiteBgAndWindow(scanline byte) {}
