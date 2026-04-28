package cpu

import (
	"fmt"

	"github.com/0xmukesh/tiny.gb/internal/helpers"
)

type Bit8ArithmeticOperandType int

const (
	Bit8R8FirstThree  Bit8ArithmeticOperandType = iota // 0bxy|abc|(rrr)
	Bit8R8MiddleThree                                  // 0bxy|(rrr)|abc
	Bit8HL
	Bit8N8
	Bit8A
)

// opcode = [ x x | y y y | z z z ]
func (c *CPU) execute(opcode uint8) {
	x := (opcode >> 6) & 0x3
	y := (opcode >> 3) & 0x7
	z := opcode & 0x7
	p := y >> 1
	q := y & 0x1

	switch x {
	case 0:
		switch z {
		case 0:
			switch {
			case p == 0 && q == 0:
				c.nop()
			case p == 0 && q == 1:
				c.ld_a16_sp()
			case p == 1 && q == 1:
				c.jr_e(true)
			case (y>>2)&0x1 == 1:
				c.jr_e(c.decodeConditionFromOpcode(opcode))
			}
		case 1:
			switch q {
			case 0:
				c.ld_r16_n16(opcode)
			case 1:
				c.add_hl_r16(opcode)
			}
		case 2:
			switch y {
			case 0:
				c.ld_bc_a()
			case 1:
				c.ld_a_bc()
			case 2:
				c.ld_de_a()
			case 3:
				c.ld_a_de()
			case 4:
				c.ld_hli_a()
			case 5:
				c.ld_a_hli()
			case 6:
				c.ld_hld_a()
			case 7:
				c.ld_a_hld()
			}
		case 3:
			switch q {
			case 0:
				c.inc_dec_r16(opcode, false)
			case 1:
				c.inc_dec_r16(opcode, true)
			}
		case 4:
			switch y {
			case 6:
				c.inc_dec_bit8(opcode, Bit8HL, false)
			default:
				c.inc_dec_bit8(opcode, Bit8R8MiddleThree, false)
			}
		case 5:
			switch y {
			case 6:
				c.inc_dec_bit8(opcode, Bit8HL, true)
			default:
				c.inc_dec_bit8(opcode, Bit8R8MiddleThree, true)
			}
		case 6:
			c.ld_r8_n8(opcode)
		case 7:
			switch y {
			case 0:
				c.rlc(opcode, Bit8A)
			case 1:
				c.rrc(opcode, Bit8A)
			case 2:
				c.rl(opcode, Bit8A)
			case 3:
				c.rr(opcode, Bit8A)
			case 4:
				c.daa()
			case 5:
				c.cpl()
			case 6:
				c.scf()
			case 7:
				c.ccf()
			}
		}
	case 1:
		switch {
		case y != 6 && z != 6:
			c.ld_r8_r8(opcode)
		case y != 6 && z == 6:
			c.ld_r8_hl(opcode)
		case y == 6 && z != 6:
			c.ld_hl_r8(opcode)
		case y == 6 && z == 6:
			c.halt()
		}
	case 2:
		switch y {
		case 0:
			switch z {
			case 6:
				c.add_sub(opcode, Bit8HL, false, false)
			default:
				c.add_sub(opcode, Bit8R8FirstThree, false, false)
			}
		case 1:
			switch z {
			case 6:
				c.add_sub(opcode, Bit8HL, true, false)
			default:
				c.add_sub(opcode, Bit8R8FirstThree, true, false)
			}
		case 2:
			switch z {
			case 6:
				c.add_sub(opcode, Bit8HL, false, true)
			default:
				c.add_sub(opcode, Bit8R8FirstThree, false, true)
			}
		case 3:
			switch z {
			case 6:
				c.add_sub(opcode, Bit8HL, true, true)
			default:
				c.add_sub(opcode, Bit8R8FirstThree, true, true)
			}
		case 4:
			switch z {
			case 6:
				c.and(opcode, Bit8HL)
			default:
				c.and(opcode, Bit8R8FirstThree)
			}
		case 5:
			switch z {
			case 6:
				c.xor(opcode, Bit8HL)
			default:
				c.xor(opcode, Bit8R8FirstThree)
			}
		case 6:
			switch z {
			case 6:
				c.or(opcode, Bit8HL)
			default:
				c.or(opcode, Bit8R8FirstThree)
			}
		case 7:
			switch z {
			case 6:
				c.cp(opcode, Bit8HL)
			default:
				c.cp(opcode, Bit8R8FirstThree)
			}
		}
	case 3:
		switch z {
		case 0:
			switch y {
			case 0, 1, 2, 3:
				c.ret(true, c.decodeConditionFromOpcode(opcode))
			case 4:
				c.ldh_n_a()
			case 5:
				c.add_sp_e()
			case 6:
				c.ldh_a_n()
			case 7:
				c.ld_hl_sp_e()
			}
		case 1:
			switch {
			case q == 0:
				c.pop_r16(opcode)
			case q == 1:
				c.ret(false, true)
			case y == 3:
				c.reti()
			case y == 5:
				c.jp_hl()
			case y == 7:
				c.ld_sp_hl()
			}
		case 2:
			switch y {
			case 0, 1, 2, 3:
				c.jp_n16(c.decodeConditionFromOpcode(opcode))
			case 4:
				c.ldh_c_a()
			case 5:
				c.ld_nn_a()
			case 6:
				c.ldh_a_c()
			case 7:
				c.ld_a_nn()
			}
		case 3:
			switch y {
			case 0:
				c.jp_n16(true)
			case 6:
				c.di()
			case 7:
				c.ei()
			}
		case 4:
			switch {
			case (y>>2)&0x1 == 0:
				c.call_n16(c.decodeConditionFromOpcode(opcode))
			}
		case 5:
			switch q {
			case 0:
				c.push_r16(opcode)
			case 1:
				c.call_n16(true)
			}
		case 6:
			switch y {
			case 0:
				c.add_sub(opcode, Bit8N8, false, false)
			case 1:
				c.add_sub(opcode, Bit8N8, true, false)
			case 2:
				c.add_sub(opcode, Bit8N8, false, true)
			case 3:
				c.add_sub(opcode, Bit8N8, true, true)
			case 4:
				c.and(opcode, Bit8N8)
			case 5:
				c.xor(opcode, Bit8N8)
			case 6:
				c.or(opcode, Bit8N8)
			case 7:
				c.cp(opcode, Bit8N8)
			}
		case 7:
			c.rst(opcode)
		}
	}
}

