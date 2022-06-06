package gb

import "bytes"

type memory struct {
	*bytes.Buffer
}

const memSize = 0xffff

func newMemory() *memory {
	buf := make([]byte, 0, memSize)
	return &memory{bytes.NewBuffer(buf)}
}

func (m *memory) readByte(n uint16) byte {
	return m.Bytes()[n]
}

func (m *memory) readWord(n uint16) uint16 {
	upper := m.readByte(n + 1)
	lower := m.readByte(n)
	return makeWord(upper, lower)
}

func (m *memory) writeByte(pos uint16, b byte) {
	m.Bytes()[pos] = b
}

func (m *memory) writeWord(pos uint16, word uint16) {
	upper, lower := splitWord(word)
	m.writeByte(pos, lower)
	m.writeByte(pos+1, upper)
}
