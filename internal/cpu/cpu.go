package cpu

import (
	"fmt"

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

	flags *helpers.Bitfield
	sp    uint16
	pc    uint16
	ticks uint

	memory []uint8
	prg    []uint8

	instructions           map[uint8]Instruction
	cbPrefixedInstructions map[uint8]Instruction // prefixed with $cb
	regs                   []*uint8

	isHalted bool
}

func NewCPU(prg []uint8) *CPU {
	c := &CPU{}
	c.regs = []*uint8{&c.b, &c.c, &c.d, &c.e, &c.h, &c.l, nil, &c.a}
	c.buildInstructions()
	c.buildCbPrefixedInstructions()
	return c
}

func (c *CPU) Step() {
	opcode := c.prg[c.pc]

	var instr Instruction
	var ok bool
	if opcode == 0xcb {
		instr, ok = c.cbPrefixedInstructions[opcode]
	} else {
		instr, ok = c.instructions[opcode]
	}

	if !ok {
		panic(fmt.Errorf("invalid instruction: 0x%04x", opcode))
	}

	instr.handler(opcode)
	c.ticks += instr.ticks
	c.pc++
}

func (c *CPU) IsHalted() bool {
	return c.isHalted
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