func (c *CPU) executeCb(opcode uint8) {
	x := (opcode >> 6) & 0x3
	y := (opcode >> 3) & 0x7
	z := opcode & 0x7

	operandType := Bit8R8FirstThree
	if z == 6 {
		operandType = Bit8HL
	}

	switch x {
	case 0:
		switch y {
		case 0:
			c.rlc(opcode, operandType)
		case 1:
			c.rrc(opcode, operandType)
		case 2:
			c.rl(opcode, operandType)
		case 3:
			c.rr(opcode, operandType)
		case 4:
			c.sla(opcode, operandType)
		case 5:
			c.sra(opcode, operandType)
		case 6:
			c.swap(opcode, operandType)
		case 7:
			c.srl(opcode, operandType)
		}
	case 1:
		c.bit(opcode, operandType)
	case 2:
		c.res(opcode, operandType)
	case 3:
		c.set(opcode, operandType)
	}
}

func (c *CPU) ld_r8_r8(opcode uint8) {
	rx := (opcode >> 3) & 0x7
	ry := opcode & 0x7
	c.writeR8(rx, c.readR8(ry))
	c.cycles += 1
}

func (c *CPU) ld_r8_n8(opcode uint8) {
	r := (opcode >> 3) & 0x7
	c.writeR8(r, c.readNextByte())
	c.cycles += 2
}

func (c *CPU) ld_r8_hl(opcode uint8) {
	r := (opcode >> 3) & 0x7
	c.writeR8(r, c.memory[c.readHL()])
	c.cycles += 2
}

