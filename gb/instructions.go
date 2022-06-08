package gb

// TODO: maybe trim this down, binary is pretty big holding
// all of this memory
type instruction struct {
	mnemonic                 string
	encoding                 string
	size                     uint16
	jumpCycles, noJumpCycles int
	z, n, h, c               string
	execute                  func(*cpu)
}

func (i instruction) String() string {
	return i.mnemonic
}

// 8-bit instructions
var nop = instruction{
	mnemonic:     "NOP",
	encoding:     "0x00",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__bc___a = instruction{
	mnemonic:     "LD (BC), A",
	encoding:     "0x02",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.bc(), cpu.a)
	},
}

var inc_bc = instruction{
	mnemonic:     "INC BC",
	encoding:     "0x03",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.setBc(cpu.bc() + 1)
	},
}

var ld_bc__d16 = instruction{
	mnemonic:     "LD BC, d16",
	encoding:     "0x01",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_b = instruction{
	mnemonic:     "INC B",
	encoding:     "0x04",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b += 1
	},
}

var ld_b__d8 = instruction{
	mnemonic:     "LD B, d8",
	encoding:     "0x06",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.readByte(cpu.pc + 1)
	},
}

var dec_b = instruction{
	mnemonic:     "DEC B",
	encoding:     "0x05",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var rlca = instruction{
	mnemonic:     "RLCA",
	encoding:     "0x07",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "0",
	n:            "0",
	h:            "0",
	c:            "A7",
	execute:      func(cpu *cpu) {},
}

var ld__a16___sp = instruction{
	mnemonic:     "LD (a16), SP",
	encoding:     "0x08",
	size:         3,
	jumpCycles:   20,
	noJumpCycles: 20,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___bc_ = instruction{
	mnemonic:     "LD A, (BC)",
	encoding:     "0x0A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.bc())
	},
}

var add_hl__bc = instruction{
	mnemonic:     "ADD HL, BC",
	encoding:     "0x09",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var inc_c = instruction{
	mnemonic:     "INC C",
	encoding:     "0x0C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var dec_bc = instruction{
	mnemonic:     "DEC BC",
	encoding:     "0x0B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var dec_c = instruction{
	mnemonic:     "DEC C",
	encoding:     "0x0D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var stop = instruction{
	mnemonic:     "STOP",
	encoding:     "0x1000",
	size:         2,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_c__d8 = instruction{
	mnemonic:     "LD C, d8",
	encoding:     "0x0E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.readByte(cpu.pc + 1)
	},
}

var ld__de___a = instruction{
	mnemonic:     "LD (DE), A",
	encoding:     "0x12",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.de(), cpu.a)
	},
}

var rrca = instruction{
	mnemonic:     "RRCA",
	encoding:     "0x0F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "0",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var ld_de__d16 = instruction{
	mnemonic:     "LD DE, d16",
	encoding:     "0x11",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_d = instruction{
	mnemonic:     "INC D",
	encoding:     "0x14",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_de = instruction{
	mnemonic:     "INC DE",
	encoding:     "0x13",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.setDe(cpu.de() + 1)
	},
}

var rla = instruction{
	mnemonic:     "RLA",
	encoding:     "0x17",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "0",
	n:            "0",
	h:            "0",
	c:            "A7",
	execute:      func(cpu *cpu) {},
}

var dec_d = instruction{
	mnemonic:     "DEC D",
	encoding:     "0x15",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var add_hl__de = instruction{
	mnemonic:     "ADD HL, DE",
	encoding:     "0x19",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ld_d__d8 = instruction{
	mnemonic:     "LD D, d8",
	encoding:     "0x16",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.readByte(cpu.pc + 1)
	},
}

var jr_s8 = instruction{
	mnemonic:     "JR s8",
	encoding:     "0x18",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var dec_de = instruction{
	mnemonic:     "DEC DE",
	encoding:     "0x1B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___de_ = instruction{
	mnemonic:     "LD A, (DE)",
	encoding:     "0x1A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.de())
	},
}

var ld_e__d8 = instruction{
	mnemonic:     "LD E, d8",
	encoding:     "0x1E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.readByte(cpu.pc + 1)
	},
}

var dec_e = instruction{
	mnemonic:     "DEC E",
	encoding:     "0x1D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var jr_nz__s8 = instruction{
	mnemonic:     "JR NZ, s8",
	encoding:     "0x20",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_e = instruction{
	mnemonic:     "INC E",
	encoding:     "0x1C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var rra = instruction{
	mnemonic:     "RRA",
	encoding:     "0x1F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "0",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var ld__hlp___a = instruction{
	mnemonic:     "LD (HL+), A",
	encoding:     "0x22",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
		cpu.setHl(cpu.hl() + 1)
	},
}

var inc_hl = instruction{
	mnemonic:     "INC HL",
	encoding:     "0x23",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.setHl(cpu.hl() + 1)
	},
}

var dec_h = instruction{
	mnemonic:     "DEC H",
	encoding:     "0x25",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_hl__d16 = instruction{
	mnemonic:     "LD HL, d16",
	encoding:     "0x21",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_h = instruction{
	mnemonic:     "INC H",
	encoding:     "0x24",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var daa = instruction{
	mnemonic:     "DAA",
	encoding:     "0x27",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "-",
	h:            "0",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var jr_z__s8 = instruction{
	mnemonic:     "JR Z, s8",
	encoding:     "0x28",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_h__d8 = instruction{
	mnemonic:     "LD H, d8",
	encoding:     "0x26",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.readByte(cpu.pc + 1)
	},
}

var add_hl__hl = instruction{
	mnemonic:     "ADD HL, HL",
	encoding:     "0x29",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var dec_l = instruction{
	mnemonic:     "DEC L",
	encoding:     "0x2D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_l = instruction{
	mnemonic:     "INC L",
	encoding:     "0x2C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_l__d8 = instruction{
	mnemonic:     "LD L, d8",
	encoding:     "0x2E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.readByte(cpu.pc + 1)
	},
}

var jr_nc__s8 = instruction{
	mnemonic:     "JR NC, s8",
	encoding:     "0x30",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var cpl = instruction{
	mnemonic:     "CPL",
	encoding:     "0x2F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "1",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__hlm____a = instruction{
	mnemonic:     "LD (HL-), A",
	encoding:     "0x32",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
		cpu.setHl(cpu.hl() - 1)
	},
}

var inc_sp = instruction{
	mnemonic:     "INC SP",
	encoding:     "0x33",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.sp += 1
	},
}

var ld_sp__d16 = instruction{
	mnemonic:     "LD SP, d16",
	encoding:     "0x31",
	size:         3,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var dec_hl = instruction{
	mnemonic:     "DEC HL",
	encoding:     "0x2B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__hl___d8 = instruction{
	mnemonic:     "LD (HL), d8",
	encoding:     "0x36",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.readByte(cpu.pc+1))
	},
}

var dec__hl_ = instruction{
	mnemonic:     "DEC (HL)",
	encoding:     "0x35",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___hlp_ = instruction{
	mnemonic:     "LD A, (HL+)",
	encoding:     "0x2A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
		cpu.setHl(cpu.hl() + 1)
	},
}

var inc__hl_ = instruction{
	mnemonic:     "INC (HL)",
	encoding:     "0x34",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var dec_sp = instruction{
	mnemonic:     "DEC SP",
	encoding:     "0x3B",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var inc_a = instruction{
	mnemonic:     "INC A",
	encoding:     "0x3C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___hlm_ = instruction{
	mnemonic:     "LD A, (HL-)",
	encoding:     "0x3A",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
		cpu.setHl(cpu.hl() - 1)
	},
}

var ld_a__d8 = instruction{
	mnemonic:     "LD A, d8",
	encoding:     "0x3E",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.pc + 1)
	},
}

var ccf = instruction{
	mnemonic:     "CCF",
	encoding:     "0x3F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "0",
	h:            "0",
	c:            "!CY",
	execute:      func(cpu *cpu) {},
}

var dec_a = instruction{
	mnemonic:     "DEC A",
	encoding:     "0x3D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_b__c = instruction{
	mnemonic:     "LD B, C",
	encoding:     "0x41",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.c
	},
}

