package cpu

import "fmt"

type Bit8ArithmeticOperandType int

const (
	Bit8R8FirstThree  = iota // 0bxy|abc|(rrr)
	Bit8R8MiddleThree        // 0bxy|(rrr)|abc
	Bit8HL
	Bit8N8
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
			case p == 0 && q == 1:
				c.ld_a16_sp()
			}
		case 1:
			switch {
			case q == 0:
				c.ld_r16_n16(opcode)
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
		case 4:
			switch y {
			case 6:
				c.inc_dec(opcode, Bit8HL, false)
			default:
				c.inc_dec(opcode, Bit8R8MiddleThree, false)
			}
		case 5:
			switch y {
			case 6:
				c.inc_dec(opcode, Bit8HL, true)
			default:
				c.inc_dec(opcode, Bit8R8MiddleThree, true)
			}
		case 6:
			switch y {
			case 6:
				c.ld_hl_n8()
			default:
				c.ld_r8_n8(opcode)
			}
		case 7:
			switch y {
			case 4:
				c.daa() // Decimal adjust accumulator
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
			case 4:
				c.ldh_n_a()
			case 6:
				c.ldh_a_n()
			case 7:
				c.ld_hl_sp_e()
			}
		case 1:
			switch {
			case q == 0:
				c.pop_r16(opcode)
			case y == 7:
				c.ld_sp_hl()
			}
		case 2:
			switch y {
			case 4:
				c.ldh_c_a()
			case 5:
				c.ld_nn_a()
			case 6:
				c.ldh_a_c()
			case 7:
				c.ld_a_nn()
			}
		case 5:
			switch {
			case q == 0:
				c.push_r16(opcode)
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
		}
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
}

func (c *CPU) add_sub(opcode uint8, operandType Bit8ArithmeticOperandType, isCarry, isSubtraction bool) {
	value := c.readBit8OperandType(opcode, operandType)
	result := c.perform8BitArithmetic(c.a, value, isCarry, isSubtraction)
	c.a = result
}

func (c *CPU) cp(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.perform8BitArithmetic(c.a, value, false, true)
}

func (c *CPU) inc_dec(opcode uint8, operandType Bit8ArithmeticOperandType, isDec bool) {
	value := c.readBit8OperandType(opcode, operandType)
	beforeCarryFlag := c.flags.IsSet(CarryFlag)
	result := c.perform8BitArithmetic(value, 1, false, isDec)
	c.writeBit8OperandType(opcode, operandType, result)

	c.flags.SetIfCondElseUnset(CarryFlag, beforeCarryFlag)
}

func (c *CPU) and(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a &= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Set(HalfCarryFlag)
	c.flags.Unset(CarryFlag)
}

func (c *CPU) or(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a |= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Unset(CarryFlag)
}

func (c *CPU) xor(opcode uint8, operandType Bit8ArithmeticOperandType) {
	value := c.readBit8OperandType(opcode, operandType)
	c.a ^= value

	c.flags.SetIfCondElseUnset(ZeroFlag, c.a == 0)
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Unset(CarryFlag)
}

func (c *CPU) ccf() {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Toggle(CarryFlag)
}

func (c *CPU) scf() {
	c.flags.Unset(SubtractFlag)
	c.flags.Unset(HalfCarryFlag)
	c.flags.Set(CarryFlag)
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
}

func (c *CPU) cpl() {
	c.a = ^c.a
	c.flags.Set(SubtractFlag)
	c.flags.Set(HalfCarryFlag)
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
	default:
		panic(fmt.Sprintf("unsupported bit8 operand type for write: %d", operandType))
	}
}
