package chip

// ALU (Arithmetic Logic Unit) represents an actual hardware chip that can handle an array of arithmetic operations
// which are executed on two 16-bit unsigned integers and which produces a single 16-bit unsigned integer as result.
type ALU struct {
	// ZX sets all bits of x to 0
	ZX Signal
	// NX negates all bits of x
	NX Signal
	// ZY sets all bits of y to 0
	ZY Signal
	// NY negates all bits of y
	NY Signal
	// F when set causes ALU operator to be a bitwise AND, otherwise operator is addition
	F Signal
	// NO negates all bits of output
	NO Signal
}

// Out is performing operations on the provided inputs as outlined by the state of the ALU
func (a *ALU) Out(x, y [16]Signal) (out [16]Signal, zr Signal, ng Signal) {
	x = And16(x, expand16(Not(a.ZX)))
	x = Xor16(x, expand16(a.NX))

	y = And16(y, expand16(Not(a.ZY)))
	y = Xor16(y, expand16(a.NY))

	out = Mux2Way16(a.F, And16(x, y), Adder16(x, y))
	out = Xor16(out, expand16(a.NO))
	ng = out[0] // If the MSB is 1 then the value is negative as per the rules of two's complement
	zr = Not(Or(out[0], Or(out[1], Or(out[2], Or(out[3], Or(out[4], Or(out[5], Or(out[6], Or(out[7], Or(out[8], Or(out[9], Or(out[10], Or(out[11], Or(out[12], Or(out[13], Or(out[14], out[15]))))))))))))))))
	return
}
