package gb

const (
	cpuHz          = 4194304
	cyclesPerFrame = 70224
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
	opcode := cpu.memory.readByte(cpu.pc)

	switch opcode {
	case 0x00:
		return nop
	case 0xC3:
		return jpnn
	case 0xAF:
		return xora
	case 0xA8:
		return xorb
	case 0xA9:
		return xorc
	case 0xAA:
		return xord
	case 0xAB:
		return xore
	case 0xAC:
		return xorh
	case 0xAD:
		return xorl
	case 0xAE:
		return xorhl
	case 0xEE:
		return xorn
	case 0x0A:
		return ldabc
	default:
		return nop
	}
}

// invoked at 60Hz
func (cpu *cpu) tick() {
	elapsedCycles := 0

	for elapsedCycles < cyclesPerFrame {
		instruction := cpu.decode()

		instruction.execute(cpu)

		if !instruction.hasJump {
			cpu.pc += instruction.size
		}

		elapsedCycles += instruction.cycles
		cpu.ppu.scanCycles += instruction.cycles

		if cpu.ppu.scanCycles >= 456 {
			cpu.ppu.tick()
			cpu.ppu.scanCycles = 0
		}
	}
}
