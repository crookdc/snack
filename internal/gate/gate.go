// Package gate provides implementations of basic logical gates in three different flavours which accept gate uint8 and
// uint16 parameters respectively. All gates are built on top of the NotAnd gate.
package gate

import (
	"github.com/crookdc/snack/internal/pin"
)

func NotAnd(a, b pin.Signal) pin.Signal {
	if a == b && a == pin.Inactive {
		return pin.Active
	}
	return pin.Inactive
}

func NotAndUint16(a, b uint16) uint16 {
	as := pin.Split16(a)
	bs := pin.Split16(b)
	return pin.Join16([16]pin.Signal{
		NotAnd(as[0], bs[0]),
		NotAnd(as[1], bs[1]),
		NotAnd(as[2], bs[2]),
		NotAnd(as[3], bs[3]),
		NotAnd(as[4], bs[4]),
		NotAnd(as[5], bs[5]),
		NotAnd(as[6], bs[6]),
		NotAnd(as[7], bs[7]),
		NotAnd(as[8], bs[8]),
		NotAnd(as[9], bs[9]),
		NotAnd(as[10], bs[10]),
		NotAnd(as[11], bs[11]),
		NotAnd(as[12], bs[12]),
		NotAnd(as[13], bs[13]),
		NotAnd(as[14], bs[14]),
		NotAnd(as[15], bs[15]),
	})
}

func Not(a pin.Signal) pin.Signal {
	return NotAnd(a, pin.Inactive)
}

func NotUint16(a uint16) uint16 {
	as := pin.Split16(a)
	return pin.Join16([16]pin.Signal{
		Not(as[0]),
		Not(as[1]),
		Not(as[2]),
		Not(as[3]),
		Not(as[4]),
		Not(as[5]),
		Not(as[6]),
		Not(as[7]),
		Not(as[8]),
		Not(as[9]),
		Not(as[10]),
		Not(as[11]),
		Not(as[12]),
		Not(as[13]),
		Not(as[14]),
		Not(as[15]),
	})
}

func And(a, b pin.Signal) pin.Signal {
	return NotAnd(Not(a), Not(b))
}

func AndUint16(a, b uint16) uint16 {
	as := pin.Split16(a)
	bs := pin.Split16(b)
	return pin.Join16([16]pin.Signal{
		And(as[0], bs[0]),
		And(as[1], bs[1]),
		And(as[2], bs[2]),
		And(as[3], bs[3]),
		And(as[4], bs[4]),
		And(as[5], bs[5]),
		And(as[6], bs[6]),
		And(as[7], bs[7]),
		And(as[8], bs[8]),
		And(as[9], bs[9]),
		And(as[10], bs[10]),
		And(as[11], bs[11]),
		And(as[12], bs[12]),
		And(as[13], bs[13]),
		And(as[14], bs[14]),
		And(as[15], bs[15]),
	})
}

func Or(a, b pin.Signal) pin.Signal {
	return Not(NotAnd(a, b))
}

func OrUint16(a, b uint16) uint16 {
	as := pin.Split16(a)
	bs := pin.Split16(b)
	return pin.Join16([16]pin.Signal{
		Or(as[0], bs[0]),
		Or(as[1], bs[1]),
		Or(as[2], bs[2]),
		Or(as[3], bs[3]),
		Or(as[4], bs[4]),
		Or(as[5], bs[5]),
		Or(as[6], bs[6]),
		Or(as[7], bs[7]),
		Or(as[8], bs[8]),
		Or(as[9], bs[9]),
		Or(as[10], bs[10]),
		Or(as[11], bs[11]),
		Or(as[12], bs[12]),
		Or(as[13], bs[13]),
		Or(as[14], bs[14]),
		Or(as[15], bs[15]),
	})
}