func (c *CPU) ld_hl_r8(opcode uint8) {
	r := opcode & 0x7
	c.memory[c.readHL()] = c.readR8(r)
	c.cycles += 2
}

func (c *CPU) ld_hl_n8() {
	c.memory[c.readHL()] = c.readNextByte()
	c.cycles += 3
}

func (c *CPU) ld_a_bc() {
	c.a = c.memory[c.readBC()]
	c.cycles += 2
}

func (c *CPU) ld_a_de() {
	c.a = c.memory[c.readDE()]
	c.cycles += 2
}

func (c *CPU) ld_bc_a() {
	c.memory[c.readBC()] = c.a
	c.cycles += 2
}

func (c *CPU) ld_de_a() {
	c.memory[c.readDE()] = c.a
	c.cycles += 2
}

func (c *CPU) ld_a_nn() {
	c.a = c.memory[c.readNextU16()]
	c.cycles += 4
}

func (c *CPU) ld_nn_a() {
	c.memory[c.readNextU16()] = c.a
	c.cycles += 4
}

func (c *CPU) ldh_a_c() {
	addr := 0xff00 + uint16(c.c)
	c.a = c.memory[addr]
	c.cycles += 2
}

func (c *CPU) ldh_c_a() {
	addr := 0xff00 + uint16(c.c)
	c.memory[addr] = c.a
	c.cycles += 2
}

func (c *CPU) ldh_a_n() {
	addr := 0xff00 + uint16(c.readNextByte())
	c.a = c.memory[addr]
	c.cycles += 2
}

func (c *CPU) ldh_n_a() {
	addr := 0xff00 + uint16(c.readNextByte())
	c.memory[addr] = c.a
	c.cycles += 2
}

func (c *CPU) ld_a_hld() {
	hl := c.readHL()
	c.a = c.memory[hl]
	c.writeHL(hl - 1)
	c.cycles += 2
}

func (c *CPU) ld_hld_a() {
	hl := c.readHL()
	c.memory[hl] = c.a
	c.writeHL(hl - 1)
	c.cycles += 2
}

func (c *CPU) ld_a_hli() {
	hl := c.readHL()
	c.a = c.memory[hl]
	c.writeHL(hl + 1)
	c.cycles += 2
}

func (c *CPU) ld_hli_a() {
	hl := c.readHL()
	c.memory[hl] = c.a
	c.writeHL(hl + 1)
	c.cycles += 2
}

func (c *CPU) ld_r16_n16(opcode uint8) {
	rr := (opcode >> 4) & 0x3
	c.writeR16(rr, c.readNextU16(), true)
	c.cycles += 3
}

func (c *CPU) ld_a16_sp() {
	addr := c.readNextU16()
	c.memory[addr] = uint8(c.sp & 0xff)
	c.memory[addr+1] = uint8((c.sp & 0xff00) >> 8)
	c.cycles += 5
}

func (c *CPU) ld_sp_hl() {
	c.sp = c.readHL()
	c.cycles += 2
}

func (c *CPU) push_r16(opcode uint8) {
	rr := (opcode >> 4) & 0x3
	c.stackPush(rr)
	c.cycles += 4
}

func (c *CPU) pop_r16(opcode uint8) {
	rr := (opcode >> 4) & 0x3
	data := c.stackPopU16()
	c.writeR16(rr, data, false)
	c.cycles += 3
}

func (c *CPU) ld_hl_sp_e() {
	e := int8(c.readNextByte())
	result := c.sp + uint16(int16(e))

	halfCarryFlag := (c.sp&0xf)+uint16(uint8(e)&0xf) > 0xf
	carryFlag := (c.sp&0xff)+uint16(uint8(e)) > 0xff

	c.writeHL(result)

	c.flags.Unset(ZeroFlag)
	c.flags.Unset(SubtractFlag)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)
	c.cycles += 3
}

