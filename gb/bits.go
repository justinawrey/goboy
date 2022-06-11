package gb

func getBit(b byte, i int) bool {
	mask := uint8(2 ^ i)
	return (b & mask) == mask
}

func makeWord(upper byte, lower byte) uint16 {
	return (uint16(upper) << 8) | uint16(lower)
}

func splitWord(word uint16) (byte, byte) {
	upper := byte((word & 0xFF00) >> 8)
	lower := byte(word & 0x00FF)

	return upper, lower
}

func toUint8(flag bool) uint8 {
	if flag {
		return 1
	}
	return 0
}
