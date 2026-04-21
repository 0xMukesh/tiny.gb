package helpers

type Bitfield uint8

func NewBitfield(n uint8) *Bitfield {
	bf := Bitfield(n)
	return &bf
}

func (bf *Bitfield) Set(flag Bitfield) {
	*bf |= flag
}

func (bf *Bitfield) Unset(flag Bitfield) {
	*bf &^= flag
}

func (bf *Bitfield) Toggle(flag Bitfield) {
	*bf ^= flag
}

func (bf *Bitfield) IsSet(flag Bitfield) bool {
	return (*bf & flag) != 0
}

func (b *Bitfield) SetIfCondElseUnset(flag Bitfield, cond bool) {
	if cond {
		b.Set(flag)
	} else {
		b.Unset(flag)
	}
}
