package chip

type Signal uint8

const (
	Inactive Signal = iota
	Active
)

// split16 transforms a 16-bit integer to its bit representation in the Signal abstraction.
func split16(n uint16) [16]Signal {
	res := [16]Signal{}
	for i := range 16 {
		res[i] = Signal(uint8(n>>(15-i)) & 1)
	}
	return res
}

// split15 transforms a 16-bit integer to its 15-bit representation in the Signal abstraction by omitting the MSB.
func split15(n uint16) [15]Signal {
	res := [15]Signal{}
	for i := range 15 {
		res[i] = Signal(uint8(n>>(14-i)) & 1)
	}
	return res
}

// expand16 takes a single bit and expands its value to cover 16 bits. That is, if the bit value is 0 then a 16-bit
// unsigned integer containing all zeroes is returned. If the input bit value is 1 then an unsigned 16-bit integer
// containing all ones is returned.
func expand16(n Signal) [16]Signal {
	return [16]Signal{
		n, n, n, n, n, n, n, n,
		n, n, n, n, n, n, n, n,
	}
}

func Join15(sigs [15]Signal) uint16 {
	n := uint16(0)
	for i, sig := range sigs {
		n = n | (uint16(sig) << (14 - i))
	}
	return n
}
