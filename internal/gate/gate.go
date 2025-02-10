// Package gate provides implementations of basic logical gates in three different flavours which accept gate uint8 and
// uint16 parameters respectively. All gates are built on top of the NotAnd gate.
package gate

import (
	"github.com/crookdc/snack"
)

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

func NotBit(a snack.Bit) snack.Bit {
	return snack.NewBit(Not(a.Bin()) & 1)
}

func And(a, b uint8) uint8 {
	return NotAnd(Not(a), Not(b))
}

func AndUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(And(am, bm), And(al, bl))
}

func AndBit(a, b snack.Bit) snack.Bit {
	return snack.NewBit(And(a.Bin(), b.Bin()))
}

func Or(a, b uint8) uint8 {
	return Not(NotAnd(a, b))
}

func OrUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(Or(am, bm), Or(al, bl))
}

func OrBit(a, b snack.Bit) snack.Bit {
	return snack.NewBit(Or(a.Bin(), b.Bin()))
}

func Xor(a, b uint8) uint8 {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func XorUint16(a, b uint16) uint16 {
	am, al := splitUint16(a)
	bm, bl := splitUint16(b)
	return joinUint16(Xor(am, bm), Xor(al, bl))
}

func XorBit(a, b snack.Bit) snack.Bit {
	return snack.NewBit(Xor(a.Bin(), b.Bin()))
}

// Mux2Way provides a multiplexer for 2 inputs and a selector. This variant of the multiplexer supports
// only binary values (0, 1) to be passed in as selector, any non-zero a is considered set (0xFF) and
// only zero is considered unset (0x00). The multiplexer will return the a of `a` if `s` is unset (0)
// and the a of `b` is `s` is set (> 0).
func Mux2Way(s uint8, a, b uint16) uint16 {
	s = selector(s)
	return OrUint16(AndUint16(NotUint16(uint16(s)|uint16(s)<<8), a), AndUint16(uint16(s)|uint16(s)<<8, b))
}

// Mux2WayBit provides a multiplexer to two single bit inputs. More information on the multiplexer is given in the
// Mux2Way function comment.
func Mux2WayBit(s uint8, a, b snack.Bit) snack.Bit {
	s = selector(s)
	return OrBit(AndBit(NotBit(snack.NewBit(s&1)), a), AndBit(snack.NewBit(s&1), b))
}

// Mux4Way provides a multiplexer for 4 inputs and a selector consisting of 2 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux4Way(s [2]uint8, a, b, c, d uint16) uint16 {
	ab := Mux2Way(s[1], a, b)
	cd := Mux2Way(s[1], c, d)
	return Mux2Way(s[0], ab, cd)
}

// Mux8Way provides a multiplexer for 8 inputs and a selector consisting of 3 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux8Way(s [3]uint8, a, b, c, d, e, f, g, h uint16) uint16 {
	abcd := Mux4Way([2]uint8{s[1], s[2]}, a, b, c, d)
	efgh := Mux4Way([2]uint8{s[1], s[2]}, e, f, g, h)
	return Mux2Way(s[0], abcd, efgh)
}

// Demux2Way provides a demultiplexer for an uint16 and a binary selector represented by a single byte.
// Non-zero values on the selector byte is considered set, and only a a of zero is considered unset.
func Demux2Way(s uint8, in uint16) (a uint16, b uint16) {
	s = selector(s)
	a = AndUint16(NotUint16(uint16(s)|uint16(s)<<8), in)
	b = AndUint16(uint16(s)|uint16(s)<<8, in)
	return a, b
}

// Demux2WayBit provides a demultiplexer for a 2-way selector, producing snack.Bit values.
func Demux2WayBit(s uint8, in snack.Bit) (a snack.Bit, b snack.Bit) {
	s = selector(s)
	a = AndBit(NotBit(snack.NewBit(s&1)), in)
	b = AndBit(snack.NewBit(s&1), in)
	return a, b
}

// Demux4Way provides a demultiplexer for 4 outputs based on a 2-byte selector.
func Demux4Way(s [2]uint8, in uint16) (a uint16, b uint16, c uint16, d uint16) {
	sh := expand(selector(s[0]))
	a, b = Demux2Way(s[1], in)
	a = AndUint16(NotUint16(sh), a)
	b = AndUint16(NotUint16(sh), b)
	c, d = Demux2Way(s[1], in)
	c = AndUint16(sh, c)
	d = AndUint16(sh, d)
	return a, b, c, d
}

// Demux8Way provides a demultiplexer for 8 outputs based on a 3-byte selector.
func Demux8Way(s [3]uint8, in uint16) (a, b, c, d, e, f, g, h uint16) {
	sh := expand(selector(s[0]))
	a, b, c, d = Demux4Way([2]uint8{s[1], s[2]}, in)
	a = AndUint16(NotUint16(sh), a)
	b = AndUint16(NotUint16(sh), b)
	c = AndUint16(NotUint16(sh), c)
	d = AndUint16(NotUint16(sh), d)
	e, f, g, h = Demux4Way([2]uint8{s[1], s[2]}, in)
	e = AndUint16(sh, e)
	f = AndUint16(sh, f)
	g = AndUint16(sh, g)
	h = AndUint16(sh, h)
	return a, b, c, d, e, f, g, h
}

func expand(n uint8) uint16 {
	return uint16(n) | uint16(n)<<8
}

func selector(n uint8) uint8 {
	if n > 0 {
		return 0xFF
	}
	return 0
}

func splitUint16(n uint16) (msb uint8, lsb uint8) {
	msb = uint8(n >> 8)
	lsb = uint8(n)
	return msb, lsb
}

func joinUint16(msb, lsb uint8) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}

func NewDFF() *DFF {
	return &DFF{
		a: snack.UnsetBit(),
		b: snack.UnsetBit(),
	}
}

type Clock interface {
	Tick() snack.Bit
}

// DFF represents a data flip-flop capable of holding a single bit of information across CPU cycles.
type DFF struct {
	l snack.Bit

	a snack.Bit
	b snack.Bit
}

func (d *DFF) IsSet() bool {
	bit := Mux2WayBit(d.l.Bin(), d.a, d.b)
	return bit.IsSet()
}

func (d *DFF) Get() snack.Bit {
	return Mux2WayBit(d.l.Bin(), d.a, d.b)
}

func (d *DFF) Set(v snack.Bit) {
	d.a = Mux2WayBit(d.l.Bin(), d.a, v)
	d.b = Mux2WayBit(d.l.Bin(), v, d.b)
}

func (d *DFF) Flip() {
	if d.l.IsSet() {
		d.l = snack.UnsetBit()
	} else {
		d.l = snack.SetBit()
	}
}
