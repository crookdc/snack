package pin

type Signal uint8

const (
	Inactive Signal = iota
	Active
)

type Pin struct {
	n Signal
}

// Split16 transforms a 16-bit integer to its bit representation in the Signal abstraction.
func Split16(n uint16) []Signal {
	res := make([]Signal, 16)
	for i := range 16 {
		res[i] = Signal(uint8(n>>(15-i)) & 1)
	}
	return res
}

// Join16 transforms a Signal slice of length 16 to a 16-bit integer
func Join16(n [16]Signal) uint16 {
	res := uint16(0)
	for i := range 16 {
		res = res | (uint16(n[i]) << (15 - i))
	}
	return res
}

// Expand16 takes a single bit and expands its value to cover 16 bits. That is, if the bit value is 0 then a 16-bit
// unsigned integer containing all zeroes is returned. If the input bit value is 1 then an unsigned 16-bit integer
// containing all ones is returned.
func Expand16(n Signal) uint16 {
	if n == Active {
		return 0xFFFF
	}
	return 0
}

func New(s Signal) Pin {
	return Pin{n: s}
}

func (b *Pin) Mask() uint8 {
	if b.Active() {
		return 0xFF
	}
	return 0
}

func (b *Pin) Set(s Signal) {
	b.n = s
}

func (b *Pin) Activate() {
	b.n = Active
}

func (b *Pin) Deactivate() {
	b.n = Inactive
}

func (b *Pin) Flip() {
	if b.n == Inactive {
		b.n = Active
	} else {
		b.n = Inactive
	}
}

func (b *Pin) Active() bool {
	return b.n == Active
}

func (b *Pin) Signal() Signal {
	return b.n
}
