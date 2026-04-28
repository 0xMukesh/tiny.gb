package cpu

import (
	"github.com/0xmukesh/tiny.gb/internal/helpers"
)

const (
	ZeroFlag      = helpers.Bitfield(1 << 7)
	SubtractFlag  = helpers.Bitfield(1 << 6)
	HalfCarryFlag = helpers.Bitfield(1 << 5)
	CarryFlag     = helpers.Bitfield(1 << 4)
)

type CPU struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8

	flags  *helpers.Bitfield
	sp     uint16
	pc     uint16
	cycles uint

	memory []uint8
	prg    []uint8

	ime      bool
	isHalted bool
}

func NewCPU(prg []uint8) *CPU {
	c := &CPU{}
	return c
}

func (c *CPU) Step() {
	opcode := c.readNextByte()
	if opcode == 0xcb {
		opcode = c.readNextByte()
		c.executeCb(opcode)
	} else {
		c.execute(opcode)
	}
}

func (c *CPU) IsHalted() bool {
	return c.isHalted
}

func (c *CPU) readR8(idx uint8) uint8 {
	switch idx {
	case 0:
		return c.b
	case 1:
		return c.c
	case 2:
		return c.d
	case 3:
		return c.e
	case 4:
		return c.h
	case 5:
		return c.l
	case 6:
		return c.memory[c.readHL()]
	case 7:
		return c.a
	default:
		panic("not possible")
	}
}

func (c *CPU) writeR8(idx, data uint8) {
	switch idx {
	case 0:
		c.b = data
	case 1:
		c.c = data
	case 2:
		c.d = data
	case 3:
		c.e = data
	case 4:
		c.h = data
	case 5:
		c.l = data
	case 6:
		c.memory[c.readHL()] = data
	case 7:
		c.a = data
	}
}

func (c *CPU) readR16(idx uint8, useSpOverAf bool) uint16 {
	switch idx {
	case 0:
		return c.readBC()
	case 1:
		return c.readDE()
	case 2:
		return c.readHL()
	case 3:
		if useSpOverAf {
			return c.sp
		} else {
			return c.readAF()
		}
	default:
		panic("not possible")
	}
}

func (c *CPU) writeR16(idx uint8, data uint16, useSpOverAf bool) {
	switch idx {
	case 0:
		c.writeBC(data)
	case 1:
		c.writeDE(data)
	case 2:
		c.writeHL(data)
	case 3:
		if useSpOverAf {
			c.sp = data
		} else {
			c.writeAF(data)
		}
	}
}

func (c *CPU) readNextByte() uint8 {
	value := c.memory[c.pc]
	c.pc++
	return value
}
func (c *CPU) readNextU16() uint16 {
	lo := c.readNextByte()
	hi := c.readNextByte()
	return (uint16(hi) << 8) | uint16(lo)
}

func (c *CPU) stackPush(value uint8) {
	c.memory[c.sp] = value
	c.sp--
}
func (c *CPU) stackPop() uint8 {
	c.sp++
	return c.memory[c.sp]
}
func (c *CPU) stackPushU16(value uint16) {
	hi := uint8(value >> 8)
	lo := uint8(value & 0xff)

	c.stackPush(hi)
	c.stackPush(lo)
}
func (c *CPU) stackPopU16() uint16 {
	hi := c.stackPop()
	lo := c.stackPop()

	return uint16(hi)<<8 | uint16(lo)
}

func (c *CPU) readBC() uint16 {
	return (uint16(c.b) << 8) | uint16(c.c)
}
func (c *CPU) writeBC(value uint16) {
	c.b = uint8(value >> 8)
	c.c = uint8(value & 0xff)
}

func (c *CPU) readDE() uint16 {
	return (uint16(c.d) << 8) | uint16(c.e)
}
func (c *CPU) writeDE(value uint16) {
	c.d = uint8(value >> 8)
	c.e = uint8(value & 0xff)
}

func (c *CPU) readHL() uint16 {
	return (uint16(c.h) << 8) | uint16(c.l)
}
func (c *CPU) writeHL(value uint16) {
	c.h = uint8(value >> 8)
	c.l = uint8(value & 0xff)
}

func (c *CPU) readAF() uint16 {
	return (uint16(c.a) << 8) | uint16(*c.flags)
}
func (c *CPU) writeAF(value uint16) {
	c.a = uint8(value >> 8)
	*c.flags = helpers.Bitfield(value & 0xff)
}
