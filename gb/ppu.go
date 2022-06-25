package gb

const (
	lcdHeight = 144
	lcdWidth  = 160
	numPixels = lcdHeight * lcdWidth
)

type Pixel = int
type ppu struct {
	*memory

	// the full 256 x 256 screen -- not entirely visible
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
	ppu := ppu{}

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

// TODO: this is terribly optimized.  does it matter?
func (ppu *ppu) drawScanline(scanline byte) {
	// TODO: these needs to default to full white
	bgTiles := []tile{}
	windowTiles := []tile{}

	// 1. Get current bg and window tiles
	if ppu.bgAndWindowEnable() {
		bgTiles = ppu.getBgTiles()

		if ppu.windowEnable() {
			windowTiles = ppu.getWindowTiles()
		}
	}

	// TODO: whut do here
	// 2. Get obj tiles
	if ppu.objEnable() {
		objTiles := ppu.getObjTiles()
	}

	// 3. Mix them together
	mixed := ppu.mixTiles(bgTiles, windowTiles, objTiles)

	// 4. Extract pixels pertaining to this scanline
	scanlinePixels := ppu.extractScanlinePixels(mixed, scanline)

	// 5. Update our screen representation
	ppu.setScanlinePixels(scanlinePixels, scanline)
}

// sets ppu.bgTiles
// TODO: this seems like it could be factored better too
func (ppu *ppu) getBgTiles() []tile {
	// 1. Choose tile map
	var tileMap []byte
	if ppu.bgTileMapSelect() {
		tileMap = ppu.tileMap1()
	} else {
		tileMap = ppu.tileMap0()
	}

	// 2. Choose lower tile data
	var lowerTileData []byte
	if ppu.bgAndWindowTileDataSelect() {
		lowerTileData = ppu.tileData0()
	} else {
		lowerTileData = ppu.tileData2()
	}

	// 3. Access tiles
	var tiles []tile
	for _, tileIndex := range tileMap {
		tileStart := tileIndex * 16

		var tileData []byte
		if tileIndex <= 127 {
			// use lowerTileData (tileData0 or tileData2)
			tileData = lowerTileData[tileStart : tileStart+16]
		} else {
			// use tileData1
			tileData = ppu.tileData1()[tileStart : tileStart+16]
		}

		tiles = append(tiles, newTile(tileData))
	}

	return tiles
}

func (ppu *ppu) getWindowTiles() []tile { return nil }
func (ppu *ppu) getObjTiles() []tile    { return nil }

// mixes bg, window, and objs together
func (ppu *ppu) mixTiles(bgTiles, windowTiles, objTiles []tile) []tile     { return nil }
func (ppu *ppu) extractScanlinePixels(mixed []tile, scanline byte) []Pixel { return nil }
func (ppu *ppu) setScanlinePixels(pixels []Pixel, scanline byte)           {}

func (ppu *ppu) visiblePixels() []Pixel { return nil }
