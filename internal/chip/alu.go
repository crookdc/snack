package chip

// ALU (Arithmetic Logic Unit) represents an actual hardware chip that can handle an array of arithmetic operations
// which are executed on two 16-bit unsigned integers and which produces a single 16-bit unsigned integer as result.
type ALU struct {
	// ZX sets all bits of x to 0
	ZX Pin
	// NX negates all bits of x
	NX Pin
	// ZY sets all bits of y to 0
	ZY Pin
	// NY negates all bits of y
	NY Pin
	// F when set causes ALU operator to be a bitwise AND, otherwise operator is addition
	F Pin
	// NO negates all bits of output
	NO Pin
}

// Call performs operations on the provided inputs as outlined by the state of the ALU
func (a *ALU) Call(x, y [16]Signal) [16]Signal {
	x = And16(x, expand16(Not(a.ZX.Signal())))
	x = Xor16(x, expand16(a.NX.Signal()))

	y = And16(y, expand16(Not(a.ZY.Signal())))
	y = Xor16(y, expand16(a.NY.Signal()))

	out := Mux2Way16(a.F.Signal(), And16(x, y), Adder16(x, y))
	return Xor16(out, expand16(a.NO.Signal()))
}