func Xor(a, b pin.Signal) pin.Signal {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func XorUint16(a, b uint16) uint16 {
	as := pin.Split16(a)
	bs := pin.Split16(b)
	return pin.Join16([16]pin.Signal{
		Xor(as[0], bs[0]),
		Xor(as[1], bs[1]),
		Xor(as[2], bs[2]),
		Xor(as[3], bs[3]),
		Xor(as[4], bs[4]),
		Xor(as[5], bs[5]),
		Xor(as[6], bs[6]),
		Xor(as[7], bs[7]),
		Xor(as[8], bs[8]),
		Xor(as[9], bs[9]),
		Xor(as[10], bs[10]),
		Xor(as[11], bs[11]),
		Xor(as[12], bs[12]),
		Xor(as[13], bs[13]),
		Xor(as[14], bs[14]),
		Xor(as[15], bs[15]),
	})
}

// Mux2Way16 provides a multiplexer for 2 inputs and a selector. This variant of the multiplexer supports
// only binary values (0, 1) to be passed in as selector, any non-zero value is considered set (0xFF) and
// only zero is considered unset (0x00). The multiplexer will return the value of `a` if `s` is unset (0)
// and the value of `b` is `s` is set (> 0).
func Mux2Way16(s pin.Signal, a, b uint16) uint16 {
	return OrUint16(AndUint16(NotUint16(pin.Expand16(s)), a), AndUint16(pin.Expand16(s), b))
}

// Mux2WaySig provides a multiplexer to two single bit inputs. More information on the multiplexer is given in the
// Mux2Way16 function comment.
func Mux2WaySig(s pin.Signal, a, b pin.Signal) pin.Signal {
	return Or(And(Not(s), a), And(s, b))
}

// Mux4Way16 provides a multiplexer for 4 inputs and a selector consisting of 2 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux4Way16(s [2]pin.Signal, a, b, c, d uint16) uint16 {
	ab := Mux2Way16(s[1], a, b)
	cd := Mux2Way16(s[1], c, d)
	return Mux2Way16(s[0], ab, cd)
}

// Mux8Way16 provides a multiplexer for 8 inputs and a selector consisting of 3 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux8Way16(s [3]pin.Signal, a, b, c, d, e, f, g, h uint16) uint16 {
	abcd := Mux4Way16([2]pin.Signal{s[1], s[2]}, a, b, c, d)
	efgh := Mux4Way16([2]pin.Signal{s[1], s[2]}, e, f, g, h)
	return Mux2Way16(s[0], abcd, efgh)
}

// Demux2Way16 provides a demultiplexer for an uint16 and a binary selector represented by a single byte.
// Non-zero values on the selector byte is considered set, and only a value of zero is considered unset.
func Demux2Way16(s pin.Signal, in uint16) (a uint16, b uint16) {
	a = AndUint16(NotUint16(pin.Expand16(s)), in)
	b = AndUint16(pin.Expand16(s), in)
	return a, b
}

// Demux4Way16 provides a demultiplexer for 4 outputs based on a 2-byte selector.
func Demux4Way16(s [2]pin.Signal, in uint16) (a uint16, b uint16, c uint16, d uint16) {
	sh := pin.Expand16(s[0])
	a, b = Demux2Way16(s[1], in)
	a = AndUint16(NotUint16(sh), a)
	b = AndUint16(NotUint16(sh), b)
	c, d = Demux2Way16(s[1], in)
	c = AndUint16(sh, c)
	d = AndUint16(sh, d)
	return a, b, c, d
}

// Demux8Way16 provides a demultiplexer for 8 outputs based on a 3-byte selector.
func Demux8Way16(s [3]pin.Signal, in uint16) (a, b, c, d, e, f, g, h uint16) {
	sh := pin.Expand16(s[0])
	a, b, c, d = Demux4Way16([2]pin.Signal{s[1], s[2]}, in)
	a = AndUint16(NotUint16(sh), a)
	b = AndUint16(NotUint16(sh), b)
	c = AndUint16(NotUint16(sh), c)
	d = AndUint16(NotUint16(sh), d)
	e, f, g, h = Demux4Way16([2]pin.Signal{s[1], s[2]}, in)
	e = AndUint16(sh, e)
	f = AndUint16(sh, f)
	g = AndUint16(sh, g)
	h = AndUint16(sh, h)
	return a, b, c, d, e, f, g, h
}

// DFF represents a data flip-flop capable of holding a single bit of information across CPU cycles.
type DFF struct {
	In  pin.Pin
	out pin.Pin
}

func (d *DFF) Out(clk pin.Signal) pin.Signal {
	if clk == pin.Active {
		d.out = d.In
	}
	return d.out.Signal()
}
