package mem

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

// Bit represents a digital snack.Signal that has been stored in a 1-bit register.
type Bit struct {
	dff gate.DFF
	In  snack.Signal
}

func (b *Bit) Out(clk snack.Signal) snack.Signal {
	b.dff.In = gate.Mux2WayBit(clk.Bin(), b.dff.Out(clk), b.In)
	return b.dff.Out(clk)
}
