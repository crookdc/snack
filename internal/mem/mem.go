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

// Register represents a simple array of 16 Bit coupled together to store a single uint16. Even though the underlying
// implementation is represented using a [16]Bit array, the API still works with actual uint16 values that it splits or
// joins depending on the operation.
type Register [16]Bit

// Out reads the currently stored uint16 value
func (r *Register) Out(clk pin.Pin) uint16 {
	sigs := [16]pin.Signal{}
	for i := range r {
		sigs[i] = r[i].Out(clk)
	}
	return pin.Join16(sigs)
}

// Set puts the provided uint16 into the input to be committed upon the next clock cycle
func (r *Register) Set(v uint16) {
	sigs := pin.Split16(v)
	for i := range sigs {
		r[i].In.Set(sigs[i])
	}
}
