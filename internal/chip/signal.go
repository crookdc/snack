package chip

type Signal uint8

const (
	Inactive Signal = iota
	Active
)

var NullWord = Wrap(&[16]Signal{})

type ReadonlyWord interface {
	Get(position int) Signal
	Copy() [16]Signal
}

func Wrap(w *[16]Signal) *Word {
	return &Word{word: w}
}

func NewWord() *Word {
	return &Word{word: &[16]Signal{}}
}

type Word struct {
	word *[16]Signal
}

func (w *Word) Get(position int) Signal {
	return w.word[position]
}

func (w *Word) Copy() [16]Signal {
	word := [16]Signal{}
	for i := range 16 {
		word[i] = w.word[i]
	}
	return word
}

func (w *Word) Set(position int, value Signal) {
	w.word[position] = value
}

func (w *Word) Address() [15]Signal {
	address := [15]Signal{}
	for i := 1; i < 16; i++ {
		address[i-1] = w.word[i]
	}
	return address
}

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
