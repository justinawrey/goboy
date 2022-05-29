package gb

type ppu struct {
	*memory
	scanCycles int
}

// invoked every 456 cycles (every scanline)
func (ppu *ppu) tick() {}
