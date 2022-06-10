package gb

import "math/bits"

// TODO: maybe trim this down, binary is pretty big holding
// all of this memory
type Instruction struct {
	Mnemonic     string
	Encoding     string
	size         int
	jumpCycles   int
	noJumpCycles int
	flags        string
	Implemented  bool
	execute      func(*cpu)
}

func (i Instruction) String() string {
	return i.Mnemonic
}

// 8-bit Instructions
var nop = Instruction{
	Mnemonic:     "NOP",
	Encoding:     "0x00",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute:      func(cpu *cpu) {},
}

var ld_bc__d16 = Instruction{
	Mnemonic:     "LD BC, d16",
	Encoding:     "0x01",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setBc(cpu.readWord(cpu.pc + 1))
	},
}

var ld__bc___a = Instruction{
	Mnemonic:     "LD (BC), A",
	Encoding:     "0x02",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.bc(), cpu.a)
	},
}

var inc_bc = Instruction{
	Mnemonic:     "INC BC",
	Encoding:     "0x03",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setBc(cpu.bc() + 1)
	},
}

var inc_b = Instruction{
	Mnemonic:     "INC B",
	Encoding:     "0x04",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.b, 1)
		cpu.b += 1
		cpu.setZ(cpu.b)
	},
}

var dec_b = Instruction{
	Mnemonic:     "DEC B",
	Encoding:     "0x05",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.b, 1)
		cpu.b -= 1
		cpu.setZ(cpu.b)
	},
}

var ld_b__d8 = Instruction{
	Mnemonic:     "LD B, d8",
	Encoding:     "0x06",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.readByte(cpu.pc + 1)
	},
}

var rlca = Instruction{
	Mnemonic:     "RLCA",
	Encoding:     "0x07",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "0 0 0 A7",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.z = false
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = (cpu.a & 0x80) == 0x80
		cpu.a = bits.RotateLeft8(cpu.a, 1)
	},
}

var ld__a16___sp = Instruction{
	Mnemonic:     "LD (a16), SP",
	Encoding:     "0x08",
	size:         3,
	jumpCycles:   20,
	noJumpCycles: 20,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		upper, lower := splitWord(cpu.sp)
		address := makeWord(cpu.readByte(cpu.pc+1), cpu.readByte(cpu.pc+2))
		cpu.writeByte(address, lower)
		cpu.writeByte(address+1, upper)
	},
}

var add_hl__bc = Instruction{
	Mnemonic:     "ADD HL, BC",
	Encoding:     "0x09",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH11Add(cpu.hl(), cpu.bc())
		cpu.setC15Add(cpu.hl(), cpu.bc())
		cpu.setHl(cpu.hl() + cpu.bc())
	},
}

var ld_a___bc_ = Instruction{
	Mnemonic:     "LD A, (BC)",
	Encoding:     "0x0A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.bc())
	},
}

var dec_bc = Instruction{
	Mnemonic:     "DEC BC",
	Encoding:     "0x0B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setBc(cpu.bc() - 1)
	},
}

var inc_c = Instruction{
	Mnemonic:     "INC C",
	Encoding:     "0x0C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.c, 1)
		cpu.c += 1
		cpu.setZ(cpu.c)
	},
}

var dec_c = Instruction{
	Mnemonic:     "DEC C",
	Encoding:     "0x0D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.c, 1)
		cpu.c -= 1
		cpu.setZ(cpu.c)
	},
}

var ld_c__d8 = Instruction{
	Mnemonic:     "LD C, d8",
	Encoding:     "0x0E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.readByte(cpu.pc + 1)
	},
}

var rrca = Instruction{
	Mnemonic:     "RRCA",
	Encoding:     "0x0F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "0 0 0 A0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.z = false
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = (cpu.a & 0x01) == 0x01
		cpu.a = bits.RotateLeft8(cpu.a, -1)
	},
}

var stop = Instruction{
	Mnemonic:     "STOP",
	Encoding:     "0x1000",
	size:         2,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_de__d16 = Instruction{
	Mnemonic:     "LD DE, d16",
	Encoding:     "0x11",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setDe(cpu.readWord(cpu.pc + 1))
	},
}

var ld__de___a = Instruction{
	Mnemonic:     "LD (DE), A",
	Encoding:     "0x12",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.de(), cpu.a)
	},
}

var inc_de = Instruction{
	Mnemonic:     "INC DE",
	Encoding:     "0x13",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setDe(cpu.de() + 1)
	},
}

var inc_d = Instruction{
	Mnemonic:     "INC D",
	Encoding:     "0x14",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.d, 1)
		cpu.d += 1
		cpu.setZ(cpu.d)
	},
}

var dec_d = Instruction{
	Mnemonic:     "DEC D",
	Encoding:     "0x15",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.d, 1)
		cpu.d -= 1
		cpu.setZ(cpu.d)
	},
}

var ld_d__d8 = Instruction{
	Mnemonic:     "LD D, d8",
	Encoding:     "0x16",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.readByte(cpu.pc + 1)
	},
}

var rla = Instruction{
	Mnemonic:     "RLA",
	Encoding:     "0x17",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "0 0 0 A7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jr_s8 = Instruction{
	Mnemonic:     "JR s8",
	Encoding:     "0x18",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_hl__de = Instruction{
	Mnemonic:     "ADD HL, DE",
	Encoding:     "0x19",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___de_ = Instruction{
	Mnemonic:     "LD A, (DE)",
	Encoding:     "0x1A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.de())
	},
}

var dec_de = Instruction{
	Mnemonic:     "DEC DE",
	Encoding:     "0x1B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setDe(cpu.de() - 1)
	},
}

var inc_e = Instruction{
	Mnemonic:     "INC E",
	Encoding:     "0x1C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.e, 1)
		cpu.e += 1
		cpu.setZ(cpu.e)
	},
}

var dec_e = Instruction{
	Mnemonic:     "DEC E",
	Encoding:     "0x1D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.e, 1)
		cpu.e -= 1
		cpu.setZ(cpu.e)
	},
}

var ld_e__d8 = Instruction{
	Mnemonic:     "LD E, d8",
	Encoding:     "0x1E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.readByte(cpu.pc + 1)
	},
}

var rra = Instruction{
	Mnemonic:     "RRA",
	Encoding:     "0x1F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "0 0 0 A0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jr_nz__s8 = Instruction{
	Mnemonic:     "JR NZ, s8",
	Encoding:     "0x20",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_hl__d16 = Instruction{
	Mnemonic:     "LD HL, d16",
	Encoding:     "0x21",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setHl(cpu.readWord(cpu.pc + 1))
	},
}

var ld__hlp___a = Instruction{
	Mnemonic:     "LD (HL+), A",
	Encoding:     "0x22",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
		cpu.setHl(cpu.hl() + 1)
	},
}

var inc_hl = Instruction{
	Mnemonic:     "INC HL",
	Encoding:     "0x23",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setHl(cpu.hl() + 1)
	},
}

var inc_h = Instruction{
	Mnemonic:     "INC H",
	Encoding:     "0x24",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.h, 1)
		cpu.h += 1
		cpu.setZ(cpu.h)
	},
}

var dec_h = Instruction{
	Mnemonic:     "DEC H",
	Encoding:     "0x25",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.h, 1)
		cpu.h -= 1
		cpu.setZ(cpu.h)
	},
}

var ld_h__d8 = Instruction{
	Mnemonic:     "LD H, d8",
	Encoding:     "0x26",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.readByte(cpu.pc + 1)
	},
}

var daa = Instruction{
	Mnemonic:     "DAA",
	Encoding:     "0x27",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z - 0 CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jr_z__s8 = Instruction{
	Mnemonic:     "JR Z, s8",
	Encoding:     "0x28",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_hl__hl = Instruction{
	Mnemonic:     "ADD HL, HL",
	Encoding:     "0x29",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___hlp_ = Instruction{
	Mnemonic:     "LD A, (HL+)",
	Encoding:     "0x2A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
		cpu.setHl(cpu.hl() + 1)
	},
}

var dec_hl = Instruction{
	Mnemonic:     "DEC HL",
	Encoding:     "0x2B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.setHl(cpu.hl() - 1)
	},
}

var inc_l = Instruction{
	Mnemonic:     "INC L",
	Encoding:     "0x2C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.l, 1)
		cpu.l += 1
		cpu.setZ(cpu.l)
	},
}

var dec_l = Instruction{
	Mnemonic:     "DEC L",
	Encoding:     "0x2D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.l, 1)
		cpu.l -= 1
		cpu.setZ(cpu.l)
	},
}

var ld_l__d8 = Instruction{
	Mnemonic:     "LD L, d8",
	Encoding:     "0x2E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.readByte(cpu.pc + 1)
	},
}

var cpl = Instruction{
	Mnemonic:     "CPL",
	Encoding:     "0x2F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- 1 1 -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.flags.h = true
		cpu.a = cpu.a ^ 0xff
	},
}

var jr_nc__s8 = Instruction{
	Mnemonic:     "JR NC, s8",
	Encoding:     "0x30",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_sp__d16 = Instruction{
	Mnemonic:     "LD SP, d16",
	Encoding:     "0x31",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.sp = cpu.readWord(cpu.pc + 1)
	},
}

var ld__hlm___a = Instruction{
	Mnemonic:     "LD (HL-), A",
	Encoding:     "0x32",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
		cpu.setHl(cpu.hl() - 1)
	},
}

var inc_sp = Instruction{
	Mnemonic:     "INC SP",
	Encoding:     "0x33",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.sp += 1
	},
}

var inc__hl_ = Instruction{
	Mnemonic:     "INC (HL)",
	Encoding:     "0x34",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		address := cpu.hl()
		b := cpu.readByte(address)
		cpu.setH3Add(b, 1)
		cpu.writeByte(address, b+1)
		cpu.setZ(cpu.readByte(cpu.hl()))
	},
}

var dec__hl_ = Instruction{
	Mnemonic:     "DEC (HL)",
	Encoding:     "0x35",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		address := cpu.hl()
		b := cpu.readByte(address)
		cpu.setH3Sub(b, 1)
		cpu.writeByte(address, b-1)
		cpu.setZ(cpu.readByte(cpu.hl()))
	},
}

