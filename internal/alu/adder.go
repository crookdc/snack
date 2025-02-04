package alu

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

// HalfAdder accepts two bits as input and produces an output consisting of a sum and a carry resulting from adding the
// provided bits together.
func HalfAdder(a, b snack.Bit) (carry snack.Bit, sum snack.Bit) {
	return gate.AndBit(a, b), gate.XorBit(a, b)
}

// FullAdder accepts three bits and produces a carry and a sum bit representing the result of adding the three bits
// together.
func FullAdder(a, b, c snack.Bit) (carry snack.Bit, sum snack.Bit) {
	ac, sum := HalfAdder(a, b)
	bc, sum := HalfAdder(sum, c)
	return gate.OrBit(ac, bc), sum
}

// Adder16 adds two 16-bit integers and returns the result. The carry bit is ignored by the adder.
func Adder16(a, b uint16) uint16 {
	ab := snack.BitSplit16(a)
	bb := snack.BitSplit16(b)

	r := make([]snack.Bit, 16)
	c, s := snack.UnsetBit(), snack.UnsetBit()
	for i := range 16 {
		c, s = FullAdder(ab[15-i], bb[15-i], c)
		r[15-i] = s
	}
	return snack.BitJoin16(r)
}
