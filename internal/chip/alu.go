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
func (a *ALU) Out(x, y ReadonlyWord) (out *Word, zr Signal, ng Signal) {
	x = And16To1(Not(a.ZX), x)
	x = Xor16To1(a.NX, x)

	y = And16To1(Not(a.ZY), y)
	y = Xor16To1(a.NY, y)

	out = Mux2Way16(a.F, And16(x, y), Adder16(x, y))
	out = Xor16To1(a.NO, out)
	ng = out.Get(0) // If the MSB is 1 then the value is negative as per the rules of two's complement
	zr = Not(Or(out.Get(0), Or(out.Get(1), Or(out.Get(2), Or(out.Get(3), Or(out.Get(4), Or(out.Get(5), Or(out.Get(6), Or(out.Get(7), Or(out.Get(8), Or(out.Get(9), Or(out.Get(10), Or(out.Get(11), Or(out.Get(12), Or(out.Get(13), Or(out.Get(14), out.Get(15)))))))))))))))))
	return
}
