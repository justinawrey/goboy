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
		return jp_a16
	case 0xAF:
		return xor_a
	case 0xA8:
		return xor_b
	case 0xA9:
		return xor_c
	case 0xAA:
		return xor_d
	case 0xAB:
		return xor_e
	case 0xAC:
		return xor_h
	case 0xAD:
		return xor_l
	case 0xAE:
		return xor__hl_
	case 0xEE:
		return xor_d8
	default:
		return nop
	}
}

// invoked at 60Hz
func (cpu *cpu) tick() {
	elapsedCycles := 0

	for elapsedCycles < cyclesPerFrame {
		instruction := cpu.decode()

		currPc := cpu.pc
		instruction.execute(cpu)

		jumped := currPc != cpu.pc
		if !jumped {
			cpu.pc += instruction.size
		}

		var cycles int
		if jumped {
			cycles = instruction.jumpCycles
		} else {
			cycles = instruction.noJumpCycles
		}

		elapsedCycles += cycles
		cpu.ppu.scanCycles += cycles

		if cpu.ppu.scanCycles >= 456 {
			cpu.ppu.tick()
			cpu.ppu.scanCycles = 0
		}
	}
}