var ld_b__d = instruction{
	mnemonic:     "LD B, D",
	encoding:     "0x42",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.d
	},
}

var scf = instruction{
	mnemonic:     "SCF",
	encoding:     "0x37",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "0",
	h:            "0",
	c:            "1",
	execute:      func(cpu *cpu) {},
}

var ld_b__h = instruction{
	mnemonic:     "LD B, H",
	encoding:     "0x44",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.h
	},
}

var ld_b__l = instruction{
	mnemonic:     "LD B, L",
	encoding:     "0x45",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.l
	},
}

var ld_b__b = instruction{
	mnemonic:     "LD B, B",
	encoding:     "0x40",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.b
	},
}

var ld_b__a = instruction{
	mnemonic:     "LD B, A",
	encoding:     "0x47",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.a
	},
}

var add_hl__sp = instruction{
	mnemonic:     "ADD HL, SP",
	encoding:     "0x39",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ld_b___hl_ = instruction{
	mnemonic:     "LD B, (HL)",
	encoding:     "0x46",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.readByte(cpu.hl())
	},
}

var jr_c__s8 = instruction{
	mnemonic:     "JR C, s8",
	encoding:     "0x38",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_c__c = instruction{
	mnemonic:     "LD C, C",
	encoding:     "0x49",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.c
	},
}

var ld_c__l = instruction{
	mnemonic:     "LD C, L",
	encoding:     "0x4D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.l
	},
}

var ld_c__e = instruction{
	mnemonic:     "LD C, E",
	encoding:     "0x4B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.e
	},
}

var ld_c__b = instruction{
	mnemonic:     "LD C, B",
	encoding:     "0x48",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.b
	},
}

var ld_c__a = instruction{
	mnemonic:     "LD C, A",
	encoding:     "0x4F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.a
	},
}

var ld_c___hl_ = instruction{
	mnemonic:     "LD C, (HL)",
	encoding:     "0x4E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.readByte(cpu.hl())
	},
}

var ld_d__d = instruction{
	mnemonic:     "LD D, D",
	encoding:     "0x52",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.d
	},
}

var ld_d__c = instruction{
	mnemonic:     "LD D, C",
	encoding:     "0x51",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.c
	},
}

var ld_d__e = instruction{
	mnemonic:     "LD D, E",
	encoding:     "0x53",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.e
	},
}

var ld_c__h = instruction{
	mnemonic:     "LD C, H",
	encoding:     "0x4C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.h
	},
}

var ld_d__h = instruction{
	mnemonic:     "LD D, H",
	encoding:     "0x54",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.h
	},
}

var ld_d__l = instruction{
	mnemonic:     "LD D, L",
	encoding:     "0x55",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.l
	},
}

var ld_d__a = instruction{
	mnemonic:     "LD D, A",
	encoding:     "0x57",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.a
	},
}

var ld_e__b = instruction{
	mnemonic:     "LD E, B",
	encoding:     "0x58",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.b
	},
}

var ld_d___hl_ = instruction{
	mnemonic:     "LD D, (HL)",
	encoding:     "0x56",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.readByte(cpu.hl())
	},
}

var ld_e__d = instruction{
	mnemonic:     "LD E, D",
	encoding:     "0x5A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.d
	},
}

var ld_e__e = instruction{
	mnemonic:     "LD E, E",
	encoding:     "0x5B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.e
	},
}

var ld_c__d = instruction{
	mnemonic:     "LD C, D",
	encoding:     "0x4A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.c = cpu.d
	},
}

var ld_e__l = instruction{
	mnemonic:     "LD E, L",
	encoding:     "0x5D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.l
	},
}

var ld_b__e = instruction{
	mnemonic:     "LD B, E",
	encoding:     "0x43",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.b = cpu.e
	},
}

var ld_h__b = instruction{
	mnemonic:     "LD H, B",
	encoding:     "0x60",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.b
	},
}

var ld_e___hl_ = instruction{
	mnemonic:     "LD E, (HL)",
	encoding:     "0x5E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.readByte(cpu.hl())
	},
}

var ld_h__d = instruction{
	mnemonic:     "LD H, D",
	encoding:     "0x62",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.d
	},
}

var ld_e__a = instruction{
	mnemonic:     "LD E, A",
	encoding:     "0x5F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.a
	},
}

var ld_h__c = instruction{
	mnemonic:     "LD H, C",
	encoding:     "0x61",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.c
	},
}

var ld_h__h = instruction{
	mnemonic:     "LD H, H",
	encoding:     "0x64",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.h
	},
}

var ld_h__e = instruction{
	mnemonic:     "LD H, E",
	encoding:     "0x63",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.e
	},
}

var ld_h__a = instruction{
	mnemonic:     "LD H, A",
	encoding:     "0x67",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.a
	},
}

var ld_h__l = instruction{
	mnemonic:     "LD H, L",
	encoding:     "0x65",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.l
	},
}

var ld_h___hl_ = instruction{
	mnemonic:     "LD H, (HL)",
	encoding:     "0x66",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.h = cpu.readByte(cpu.hl())
	},
}

var ld_l__c = instruction{
	mnemonic:     "LD L, C",
	encoding:     "0x69",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.c
	},
}

var ld_l__d = instruction{
	mnemonic:     "LD L, D",
	encoding:     "0x6A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.d
	},
}

var ld_l__b = instruction{
	mnemonic:     "LD L, B",
	encoding:     "0x68",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.b
	},
}

var ld_l__h = instruction{
	mnemonic:     "LD L, H",
	encoding:     "0x6C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.h
	},
}

var ld_l__l = instruction{
	mnemonic:     "LD L, L",
	encoding:     "0x6D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.l
	},
}

var ld_e__c = instruction{
	mnemonic:     "LD E, C",
	encoding:     "0x59",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.c
	},
}

var ld_l__a = instruction{
	mnemonic:     "LD L, A",
	encoding:     "0x6F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.a
	},
}

var ld_l___hl_ = instruction{
	mnemonic:     "LD L, (HL)",
	encoding:     "0x6E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.readByte(cpu.hl())
	},
}

var ld__hl___b = instruction{
	mnemonic:     "LD (HL), B",
	encoding:     "0x70",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.b)
	},
}

var ld_d__b = instruction{
	mnemonic:     "LD D, B",
	encoding:     "0x50",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.d = cpu.b
	},
}

var ld__hl___c = instruction{
	mnemonic:     "LD (HL), C",
	encoding:     "0x71",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.c)
	},
}

var ld__hl___h = instruction{
	mnemonic:     "LD (HL), H",
	encoding:     "0x74",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.h)
	},
}

