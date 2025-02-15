package mem

import (
	"github.com/crookdc/snack/internal/gate"
	"github.com/crookdc/snack/internal/pin"
)

// Bit represents a digital snack.Pin that has been stored in a 1-bit register.
type Bit struct {
	dff gate.DFF
}

func (b *Bit) Out(clk pin.Pin, in pin.Pin) pin.Signal {
	b.dff.In.Set(gate.Mux2Way1(clk.Signal(), b.dff.Out(clk.Signal()), in.Signal()))
	return b.dff.Out(clk.Signal())
}

// Register represents a simple array of 16 Bit coupled together to store a single 16 bit value.
type Register struct {
	bits [16]Bit
}

// Out reads the currently stored 16 bit value
func (r *Register) Out(clk pin.Pin, in [16]pin.Pin) [16]pin.Signal {
	return [16]pin.Signal{
		r.bits[0].Out(clk, in[0]),
		r.bits[1].Out(clk, in[1]),
		r.bits[2].Out(clk, in[2]),
		r.bits[3].Out(clk, in[3]),
		r.bits[4].Out(clk, in[4]),
		r.bits[5].Out(clk, in[5]),
		r.bits[6].Out(clk, in[6]),
		r.bits[7].Out(clk, in[7]),
		r.bits[8].Out(clk, in[8]),
		r.bits[9].Out(clk, in[9]),
		r.bits[10].Out(clk, in[10]),
		r.bits[11].Out(clk, in[11]),
		r.bits[12].Out(clk, in[12]),
		r.bits[13].Out(clk, in[13]),
		r.bits[14].Out(clk, in[14]),
		r.bits[15].Out(clk, in[15]),
	}
}

type RAM8 struct {
	Registers [8]Register
}

func (r *RAM8) Out(clk pin.Pin, addr [3]pin.Pin, in [16]pin.Pin) [16]pin.Signal {
	al, bl, cl, dl, el, fl, gl, hl := gate.DMux8Way1(
		[3]pin.Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		clk.Signal(),
	)
	return gate.Mux8Way16(
		[3]pin.Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.Registers[0].Out(pin.New(al), in),
		r.Registers[1].Out(pin.New(bl), in),
		r.Registers[2].Out(pin.New(cl), in),
		r.Registers[3].Out(pin.New(dl), in),
		r.Registers[4].Out(pin.New(el), in),
		r.Registers[5].Out(pin.New(fl), in),
		r.Registers[6].Out(pin.New(gl), in),
		r.Registers[7].Out(pin.New(hl), in),
	)
}
