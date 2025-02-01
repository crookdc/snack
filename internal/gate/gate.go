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

func Or(a, b uint8) uint8 {
	return Not(NotAnd(a, b))
}

func Xor(a, b uint8) uint8 {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func splitUint16(n uint16) (msb uint8, lsb uint8) {
	msb = uint8(n >> 8)
	lsb = uint8(n)
	return msb, lsb
}

func joinUint16(msb, lsb uint8) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}