func (c *CPU) add_sub(opcode uint8, operandType Bit8ArithmeticOperandType, isCarry, isSubtraction bool) {
	value := c.readBit8OperandType(opcode, operandType)
	result := c.perform8BitArithmetic(c.a, value, isCarry, isSubtraction)
	c.a = result

	if operandType == Bit8HL || operandType == Bit8N8 {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) cp(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.perform8BitArithmetic(c.a, value, false, true)

	if operandType == Bit8HL || operandType == Bit8N8 {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) inc_dec_bit8(opcode uint8, operandType Bit8ArithmeticOperandType, isDec bool) {
	value := c.readBit8OperandType(opcode, operandType)
	beforeCarryFlag := c.flags.IsSet(CarryFlag)
	result := c.perform8BitArithmetic(value, 1, false, isDec)
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(CarryFlag, beforeCarryFlag)

	if operandType == Bit8HL {
		c.cycles += 3
	} else {
		c.cycles += 1
	}
}

func (c *CPU) and(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a &= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Set(HalfCarryFlag)
	c.flags.Unset(CarryFlag)

	if operandType == Bit8HL || operandType == Bit8N8 {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) or(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a |= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Unset(CarryFlag)

	if operandType == Bit8HL || operandType == Bit8N8 {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) xor(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a ^= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Unset(CarryFlag)

	if operandType == Bit8HL || operandType == Bit8N8 {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) ccf() {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Toggle(CarryFlag)

	c.cycles += 1
}

func (c *CPU) scf() {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Set(CarryFlag)

	c.cycles += 1
}

// https://rgbds.gbdev.io/docs/v1.0.1/gbz80.7#DAA
func (c *CPU) daa() {
	adjustment := uint8(0)

	if c.flags.IsSet(SubtractFlag) {
		if c.flags.IsSet(HalfCarryFlag) {
			adjustment += 0x6
		}

		if c.flags.IsSet(CarryFlag) {
			adjustment += 0x60
		}

		c.a -= adjustment
	} else {
		if c.flags.IsSet(HalfCarryFlag) || (c.a&0xf > 0x9) {
			adjustment += 0x6
		}

		if c.flags.IsSet(CarryFlag) || c.a > 0x99 {
			adjustment += 0x60
			c.flags.Set(CarryFlag)
		}

		c.a += adjustment
	}

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(HalfCarryFlag)

	c.cycles += 1
}

func (c *CPU) cpl() {
	c.a = ^c.a
	c.flags.Set(SubtractFlag)
	c.flags.Set(HalfCarryFlag)

	c.cycles += 1
}

func (c *CPU) inc_dec_r16(opcode uint8, isDec bool) {
	rr := (opcode >> 4) & 0x3
	value := c.readR16(rr, true)

	data := value
	if isDec {
		data -= 1
	} else {
		data += 1
	}

	c.writeR16(rr, data, true)
	c.cycles += 2
}

func (c *CPU) add_hl_r16(opcode uint8) {
	rr := (opcode >> 4) & 0x3
	value := c.readR16((opcode>>4)&0x3, true)
	hl := c.readHL()

	sum := uint32(hl) + uint32(value)
	result := uint16(sum)

	halfCarryFlag := (hl&0xfff)+(value&0xfff) > 0xfff
	carryFlag := sum > 0xffff

	c.writeR16(rr, result, true)
	c.flags.Unset(SubtractFlag)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	c.cycles += 2
}

func (c *CPU) add_sp_e() {
	e := int8(c.readNextByte())
	result := c.sp + uint16(int16(e))

	halfCarryFlag := (c.sp&0xf)+uint16(uint8(e)&0xf) > 0xf
	carryFlag := (result & 0xff) < (c.sp & 0xff)

	c.sp = result
	c.flags.Unset(ZeroFlag)
	c.flags.Unset(SubtractFlag)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	c.cycles += 4
}

func (c *CPU) rlc(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b7 := (value >> 7) & 0x1

	result := (value << 1) | b7
	c.writeBit8OperandType(opcode, operandType, result)

	if operandType == Bit8A {
		c.flags.Unset(ZeroFlag)
	} else {
		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	}
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b7 != 0)

	switch operandType {
	case Bit8A:
		c.cycles += 1
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) rrc(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b0 := value & 0x1

	result := (value >> 1) | (b0 << 7)
	c.writeBit8OperandType(opcode, operandType, result)

	if operandType == Bit8A {
		c.flags.Unset(ZeroFlag)
	} else {
		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	}
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)

	switch operandType {
	case Bit8A:
		c.cycles += 1
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) rl(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b7 := (value >> 7) & 0x1
	cf := uint8(0)
	if c.flags.IsSet(CarryFlag) {
		cf = 1
	}

	result := (value << 1) | cf
	c.writeBit8OperandType(opcode, operandType, result)

	if operandType == Bit8A {
		c.flags.Unset(ZeroFlag)
	} else {
		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	}
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b7 != 0)

	switch operandType {
	case Bit8A:
		c.cycles += 1
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) rr(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b0 := value & 0x1
	cf := uint8(0)
	if c.flags.IsSet(CarryFlag) {
		cf = 1
	}

	result := (value >> 1) | (cf << 7)
	c.writeBit8OperandType(opcode, operandType, result)

	if operandType == Bit8A {
		c.flags.Unset(ZeroFlag)
	} else {
		c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	}
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)

	switch operandType {
	case Bit8A:
		c.cycles += 1
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) sla(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b7 := (value >> 7) & 0x1

	result := value << 1
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b7 != 0)

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) sra(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b7 := value & (0x1 << 7)
	b0 := value & 0x1

	result := (value >> 1) | b7
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) swap(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	lower := value & 0x0f
	upper := (value & 0xf0) >> 4

	result := (lower << 4) | upper
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Unset(CarryFlag)

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) srl(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	b0 := value & 0x1

	result := value >> 1
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, b0 != 0)

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) bit(opcode uint8, operandType Bit8ArithmeticOperandType) {
	bit := c.readBit8OperandType(opcode, Bit8R8MiddleThree)
	value := c.readBit8OperandType(opcode, operandType)

	bf := helpers.NewBitfield(value)

	c.flags.SetIfCondElseUnset(ZeroFlag, !bf.IsSet(1<<bit))
	c.flags.Unset(SubtractFlag)
	c.flags.Set(HalfCarryFlag)

	switch operandType {
	case Bit8HL:
		c.cycles += 3
	default:
		c.cycles += 2
	}
}

func (c *CPU) res(opcode uint8, operandType Bit8ArithmeticOperandType) {
	bit := c.readBit8OperandType(opcode, Bit8R8MiddleThree)
	value := c.readBit8OperandType(opcode, operandType)

	bf := helpers.NewBitfield(value)
	bf.Unset(1 << bit)

	c.writeBit8OperandType(opcode, operandType, uint8(*bf))

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) set(opcode uint8, operandType Bit8ArithmeticOperandType) {
	bit := c.readBit8OperandType(opcode, Bit8R8MiddleThree)
	value := c.readBit8OperandType(opcode, operandType)

	bf := helpers.NewBitfield(value)
	bf.Set(1 << bit)

	c.writeBit8OperandType(opcode, operandType, uint8(*bf))

	switch operandType {
	case Bit8HL:
		c.cycles += 4
	default:
		c.cycles += 2
	}
}

func (c *CPU) jp_n16(condition bool) {
	if condition {
		nn := c.readNextU16()
		c.pc = nn
		c.cycles += 1
	}

	c.cycles += 3
}

func (c *CPU) jp_hl() {
	c.pc = c.readHL()
	c.cycles += 1
}

func (c *CPU) jr_e(condition bool) {
	if condition {
		e := int(c.readNextByte())
		c.pc = uint16(int(c.pc) + e)
		c.cycles += 1
	}

	c.cycles += 2
}

func (c *CPU) call_n16(condition bool) {
	if condition {
		nn := c.readNextU16()
		c.stackPushU16(nn)
		c.pc = nn
		c.cycles += 3
	}

	c.cycles += 3
}

func (c *CPU) ret(isConditional, condition bool) {
	if condition {
		nn := c.stackPopU16()
		c.pc = nn
		c.cycles += 3
	}

	if isConditional {
		c.cycles += 2
	} else {
		c.cycles += 1
	}
}

func (c *CPU) reti() {
	nn := c.stackPopU16()
	c.pc = nn
	c.ime = true
	c.cycles += 4
}

func (c *CPU) rst(opcode uint8) {
	idx := (opcode >> 3) & 0x7
	addr := idx * 8
	c.stackPushU16(c.pc)
	c.pc = uint16(addr)
	c.cycles += 4
}

func (c *CPU) di() {
	c.ime = false
	c.cycles += 1
}

func (c *CPU) ei() {
	c.ime = true
	c.cycles += 1
}

func (c *CPU) nop() {
	c.cycles += 1
}

func (c *CPU) halt() {
	c.isHalted = true
	c.cycles += 1
}

func (c *CPU) perform8BitArithmetic(operand1, operand2 uint8, toCarry, isSubtraction bool) uint8 {
	carry := uint8(0)
	if toCarry && c.flags.IsSet(CarryFlag) {
		carry = uint8(1)
	}

	var sum uint16
	if isSubtraction {
		sum = uint16(operand1) - uint16(operand2)
		if toCarry {
			sum -= uint16(carry)
		}
	} else {
		sum = uint16(operand1) + uint16(operand2)
		if toCarry {
			sum += uint16(carry)
		}
	}
	result := uint8(sum)

	halfCarryFlag := false
	if isSubtraction {
		halfCarryFlag = (operand1 & 0xf) < (operand2&0xf)+carry
	} else {
		halfCarryFlag = (operand1&0xf)+(operand2&0xf)+carry > 0xf
	}

	carryFlag := false
	if isSubtraction {
		carryFlag = operand1 < operand2
	} else {
		carryFlag = sum > 0xff
	}

	c.flags.SetIfCondElseUnset(ZeroFlag, result == 0)
	c.flags.SetIfCondElseUnset(SubtractFlag, isSubtraction)
	c.flags.SetIfCondElseUnset(HalfCarryFlag, halfCarryFlag)
	c.flags.SetIfCondElseUnset(CarryFlag, carryFlag)

	return result
}

func (c *CPU) readBit8OperandType(opcode uint8, operandType Bit8ArithmeticOperandType) uint8 {
	switch operandType {
	case Bit8R8FirstThree:
		return c.readR8(opcode & 0x7)
	case Bit8R8MiddleThree:
		return c.readR8((opcode >> 3) & 0x7)
	case Bit8HL:
		return c.memory[c.readHL()]
	case Bit8N8:
		return c.readNextByte()
	case Bit8A:
		return c.a
	default:
		panic(fmt.Sprintf("unsupported bit8 operand type for read: %d", operandType))
	}
}

func (c *CPU) writeBit8OperandType(opcode uint8, operandType Bit8ArithmeticOperandType, value uint8) {
	switch operandType {
	case Bit8R8FirstThree:
		c.writeR8(opcode&0x7, value)
	case Bit8R8MiddleThree:
		c.writeR8((opcode>>3)&0x7, value)
	case Bit8HL:
		c.memory[c.readHL()] = value
	case Bit8A:
		c.a = value
	default:
		panic(fmt.Sprintf("unsupported bit8 operand type for write: %d", operandType))
	}
}

func (c *CPU) decodeConditionFromOpcode(opcode uint8) bool {
	cc := (opcode >> 3) & 0x3

	switch cc {
	case 0: // NZ
		return !c.flags.IsSet(ZeroFlag)
	case 1: // Z
		return c.flags.IsSet(ZeroFlag)
	case 2: // NC
		return !c.flags.IsSet(CarryFlag)
	case 3: // C
		return c.flags.IsSet(CarryFlag)
	default:
		panic("not possible")
	}
}
