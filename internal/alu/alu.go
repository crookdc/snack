package alu

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

// ALU (Arithmetic Logic Unit) represents an actual hardware chip that can handle an array of arithmetic operations
// which are executed on two 16-bit unsigned integers and which produces a single 16-bit unsigned integer as result.
type ALU struct {
	// ZX sets all bits of x to 0
	ZX snack.Bit
	// NX negates all bits of x
	NX snack.Bit
	// ZY sets all bits of y to 0
	ZY snack.Bit
	// NY negates all bits of y
	NY snack.Bit
	// F when set causes ALU operator to be a bitwise AND, otherwise operator is addition
	F snack.Bit
	// NO negates all bits of output
	NO snack.Bit
}

// Call performs operations on the provided inputs as outlined by the state of the ALU
func (a *ALU) Call(x, y uint16) uint16 {
	x = gate.AndUint16(x, snack.Expand16(gate.NotBit(a.ZX)))
	x = gate.XorUint16(x, snack.Expand16(a.NX))

	y = gate.AndUint16(y, snack.Expand16(gate.NotBit(a.ZY)))
	y = gate.XorUint16(y, snack.Expand16(a.NY))

	out := gate.Mux2Way(a.F.Bin(), gate.AndUint16(x, y), Adder16(x, y))
	return gate.XorUint16(out, snack.Expand16(a.NO))
}