var ld__hl___d8 = Instruction{
	Mnemonic:     "LD (HL), d8",
	Encoding:     "0x36",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.readByte(cpu.pc+1))
	},
}

var scf = Instruction{
	Mnemonic:     "SCF",
	Encoding:     "0x37",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- 0 0 1",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = true
	},
}

var jr_c__s8 = Instruction{
	Mnemonic:     "JR C, s8",
	Encoding:     "0x38",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_hl__sp = Instruction{
	Mnemonic:     "ADD HL, SP",
	Encoding:     "0x39",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___hlm_ = Instruction{
	Mnemonic:     "LD A, (HL-)",
	Encoding:     "0x3A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
		cpu.setHl(cpu.hl() - 1)
	},
}

var dec_sp = Instruction{
	Mnemonic:     "DEC SP",
	Encoding:     "0x3B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.sp = cpu.hl() - 1
	},
}

var inc_a = Instruction{
	Mnemonic:     "INC A",
	Encoding:     "0x3C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, 1)
		cpu.a += 1
		cpu.setZ(cpu.a)
	},
}

var dec_a = Instruction{
	Mnemonic:     "DEC A",
	Encoding:     "0x3D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, 1)
		cpu.a -= 1
		cpu.setZ(cpu.a)
	},
}

var ld_a__d8 = Instruction{
	Mnemonic:     "LD A, d8",
	Encoding:     "0x3E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.pc + 1)
	},
}

var ccf = Instruction{
	Mnemonic:     "CCF",
	Encoding:     "0x3F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- 0 0 !CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = !cpu.flags.c
	},
}

var ld_b__b = Instruction{
	Mnemonic:     "LD B, B",
	Encoding:     "0x40",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.b
	},
}

var ld_b__c = Instruction{
	Mnemonic:     "LD B, C",
	Encoding:     "0x41",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.c
	},
}

var ld_b__d = Instruction{
	Mnemonic:     "LD B, D",
	Encoding:     "0x42",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.d
	},
}

var ld_b__e = Instruction{
	Mnemonic:     "LD B, E",
	Encoding:     "0x43",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.e
	},
}

var ld_b__h = Instruction{
	Mnemonic:     "LD B, H",
	Encoding:     "0x44",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.h
	},
}

var ld_b__l = Instruction{
	Mnemonic:     "LD B, L",
	Encoding:     "0x45",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.l
	},
}

var ld_b___hl_ = Instruction{
	Mnemonic:     "LD B, (HL)",
	Encoding:     "0x46",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.readByte(cpu.hl())
	},
}

var ld_b__a = Instruction{
	Mnemonic:     "LD B, A",
	Encoding:     "0x47",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.b = cpu.a
	},
}

var ld_c__b = Instruction{
	Mnemonic:     "LD C, B",
	Encoding:     "0x48",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.b
	},
}

var ld_c__c = Instruction{
	Mnemonic:     "LD C, C",
	Encoding:     "0x49",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.c
	},
}

var ld_c__d = Instruction{
	Mnemonic:     "LD C, D",
	Encoding:     "0x4A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.d
	},
}

var ld_c__e = Instruction{
	Mnemonic:     "LD C, E",
	Encoding:     "0x4B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.e
	},
}

var ld_c__h = Instruction{
	Mnemonic:     "LD C, H",
	Encoding:     "0x4C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.h
	},
}

var ld_c__l = Instruction{
	Mnemonic:     "LD C, L",
	Encoding:     "0x4D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.l
	},
}

var ld_c___hl_ = Instruction{
	Mnemonic:     "LD C, (HL)",
	Encoding:     "0x4E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.readByte(cpu.hl())
	},
}

var ld_c__a = Instruction{
	Mnemonic:     "LD C, A",
	Encoding:     "0x4F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.c = cpu.a
	},
}

var ld_d__b = Instruction{
	Mnemonic:     "LD D, B",
	Encoding:     "0x50",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.b
	},
}

var ld_d__c = Instruction{
	Mnemonic:     "LD D, C",
	Encoding:     "0x51",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.c
	},
}

var ld_d__d = Instruction{
	Mnemonic:     "LD D, D",
	Encoding:     "0x52",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.d
	},
}

var ld_d__e = Instruction{
	Mnemonic:     "LD D, E",
	Encoding:     "0x53",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.e
	},
}

var ld_d__h = Instruction{
	Mnemonic:     "LD D, H",
	Encoding:     "0x54",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.h
	},
}

var ld_d__l = Instruction{
	Mnemonic:     "LD D, L",
	Encoding:     "0x55",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.l
	},
}

var ld_d___hl_ = Instruction{
	Mnemonic:     "LD D, (HL)",
	Encoding:     "0x56",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.readByte(cpu.hl())
	},
}

var ld_d__a = Instruction{
	Mnemonic:     "LD D, A",
	Encoding:     "0x57",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.d = cpu.a
	},
}

var ld_e__b = Instruction{
	Mnemonic:     "LD E, B",
	Encoding:     "0x58",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.b
	},
}

var ld_e__c = Instruction{
	Mnemonic:     "LD E, C",
	Encoding:     "0x59",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.c
	},
}

var ld_e__d = Instruction{
	Mnemonic:     "LD E, D",
	Encoding:     "0x5A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.d
	},
}

var ld_e__e = Instruction{
	Mnemonic:     "LD E, E",
	Encoding:     "0x5B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.e
	},
}

var ld_e__h = Instruction{
	Mnemonic:     "LD E, H",
	Encoding:     "0x5C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.h
	},
}

var ld_e__l = Instruction{
	Mnemonic:     "LD E, L",
	Encoding:     "0x5D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.l
	},
}

var ld_e___hl_ = Instruction{
	Mnemonic:     "LD E, (HL)",
	Encoding:     "0x5E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.readByte(cpu.hl())
	},
}

var ld_e__a = Instruction{
	Mnemonic:     "LD E, A",
	Encoding:     "0x5F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.e = cpu.a
	},
}

var ld_h__b = Instruction{
	Mnemonic:     "LD H, B",
	Encoding:     "0x60",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.b
	},
}

var ld_h__c = Instruction{
	Mnemonic:     "LD H, C",
	Encoding:     "0x61",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.c
	},
}

var ld_h__d = Instruction{
	Mnemonic:     "LD H, D",
	Encoding:     "0x62",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.d
	},
}

var ld_h__e = Instruction{
	Mnemonic:     "LD H, E",
	Encoding:     "0x63",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.e
	},
}

var ld_h__h = Instruction{
	Mnemonic:     "LD H, H",
	Encoding:     "0x64",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.h
	},
}

var ld_h__l = Instruction{
	Mnemonic:     "LD H, L",
	Encoding:     "0x65",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.l
	},
}

var ld_h___hl_ = Instruction{
	Mnemonic:     "LD H, (HL)",
	Encoding:     "0x66",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.readByte(cpu.hl())
	},
}

var ld_h__a = Instruction{
	Mnemonic:     "LD H, A",
	Encoding:     "0x67",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.h = cpu.a
	},
}

var ld_l__b = Instruction{
	Mnemonic:     "LD L, B",
	Encoding:     "0x68",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.b
	},
}

var ld_l__c = Instruction{
	Mnemonic:     "LD L, C",
	Encoding:     "0x69",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.c
	},
}

var ld_l__d = Instruction{
	Mnemonic:     "LD L, D",
	Encoding:     "0x6A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.d
	},
}

var ld_l__e = Instruction{
	Mnemonic:     "LD L, E",
	Encoding:     "0x6B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.e
	},
}

var ld_l__h = Instruction{
	Mnemonic:     "LD L, H",
	Encoding:     "0x6C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.h
	},
}

var ld_l__l = Instruction{
	Mnemonic:     "LD L, L",
	Encoding:     "0x6D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.l
	},
}

var ld_l___hl_ = Instruction{
	Mnemonic:     "LD L, (HL)",
	Encoding:     "0x6E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.readByte(cpu.hl())
	},
}

var ld_l__a = Instruction{
	Mnemonic:     "LD L, A",
	Encoding:     "0x6F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.l = cpu.a
	},
}

var ld__hl___b = Instruction{
	Mnemonic:     "LD (HL), B",
	Encoding:     "0x70",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.b)
	},
}

var ld__hl___c = Instruction{
	Mnemonic:     "LD (HL), C",
	Encoding:     "0x71",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.c)
	},
}

var ld__hl___d = Instruction{
	Mnemonic:     "LD (HL), D",
	Encoding:     "0x72",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.d)
	},
}

var ld__hl___e = Instruction{
	Mnemonic:     "LD (HL), E",
	Encoding:     "0x73",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.e)
	},
}

var ld__hl___h = Instruction{
	Mnemonic:     "LD (HL), H",
	Encoding:     "0x74",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.h)
	},
}

var ld__hl___l = Instruction{
	Mnemonic:     "LD (HL), L",
	Encoding:     "0x75",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.l)
	},
}

var halt = Instruction{
	Mnemonic:     "HALT",
	Encoding:     "0x76",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld__hl___a = Instruction{
	Mnemonic:     "LD (HL), A",
	Encoding:     "0x77",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
	},
}

var ld_a__b = Instruction{
	Mnemonic:     "LD A, B",
	Encoding:     "0x78",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.b
	},
}

var ld_a__c = Instruction{
	Mnemonic:     "LD A, C",
	Encoding:     "0x79",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.c
	},
}

var ld_a__d = Instruction{
	Mnemonic:     "LD A, D",
	Encoding:     "0x7A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.d
	},
}

var ld_a__e = Instruction{
	Mnemonic:     "LD A, E",
	Encoding:     "0x7B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.e
	},
}

var ld_a__h = Instruction{
	Mnemonic:     "LD A, H",
	Encoding:     "0x7C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.h
	},
}

var ld_a__l = Instruction{
	Mnemonic:     "LD A, L",
	Encoding:     "0x7D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.l
	},
}

var ld_a___hl_ = Instruction{
	Mnemonic:     "LD A, (HL)",
	Encoding:     "0x7E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
	},
}

var ld_a__a = Instruction{
	Mnemonic:     "LD A, A",
	Encoding:     "0x7F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = cpu.a
	},
}

