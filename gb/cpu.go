package gb

const (
	cpuHz             = 4194304
	cyclesPerFrame    = 70224
	cyclesPerScanline = 456
)

type cpu struct {
	*memory
	*ppu
	a, b, c, d, e, h, l byte
	sp                  uint16
	pc                  uint16
	flags
}

type flags struct {
	z, n, h, c bool
}

func (cpu *cpu) bc() uint16 {
	return makeWord(cpu.b, cpu.c)
}

func (cpu *cpu) de() uint16 {
	return makeWord(cpu.d, cpu.e)
}

func (cpu *cpu) hl() uint16 {
	return makeWord(cpu.h, cpu.l)
}

func (cpu *cpu) setBc(word uint16) {
	cpu.b, cpu.c = splitWord(word)
}

func (cpu *cpu) setDe(word uint16) {
	cpu.d, cpu.e = splitWord(word)
}

func (cpu *cpu) setHl(word uint16) {
	cpu.h, cpu.l = splitWord(word)
}

// TODO: half carry logic maybe could be streamlined?
func (f *flags) setH8Add(b1 byte, b2 byte) {
	f.h = (((b1 & 0x0f) + (b2 & 0x0f)) & 0x10) == 0x10
}

func (f *flags) setH8Sub(b1 byte, b2 byte) {
	f.h = (((b1 & 0x0f) - (b2 & 0x0f)) & 0x10) == 0x10
}

func (f *flags) setH11Add(w1 uint16, w2 uint16) {
	f.h = (((w1 & 0x0fff) + (w2 & 0x0fff)) & 0x1000) == 0x1000
}

func (f *flags) setH11Sub(w1 uint16, w2 uint16) {
	f.h = (((w1 & 0x0fff) - (w2 & 0x0fff)) & 0x1000) == 0x1000
}

func (f *flags) setH15Add(w1 uint16, w2 uint16) {
	wA := uint32(w1)
	wB := uint32(w1)
	f.h = (((wA & 0xffff) + (wB & 0xffff)) & 0x10000) == 0x10000
}

func (f *flags) setZ(test byte) {
	f.z = test == 0
}

func (cpu *cpu) decode() instruction {
	b1 := cpu.readByte(cpu.pc)

	// 16-bit instructions
	if b1 == 0x10 || b1 == 0xCB {
		b2 := cpu.readByte(cpu.pc + 1)
		return instructionTable16[makeWord(b1, b2)]
	}

	// 8-bit instruction
	return instructionTable8[b1]
}

// invoked at 60Hz
func (cpu *cpu) tick() {
	elapsedCycles := 0
	scanCycles := 0

	for elapsedCycles < cyclesPerFrame {
		cycles := cpu.cpuCycle()

		elapsedCycles += cycles
		scanCycles += cycles

		if scanCycles >= cyclesPerScanline {
			cpu.ppu.tick()
			scanCycles = 0
		}
	}
}

// does a decode, execute, move pc cycle
// returns number of cycles elapsed
func (cpu *cpu) cpuCycle() (cycles int) {
	instruction := cpu.decode()

	currPc := cpu.pc
	instruction.execute(cpu)

	jumped := currPc != cpu.pc
	if !jumped {
		cpu.pc += instruction.size
	}

	if jumped {
		return instruction.jumpCycles
	}

	return instruction.noJumpCycles
}
