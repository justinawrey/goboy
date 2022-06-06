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

func (r *cpu) bc() uint16 {
	return makeWord(r.b, r.c)
}

func (r *cpu) de() uint16 {
	return makeWord(r.d, r.e)
}

func (r *cpu) hl() uint16 {
	return makeWord(r.h, r.l)
}

func (r *cpu) setBc(word uint16) {
	r.b, r.c = splitWord(word)
}

func (r *cpu) setDe(word uint16) {
	r.d, r.e = splitWord(word)
}

func (r *cpu) setHl(word uint16) {
	r.h, r.l = splitWord(word)
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