var add_a__b = Instruction{
	Mnemonic:     "ADD A, B",
	Encoding:     "0x80",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.b)
		cpu.setC7Add(cpu.a, cpu.b)
		cpu.a += cpu.b
		cpu.setZ(cpu.a)
	},
}

var add_a__c = Instruction{
	Mnemonic:     "ADD A, C",
	Encoding:     "0x81",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.c)
		cpu.setC7Add(cpu.a, cpu.c)
		cpu.a += cpu.c
		cpu.setZ(cpu.a)
	},
}

var add_a__d = Instruction{
	Mnemonic:     "ADD A, D",
	Encoding:     "0x82",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.d)
		cpu.setC7Add(cpu.a, cpu.d)
		cpu.a += cpu.d
		cpu.setZ(cpu.a)
	},
}

var add_a__e = Instruction{
	Mnemonic:     "ADD A, E",
	Encoding:     "0x83",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.e)
		cpu.setC7Add(cpu.a, cpu.e)
		cpu.a += cpu.e
		cpu.setZ(cpu.a)
	},
}

var add_a__h = Instruction{
	Mnemonic:     "ADD A, H",
	Encoding:     "0x84",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.h)
		cpu.setC7Add(cpu.a, cpu.h)
		cpu.a += cpu.h
		cpu.setZ(cpu.a)
	},
}

var add_a__l = Instruction{
	Mnemonic:     "ADD A, L",
	Encoding:     "0x85",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.l)
		cpu.setC7Add(cpu.a, cpu.l)
		cpu.a += cpu.l
		cpu.setZ(cpu.a)
	},
}

var add_a___hl_ = Instruction{
	Mnemonic:     "ADD A, (HL)",
	Encoding:     "0x86",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_a__a = Instruction{
	Mnemonic:     "ADD A, A",
	Encoding:     "0x87",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, cpu.a)
		cpu.setC7Add(cpu.a, cpu.a)
		cpu.a += cpu.a
		cpu.setZ(cpu.a)
	},
}

var adc_a__b = Instruction{
	Mnemonic:     "ADC A, B",
	Encoding:     "0x88",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__c = Instruction{
	Mnemonic:     "ADC A, C",
	Encoding:     "0x89",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__d = Instruction{
	Mnemonic:     "ADC A, D",
	Encoding:     "0x8A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__e = Instruction{
	Mnemonic:     "ADC A, E",
	Encoding:     "0x8B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__h = Instruction{
	Mnemonic:     "ADC A, H",
	Encoding:     "0x8C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__l = Instruction{
	Mnemonic:     "ADC A, L",
	Encoding:     "0x8D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a___hl_ = Instruction{
	Mnemonic:     "ADC A, (HL)",
	Encoding:     "0x8E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__a = Instruction{
	Mnemonic:     "ADC A, A",
	Encoding:     "0x8F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sub_b = Instruction{
	Mnemonic:     "SUB B",
	Encoding:     "0x90",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.b)
		cpu.setC7Sub(cpu.a, cpu.b)
		cpu.a -= cpu.b
		cpu.setZ(cpu.a)
	},
}

var sub_c = Instruction{
	Mnemonic:     "SUB C",
	Encoding:     "0x91",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.c)
		cpu.setC7Sub(cpu.a, cpu.c)
		cpu.a -= cpu.c
		cpu.setZ(cpu.a)
	},
}

var sub_d = Instruction{
	Mnemonic:     "SUB D",
	Encoding:     "0x92",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.d)
		cpu.setC7Sub(cpu.a, cpu.d)
		cpu.a -= cpu.d
		cpu.setZ(cpu.a)
	},
}

var sub_e = Instruction{
	Mnemonic:     "SUB E",
	Encoding:     "0x93",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.e)
		cpu.setC7Sub(cpu.a, cpu.e)
		cpu.a -= cpu.e
		cpu.setZ(cpu.a)
	},
}

var sub_h = Instruction{
	Mnemonic:     "SUB H",
	Encoding:     "0x94",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.h)
		cpu.setC7Sub(cpu.a, cpu.h)
		cpu.a -= cpu.h
		cpu.setZ(cpu.a)
	},
}

var sub_l = Instruction{
	Mnemonic:     "SUB L",
	Encoding:     "0x95",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.l)
		cpu.setC7Sub(cpu.a, cpu.l)
		cpu.a -= cpu.l
		cpu.setZ(cpu.a)
	},
}

var sub__hl_ = Instruction{
	Mnemonic:     "SUB (HL)",
	Encoding:     "0x96",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sub_a = Instruction{
	Mnemonic:     "SUB A",
	Encoding:     "0x97",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = true
		cpu.setH3Sub(cpu.a, cpu.a)
		cpu.setC7Sub(cpu.a, cpu.a)
		cpu.a -= cpu.a
		cpu.setZ(cpu.a)
	},
}

var sbc_a__b = Instruction{
	Mnemonic:     "SBC A, B",
	Encoding:     "0x98",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__c = Instruction{
	Mnemonic:     "SBC A, C",
	Encoding:     "0x99",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__d = Instruction{
	Mnemonic:     "SBC A, D",
	Encoding:     "0x9A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__e = Instruction{
	Mnemonic:     "SBC A, E",
	Encoding:     "0x9B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__h = Instruction{
	Mnemonic:     "SBC A, H",
	Encoding:     "0x9C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__l = Instruction{
	Mnemonic:     "SBC A, L",
	Encoding:     "0x9D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a___hl_ = Instruction{
	Mnemonic:     "SBC A, (HL)",
	Encoding:     "0x9E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__a = Instruction{
	Mnemonic:     "SBC A, A",
	Encoding:     "0x9F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var and_b = Instruction{
	Mnemonic:     "AND B",
	Encoding:     "0xA0",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.b
		cpu.setZ(cpu.a)
	},
}

var and_c = Instruction{
	Mnemonic:     "AND C",
	Encoding:     "0xA1",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.c
		cpu.setZ(cpu.a)
	},
}

var and_d = Instruction{
	Mnemonic:     "AND D",
	Encoding:     "0xA2",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.d
		cpu.setZ(cpu.a)
	},
}

var and_e = Instruction{
	Mnemonic:     "AND E",
	Encoding:     "0xA3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.e
		cpu.setZ(cpu.a)
	},
}

var and_h = Instruction{
	Mnemonic:     "AND H",
	Encoding:     "0xA4",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.h
		cpu.setZ(cpu.a)
	},
}

var and_l = Instruction{
	Mnemonic:     "AND L",
	Encoding:     "0xA5",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.l
		cpu.setZ(cpu.a)
	},
}

var and__hl_ = Instruction{
	Mnemonic:     "AND (HL)",
	Encoding:     "0xA6",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 1 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var and_a = Instruction{
	Mnemonic:     "AND A",
	Encoding:     "0xA7",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 1 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.flags.n = false
		cpu.flags.h = true
		cpu.flags.c = false
		cpu.a &= cpu.a
		cpu.setZ(cpu.a)
	},
}

