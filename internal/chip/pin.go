package chip

type Signal uint8

const (
	Inactive Signal = iota
	Active
)

type Pin struct {
	n Signal
}

// split16 transforms a 16-bit integer to its bit representation in the Signal abstraction.
func split16(n uint16) [16]Signal {
	res := [16]Signal{}
	for i := range 16 {
		res[i] = Signal(uint8(n>>(15-i)) & 1)
	}
	return res
}

func NewPin16(n [16]Signal) [16]Pin {
	p := [16]Pin{}
	for i := range 16 {
		p[i] = NewPin(n[i])
	}
	return p
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

func NewPin(s Signal) Pin {
	return Pin{n: s}
}

func (p *Pin) Mask() uint8 {
	if p.Active() {
		return 0xFF
	}
	return 0
}

func (p *Pin) Set(s Signal) {
	p.n = s
}

func (p *Pin) Activate() {
	p.Set(Active)
}

func (p *Pin) Deactivate() {
	p.Set(Inactive)
}

func (p *Pin) Flip() {
	if p.n == Inactive {
		p.Set(Active)
	} else {
		p.Set(Inactive)
	}
}

func (p *Pin) Active() bool {
	return p.n == Active
}

func (p *Pin) Signal() Signal {
	return p.n
}
