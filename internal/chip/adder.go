package chip

// HalfAdder accepts two bits as input and produces an output consisting of a sum and a carry resulting from adding the
// provided bits together.
func HalfAdder(a, b Signal) (carry Signal, sum Signal) {
	return And(a, b), Xor(a, b)
}

// FullAdder accepts three bits and produces a carry and a sum bit representing the result of adding the three bits
// together.
func FullAdder(a, b, c Signal) (carry Signal, sum Signal) {
	ac, sum := HalfAdder(a, b)
	bc, sum := HalfAdder(sum, c)
	return Or(ac, bc), sum
}

// Adder16 adds two 16-bit integers and returns the result. The carry bit is ignored by the adder.
func Adder16(a, b [16]Signal) [16]Signal {
	r := [16]Signal{}
	c, s := Inactive, Inactive
	for i := range 16 {
		c, s = FullAdder(a[15-i], b[15-i], c)
		r[15-i] = s
	}
	return r
}