var xor_b = Instruction{
	Mnemonic:     "XOR B",
	Encoding:     "0xA8",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.b
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_c = Instruction{
	Mnemonic:     "XOR C",
	Encoding:     "0xA9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.c
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_d = Instruction{
	Mnemonic:     "XOR D",
	Encoding:     "0xAA",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.d
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_e = Instruction{
	Mnemonic:     "XOR E",
	Encoding:     "0xAB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.e
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_h = Instruction{
	Mnemonic:     "XOR H",
	Encoding:     "0xAC",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.h
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_l = Instruction{
	Mnemonic:     "XOR L",
	Encoding:     "0xAD",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.l
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor__hl_ = Instruction{
	Mnemonic:     "XOR (HL)",
	Encoding:     "0xAE",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a = byte(uint16(cpu.a) ^ cpu.hl())
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var xor_a = Instruction{
	Mnemonic:     "XOR A",
	Encoding:     "0xAF",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.a
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var or_b = Instruction{
	Mnemonic:     "OR B",
	Encoding:     "0xB0",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_c = Instruction{
	Mnemonic:     "OR C",
	Encoding:     "0xB1",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_d = Instruction{
	Mnemonic:     "OR D",
	Encoding:     "0xB2",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_e = Instruction{
	Mnemonic:     "OR E",
	Encoding:     "0xB3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_h = Instruction{
	Mnemonic:     "OR H",
	Encoding:     "0xB4",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_l = Instruction{
	Mnemonic:     "OR L",
	Encoding:     "0xB5",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or__hl_ = Instruction{
	Mnemonic:     "OR (HL)",
	Encoding:     "0xB6",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_a = Instruction{
	Mnemonic:     "OR A",
	Encoding:     "0xB7",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_b = Instruction{
	Mnemonic:     "CP B",
	Encoding:     "0xB8",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_c = Instruction{
	Mnemonic:     "CP C",
	Encoding:     "0xB9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_d = Instruction{
	Mnemonic:     "CP D",
	Encoding:     "0xBA",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_e = Instruction{
	Mnemonic:     "CP E",
	Encoding:     "0xBB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_h = Instruction{
	Mnemonic:     "CP H",
	Encoding:     "0xBC",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_l = Instruction{
	Mnemonic:     "CP L",
	Encoding:     "0xBD",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp__hl_ = Instruction{
	Mnemonic:     "CP (HL)",
	Encoding:     "0xBE",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_a = Instruction{
	Mnemonic:     "CP A",
	Encoding:     "0xBF",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ret_nz = Instruction{
	Mnemonic:     "RET NZ",
	Encoding:     "0xC0",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var pop_bc = Instruction{
	Mnemonic:     "POP BC",
	Encoding:     "0xC1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_nz__a16 = Instruction{
	Mnemonic:     "JP NZ, a16",
	Encoding:     "0xC2",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_a16 = Instruction{
	Mnemonic:     "JP a16",
	Encoding:     "0xC3",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.pc = cpu.readWord(cpu.pc + 1)
	},
}

var call_nz__a16 = Instruction{
	Mnemonic:     "CALL NZ, a16",
	Encoding:     "0xC4",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var push_bc = Instruction{
	Mnemonic:     "PUSH BC",
	Encoding:     "0xC5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_a__d8 = Instruction{
	Mnemonic:     "ADD A, d8",
	Encoding:     "0xC6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		d8 := cpu.readByte(cpu.pc + 1)
		cpu.flags.n = false
		cpu.setH3Add(cpu.a, d8)
		cpu.setC7Add(cpu.a, d8)
		cpu.a += d8
		cpu.setZ(cpu.a)
	},
}

var rst_0 = Instruction{
	Mnemonic:     "RST 0",
	Encoding:     "0xC7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ret_z = Instruction{
	Mnemonic:     "RET Z",
	Encoding:     "0xC8",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ret = Instruction{
	Mnemonic:     "RET",
	Encoding:     "0xC9",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_z__a16 = Instruction{
	Mnemonic:     "JP Z, a16",
	Encoding:     "0xCA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var call_z__a16 = Instruction{
	Mnemonic:     "CALL Z, a16",
	Encoding:     "0xCC",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var call_a16 = Instruction{
	Mnemonic:     "CALL a16",
	Encoding:     "0xCD",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 24,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var adc_a__d8 = Instruction{
	Mnemonic:     "ADC A, d8",
	Encoding:     "0xCE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rst_1 = Instruction{
	Mnemonic:     "RST 1",
	Encoding:     "0xCF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ret_nc = Instruction{
	Mnemonic:     "RET NC",
	Encoding:     "0xD0",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var pop_de = Instruction{
	Mnemonic:     "POP DE",
	Encoding:     "0xD1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_nc__a16 = Instruction{
	Mnemonic:     "JP NC, a16",
	Encoding:     "0xD2",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var call_nc__a16 = Instruction{
	Mnemonic:     "CALL NC, a16",
	Encoding:     "0xD4",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var push_de = Instruction{
	Mnemonic:     "PUSH DE",
	Encoding:     "0xD5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sub_d8 = Instruction{
	Mnemonic:     "SUB d8",
	Encoding:     "0xD6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  true,
	execute: func(cpu *cpu) {
		d8 := cpu.readByte(cpu.pc + 1)
		cpu.flags.n = true
		cpu.setH3Add(cpu.a, d8)
		cpu.setC7Add(cpu.a, d8)
		cpu.a -= d8
		cpu.setZ(cpu.a)
	},
}

var rst_2 = Instruction{
	Mnemonic:     "RST 2",
	Encoding:     "0xD7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ret_c = Instruction{
	Mnemonic:     "RET C",
	Encoding:     "0xD8",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var reti = Instruction{
	Mnemonic:     "RETI",
	Encoding:     "0xD9",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_c__a16 = Instruction{
	Mnemonic:     "JP C, a16",
	Encoding:     "0xDA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var call_c__a16 = Instruction{
	Mnemonic:     "CALL C, a16",
	Encoding:     "0xDC",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sbc_a__d8 = Instruction{
	Mnemonic:     "SBC A, d8",
	Encoding:     "0xDE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rst_3 = Instruction{
	Mnemonic:     "RST 3",
	Encoding:     "0xDF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld__a8___a = Instruction{
	Mnemonic:     "LD (a8), A",
	Encoding:     "0xE0",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var pop_hl = Instruction{
	Mnemonic:     "POP HL",
	Encoding:     "0xE1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld__c___a = Instruction{
	Mnemonic:     "LD (C), A",
	Encoding:     "0xE2",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var push_hl = Instruction{
	Mnemonic:     "PUSH HL",
	Encoding:     "0xE5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var and_d8 = Instruction{
	Mnemonic:     "AND d8",
	Encoding:     "0xE6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 1 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rst_4 = Instruction{
	Mnemonic:     "RST 4",
	Encoding:     "0xE7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var add_sp__s8 = Instruction{
	Mnemonic:     "ADD SP, s8",
	Encoding:     "0xE8",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "0 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var jp_hl = Instruction{
	Mnemonic:     "JP HL",
	Encoding:     "0xE9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld__a16___a = Instruction{
	Mnemonic:     "LD (a16), A",
	Encoding:     "0xEA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var xor_d8 = Instruction{
	Mnemonic:     "XOR d8",
	Encoding:     "0xEE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  true,
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.memory.readByte(cpu.pc + 1)
		cpu.setZ(cpu.a)
		cpu.flags.n = false
		cpu.flags.h = false
		cpu.flags.c = false
	},
}

var rst_5 = Instruction{
	Mnemonic:     "RST 5",
	Encoding:     "0xEF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___a8_ = Instruction{
	Mnemonic:     "LD A, (a8)",
	Encoding:     "0xF0",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var pop_af = Instruction{
	Mnemonic:     "POP AF",
	Encoding:     "0xF1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___c_ = Instruction{
	Mnemonic:     "LD A, (C)",
	Encoding:     "0xF2",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var di = Instruction{
	Mnemonic:     "DI",
	Encoding:     "0xF3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var push_af = Instruction{
	Mnemonic:     "PUSH AF",
	Encoding:     "0xF5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var or_d8 = Instruction{
	Mnemonic:     "OR d8",
	Encoding:     "0xF6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rst_6 = Instruction{
	Mnemonic:     "RST 6",
	Encoding:     "0xF7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_hl__sp_s8 = Instruction{
	Mnemonic:     "LD HL, SP+s8",
	Encoding:     "0xF8",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "0 0 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_sp__hl = Instruction{
	Mnemonic:     "LD SP, HL",
	Encoding:     "0xF9",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ld_a___a16_ = Instruction{
	Mnemonic:     "LD A, (a16)",
	Encoding:     "0xFA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var ei = Instruction{
	Mnemonic:     "EI",
	Encoding:     "0xFB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var cp_d8 = Instruction{
	Mnemonic:     "CP d8",
	Encoding:     "0xFE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 1 H CY",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rst_7 = Instruction{
	Mnemonic:     "RST 7",
	Encoding:     "0xFF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

// 16-bit Instructions
var rlc_b = Instruction{
	Mnemonic:     "RLC B",
	Encoding:     "0xCB00",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_c = Instruction{
	Mnemonic:     "RLC C",
	Encoding:     "0xCB01",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_d = Instruction{
	Mnemonic:     "RLC D",
	Encoding:     "0xCB02",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_e = Instruction{
	Mnemonic:     "RLC E",
	Encoding:     "0xCB03",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_h = Instruction{
	Mnemonic:     "RLC H",
	Encoding:     "0xCB04",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_l = Instruction{
	Mnemonic:     "RLC L",
	Encoding:     "0xCB05",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc__hl_ = Instruction{
	Mnemonic:     "RLC (HL)",
	Encoding:     "0xCB06",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rlc_a = Instruction{
	Mnemonic:     "RLC A",
	Encoding:     "0xCB07",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_b = Instruction{
	Mnemonic:     "RRC B",
	Encoding:     "0xCB08",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_c = Instruction{
	Mnemonic:     "RRC C",
	Encoding:     "0xCB09",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_d = Instruction{
	Mnemonic:     "RRC D",
	Encoding:     "0xCB0A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_e = Instruction{
	Mnemonic:     "RRC E",
	Encoding:     "0xCB0B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_h = Instruction{
	Mnemonic:     "RRC H",
	Encoding:     "0xCB0C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_l = Instruction{
	Mnemonic:     "RRC L",
	Encoding:     "0xCB0D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc__hl_ = Instruction{
	Mnemonic:     "RRC (HL)",
	Encoding:     "0xCB0E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rrc_a = Instruction{
	Mnemonic:     "RRC A",
	Encoding:     "0xCB0F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_b = Instruction{
	Mnemonic:     "RL B",
	Encoding:     "0xCB10",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_c = Instruction{
	Mnemonic:     "RL C",
	Encoding:     "0xCB11",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_d = Instruction{
	Mnemonic:     "RL D",
	Encoding:     "0xCB12",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_e = Instruction{
	Mnemonic:     "RL E",
	Encoding:     "0xCB13",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_h = Instruction{
	Mnemonic:     "RL H",
	Encoding:     "0xCB14",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_l = Instruction{
	Mnemonic:     "RL L",
	Encoding:     "0xCB15",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl__hl_ = Instruction{
	Mnemonic:     "RL (HL)",
	Encoding:     "0xCB16",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rl_a = Instruction{
	Mnemonic:     "RL A",
	Encoding:     "0xCB17",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_b = Instruction{
	Mnemonic:     "RR B",
	Encoding:     "0xCB18",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_c = Instruction{
	Mnemonic:     "RR C",
	Encoding:     "0xCB19",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_d = Instruction{
	Mnemonic:     "RR D",
	Encoding:     "0xCB1A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_e = Instruction{
	Mnemonic:     "RR E",
	Encoding:     "0xCB1B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_h = Instruction{
	Mnemonic:     "RR H",
	Encoding:     "0xCB1C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_l = Instruction{
	Mnemonic:     "RR L",
	Encoding:     "0xCB1D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr__hl_ = Instruction{
	Mnemonic:     "RR (HL)",
	Encoding:     "0xCB1E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var rr_a = Instruction{
	Mnemonic:     "RR A",
	Encoding:     "0xCB1F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_b = Instruction{
	Mnemonic:     "SLA B",
	Encoding:     "0xCB20",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_c = Instruction{
	Mnemonic:     "SLA C",
	Encoding:     "0xCB21",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_d = Instruction{
	Mnemonic:     "SLA D",
	Encoding:     "0xCB22",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_e = Instruction{
	Mnemonic:     "SLA E",
	Encoding:     "0xCB23",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_h = Instruction{
	Mnemonic:     "SLA H",
	Encoding:     "0xCB24",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_l = Instruction{
	Mnemonic:     "SLA L",
	Encoding:     "0xCB25",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla__hl_ = Instruction{
	Mnemonic:     "SLA (HL)",
	Encoding:     "0xCB26",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sla_a = Instruction{
	Mnemonic:     "SLA A",
	Encoding:     "0xCB27",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A7",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_b = Instruction{
	Mnemonic:     "SRA B",
	Encoding:     "0xCB28",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_c = Instruction{
	Mnemonic:     "SRA C",
	Encoding:     "0xCB29",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_d = Instruction{
	Mnemonic:     "SRA D",
	Encoding:     "0xCB2A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_e = Instruction{
	Mnemonic:     "SRA E",
	Encoding:     "0xCB2B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_h = Instruction{
	Mnemonic:     "SRA H",
	Encoding:     "0xCB2C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_l = Instruction{
	Mnemonic:     "SRA L",
	Encoding:     "0xCB2D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra__hl_ = Instruction{
	Mnemonic:     "SRA (HL)",
	Encoding:     "0xCB2E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var sra_a = Instruction{
	Mnemonic:     "SRA A",
	Encoding:     "0xCB2F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_b = Instruction{
	Mnemonic:     "SWAP B",
	Encoding:     "0xCB30",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_c = Instruction{
	Mnemonic:     "SWAP C",
	Encoding:     "0xCB31",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_d = Instruction{
	Mnemonic:     "SWAP D",
	Encoding:     "0xCB32",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_e = Instruction{
	Mnemonic:     "SWAP E",
	Encoding:     "0xCB33",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_h = Instruction{
	Mnemonic:     "SWAP H",
	Encoding:     "0xCB34",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_l = Instruction{
	Mnemonic:     "SWAP L",
	Encoding:     "0xCB35",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap__hl_ = Instruction{
	Mnemonic:     "SWAP (HL)",
	Encoding:     "0xCB36",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var swap_a = Instruction{
	Mnemonic:     "SWAP A",
	Encoding:     "0xCB37",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_b = Instruction{
	Mnemonic:     "SRL B",
	Encoding:     "0xCB38",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 B0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_c = Instruction{
	Mnemonic:     "SRL C",
	Encoding:     "0xCB39",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 C0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_d = Instruction{
	Mnemonic:     "SRL D",
	Encoding:     "0xCB3A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 D0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_e = Instruction{
	Mnemonic:     "SRL E",
	Encoding:     "0xCB3B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 E0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_h = Instruction{
	Mnemonic:     "SRL H",
	Encoding:     "0xCB3C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 H0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_l = Instruction{
	Mnemonic:     "SRL L",
	Encoding:     "0xCB3D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 L0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl__hl_ = Instruction{
	Mnemonic:     "SRL (HL)",
	Encoding:     "0xCB3E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "Z 0 0 (HL)0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var srl_a = Instruction{
	Mnemonic:     "SRL A",
	Encoding:     "0xCB3F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "Z 0 0 A0",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__b = Instruction{
	Mnemonic:     "BIT 0, B",
	Encoding:     "0xCB40",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__c = Instruction{
	Mnemonic:     "BIT 0, C",
	Encoding:     "0xCB41",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__d = Instruction{
	Mnemonic:     "BIT 0, D",
	Encoding:     "0xCB42",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__e = Instruction{
	Mnemonic:     "BIT 0, E",
	Encoding:     "0xCB43",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__h = Instruction{
	Mnemonic:     "BIT 0, H",
	Encoding:     "0xCB44",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__l = Instruction{
	Mnemonic:     "BIT 0, L",
	Encoding:     "0xCB45",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0___hl_ = Instruction{
	Mnemonic:     "BIT 0, (HL)",
	Encoding:     "0xCB46",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_0__a = Instruction{
	Mnemonic:     "BIT 0, A",
	Encoding:     "0xCB47",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r0 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__b = Instruction{
	Mnemonic:     "BIT 1, B",
	Encoding:     "0xCB48",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__c = Instruction{
	Mnemonic:     "BIT 1, C",
	Encoding:     "0xCB49",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__d = Instruction{
	Mnemonic:     "BIT 1, D",
	Encoding:     "0xCB4A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__e = Instruction{
	Mnemonic:     "BIT 1, E",
	Encoding:     "0xCB4B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__h = Instruction{
	Mnemonic:     "BIT 1, H",
	Encoding:     "0xCB4C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__l = Instruction{
	Mnemonic:     "BIT 1, L",
	Encoding:     "0xCB4D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1___hl_ = Instruction{
	Mnemonic:     "BIT 1, (HL)",
	Encoding:     "0xCB4E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_1__a = Instruction{
	Mnemonic:     "BIT 1, A",
	Encoding:     "0xCB4F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r1 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__b = Instruction{
	Mnemonic:     "BIT 2, B",
	Encoding:     "0xCB50",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__c = Instruction{
	Mnemonic:     "BIT 2, C",
	Encoding:     "0xCB51",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__d = Instruction{
	Mnemonic:     "BIT 2, D",
	Encoding:     "0xCB52",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__e = Instruction{
	Mnemonic:     "BIT 2, E",
	Encoding:     "0xCB53",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__h = Instruction{
	Mnemonic:     "BIT 2, H",
	Encoding:     "0xCB54",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__l = Instruction{
	Mnemonic:     "BIT 2, L",
	Encoding:     "0xCB55",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2___hl_ = Instruction{
	Mnemonic:     "BIT 2, (HL)",
	Encoding:     "0xCB56",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_2__a = Instruction{
	Mnemonic:     "BIT 2, A",
	Encoding:     "0xCB57",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r2 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__b = Instruction{
	Mnemonic:     "BIT 3, B",
	Encoding:     "0xCB58",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__c = Instruction{
	Mnemonic:     "BIT 3, C",
	Encoding:     "0xCB59",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__d = Instruction{
	Mnemonic:     "BIT 3, D",
	Encoding:     "0xCB5A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__e = Instruction{
	Mnemonic:     "BIT 3, E",
	Encoding:     "0xCB5B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__h = Instruction{
	Mnemonic:     "BIT 3, H",
	Encoding:     "0xCB5C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__l = Instruction{
	Mnemonic:     "BIT 3, L",
	Encoding:     "0xCB5D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3___hl_ = Instruction{
	Mnemonic:     "BIT 3, (HL)",
	Encoding:     "0xCB5E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_3__a = Instruction{
	Mnemonic:     "BIT 3, A",
	Encoding:     "0xCB5F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r3 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__b = Instruction{
	Mnemonic:     "BIT 4, B",
	Encoding:     "0xCB60",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__c = Instruction{
	Mnemonic:     "BIT 4, C",
	Encoding:     "0xCB61",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__d = Instruction{
	Mnemonic:     "BIT 4, D",
	Encoding:     "0xCB62",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__e = Instruction{
	Mnemonic:     "BIT 4, E",
	Encoding:     "0xCB63",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__h = Instruction{
	Mnemonic:     "BIT 4, H",
	Encoding:     "0xCB64",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__l = Instruction{
	Mnemonic:     "BIT 4, L",
	Encoding:     "0xCB65",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4___hl_ = Instruction{
	Mnemonic:     "BIT 4, (HL)",
	Encoding:     "0xCB66",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_4__a = Instruction{
	Mnemonic:     "BIT 4, A",
	Encoding:     "0xCB67",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r4 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__b = Instruction{
	Mnemonic:     "BIT 5, B",
	Encoding:     "0xCB68",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__c = Instruction{
	Mnemonic:     "BIT 5, C",
	Encoding:     "0xCB69",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__d = Instruction{
	Mnemonic:     "BIT 5, D",
	Encoding:     "0xCB6A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__e = Instruction{
	Mnemonic:     "BIT 5, E",
	Encoding:     "0xCB6B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__h = Instruction{
	Mnemonic:     "BIT 5, H",
	Encoding:     "0xCB6C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__l = Instruction{
	Mnemonic:     "BIT 5, L",
	Encoding:     "0xCB6D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5___hl_ = Instruction{
	Mnemonic:     "BIT 5, (HL)",
	Encoding:     "0xCB6E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_5__a = Instruction{
	Mnemonic:     "BIT 5, A",
	Encoding:     "0xCB6F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r5 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__b = Instruction{
	Mnemonic:     "BIT 6, B",
	Encoding:     "0xCB70",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__c = Instruction{
	Mnemonic:     "BIT 6, C",
	Encoding:     "0xCB71",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__d = Instruction{
	Mnemonic:     "BIT 6, D",
	Encoding:     "0xCB72",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__e = Instruction{
	Mnemonic:     "BIT 6, E",
	Encoding:     "0xCB73",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__h = Instruction{
	Mnemonic:     "BIT 6, H",
	Encoding:     "0xCB74",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__l = Instruction{
	Mnemonic:     "BIT 6, L",
	Encoding:     "0xCB75",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6___hl_ = Instruction{
	Mnemonic:     "BIT 6, (HL)",
	Encoding:     "0xCB76",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_6__a = Instruction{
	Mnemonic:     "BIT 6, A",
	Encoding:     "0xCB77",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r6 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__b = Instruction{
	Mnemonic:     "BIT 7, B",
	Encoding:     "0xCB78",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__c = Instruction{
	Mnemonic:     "BIT 7, C",
	Encoding:     "0xCB79",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__d = Instruction{
	Mnemonic:     "BIT 7, D",
	Encoding:     "0xCB7A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__e = Instruction{
	Mnemonic:     "BIT 7, E",
	Encoding:     "0xCB7B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__h = Instruction{
	Mnemonic:     "BIT 7, H",
	Encoding:     "0xCB7C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__l = Instruction{
	Mnemonic:     "BIT 7, L",
	Encoding:     "0xCB7D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7___hl_ = Instruction{
	Mnemonic:     "BIT 7, (HL)",
	Encoding:     "0xCB7E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	flags:        "!(HL)7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var bit_7__a = Instruction{
	Mnemonic:     "BIT 7, A",
	Encoding:     "0xCB7F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "!r7 0 1 -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__b = Instruction{
	Mnemonic:     "RES 0, B",
	Encoding:     "0xCB80",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__c = Instruction{
	Mnemonic:     "RES 0, C",
	Encoding:     "0xCB81",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__d = Instruction{
	Mnemonic:     "RES 0, D",
	Encoding:     "0xCB82",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__e = Instruction{
	Mnemonic:     "RES 0, E",
	Encoding:     "0xCB83",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__h = Instruction{
	Mnemonic:     "RES 0, H",
	Encoding:     "0xCB84",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__l = Instruction{
	Mnemonic:     "RES 0, L",
	Encoding:     "0xCB85",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0___hl_ = Instruction{
	Mnemonic:     "RES 0, (HL)",
	Encoding:     "0xCB86",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_0__a = Instruction{
	Mnemonic:     "RES 0, A",
	Encoding:     "0xCB87",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__b = Instruction{
	Mnemonic:     "RES 1, B",
	Encoding:     "0xCB88",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__c = Instruction{
	Mnemonic:     "RES 1, C",
	Encoding:     "0xCB89",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__d = Instruction{
	Mnemonic:     "RES 1, D",
	Encoding:     "0xCB8A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__e = Instruction{
	Mnemonic:     "RES 1, E",
	Encoding:     "0xCB8B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__h = Instruction{
	Mnemonic:     "RES 1, H",
	Encoding:     "0xCB8C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__l = Instruction{
	Mnemonic:     "RES 1, L",
	Encoding:     "0xCB8D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1___hl_ = Instruction{
	Mnemonic:     "RES 1, (HL)",
	Encoding:     "0xCB8E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_1__a = Instruction{
	Mnemonic:     "RES 1, A",
	Encoding:     "0xCB8F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__b = Instruction{
	Mnemonic:     "RES 2, B",
	Encoding:     "0xCB90",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__c = Instruction{
	Mnemonic:     "RES 2, C",
	Encoding:     "0xCB91",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__d = Instruction{
	Mnemonic:     "RES 2, D",
	Encoding:     "0xCB92",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__e = Instruction{
	Mnemonic:     "RES 2, E",
	Encoding:     "0xCB93",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__h = Instruction{
	Mnemonic:     "RES 2, H",
	Encoding:     "0xCB94",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__l = Instruction{
	Mnemonic:     "RES 2, L",
	Encoding:     "0xCB95",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2___hl_ = Instruction{
	Mnemonic:     "RES 2, (HL)",
	Encoding:     "0xCB96",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_2__a = Instruction{
	Mnemonic:     "RES 2, A",
	Encoding:     "0xCB97",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__b = Instruction{
	Mnemonic:     "RES 3, B",
	Encoding:     "0xCB98",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__c = Instruction{
	Mnemonic:     "RES 3, C",
	Encoding:     "0xCB99",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__d = Instruction{
	Mnemonic:     "RES 3, D",
	Encoding:     "0xCB9A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__e = Instruction{
	Mnemonic:     "RES 3, E",
	Encoding:     "0xCB9B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__h = Instruction{
	Mnemonic:     "RES 3, H",
	Encoding:     "0xCB9C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__l = Instruction{
	Mnemonic:     "RES 3, L",
	Encoding:     "0xCB9D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3___hl_ = Instruction{
	Mnemonic:     "RES 3, (HL)",
	Encoding:     "0xCB9E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_3__a = Instruction{
	Mnemonic:     "RES 3, A",
	Encoding:     "0xCB9F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__b = Instruction{
	Mnemonic:     "RES 4, B",
	Encoding:     "0xCBA0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__c = Instruction{
	Mnemonic:     "RES 4, C",
	Encoding:     "0xCBA1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__d = Instruction{
	Mnemonic:     "RES 4, D",
	Encoding:     "0xCBA2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__e = Instruction{
	Mnemonic:     "RES 4, E",
	Encoding:     "0xCBA3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__h = Instruction{
	Mnemonic:     "RES 4, H",
	Encoding:     "0xCBA4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__l = Instruction{
	Mnemonic:     "RES 4, L",
	Encoding:     "0xCBA5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4___hl_ = Instruction{
	Mnemonic:     "RES 4, (HL)",
	Encoding:     "0xCBA6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_4__a = Instruction{
	Mnemonic:     "RES 4, A",
	Encoding:     "0xCBA7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__b = Instruction{
	Mnemonic:     "RES 5, B",
	Encoding:     "0xCBA8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__c = Instruction{
	Mnemonic:     "RES 5, C",
	Encoding:     "0xCBA9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__d = Instruction{
	Mnemonic:     "RES 5, D",
	Encoding:     "0xCBAA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__e = Instruction{
	Mnemonic:     "RES 5, E",
	Encoding:     "0xCBAB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__h = Instruction{
	Mnemonic:     "RES 5, H",
	Encoding:     "0xCBAC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__l = Instruction{
	Mnemonic:     "RES 5, L",
	Encoding:     "0xCBAD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5___hl_ = Instruction{
	Mnemonic:     "RES 5, (HL)",
	Encoding:     "0xCBAE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_5__a = Instruction{
	Mnemonic:     "RES 5, A",
	Encoding:     "0xCBAF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__b = Instruction{
	Mnemonic:     "RES 6, B",
	Encoding:     "0xCBB0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__c = Instruction{
	Mnemonic:     "RES 6, C",
	Encoding:     "0xCBB1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__d = Instruction{
	Mnemonic:     "RES 6, D",
	Encoding:     "0xCBB2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__e = Instruction{
	Mnemonic:     "RES 6, E",
	Encoding:     "0xCBB3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__h = Instruction{
	Mnemonic:     "RES 6, H",
	Encoding:     "0xCBB4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__l = Instruction{
	Mnemonic:     "RES 6, L",
	Encoding:     "0xCBB5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6___hl_ = Instruction{
	Mnemonic:     "RES 6, (HL)",
	Encoding:     "0xCBB6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_6__a = Instruction{
	Mnemonic:     "RES 6, A",
	Encoding:     "0xCBB7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__b = Instruction{
	Mnemonic:     "RES 7, B",
	Encoding:     "0xCBB8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__c = Instruction{
	Mnemonic:     "RES 7, C",
	Encoding:     "0xCBB9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__d = Instruction{
	Mnemonic:     "RES 7, D",
	Encoding:     "0xCBBA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__e = Instruction{
	Mnemonic:     "RES 7, E",
	Encoding:     "0xCBBB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__h = Instruction{
	Mnemonic:     "RES 7, H",
	Encoding:     "0xCBBC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__l = Instruction{
	Mnemonic:     "RES 7, L",
	Encoding:     "0xCBBD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7___hl_ = Instruction{
	Mnemonic:     "RES 7, (HL)",
	Encoding:     "0xCBBE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var res_7__a = Instruction{
	Mnemonic:     "RES 7, A",
	Encoding:     "0xCBBF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__b = Instruction{
	Mnemonic:     "SET 0, B",
	Encoding:     "0xCBC0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__c = Instruction{
	Mnemonic:     "SET 0, C",
	Encoding:     "0xCBC1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__d = Instruction{
	Mnemonic:     "SET 0, D",
	Encoding:     "0xCBC2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__e = Instruction{
	Mnemonic:     "SET 0, E",
	Encoding:     "0xCBC3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__h = Instruction{
	Mnemonic:     "SET 0, H",
	Encoding:     "0xCBC4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__l = Instruction{
	Mnemonic:     "SET 0, L",
	Encoding:     "0xCBC5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0___hl_ = Instruction{
	Mnemonic:     "SET 0, (HL)",
	Encoding:     "0xCBC6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_0__a = Instruction{
	Mnemonic:     "SET 0, A",
	Encoding:     "0xCBC7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__b = Instruction{
	Mnemonic:     "SET 1, B",
	Encoding:     "0xCBC8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__c = Instruction{
	Mnemonic:     "SET 1, C",
	Encoding:     "0xCBC9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__d = Instruction{
	Mnemonic:     "SET 1, D",
	Encoding:     "0xCBCA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__e = Instruction{
	Mnemonic:     "SET 1, E",
	Encoding:     "0xCBCB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__h = Instruction{
	Mnemonic:     "SET 1, H",
	Encoding:     "0xCBCC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__l = Instruction{
	Mnemonic:     "SET 1, L",
	Encoding:     "0xCBCD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1___hl_ = Instruction{
	Mnemonic:     "SET 1, (HL)",
	Encoding:     "0xCBCE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_1__a = Instruction{
	Mnemonic:     "SET 1, A",
	Encoding:     "0xCBCF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__b = Instruction{
	Mnemonic:     "SET 2, B",
	Encoding:     "0xCBD0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__c = Instruction{
	Mnemonic:     "SET 2, C",
	Encoding:     "0xCBD1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__d = Instruction{
	Mnemonic:     "SET 2, D",
	Encoding:     "0xCBD2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__e = Instruction{
	Mnemonic:     "SET 2, E",
	Encoding:     "0xCBD3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__h = Instruction{
	Mnemonic:     "SET 2, H",
	Encoding:     "0xCBD4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__l = Instruction{
	Mnemonic:     "SET 2, L",
	Encoding:     "0xCBD5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2___hl_ = Instruction{
	Mnemonic:     "SET 2, (HL)",
	Encoding:     "0xCBD6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_2__a = Instruction{
	Mnemonic:     "SET 2, A",
	Encoding:     "0xCBD7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__b = Instruction{
	Mnemonic:     "SET 3, B",
	Encoding:     "0xCBD8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__c = Instruction{
	Mnemonic:     "SET 3, C",
	Encoding:     "0xCBD9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__d = Instruction{
	Mnemonic:     "SET 3, D",
	Encoding:     "0xCBDA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__e = Instruction{
	Mnemonic:     "SET 3, E",
	Encoding:     "0xCBDB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__h = Instruction{
	Mnemonic:     "SET 3, H",
	Encoding:     "0xCBDC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__l = Instruction{
	Mnemonic:     "SET 3, L",
	Encoding:     "0xCBDD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3___hl_ = Instruction{
	Mnemonic:     "SET 3, (HL)",
	Encoding:     "0xCBDE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_3__a = Instruction{
	Mnemonic:     "SET 3, A",
	Encoding:     "0xCBDF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__b = Instruction{
	Mnemonic:     "SET 4, B",
	Encoding:     "0xCBE0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__c = Instruction{
	Mnemonic:     "SET 4, C",
	Encoding:     "0xCBE1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__d = Instruction{
	Mnemonic:     "SET 4, D",
	Encoding:     "0xCBE2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__e = Instruction{
	Mnemonic:     "SET 4, E",
	Encoding:     "0xCBE3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__h = Instruction{
	Mnemonic:     "SET 4, H",
	Encoding:     "0xCBE4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__l = Instruction{
	Mnemonic:     "SET 4, L",
	Encoding:     "0xCBE5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4___hl_ = Instruction{
	Mnemonic:     "SET 4, (HL)",
	Encoding:     "0xCBE6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_4__a = Instruction{
	Mnemonic:     "SET 4, A",
	Encoding:     "0xCBE7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__b = Instruction{
	Mnemonic:     "SET 5, B",
	Encoding:     "0xCBE8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__c = Instruction{
	Mnemonic:     "SET 5, C",
	Encoding:     "0xCBE9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__d = Instruction{
	Mnemonic:     "SET 5, D",
	Encoding:     "0xCBEA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__e = Instruction{
	Mnemonic:     "SET 5, E",
	Encoding:     "0xCBEB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__h = Instruction{
	Mnemonic:     "SET 5, H",
	Encoding:     "0xCBEC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__l = Instruction{
	Mnemonic:     "SET 5, L",
	Encoding:     "0xCBED",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5___hl_ = Instruction{
	Mnemonic:     "SET 5, (HL)",
	Encoding:     "0xCBEE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_5__a = Instruction{
	Mnemonic:     "SET 5, A",
	Encoding:     "0xCBEF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__b = Instruction{
	Mnemonic:     "SET 6, B",
	Encoding:     "0xCBF0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__c = Instruction{
	Mnemonic:     "SET 6, C",
	Encoding:     "0xCBF1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__d = Instruction{
	Mnemonic:     "SET 6, D",
	Encoding:     "0xCBF2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__e = Instruction{
	Mnemonic:     "SET 6, E",
	Encoding:     "0xCBF3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__h = Instruction{
	Mnemonic:     "SET 6, H",
	Encoding:     "0xCBF4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__l = Instruction{
	Mnemonic:     "SET 6, L",
	Encoding:     "0xCBF5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6___hl_ = Instruction{
	Mnemonic:     "SET 6, (HL)",
	Encoding:     "0xCBF6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_6__a = Instruction{
	Mnemonic:     "SET 6, A",
	Encoding:     "0xCBF7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__b = Instruction{
	Mnemonic:     "SET 7, B",
	Encoding:     "0xCBF8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__c = Instruction{
	Mnemonic:     "SET 7, C",
	Encoding:     "0xCBF9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__d = Instruction{
	Mnemonic:     "SET 7, D",
	Encoding:     "0xCBFA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__e = Instruction{
	Mnemonic:     "SET 7, E",
	Encoding:     "0xCBFB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__h = Instruction{
	Mnemonic:     "SET 7, H",
	Encoding:     "0xCBFC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__l = Instruction{
	Mnemonic:     "SET 7, L",
	Encoding:     "0xCBFD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7___hl_ = Instruction{
	Mnemonic:     "SET 7, (HL)",
	Encoding:     "0xCBFE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var set_7__a = Instruction{
	Mnemonic:     "SET 7, A",
	Encoding:     "0xCBFF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	flags:        "- - - -",
	Implemented:  false,
	execute:      func(cpu *cpu) {},
}

var InstructionTable8 = map[uint8]Instruction{
	0x00: nop,
	0x01: ld_bc__d16,
	0x02: ld__bc___a,
	0x03: inc_bc,
	0x04: inc_b,
	0x05: dec_b,
	0x06: ld_b__d8,
	0x07: rlca,
	0x08: ld__a16___sp,
	0x09: add_hl__bc,
	0x0A: ld_a___bc_,
	0x0B: dec_bc,
	0x0C: inc_c,
	0x0D: dec_c,
	0x0E: ld_c__d8,
	0x0F: rrca,
	0x11: ld_de__d16,
	0x12: ld__de___a,
	0x13: inc_de,
	0x14: inc_d,
	0x15: dec_d,
	0x16: ld_d__d8,
	0x17: rla,
	0x18: jr_s8,
	0x19: add_hl__de,
	0x1A: ld_a___de_,
	0x1B: dec_de,
	0x1C: inc_e,
	0x1D: dec_e,
	0x1E: ld_e__d8,
	0x1F: rra,
	0x20: jr_nz__s8,
	0x21: ld_hl__d16,
	0x22: ld__hlp___a,
	0x23: inc_hl,
	0x24: inc_h,
	0x25: dec_h,
	0x26: ld_h__d8,
	0x27: daa,
	0x28: jr_z__s8,
	0x29: add_hl__hl,
	0x2A: ld_a___hlp_,
	0x2B: dec_hl,
	0x2C: inc_l,
	0x2D: dec_l,
	0x2E: ld_l__d8,
	0x2F: cpl,
	0x30: jr_nc__s8,
	0x31: ld_sp__d16,
	0x32: ld__hlm___a,
	0x33: inc_sp,
	0x34: inc__hl_,
	0x35: dec__hl_,
	0x36: ld__hl___d8,
	0x37: scf,
	0x38: jr_c__s8,
	0x39: add_hl__sp,
	0x3A: ld_a___hlm_,
	0x3B: dec_sp,
	0x3C: inc_a,
	0x3D: dec_a,
	0x3E: ld_a__d8,
	0x3F: ccf,
	0x40: ld_b__b,
	0x41: ld_b__c,
	0x42: ld_b__d,
	0x43: ld_b__e,
	0x44: ld_b__h,
	0x45: ld_b__l,
	0x46: ld_b___hl_,
	0x47: ld_b__a,
	0x48: ld_c__b,
	0x49: ld_c__c,
	0x4A: ld_c__d,
	0x4B: ld_c__e,
	0x4C: ld_c__h,
	0x4D: ld_c__l,
	0x4E: ld_c___hl_,
	0x4F: ld_c__a,
	0x50: ld_d__b,
	0x51: ld_d__c,
	0x52: ld_d__d,
	0x53: ld_d__e,
	0x54: ld_d__h,
	0x55: ld_d__l,
	0x56: ld_d___hl_,
	0x57: ld_d__a,
	0x58: ld_e__b,
	0x59: ld_e__c,
	0x5A: ld_e__d,
	0x5B: ld_e__e,
	0x5C: ld_e__h,
	0x5D: ld_e__l,
	0x5E: ld_e___hl_,
	0x5F: ld_e__a,
	0x60: ld_h__b,
	0x61: ld_h__c,
	0x62: ld_h__d,
	0x63: ld_h__e,
	0x64: ld_h__h,
	0x65: ld_h__l,
	0x66: ld_h___hl_,
	0x67: ld_h__a,
	0x68: ld_l__b,
	0x69: ld_l__c,
	0x6A: ld_l__d,
	0x6B: ld_l__e,
	0x6C: ld_l__h,
	0x6D: ld_l__l,
	0x6E: ld_l___hl_,
	0x6F: ld_l__a,
	0x70: ld__hl___b,
	0x71: ld__hl___c,
	0x72: ld__hl___d,
	0x73: ld__hl___e,
	0x74: ld__hl___h,
	0x75: ld__hl___l,
	0x76: halt,
	0x77: ld__hl___a,
	0x78: ld_a__b,
	0x79: ld_a__c,
	0x7A: ld_a__d,
	0x7B: ld_a__e,
	0x7C: ld_a__h,
	0x7D: ld_a__l,
	0x7E: ld_a___hl_,
	0x7F: ld_a__a,
	0x80: add_a__b,
	0x81: add_a__c,
	0x82: add_a__d,
	0x83: add_a__e,
	0x84: add_a__h,
	0x85: add_a__l,
	0x86: add_a___hl_,
	0x87: add_a__a,
	0x88: adc_a__b,
	0x89: adc_a__c,
	0x8A: adc_a__d,
	0x8B: adc_a__e,
	0x8C: adc_a__h,
	0x8D: adc_a__l,
	0x8E: adc_a___hl_,
	0x8F: adc_a__a,
	0x90: sub_b,
	0x91: sub_c,
	0x92: sub_d,
	0x93: sub_e,
	0x94: sub_h,
	0x95: sub_l,
	0x96: sub__hl_,
	0x97: sub_a,
	0x98: sbc_a__b,
	0x99: sbc_a__c,
	0x9A: sbc_a__d,
	0x9B: sbc_a__e,
	0x9C: sbc_a__h,
	0x9D: sbc_a__l,
	0x9E: sbc_a___hl_,
	0x9F: sbc_a__a,
	0xA0: and_b,
	0xA1: and_c,
	0xA2: and_d,
	0xA3: and_e,
	0xA4: and_h,
	0xA5: and_l,
	0xA6: and__hl_,
	0xA7: and_a,
	0xA8: xor_b,
	0xA9: xor_c,
	0xAA: xor_d,
	0xAB: xor_e,
	0xAC: xor_h,
	0xAD: xor_l,
	0xAE: xor__hl_,
	0xAF: xor_a,
	0xB0: or_b,
	0xB1: or_c,
	0xB2: or_d,
	0xB3: or_e,
	0xB4: or_h,
	0xB5: or_l,
	0xB6: or__hl_,
	0xB7: or_a,
	0xB8: cp_b,
	0xB9: cp_c,
	0xBA: cp_d,
	0xBB: cp_e,
	0xBC: cp_h,
	0xBD: cp_l,
	0xBE: cp__hl_,
	0xBF: cp_a,
	0xC0: ret_nz,
	0xC1: pop_bc,
	0xC2: jp_nz__a16,
	0xC3: jp_a16,
	0xC4: call_nz__a16,
	0xC5: push_bc,
	0xC6: add_a__d8,
	0xC7: rst_0,
	0xC8: ret_z,
	0xC9: ret,
	0xCA: jp_z__a16,
	0xCC: call_z__a16,
	0xCD: call_a16,
	0xCE: adc_a__d8,
	0xCF: rst_1,
	0xD0: ret_nc,
	0xD1: pop_de,
	0xD2: jp_nc__a16,
	0xD4: call_nc__a16,
	0xD5: push_de,
	0xD6: sub_d8,
	0xD7: rst_2,
	0xD8: ret_c,
	0xD9: reti,
	0xDA: jp_c__a16,
	0xDC: call_c__a16,
	0xDE: sbc_a__d8,
	0xDF: rst_3,
	0xE0: ld__a8___a,
	0xE1: pop_hl,
	0xE2: ld__c___a,
	0xE5: push_hl,
	0xE6: and_d8,
	0xE7: rst_4,
	0xE8: add_sp__s8,
	0xE9: jp_hl,
	0xEA: ld__a16___a,
	0xEE: xor_d8,
	0xEF: rst_5,
	0xF0: ld_a___a8_,
	0xF1: pop_af,
	0xF2: ld_a___c_,
	0xF3: di,
	0xF5: push_af,
	0xF6: or_d8,
	0xF7: rst_6,
	0xF8: ld_hl__sp_s8,
	0xF9: ld_sp__hl,
	0xFA: ld_a___a16_,
	0xFB: ei,
	0xFE: cp_d8,
	0xFF: rst_7,
}

var InstructionTable16 = map[uint16]Instruction{
	0x1000: stop,
	0xCB00: rlc_b,
	0xCB01: rlc_c,
	0xCB02: rlc_d,
	0xCB03: rlc_e,
	0xCB04: rlc_h,
	0xCB05: rlc_l,
	0xCB06: rlc__hl_,
	0xCB07: rlc_a,
	0xCB08: rrc_b,
	0xCB09: rrc_c,
	0xCB0A: rrc_d,
	0xCB0B: rrc_e,
	0xCB0C: rrc_h,
	0xCB0D: rrc_l,
	0xCB0E: rrc__hl_,
	0xCB0F: rrc_a,
	0xCB10: rl_b,
	0xCB11: rl_c,
	0xCB12: rl_d,
	0xCB13: rl_e,
	0xCB14: rl_h,
	0xCB15: rl_l,
	0xCB16: rl__hl_,
	0xCB17: rl_a,
	0xCB18: rr_b,
	0xCB19: rr_c,
	0xCB1A: rr_d,
	0xCB1B: rr_e,
	0xCB1C: rr_h,
	0xCB1D: rr_l,
	0xCB1E: rr__hl_,
	0xCB1F: rr_a,
	0xCB20: sla_b,
	0xCB21: sla_c,
	0xCB22: sla_d,
	0xCB23: sla_e,
	0xCB24: sla_h,
	0xCB25: sla_l,
	0xCB26: sla__hl_,
	0xCB27: sla_a,
	0xCB28: sra_b,
	0xCB29: sra_c,
	0xCB2A: sra_d,
	0xCB2B: sra_e,
	0xCB2C: sra_h,
	0xCB2D: sra_l,
	0xCB2E: sra__hl_,
	0xCB2F: sra_a,
	0xCB30: swap_b,
	0xCB31: swap_c,
	0xCB32: swap_d,
	0xCB33: swap_e,
	0xCB34: swap_h,
	0xCB35: swap_l,
	0xCB36: swap__hl_,
	0xCB37: swap_a,
	0xCB38: srl_b,
	0xCB39: srl_c,
	0xCB3A: srl_d,
	0xCB3B: srl_e,
	0xCB3C: srl_h,
	0xCB3D: srl_l,
	0xCB3E: srl__hl_,
	0xCB3F: srl_a,
	0xCB40: bit_0__b,
	0xCB41: bit_0__c,
	0xCB42: bit_0__d,
	0xCB43: bit_0__e,
	0xCB44: bit_0__h,
	0xCB45: bit_0__l,
	0xCB46: bit_0___hl_,
	0xCB47: bit_0__a,
	0xCB48: bit_1__b,
	0xCB49: bit_1__c,
	0xCB4A: bit_1__d,
	0xCB4B: bit_1__e,
	0xCB4C: bit_1__h,
	0xCB4D: bit_1__l,
	0xCB4E: bit_1___hl_,
	0xCB4F: bit_1__a,
	0xCB50: bit_2__b,
	0xCB51: bit_2__c,
	0xCB52: bit_2__d,
	0xCB53: bit_2__e,
	0xCB54: bit_2__h,
	0xCB55: bit_2__l,
	0xCB56: bit_2___hl_,
	0xCB57: bit_2__a,
	0xCB58: bit_3__b,
	0xCB59: bit_3__c,
	0xCB5A: bit_3__d,
	0xCB5B: bit_3__e,
	0xCB5C: bit_3__h,
	0xCB5D: bit_3__l,
	0xCB5E: bit_3___hl_,
	0xCB5F: bit_3__a,
	0xCB60: bit_4__b,
	0xCB61: bit_4__c,
	0xCB62: bit_4__d,
	0xCB63: bit_4__e,
	0xCB64: bit_4__h,
	0xCB65: bit_4__l,
	0xCB66: bit_4___hl_,
	0xCB67: bit_4__a,
	0xCB68: bit_5__b,
	0xCB69: bit_5__c,
	0xCB6A: bit_5__d,
	0xCB6B: bit_5__e,
	0xCB6C: bit_5__h,
	0xCB6D: bit_5__l,
	0xCB6E: bit_5___hl_,
	0xCB6F: bit_5__a,
	0xCB70: bit_6__b,
	0xCB71: bit_6__c,
	0xCB72: bit_6__d,
	0xCB73: bit_6__e,
	0xCB74: bit_6__h,
	0xCB75: bit_6__l,
	0xCB76: bit_6___hl_,
	0xCB77: bit_6__a,
	0xCB78: bit_7__b,
	0xCB79: bit_7__c,
	0xCB7A: bit_7__d,
	0xCB7B: bit_7__e,
	0xCB7C: bit_7__h,
	0xCB7D: bit_7__l,
	0xCB7E: bit_7___hl_,
	0xCB7F: bit_7__a,
	0xCB80: res_0__b,
	0xCB81: res_0__c,
	0xCB82: res_0__d,
	0xCB83: res_0__e,
	0xCB84: res_0__h,
	0xCB85: res_0__l,
	0xCB86: res_0___hl_,
	0xCB87: res_0__a,
	0xCB88: res_1__b,
	0xCB89: res_1__c,
	0xCB8A: res_1__d,
	0xCB8B: res_1__e,
	0xCB8C: res_1__h,
	0xCB8D: res_1__l,
	0xCB8E: res_1___hl_,
	0xCB8F: res_1__a,
	0xCB90: res_2__b,
	0xCB91: res_2__c,
	0xCB92: res_2__d,
	0xCB93: res_2__e,
	0xCB94: res_2__h,
	0xCB95: res_2__l,
	0xCB96: res_2___hl_,
	0xCB97: res_2__a,
	0xCB98: res_3__b,
	0xCB99: res_3__c,
	0xCB9A: res_3__d,
	0xCB9B: res_3__e,
	0xCB9C: res_3__h,
	0xCB9D: res_3__l,
	0xCB9E: res_3___hl_,
	0xCB9F: res_3__a,
	0xCBA0: res_4__b,
	0xCBA1: res_4__c,
	0xCBA2: res_4__d,
	0xCBA3: res_4__e,
	0xCBA4: res_4__h,
	0xCBA5: res_4__l,
	0xCBA6: res_4___hl_,
	0xCBA7: res_4__a,
	0xCBA8: res_5__b,
	0xCBA9: res_5__c,
	0xCBAA: res_5__d,
	0xCBAB: res_5__e,
	0xCBAC: res_5__h,
	0xCBAD: res_5__l,
	0xCBAE: res_5___hl_,
	0xCBAF: res_5__a,
	0xCBB0: res_6__b,
	0xCBB1: res_6__c,
	0xCBB2: res_6__d,
	0xCBB3: res_6__e,
	0xCBB4: res_6__h,
	0xCBB5: res_6__l,
	0xCBB6: res_6___hl_,
	0xCBB7: res_6__a,
	0xCBB8: res_7__b,
	0xCBB9: res_7__c,
	0xCBBA: res_7__d,
	0xCBBB: res_7__e,
	0xCBBC: res_7__h,
	0xCBBD: res_7__l,
	0xCBBE: res_7___hl_,
	0xCBBF: res_7__a,
	0xCBC0: set_0__b,
	0xCBC1: set_0__c,
	0xCBC2: set_0__d,
	0xCBC3: set_0__e,
	0xCBC4: set_0__h,
	0xCBC5: set_0__l,
	0xCBC6: set_0___hl_,
	0xCBC7: set_0__a,
	0xCBC8: set_1__b,
	0xCBC9: set_1__c,
	0xCBCA: set_1__d,
	0xCBCB: set_1__e,
	0xCBCC: set_1__h,
	0xCBCD: set_1__l,
	0xCBCE: set_1___hl_,
	0xCBCF: set_1__a,
	0xCBD0: set_2__b,
	0xCBD1: set_2__c,
	0xCBD2: set_2__d,
	0xCBD3: set_2__e,
	0xCBD4: set_2__h,
	0xCBD5: set_2__l,
	0xCBD6: set_2___hl_,
	0xCBD7: set_2__a,
	0xCBD8: set_3__b,
	0xCBD9: set_3__c,
	0xCBDA: set_3__d,
	0xCBDB: set_3__e,
	0xCBDC: set_3__h,
	0xCBDD: set_3__l,
	0xCBDE: set_3___hl_,
	0xCBDF: set_3__a,
	0xCBE0: set_4__b,
	0xCBE1: set_4__c,
	0xCBE2: set_4__d,
	0xCBE3: set_4__e,
	0xCBE4: set_4__h,
	0xCBE5: set_4__l,
	0xCBE6: set_4___hl_,
	0xCBE7: set_4__a,
	0xCBE8: set_5__b,
	0xCBE9: set_5__c,
	0xCBEA: set_5__d,
	0xCBEB: set_5__e,
	0xCBEC: set_5__h,
	0xCBED: set_5__l,
	0xCBEE: set_5___hl_,
	0xCBEF: set_5__a,
	0xCBF0: set_6__b,
	0xCBF1: set_6__c,
	0xCBF2: set_6__d,
	0xCBF3: set_6__e,
	0xCBF4: set_6__h,
	0xCBF5: set_6__l,
	0xCBF6: set_6___hl_,
	0xCBF7: set_6__a,
	0xCBF8: set_7__b,
	0xCBF9: set_7__c,
	0xCBFA: set_7__d,
	0xCBFB: set_7__e,
	0xCBFC: set_7__h,
	0xCBFD: set_7__l,
	0xCBFE: set_7___hl_,
	0xCBFF: set_7__a,
}
