package alu

import (
	"github.com/crookdc/snack/internal/gate"
	"github.com/crookdc/snack/internal/pin"
)

// HalfAdder accepts two bits as input and produces an output consisting of a sum and a carry resulting from adding the
// provided bits together.
func HalfAdder(a, b pin.Signal) (carry pin.Signal, sum pin.Signal) {
	return gate.And(a, b), gate.Xor(a, b)
}

// FullAdder accepts three bits and produces a carry and a sum bit representing the result of adding the three bits
// together.
func FullAdder(a, b, c pin.Signal) (carry pin.Signal, sum pin.Signal) {
	ac, sum := HalfAdder(a, b)
	bc, sum := HalfAdder(sum, c)
	return gate.Or(ac, bc), sum
}

// Adder16 adds two 16-bit integers and returns the result. The carry bit is ignored by the adder.
func Adder16(a, b uint16) uint16 {
	ab := pin.Split16(a)
	bb := pin.Split16(b)

	r := [16]pin.Signal{}
	c, s := pin.Inactive, pin.Inactive
	for i := range 16 {
		c, s = FullAdder(ab[15-i], bb[15-i], c)
		r[15-i] = s
	}
	return pin.Join16(r)
}
