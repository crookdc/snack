package alu

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

// HalfAdder is a so-called half alu that accepts two bits as input and produces an output consisting of a sum and a carry.
func HalfAdder(a, b snack.Bit) (carry snack.Bit, sum snack.Bit) {
	return gate.AndBit(a, b), gate.XorBit(a, b)
}
