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
	buf := m.Bytes()
	return makeWord(buf[n+1], buf[n])
}
