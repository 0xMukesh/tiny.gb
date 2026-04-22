package cpu

import (
	"github.com/0xmukesh/tiny.gb/internal/helpers"
)

type Op8Type int

const (
	Op8R8Src  Op8Type = iota // r8 index is encoded within bit 0-2
	Op8R8Dest                // r8 index is encoded within bit 3-5
	Op8HL
	Op8N8
)

type Instruction struct {
	mnemonic string
	bytes    uint
	ticks    uint
	handler  func(opcode uint8)
}

func NewInstruction(mnemonic string, bytes, ticks uint, handler func(opcode uint8)) Instruction {
	return Instruction{
		mnemonic: mnemonic,
		bytes:    bytes,
		ticks:    ticks,
		handler:  handler,
	}
}

func (c *CPU) buildInstructions() {
	c.instructions = map[uint8]Instruction{
		0x00: NewInstruction("NOP", 1, 4, c.nop),
		0x76: NewInstruction("HALT", 1, 4, c.halt),

		0x40: NewInstruction("LD B, B", 1, 4, c.nop),
		0x41: NewInstruction("LD B, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x42: NewInstruction("LD B, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x43: NewInstruction("LD B, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x44: NewInstruction("LD B, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x45: NewInstruction("LD B, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x46: NewInstruction("LD B, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x47: NewInstruction("LD B, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x06: NewInstruction("LD B, n8", 2, 8, c.ld_r8(Op8N8)),

		0x48: NewInstruction("LD C, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x49: NewInstruction("LD C, C", 1, 4, c.nop),
		0x4a: NewInstruction("LD C, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x4b: NewInstruction("LD C, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x4c: NewInstruction("LD C, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x4d: NewInstruction("LD C, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x4e: NewInstruction("LD C, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x4f: NewInstruction("LD C, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x0e: NewInstruction("LD C, n8", 2, 8, c.ld_r8(Op8N8)),

		0x50: NewInstruction("LD D, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x51: NewInstruction("LD D, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x52: NewInstruction("LD D, D", 1, 4, c.nop),
		0x53: NewInstruction("LD D, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x54: NewInstruction("LD D, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x55: NewInstruction("LD D, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x56: NewInstruction("LD D, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x57: NewInstruction("LD D, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x16: NewInstruction("LD D, n8", 2, 8, c.ld_r8(Op8N8)),

		0x58: NewInstruction("LD E, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x59: NewInstruction("LD E, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x5a: NewInstruction("LD E, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x5b: NewInstruction("LD E, E", 1, 4, c.nop),
		0x5c: NewInstruction("LD E, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x5d: NewInstruction("LD E, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x5e: NewInstruction("LD E, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x5f: NewInstruction("LD E, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x1e: NewInstruction("LD E, n8", 2, 8, c.ld_r8(Op8N8)),

		0x60: NewInstruction("LD H, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x61: NewInstruction("LD H, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x62: NewInstruction("LD H, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x63: NewInstruction("LD H, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x64: NewInstruction("LD H, H", 1, 4, c.nop),
		0x65: NewInstruction("LD H, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x66: NewInstruction("LD H, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x67: NewInstruction("LD H, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x26: NewInstruction("LD H, n8", 2, 8, c.ld_r8(Op8N8)),

		0x68: NewInstruction("LD L, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x69: NewInstruction("LD L, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x6a: NewInstruction("LD L, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x6b: NewInstruction("LD L, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x6c: NewInstruction("LD L, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x6d: NewInstruction("LD L, L", 1, 4, c.nop),
		0x6e: NewInstruction("LD L, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x6f: NewInstruction("LD L, A", 1, 4, c.ld_r8(Op8R8Src)),
		0x2e: NewInstruction("LD L, n8", 2, 8, c.ld_r8(Op8N8)),

		0x70: NewInstruction("LD [HL], B", 1, 8, c.ld_hl(Op8R8Src)),
		0x71: NewInstruction("LD [HL], C", 1, 8, c.ld_hl(Op8R8Src)),
		0x72: NewInstruction("LD [HL], D", 1, 8, c.ld_hl(Op8R8Src)),
		0x73: NewInstruction("LD [HL], E", 1, 8, c.ld_hl(Op8R8Src)),
		0x74: NewInstruction("LD [HL], H", 1, 8, c.ld_hl(Op8R8Src)),
		0x75: NewInstruction("LD [HL], L", 1, 8, c.ld_hl(Op8R8Src)),
		0x77: NewInstruction("LD [HL], A", 1, 8, c.ld_hl(Op8R8Src)),
		0x36: NewInstruction("LD [HL], n8", 2, 12, c.ld_hl(Op8N8)),

		0x78: NewInstruction("LD A, B", 1, 4, c.ld_r8(Op8R8Src)),
		0x79: NewInstruction("LD A, C", 1, 4, c.ld_r8(Op8R8Src)),
		0x7a: NewInstruction("LD A, D", 1, 4, c.ld_r8(Op8R8Src)),
		0x7b: NewInstruction("LD A, E", 1, 4, c.ld_r8(Op8R8Src)),
		0x7c: NewInstruction("LD A, H", 1, 4, c.ld_r8(Op8R8Src)),
		0x7d: NewInstruction("LD A, L", 1, 4, c.ld_r8(Op8R8Src)),
		0x7e: NewInstruction("LD A, [HL]", 1, 8, c.ld_r8(Op8HL)),
		0x7f: NewInstruction("LD A, A", 1, 4, c.nop),
		0x3e: NewInstruction("LD A, n8", 2, 8, c.ld_r8(Op8N8)),

		0x02: NewInstruction("LD [BC], A", 1, 8, c.ld_r16_a),
		0x12: NewInstruction("LD [DE], A", 1, 8, c.ld_r16_a),
		0x22: NewInstruction("LDI [HL], A", 1, 8, c.ldi_hl_a),
		0x32: NewInstruction("LDD [HL], A", 1, 8, c.ldd_hl_a),

		0x0a: NewInstruction("LD A, [BC]", 1, 8, c.ld_a_r16),
		0x1a: NewInstruction("LD A, [DE]", 1, 8, c.ld_a_r16),
		0x2a: NewInstruction("LDI A, HL", 1, 8, c.ldi_a_hl),
		0x3a: NewInstruction("LDD A, HL", 1, 8, c.ldd_a_hl),

		0xe0: NewInstruction("LDH [a8], A", 2, 12, c.ldh_a8_a),
		0xf0: NewInstruction("LDH A, [a8]", 2, 12, c.ldh_a_a8),
		0xe2: NewInstruction("LDH [C], A", 1, 8, c.ldh_c_a),
		0xf2: NewInstruction("LDH A, [C]", 1, 8, c.ldh_a_c),

		0x01: NewInstruction("LD BC, n16", 3, 12, c.ld_r16_n16),
		0x11: NewInstruction("LD DE, n16", 3, 12, c.ld_r16_n16),
		0x21: NewInstruction("LD HL, n16", 3, 12, c.ld_r16_n16),
		0x31: NewInstruction("LD SP, n16", 3, 12, c.ld_r16_n16),

		0xea: NewInstruction("LD [a16], A", 3, 16, c.ld_a16_a),
		0xfa: NewInstruction("LD A, [a16]", 3, 16, c.la_a_a16),
		0x08: NewInstruction("LD [a16], SP", 3, 20, c.ld_a16_sp),
		0xf9: NewInstruction("LD SP, HL", 1, 8, c.ld_sp_hl),

		0xc5: NewInstruction("PUSH BC", 1, 16, c.push_r16),
		0xd5: NewInstruction("PUSH DE", 1, 16, c.push_r16),
		0xe5: NewInstruction("PUSH HL", 1, 16, c.push_r16),
		0xf5: NewInstruction("PUSH AF", 1, 16, c.push_r16),

		0xc1: NewInstruction("POP BC", 1, 16, c.pop_r16),
		0xd1: NewInstruction("POP DE", 1, 16, c.pop_r16),
		0xe1: NewInstruction("POP HL", 1, 16, c.pop_r16),
		0xf1: NewInstruction("POP AF", 1, 16, c.pop_r16),

		0xf8: NewInstruction("LD HL, SP + e8", 2, 12, c.ld_hl_sp_e),

		0x80: NewInstruction("ADD A, B", 1, 4, c.add_a(Op8R8Src, false)),
		0x81: NewInstruction("ADD A, C", 1, 4, c.add_a(Op8R8Src, false)),
		0x82: NewInstruction("ADD A, D", 1, 4, c.add_a(Op8R8Src, false)),
		0x83: NewInstruction("ADD A, E", 1, 4, c.add_a(Op8R8Src, false)),
		0x84: NewInstruction("ADD A, H", 1, 4, c.add_a(Op8R8Src, false)),
		0x85: NewInstruction("ADD A, L", 1, 4, c.add_a(Op8R8Src, false)),
		0x86: NewInstruction("ADD A, [HL]", 1, 8, c.add_a(Op8HL, false)),
		0x87: NewInstruction("ADD A, A", 1, 4, c.add_a(Op8R8Src, false)),
		0xc6: NewInstruction("ADD A, n8", 2, 8, c.add_a(Op8N8, false)),

		0x88: NewInstruction("ADC A, B", 1, 4, c.add_a(Op8R8Src, true)),
		0x89: NewInstruction("ADC A, C", 1, 4, c.add_a(Op8R8Src, true)),
		0x8a: NewInstruction("ADC A, D", 1, 4, c.add_a(Op8R8Src, true)),
		0x8b: NewInstruction("ADC A, E", 1, 4, c.add_a(Op8R8Src, true)),
		0x8c: NewInstruction("ADC A, H", 1, 4, c.add_a(Op8R8Src, true)),
		0x8d: NewInstruction("ADC A, L", 1, 4, c.add_a(Op8R8Src, true)),
		0x8e: NewInstruction("ADC A, [HL]", 1, 8, c.add_a(Op8HL, true)),
		0x8f: NewInstruction("ADC A, A", 1, 4, c.add_a(Op8R8Src, true)),
		0xce: NewInstruction("ADC A, n8", 2, 8, c.add_a(Op8N8, true)),

		0x90: NewInstruction("SUB A, B", 1, 4, c.sub_a(Op8R8Src, false)),
		0x91: NewInstruction("SUB A, C", 1, 4, c.sub_a(Op8R8Src, false)),
		0x92: NewInstruction("SUB A, D", 1, 4, c.sub_a(Op8R8Src, false)),
		0x93: NewInstruction("SUB A, E", 1, 4, c.sub_a(Op8R8Src, false)),
		0x94: NewInstruction("SUB A, H", 1, 4, c.sub_a(Op8R8Src, false)),
		0x95: NewInstruction("SUB A, L", 1, 4, c.sub_a(Op8R8Src, false)),
		0x96: NewInstruction("SUB A, [HL]", 1, 8, c.sub_a(Op8HL, false)),
		0x97: NewInstruction("SUB A, A", 1, 4, c.sub_a(Op8R8Src, false)),
		0xd6: NewInstruction("SUB A, n8", 2, 8, c.sub_a(Op8N8, false)),

		0x98: NewInstruction("SBC A, B", 1, 4, c.sub_a(Op8R8Src, true)),
		0x99: NewInstruction("SBC A, C", 1, 4, c.sub_a(Op8R8Src, true)),
		0x9a: NewInstruction("SBC A, D", 1, 4, c.sub_a(Op8R8Src, true)),
		0x9b: NewInstruction("SBC A, E", 1, 4, c.sub_a(Op8R8Src, true)),
		0x9c: NewInstruction("SBC A, H", 1, 4, c.sub_a(Op8R8Src, true)),
		0x9d: NewInstruction("SBC A, L", 1, 4, c.sub_a(Op8R8Src, true)),
		0x9e: NewInstruction("SBC A, [HL]", 1, 8, c.sub_a(Op8HL, true)),
		0x9f: NewInstruction("SBC A, A", 1, 4, c.sub_a(Op8R8Src, true)),
		0xde: NewInstruction("SBC A, n8", 2, 8, c.sub_a(Op8N8, true)),

		0xb8: NewInstruction("CP A, B", 1, 4, c.cp_a(Op8R8Src)),
		0xb9: NewInstruction("CP A, C", 1, 4, c.cp_a(Op8R8Src)),
		0xba: NewInstruction("CP A, D", 1, 4, c.cp_a(Op8R8Src)),
		0xbb: NewInstruction("CP A, E", 1, 4, c.cp_a(Op8R8Src)),
		0xbc: NewInstruction("CP A, H", 1, 4, c.cp_a(Op8R8Src)),
		0xbd: NewInstruction("CP A, L", 1, 4, c.cp_a(Op8R8Src)),
		0xbe: NewInstruction("CP A, [HL]", 1, 8, c.cp_a(Op8HL)),
		0xbf: NewInstruction("CP A, A", 1, 4, c.cp_a(Op8R8Src)),
		0xfe: NewInstruction("CP A, n8", 2, 8, c.cp_a(Op8N8)),

		0x04: NewInstruction("INC B", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x0c: NewInstruction("INC C", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x14: NewInstruction("INC D", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x1c: NewInstruction("INC E", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x24: NewInstruction("INC H", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x2c: NewInstruction("INC L", 1, 4, c.inc_8bit(Op8R8Dest)),
		0x34: NewInstruction("INC [HL]", 1, 12, c.inc_8bit(Op8HL)),
		0x3c: NewInstruction("INC A", 1, 4, c.inc_8bit(Op8R8Dest)),

		0x05: NewInstruction("DEC B", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x0d: NewInstruction("DEC C", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x15: NewInstruction("DEC D", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x1d: NewInstruction("DEC E", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x25: NewInstruction("DEC H", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x2d: NewInstruction("DEC L", 1, 4, c.dec_8bit(Op8R8Dest)),
		0x35: NewInstruction("DEC [HL]", 1, 12, c.dec_8bit(Op8HL)),
		0x3d: NewInstruction("DEC A", 1, 4, c.dec_8bit(Op8R8Dest)),

		0xa0: NewInstruction("AND A, B", 1, 4, c.and_a(Op8R8Src)),
		0xa1: NewInstruction("AND A, C", 1, 4, c.and_a(Op8R8Src)),
		0xa2: NewInstruction("AND A, D", 1, 4, c.and_a(Op8R8Src)),
		0xa3: NewInstruction("AND A, E", 1, 4, c.and_a(Op8R8Src)),
		0xa4: NewInstruction("AND A, H", 1, 4, c.and_a(Op8R8Src)),
		0xa5: NewInstruction("AND A, L", 1, 4, c.and_a(Op8R8Src)),
		0xa6: NewInstruction("AND A, [HL]", 1, 8, c.and_a(Op8HL)),
		0xa7: NewInstruction("AND A, A", 1, 4, c.and_a(Op8R8Src)),
		0xe6: NewInstruction("AND A, n8", 2, 8, c.and_a(Op8N8)),

		0xb0: NewInstruction("OR A, B", 1, 4, c.or_a(Op8R8Src)),
		0xb1: NewInstruction("OR A, C", 1, 4, c.or_a(Op8R8Src)),
		0xb2: NewInstruction("OR A, D", 1, 4, c.or_a(Op8R8Src)),
		0xb3: NewInstruction("OR A, E", 1, 4, c.or_a(Op8R8Src)),
		0xb4: NewInstruction("OR A, H", 1, 4, c.or_a(Op8R8Src)),
		0xb5: NewInstruction("OR A, L", 1, 4, c.or_a(Op8R8Src)),
		0xb6: NewInstruction("OR A, [HL]", 1, 8, c.or_a(Op8HL)),
		0xb7: NewInstruction("OR A, A", 1, 4, c.or_a(Op8R8Src)),
		0xf6: NewInstruction("OR A, n8", 2, 8, c.or_a(Op8N8)),

		0xa8: NewInstruction("XOR A, B", 1, 4, c.xor_a(Op8R8Src)),
		0xa9: NewInstruction("XOR A, C", 1, 4, c.xor_a(Op8R8Src)),
		0xaa: NewInstruction("XOR A, D", 1, 4, c.xor_a(Op8R8Src)),
		0xab: NewInstruction("XOR A, E", 1, 4, c.xor_a(Op8R8Src)),
		0xac: NewInstruction("XOR A, H", 1, 4, c.xor_a(Op8R8Src)),
		0xad: NewInstruction("XOR A, L", 1, 4, c.xor_a(Op8R8Src)),
		0xae: NewInstruction("XOR A, [HL]", 1, 8, c.xor_a(Op8HL)),
		0xaf: NewInstruction("XOR A, A", 1, 4, c.xor_a(Op8R8Src)),
		0xee: NewInstruction("XOR A, n8", 2, 8, c.xor_a(Op8N8)),

		0x3f: NewInstruction("CCF", 1, 4, c.ccf),
		0x37: NewInstruction("SCF", 1, 4, c.scf),
		0x27: NewInstruction("CPL", 1, 4, c.cpl),

		0x03: NewInstruction("INC BC", 1, 8, c.inc_16bit),
		0x13: NewInstruction("INC DE", 1, 8, c.inc_16bit),
		0x23: NewInstruction("INC HL", 1, 8, c.inc_16bit),
		0x33: NewInstruction("INC SP", 1, 8, c.inc_16bit),

		0x0b: NewInstruction("DEC BC", 1, 8, c.dec_16bit),
		0x1b: NewInstruction("DEC DE", 1, 8, c.dec_16bit),
		0x2b: NewInstruction("DEC HL", 1, 8, c.dec_16bit),
		0x3b: NewInstruction("DEC SP", 1, 8, c.dec_16bit),

		0x09: NewInstruction("ADD HL, BC", 1, 8, c.add_hl),
		0x19: NewInstruction("ADD HL, DE", 1, 8, c.add_hl),
		0x29: NewInstruction("ADD HL, HL", 1, 8, c.add_hl),
		0x39: NewInstruction("ADD HL, SP", 1, 8, c.add_hl),
		0xe8: NewInstruction("ADD SP, e8", 2, 16, c.add_sp_e),

		0x07: NewInstruction("RLCA", 1, 4, c.rl(Op8R8Src, true)),
		0x0f: NewInstruction("RRCA", 1, 4, c.rr(Op8R8Src, true)),
		0x17: NewInstruction("RLA", 1, 4, c.rl(Op8R8Src, false)),
		0x1f: NewInstruction("RRA", 1, 4, c.rr(Op8R8Src, false)),
	}
}

func (c *CPU) buildCbPrefixedInstructions() {
	c.cbPrefixedInstructions = map[uint8]Instruction{
		0x00: NewInstruction("RLC B", 2, 8, c.rl(Op8R8Src, true)),
		0x01: NewInstruction("RLC C", 2, 8, c.rl(Op8R8Src, true)),
		0x02: NewInstruction("RLC D", 2, 8, c.rl(Op8R8Src, true)),
		0x03: NewInstruction("RLC E", 2, 8, c.rl(Op8R8Src, true)),
		0x04: NewInstruction("RLC H", 2, 8, c.rl(Op8R8Src, true)),
		0x05: NewInstruction("RLC L", 2, 8, c.rl(Op8R8Src, true)),
		0x06: NewInstruction("RLC [HL]", 2, 16, c.rl(Op8HL, true)),
		0x07: NewInstruction("RLC A", 2, 8, c.rl(Op8R8Src, true)),

		0x08: NewInstruction("RRC B", 2, 8, c.rr(Op8R8Src, true)),
		0x09: NewInstruction("RRC C", 2, 8, c.rr(Op8R8Src, true)),
		0x0a: NewInstruction("RRC D", 2, 8, c.rr(Op8R8Src, true)),
		0x0b: NewInstruction("RRC E", 2, 8, c.rr(Op8R8Src, true)),
		0x0c: NewInstruction("RRC H", 2, 8, c.rr(Op8R8Src, true)),
		0x0d: NewInstruction("RRC L", 2, 8, c.rr(Op8R8Src, true)),
		0x0e: NewInstruction("RRC [HL]", 2, 16, c.rr(Op8HL, true)),
		0x0f: NewInstruction("RRC A", 2, 8, c.rr(Op8R8Src, true)),

		0x10: NewInstruction("RL B", 2, 8, c.rl(Op8R8Src, false)),
		0x11: NewInstruction("RL C", 2, 8, c.rl(Op8R8Src, false)),
		0x12: NewInstruction("RL D", 2, 8, c.rl(Op8R8Src, false)),
		0x13: NewInstruction("RL E", 2, 8, c.rl(Op8R8Src, false)),
		0x14: NewInstruction("RL H", 2, 8, c.rl(Op8R8Src, false)),
		0x15: NewInstruction("RL L", 2, 8, c.rl(Op8R8Src, false)),
		0x16: NewInstruction("RL [HL]", 2, 16, c.rl(Op8HL, false)),
		0x17: NewInstruction("RL A", 2, 8, c.rl(Op8R8Src, false)),

		0x18: NewInstruction("RR B", 2, 8, c.rr(Op8R8Src, false)),
		0x19: NewInstruction("RR C", 2, 8, c.rr(Op8R8Src, false)),
		0x1a: NewInstruction("RR D", 2, 8, c.rr(Op8R8Src, false)),
		0x1b: NewInstruction("RR E", 2, 8, c.rr(Op8R8Src, false)),
		0x1c: NewInstruction("RR H", 2, 8, c.rr(Op8R8Src, false)),
		0x1d: NewInstruction("RR L", 2, 8, c.rr(Op8R8Src, false)),
		0x1e: NewInstruction("RR [HL]", 2, 16, c.rr(Op8HL, false)),
		0x1f: NewInstruction("RR A", 2, 8, c.rr(Op8R8Src, false)),

		0x20: NewInstruction("SLA B", 2, 8, c.sla(Op8R8Src)),
		0x21: NewInstruction("SLA C", 2, 8, c.sla(Op8R8Src)),
		0x22: NewInstruction("SLA D", 2, 8, c.sla(Op8R8Src)),
		0x23: NewInstruction("SLA E", 2, 8, c.sla(Op8R8Src)),
		0x24: NewInstruction("SLA H", 2, 8, c.sla(Op8R8Src)),
		0x25: NewInstruction("SLA L", 2, 8, c.sla(Op8R8Src)),
		0x26: NewInstruction("SLA [HL]", 2, 16, c.sla(Op8HL)),
		0x27: NewInstruction("SLA A", 2, 8, c.sla(Op8R8Src)),

		0x28: NewInstruction("SRA B", 2, 8, c.sra(Op8R8Src)),
		0x29: NewInstruction("SRA C", 2, 8, c.sra(Op8R8Src)),
		0x2a: NewInstruction("SRA D", 2, 8, c.sra(Op8R8Src)),
		0x2b: NewInstruction("SRA E", 2, 8, c.sra(Op8R8Src)),
		0x2c: NewInstruction("SRA H", 2, 8, c.sra(Op8R8Src)),
		0x2d: NewInstruction("SRA L", 2, 8, c.sra(Op8R8Src)),
		0x2e: NewInstruction("SRA [HL]", 2, 16, c.sra(Op8HL)),
		0x2f: NewInstruction("SRA A", 2, 8, c.sra(Op8R8Src)),

		0x30: NewInstruction("SWAP B", 2, 8, c.swap(Op8R8Src)),
		0x31: NewInstruction("SWAP C", 2, 8, c.swap(Op8R8Src)),
		0x32: NewInstruction("SWAP D", 2, 8, c.swap(Op8R8Src)),
		0x33: NewInstruction("SWAP E", 2, 8, c.swap(Op8R8Src)),
		0x34: NewInstruction("SWAP H", 2, 8, c.swap(Op8R8Src)),
		0x35: NewInstruction("SWAP L", 2, 8, c.swap(Op8R8Src)),
		0x36: NewInstruction("SWAP [HL]", 2, 16, c.swap(Op8HL)),
		0x37: NewInstruction("SWAP A", 2, 8, c.swap(Op8R8Src)),

		0x38: NewInstruction("SRL B", 2, 8, c.srl(Op8R8Src)),
		0x39: NewInstruction("SRL C", 2, 8, c.srl(Op8R8Src)),
		0x3a: NewInstruction("SRL D", 2, 8, c.srl(Op8R8Src)),
		0x3b: NewInstruction("SRL E", 2, 8, c.srl(Op8R8Src)),
		0x3c: NewInstruction("SRL H", 2, 8, c.srl(Op8R8Src)),
		0x3d: NewInstruction("SRL L", 2, 8, c.srl(Op8R8Src)),
		0x3e: NewInstruction("SRL [HL]", 2, 16, c.srl(Op8HL)),
		0x3f: NewInstruction("SRL A", 2, 8, c.srl(Op8R8Src)),

		0x40: NewInstruction("BIT 0, B", 2, 8, c.bit(Op8R8Src)),
		0x41: NewInstruction("BIT 0, C", 2, 8, c.bit(Op8R8Src)),
		0x42: NewInstruction("BIT 0, D", 2, 8, c.bit(Op8R8Src)),
		0x43: NewInstruction("BIT 0, E", 2, 8, c.bit(Op8R8Src)),
		0x44: NewInstruction("BIT 0, H", 2, 8, c.bit(Op8R8Src)),
		0x45: NewInstruction("BIT 0, L", 2, 8, c.bit(Op8R8Src)),
		0x46: NewInstruction("BIT 0, [HL]", 2, 12, c.bit(Op8HL)),
		0x47: NewInstruction("BIT 0, A", 2, 8, c.bit(Op8R8Src)),

		0x48: NewInstruction("BIT 1, B", 2, 8, c.bit(Op8R8Src)),
		0x49: NewInstruction("BIT 1, C", 2, 8, c.bit(Op8R8Src)),
		0x4a: NewInstruction("BIT 1, D", 2, 8, c.bit(Op8R8Src)),
		0x4b: NewInstruction("BIT 1, E", 2, 8, c.bit(Op8R8Src)),
		0x4c: NewInstruction("BIT 1, H", 2, 8, c.bit(Op8R8Src)),
		0x4d: NewInstruction("BIT 1, L", 2, 8, c.bit(Op8R8Src)),
		0x4e: NewInstruction("BIT 1, [HL]", 2, 12, c.bit(Op8HL)),
		0x4f: NewInstruction("BIT 1, A", 2, 8, c.bit(Op8R8Src)),

		0x50: NewInstruction("BIT 2, B", 2, 8, c.bit(Op8R8Src)),
		0x51: NewInstruction("BIT 2, C", 2, 8, c.bit(Op8R8Src)),
		0x52: NewInstruction("BIT 2, D", 2, 8, c.bit(Op8R8Src)),
		0x53: NewInstruction("BIT 2, E", 2, 8, c.bit(Op8R8Src)),
		0x54: NewInstruction("BIT 2, H", 2, 8, c.bit(Op8R8Src)),
		0x55: NewInstruction("BIT 2, L", 2, 8, c.bit(Op8R8Src)),
		0x56: NewInstruction("BIT 2, [HL]", 2, 12, c.bit(Op8HL)),
		0x57: NewInstruction("BIT 2, A", 2, 8, c.bit(Op8R8Src)),

		0x58: NewInstruction("BIT 3, B", 2, 8, c.bit(Op8R8Src)),
		0x59: NewInstruction("BIT 3, C", 2, 8, c.bit(Op8R8Src)),
		0x5a: NewInstruction("BIT 3, D", 2, 8, c.bit(Op8R8Src)),
		0x5b: NewInstruction("BIT 3, E", 2, 8, c.bit(Op8R8Src)),
		0x5c: NewInstruction("BIT 3, H", 2, 8, c.bit(Op8R8Src)),
		0x5d: NewInstruction("BIT 3, L", 2, 8, c.bit(Op8R8Src)),
		0x5e: NewInstruction("BIT 3, [HL]", 2, 12, c.bit(Op8HL)),
		0x5f: NewInstruction("BIT 3, A", 2, 8, c.bit(Op8R8Src)),

		0x60: NewInstruction("BIT 4, B", 2, 8, c.bit(Op8R8Src)),
		0x61: NewInstruction("BIT 4, C", 2, 8, c.bit(Op8R8Src)),
		0x62: NewInstruction("BIT 4, D", 2, 8, c.bit(Op8R8Src)),
		0x63: NewInstruction("BIT 4, E", 2, 8, c.bit(Op8R8Src)),
		0x64: NewInstruction("BIT 4, H", 2, 8, c.bit(Op8R8Src)),
		0x65: NewInstruction("BIT 4, L", 2, 8, c.bit(Op8R8Src)),
		0x66: NewInstruction("BIT 4, [HL]", 2, 12, c.bit(Op8HL)),
		0x67: NewInstruction("BIT 4, A", 2, 8, c.bit(Op8R8Src)),

		0x68: NewInstruction("BIT 5, B", 2, 8, c.bit(Op8R8Src)),
		0x69: NewInstruction("BIT 5, C", 2, 8, c.bit(Op8R8Src)),
		0x6a: NewInstruction("BIT 5, D", 2, 8, c.bit(Op8R8Src)),
		0x6b: NewInstruction("BIT 5, E", 2, 8, c.bit(Op8R8Src)),
		0x6c: NewInstruction("BIT 5, H", 2, 8, c.bit(Op8R8Src)),
		0x6d: NewInstruction("BIT 5, L", 2, 8, c.bit(Op8R8Src)),
		0x6e: NewInstruction("BIT 5, [HL]", 2, 12, c.bit(Op8HL)),
		0x6f: NewInstruction("BIT 5, A", 2, 8, c.bit(Op8R8Src)),

		0x70: NewInstruction("BIT 6, B", 2, 8, c.bit(Op8R8Src)),
		0x71: NewInstruction("BIT 6, C", 2, 8, c.bit(Op8R8Src)),
		0x72: NewInstruction("BIT 6, D", 2, 8, c.bit(Op8R8Src)),
		0x73: NewInstruction("BIT 6, E", 2, 8, c.bit(Op8R8Src)),
		0x74: NewInstruction("BIT 6, H", 2, 8, c.bit(Op8R8Src)),
		0x75: NewInstruction("BIT 6, L", 2, 8, c.bit(Op8R8Src)),
		0x76: NewInstruction("BIT 6, [HL]", 2, 12, c.bit(Op8HL)),
		0x77: NewInstruction("BIT 6, A", 2, 8, c.bit(Op8R8Src)),

		0x78: NewInstruction("BIT 7, B", 2, 8, c.bit(Op8R8Src)),
		0x79: NewInstruction("BIT 7, C", 2, 8, c.bit(Op8R8Src)),
		0x7a: NewInstruction("BIT 7, D", 2, 8, c.bit(Op8R8Src)),
		0x7b: NewInstruction("BIT 7, E", 2, 8, c.bit(Op8R8Src)),
		0x7c: NewInstruction("BIT 7, H", 2, 8, c.bit(Op8R8Src)),
		0x7d: NewInstruction("BIT 7, L", 2, 8, c.bit(Op8R8Src)),
		0x7e: NewInstruction("BIT 7, [HL]", 2, 12, c.bit(Op8HL)),
		0x7f: NewInstruction("BIT 7, A", 2, 8, c.bit(Op8R8Src)),

		0x80: NewInstruction("RES 0, B", 2, 8, c.res(Op8R8Src)),
		0x81: NewInstruction("RES 0, C", 2, 8, c.res(Op8R8Src)),
		0x82: NewInstruction("RES 0, D", 2, 8, c.res(Op8R8Src)),
		0x83: NewInstruction("RES 0, E", 2, 8, c.res(Op8R8Src)),
		0x84: NewInstruction("RES 0, H", 2, 8, c.res(Op8R8Src)),
		0x85: NewInstruction("RES 0, L", 2, 8, c.res(Op8R8Src)),
		0x86: NewInstruction("RES 0, [HL]", 2, 16, c.res(Op8HL)),
		0x87: NewInstruction("RES 0, A", 2, 8, c.res(Op8R8Src)),

		0x88: NewInstruction("RES 1, B", 2, 8, c.res(Op8R8Src)),
		0x89: NewInstruction("RES 1, C", 2, 8, c.res(Op8R8Src)),
		0x8a: NewInstruction("RES 1, D", 2, 8, c.res(Op8R8Src)),
		0x8b: NewInstruction("RES 1, E", 2, 8, c.res(Op8R8Src)),
		0x8c: NewInstruction("RES 1, H", 2, 8, c.res(Op8R8Src)),
		0x8d: NewInstruction("RES 1, L", 2, 8, c.res(Op8R8Src)),
		0x8e: NewInstruction("RES 1, [HL]", 2, 16, c.res(Op8HL)),
		0x8f: NewInstruction("RES 1, A", 2, 8, c.res(Op8R8Src)),

		0x90: NewInstruction("RES 2, B", 2, 8, c.res(Op8R8Src)),
		0x91: NewInstruction("RES 2, C", 2, 8, c.res(Op8R8Src)),
		0x92: NewInstruction("RES 2, D", 2, 8, c.res(Op8R8Src)),
		0x93: NewInstruction("RES 2, E", 2, 8, c.res(Op8R8Src)),
		0x94: NewInstruction("RES 2, H", 2, 8, c.res(Op8R8Src)),
		0x95: NewInstruction("RES 2, L", 2, 8, c.res(Op8R8Src)),
		0x96: NewInstruction("RES 2, [HL]", 2, 16, c.res(Op8HL)),
		0x97: NewInstruction("RES 2, A", 2, 8, c.res(Op8R8Src)),

		0x98: NewInstruction("RES 3, B", 2, 8, c.res(Op8R8Src)),
		0x99: NewInstruction("RES 3, C", 2, 8, c.res(Op8R8Src)),
		0x9a: NewInstruction("RES 3, D", 2, 8, c.res(Op8R8Src)),
		0x9b: NewInstruction("RES 3, E", 2, 8, c.res(Op8R8Src)),
		0x9c: NewInstruction("RES 3, H", 2, 8, c.res(Op8R8Src)),
		0x9d: NewInstruction("RES 3, L", 2, 8, c.res(Op8R8Src)),
		0x9e: NewInstruction("RES 3, [HL]", 2, 16, c.res(Op8HL)),
		0x9f: NewInstruction("RES 3, A", 2, 8, c.res(Op8R8Src)),

		0xa0: NewInstruction("RES 4, B", 2, 8, c.res(Op8R8Src)),
		0xa1: NewInstruction("RES 4, C", 2, 8, c.res(Op8R8Src)),
		0xa2: NewInstruction("RES 4, D", 2, 8, c.res(Op8R8Src)),
		0xa3: NewInstruction("RES 4, E", 2, 8, c.res(Op8R8Src)),
		0xa4: NewInstruction("RES 4, H", 2, 8, c.res(Op8R8Src)),
		0xa5: NewInstruction("RES 4, L", 2, 8, c.res(Op8R8Src)),
		0xa6: NewInstruction("RES 4, [HL]", 2, 16, c.res(Op8HL)),
		0xa7: NewInstruction("RES 4, A", 2, 8, c.res(Op8R8Src)),

		0xa8: NewInstruction("RES 5, B", 2, 8, c.res(Op8R8Src)),
		0xa9: NewInstruction("RES 5, C", 2, 8, c.res(Op8R8Src)),
		0xaa: NewInstruction("RES 5, D", 2, 8, c.res(Op8R8Src)),
		0xab: NewInstruction("RES 5, E", 2, 8, c.res(Op8R8Src)),
		0xac: NewInstruction("RES 5, H", 2, 8, c.res(Op8R8Src)),
		0xad: NewInstruction("RES 5, L", 2, 8, c.res(Op8R8Src)),
		0xae: NewInstruction("RES 5, [HL]", 2, 16, c.res(Op8HL)),
		0xaf: NewInstruction("RES 5, A", 2, 8, c.res(Op8R8Src)),

		0xb0: NewInstruction("RES 6, B", 2, 8, c.res(Op8R8Src)),
		0xb1: NewInstruction("RES 6, C", 2, 8, c.res(Op8R8Src)),
		0xb2: NewInstruction("RES 6, D", 2, 8, c.res(Op8R8Src)),
		0xb3: NewInstruction("RES 6, E", 2, 8, c.res(Op8R8Src)),
		0xb4: NewInstruction("RES 6, H", 2, 8, c.res(Op8R8Src)),
		0xb5: NewInstruction("RES 6, L", 2, 8, c.res(Op8R8Src)),
		0xb6: NewInstruction("RES 6, [HL]", 2, 16, c.res(Op8HL)),
		0xb7: NewInstruction("RES 6, A", 2, 8, c.res(Op8R8Src)),

		0xb8: NewInstruction("RES 7, B", 2, 8, c.res(Op8R8Src)),
		0xb9: NewInstruction("RES 7, C", 2, 8, c.res(Op8R8Src)),
		0xba: NewInstruction("RES 7, D", 2, 8, c.res(Op8R8Src)),
		0xbb: NewInstruction("RES 7, E", 2, 8, c.res(Op8R8Src)),
		0xbc: NewInstruction("RES 7, H", 2, 8, c.res(Op8R8Src)),
		0xbd: NewInstruction("RES 7, L", 2, 8, c.res(Op8R8Src)),
		0xbe: NewInstruction("RES 7, [HL]", 2, 16, c.res(Op8HL)),
		0xbf: NewInstruction("RES 7, A", 2, 8, c.res(Op8R8Src)),

		0xc0: NewInstruction("SET 0, B", 2, 8, c.set(Op8R8Src)),
		0xc1: NewInstruction("SET 0, C", 2, 8, c.set(Op8R8Src)),
		0xc2: NewInstruction("SET 0, D", 2, 8, c.set(Op8R8Src)),
		0xc3: NewInstruction("SET 0, E", 2, 8, c.set(Op8R8Src)),
		0xc4: NewInstruction("SET 0, H", 2, 8, c.set(Op8R8Src)),
		0xc5: NewInstruction("SET 0, L", 2, 8, c.set(Op8R8Src)),
		0xc6: NewInstruction("SET 0, [HL]", 2, 16, c.set(Op8HL)),
		0xc7: NewInstruction("SET 0, A", 2, 8, c.set(Op8R8Src)),

		0xc8: NewInstruction("SET 1, B", 2, 8, c.set(Op8R8Src)),
		0xc9: NewInstruction("SET 1, C", 2, 8, c.set(Op8R8Src)),
		0xca: NewInstruction("SET 1, D", 2, 8, c.set(Op8R8Src)),
		0xcb: NewInstruction("SET 1, E", 2, 8, c.set(Op8R8Src)),
		0xcc: NewInstruction("SET 1, H", 2, 8, c.set(Op8R8Src)),
		0xcd: NewInstruction("SET 1, L", 2, 8, c.set(Op8R8Src)),
		0xce: NewInstruction("SET 1, [HL]", 2, 16, c.set(Op8HL)),
		0xcf: NewInstruction("SET 1, A", 2, 8, c.set(Op8R8Src)),

		0xd0: NewInstruction("SET 2, B", 2, 8, c.set(Op8R8Src)),
		0xd1: NewInstruction("SET 2, C", 2, 8, c.set(Op8R8Src)),
		0xd2: NewInstruction("SET 2, D", 2, 8, c.set(Op8R8Src)),
		0xd3: NewInstruction("SET 2, E", 2, 8, c.set(Op8R8Src)),
		0xd4: NewInstruction("SET 2, H", 2, 8, c.set(Op8R8Src)),
		0xd5: NewInstruction("SET 2, L", 2, 8, c.set(Op8R8Src)),
		0xd6: NewInstruction("SET 2, [HL]", 2, 16, c.set(Op8HL)),
		0xd7: NewInstruction("SET 2, A", 2, 8, c.set(Op8R8Src)),

		0xd8: NewInstruction("SET 3, B", 2, 8, c.set(Op8R8Src)),
		0xd9: NewInstruction("SET 3, C", 2, 8, c.set(Op8R8Src)),
		0xda: NewInstruction("SET 3, D", 2, 8, c.set(Op8R8Src)),
		0xdb: NewInstruction("SET 3, E", 2, 8, c.set(Op8R8Src)),
		0xdc: NewInstruction("SET 3, H", 2, 8, c.set(Op8R8Src)),
		0xdd: NewInstruction("SET 3, L", 2, 8, c.set(Op8R8Src)),
		0xde: NewInstruction("SET 3, [HL]", 2, 16, c.set(Op8HL)),
		0xdf: NewInstruction("SET 3, A", 2, 8, c.set(Op8R8Src)),

		0xe0: NewInstruction("SET 4, B", 2, 8, c.set(Op8R8Src)),
		0xe1: NewInstruction("SET 4, C", 2, 8, c.set(Op8R8Src)),
		0xe2: NewInstruction("SET 4, D", 2, 8, c.set(Op8R8Src)),
		0xe3: NewInstruction("SET 4, E", 2, 8, c.set(Op8R8Src)),
		0xe4: NewInstruction("SET 4, H", 2, 8, c.set(Op8R8Src)),
		0xe5: NewInstruction("SET 4, L", 2, 8, c.set(Op8R8Src)),
		0xe6: NewInstruction("SET 4, [HL]", 2, 16, c.set(Op8HL)),
		0xe7: NewInstruction("SET 4, A", 2, 8, c.set(Op8R8Src)),

		0xe8: NewInstruction("SET 5, B", 2, 8, c.set(Op8R8Src)),
		0xe9: NewInstruction("SET 5, C", 2, 8, c.set(Op8R8Src)),
		0xea: NewInstruction("SET 5, D", 2, 8, c.set(Op8R8Src)),
		0xeb: NewInstruction("SET 5, E", 2, 8, c.set(Op8R8Src)),
		0xec: NewInstruction("SET 5, H", 2, 8, c.set(Op8R8Src)),
		0xed: NewInstruction("SET 5, L", 2, 8, c.set(Op8R8Src)),
		0xee: NewInstruction("SET 5, [HL]", 2, 16, c.set(Op8HL)),
		0xef: NewInstruction("SET 5, A", 2, 8, c.set(Op8R8Src)),

		0xf0: NewInstruction("SET 6, B", 2, 8, c.set(Op8R8Src)),
		0xf1: NewInstruction("SET 6, C", 2, 8, c.set(Op8R8Src)),
		0xf2: NewInstruction("SET 6, D", 2, 8, c.set(Op8R8Src)),
		0xf3: NewInstruction("SET 6, E", 2, 8, c.set(Op8R8Src)),
		0xf4: NewInstruction("SET 6, H", 2, 8, c.set(Op8R8Src)),
		0xf5: NewInstruction("SET 6, L", 2, 8, c.set(Op8R8Src)),
		0xf6: NewInstruction("SET 6, [HL]", 2, 16, c.set(Op8HL)),
		0xf7: NewInstruction("SET 6, A", 2, 8, c.set(Op8R8Src)),

		0xf8: NewInstruction("SET 7, B", 2, 8, c.set(Op8R8Src)),
		0xf9: NewInstruction("SET 7, C", 2, 8, c.set(Op8R8Src)),
		0xfa: NewInstruction("SET 7, D", 2, 8, c.set(Op8R8Src)),
		0xfb: NewInstruction("SET 7, E", 2, 8, c.set(Op8R8Src)),
		0xfc: NewInstruction("SET 7, H", 2, 8, c.set(Op8R8Src)),
		0xfd: NewInstruction("SET 7, L", 2, 8, c.set(Op8R8Src)),
		0xfe: NewInstruction("SET 7, [HL]", 2, 16, c.set(Op8HL)),
		0xff: NewInstruction("SET 7, A", 2, 8, c.set(Op8R8Src)),
	}
}

// ixns.
func (c *CPU) nop(uint8) {}
func (c *CPU) halt(uint8) {
	c.isHalted = true
}
func (c *CPU) ld_r8(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		c.writeOp8Value(opcode, operandType, value)
	}
}
func (c *CPU) ld_hl(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		c.memory[c.readHL()] = c.getOp8Value(opcode, operandType)
	}
}
func (c *CPU) ld_r16_a(opcode uint8) {
	v := (opcode >> 4) & 1

	if v == 0 {
		c.memory[c.readBC()] = c.a
	} else {
		c.memory[c.readDE()] = c.a
	}
}
func (c *CPU) ld_a_r16(opcode uint8) {
	v := (opcode >> 4) & 1

	if v == 0 {
		c.a = c.memory[c.readBC()]
	} else {
		c.a = c.memory[c.readDE()]
	}
}
func (c *CPU) ldi_hl_a(uint8) {
	c.memory[c.readHL()] = c.a
	c.writeHL(c.readHL() + 1)
}
func (c *CPU) ldi_a_hl(uint8) {
	c.a = c.memory[c.readHL()]
	c.writeHL(c.readHL() + 1)
}
func (c *CPU) ldd_hl_a(uint8) {
	c.memory[c.readHL()] = c.a
	c.writeHL(c.readHL() - 1)
}
func (c *CPU) ldd_a_hl(uint8) {
	c.a = c.memory[c.readHL()]
	c.writeHL(c.readHL() - 1)
}
func (c *CPU) ldh_a8_a(uint8) {
	v := c.readNextByte()
	addr := 0xff00 + uint16(v)
	c.memory[addr] = c.a
}
func (c *CPU) ldh_a_a8(uint8) {
	v := c.readNextByte()
	addr := 0xff00 + uint16(v)
	c.a = c.memory[addr]
}
func (c *CPU) ldh_c_a(uint8) {
	addr := 0xff00 + uint16(c.c)
	c.memory[addr] = c.a
}
func (c *CPU) ldh_a_c(uint8) {
	addr := 0xff00 + uint16(c.c)
	c.a = c.memory[addr]
}
func (c *CPU) ld_r16_n16(opcode uint8) {
	v := (opcode & 0x30) >> 4
	n := c.readNextU16()

	switch v {
	case 0:
		c.writeBC(n)
	case 1:
		c.writeDE(n)
	case 2:
		c.writeHL(n)
	case 3:
		c.sp = n
	}
}
func (c *CPU) ld_a16_a(uint8) {
	nn := c.readNextU16()
	c.memory[nn] = c.a
}
func (c *CPU) la_a_a16(uint8) {
	nn := c.readNextU16()
	c.a = c.memory[nn]
}
func (c *CPU) ld_a16_sp(uint8) {
	nn := c.readNextU16()
	c.memory[nn] = uint8(c.sp & 0xff)
	c.memory[nn+1] = uint8((c.sp & 0xff00) >> 8)
}
func (c *CPU) ld_sp_hl(uint8) {
	c.sp = c.readHL()
}
func (c *CPU) push_r16(opcode uint8) {
	v := ((opcode & 0x30) >> 4) & 0b11

	switch v {
	case 0:
		c.stackPushU16(c.readBC())
	case 1:
		c.stackPushU16(c.readDE())
	case 2:
		c.stackPushU16(c.readHL())
	case 3:
		c.stackPushU16(c.readAF())
	}
}
func (c *CPU) pop_r16(opcode uint8) {
	v := ((opcode & 0x30) >> 4) & 0b11

	switch v {
	case 0:
		c.writeBC(c.stackPopU16())
	case 1:
		c.writeDE(c.stackPopU16())
	case 2:
		c.writeHL(c.stackPopU16())
	case 3:
		c.writeAF(c.stackPopU16())
	}
}
func (c *CPU) ld_hl_sp_e(uint8) {
	e := int8(c.readNextByte())
	result := uint16(int32(c.sp) + int32(e))

	halfCarryFlag := (uint16(c.sp&0xf) + uint16(e&0xf)) > 0xf
	carryFlag := (result & 0xff) < (c.sp & 0xff)

	c.flags.Unset(ZeroFlag)
	c.flags.Unset(SubtractFlag)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	c.writeHL(result)
}
func (c *CPU) add_a(operandType Op8Type, toCarry bool) func(uint8) {
	return func(opcode uint8) {
		operand := c.getOp8Value(opcode, operandType)
		c.addToRegA(operand, toCarry)
	}
}
func (c *CPU) sub_a(operandType Op8Type, toCarry bool) func(uint8) {
	return func(opcode uint8) {
		operand := c.getOp8Value(opcode, operandType)
		c.subFromRegA(operand, toCarry)
	}
}
func (c *CPU) cp_a(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		currA := c.a
		operand := c.getOp8Value(opcode, operandType)
		c.subFromRegA(operand, false)
		c.a = currA
	}
}
func (c *CPU) inc_8bit(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		result := c.perform8BitArithmetic(value, 1, false, false)
		c.writeOp8Value(opcode, operandType, result)
	}
}
func (c *CPU) dec_8bit(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		result := c.perform8BitArithmetic(value, 1, false, true)
		c.writeOp8Value(opcode, operandType, result)
	}
}
func (c *CPU) and_a(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		result := c.a & value
		c.a = result

		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Set(HalfCarryFlag)
		c.flags.Unset(CarryFlag)
	}
}
func (c *CPU) or_a(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		result := c.a | value
		c.a = result

		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.Unset(CarryFlag)
	}
}
func (c *CPU) xor_a(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		value := c.getOp8Value(opcode, operandType)
		result := c.a ^ value
		c.a = result

		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.Unset(CarryFlag)
	}
}
func (c *CPU) ccf(uint8) {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Toggle(CarryFlag)
}
func (c *CPU) scf(uint8) {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Set(CarryFlag)
}
func (c *CPU) cpl(uint8) {
	c.a = ^c.a
	c.flags.Set(SubtractFlag)
	c.flags.Set(HalfCarryFlag)
}
func (c *CPU) inc_16bit(opcode uint8) {
	v := (opcode & 0x30) >> 4

	switch v {
	case 0:
		c.writeBC(c.readBC() + 1)
	case 1:
		c.writeDE(c.readDE() + 1)
	case 2:
		c.writeHL(c.readHL() + 1)
	case 3:
		c.sp++
	}
}
func (c *CPU) dec_16bit(opcode uint8) {
	v := (opcode & 0x30) >> 4

	switch v {
	case 0:
		c.writeBC(c.readBC() - 1)
	case 1:
		c.writeDE(c.readDE() - 1)
	case 2:
		c.writeHL(c.readHL() - 1)
	case 3:
		c.sp--
	}
}
func (c *CPU) add_hl(opcode uint8) {
	v := (opcode & 0x30) >> 4
	value := uint16(0)

	switch v {
	case 0:
		value = c.readBC()
	case 1:
		value = c.readDE()
	case 2:
		value = c.readHL()
	case 3:
		value = c.sp
	}

	// ADD HL, rr doesn't set zero flag
	result := c.perform16BitAddition(c.readHL(), value, false)
	c.writeHL(result)
}
func (c *CPU) add_sp_e(uint8) {
	e := int8(c.readNextByte())
	result := uint16(int32(c.sp) + int32(e))

	halfCarryFlag := (uint16(c.sp&0xf) + uint16(e&0xf)) > 0xf
	carryFlag := (result & 0xff) < (c.sp & 0xff)

	c.flags.Unset(ZeroFlag)
	c.flags.Unset(SubtractFlag)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	c.sp = result
}
func (c *CPU) rl(operandType Op8Type, isCircular bool) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		b7 := (src >> 7) & 1

		value := (src << 1) & 0xff
		if isCircular {
			value |= b7
		} else {
			value |= ((uint8(*c.flags) >> 4) & 1)
		}

		c.writeOp8Value(opcode, operandType, value)
		c.flags.Unset(ZeroFlag)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.SetIfCondElseUnset(CarryFlag, b7 != 0)
	}
}
func (c *CPU) rr(operandType Op8Type, isCircular bool) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		b0 := src & 1

		value := src >> 1
		if isCircular {
			value |= b0 << 7
		} else {
			value |= ((uint8(*c.flags) >> 4) & 1) << 7
		}

		c.writeOp8Value(opcode, operandType, value)
		c.flags.Unset(ZeroFlag)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)
	}
}
func (c *CPU) sla(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		b7 := (src >> 7) & 1
		value := src << 1

		c.writeOp8Value(opcode, operandType, value)
		c.flags.SetIfCondElseUnset(ZeroFlag, value == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.SetIfCondElseUnset(CarryFlag, b7 != 0)
	}
}
func (c *CPU) sra(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		b0 := src & 1
		b7 := src & (1 << 7)
		value := b7 | (src >> 1)

		c.writeOp8Value(opcode, operandType, value)
		c.flags.SetIfCondElseUnset(ZeroFlag, value == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)
	}
}
func (c *CPU) swap(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		value := ((src & 0x0f) << 4) | ((src & 0xf0) >> 4)
		c.writeOp8Value(opcode, operandType, value)
	}
}
func (c *CPU) srl(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		src := c.getOp8Value(opcode, operandType)
		b0 := src & 1
		value := src >> 1

		c.writeOp8Value(opcode, operandType, value)
		c.flags.SetIfCondElseUnset(ZeroFlag, value == 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
		c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)
	}
}
func (c *CPU) bit(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		num := (opcode & 0x38) >> 3
		value := c.getOp8Value(opcode, operandType)

		c.flags.SetIfCondElseUnset(ZeroFlag, value&(1<<num) != 0)
		c.flags.Unset(SubtractFlag)
		c.flags.Unset(HalfCarryFlag)
	}
}
func (c *CPU) res(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		num := (opcode & 0x38) >> 3
		value := c.getOp8Value(opcode, operandType)
		bf := helpers.NewBitfield(value)
		bf.Unset(1 << num)

		c.writeOp8Value(opcode, operandType, uint8(*bf))
	}
}
func (c *CPU) set(operandType Op8Type) func(uint8) {
	return func(opcode uint8) {
		num := (opcode & 0x38) >> 3
		value := c.getOp8Value(opcode, operandType)
		bf := helpers.NewBitfield(value)
		bf.Set(1 << num)

		c.writeOp8Value(opcode, operandType, uint8(*bf))
	}
}

// utils
func (c *CPU) perform8BitArithmetic(operand1, operand2 uint8, toCarry, isSubtract bool) uint8 {
	carry := uint8(0)
	if toCarry && c.flags.IsSet(CarryFlag) {
		carry = uint8(1)
	}

	var sum uint16
	if isSubtract {
		sum = uint16(operand1) - uint16(operand2) - uint16(carry)
	} else {
		sum = uint16(operand1) + uint16(operand2) + uint16(carry)
	}
	result := uint8(sum)

	halfCarryFlag := false
	if isSubtract {
		halfCarryFlag = (operand1 & 0xf) < ((operand2 & 0xf) + carry)
	} else {
		halfCarryFlag = ((operand1 & 0xf) + (operand2 & 0xf) + carry) > 0xf
	}

	carryFlag := false
	if toCarry {
		if isSubtract {
			carryFlag = uint16(operand1) < uint16(operand2)
		} else {
			carryFlag = sum > 0xff
		}
	}

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.SetIfCondElseUnset(SubtractFlag, isSubtract)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	return result
}
func (c *CPU) perform16BitAddition(operand1, operand2 uint16, setZeroFlag bool) uint16 {
	sum := uint32(operand1) + uint32(operand2)
	result := uint16(sum)

	halfCarryFlag := ((operand1 & 0xfff) + (operand2 & 0xfff)) > 0xfff
	carryFlag := sum > 0xffff

	if setZeroFlag {
		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	}
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	return result
}
func (c *CPU) addToRegA(operand uint8, toCarry bool) {
	result := c.perform8BitArithmetic(c.a, operand, toCarry, false)
	c.a = result
}
func (c *CPU) subFromRegA(operand uint8, toCarry bool) {
	result := c.perform8BitArithmetic(c.a, operand, toCarry, true)
	c.a = result
}
func (c *CPU) getOp8Value(opcode uint8, operandType Op8Type) uint8 {
	switch operandType {
	case Op8R8Src:
		src := opcode & 0x07
		return *c.regs[src]
	case Op8R8Dest:
		dest := (opcode & 0x38) >> 3
		return *c.regs[dest]
	case Op8N8:
		return c.readNextByte()
	case Op8HL:
		return c.memory[c.readHL()]
	default:
		panic("unsupported `Op8Type` type for getOp8Value")
	}
}
func (c *CPU) writeOp8Value(opcode uint8, operandType Op8Type, value uint8) {
	switch operandType {
	case Op8R8Src:
		idx := opcode & 0x07
		*c.regs[idx] = value
	case Op8R8Dest:
		idx := (opcode & 0x38) >> 3
		*c.regs[idx] = value
	case Op8HL:
		c.memory[c.readHL()] = value
	default:
		panic("unsupported `Op8Type` type for writeOp8Value")
	}
}
