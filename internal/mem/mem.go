package mem

import (
	"github.com/crookdc/snack/internal/gate"
	"github.com/crookdc/snack/internal/pin"
)

// Bit represents a digital snack.Pin that has been stored in a 1-bit register.
type Bit struct {
	dff gate.DFF
	In  pin.Pin
}

func (b *Bit) Out(clk pin.Pin) pin.Signal {
	b.dff.In.Set(gate.Mux2WaySig(clk.Signal(), b.dff.Out(clk.Signal()), b.In.Signal()))
	return b.dff.Out(clk.Signal())
}

// Register represents a simple array of 16 Bit coupled together to store a single 16 bit value.
type Register [16]Bit

// Out reads the currently stored 16 bit value
func (r *Register) Out(clk pin.Pin) [16]pin.Signal {
	return [16]pin.Signal{
		r[0].Out(clk),
		r[1].Out(clk),
		r[2].Out(clk),
		r[3].Out(clk),
		r[4].Out(clk),
		r[5].Out(clk),
		r[6].Out(clk),
		r[7].Out(clk),
		r[8].Out(clk),
		r[9].Out(clk),
		r[10].Out(clk),
		r[11].Out(clk),
		r[12].Out(clk),
		r[13].Out(clk),
		r[14].Out(clk),
		r[15].Out(clk),
	}
}

// Set puts the provided 16 bit value into the input to be committed upon the next clock cycle
func (r *Register) Set(s [16]pin.Signal) {
	r[0].In.Set(s[0])
	r[1].In.Set(s[1])
	r[2].In.Set(s[2])
	r[3].In.Set(s[3])
	r[4].In.Set(s[4])
	r[5].In.Set(s[5])
	r[6].In.Set(s[6])
	r[7].In.Set(s[7])
	r[8].In.Set(s[8])
	r[9].In.Set(s[9])
	r[10].In.Set(s[10])
	r[11].In.Set(s[11])
	r[12].In.Set(s[12])
	r[13].In.Set(s[13])
	r[14].In.Set(s[14])
	r[15].In.Set(s[15])
}
