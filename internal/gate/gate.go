// Package gate provides implementations of basic logical gates in three different flavours which accept gate uint8 and
// uint16 parameters respectively. All gates are built on top of the NotAnd gate.
package gate

func NotAnd(a, b uint8) uint8 {
	res := uint8(0)
	for i := range 8 {
		if (a>>i)&1 == 0 && (b>>i)&1 == 0 {
			res = res | 1<<i
		}
	}
	return res
}

func NotAndUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(NotAnd(am, bm), NotAnd(al, bl))
}

func Not(a uint8) uint8 {
	return NotAnd(a, 0)
}

func NotUint16(a uint16) uint16 {
	msb, lsb := splitUint16(a)
	return joinUint16(Not(msb), Not(lsb))
}

func And(a, b uint8) uint8 {
	return NotAnd(Not(a), Not(b))
}

func AndUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(And(am, bm), And(al, bl))
}

func Or(a, b uint8) uint8 {
	return Not(NotAnd(a, b))
}

func OrUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(Or(am, bm), Or(al, bl))
}

func Xor(a, b uint8) uint8 {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func XorUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(Xor(am, bm), Xor(al, bl))
}

// Mux2Way provides a multiplexer for 2 inputs and a selector. This variant of the multiplexer supports
// only binary values (0, 1) to be passed in, otherwise the returned value is for all intents considered
// undefined. This hard requirements is made clear by the `panic` calls made for any other type of value.
// Given that the supplied input is correct, this multiplexer will return the value of `a` is `sel` is
// unset (0) and the value of `b` is `sel` is set (1).
func Mux2Way(sel, a, b uint8) uint8 {
	if !binary(sel) || !binary(a) || !binary(b) {
		panic("passed non-binary input signal to binary multiplexer")
	}
	return Or(And(Not(sel), a), And(sel, b))
}

// Demux2Way provides the opposite function of Mux2Way, that is, it accepts a selection and an input
// signal and produces two values, one of which will always be 0 and the other will take the value of
// input. When `sel` is unset then the first output is given the value of `in`, and when `sel` is set then
// the second output is given the value of `in`.
func Demux2Way(sel, in uint8) (uint8, uint8) {
	if !binary(sel) || !binary(sel) {
		panic("passed non-binary input signal to binary multiplexer")
	}
	return And(in, Not(sel)), And(in, sel)
}

func binary(n uint8) bool {
	return n == 0 || n == 1
}

func splitUint16(n uint16) (msb uint8, lsb uint8) {
	msb = uint8(n >> 8)
	lsb = uint8(n)
	return msb, lsb
}

func joinUint16(msb, lsb uint8) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}
