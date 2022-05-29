package gb

type instruction struct {
	mnemonic, encoding string
	size uint16
	cycles int
	hasJump bool
	execute func(*cpu)
}

func (i instruction) String () string {
	return i.mnemonic
}

// TODO: add flag ramifications
var ldabc = instruction{
	mnemonic: "ld A,(BC)",
	encoding: "0A",
	size: 1,
	cycles: 8,
	hasJump: false,
	execute: func(cpu *cpu) {},
} 

var nop = instruction{
	mnemonic: "nop",
	encoding: "00",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {},
} 

var jpnn = instruction{
	mnemonic: "jp nn",
	encoding: "C3 nn nn",
	size: 3,
	cycles: 16,
	hasJump: true,
	execute: func(cpu *cpu) {
		cpu.pc = cpu.memory.readWord(cpu.pc + 1)
	},
}

var xora = instruction{
	mnemonic: "xor A",
	encoding: "AF",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.a
	},
}

var xorb = instruction{
	mnemonic: "xor B",
	encoding: "A8",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.b
	},
}

var xorc = instruction{
	mnemonic: "xor C",
	encoding: "A9",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.c
	},
}

var xord = instruction{
	mnemonic: "xor D",
	encoding: "AA",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.d
	},
}

var xore = instruction{
	mnemonic: "xor E",
	encoding: "AB",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.e
	},
}

var xorh = instruction{
	mnemonic: "xor H",
	encoding: "AC",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.h
	},
}

var xorl = instruction{
	mnemonic: "xor L",
	encoding: "AD",
	size: 1,
	cycles: 4,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.l
	},
}

var xorhl = instruction{
	mnemonic: "xor (HL)",
	encoding: "AE",
	size: 1,
	cycles: 8,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a = byte(uint16(cpu.a) ^ cpu.hl())
	},
}

var xorn = instruction{
	mnemonic: "xor n",
	encoding: "EE nn",
	size: 2,
	cycles: 8,
	hasJump: false,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.memory.readByte(cpu.pc + 1)
	},
}