var ld__hl___l = instruction{
	mnemonic:     "LD (HL), L",
	encoding:     "0x75",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.l)
	},
}

var ld__hl___e = instruction{
	mnemonic:     "LD (HL), E",
	encoding:     "0x73",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.e)
	},
}

var ld_l__e = instruction{
	mnemonic:     "LD L, E",
	encoding:     "0x6B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.l = cpu.e
	},
}

var ld_a__b = instruction{
	mnemonic:     "LD A, B",
	encoding:     "0x78",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.b
	},
}

var ld_e__h = instruction{
	mnemonic:     "LD E, H",
	encoding:     "0x5C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.e = cpu.h
	},
}

var ld_a__d = instruction{
	mnemonic:     "LD A, D",
	encoding:     "0x7A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.d
	},
}

var ld_a__e = instruction{
	mnemonic:     "LD A, E",
	encoding:     "0x7B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.e
	},
}

var ld_a__c = instruction{
	mnemonic:     "LD A, C",
	encoding:     "0x79",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.c
	},
}

var ld_a__l = instruction{
	mnemonic:     "LD A, L",
	encoding:     "0x7D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.l
	},
}

var ld_a___hl_ = instruction{
	mnemonic:     "LD A, (HL)",
	encoding:     "0x7E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.readByte(cpu.hl())
	},
}

var ld_a__h = instruction{
	mnemonic:     "LD A, H",
	encoding:     "0x7C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.h
	},
}

