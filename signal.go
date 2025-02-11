package snack

import (
	"fmt"
)

type Signal struct {
	n uint8
}

// SignalSplit16 transforms a 16-bit integer to its bit representation in the Signal abstraction. The index of the slice is
// the position of the bit, meaning that in the 0th index you will find the 2^0 position bit, and on the 15th index you
// will find the 2^15 positioned bit.
func SignalSplit16(n uint16) []Signal {
	res := make([]Signal, 16)
	for i := range 16 {
		res[i] = NewSignal(uint8(n>>(15-i)) & 1)
	}
	return res
}

// SignalJoin16 transforms a Signal slice of length 16 to a 16-bit integer
func SignalJoin16(n []Signal) uint16 {
	if len(n) != 16 {
		panic(fmt.Errorf("invalid bit join length %d", len(n)))
	}
	res := uint16(0)
	for i := range 16 {
		res = res | (uint16(n[i].Bin()) << (15 - i))
	}
	return res
}

// Expand16 takes a single bit and expands its value to cover 16 bits. That is, if the bit value is 0 then a 16-bit
// unsigned integer containing all zeroes is returned. If the input bit value is 1 then an unsigned 16-bit integer
// containing all ones is returned.
func Expand16(n Signal) uint16 {
	if n.IsActive() {
		return 0xFFFF
	}
	return 0
}

func NewSignal(n uint8) Signal {
	if n == 0 {
		return InactiveSignal()
	} else if n == 1 {
		return ActiveSignal()
	}
	panic(fmt.Errorf("invalid bit value %d", n))
}

func ActiveSignal() Signal {
	return Signal{n: 1}
}

func InactiveSignal() Signal {
	return Signal{n: 0}
}

func (b *Signal) Mask() uint8 {
	if b.IsActive() {
		return 0xFF
	}
	return 0
}

func (b *Signal) Activate() {
	b.n = 1
}

func (b *Signal) Deactivate() {
	b.n = 0
}

func (b *Signal) Flip() {
	if b.n == 0 {
		b.n = 1
	} else {
		b.n = 0
	}
}

func (b *Signal) IsActive() bool {
	return b.n == 1
}

func (b *Signal) Bin() uint8 {
	return b.n
}
