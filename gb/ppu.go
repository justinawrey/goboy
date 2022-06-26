package gb

const (
	lcdHeight = 144
	lcdWidth  = 160
	numPixels = lcdHeight * lcdWidth
)

type Pixel = int
type ppu struct {
	*memory

	// visible area only
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

func (tile *tile) getPixelsAt(row byte) []Pixel {
	offset := row * 8
	return tile.pixels[offset : offset+8]
}

// constructs a well-formed 'tile' from 16 bytes
// TODO: we don't actually need to do all this
// because we only care about specific rows
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
	if scanline >= 144 {
		return
	}

	mixed := ppu.mixPixels(
		ppu.getBgPixels(scanline),
		ppu.getWindowPixels(scanline),
		ppu.getObjPixels(scanline),
	)

	cropped := ppu.cropPixels(mixed)

	// actually save to our screen representation
	// TODO: maybe this could be more idiomatic
	offset := scanline * lcdWidth
	for i := byte(0); i < lcdWidth; i++ {
		ppu.pixels[offset+i] = cropped[i]
	}
}

func (ppu *ppu) getTileInfo(useTileMap1 bool) (tileMap []byte, lowerTileData []byte) {
	if useTileMap1 {
		tileMap = ppu.tileMap1()
	} else {
		tileMap = ppu.tileMap0()
	}

	if ppu.bgAndWindowTileDataSelect() {
		lowerTileData = ppu.tileData0()
	} else {
		lowerTileData = ppu.tileData2()
	}

	return tileMap, lowerTileData
}

// returns 32 tile objects -- a row
func (ppu *ppu) createTileRow(dataIndices []byte, lowerTileData []byte) []tile {
	tiles := make([]tile, 32)

	for _, i := range dataIndices {
		// each tile is encoded into 16 bytes
		tileStart := i * 16

		var tileData []byte
		if i <= 127 {
			// use lowerTileData (tileData0 or tileData2)
			tileData = lowerTileData[tileStart : tileStart+16]
		} else {
			// always use tileData1
			tileData = ppu.tileData1()[tileStart : tileStart+16]
		}

		tiles = append(tiles, newTile(tileData))
	}

	return tiles
}

func (ppu *ppu) getPixelsFromTiles(tiles []tile, row byte) []Pixel {
	pixels := make([]Pixel, 256)

	for _, tile := range tiles {
		pixels = append(pixels, tile.getPixelsAt(row)...)
	}

	return pixels
}

func (ppu *ppu) getScanlinePixels(scanline byte, useTileMap1 bool) []Pixel {
	// 1. Get tile map info
	tileMap, lowerTileData := ppu.getTileInfo(useTileMap1)

	// 2. Which tiles do we actually care about?
	absoluteY := scanline + ppu.scy.get() // transform to 256x256 space
	tileOffset := (absoluteY / 8) * 32
	dataIndices := tileMap[tileOffset : tileOffset+32]

	// 3. Access and create tiles
	tiles := ppu.createTileRow(dataIndices, lowerTileData)

	// 4. We should now have 32 tile objects (a row of tiles),
	// but we only care about pixels from one row of pixels
	return ppu.getPixelsFromTiles(tiles, absoluteY%8)
}

func (ppu *ppu) getBgPixels(scanline byte) []Pixel {
	return ppu.getScanlinePixels(scanline, ppu.bgTileMapSelect())
}

// TODO: I think this is wrong!!
func (ppu *ppu) getWindowPixels(scanline byte) []Pixel {
	return ppu.getScanlinePixels(scanline, ppu.windowTileMapSelect())
}

func (ppu *ppu) getObjPixels(scanline byte) []Pixel                          { return nil }
func (ppu *ppu) mixPixels(bgPixels, windowPixels, objPixels []Pixel) []Pixel { return nil }

// given an uncropped "row" of 256 pixels, crops it to a visible 160 according to window
func (ppu *ppu) cropPixels(pixels []Pixel) []Pixel {
	// double to account for wrap around
	copied := make([]Pixel, len(pixels))
	copy(pixels, copied)
	pixels = append(pixels, copied...)

	return pixels[ppu.scx.get():lcdWidth]
}