var add_a__b = instruction{
	mnemonic:     "ADD A, B",
	encoding:     "0x80",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a__c = instruction{
	mnemonic:     "ADD A, C",
	encoding:     "0x81",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ld_a__a = instruction{
	mnemonic:     "LD A, A",
	encoding:     "0x7F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.a = cpu.a
	},
}

var add_a__e = instruction{
	mnemonic:     "ADD A, E",
	encoding:     "0x83",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a__h = instruction{
	mnemonic:     "ADD A, H",
	encoding:     "0x84",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a__d = instruction{
	mnemonic:     "ADD A, D",
	encoding:     "0x82",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a___hl_ = instruction{
	mnemonic:     "ADD A, (HL)",
	encoding:     "0x86",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a__a = instruction{
	mnemonic:     "ADD A, A",
	encoding:     "0x87",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var add_a__l = instruction{
	mnemonic:     "ADD A, L",
	encoding:     "0x85",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a__c = instruction{
	mnemonic:     "ADC A, C",
	encoding:     "0x89",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a__b = instruction{
	mnemonic:     "ADC A, B",
	encoding:     "0x88",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var halt = instruction{
	mnemonic:     "HALT",
	encoding:     "0x76",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var adc_a__e = instruction{
	mnemonic:     "ADC A, E",
	encoding:     "0x8B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a__d = instruction{
	mnemonic:     "ADC A, D",
	encoding:     "0x8A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a___hl_ = instruction{
	mnemonic:     "ADC A, (HL)",
	encoding:     "0x8E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a__l = instruction{
	mnemonic:     "ADC A, L",
	encoding:     "0x8D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_b = instruction{
	mnemonic:     "SUB B",
	encoding:     "0x90",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_c = instruction{
	mnemonic:     "SUB C",
	encoding:     "0x91",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var adc_a__a = instruction{
	mnemonic:     "ADC A, A",
	encoding:     "0x8F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_e = instruction{
	mnemonic:     "SUB E",
	encoding:     "0x93",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_h = instruction{
	mnemonic:     "SUB H",
	encoding:     "0x94",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ld__hl___a = instruction{
	mnemonic:     "LD (HL), A",
	encoding:     "0x77",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.a)
	},
}

var sub__hl_ = instruction{
	mnemonic:     "SUB (HL)",
	encoding:     "0x96",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_a = instruction{
	mnemonic:     "SUB A",
	encoding:     "0x97",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_l = instruction{
	mnemonic:     "SUB L",
	encoding:     "0x95",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__c = instruction{
	mnemonic:     "SBC A, C",
	encoding:     "0x99",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__d = instruction{
	mnemonic:     "SBC A, D",
	encoding:     "0x9A",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__b = instruction{
	mnemonic:     "SBC A, B",
	encoding:     "0x98",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__h = instruction{
	mnemonic:     "SBC A, H",
	encoding:     "0x9C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__l = instruction{
	mnemonic:     "SBC A, L",
	encoding:     "0x9D",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__e = instruction{
	mnemonic:     "SBC A, E",
	encoding:     "0x9B",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sbc_a__a = instruction{
	mnemonic:     "SBC A, A",
	encoding:     "0x9F",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var and_b = instruction{
	mnemonic:     "AND B",
	encoding:     "0xA0",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var sbc_a___hl_ = instruction{
	mnemonic:     "SBC A, (HL)",
	encoding:     "0x9E",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var and_d = instruction{
	mnemonic:     "AND D",
	encoding:     "0xA2",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var and_e = instruction{
	mnemonic:     "AND E",
	encoding:     "0xA3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var and_c = instruction{
	mnemonic:     "AND C",
	encoding:     "0xA1",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var and_l = instruction{
	mnemonic:     "AND L",
	encoding:     "0xA5",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var and__hl_ = instruction{
	mnemonic:     "AND (HL)",
	encoding:     "0xA6",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var and_h = instruction{
	mnemonic:     "AND H",
	encoding:     "0xA4",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var xor_b = instruction{
	mnemonic:     "XOR B",
	encoding:     "0xA8",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.b
	},
}

var and_a = instruction{
	mnemonic:     "AND A",
	encoding:     "0xA7",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var xor_e = instruction{
	mnemonic:     "XOR E",
	encoding:     "0xAB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.e
	},
}

var xor_c = instruction{
	mnemonic:     "XOR C",
	encoding:     "0xA9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.c
	},
}

var xor_d = instruction{
	mnemonic:     "XOR D",
	encoding:     "0xAA",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.d
	},
}

var xor_l = instruction{
	mnemonic:     "XOR L",
	encoding:     "0xAD",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.l
	},
}

var xor__hl_ = instruction{
	mnemonic:     "XOR (HL)",
	encoding:     "0xAE",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a = byte(uint16(cpu.a) ^ cpu.hl())
	},
}

var xor_h = instruction{
	mnemonic:     "XOR H",
	encoding:     "0xAC",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.h
	},
}

var or_b = instruction{
	mnemonic:     "OR B",
	encoding:     "0xB0",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var xor_a = instruction{
	mnemonic:     "XOR A",
	encoding:     "0xAF",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.a
	},
}

var or_c = instruction{
	mnemonic:     "OR C",
	encoding:     "0xB1",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var or_d = instruction{
	mnemonic:     "OR D",
	encoding:     "0xB2",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var or_l = instruction{
	mnemonic:     "OR L",
	encoding:     "0xB5",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var adc_a__h = instruction{
	mnemonic:     "ADC A, H",
	encoding:     "0x8C",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var or_h = instruction{
	mnemonic:     "OR H",
	encoding:     "0xB4",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var or_a = instruction{
	mnemonic:     "OR A",
	encoding:     "0xB7",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var cp_b = instruction{
	mnemonic:     "CP B",
	encoding:     "0xB8",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var or__hl_ = instruction{
	mnemonic:     "OR (HL)",
	encoding:     "0xB6",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var cp_d = instruction{
	mnemonic:     "CP D",
	encoding:     "0xBA",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var cp_e = instruction{
	mnemonic:     "CP E",
	encoding:     "0xBB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var cp_c = instruction{
	mnemonic:     "CP C",
	encoding:     "0xB9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var cp_l = instruction{
	mnemonic:     "CP L",
	encoding:     "0xBD",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var cp_h = instruction{
	mnemonic:     "CP H",
	encoding:     "0xBC",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var cp__hl_ = instruction{
	mnemonic:     "CP (HL)",
	encoding:     "0xBE",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ret_nz = instruction{
	mnemonic:     "RET NZ",
	encoding:     "0xC0",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var cp_a = instruction{
	mnemonic:     "CP A",
	encoding:     "0xBF",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var jp_nz__a16 = instruction{
	mnemonic:     "JP NZ, a16",
	encoding:     "0xC2",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var jp_a16 = instruction{
	mnemonic:     "JP a16",
	encoding:     "0xC3",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.pc = cpu.readWord(cpu.pc + 1)
	},
}

var pop_bc = instruction{
	mnemonic:     "POP BC",
	encoding:     "0xC1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var push_bc = instruction{
	mnemonic:     "PUSH BC",
	encoding:     "0xC5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var call_nz__a16 = instruction{
	mnemonic:     "CALL NZ, a16",
	encoding:     "0xC4",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var add_a__d8 = instruction{
	mnemonic:     "ADD A, d8",
	encoding:     "0xC6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ret_z = instruction{
	mnemonic:     "RET Z",
	encoding:     "0xC8",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var rst_0 = instruction{
	mnemonic:     "RST 0",
	encoding:     "0xC7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var jp_z__a16 = instruction{
	mnemonic:     "JP Z, a16",
	encoding:     "0xCA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var call_z__a16 = instruction{
	mnemonic:     "CALL Z, a16",
	encoding:     "0xCC",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ret = instruction{
	mnemonic:     "RET",
	encoding:     "0xC9",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var adc_a__d8 = instruction{
	mnemonic:     "ADC A, d8",
	encoding:     "0xCE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var rst_1 = instruction{
	mnemonic:     "RST 1",
	encoding:     "0xCF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__hl___d = instruction{
	mnemonic:     "LD (HL), D",
	encoding:     "0x72",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute: func(cpu *cpu) {
		cpu.writeByte(cpu.hl(), cpu.d)
	},
}

var pop_de = instruction{
	mnemonic:     "POP DE",
	encoding:     "0xD1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ret_nc = instruction{
	mnemonic:     "RET NC",
	encoding:     "0xD0",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var push_de = instruction{
	mnemonic:     "PUSH DE",
	encoding:     "0xD5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var jp_nc__a16 = instruction{
	mnemonic:     "JP NC, a16",
	encoding:     "0xD2",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var rst_2 = instruction{
	mnemonic:     "RST 2",
	encoding:     "0xD7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var call_nc__a16 = instruction{
	mnemonic:     "CALL NC, a16",
	encoding:     "0xD4",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var sub_d8 = instruction{
	mnemonic:     "SUB d8",
	encoding:     "0xD6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var sub_d = instruction{
	mnemonic:     "SUB D",
	encoding:     "0x92",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ret_c = instruction{
	mnemonic:     "RET C",
	encoding:     "0xD8",
	size:         1,
	jumpCycles:   20,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var jp_c__a16 = instruction{
	mnemonic:     "JP C, a16",
	encoding:     "0xDA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var call_c__a16 = instruction{
	mnemonic:     "CALL C, a16",
	encoding:     "0xDC",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var sbc_a__d8 = instruction{
	mnemonic:     "SBC A, d8",
	encoding:     "0xDE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var rst_3 = instruction{
	mnemonic:     "RST 3",
	encoding:     "0xDF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__a8___a = instruction{
	mnemonic:     "LD (a8), A",
	encoding:     "0xE0",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var pop_hl = instruction{
	mnemonic:     "POP HL",
	encoding:     "0xE1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__c___a = instruction{
	mnemonic:     "LD (C), A",
	encoding:     "0xE2",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var reti = instruction{
	mnemonic:     "RETI",
	encoding:     "0xD9",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var and_d8 = instruction{
	mnemonic:     "AND d8",
	encoding:     "0xE6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "1",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var rst_4 = instruction{
	mnemonic:     "RST 4",
	encoding:     "0xE7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var add_sp__s8 = instruction{
	mnemonic:     "ADD SP, s8",
	encoding:     "0xE8",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "0",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var jp_hl = instruction{
	mnemonic:     "JP HL",
	encoding:     "0xE9",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld__a16___a = instruction{
	mnemonic:     "LD (a16), A",
	encoding:     "0xEA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var xor_d8 = instruction{
	mnemonic:     "XOR d8",
	encoding:     "0xEE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute: func(cpu *cpu) {
		cpu.a ^= cpu.memory.readByte(cpu.pc + 1)
	},
}

var rst_5 = instruction{
	mnemonic:     "RST 5",
	encoding:     "0xEF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___a8_ = instruction{
	mnemonic:     "LD A, (a8)",
	encoding:     "0xF0",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var pop_af = instruction{
	mnemonic:     "POP AF",
	encoding:     "0xF1",
	size:         1,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___c_ = instruction{
	mnemonic:     "LD A, (C)",
	encoding:     "0xF2",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var di = instruction{
	mnemonic:     "DI",
	encoding:     "0xF3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var push_hl = instruction{
	mnemonic:     "PUSH HL",
	encoding:     "0xE5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var or_d8 = instruction{
	mnemonic:     "OR d8",
	encoding:     "0xF6",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var rst_6 = instruction{
	mnemonic:     "RST 6",
	encoding:     "0xF7",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_sp__hl = instruction{
	mnemonic:     "LD SP, HL",
	encoding:     "0xF9",
	size:         1,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_hl__sp_s8 = instruction{
	mnemonic:     "LD HL, SP+s8",
	encoding:     "0xF8",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "0",
	n:            "0",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var ei = instruction{
	mnemonic:     "EI",
	encoding:     "0xFB",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var ld_a___a16_ = instruction{
	mnemonic:     "LD A, (a16)",
	encoding:     "0xFA",
	size:         3,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var cp_d8 = instruction{
	mnemonic:     "CP d8",
	encoding:     "0xFE",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "1",
	h:            "H",
	c:            "CY",
	execute:      func(cpu *cpu) {},
}

var rst_7 = instruction{
	mnemonic:     "RST 7",
	encoding:     "0xFF",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

// 16-bit instructions
var rlc_b = instruction{
	mnemonic:     "RLC B",
	encoding:     "0xCB00",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B7",
	execute:      func(cpu *cpu) {},
}

var rlc_c = instruction{
	mnemonic:     "RLC C",
	encoding:     "0xCB01",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C7",
	execute:      func(cpu *cpu) {},
}

var or_e = instruction{
	mnemonic:     "OR E",
	encoding:     "0xB3",
	size:         1,
	jumpCycles:   4,
	noJumpCycles: 4,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var push_af = instruction{
	mnemonic:     "PUSH AF",
	encoding:     "0xF5",
	size:         1,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var call_a16 = instruction{
	mnemonic:     "CALL a16",
	encoding:     "0xCD",
	size:         3,
	jumpCycles:   24,
	noJumpCycles: 24,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var rlc__hl_ = instruction{
	mnemonic:     "RLC (HL)",
	encoding:     "0xCB06",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)7",
	execute:      func(cpu *cpu) {},
}

var rlc_l = instruction{
	mnemonic:     "RLC L",
	encoding:     "0xCB05",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L7",
	execute:      func(cpu *cpu) {},
}

var rlc_a = instruction{
	mnemonic:     "RLC A",
	encoding:     "0xCB07",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A7",
	execute:      func(cpu *cpu) {},
}

var rrc_b = instruction{
	mnemonic:     "RRC B",
	encoding:     "0xCB08",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B0",
	execute:      func(cpu *cpu) {},
}

var rrc_c = instruction{
	mnemonic:     "RRC C",
	encoding:     "0xCB09",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C0",
	execute:      func(cpu *cpu) {},
}

var rlc_e = instruction{
	mnemonic:     "RLC E",
	encoding:     "0xCB03",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E7",
	execute:      func(cpu *cpu) {},
}

var rrc_d = instruction{
	mnemonic:     "RRC D",
	encoding:     "0xCB0A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D0",
	execute:      func(cpu *cpu) {},
}

var rrc_h = instruction{
	mnemonic:     "RRC H",
	encoding:     "0xCB0C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H0",
	execute:      func(cpu *cpu) {},
}

var rrc__hl_ = instruction{
	mnemonic:     "RRC (HL)",
	encoding:     "0xCB0E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)0",
	execute:      func(cpu *cpu) {},
}

var rrc_l = instruction{
	mnemonic:     "RRC L",
	encoding:     "0xCB0D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L0",
	execute:      func(cpu *cpu) {},
}

var rl_b = instruction{
	mnemonic:     "RL B",
	encoding:     "0xCB10",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B7",
	execute:      func(cpu *cpu) {},
}

var rl_c = instruction{
	mnemonic:     "RL C",
	encoding:     "0xCB11",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C7",
	execute:      func(cpu *cpu) {},
}

var rlc_h = instruction{
	mnemonic:     "RLC H",
	encoding:     "0xCB04",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H7",
	execute:      func(cpu *cpu) {},
}

var rl_e = instruction{
	mnemonic:     "RL E",
	encoding:     "0xCB13",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E7",
	execute:      func(cpu *cpu) {},
}

var rl_h = instruction{
	mnemonic:     "RL H",
	encoding:     "0xCB14",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H7",
	execute:      func(cpu *cpu) {},
}

var rl_d = instruction{
	mnemonic:     "RL D",
	encoding:     "0xCB12",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D7",
	execute:      func(cpu *cpu) {},
}

var rl__hl_ = instruction{
	mnemonic:     "RL (HL)",
	encoding:     "0xCB16",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)7",
	execute:      func(cpu *cpu) {},
}

var rl_a = instruction{
	mnemonic:     "RL A",
	encoding:     "0xCB17",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A7",
	execute:      func(cpu *cpu) {},
}

var rl_l = instruction{
	mnemonic:     "RL L",
	encoding:     "0xCB15",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L7",
	execute:      func(cpu *cpu) {},
}

var rr_c = instruction{
	mnemonic:     "RR C",
	encoding:     "0xCB19",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C0",
	execute:      func(cpu *cpu) {},
}

var rr_d = instruction{
	mnemonic:     "RR D",
	encoding:     "0xCB1A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D0",
	execute:      func(cpu *cpu) {},
}

var rr_b = instruction{
	mnemonic:     "RR B",
	encoding:     "0xCB18",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B0",
	execute:      func(cpu *cpu) {},
}

var rr_h = instruction{
	mnemonic:     "RR H",
	encoding:     "0xCB1C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H0",
	execute:      func(cpu *cpu) {},
}

var rr_l = instruction{
	mnemonic:     "RR L",
	encoding:     "0xCB1D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L0",
	execute:      func(cpu *cpu) {},
}

var rr_e = instruction{
	mnemonic:     "RR E",
	encoding:     "0xCB1B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E0",
	execute:      func(cpu *cpu) {},
}

var rr_a = instruction{
	mnemonic:     "RR A",
	encoding:     "0xCB1F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var sla_b = instruction{
	mnemonic:     "SLA B",
	encoding:     "0xCB20",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B7",
	execute:      func(cpu *cpu) {},
}

var rr__hl_ = instruction{
	mnemonic:     "RR (HL)",
	encoding:     "0xCB1E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)0",
	execute:      func(cpu *cpu) {},
}

var sla_d = instruction{
	mnemonic:     "SLA D",
	encoding:     "0xCB22",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D7",
	execute:      func(cpu *cpu) {},
}

var sla_e = instruction{
	mnemonic:     "SLA E",
	encoding:     "0xCB23",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E7",
	execute:      func(cpu *cpu) {},
}

var rrc_a = instruction{
	mnemonic:     "RRC A",
	encoding:     "0xCB0F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var rlc_d = instruction{
	mnemonic:     "RLC D",
	encoding:     "0xCB02",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D7",
	execute:      func(cpu *cpu) {},
}

var sla_h = instruction{
	mnemonic:     "SLA H",
	encoding:     "0xCB24",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H7",
	execute:      func(cpu *cpu) {},
}

var sla__hl_ = instruction{
	mnemonic:     "SLA (HL)",
	encoding:     "0xCB26",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)7",
	execute:      func(cpu *cpu) {},
}

var sla_a = instruction{
	mnemonic:     "SLA A",
	encoding:     "0xCB27",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A7",
	execute:      func(cpu *cpu) {},
}

var sra_b = instruction{
	mnemonic:     "SRA B",
	encoding:     "0xCB28",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B0",
	execute:      func(cpu *cpu) {},
}

var sra_c = instruction{
	mnemonic:     "SRA C",
	encoding:     "0xCB29",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C0",
	execute:      func(cpu *cpu) {},
}

var sra_d = instruction{
	mnemonic:     "SRA D",
	encoding:     "0xCB2A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D0",
	execute:      func(cpu *cpu) {},
}

var sra_e = instruction{
	mnemonic:     "SRA E",
	encoding:     "0xCB2B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E0",
	execute:      func(cpu *cpu) {},
}

var sla_c = instruction{
	mnemonic:     "SLA C",
	encoding:     "0xCB21",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C7",
	execute:      func(cpu *cpu) {},
}

var sra_l = instruction{
	mnemonic:     "SRA L",
	encoding:     "0xCB2D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L0",
	execute:      func(cpu *cpu) {},
}

var sra__hl_ = instruction{
	mnemonic:     "SRA (HL)",
	encoding:     "0xCB2E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)0",
	execute:      func(cpu *cpu) {},
}

var sra_a = instruction{
	mnemonic:     "SRA A",
	encoding:     "0xCB2F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var swap_b = instruction{
	mnemonic:     "SWAP B",
	encoding:     "0xCB30",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_c = instruction{
	mnemonic:     "SWAP C",
	encoding:     "0xCB31",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_d = instruction{
	mnemonic:     "SWAP D",
	encoding:     "0xCB32",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_e = instruction{
	mnemonic:     "SWAP E",
	encoding:     "0xCB33",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_h = instruction{
	mnemonic:     "SWAP H",
	encoding:     "0xCB34",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_l = instruction{
	mnemonic:     "SWAP L",
	encoding:     "0xCB35",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap__hl_ = instruction{
	mnemonic:     "SWAP (HL)",
	encoding:     "0xCB36",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var swap_a = instruction{
	mnemonic:     "SWAP A",
	encoding:     "0xCB37",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "0",
	execute:      func(cpu *cpu) {},
}

var sra_h = instruction{
	mnemonic:     "SRA H",
	encoding:     "0xCB2C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H0",
	execute:      func(cpu *cpu) {},
}

var sla_l = instruction{
	mnemonic:     "SLA L",
	encoding:     "0xCB25",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L7",
	execute:      func(cpu *cpu) {},
}

var rrc_e = instruction{
	mnemonic:     "RRC E",
	encoding:     "0xCB0B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E0",
	execute:      func(cpu *cpu) {},
}

var srl_b = instruction{
	mnemonic:     "SRL B",
	encoding:     "0xCB38",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "B0",
	execute:      func(cpu *cpu) {},
}

var srl_h = instruction{
	mnemonic:     "SRL H",
	encoding:     "0xCB3C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "H0",
	execute:      func(cpu *cpu) {},
}

var srl_l = instruction{
	mnemonic:     "SRL L",
	encoding:     "0xCB3D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "L0",
	execute:      func(cpu *cpu) {},
}

var srl_e = instruction{
	mnemonic:     "SRL E",
	encoding:     "0xCB3B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "E0",
	execute:      func(cpu *cpu) {},
}

var srl_a = instruction{
	mnemonic:     "SRL A",
	encoding:     "0xCB3F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "A0",
	execute:      func(cpu *cpu) {},
}

var bit_0__b = instruction{
	mnemonic:     "BIT 0, B",
	encoding:     "0xCB40",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__c = instruction{
	mnemonic:     "BIT 0, C",
	encoding:     "0xCB41",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__d = instruction{
	mnemonic:     "BIT 0, D",
	encoding:     "0xCB42",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__e = instruction{
	mnemonic:     "BIT 0, E",
	encoding:     "0xCB43",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__h = instruction{
	mnemonic:     "BIT 0, H",
	encoding:     "0xCB44",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__l = instruction{
	mnemonic:     "BIT 0, L",
	encoding:     "0xCB45",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0___hl_ = instruction{
	mnemonic:     "BIT 0, (HL)",
	encoding:     "0xCB46",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_0__a = instruction{
	mnemonic:     "BIT 0, A",
	encoding:     "0xCB47",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r0",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__b = instruction{
	mnemonic:     "BIT 1, B",
	encoding:     "0xCB48",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__c = instruction{
	mnemonic:     "BIT 1, C",
	encoding:     "0xCB49",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__d = instruction{
	mnemonic:     "BIT 1, D",
	encoding:     "0xCB4A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__e = instruction{
	mnemonic:     "BIT 1, E",
	encoding:     "0xCB4B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__h = instruction{
	mnemonic:     "BIT 1, H",
	encoding:     "0xCB4C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__l = instruction{
	mnemonic:     "BIT 1, L",
	encoding:     "0xCB4D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1___hl_ = instruction{
	mnemonic:     "BIT 1, (HL)",
	encoding:     "0xCB4E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_1__a = instruction{
	mnemonic:     "BIT 1, A",
	encoding:     "0xCB4F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r1",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__b = instruction{
	mnemonic:     "BIT 2, B",
	encoding:     "0xCB50",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__c = instruction{
	mnemonic:     "BIT 2, C",
	encoding:     "0xCB51",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__d = instruction{
	mnemonic:     "BIT 2, D",
	encoding:     "0xCB52",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__e = instruction{
	mnemonic:     "BIT 2, E",
	encoding:     "0xCB53",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__h = instruction{
	mnemonic:     "BIT 2, H",
	encoding:     "0xCB54",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__l = instruction{
	mnemonic:     "BIT 2, L",
	encoding:     "0xCB55",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2___hl_ = instruction{
	mnemonic:     "BIT 2, (HL)",
	encoding:     "0xCB56",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_2__a = instruction{
	mnemonic:     "BIT 2, A",
	encoding:     "0xCB57",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r2",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var srl_d = instruction{
	mnemonic:     "SRL D",
	encoding:     "0xCB3A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "D0",
	execute:      func(cpu *cpu) {},
}

var bit_3__c = instruction{
	mnemonic:     "BIT 3, C",
	encoding:     "0xCB59",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__d = instruction{
	mnemonic:     "BIT 3, D",
	encoding:     "0xCB5A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__e = instruction{
	mnemonic:     "BIT 3, E",
	encoding:     "0xCB5B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__h = instruction{
	mnemonic:     "BIT 3, H",
	encoding:     "0xCB5C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__l = instruction{
	mnemonic:     "BIT 3, L",
	encoding:     "0xCB5D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3___hl_ = instruction{
	mnemonic:     "BIT 3, (HL)",
	encoding:     "0xCB5E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var srl__hl_ = instruction{
	mnemonic:     "SRL (HL)",
	encoding:     "0xCB3E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "(HL)0",
	execute:      func(cpu *cpu) {},
}

var bit_4__b = instruction{
	mnemonic:     "BIT 4, B",
	encoding:     "0xCB60",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4__c = instruction{
	mnemonic:     "BIT 4, C",
	encoding:     "0xCB61",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4__d = instruction{
	mnemonic:     "BIT 4, D",
	encoding:     "0xCB62",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4__e = instruction{
	mnemonic:     "BIT 4, E",
	encoding:     "0xCB63",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4__h = instruction{
	mnemonic:     "BIT 4, H",
	encoding:     "0xCB64",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4__l = instruction{
	mnemonic:     "BIT 4, L",
	encoding:     "0xCB65",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var srl_c = instruction{
	mnemonic:     "SRL C",
	encoding:     "0xCB39",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "Z",
	n:            "0",
	h:            "0",
	c:            "C0",
	execute:      func(cpu *cpu) {},
}

var bit_4__a = instruction{
	mnemonic:     "BIT 4, A",
	encoding:     "0xCB67",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__b = instruction{
	mnemonic:     "BIT 5, B",
	encoding:     "0xCB68",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__c = instruction{
	mnemonic:     "BIT 5, C",
	encoding:     "0xCB69",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__d = instruction{
	mnemonic:     "BIT 5, D",
	encoding:     "0xCB6A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__e = instruction{
	mnemonic:     "BIT 5, E",
	encoding:     "0xCB6B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__h = instruction{
	mnemonic:     "BIT 5, H",
	encoding:     "0xCB6C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__b = instruction{
	mnemonic:     "BIT 3, B",
	encoding:     "0xCB58",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_3__a = instruction{
	mnemonic:     "BIT 3, A",
	encoding:     "0xCB5F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r3",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__a = instruction{
	mnemonic:     "BIT 5, A",
	encoding:     "0xCB6F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__b = instruction{
	mnemonic:     "BIT 6, B",
	encoding:     "0xCB70",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__c = instruction{
	mnemonic:     "BIT 6, C",
	encoding:     "0xCB71",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__d = instruction{
	mnemonic:     "BIT 6, D",
	encoding:     "0xCB72",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__e = instruction{
	mnemonic:     "BIT 6, E",
	encoding:     "0xCB73",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__h = instruction{
	mnemonic:     "BIT 6, H",
	encoding:     "0xCB74",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_4___hl_ = instruction{
	mnemonic:     "BIT 4, (HL)",
	encoding:     "0xCB66",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)4",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5__l = instruction{
	mnemonic:     "BIT 5, L",
	encoding:     "0xCB6D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__a = instruction{
	mnemonic:     "BIT 6, A",
	encoding:     "0xCB77",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_5___hl_ = instruction{
	mnemonic:     "BIT 5, (HL)",
	encoding:     "0xCB6E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)5",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__c = instruction{
	mnemonic:     "BIT 7, C",
	encoding:     "0xCB79",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__d = instruction{
	mnemonic:     "BIT 7, D",
	encoding:     "0xCB7A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__e = instruction{
	mnemonic:     "BIT 7, E",
	encoding:     "0xCB7B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6___hl_ = instruction{
	mnemonic:     "BIT 6, (HL)",
	encoding:     "0xCB76",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__l = instruction{
	mnemonic:     "BIT 7, L",
	encoding:     "0xCB7D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7___hl_ = instruction{
	mnemonic:     "BIT 7, (HL)",
	encoding:     "0xCB7E",
	size:         2,
	jumpCycles:   12,
	noJumpCycles: 12,
	z:            "!(HL)7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__a = instruction{
	mnemonic:     "BIT 7, A",
	encoding:     "0xCB7F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_6__l = instruction{
	mnemonic:     "BIT 6, L",
	encoding:     "0xCB75",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r6",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__c = instruction{
	mnemonic:     "RES 0, C",
	encoding:     "0xCB81",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__d = instruction{
	mnemonic:     "RES 0, D",
	encoding:     "0xCB82",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__e = instruction{
	mnemonic:     "RES 0, E",
	encoding:     "0xCB83",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__h = instruction{
	mnemonic:     "RES 0, H",
	encoding:     "0xCB84",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__l = instruction{
	mnemonic:     "RES 0, L",
	encoding:     "0xCB85",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0___hl_ = instruction{
	mnemonic:     "RES 0, (HL)",
	encoding:     "0xCB86",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__a = instruction{
	mnemonic:     "RES 0, A",
	encoding:     "0xCB87",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__b = instruction{
	mnemonic:     "RES 1, B",
	encoding:     "0xCB88",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__c = instruction{
	mnemonic:     "RES 1, C",
	encoding:     "0xCB89",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__d = instruction{
	mnemonic:     "RES 1, D",
	encoding:     "0xCB8A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__e = instruction{
	mnemonic:     "RES 1, E",
	encoding:     "0xCB8B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__h = instruction{
	mnemonic:     "RES 1, H",
	encoding:     "0xCB8C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__l = instruction{
	mnemonic:     "RES 1, L",
	encoding:     "0xCB8D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1___hl_ = instruction{
	mnemonic:     "RES 1, (HL)",
	encoding:     "0xCB8E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_1__a = instruction{
	mnemonic:     "RES 1, A",
	encoding:     "0xCB8F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__b = instruction{
	mnemonic:     "RES 2, B",
	encoding:     "0xCB90",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__c = instruction{
	mnemonic:     "RES 2, C",
	encoding:     "0xCB91",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__d = instruction{
	mnemonic:     "RES 2, D",
	encoding:     "0xCB92",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__e = instruction{
	mnemonic:     "RES 2, E",
	encoding:     "0xCB93",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__h = instruction{
	mnemonic:     "RES 2, H",
	encoding:     "0xCB94",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__l = instruction{
	mnemonic:     "RES 2, L",
	encoding:     "0xCB95",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_0__b = instruction{
	mnemonic:     "RES 0, B",
	encoding:     "0xCB80",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__h = instruction{
	mnemonic:     "BIT 7, H",
	encoding:     "0xCB7C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var bit_7__b = instruction{
	mnemonic:     "BIT 7, B",
	encoding:     "0xCB78",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "!r7",
	n:            "0",
	h:            "1",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__c = instruction{
	mnemonic:     "RES 3, C",
	encoding:     "0xCB99",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__d = instruction{
	mnemonic:     "RES 3, D",
	encoding:     "0xCB9A",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__e = instruction{
	mnemonic:     "RES 3, E",
	encoding:     "0xCB9B",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__h = instruction{
	mnemonic:     "RES 3, H",
	encoding:     "0xCB9C",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__l = instruction{
	mnemonic:     "RES 3, L",
	encoding:     "0xCB9D",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3___hl_ = instruction{
	mnemonic:     "RES 3, (HL)",
	encoding:     "0xCB9E",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__b = instruction{
	mnemonic:     "RES 3, B",
	encoding:     "0xCB98",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2__a = instruction{
	mnemonic:     "RES 2, A",
	encoding:     "0xCB97",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__c = instruction{
	mnemonic:     "RES 4, C",
	encoding:     "0xCBA1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__d = instruction{
	mnemonic:     "RES 4, D",
	encoding:     "0xCBA2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__b = instruction{
	mnemonic:     "RES 4, B",
	encoding:     "0xCBA0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__e = instruction{
	mnemonic:     "RES 4, E",
	encoding:     "0xCBA3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__l = instruction{
	mnemonic:     "RES 4, L",
	encoding:     "0xCBA5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4___hl_ = instruction{
	mnemonic:     "RES 4, (HL)",
	encoding:     "0xCBA6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__h = instruction{
	mnemonic:     "RES 4, H",
	encoding:     "0xCBA4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__b = instruction{
	mnemonic:     "RES 5, B",
	encoding:     "0xCBA8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__c = instruction{
	mnemonic:     "RES 5, C",
	encoding:     "0xCBA9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__d = instruction{
	mnemonic:     "RES 5, D",
	encoding:     "0xCBAA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__e = instruction{
	mnemonic:     "RES 5, E",
	encoding:     "0xCBAB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__h = instruction{
	mnemonic:     "RES 5, H",
	encoding:     "0xCBAC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__l = instruction{
	mnemonic:     "RES 5, L",
	encoding:     "0xCBAD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5___hl_ = instruction{
	mnemonic:     "RES 5, (HL)",
	encoding:     "0xCBAE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_5__a = instruction{
	mnemonic:     "RES 5, A",
	encoding:     "0xCBAF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__b = instruction{
	mnemonic:     "RES 6, B",
	encoding:     "0xCBB0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__c = instruction{
	mnemonic:     "RES 6, C",
	encoding:     "0xCBB1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__d = instruction{
	mnemonic:     "RES 6, D",
	encoding:     "0xCBB2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__e = instruction{
	mnemonic:     "RES 6, E",
	encoding:     "0xCBB3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__h = instruction{
	mnemonic:     "RES 6, H",
	encoding:     "0xCBB4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_2___hl_ = instruction{
	mnemonic:     "RES 2, (HL)",
	encoding:     "0xCB96",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6___hl_ = instruction{
	mnemonic:     "RES 6, (HL)",
	encoding:     "0xCBB6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__a = instruction{
	mnemonic:     "RES 6, A",
	encoding:     "0xCBB7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__b = instruction{
	mnemonic:     "RES 7, B",
	encoding:     "0xCBB8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__c = instruction{
	mnemonic:     "RES 7, C",
	encoding:     "0xCBB9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__d = instruction{
	mnemonic:     "RES 7, D",
	encoding:     "0xCBBA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__e = instruction{
	mnemonic:     "RES 7, E",
	encoding:     "0xCBBB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__h = instruction{
	mnemonic:     "RES 7, H",
	encoding:     "0xCBBC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__l = instruction{
	mnemonic:     "RES 7, L",
	encoding:     "0xCBBD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7___hl_ = instruction{
	mnemonic:     "RES 7, (HL)",
	encoding:     "0xCBBE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_4__a = instruction{
	mnemonic:     "RES 4, A",
	encoding:     "0xCBA7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__b = instruction{
	mnemonic:     "SET 0, B",
	encoding:     "0xCBC0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__c = instruction{
	mnemonic:     "SET 0, C",
	encoding:     "0xCBC1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__d = instruction{
	mnemonic:     "SET 0, D",
	encoding:     "0xCBC2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__e = instruction{
	mnemonic:     "SET 0, E",
	encoding:     "0xCBC3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__h = instruction{
	mnemonic:     "SET 0, H",
	encoding:     "0xCBC4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__l = instruction{
	mnemonic:     "SET 0, L",
	encoding:     "0xCBC5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0___hl_ = instruction{
	mnemonic:     "SET 0, (HL)",
	encoding:     "0xCBC6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_0__a = instruction{
	mnemonic:     "SET 0, A",
	encoding:     "0xCBC7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__b = instruction{
	mnemonic:     "SET 1, B",
	encoding:     "0xCBC8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__c = instruction{
	mnemonic:     "SET 1, C",
	encoding:     "0xCBC9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_6__l = instruction{
	mnemonic:     "RES 6, L",
	encoding:     "0xCBB5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_3__a = instruction{
	mnemonic:     "RES 3, A",
	encoding:     "0xCB9F",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__h = instruction{
	mnemonic:     "SET 1, H",
	encoding:     "0xCBCC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__l = instruction{
	mnemonic:     "SET 1, L",
	encoding:     "0xCBCD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1___hl_ = instruction{
	mnemonic:     "SET 1, (HL)",
	encoding:     "0xCBCE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__a = instruction{
	mnemonic:     "SET 1, A",
	encoding:     "0xCBCF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__b = instruction{
	mnemonic:     "SET 2, B",
	encoding:     "0xCBD0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__c = instruction{
	mnemonic:     "SET 2, C",
	encoding:     "0xCBD1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__d = instruction{
	mnemonic:     "SET 2, D",
	encoding:     "0xCBD2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__d = instruction{
	mnemonic:     "SET 1, D",
	encoding:     "0xCBCA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__h = instruction{
	mnemonic:     "SET 2, H",
	encoding:     "0xCBD4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__l = instruction{
	mnemonic:     "SET 2, L",
	encoding:     "0xCBD5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2___hl_ = instruction{
	mnemonic:     "SET 2, (HL)",
	encoding:     "0xCBD6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__a = instruction{
	mnemonic:     "SET 2, A",
	encoding:     "0xCBD7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__b = instruction{
	mnemonic:     "SET 3, B",
	encoding:     "0xCBD8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__c = instruction{
	mnemonic:     "SET 3, C",
	encoding:     "0xCBD9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__d = instruction{
	mnemonic:     "SET 3, D",
	encoding:     "0xCBDA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__e = instruction{
	mnemonic:     "SET 3, E",
	encoding:     "0xCBDB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__h = instruction{
	mnemonic:     "SET 3, H",
	encoding:     "0xCBDC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_1__e = instruction{
	mnemonic:     "SET 1, E",
	encoding:     "0xCBCB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var res_7__a = instruction{
	mnemonic:     "RES 7, A",
	encoding:     "0xCBBF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_2__e = instruction{
	mnemonic:     "SET 2, E",
	encoding:     "0xCBD3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__b = instruction{
	mnemonic:     "SET 4, B",
	encoding:     "0xCBE0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__c = instruction{
	mnemonic:     "SET 4, C",
	encoding:     "0xCBE1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__d = instruction{
	mnemonic:     "SET 4, D",
	encoding:     "0xCBE2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__e = instruction{
	mnemonic:     "SET 4, E",
	encoding:     "0xCBE3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__h = instruction{
	mnemonic:     "SET 4, H",
	encoding:     "0xCBE4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__l = instruction{
	mnemonic:     "SET 4, L",
	encoding:     "0xCBE5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4___hl_ = instruction{
	mnemonic:     "SET 4, (HL)",
	encoding:     "0xCBE6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_4__a = instruction{
	mnemonic:     "SET 4, A",
	encoding:     "0xCBE7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__b = instruction{
	mnemonic:     "SET 5, B",
	encoding:     "0xCBE8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__c = instruction{
	mnemonic:     "SET 5, C",
	encoding:     "0xCBE9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__d = instruction{
	mnemonic:     "SET 5, D",
	encoding:     "0xCBEA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__e = instruction{
	mnemonic:     "SET 5, E",
	encoding:     "0xCBEB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__h = instruction{
	mnemonic:     "SET 5, H",
	encoding:     "0xCBEC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__l = instruction{
	mnemonic:     "SET 5, L",
	encoding:     "0xCBED",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__a = instruction{
	mnemonic:     "SET 3, A",
	encoding:     "0xCBDF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3__l = instruction{
	mnemonic:     "SET 3, L",
	encoding:     "0xCBDD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_3___hl_ = instruction{
	mnemonic:     "SET 3, (HL)",
	encoding:     "0xCBDE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__c = instruction{
	mnemonic:     "SET 6, C",
	encoding:     "0xCBF1",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__d = instruction{
	mnemonic:     "SET 6, D",
	encoding:     "0xCBF2",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__e = instruction{
	mnemonic:     "SET 6, E",
	encoding:     "0xCBF3",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5___hl_ = instruction{
	mnemonic:     "SET 5, (HL)",
	encoding:     "0xCBEE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__l = instruction{
	mnemonic:     "SET 6, L",
	encoding:     "0xCBF5",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6___hl_ = instruction{
	mnemonic:     "SET 6, (HL)",
	encoding:     "0xCBF6",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__h = instruction{
	mnemonic:     "SET 6, H",
	encoding:     "0xCBF4",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__b = instruction{
	mnemonic:     "SET 7, B",
	encoding:     "0xCBF8",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__b = instruction{
	mnemonic:     "SET 6, B",
	encoding:     "0xCBF0",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__d = instruction{
	mnemonic:     "SET 7, D",
	encoding:     "0xCBFA",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__e = instruction{
	mnemonic:     "SET 7, E",
	encoding:     "0xCBFB",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__h = instruction{
	mnemonic:     "SET 7, H",
	encoding:     "0xCBFC",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_6__a = instruction{
	mnemonic:     "SET 6, A",
	encoding:     "0xCBF7",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_5__a = instruction{
	mnemonic:     "SET 5, A",
	encoding:     "0xCBEF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__c = instruction{
	mnemonic:     "SET 7, C",
	encoding:     "0xCBF9",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__l = instruction{
	mnemonic:     "SET 7, L",
	encoding:     "0xCBFD",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7___hl_ = instruction{
	mnemonic:     "SET 7, (HL)",
	encoding:     "0xCBFE",
	size:         2,
	jumpCycles:   16,
	noJumpCycles: 16,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var set_7__a = instruction{
	mnemonic:     "SET 7, A",
	encoding:     "0xCBFF",
	size:         2,
	jumpCycles:   8,
	noJumpCycles: 8,
	z:            "-",
	n:            "-",
	h:            "-",
	c:            "-",
	execute:      func(cpu *cpu) {},
}

var instructionTable8 = map[uint8]instruction{
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
	0x32: ld__hlm____a,
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

var instructionTable16 = map[uint16]instruction{
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
