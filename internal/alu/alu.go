package alu

import (
	"github.com/crookdc/snack/internal/gate"
	"github.com/crookdc/snack/internal/pin"
)

// ALU (Arithmetic Logic Unit) represents an actual hardware chip that can handle an array of arithmetic operations
// which are executed on two 16-bit unsigned integers and which produces a single 16-bit unsigned integer as result.
type ALU struct {
	// ZX sets all bits of x to 0
	ZX pin.Pin
	// NX negates all bits of x
	NX pin.Pin
	// ZY sets all bits of y to 0
	ZY pin.Pin
	// NY negates all bits of y
	NY pin.Pin
	// F when set causes ALU operator to be a bitwise AND, otherwise operator is addition
	F pin.Pin
	// NO negates all bits of output
	NO pin.Pin
}

// Call performs operations on the provided inputs as outlined by the state of the ALU
func (a *ALU) Call(x, y [16]pin.Signal) [16]pin.Signal {
	x = gate.And16(x, pin.Expand16(gate.Not(a.ZX.Signal())))
	x = gate.Xor16(x, pin.Expand16(a.NX.Signal()))

	y = gate.And16(y, pin.Expand16(gate.Not(a.ZY.Signal())))
	y = gate.Xor16(y, pin.Expand16(a.NY.Signal()))

	out := gate.Mux2Way16(a.F.Signal(), gate.And16(x, y), Adder16(x, y))
	return gate.Xor16(out, pin.Expand16(a.NO.Signal()))
}
